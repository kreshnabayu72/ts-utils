package migration

//using read sql file is a better practice but it currently doesnt work

var PositionPostQuery string = `
INSERT INTO positions (id,name)
VALUES (1, 'Default Position')
`
