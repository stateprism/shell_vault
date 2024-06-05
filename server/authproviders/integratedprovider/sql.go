package integratedprovider

const SetupTables = `
CREATE TABLE IF NOT EXISTS users (
    principal TEXT PRIMARY KEY,
    username TEXT NOT NULL,
    realm TEXT NOT NULL,
    extra_data BLOB NOT NULL
);
CREATE INDEX IF NOT EXISTS users_username ON users (username);
`

const ListUsersQuery = `SELECT principal FROM users;`
const GetUserQuery = `SELECT username, realm, extra_data FROM users WHERE principal = ?;`
const GetExtraDataQuery = `SELECT extra_data FROM users WHERE principal = ?;`
const AddUserQuery = `INSERT INTO users (principal, username, realm, extra_data) VALUES (?, ?, ?, ?);`
const UpdateUserQuery = `UPDATE users SET username = ?, realm = ?, extra_data = ? WHERE principal = ?;`
const DeleteUserQuery = `DELETE FROM users WHERE principal = ?;`
