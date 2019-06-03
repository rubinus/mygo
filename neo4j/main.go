package main

import (
	"fmt"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
	driver  neo4j.Driver
	session neo4j.Session
	result  neo4j.Result
	err     error
)

func main() {
	err := TestNeo4j()
	if err != nil {
		panic(err)
	}
}

func TestNeo4j() error {
	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "12345678", "")); err != nil {
		return err // handle error
	}
	// handle driver lifetime based on your application lifetime requirements
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per application
	defer driver.Close()

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return err
	}
	defer session.Close()

	result, err = session.Run("CREATE (n:Person { id: $id, name: $name }) RETURN n.id, n.name", map[string]interface{}{
		"id":   "1",
		"name": "朱大仙儿②",
	})
	if err != nil {
		return err // handle error
	}

	for result.Next() {
		fmt.Printf("Created Item with Id = '%s' and Name = '%s'\n", result.Record().GetByIndex(0).(string), result.Record().GetByIndex(1).(string))
	}
	if err = result.Err(); err != nil {
		return err // handle error
	}
	return nil
}
