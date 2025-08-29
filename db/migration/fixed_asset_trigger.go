package migration

var FixedAssetTrigger string = `

CREATE OR REPLACE FUNCTION set_default_approval_status()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status IS NULL THEN
        IF NEW.level = 1 THEN
            NEW.status := 'Need Approval';
        ELSE
            NEW.status := 'Pending';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

`
