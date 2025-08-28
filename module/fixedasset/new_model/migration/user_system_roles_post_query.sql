INSERT INTO user_system_roles (
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
  (16, 'Global', NULL, 1, 'Open API', 'openAPI');