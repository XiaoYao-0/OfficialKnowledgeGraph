package conf

import (
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	Neo4JURI      string `json:"neo4j_uri"`
	Neo4JUsername string `json:"neo4j_username"`
	Neo4JPassword string `json:"neo4j_password"`
}

func getConf() (Conf, error) {
	var conf Conf
	confJson, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		return conf, err
	}
	err = json.Unmarshal(confJson, &conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}

func Neo4JURI() (string, error) {
	neo4jURI := ""
	conf, err := getConf()
	if err != nil {
		return neo4jURI, err
	}
	neo4jURI = conf.Neo4JURI
	return neo4jURI, err
}

func Neo4JUsername() (string, error) {
	neo4JUsername := ""
	conf, err := getConf()
	if err != nil {
		return neo4JUsername, err
	}
	neo4JUsername = conf.Neo4JUsername
	return neo4JUsername, err
}

func Neo4JPassword() (string, error) {
	neo4JPassword := ""
	conf, err := getConf()
	if err != nil {
		return neo4JPassword, err
	}
	neo4JPassword = conf.Neo4JPassword
	return neo4JPassword, err
}
