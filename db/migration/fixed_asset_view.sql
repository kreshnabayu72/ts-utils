CREATE OR REPLACE VIEW "category_tree" AS
SELECT build_category_tree(NULL, ARRAY_AGG(node)) AS tree
FROM (
    SELECT jsonb_build_object(
        '_id', id,
        'name', name,
        'code', code,
        'key', key,
        'parent_id', parent_id,
        'children', '[]'::jsonb,
        'is_deleted', is_deleted
    ) AS node
    FROM categories
) AS all_nodes;