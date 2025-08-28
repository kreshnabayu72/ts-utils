package migration

//using read sql file is a better practice but it currently doesnt work

var DivisionPostQuery string = `
INSERT INTO divisions (id,name)
VALUES (1, 'Default Division')
`
