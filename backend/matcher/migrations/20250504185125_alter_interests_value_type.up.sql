ALTER TABLE interests
    ALTER COLUMN value TYPE int USING (value::integer);