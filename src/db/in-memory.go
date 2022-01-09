package db

import "fmt"

var Datastore map[string]string

// initialize datastore
func init() {
	Datastore = make(map[string]string)
	fmt.Println("Created In-Memory DB!")
}
