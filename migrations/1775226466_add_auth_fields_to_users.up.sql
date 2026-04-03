-- Add auth fields to users table
ALTER TABLE users ADD COLUMN IF NOT EXISTS password VARCHAR(255);
ALTER TABLE users ADD COLUMN IF NOT EXISTS is_active BOOLEAN DEFAULT true;
ALTER TABLE users ADD COLUMN IF NOT EXISTS last_login TIMESTAMP;

-- Set password to empty string for existing rows if null
UPDATE users SET password = '' WHERE password IS NULL;

-- Add NOT NULL constraint after setting default value
ALTER TABLE users ALTER COLUMN password SET NOT NULL;

-- Create indexes untuk performa query
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_users_is_active ON users(is_active);