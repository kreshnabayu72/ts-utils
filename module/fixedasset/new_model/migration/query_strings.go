package migration

// Note: Method of reading sql file doesnt work on server, so this strings were created

var AssignmentApprovalPostQuery string = `DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_assignment_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_assignment_approval_status
        BEFORE INSERT ON assignment_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;`

var BorrowingApprovalPostQuery string = `DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_borrowing_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_borrowing_approval_status
        BEFORE INSERT ON borrowing_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;`

var CategoryPostQuery string = `DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM categories) THEN
    INSERT INTO categories (id, name, code, parent_id,key)
    VALUES (0, 'All', '', NULL,-1); --ID needs to be 0 
  END IF;
END
$$;`

var DisposalApprovalPostQuery string = `DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_disposal_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_disposal_approval_status
        BEFORE INSERT ON disposal_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;`

var GroupPostQuery string = `INSERT INTO groups (
    id, name, code, parent_id, parent_quota_source_id, level,
) VALUES
(0, 'All', 'ALL', NULL, 0, -1),
(1, 'Default','0000', 0, 0, 1);`

var RepairApprovalPostQuery string = `DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_repair_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_repair_approval_status
        BEFORE INSERT ON repair_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;`

var RoutineApprovalPostQuery string = `DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_routine_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_routine_approval_status
        BEFORE INSERT ON routine_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;`

var ServiceCenterPostQuery string = `INSERT INTO service_centers (id, name, is_default)
VALUES (0, 'Default Service Center', TRUE);`

var TrackingApprovalPostQuery string = `DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_tracking_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_tracking_approval_status
        BEFORE INSERT ON tracking_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;`

var TransferApprovalPostQuery string = `DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_transfer_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_transfer_approval_status
        BEFORE INSERT ON transfer_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;`

var UserPostQuery string = `INSERT INTO users (id, first_name, last_name, phone_number, email, division, position,
 modified_by_id, modified_by_first_name, modified_by_last_name, modified_by_full_name, is_email_confirmed,
 read_only_is_active, read_only_is_assigned, total_control_is_active, total_control_is_assigned
),
VALUES (1, 'User', 'Satu', '081333','user1@qwertysystem.net', 'Boss', 'Boss',
 1, 'User', 'Satu', 'User Satu',TRUE,
 FALSE, FALSE, TRUE, TRUE
)
`
var UserSystemRolePostQuery string = `INSERT INTO user_system_roles (
  id, role_type, group_id, user_id, role_name, role_attribute
) VALUES
  (1, 'Group',  1,    1, 'Register assets', 'registerAsset'),
  (2, 'Group',  1,    1, 'Manage group', 'group'),
  (3, 'Global',  NULL, 1, 'Manage asset names, brands, categories, and custom fields', 'assetAttribute')
  (4 , 'Global', NULL, 1, 'Import data', 'importData'),
  (5 , 'Global', NULL, 1, 'Manage asset policy', 'assetPolicies'),
  (6 , 'Global', NULL, 1, 'Manage depreciation group', 'depreciationGroup'),
  (7 , 'Global', NULL, 1, 'Manage depreciation method', 'depreciationMethod'),
  (8 , 'Group',  1, 1, 'Manage document deletion', 'documentDeletion'),
  (9 , 'Group',  1, 1, 'Manage user and role', 'userAndRole'),
  (10, 'Group',  1, 1, 'Manage IOT and reader', 'iotReader'),
  (11, 'Group',  1, 1, 'Manage license', 'license'),
  (12, 'Group',  1,    1, 'Manage purchase document', 'purchaseDocument'),
  (13, 'Group',  1,    1, 'Manage purchase information', 'purchaseInformation'),
  (14, 'Group',  1,    1, 'Manage TAG', 'tag'),
  (15, 'Global',  NULL,    1, 'Manage TAG type', 'tagType'),
  (16, 'Global', NULL, 1, 'Open API', 'openAPI');`

var UserTransactionRolePostQuery string = `
INSERT INTO user_transaction_roles (
  id, user_id, role_attribute, role_name,
  role_type, group_id
) VALUES
  (1, 1, 'borrowingRole', 'Borrowing/Assignment', 'Manager', 1),
  (2, 1, 'disposalRole', 'Disposal', 'Manager', 1),
  (3, 1, 'transferRole', 'Transfer', 'Manager', 1),
  (4, 1, 'auditRole', 'Audit', 'Manager', 1),
  (5, 1, 'maintenanceRole', 'Maintenance Routine', 'Manager', 1),
  (6, 1, 'repairRole', 'Repair Ticketing', 'Manager', 1),
  (7, 1, 'assetControlRole', 'Asset Control', 'Manager', 1),
  (8, 1, 'trackingRole', 'Tracking', 'Manager', 1),
  (9, 1, 'damagedRole', 'Damaged', 'Manager', 1),
  (10, 1, 'missingRole', 'Missing', 'Manager', 1);
`
