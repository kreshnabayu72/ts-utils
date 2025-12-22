package dto

import "gitea.tagsamurai.local/BETS-V2/ts-utils/v2/enum"

type ClientPlan struct {
	PlanFixedAsset       enum.PlanTypes `json:"planFixedAsset"       column:"plan_fixed_asset"`
	PlanSupplyAsset      enum.PlanTypes `json:"planSupplyAsset"      column:"plan_supply_asset"`
	PlanInventoryAsset   enum.PlanTypes `json:"planInventoryAsset"   column:"plan_inventory_asset"`
	PlanProcurementAsset enum.PlanTypes `json:"planProcurementAsset" column:"plan_procurement_asset"`
}
