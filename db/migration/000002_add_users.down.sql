-- Drop the unique constraint on the combination of "owner" and "currency"
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- Drop the foreign key from "accounts" referencing "users"
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drop the "users" table
DROP TABLE IF EXISTS "users";
