package cache

// Record ...
type Record struct {
	Name   string `redis:"name"`
	Number int    `redis:"number"`
	JSON   []byte `redis:"json"`
}

// Handler ...
type Handler interface {
	// Health of the cache handler
	Ping() (string, error)
	Increment(counterKey string) (int, error)
	GetKey(key string) (interface{}, error)
	SetKey(key string, value interface{}) error
	GetRecord(keys []string) (Record, error)
	Close() error
}
