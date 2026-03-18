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
	// PeerDB sets source to "{schema}.{table}" (e.g., "qtera.users" or "public.device_names").
	// We use the source schema name with the ReplicationFlow to determine
	// the actual target schema and set search_path before executing queries.
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
	case api.ReplicationFlowTenantToTenant: // Target is same as source (tenant to tenant)
		if !isValidClients(tenantIdSource) {
			return nil
		}
		// Ensure both oldSchema and newSchema assigned for upsert & delete
		targetSchemas = []CDCSchemaPair{{oldSchema: tenantIdSource, newSchema: tenantIdSource}}

	// case api.ReplicationFlowNonTenantToNonTenant: // Source should be "public" since this is `from non-tenant` flow
	// 	if tenantIdSource != "public" {
	// 		return nil
	// 	}
	// 	// Ensure both oldSchema and newSchema assigned for upsert & delete
	// 	targetSchemas = []CDCSchemaPair{{oldSchema: "public", newSchema: "public"}}

	// case api.ReplicationFlowNonTenantToTenant:
	// 	newData := msg.New
	// 	oldData := msg.Old

	// 	newClientIdVal, newExists := newData["client_id"]
	// 	oldClientIdVal, oldExists := oldData["client_id"]

	// 	if !newExists && !oldExists {
	// 		// Rule: No client_id field AND flow from non-tenant (public) -> Broadcast to all tenants (e.g., device_names & skus)
	// 		if tenantIdSource == "public" {
	// 			targetSchemas = make([]CDCSchemaPair, len(r.replicationConfig.ValidClients))
	// 			for i, v := range r.replicationConfig.ValidClients {
	// 				targetSchemas[i] = CDCSchemaPair{
	// 					oldSchema: v, // To handle deletion for this Rule (e.g., delete from schema "qtera")
	// 					newSchema: v, // To handle insertion for this Rule (e.g., insert into schema "qtera")
	// 				}
	// 			}
	// 		}
	// 	} else {
	// 		// Rule: client_id field exists but empty/null -> Skip this row (skipped by if else conditioning)
	// 		var newClientIdStr, oldClientIdStr string
	// 		if newClientIdVal != nil {
	// 			newClientIdStr = fmt.Sprintf("%v", newClientIdVal)
	// 		}
	// 		if oldClientIdVal != nil {
	// 			oldClientIdStr = fmt.Sprintf("%v", oldClientIdVal)
	// 		}

	// 		// Pure Deletion, no need newSchema.
	// 		// Update with no client_id, it means the record is deleted from target schema.
	// 		if msg.Kind == CDCKindDelete || (msg.Kind == CDCKindUpdate && !isValidClients(newClientIdStr)) {
	// 			if oldExists && !isValidClients(oldClientIdStr) {
	// 				return nil
	// 			}
	// 			targetSchemas = []CDCSchemaPair{{
	// 				oldSchema: oldClientIdStr, // e.g., delete from schema "qtera"
	// 				newSchema: "",
	// 			}}
	// 			msg.Kind = CDCKindDelete
	// 		} else if newExists && oldExists && isValidClients(oldClientIdStr) && isValidClients(newClientIdStr) && newClientIdStr != oldClientIdStr {
	// 			// Move, both field in new & old exist but value differ
	// 			targetSchemas = []CDCSchemaPair{{
	// 				oldSchema: oldClientIdStr, // e.g., delete from schema "0001"
	// 				newSchema: newClientIdStr, // e.g., insert into schema "0002"
	// 			}}
	// 			msg.Kind = CDCCustomKindMove
	// 		} else if newExists && isValidClients(newClientIdStr) {
	// 			// Pure Insertion, no need oldSchema
	// 			targetSchemas = []CDCSchemaPair{{
	// 				oldSchema: "",
	// 				newSchema: newClientIdStr, // e.g., insert into schema "qtera"
	// 			}}
	// 		}
	// 	}

	// case api.ReplicationFlowTenantToNonTenant: // Target is "public"
	// 	if !isValidClients(tenantIdSource) {
	// 		return nil
	// 	}
	// 	// Ensure both oldSchema and newSchema assigned for upsert & delete
	// 	targetSchemas = []CDCSchemaPair{{newSchema: "public", oldSchema: "public"}}

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
