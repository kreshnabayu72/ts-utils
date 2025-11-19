package dto

import "gitea.qwertysystem.net/BETS/ts-utils/v2/module/fixedasset/new_model"

type ClientPlan struct {
	PlanFixedAsset       new_model.PlanTypes `json:"planFixedAsset"              column:"plan_fixed_asset"`
	PlanSupplyAsset      new_model.PlanTypes `json:"planSupplyAsset"             column:"plan_supply_asset"`
	PlanInventoryAsset   new_model.PlanTypes `json:"planInventoryAsset"          column:"plan_inventory_asset"`
	PlanProcurementAsset new_model.PlanTypes `json:"planProcurementAsset"        column:"plan_procurement_asset"`
}
