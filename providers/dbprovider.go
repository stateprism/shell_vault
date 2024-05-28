package providers

type DatabaseProvider interface {
	// Authenticate checks the provided username and password against the database
	Authenticate(username, password string) (bool, error)
	// GetUser, used by some auth providers to get the user's data
	GetUser(username string) (string, error)
	// CreateUser creates a new user in the database, used by some auth providers
	CreateUser(username, password string) error
	// GetSession retrieves the session for the provided token
	GetSession(token string) (string, error)
	// CreateSession creates a new session for the provided token
	CreateSession(token, username string) error
	// DeleteSession deletes the session for the provided token
	DeleteSession(token string) error
	// Close closes the connection to the database
	Close() error
}
