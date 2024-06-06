package localkeychain

const CreateTables string = `
CREATE TABLE IF NOT EXISTS keychain (
    	key_name VARCHAR(255) UNIQUE PRIMARY KEY NOT NULL,
    	key_type VARCHAR(255) NOT NULL,
    	key_value BLOB NOT NULL,
    	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	ttl BIGINT NOT NULL
);
`

const InsertKey string = `
INSERT INTO keychain (key_name, key_type, key_value, ttl) VALUES (?, ?, ?, ?) RETURNING key_name;
`

const SelectKey string = `
SELECT key_value, key_type, ttl FROM keychain WHERE key_name = ?;
`

const KeychainExists string = `
SELECT key_name FROM keychain LIMIT 1;
`

const DropKey string = `
DELETE FROM keychain WHERE key_name = ?;
`
