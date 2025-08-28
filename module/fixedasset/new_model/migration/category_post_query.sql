DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM categories) THEN
    INSERT INTO categories (id, name, code, parent_id,key)
    VALUES (0, 'All', '', NULL,-1); --ID needs to be 0 
  END IF;
END
$$;