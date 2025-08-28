package migration

//using read sql file is a better practice but it currently doesnt work

var GroupPostQuery string = `
DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM groups) THEN
    INSERT INTO groups (id, name,  parent_id,level)
    VALUES (0, 'All', NULL,0);--ID needs to be 0 
  END IF;
END
$$;
`
