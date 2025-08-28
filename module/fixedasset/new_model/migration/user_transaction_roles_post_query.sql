
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
