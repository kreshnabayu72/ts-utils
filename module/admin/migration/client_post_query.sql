INSERT INTO clients (
    id,
    name,
    status,
    pin,
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
    pending_email_change,
    image,
    login_session_minutes,

    plan_fixed_asset,
    plan_supply_asset,
    plan_inventory_asset,
    plan_procurement_asset,

    depreciation_method,
    pending_depreciation_method,

    open_api_status,
    open_api_requested_at,
    open_api_approved_at,

    addon_repair_ticketing,
    addon_maintenance,
    addon_tracking,
    addon_asset_control,
    addon_audit,

    date_format,
    time_format,
    geolocation,
    time_zone,

    currency_symbol,
    currency_currency,
    currency_locale,
    currency_label,

    audit_attrs_max_grace_period,
    transfer_attrs_transaction_timeout,
    transfer_attrs_asset_control_to_missing,
    borrowing_attrs_max_borrowing_time,
    borrowing_attrs_default_borrowing_time,
    borrowing_attrs_confirmation_timeout,
    borrowing_attrs_transaction_timeout,
    disposal_attrs_transaction_timeout,

    file_storage_total_kb,
    file_storage_available_kb,
    file_storage_used_kb,
    file_storage_attachment_kb,
    file_storage_damage_kb,
    file_storage_disposal_kb,
    file_storage_repair_ticketing_kb,
    file_storage_routine_kb,

    expired_at,
    created_at,
    updated_at
)
VALUES (
    1,
    'Acme Corp',
    'Active',
    '123456',
    'CLNT-0001',

    101,

    'admin@acmecorp.com',
    TRUE,

    'USA',
    'California',
    'Los Angeles',
    '90001',
    '+1-213-555-0100',
    '+1-213-555-0199',
    '123 Sunset Blvd, Los Angeles, CA',

    'Jane Doe',
    'jane.doe@acmecorp.com',
    NULL,
    'https://example.com/image.png',
    60,

    'Basic',
    'Basic',
    'Basic',
    'Basic',

    'Straight-Line',
    NULL,

    'Approved',
    '2025-07-01T10:00:00+00:00',
    '2025-07-02T15:30:00+00:00',

    TRUE,
    TRUE,
    FALSE,
    TRUE,
    FALSE,

    'DD-MM-YYYY',
    'HH:mm:ss',
    TRUE,
    'America/Los_Angeles',

    '$',
    'USD',
    'en-US',
    'US Dollar',

    14,
    30,
    60,
    1440,
    720,
    10,
    15,
    20,

    100000.0,
    75000.0,
    25000.0,
    5000.0,
    300.0,
    100.0,
    250.0,
    120.0,

    NULL,
    NOW(),
    NOW()
);