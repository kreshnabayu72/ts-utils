package helper

import (
	"context"
	"errors"
	"log"
	"strconv"

	"gitea.qwertysystem.net/BETS/ts-utils/db"
	"gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/new_model"
	"gitea.qwertysystem.net/BETS/ts-utils/service"
	"gitea.qwertysystem.net/BETS/ts-utils/sql_query"
)

// ChangeFixedAssetQuota adjusts the quota of a group and, if necessary, bubbles the
// change up the parent hierarchy until it finds a parent group that has available quota.
//
// Behavior:
//  1. Update the current group's quota:
//     - Decrease `available` (if not NULL).
//     - Increase `used_by_group`.
//     - Increase `used`.
//  2. If the current group still has available quota left, stop here.
//  3. Otherwise, walk up the parent chain:
//     - Increment `allocated_to_sub_groups` for each parent.
//     - Stop once a parent with `available > 0` is found.
//  4. Deduct from that parent's `available` and increment its `used`.
//
// Example usage:
//
//	// Moving quota between origin and destination groups
//	if err := util_fixedasset_helper.ChangeFixedAssetQuota(ctx, service, originGroupId, -quotaChanges, "Fixed Asset"); err != nil {
//		return err
//	}
//	if err := util_fixedasset_helper.ChangeFixedAssetQuota(ctx, service, destinationGroupId, quotaChanges, "Fixed Asset"); err != nil {
//		return err
//	}
//
// Parameters:
//   - ctx: context for timeout/cancel control
//   - service: PostgreSQL service abstraction
//   - groupId: ID of the group whose quota is being changed
//   - amount: positive or negative integer to add/subtract from quota
//   - quotaType: the specific quota type (e.g. Fixed Asset, Tracking, etc.)
func ChangeFixedAssetQuota(
	ctx context.Context,
	service service.PostgreSqlService,
	groupId string,
	amount int,
	quotaType new_model.GroupQuotaType,
) error {
	groupData, groupQuota, err := GetGroupWithQuota(ctx, service, groupId, quotaType)
	if err != nil {
		return err
	}

	// Step 1: update current group
	setValue := map[string]any{
		"available": sql_query.UpdateRawSQL{
			Expr: "CASE WHEN available IS NULL THEN NULL ELSE available - ? END",
			Args: []any{amount},
		},
		"used_by_this_group": sql_query.UpdateRawSQL{
			Expr: "used_by_this_group + ?",
			Args: []interface{}{amount},
		},
		"used": sql_query.UpdateRawSQL{Expr: "used + ?", Args: []interface{}{amount}},
	}
	if err := UpdateGroupQuota(service, ctx, setValue, groupData.ID, quotaType); err != nil {
		return err
	}

	// Step 2: if current group had available > 0, stop
	if groupQuota.Available != nil {
		if groupQuota.Type == quotaType && *groupQuota.Available > 0 {
			return nil
		}
	}

	// Step 3: bubble up parents
	parentId := groupData.ParentID
	var groupToUpdate *new_model.Group
	for parentId != nil {
		// Add to allocatedToSubGroups
		setParent := map[string]any{
			"allocated_to_sub_groups": sql_query.UpdateRawSQL{
				Expr: "allocated_to_sub_groups + ?",
				Args: []interface{}{amount},
			},
		}
		if err := UpdateGroupQuota(service, ctx, setParent, *parentId, quotaType); err != nil {
			return err
		}

		parentGroupData, parentGroupQuota, err := GetGroupWithQuota(
			ctx,
			service,
			strconv.Itoa(*parentId),
			quotaType,
		)
		if err != nil {
			return err
		}
		groupToUpdate = &parentGroupData

		// Stop if parent has available > 0
		if parentGroupQuota.Available != nil {
			if parentGroupQuota.Type == quotaType && *parentGroupQuota.Available > 0 {
				break
			}
		}

		// Keep climbing up
		parentId = parentGroupData.ParentID
	}

	// Step 4: deduct from the parent with available > 0
	if groupToUpdate != nil {
		setFinal := map[string]any{
			"available": sql_query.UpdateRawSQL{Expr: "available - ?", Args: []interface{}{amount}},
			"used":      sql_query.UpdateRawSQL{Expr: "used + ?", Args: []interface{}{amount}},
		}
		if err := UpdateGroupQuota(service, ctx, setFinal, groupToUpdate.ID, quotaType); err != nil {
			return err
		}
	}

	return nil
}

// UpdateGroupQuota builds and executes an UPDATE statement on the group_quotas table,
// incrementing the specified fields.
//
// Parameters:
//   - service: PostgreSQL service abstraction
//   - ctx: context for timeout/cancel control
//   - setValue: map of field names → SQL expressions to update
//   - groupId: group ID to update
//   - licenseType: quota type to filter by
func UpdateGroupQuota(
	service service.PostgreSqlService,
	ctx context.Context,
	setValue map[string]any,
	groupId int,
	licenseType new_model.GroupQuotaType,
) error {
	queryQuota, argsQuota, _ := sql_query.NewSQLUpdateBuilder(
		db.GroupQuotaTableName,
	).
		Update(setValue).
		Where(map[string]sql_query.SQLCondition{
			"group_id": {Operator: sql_query.SQLOperatorEqual, Value: groupId},
			"type":     {Operator: sql_query.SQLOperatorEqual, Value: licenseType},
		}).
		Build()
	if _, err := service.UpdateOne(
		ctx,
		queryQuota,
		argsQuota...,
	); err != nil {
		log.Println("err", err)
		return err
	}

	return nil
}

// GetGroupWithQuota fetches both the group row and its associated quota row
// for the specified quota type.
//
// Returns:
//   - new_model.Group containing group info
//   - new_model.GroupQuota containing quota info
//   - error if either group or quota was not found
func GetGroupWithQuota(
	ctx context.Context,
	service service.PostgreSqlService,
	groupId string,
	quotaType new_model.GroupQuotaType,
) (new_model.Group, new_model.GroupQuota, error) {
	queryGroup, argsGroup, _ := sql_query.NewSQLSelectBuilder[new_model.Group](db.GroupTableName).
		Where(map[string]sql_query.SQLCondition{
			"id": {Operator: sql_query.SQLOperatorEqual, Value: groupId},
		}).
		Build()
	groupData := new_model.Group{}
	if err := service.SelectOne(&groupData, ctx, queryGroup, argsGroup...); err != nil {
		return new_model.Group{}, new_model.GroupQuota{}, errors.New("group not found")
	}
	queryQuota, argsQuota, _ := sql_query.NewSQLSelectBuilder[new_model.GroupQuota](
		db.GroupQuotaTableName,
	).
		Where(map[string]sql_query.SQLCondition{
			"group_id": {Operator: sql_query.SQLOperatorEqual, Value: groupId},
			"type":     {Operator: sql_query.SQLOperatorEqual, Value: quotaType},
		}).
		Build()
	groupQuota := new_model.GroupQuota{}
	if err := service.SelectOne(&groupQuota, ctx, queryQuota, argsQuota...); err != nil {
		return new_model.Group{}, new_model.GroupQuota{}, errors.New("group quota not found")
	}

	return groupData, groupQuota, nil
}
