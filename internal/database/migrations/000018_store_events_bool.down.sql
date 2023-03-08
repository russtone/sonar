ALTER TABLE payloads
  ALTER COLUMN store_events DROP DEFAULT,
  ALTER COLUMN store_events TYPE INT USING (CASE WHEN store_events THEN 1000 ELSE 0 END),
  ALTER COLUMN store_events SET DEFAULT 0;
