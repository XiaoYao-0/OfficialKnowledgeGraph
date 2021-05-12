package storage

import (
	"OfficialKnowledgeGraph/item"
	"errors"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"sync"
)

const MAX_GO = 500

func QueryAreaIDByName(name string) int64 {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	id, err := session.ReadTransaction(func(tx neo4j.Transaction) (i interface{}, e error) {
		result, err := tx.Run(
			"MATCH (a:Area) WHERE a.name = $name RETURN a.id",
			map[string]interface{}{"name": name})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, errors.New("no result")
	})
	if err != nil {
		fmt.Printf("area.name: %v, query failed; error: %v\n", name, err)
		return -1
	}
	return id.(int64)
}

func QueryAreaLevelByName(name string) int {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	level, err := session.ReadTransaction(func(tx neo4j.Transaction) (i interface{}, e error) {
		result, err := tx.Run(
			"MATCH (a:Area) WHERE a.name = $name RETURN a.level",
			map[string]interface{}{"name": name})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, errors.New("no result")
	})
	if err != nil {
		fmt.Printf("area.name: %v, query failed; error: %v\n", name, err)
		return -1
	}
	return int(level.(int64))
}

func QueryAreaLevelByID(id int64) int {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	level, err := session.ReadTransaction(func(tx neo4j.Transaction) (i interface{}, e error) {
		result, err := tx.Run(
			"MATCH (a:Area) WHERE a.id = $id RETURN a.level",
			map[string]interface{}{"id": id})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, errors.New("no result")
	})
	if err != nil {
		fmt.Printf("area.id: %v, query failed; error: %v\n", id, err)
		return -1
	}
	return int(level.(int64))
}

func QueryPositionIDByName(name string) int64 {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	id, err := session.ReadTransaction(func(tx neo4j.Transaction) (i interface{}, e error) {
		result, err := tx.Run(
			"MATCH (p:Position) WHERE a.name = $name RETURN p.id",
			map[string]interface{}{"name": name})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, errors.New("no result")
	})
	if err != nil {
		fmt.Printf("position.name: %v, query failed; error: %v\n", name, err)
		return -1
	}
	return id.(int64)
}

func QueryPositionLevelByName(name string) int {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	level, err := session.ReadTransaction(func(tx neo4j.Transaction) (i interface{}, e error) {
		result, err := tx.Run(
			"MATCH (p:Position) WHERE a.name = $name RETURN p.level",
			map[string]interface{}{"name": name})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], nil
		}
		return nil, errors.New("no result")
	})
	if err != nil {
		fmt.Printf("position.name: %v, query failed; error: %v\n", name, err)
		return -1
	}
	return level.(int)
}

func MInsertArea(areaList []item.Area) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	ch := make(chan struct{}, MAX_GO)
	for _, area := range areaList {
		wg.Add(1)
		ch <- struct{}{}
		go func(area item.Area) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"CREATE (a:Area{id: $id, name: $name, level: $level})",
					map[string]interface{}{"id": area.ID, "name": area.Name, "level": area.Level})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("area: %v, insert failed; error: %v\n", area, err)
			}
			<-ch
		}(area)
	}
	wg.Wait()
	return
}

func MInsertUniversity(universityList []item.University) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	ch := make(chan struct{}, MAX_GO)
	for _, university := range universityList {
		wg.Add(1)
		ch <- struct{}{}
		go func(university item.University) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"CREATE (u:University{id: $id, name: $name})",
					map[string]interface{}{"id": university.ID, "name": university.Name})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("university: %v, insert failed; error: %v\n", university, err)
			}
			<-ch
		}(university)
	}
	wg.Wait()
	return
}

func MInsertOfficial(officialList []item.Official) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	ch := make(chan struct{}, MAX_GO)

	for _, official := range officialList {
		wg.Add(1)
		ch <- struct{}{}

		go func(official item.Official) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"CREATE (o:Official{id: $id, name: $name, gender: $gender, birth_year: $birth_year, nationality: $nationality})",
					map[string]interface{}{"id": official.ID, "name": official.Name, "gender": official.Gender, "birth_year": official.BirthYear, "nationality": official.Nationality})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("official: %v, insert failed; error: %v\n", official, err)
			}
			<-ch
		}(official)
	}
	wg.Wait()
	return
}

func MInsertPosition(positionList []item.Position) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	ch := make(chan struct{}, MAX_GO)
	for _, position := range positionList {
		wg.Add(1)
		ch <- struct{}{}
		go func(position item.Position) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"CREATE (p:Position{id: $id, name: $name, level: $level})",
					map[string]interface{}{"id": position.ID, "name": position.Name, "level": position.Level})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("position: %v, insert failed; error: %v\n", position, err)
			}
			<-ch
		}(position)
	}
	wg.Wait()
	return
}

func MDeleteOfficial() int {
	var officialIDList []int64
	count := 0
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	_, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(
			"MATCH (o:Official) RETURN o.id",
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		for result.Next() {
			officialIDList = append(officialIDList, result.Record().Values[0].(int64))
		}
		if err = result.Err(); err != nil {
			return nil, err
		}

		return officialIDList, result.Err()
	})
	if err != nil {
		fmt.Printf("official delete failed; error: %v\n", err)
		return 0
	}
	session.Close()
	var wg sync.WaitGroup
	var mutex sync.Mutex
	ch := make(chan struct{}, MAX_GO)
	for _, id := range officialIDList {
		wg.Add(1)
		ch <- struct{}{}
		go func(id int64) {
			defer wg.Done()
			if !QueryIsExistOfficialPositionByOfficialID(id) {
				deleteOfficial(id)
				mutex.Lock()
				count++
				mutex.Unlock()
			}
			<-ch
		}(id)
	}
	wg.Wait()
	return count
}

func deleteOfficial(id int64) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (i interface{}, e error) {
		_, err := tx.Run(
			"MATCH (o:Official) WHERE o.id = $oid DETACH DELETE o",
			map[string]interface{}{"oid": id})
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		fmt.Printf("official.id: %d delete failed; error: %v\n", id, err)
		return
	}
	return
}
