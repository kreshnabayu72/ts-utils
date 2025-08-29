package migration

var AdminConsoleEnum string = `-- SAFE ENUM CREATION (PostgreSQL)
-- client_status_enum
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'client_status_enum') THEN
    CREATE TYPE client_status_enum AS ENUM ('Active', 'Expired', 'Archived');
  END IF;
END $$;

-- open_api_status_enum
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'open_api_status_enum') THEN
    CREATE TYPE open_api_status_enum AS ENUM ('Approved', 'Waiting for Approval');
  END IF;
END $$;

-- license_plan_enum
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'license_plan_enum') THEN
    CREATE TYPE license_plan_enum AS ENUM ('Basic', 'Standard', 'Enterprise');
  END IF;
END $$;

-- depreciation_method_enum
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'depreciation_method_enum') THEN
    CREATE TYPE depreciation_method_enum AS ENUM ('Straight-Line', 'Double-Declining Balance');
  END IF;
END $$;

-- license_purchase_type_enum  (use this consistently; remove any license_purchase_type_enum usage)
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'license_purchase_type_enum') THEN
    CREATE TYPE license_purchase_type_enum AS ENUM ('Purchase', 'Trial');
  END IF;
END $$;

-- license_type_enum
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'license_type_enum') THEN
    CREATE TYPE license_type_enum AS ENUM (
      'concurrent_user',
      'fixed_asset',
      'fixed_asset_audit',
      'fixed_asset_asset_control',
      'fixed_asset_maintenance',
      'fixed_asset_repair_ticketing',
      'fixed_asset_tracking',
      'supply_asset',
      'supply_asset_asset_control',
      'supply_asset_stock_opname',
      'supply_asset_tracking',
      'inventory'
    );
  END IF;
END $$;
`
