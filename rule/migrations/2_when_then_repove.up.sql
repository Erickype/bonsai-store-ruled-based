-- Remove the "when" column from the "rule" table
ALTER TABLE rule
DROP COLUMN "when";

-- Remove the "then" column from the "rule" table
ALTER TABLE rule
DROP COLUMN "then";
