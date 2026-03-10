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

type ReplicationFlow string

const (
	ReplicationFlowTenantToTenant       ReplicationFlow = "Tenant to Tenant"
	ReplicationFlowNonTenantToTenant    ReplicationFlow = "Non-Tenant to Tenant"
	ReplicationFlowNonTenantToNonTenant ReplicationFlow = "Non-Tenant to Non-Tenant"
	ReplicationFlowTenantToNonTenant    ReplicationFlow = "Tenant to Non-Tenant"
)
