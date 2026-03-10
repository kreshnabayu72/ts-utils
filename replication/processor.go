package replication

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/replication/api"
	"github.com/goccy/go-json"
)

// processMessage handles a single CDC message from Kafka.
//
// Flow:
//  1. Parse the Kafka topic name to extract the source table name
//     Example: topic "tagsamurai_iam.users" -> table "users"
//  2. Look up the TargetTableMapping for this table
//  3. Parse the JSON message into a CDCMessage struct
//  4. Extract the source schema from CDCMessage.Source
//     Example: source "qtera.users" -> schema "qtera"
//     Example: source "public.device_names" -> schema "public"
//  5. Determine the actual target schema using the configured ReplicationFlow.
//     For example, ReplicationFlowNonTenantToTenant extracts a client_id from the record
//     to route to a specific schema instead of using the source schema.
//  6. Route to executeReplication (via r.executeReplication) which then routes to
//     upsertRecord (for insert/update) or deleteRecord (for delete)
//     with the target schema, so data is written to the correct location.
func (r *CDCReplicator) processMessage(topic string, value []byte) error {
	// Parse topic to get source table name
	source, err := r.parseTopic(topic)
	if err != nil {
		return err
	}

	tableMap, ok := r.replicationConfig.TableMappings[source]
	if !ok {
		return fmt.Errorf("table mapping for %s not found", source)
	}
	slog.Debug("Table mapping found for %s: %+v", source, tableMap)

	var msg CDCMessage
	if err := json.Unmarshal(value, &msg); err != nil {
		return fmt.Errorf("parse failed: %w", err)
	}

	// Extract schema name from the CDC message source field.
	splitSource := strings.Split(msg.Source, ".")
	var tenantIdSource string
	if len(splitSource) == 2 && splitSource[0] != "" {
		tenantIdSource = splitSource[0] // Tenant schema (e.g., "qtera")
	} else {
		slog.Error("Invalid source format", slog.String("source", msg.Source))
		return fmt.Errorf("invalid source format: %s", msg.Source)
	}

	slog.Debug("Processing CDC message for table:", slog.String("table", source), slog.Any("message", msg))

	isValidClients := func(clientId string) bool {
		return slices.Contains(r.replicationConfig.ValidClients, clientId)
	}

	var targetSchemas []CDCSchemaPair
	switch r.replicationConfig.ReplicationFlow {
	case api.ReplicationFlowTenantToTenant:
		if !isValidClients(tenantIdSource) {
			return nil
		}
		targetSchemas = []CDCSchemaPair{{oldSchema: tenantIdSource, newSchema: tenantIdSource}}

	case api.ReplicationFlowNonTenantToNonTenant:
		if tenantIdSource != "public" {
			return nil
		}
		targetSchemas = []CDCSchemaPair{{oldSchema: "public", newSchema: "public"}}

	case api.ReplicationFlowNonTenantToTenant:
		newData := msg.New
		oldData := msg.Old

		newClientIdVal, newExists := newData["client_id"]
		oldClientIdVal, oldExists := oldData["client_id"]

		if !newExists && !oldExists {
			if tenantIdSource == "public" {
				targetSchemas = make([]CDCSchemaPair, len(r.replicationConfig.ValidClients))
				for i, v := range r.replicationConfig.ValidClients {
					targetSchemas[i] = CDCSchemaPair{oldSchema: v, newSchema: v}
				}
			}
		} else {
			var newClientIdStr, oldClientIdStr string
			if newClientIdVal != nil {
				newClientIdStr = fmt.Sprintf("%v", newClientIdVal)
			}
			if oldClientIdVal != nil {
				oldClientIdStr = fmt.Sprintf("%v", oldClientIdVal)
			}

			if msg.Kind == CDCKindDelete || (msg.Kind == CDCKindUpdate && !isValidClients(newClientIdStr)) {
				if oldExists && !isValidClients(oldClientIdStr) {
					return nil
				}
				targetSchemas = []CDCSchemaPair{{oldSchema: oldClientIdStr, newSchema: ""}}
				msg.Kind = CDCKindDelete
			} else if newExists && oldExists && isValidClients(oldClientIdStr) && isValidClients(newClientIdStr) && newClientIdStr != oldClientIdStr {
				targetSchemas = []CDCSchemaPair{{oldSchema: oldClientIdStr, newSchema: newClientIdStr}}
				msg.Kind = CDCCustomKindMove
			} else if newExists && isValidClients(newClientIdStr) {
				targetSchemas = []CDCSchemaPair{{oldSchema: "", newSchema: newClientIdStr}}
			}
		}

	case api.ReplicationFlowTenantToNonTenant:
		if !isValidClients(tenantIdSource) {
			return nil
		}
		targetSchemas = []CDCSchemaPair{{newSchema: "public", oldSchema: "public"}}

	default:
		slog.Error("Unsupported replication flow", slog.String("flow", string(r.replicationConfig.ReplicationFlow)))
		return nil
	}

	return r.executeReplication(targetSchemas, tableMap, msg)
}

// parseTopic extracts the table name from a Kafka topic name.
//
// Topic format: "{sourceDBName}.{tableName}"
// Example: "tagsamurai_iam.users" -> returns "users"
// Example: "tagsamurai_hardware_tagging.handheld_stocks" -> returns "handheld_stocks"
//
// The returned table name is used to look up the TargetTableMapping.
func (r *CDCReplicator) parseTopic(topic string) (source string, err error) {
	parts := strings.SplitN(topic, ".", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid topic format: %s", topic)
	}

	return parts[1], nil
}
