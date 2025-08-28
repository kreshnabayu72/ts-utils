package migration

//using read sql file is a better practice but it currently doesnt work

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
    contact_person_email
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
    'user1@qwertysystem.net'
);
`
