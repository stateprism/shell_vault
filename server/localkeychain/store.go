package localkeychain

import "sync"

type keyStoreSer struct {
	Keys *sync.Map
}

type keyStore struct {
	*sync.Map
}
