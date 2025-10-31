package sql_query

import (
	"fmt"
	"strings"
)

// quoteSingleQualifiedAttr handles the dot-separation and quoting for one attribute (e.g., "users.firstName").
func quoteSingleQualifiedAttr(qualifiedAttr string) string {
	// 1. Split the string by the dot ('.')
	segments := strings.Split(qualifiedAttr, ".")

	// 2. Quote each segment individually
	quotedSegments := make([]string, 0, len(segments))
	for _, segment := range segments {
		// Only quote non-empty segments
		segment = strings.TrimSpace(segment)
		if segment != "" {
			// Quoting the segment: e.g., "users", "firstName"
			quotedSegments = append(quotedSegments, fmt.Sprintf("\"%s\"", segment))
		}
	}

	// 3. Rejoin the segments with a dot
	return strings.Join(quotedSegments, ".")
}

// QuoteAttributeList processes a comma-separated list of attributes (e.g., "updatedAt, user.fullName")
// and quotes all identifiers correctly for PostgreSQL case-sensitivity.
func QuoteAttributeList(attrList string) string {
	// 1. Split the entire list by the comma (',').
	// Use strings.Split (not strings.SplitFunc) to handle spaces later.
	attributes := strings.Split(attrList, ",")

	finalQuotedList := make([]string, 0, len(attributes))

	// 2. Iterate over each attribute in the list.
	for _, attr := range attributes {
		// Remove leading/trailing whitespace from the attribute name
		trimmedAttr := strings.TrimSpace(attr)

		if trimmedAttr != "" {
			// 3. Apply the single attribute quoting logic (handles the dot '.')
			quotedAttr := quoteSingleQualifiedAttr(trimmedAttr)
			finalQuotedList = append(finalQuotedList, quotedAttr)
		}
	}

	// 4. Rejoin all fully quoted attributes with a comma and a space.
	return strings.Join(finalQuotedList, ", ")
}
