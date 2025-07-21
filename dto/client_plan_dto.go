package dto

import "gitea.qwertysystem.net/BETS/ts-utils/module/fixedasset/model"

type ClientPlan struct {
	PlanFixedAsset       model.PlanTypes `json:"planFixedAsset"              column:"plan_fixed_asset"`
	PlanSupplyAsset      model.PlanTypes `json:"planSupplyAsset"             column:"plan_supply_asset"`
	PlanInventoryAsset   model.PlanTypes `json:"planInventoryAsset"          column:"plan_inventory_asset"`
	PlanProcurementAsset model.PlanTypes `json:"planProcurementAsset"        column:"plan_procurement_asset"`
}
