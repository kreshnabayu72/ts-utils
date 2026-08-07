package query_builder

// import (
// 	"strings"
// 	"testing"

// 	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
// )

// func TestAuthUserRolesQueryBuilder(t *testing.T) {
// 	query, args, err := AuthUserRolesQueryBuilder("42", "acme")
// 	if err != nil {
// 		t.Fatalf("AuthUserRolesQueryBuilder: %v", err)
// 	}

// 	// The user id must be parameterized, not interpolated.
// 	foundUserID := false
// 	for _, a := range args {
// 		if a == "42" {
// 			foundUserID = true
// 		}
// 	}
// 	if !foundUserID {
// 		t.Errorf("user id '42' not passed as a bind arg; args=%v", args)
// 	}

// 	mustContain := []string{
// 		"user_system_roles",
// 		"fixed_asset_user_transaction_roles",
// 		"group_role_cte",
// 		"transaction_role_cte",
// 		"jsonb_agg",   // group_id aggregation for group-scoped permissions
// 		"FILTER",      // per-permission filtering
// 		`as "clientId"`,
// 	}
// 	for _, frag := range mustContain {
// 		if !strings.Contains(query, frag) {
// 			t.Errorf("query missing %q\n---\n%s", frag, query)
// 		}
// 	}

// 	// Approver transaction roles are excluded from the snapshot.
// 	if !strings.Contains(query, "!=") && !strings.Contains(query, "<>") {
// 		t.Errorf("expected an inequality filter (Approver exclusion) in query:\n%s", query)
// 	}

// 	// user_system_roles is shared across products and role keys collide between them
// 	// (importData, tag, userAndRole), so BOTH reads of it — the group CTE and the
// 	// global LEFT JOIN — must be constrained to Fixed Asset. Dropping either filter
// 	// grants a user another product's permission of the same name.
// 	if got := strings.Count(query, `."product"`); got != 2 {
// 		t.Errorf("expected 2 product filters on user_system_roles, got %d:\n%s", got, query)
// 	}
// 	productArgs := 0
// 	for _, a := range args {
// 		if a == string(enum.ModuleAccessFixedAsset) {
// 			productArgs++
// 		}
// 	}
// 	if productArgs != 2 {
// 		t.Errorf("expected both product filters to bind %q, got %d; args=%v",
// 			enum.ModuleAccessFixedAsset, productArgs, args)
// 	}
// }
