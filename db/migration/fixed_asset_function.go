package migration

var FixedAssetFunction string = `

CREATE OR REPLACE FUNCTION build_category_tree(parent VARCHAR, nodes JSONB[])
RETURNS JSONB LANGUAGE plpgsql AS $$
DECLARE
    result JSONB := '[]'::jsonb;
    node JSONB;
    children JSONB;
    current_id VARCHAR;
    parent_id VARCHAR;
BEGIN
    FOREACH node IN ARRAY nodes LOOP
        BEGIN
            parent_id := (node->>'parent_id')::VARCHAR;
        EXCEPTION WHEN others THEN
            RAISE NOTICE 'Skipping node: %', node;
            CONTINUE;
        END;

        -- Skip deleted nodes
      IF (node->'is_deleted')::BOOLEAN IS TRUE THEN CONTINUE;
        END IF;

        IF parent_id IS NOT DISTINCT FROM parent THEN
            current_id := (node->>'_id')::TEXT;
            children := build_category_tree(current_id, nodes);

            result := result || jsonb_build_object(
                '_id', current_id,
                'name', node->>'name',
                'code', node->>'code',
                'key', node->'key',
                'children', children,
				'is_deleted',(node->'is_deleted')::BOOLEAN
            );
        END IF;
    END LOOP;

    RETURN result;
END;
$$;

CREATE OR REPLACE FUNCTION build_group_tree(parent VARCHAR, nodes JSONB[])
RETURNS JSONB LANGUAGE plpgsql AS $$
DECLARE
  result      JSONB := '[]'::jsonb;
  node        JSONB;
  children    JSONB;
  current_id  varchar;
  parent_id   varchar;
BEGIN
  FOREACH node IN ARRAY nodes LOOP
    BEGIN
      parent_id := (node ->> 'parent_id')::varchar;
    EXCEPTION WHEN others THEN
      RAISE NOTICE 'Skipping node: %', node;
      CONTINUE;
    END;

    -- Skip deleted nodes
    IF (node -> 'is_deleted')::boolean IS TRUE THEN
      CONTINUE;
    END IF;

    IF parent_id IS NOT DISTINCT FROM parent THEN
      current_id := (node ->> '_id')::text;
      children   := build_group_tree_with_role(current_id, nodes);

      result := result || jsonb_build_object(
        '_id',                      current_id,
        'name',                     node ->> 'name',
        'key',                      node -> 'key',
        'children',                 children,
        'is_deleted',               (node -> 'is_deleted')::boolean,
        'hasSystemRoles',           node -> 'hasSystemRoles',
        'fixedAssetQuota',          node -> 'fixedAssetQuota',
        'trackingQuota',            node -> 'trackingQuota',
        'assetControlQuota',        node -> 'assetControlQuota',
        'auditQuota',               node -> 'auditQuota',
        'maintenanceControlQuota',  node -> 'maintenanceControlQuota',
        'repairTicketingQuota',     node -> 'repairTicketingQuota'
      );
    END IF;
  END LOOP;

  RETURN result;
END;
$$;`
