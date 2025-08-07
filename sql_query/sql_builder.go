package sql_query

import "fmt"

func GenerateCTEOption(
	sourceBuilder *SelectBuilder,
	mainBuilder *SelectBuilder,
	aliasName string,
	refTable string,
	optionsQuery bool,
	labelKey string,
	labelValue string,
) {
	if optionsQuery {
		labelKey = escapeQuoteColumns(labelKey)
		labelValue = escapeQuoteColumns(labelValue)
		snakeAliasName := CamelToSnake(aliasName)

		cteBuilder := sourceBuilder.
			Distinct(
				fmt.Sprintf(`%s AS "value"`, labelValue),
				labelValue,
			).
			Select(
				fmt.Sprintf(`%s AS "key"`, labelKey),
				fmt.Sprintf(`"%s"."id" AS "id"`, refTable),
			)

		mainBuilder.WithCTEBuilder(snakeAliasName, cteBuilder.(*SelectBuilder).SQLEloquentQuery)
		mainBuilder.LeftJoin(snakeAliasName, fmt.Sprintf(`"%s"."id" = "%s"."id"`, snakeAliasName, refTable))

		mainBuilder.Select(
			fmt.Sprintf(
				`jsonb_agg(
					jsonb_build_object(
						'label', "%s"."label",
						'value', "%s"."value"
					)
				) filter (WHERE "%s"."value" IS NOT NULL) AS "%s"`,
				snakeAliasName,
				snakeAliasName,
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
