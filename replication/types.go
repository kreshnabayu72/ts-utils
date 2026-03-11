package replication

import (
	"context"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/config"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/replication/api"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/service"
)

// CDCKind represents the type of database change captured by Change Data Capture.
type CDCKind string

const (
	CDCKindInsert CDCKind = "insert" // A new row was inserted
	CDCKindUpdate CDCKind = "update" // An existing row was modified
	CDCKindDelete CDCKind = "delete" // A row was deleted

	// CDCCustomKindMove is a synthetic CDC kind, NOT produced by PeerDB directly.
	// It is derived internally by the processor when it detects that an "update" event
	// has changed the `client_id` field between `msg.Old` and `msg.New`, meaning the
	// record has been reassigned from one tenant to another.
	//
	// REQUIREMENT: Tables that can trigger a Move event MUST have:
	//
	//	ALTER TABLE IF EXISTS <name> REPLICA IDENTITY FULL;
	//
	// Without this, PostgreSQL omits `OLD` row data from WAL `UPDATE` events.
	// PeerDB will then send `msg.Old = {}`, and the processor cannot determine
	// which old tenant schema to delete from - causing ghost records.
	//
	// Tables that qualify: any table with `ReplicationFlowNonTenantToTenant` AND a `client_id` column.
	// See README.md in this package for the full explanation.
	CDCCustomKindMove CDCKind = "move" // A row was moved from one tenant schema to another
)

// ReplicationFlow is a type alias for api.ReplicationFlow
type ReplicationFlow = api.ReplicationFlow

const (
	ReplicationFlowTenantToTenant       ReplicationFlow = api.ReplicationFlowTenantToTenant
	ReplicationFlowNonTenantToTenant    ReplicationFlow = api.ReplicationFlowNonTenantToTenant
	ReplicationFlowNonTenantToNonTenant ReplicationFlow = api.ReplicationFlowNonTenantToNonTenant
	ReplicationFlowTenantToNonTenant    ReplicationFlow = api.ReplicationFlowTenantToNonTenant
)

// CDCMessage represents a Change Data Capture message produced by PeerDB.
// PeerDB captures PostgreSQL WAL (Write-Ahead Log) changes and publishes them to Kafka.
//
// Example message from topic "tagsamurai_iam.users":
//
//	{
//	    "kind": "insert",
//	    "new": {
//	        "id": 2029086878094856200,
//	        "email": "tyler@example.com",
//	        "full_name": "Tyler Lang"
//	    },
//	    "source": "qtera.users"
//	}
//
// Fields:
//   - Kind: The type of change (insert, update, delete)
//   - New:  The new row data (present for insert and update)
//   - Old:  The old row data (present for delete; contains the row before deletion)
//   - Source: The schema-qualified table name in the source database (e.g., "qtera.users")
//     The source schema part ("qtera" or "public") is used in combination with ReplicationFlow
//     to determine the actual target schema.
type CDCMessage struct {
	Kind   CDCKind        `json:"kind"`
	New    map[string]any `json:"new,omitempty"`
	Old    map[string]any `json:"old,omitempty"`
	Source string         `json:"source"`
}

// ReplicationConfig holds configuration for a single CDC replication instance.
// One CDCReplicator gets one ReplicationConfig.
//
// See README.md in this package for guidance on which ReplicationFlow to use and
// when REPLICA IDENTITY FULL is required on the source tables.
type ReplicationConfig struct {
	SourceDBName    string                        // Source database name, used to build Kafka topic names (e.g., "tagsamurai_iam")
	TableMappings   map[string]TargetTableMapping // Maps source table name (e.g., "users") -> how to replicate it
	ConsumerGroup   string                        // Kafka consumer group ID for offset tracking
	ValidClients    []string                      // List of valid clients for this table.
	ReplicationFlow api.ReplicationFlow
}

// TargetTableMapping describes how to replicate a single table.
// It defines the target table structure and optional transformation logic.
type TargetTableMapping struct {
	PrimaryKey      string                              // Primary key column name. Defaults to "id" if empty.
	Columns         []string                            // Columns to replicate. Only these columns are extracted from CDC messages.
	ColumnTypes     map[string]string                   // Maps column name -> SQL type (e.g., "id" -> "BIGINT"). Used for table creation.
	TargetTableName string                              // The table name in the target database to write to.
	Transform       func(map[string]any) map[string]any // Optional: transforms the CDC data before writing. If nil, data is used as-is.
}

// CDCReplicator consumes CDC messages from Kafka and replicates them into PostgreSQL.
type CDCReplicator struct {
	ctx               context.Context          // Context for managing lifecycle; cancelled on Stop()
	cancel            context.CancelFunc       // Cancel function to stop the replicator
	pool              service.PgxPoolInterface // Postgres connection pool for the TARGET database
	brokers           string                   // Kafka brokers list (e.g., "redpanda:9092"). Comma-separated for multiple.
	replicationConfig *ReplicationConfig       // Configuration: which topics to consume, how to map tables
	configOptions     []config.ConfigOptions   // Optional configuration options
	sem               chan struct{}            // Semaphore to limit concurrent DB operations
}

type CDCSchemaPair struct {
	oldSchema string
	newSchema string
}
