package storage

import (
	"OfficialKnowledgeGraph/item"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"sync"
)

func MInsertAreaArea(areaAreaList []item.AreaArea) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	for _, areaArea := range areaAreaList {
		wg.Add(1)
		go func(areaArea item.AreaArea) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"MATCH (c:Area),(r:Area) WHERE c.id = $cid and r.id = $rid CREATE (c)-[belong:Belong]->(r)",
					map[string]interface{}{"cid": areaArea.ChildAreaID, "rid": areaArea.RootAreaID})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("areaArea: %v, insert failed; error: %v\n", areaArea, err)
			}
		}(areaArea)
	}
	wg.Wait()
	return
}

func MInsertUniversityArea(universityAreaList []item.UniversityArea) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	for _, universityArea := range universityAreaList {
		wg.Add(1)
		go func(universityArea item.UniversityArea) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"MATCH (u:University),(a:Area) WHERE u.id = $uid and a.id = $aid CREATE (u)-[locate:Locate]->(a)",
					map[string]interface{}{"uid": universityArea.UniversityID, "aid": universityArea.AreaID})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("universityArea: %v, insert failed; error: %v\n", universityArea, err)
			}
		}(universityArea)
	}
	wg.Wait()
	return
}

func MInsertPositionArea(positionAreaList []item.PositionArea) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	for _, positionArea := range positionAreaList {
		wg.Add(1)
		go func(positionArea item.PositionArea) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"MATCH (p:Position),(a:Area) WHERE p.id = $pid and a.id = $aid CREATE (p)-[base:Base]->(a)",
					map[string]interface{}{"pid": positionArea.PositionID, "aid": positionArea.AreaID})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("positionArea: %v, insert failed; error: %v\n", positionArea, err)
			}
		}(positionArea)
	}
	wg.Wait()
	return
}

func MInsertOfficialArea(officialAreaList []item.OfficialArea) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	for _, officialArea := range officialAreaList {
		wg.Add(1)
		go func(officialArea item.OfficialArea) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"MATCH (o:Official),(a:Area) WHERE o.id = $oid and a.id = $aid CREATE (o)-[grow:Grow]->(a)",
					map[string]interface{}{"oid": officialArea.OfficialID, "aid": officialArea.AreaID})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("officiaArea: %v, insert failed; error: %v\n", officialArea, err)
			}
		}(officialArea)
	}
	wg.Wait()
	return
}

func MInsertOfficialPosition(officialPositionList []item.OfficialPosition) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	for _, officialPosition := range officialPositionList {
		wg.Add(1)
		go func(officialPosition item.OfficialPosition) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"MATCH (o:Official),(p:Position) WHERE o.id = $oid and p.id = $pid CREATE (o)-[office:Office]->(p)",
					map[string]interface{}{"oid": officialPosition.OfficialID, "rid": officialPosition.PositionID})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("officialPosition: %v, insert failed; error: %v\n", officialPosition, err)
			}
		}(officialPosition)
	}
	wg.Wait()
	return
}

func MInsertOfficialUniversity(officialUniversityList []item.OfficialUniversity) {
	check()
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var wg sync.WaitGroup
	for _, officialUniversity := range officialUniversityList {
		wg.Add(1)
		go func(officialUniversity item.OfficialUniversity) {
			defer wg.Done()
			_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
				_, err := transaction.Run(
					"MATCH (o:Official),(u:University) WHERE o.id = $oid and u.id = $uid CREATE (o)-[graduate:Graduate]->(u)",
					map[string]interface{}{"oid": officialUniversity.OfficialID, "uid": officialUniversity.UniversityID})
				if err != nil {
					return nil, err
				}
				return nil, nil
			})
			if err != nil {
				fmt.Printf("officialUniversity: %v, insert failed; error: %v\n", officialUniversity, err)
			}
		}(officialUniversity)
	}
	wg.Wait()
	return
}
