package storage

import (
	"OfficialKnowledgeGraph/conf"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

var driver neo4j.Driver

func InitDB() error {
	uri, err := conf.Neo4JURI()
	if err != nil {
		return err
	}
	username, err := conf.Neo4JUsername()
	if err != nil {
		return err
	}
	password, err := conf.Neo4JPassword()
	if err != nil {
		return err
	}
	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return err
	}
	return nil
}

func check() {
	if driver == nil {
		panic(errors.New("not init"))
	}
	return
}
