package migration

var FixedAssetEnum string = `-- NOTE: WITHOUT CHECKING, IT WILL THROW ERROR IF TYPE ALREADY EXIST
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'approval_type_enum') THEN
        CREATE TYPE approval_type_enum AS ENUM ('Or', 'And');
    END IF;
END$$;

-- Create yes_no_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'yes_no_enum') THEN
        CREATE TYPE yes_no_enum AS ENUM ('Yes', 'No');
    END IF;
END$$;

-- Create tag_type_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tag_type_enum') THEN
        CREATE TYPE tag_type_enum AS ENUM ('QR', 'RFID', 'RFID & QR', 'Non TAG');
    END IF;
END$$;

-- Create asset_status_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'asset_status_enum') THEN
        CREATE TYPE asset_status_enum AS ENUM (
            'Non-Transactionable',
            'Available',
            'Assigned',
            'Borrowed',
            'Disposed',
            'Damaged',
            'Missing',
            'On Assign Process',
            'On Audit Process',
            'On Borrowing Process',
            'On Disposal Process',
            'Reported Disposal',
            'On Transfer',
            'Expired License',
            'TAG Reported',
            'Not Paired Yet',
            'TAG Not Set'
        );

    END IF;
END$$;

-- Create asset_policy_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'asset_policy_enum') THEN
        CREATE TYPE asset_policy_enum AS ENUM ('Need Handover', 'Self-Service');
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'asset_default_image_enum') THEN
        CREATE TYPE asset_default_image_enum AS ENUM ('first_image', 'second_image');
    END IF;
END$$;


-- Create custom_field_type_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'custom_field_type_enum') THEN
        CREATE TYPE custom_field_type_enum AS ENUM ('General Information', 'Purchase','Accounting','Maintenance Routine','Audit');
    END IF;
END$$;

-- Create custom_field_data_type_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'custom_field_data_type_enum') THEN
        CREATE TYPE custom_field_data_type_enum AS ENUM ('Text',
         'Textarea','Dropdown','MultiDropdown','Radio button','Checkbox',
         'Date','Datetime','Currency','Numeric','Phone','Percentage','Email','URL','Document');
    END IF;
END$$;


-- Create system_role_name_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'system_role_name_enum') THEN
        CREATE TYPE system_role_name_enum AS ENUM (
            'Import data',
            'Manage asset names, brands, categories, and custom fields',
            'Manage asset policy',
            'Manage depreciation group',
            'Manage depreciation method',
            'Manage document deletion',
            'Manage group',
            'Manage user and role',
            'Manage IOT and reader',
            'Manage license',
            'Manage purchase document',
            'Manage purchase information',
            'Manage TAG',
            'Register assets',
            'Manage TAG type',
            'Open API'
        );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'system_role_type_enum') THEN
        CREATE TYPE system_role_type_enum AS ENUM (
            'Global',
            'Group'
        );
    END IF;
END$$;

-- Create system_role_attribute_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'system_role_attribute_enum') THEN
        CREATE TYPE system_role_attribute_enum AS ENUM (
            'importData',
            'assetAttribute',
            'assetPolicies',
            'depreciationGroup',
            'depreciationMethod',
            'documentDeletion',
            'group',
            'userAndRole',
            'iotReader',
            'license',
            'purchaseDocument',
            'purchaseInformation',
            'tag',
            'registerAsset',
            'tagType',
            'openAPI'
        );
    END IF;
END$$;

-- Create transaction_role_name_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_role_name_enum') THEN
        CREATE TYPE transaction_role_name_enum AS ENUM (
            'Borrowing/Assignment',
            'Disposal',
            'Transfer',
            'Audit',
            'Maintenance Routine',
            'Repair Ticketing',
            'Asset Control',
            'Tracking',
            'Damaged',
            'Missing'
        );
    END IF;
END$$;


-- Create transaction_role_attribute_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_role_attribute_enum') THEN
        CREATE TYPE transaction_role_attribute_enum AS ENUM (
            'borrowingRole',
            'disposalRole',
            'transferRole',
            'auditRole',
            'maintenanceRole',
            'repairRole',
            'assetControlRole',
            'trackingRole',
            'damagedRole',
            'missingRole'
        );
    END IF;
END$$;

-- Create transaction_role_type_enum if not exists
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_role_type_enum') THEN
        CREATE TYPE transaction_role_type_enum AS ENUM (
            'Manager',
            'Monitoring',
            'Approval',
            'Staff'
        );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'group_transaction_type_enum') THEN
        CREATE TYPE group_transaction_type_enum AS ENUM (
            'Borrowing/Assignment',
            'Transfer',
            'Disposal',
            'Damage',
            'Missing',
            'Repair Ticketing',
            'Tracking',
            'Audit',
            'Maintenance',
            'Asset Control'
        );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'group_quota_type_enum') THEN
        CREATE TYPE group_quota_type_enum AS ENUM (
            'Fixed Asset',
            'Repair Ticketing',
            'Tracking',
            'Audit',
            'Maintenance',
            'Asset Control'
        );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'module_access_enum') THEN
        CREATE TYPE module_access_enum AS ENUM ('Fixed Asset', 'Global Settings','Admin Console','Supply Asset');
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'license_transaction_type_enum') THEN
        CREATE TYPE license_transaction_type_enum AS ENUM (
            'Purchase', 'Trial'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'license_transaction_purchase_type_enum') THEN
        CREATE TYPE license_transaction_purchase_type_enum AS ENUM (
            'Extend','Additional','New'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'license_module_enum') THEN
        CREATE TYPE license_module_enum AS ENUM (
            'Fixed Asset',
            '+Audit',
            '+Asset Control',
            '+Maintenance Routine',
            '+Repair Ticketing',
            '+Tracking'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tag_status_enum') THEN
        CREATE TYPE tag_status_enum AS ENUM (
            'Available',
            'Paired',
            'Reported',
            'Damaged / Missing',
            'To be Returned'
            );
    END IF;
END$$;
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'attachment_type_enum') THEN
        CREATE TYPE attachment_type_enum AS ENUM (
            'Purchase',
            'Product'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'file_type_enum') THEN
        CREATE TYPE file_type_enum AS ENUM (
            'Attachment',
            'Damage',
            'Disposal',
            'Maintenance Routine',
            'Repair Ticketing'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'reader_status_enum') THEN
        CREATE TYPE reader_status_enum AS ENUM (
            'Available',
            'Damaged',
            'Missing'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'damage_report_status_enum') THEN
        CREATE TYPE damage_report_status_enum AS ENUM (
            'Damaged',
            'Repaired',
            'Reported Disposal',
            'On Disposal Process',
            'Disposed'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'missing_report_status_enum') THEN
        CREATE TYPE missing_report_status_enum AS ENUM (
            'Missing',
            'Found',
            'Reported Disposal',
            'On Disposal Process',
            'Disposed'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'asset_report_source_enum') THEN
        CREATE TYPE asset_report_source_enum AS ENUM (
            'Available Asset',
            'Audit Asset',
            'Borrowed Asset',
            'Assigned Asset',
            'Transfer Asset',
            'Tracking Asset',
            'Maintenance Asset'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'network_status_enum') THEN
        CREATE TYPE network_status_enum AS ENUM (
            'Online',
            'Offline'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'borrowing_transaction_status_enum') THEN
        CREATE TYPE borrowing_transaction_status_enum AS ENUM (
            'Canceled',
            'Waiting for Approval',
            'Waiting for Handover',
            'Borrowed',
            'Returned',
            'Extend Approval',
            'Request Extension',
            'Rejected'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'borrowing_request_status_enum') THEN
        CREATE TYPE borrowing_request_status_enum AS ENUM (
            'Canceled',
            'Waiting for Approval',
            'Waiting for Handover',
            'Waiting for Confirmation',
            'Borrowed',
            'Returned',
            'Approved',
            'Rejected',
            'Reported Damaged',
            'Reported Missing',
            'Reported Disposal',
            'On Disposal Process',
            'Disposed'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'assignment_transaction_status_enum') THEN
        CREATE TYPE assignment_transaction_status_enum AS ENUM (
            'Canceled',
            'Waiting for Approval',
            'Waiting for Handover',
            'Assigned',
            'Unassigned',
            'Rejected'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'assignment_request_status_enum') THEN
        CREATE TYPE assignment_request_status_enum AS ENUM (
            'Canceled',
            'Waiting for Approval',
            'Waiting for Handover',
            'Waiting for Confirmation',
            'Assigned',
            'Unassigned',
            'Approved',
            'Rejected',
            'Reported Damaged',
            'Reported Missing',
            'Reported Disposal',
            'On Disposal Process',
            'Disposed'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'approval_status_enum') THEN
        CREATE TYPE approval_status_enum AS ENUM (
            'Need Approval',
            'Finished Approval',
            'Pending'            
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'assignment_prelist_type_enum') THEN
        CREATE TYPE assignment_prelist_type_enum AS ENUM (
            'Assign',
            'Unassign'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transfer_group_type_enum') THEN
        CREATE TYPE transfer_group_type_enum AS ENUM (
            'Origin',
            'Destination'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transfer_transaction_status_enum') THEN
        CREATE TYPE transfer_transaction_status_enum AS ENUM (
            'Canceled',
            'Waiting for Approval',
            'Waiting for Handover',
            'Completed',
            'Rejected'   
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transfer_request_status_enum') THEN
        CREATE TYPE transfer_request_status_enum AS ENUM (
            'Canceled',
            'Reported Missing',
            'Waiting for Approval',
            'Waiting for Handover',
            'Completed',
            'Approved',
            'Rejected'           
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transfer_request_asset_tag_status_enum') THEN
        CREATE TYPE transfer_request_asset_tag_status_enum AS ENUM (
            'Verified',
            'Unverified',
            'TAG Reported'           
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transfer_request_asset_status_enum') THEN
        CREATE TYPE transfer_request_asset_status_enum AS ENUM (
            'Normal',
            'Damaged',
            'Missing'           
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'disposal_last_condition_enum') THEN
        CREATE TYPE disposal_last_condition_enum AS ENUM (
            'No Salvage Value',
            'Damaged',
            'Missing'           
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'disposal_method_enum') THEN
        CREATE TYPE disposal_method_enum AS ENUM (
            'Donate',
            'Sell',
            'Destroy',           
            'Clearance'           
            );
    END IF;
END$$;


DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'disposal_transaction_status_enum') THEN
        CREATE TYPE disposal_transaction_status_enum AS ENUM (
            'Canceled',
            'Waiting for Approval',
            'Waiting for Verification',
            'On Progress',
            'Completion APproval',
            'Completed',
            'Rejected'   
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'disposal_request_status_enum') THEN
        CREATE TYPE disposal_request_status_enum AS ENUM (
            'Canceled',
            'Waiting for Approval',
            'Waiting for Verification',
            'On Progress',
            'Completed',
            'Approved',
            'Rejected'           
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'disposal_group_type_enum') THEN
        CREATE TYPE disposal_group_type_enum AS ENUM (
            'Origin',
            'Disposal'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'repair_ticketing_status_enum') THEN
        CREATE TYPE repair_ticketing_status_enum AS ENUM (
            'Reported Damaged',
            'Under Repair',
            'Waiting for Confirmation',
            'Waiting for Approval',
            'Repaired',
            'Pemranently Damaged',
            'Reported Disposal',
            'On Disposal Process',
            'Disposed'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'repair_ticketing_method_enum') THEN
        CREATE TYPE repair_ticketing_method_enum AS ENUM (
            'Internal Repair',
            'Third Party',
            'Warranty Claim'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tracking_status_enum') THEN
        CREATE TYPE tracking_status_enum AS ENUM (
            'On Tracking',
            'Tracked by User',
            'Tracked by Maintenance',
            'Tracked by Audit',
            'Tracked by IOT',
            'Tracked by Tracking',
            'Reported as Permanently Missing',
            'Permanently Missing',
            'Reported Disposal',
            'On Disposal Process',
            'Disposed'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tracking_condition_enum') THEN
        CREATE TYPE tracking_condition_enum AS ENUM (
            'Asset found, normal condition',
            'Asset found, damaged condition',
            'TAG found, asset missing'
            );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'audit_type_enum') THEN
        CREATE TYPE audit_type_enum AS ENUM (
            'Existence',
            'Condition'
        );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'audit_task_status_enum') THEN
        CREATE TYPE audit_task_status_enum AS ENUM (
            'On Schedule',
            'In Audit',
            'Waiting for Approval',
            'Grace Period',
            'Overdue',
            'Completed',
            'Canceled',
            'Stopped'
        );
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'audit_task_asset_status_enum') THEN
        CREATE TYPE audit_task_asset_status_enum AS ENUM (
            'In Audit',
            'Audited'
        );
    END IF;
END$$;
`
