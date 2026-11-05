package rbac

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/constant"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"
	"gitea.tagsamurai.local/BETS-V2/ts-utils/v2/middleware"
	"gitea.tagsamurai.local/Wangsit-Developer/be-core-utils/entity"

	"github.com/gofiber/fiber/v3"
)

type approverScope struct {
	GroupManagerGroupIDs       []string `json:"groupManagerGroupIds"`
	AdditionalApproverGroupIDs []string `json:"additionalApproverGroupIds"`
	GlobalAdditionalApprover   bool     `json:"globalAdditionalApprover"`
}

type approverScopeKey string

// AdminConsoleApprover passes for an admin-console total-control caller, otherwise when iam's live
// approver pool for roleKey names the caller: a global or group additional approver, or — with
// withGroupManagers — a group manager. The scope is fetched from iam (not the snapshot) and cached
// in Locals per request+roleKey, so composing several approver rules costs one call.
//
// Needs LoadUserRoles mounted: it supplies the iam base URL, service token and client.
func AdminConsoleApprover(roleKey enum.SystemRoleKey, withGroupManagers bool) Rule {
	return func(c fiber.Ctx, roles middleware.UserAccessCache) (bool, error) {
		if roles.IsTotalControl(enum.ModuleAccessAdminConsole) {
			return true, nil
		}
		scope, err := loadApproverScope(c, roleKey)
		if err != nil {
			return false, err
		}
		return scope.GlobalAdditionalApprover ||
			len(scope.AdditionalApproverGroupIDs) > 0 ||
			(withGroupManagers && len(scope.GroupManagerGroupIDs) > 0), nil
	}
}

func loadApproverScope(c fiber.Ctx, roleKey enum.SystemRoleKey) (approverScope, error) {
	key := approverScopeKey(roleKey)
	if cached, ok := c.Locals(key).(approverScope); ok {
		return cached, nil
	}
	ep, ok := middleware.GetIAMEndpoint(c)
	if !ok {
		return approverScope{}, fmt.Errorf("approver scope: iam endpoint not configured (LoadUserRoles not mounted)")
	}
	userID := c.Get(constant.HEADER_USER_ID)
	clientID := c.Get(constant.HEADER_CLIENT_ID)

	url := fmt.Sprintf("%s/internal/users/%s/approver-scope?roleKey=%s", ep.BaseURL, userID, roleKey)
	req, err := http.NewRequestWithContext(c.Context(), http.MethodGet, url, nil)
	if err != nil {
		return approverScope{}, err
	}
	req.Header.Set(constant.HEADER_USER_ID, userID)
	req.Header.Set(constant.HEADER_CLIENT_ID, clientID)
	req.Header.Set(constant.HEADER_SERVICE_TOKEN, ep.ServiceToken)

	resp, err := ep.Client.Do(req)
	if err != nil {
		return approverScope{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return approverScope{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return approverScope{}, &entity.HttpError{Code: resp.StatusCode, Message: fmt.Sprintf("iam returned %d: %s", resp.StatusCode, string(body))}
	}
	var envelope struct {
		Data approverScope `json:"data"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return approverScope{}, err
	}
	c.Locals(key, envelope.Data)
	return envelope.Data, nil
}
