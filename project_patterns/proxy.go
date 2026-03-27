package projectpatterns

import "fmt"

type Database interface {
	Connect() string
	Query(query string) string
}

type RealDatabase struct{}

func (rdb *RealDatabase) Connect() string {
	return "Connecting to DB"
}

func (rdb *RealDatabase) Query(query string) string {
	return fmt.Sprintf("Query: %s", query)
}

type ProxyDB struct {
	RealDB Database
	user   string
}

func (proxy *ProxyDB) Connect() string {
	if proxy.user != "admin" {
		return "Error user role"
	}
	return proxy.RealDB.Connect()
}

func (proxy *ProxyDB) Query(query string) string {
	if proxy.user != "admin" {
		return "Error user role"
	}
	return proxy.RealDB.Query(query)
}
