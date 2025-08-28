DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger
        WHERE tgname = 'trigger_set_default_disposal_approval_status'
    ) THEN
        CREATE TRIGGER trigger_set_default_disposal_approval_status
        BEFORE INSERT ON disposal_approvals
        FOR EACH ROW
        EXECUTE FUNCTION set_default_approval_status();
    END IF;
END;
$$;