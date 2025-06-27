package functions

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	if len(parts) == 0 {
		return s
	}

	builder := strings.Builder{}
	builder.WriteString(strings.ToLower(parts[0])) // bagian pertama lowercase

	titleCaser := cases.Title(language.English)
	for _, part := range parts[1:] {
		builder.WriteString(titleCaser.String(part)) // kapitalisasi dengan Unicode-aware
	}

	return builder.String()
}
