package migration

var FixedAssetView string = `CREATE OR REPLACE VIEW "category_tree" AS
SELECT build_category_tree(NULL, ARRAY_AGG(node  ORDER BY  (node->>'name'))) AS tree
FROM (
    SELECT jsonb_build_object(
        '_id', id::text,
        'name', name,
        'code', code,
        'key', key,
        'parent_id', parent_id,
        'children', '[]'::jsonb,
        'is_deleted', is_deleted
    ) AS node
    FROM categories
) AS all_nodes;


CREATE OR REPLACE VIEW "group_tree" AS
SELECT build_group_tree(NULL, ARRAY_AGG(node  ORDER BY  (node->>'name'))) AS tree
FROM (
    SELECT jsonb_build_object(
        '_id', id::text,
        'name', name,
        'code', code,
        'key', key,
        'parent_id', parent_id,
        'children', '[]'::jsonb,
        'is_deleted', is_deleted
    ) AS node
    FROM groups
) AS all_nodes;

CREATE OR REPLACE VIEW "my_asset_histories" AS
SELECT *
FROM (
    SELECT
        borrowing_requests.id AS id,
        borrowing_transactions.borrower_id AS user_id,
        borrowing_requests.asset AS asset,
        borrowing_requests.status::text AS status,
        borrowing_requests.updated_at AS updated_at,
        borrowing_requests.created_at AS created_at
    FROM borrowing_requests
        LEFT JOIN borrowing_transactions ON borrowing_transactions.id = borrowing_requests.transaction
    WHERE borrowing_requests.status::text IN ('Returned', 'Rejected')
    UNION
    SELECT
        assignment_requests.id AS id,
        assignment_transactions.assigned_to_id AS user_id,
        assignment_requests.asset AS asset,
        assignment_requests.status::text AS status,
        assignment_requests.updated_at AS updated_at,
        assignment_requests.created_at AS created_at
    FROM assignment_requests
        LEFT JOIN assignment_transactions ON assignment_transactions.id = assignment_requests.transaction
    WHERE assignment_requests.status::text IN ('Unassigned', 'Rejected')
) AS my_asset_histories;
 
 `
