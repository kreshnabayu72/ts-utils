package sql_query

import "fmt"

func FormatPaginateQuery(dataQuery, totalQuery string) string {
	paginationQuery := fmt.Sprintf(`
		WITH
			paginated AS (%s),
			total AS (%s)
		SELECT
			COALESCE((SELECT json_agg(paginated) FROM paginated), '[]') AS data,
			(SELECT COUNT FROM total) AS totalRecords;
	`, dataQuery, totalQuery)

	fmt.Printf("build query pagination raw: %v, type: %T\n", paginationQuery, paginationQuery)
	return paginationQuery
}
