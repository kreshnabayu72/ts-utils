package sql_query

import (
	"fmt"
	"strings"
)

func GenerateCTEOption(
	sourceBuilder *SelectBuilder,
	mainBuilder *SelectBuilder,
	aliasName string,
	refTable string,
	optionsQuery bool,
	labelKey string,
	labelValue string,
	customSort ...Sort,
) {
	if optionsQuery {
		if !strings.Contains(labelKey, "->>") {
			labelKey = escapeQuoteColumns(labelKey)
		}
		if !strings.Contains(labelValue, "->>") {
			labelValue = escapeQuoteColumns(labelValue)
		}
		snakeAliasName := CamelToSnake(aliasName)

		orderString := fmt.Sprintf(`%s ASC`, labelKey)
		if len(customSort) > 0 {
			orderString = fmt.Sprintf(`%s %s`, customSort[0].SortBy, "ASC")
		}
		cteBuilder := sourceBuilder.
			ClearSelects().
			Distinct(
				fmt.Sprintf(`%s AS "value"`, labelValue),
				labelValue,
			).
			Select(
				fmt.Sprintf(`%s AS "key"`, labelKey),
				fmt.Sprintf(`"%s"."id" AS "id"`, refTable),
			)

		mainBuilder.WithCTEBuilder(snakeAliasName, cteBuilder.(*SelectBuilder).SQLEloquentQuery)
		mainBuilder.LeftJoin(
			snakeAliasName,
			fmt.Sprintf(`"%s"."id" = "%s"."id"`, snakeAliasName, refTable),
		)

		mainBuilder.Select(
			fmt.Sprintf(
				`jsonb_agg(
					jsonb_build_object(
						'label', "%s"."key",
						'value', "%s"."value"
					)
					ORDER BY %s
				) FILTER (WHERE "%s"."value" IS NOT NULL) AS "%s"`,
				snakeAliasName,
				snakeAliasName,
				orderString,
				snakeAliasName,
				aliasName,
			),
		)

	} else {
		mainBuilder.Select(
			fmt.Sprintf(
				`'[]'::jsonb AS "%s"`,
				aliasName,
			),
		)
	}
}
