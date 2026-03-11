# Replication Package

This package implements **CDC (Change Data Capture) Replication** powered by [PeerDB](https://www.peerdb.io/) and Kafka.

It listens to WAL (Write-Ahead Log) changes from one PostgreSQL database and replays them into another.

## Architecture Overview

```text
[Source DB] -> [PeerDB] -> [Kafka] -> [CDCReplicator] -> [Target DB]
```

## Key Concepts

### ReplicationFlow

Defines **the directional relationship** between the source and target schemas.

| Flow                   | Description                                                                                                   | Example                                             |
| ---------------------- | ------------------------------------------------------------------------------------------------------------- | --------------------------------------------------- |
| `TenantToTenant`       | Data is in a tenant schema and stays in the same tenant schema                                                | `qtera.users` -> `qtera.users`                      |
| `NonTenantToNonTenant` | Data is in `public` and stays in `public`                                                                     | `public.some_table` -> `public.some_table`          |
| `NonTenantToTenant`    | Data is in `public` but routed to specific tenant schema(s) by `client_id` field                              | `public.handheld_stocks` -> `qtera.handheld_stocks` |
| `TenantToNonTenant`    | Data is in a tenant schema but written to `public` (currently not used, but replicator already supports this) | `qtera.some_table` -> `public.some_table`           |

---

## Why `REPLICA IDENTITY FULL` is Required

> **TL;DR:** Without `REPLICA IDENTITY FULL`, PostgreSQL's WAL does not include the
> **old row data** (`OLD`) in `UPDATE` events. This breaks the `NonTenantToTenant`
> Move detection logic, because we rely on `OLD.client_id` to know which tenant
> schema to delete from.

### The Problem

By default, PostgreSQL's WAL only records:

- **INSERT** -> the full new row ✅
- **UPDATE** -> only the **new** row values ❌ (old values are omitted unless the column is a replica identity)
- **DELETE** -> only the **primary key** from old values ❌ (new values are omitted)

For `NonTenantToTenant` tables that have a `client_id` column, the replicator
uses the **change in `client_id`** between `OLD` and `NEW` to detect that a record
has been moved from one tenant to another (the `CDCCustomKindMove` event).

If `OLD` is empty, the replicator cannot detect the move, and the record will
simply be upserted into the new tenant schema without being deleted from the old one,
causing **ghost records** (duplicate data in the old tenant).

### The Fix

Run `ALTER TABLE IF EXISTS <name> REPLICA IDENTITY FULL` on every table that:

1. Lives in the `public` (non-tenant) source schema, AND
2. Uses `ReplicationFlowNonTenantToTenant`, AND
3. Has a `client_id` column (meaning it can be "moved" between tenants)

This tells PostgreSQL to log the **full old row** in every `UPDATE` and `DELETE`
event, which PeerDB then forwards in the `old` JSON field of the CDC message.

### Which Tables Need It?

Only tables where `client_id` can change across tenants. Tables like `device_names` or `skus` that are broadcast to **all** tenants do NOT need it, because they never trigger a Move event.

**Current tables configured with `REPLICA IDENTITY FULL`** (see migration `20260311030942_alter_table_in_public_schema...`):

- `handheld_stocks`
- `smart_fixed_stocks`
- `tablet_stocks`

---

## The Custom `move` CDC Kind

PeerDB natively produces three CDC kinds: `insert`, `update`, `delete`.

This package introduces a **custom `move` kind** (`CDCCustomKindMove`) which is derived from a standard `update` event where the `client_id` field changes between `OLD` and `NEW`.

### Why a Custom Kind?

The `move` kind allows the executor to handle the operation atomically in a **single transaction**:

```sql
BEGIN;
  SET LOCAL search_path = <old_tenant>;
  DELETE FROM <table> WHERE id = $1;   -- remove from old tenant

  SET LOCAL search_path = <new_tenant>;
  INSERT INTO <table> ... ON CONFLICT DO UPDATE SET ...;  -- add to new tenant
COMMIT;
```

Without this, a naive `update` handler would only upsert into the new tenant
schema, leaving the stale record in the old tenant schema forever.

### Move Detection Logic (in `processor.go`)

```text
msg.Kind == "update"
  AND msg.New["client_id"] != msg.Old["client_id"]
  AND both values are valid (non-empty, in ValidClients list)
  -> promoted to CDCCustomKindMove
```

---

## Adding a New Replicated Table

1. Create the replica model struct in `module/<service>/replica/`.
   - Implement `GetTableName()`, `GetColumns()`, `GetColumnTypes()`.
2. Register it in `registry/registry.go` under the appropriate service and flow.
3. Create a migration in the **target** service's database to create the replica table.
   - `migrate.go` automatically iterates all tenant schemas - just write a standard `CREATE TABLE IF NOT EXISTS` migration and the ts-migrations will apply it to every tenant schema.
   - The table definition must include at minimum all columns listed in `GetColumns()` + the primary key.
   - Run `go test ./internal/validator/... -v` and the integration tests to verify schema parity.
4. If the table uses `NonTenantToTenant` AND has a `client_id` field in the source:
   - Add `ALTER TABLE IF EXISTS <name> REPLICA IDENTITY FULL;` to the **source DB** migration.
   - Add `ALTER TABLE IF EXISTS <name> REPLICA IDENTITY DEFAULT;` to the down migration.
5. Configure PeerDB to include the new source table in the publication/mirror if it is not already covered.
