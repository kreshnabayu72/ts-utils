package enum

// AccessLevel is a user's blanket access level for one product.
//
// There is deliberately no "normal" constant: a user with neither level simply has no
// user_access_levels row, so the column stays NOT NULL and every row it holds means
// something. Absence is the third state.
type AccessLevel string

const (
	AccessLevelTotalControl AccessLevel = "Total Control"
	AccessLevelReadOnly     AccessLevel = "Read Only"
)
