package redis

import (
	"log"
	"redistest/cache"

	"github.com/gomodule/redigo/redis"
)

const get string = "GET"
const set string = "SET"

type server struct {
	conn redis.Conn // connection handle to redis server
}

// Connect ...
func Connect() (cache.Handler, error) {

	log.Println("Opening connection")
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, err
	}
	return &server{conn: conn}, nil
}

func (s *server) Ping() (string, error) {
	return redis.String(s.conn.Do("PING"))
}

func (s *server) GetKey(key string) (interface{}, error) {

	switch key {
	case "number":
		return redis.Int64(s.conn.Do(get, key))
	case "name":
		return redis.String(s.conn.Do(get, key))
	case "json":
		return redis.Bytes(s.conn.Do(get, key))
	default:
		return s.conn.Do(get, key)
	}
}

func (s *server) SetKey(key string, value interface{}) error {
	reply, err := s.conn.Do(set, key, value)
	log.Println(reply)
	return err
}

func (s *server) Close() error {
	log.Println("Closing connection")
	return s.conn.Close()
}

func (s *server) Increment(counterKey string) (int, error) {
	return redis.Int(s.conn.Do("INCR", counterKey))
}

func (s *server) GetRecord(keys []string) (cache.Record, error) {

	reply, err := s.conn.Do("HMSET", "data:1", "name", "some", "number", 20)
	log.Println(reply, err)

	v, err := redis.Values(s.conn.Do("HGETALL", "data:1"))
	var r cache.Record
	err = redis.ScanStruct(v, &r)
	return r, err
}
