ALTER TABLE festival
  DROP COLUMN date RESTRICT,
  ADD COLUMN start_date date,
  ADD COLUMN end_date date,
  DROP COLUMN location RESTRICT,
  ADD COLUMN country VARCHAR,
  ADD COLUMN state VARCHAR,
  ADD COLUMN city VARCHAR;
