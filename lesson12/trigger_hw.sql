CREATE TABLE IF NOT EXISTS audit_log(
    id SERIAL,
    table_name TEXT,
    operation TEXT,
    record_id INT,
    old_data JSONB,
    new_data JSONB,
    mod_time TIMESTAMP
);




CREATE OR REPLACE FUNCTION log_audit_changes()
RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO audit_log (table_name, operation, record_id, old_data, new_data, mod_time)
    VALUES (TG_TABLE_NAME, TG_OP, COALESCE(OLD.id, NEW.id), row_to_json(OLD), row_to_json(NEW), current_timestamp);
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER products_audit_trigger
AFTER INSERT OR UPDATE OR DELETE ON newtab
FOR EACH ROW
EXECUTE FUNCTION log_audit_changes();

CREATE TABLE newtab(id INT,name TEXT);
INSERT INTO newtab(id,name) VALUES(2,'czvgt');

