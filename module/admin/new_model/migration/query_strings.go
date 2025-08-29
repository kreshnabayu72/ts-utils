package migration

var ClientPostQuery string = `
INSERT INTO clients (
    id,
    name,
    client_id,
    modified_by_id,
    email,
    is_email_confirmed,
    country,
    state,
    city,
    zip,
    phone_number,
    fax,
    address,
    contact_person_name,
    contact_person_email,
    is_default
)
VALUES (
    1,
    'PT ASDF',
    '80000001',
    101,
    'asdf@asdf.com',
    TRUE,
    'Indonesia',
    'Jakarta',
    'Jakarta Utara',
    '90001',
    '+62-1234',
    '+62-1234',
    'Kelapa gading',
    'User 1',
    'user1@qwertysystem.net',
    TRUE
);
`
var ClientLicensePostQuery string = `INSERT INTO client_licenses (
    id,
    client_id,
    type,
    managed,
    unmanaged,
    available,
    total,
    purchase_type
)
VALUES 
(1 , 1, 'concurrent_user',                0, 0, 100, 100, 'Purchase'),
(2 , 1, 'fixed_asset',                     0, 0,  100,  100, 'Purchase'),
(3 , 1, 'fixed_asset_audit',               0, 0,  20,  20, 'Purchase'),
(4 , 1, 'fixed_asset_asset_control',       0, 0,  30,  30, 'Purchase'),
(5 , 1, 'fixed_asset_maintenance',         0, 0,  15,  15, 'Purchase'),
(6 , 1, 'fixed_asset_repair_ticketing',    0, 0,  10,  10, 'Purchase'),
(7 , 1, 'fixed_asset_tracking',            0, 0,  25,  25, 'Purchase'),
(8 , 1, 'supply_asset',                    0, 0,  0,  0, 'Purchase'),
(9 , 1, 'supply_asset_asset_control',      0, 0,  0,  0, 'Purchase'),
(10, 1, 'supply_asset_stock_opname',       0, 0,  0,  0, 'Purchase'),
(11, 1, 'supply_asset_tracking',           0, 0,  0,  0, 'Purchase'),
(12, 1, 'inventory',                       0, 0,  0,  0, 'Purchase');`

var GroupPostQuery string = `

INSERT INTO groups (id, name,  parent_id,level)
VALUES (0, 'All', NULL,0);--ID needs to be 0

`
