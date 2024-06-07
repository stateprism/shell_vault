package localkeychain

const SetupDB string = `
CREATE TABLE IF NOT EXISTS keychain (
    	key_name VARCHAR(255) UNIQUE PRIMARY KEY NOT NULL,
    	key_type VARCHAR(255) NOT NULL,
    	key_value BLOB NOT NULL,
    	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	expiry_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	ttl BIGINT NOT NULL
);

CREATE TABLE IF NOT EXISTS expired_keys (
    	key_name VARCHAR(255) UNIQUE PRIMARY KEY NOT NULL,
    	key_type VARCHAR(255) NOT NULL,
    	expired_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS issued_certificates (
	cert_id VARCHAR(255) UNIQUE PRIMARY KEY NOT NULL,
    principal VARCHAR(255) NOT NULL,
	cert_type VARCHAR(255) NOT NULL,
	cert_value BLOB NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expiry_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ttl BIGINT NOT NULL
);

CREATE INDEX IF NOT EXISTS keychain_expiry_idx ON keychain (expiry_at);
CREATE INDEX IF NOT EXISTS certificate_principal_idx ON issued_certificates (principal);
`

const InsertKey string = `
INSERT INTO keychain (key_name, key_type, key_value, expiry_at, ttl) VALUES (?, ?, ?, ?, ?) RETURNING key_name;
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

const SelectExpiredKeys = `
SELECT key_name FROM keychain WHERE CURRENT_TIMESTAMP > expiry_at;
`

const RegisterNewCertificate string = `
INSERT INTO issued_certificates (cert_id, principal, cert_type, cert_value, expiry_at, ttl) VALUES (?, ?, ?, ?, ?, ?) RETURNING cert_id;
`

const SelectCertificate string = `
SELECT cert_value, cert_type, ttl FROM issued_certificates WHERE cert_id = ?;
`

const SelectCertificateByPrincipal string = `
SELECT cert_value, cert_type, ttl FROM issued_certificates WHERE principal = ?;
`
