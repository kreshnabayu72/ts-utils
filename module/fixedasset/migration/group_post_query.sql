DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM groups) THEN
    INSERT INTO groups (id, name, code, parent_id,level,key)
    VALUES (0, 'All', 'ALL', NULL,0,-1);--ID needs to be 0 
  END IF;
END
$$;