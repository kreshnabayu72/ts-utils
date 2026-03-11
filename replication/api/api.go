// Package api defines the public contracts (interfaces and types) for the CDC replication system.
// These types are shared between ts-utils (the replicator) and ts-migrations (the schema validator).
package api

// ReplicatedTable defines the contract for a table that can be replicated via CDC.
// Each struct that implements this interface represents one database table
// and describes its structure so the replicator knows which columns to read and write.
type ReplicatedTable interface {
	GetTableName() string              // The SQL table name (e.g., "users", "rfids")
	GetColumnTypes() map[string]string // Maps column name -> SQL type (e.g., "id" -> "BIGINT")
	GetColumns() []string              // Ordered list of column names to replicate
}

// CDCReplicatorInterface defines the contract for a CDC replicator.
type CDCReplicatorInterface interface {
	Start() error // Blocks and consumes messages until context is cancelled
	Stop()        // Cancels the context, causing Start() to return
}

// ReplicationFlow describes the directional relationship between source and target schemas.
// It determines how the CDCReplicator routes data from the source database to the target database.
// Choose the correct flow based on whether the source/target schemas are tenant-specific or global (public).
// See ts-utils/replication/README.md for a full explanation with examples.
type ReplicationFlow string

const (
	// ReplicationFlowTenantToTenant: source is tenant schema, target is the same tenant schema.
	// Example: qtera.users (IAM DB) -> qtera.users (hardware-tagging DB)
	ReplicationFlowTenantToTenant ReplicationFlow = "Tenant to Tenant"

	// ReplicationFlowNonTenantToTenant: source is public schema, target is routed to tenant schema(s).
	// There are two sub-cases depending on whether the table has a client_id column:
	//   - WITH client_id: data is routed to the specific tenant schema matching client_id.
	//     These tables MUST have `REPLICA IDENTITY FULL` set to support tenant move detection.
	//     Kind `Move` is manually triggered when CDC send UPDATE value of column client_id has changed to different valid client id.
	//     Example: public.handheld_stocks -> qtera.handheld_stocks (client_id determines target schema)
	//   - WITHOUT client_id: data is broadcast to ALL tenant schemas (ValidClients).
	//     No REPLICA IDENTITY FULL needed since move is never triggered.
	//     Example: public.device_names -> qtera.device_names, acme.device_names, ...
	ReplicationFlowNonTenantToTenant ReplicationFlow = "Non-Tenant to Tenant"

	// ReplicationFlowNonTenantToNonTenant: source is public schema, target is also public schema.
	// Example: public.some_table -> public.some_table
	// Currently supported by the replicator but not actively used.
	ReplicationFlowNonTenantToNonTenant ReplicationFlow = "Non-Tenant to Non-Tenant"

	// ReplicationFlowTenantToNonTenant: source is a tenant schema, target is written to public schema.
	// Example: qtera.some_table -> public.some_table
	// Currently supported by the replicator but not actively used.
	ReplicationFlowTenantToNonTenant ReplicationFlow = "Tenant to Non-Tenant"
)
