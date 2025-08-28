package migration

// Note: Method of reading sql file doesnt work on server, so this strings were created
var DivisionPostQuery string = `INSERT INTO divisions (id,name)
VALUES (1, 'Default Division')`

var PositionPostQuery string = `INSERT INTO positions (id,name)
VALUES (1, 'Default Position')`

var UserPostQuery string = `INSERT INTO users (id, first_name, last_name, phone_number, email, position_id, division_id, access, modified_by_id)
VALUES (1, 'User', 'Satu', '081333','user1@qwertysystem.net', 1, 1, ARRAY['Global Settings', 'Fixed Asset'], 1)`
