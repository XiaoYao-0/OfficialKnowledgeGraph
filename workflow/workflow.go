package workflow

import (
	"OfficialKnowledgeGraph/collector"
	"OfficialKnowledgeGraph/extractor"
	"OfficialKnowledgeGraph/storage"
	"fmt"
)

// 初始化neo4j数据库
func initDB() {

}

// 批量写入Area实体和AreaArea关系
func areaWork() {
	areaList, areaAreaList := extractor.ExtractArea(collector.CollectAreaJson())
	err := storage.MInsertArea(areaList)
	if err != nil {
		panic(err)
	}
	err = storage.MInsertAreaArea(areaAreaList)
	if err != nil {
		panic(err)
	}
	return
}

// 批量写入University实体和UniversityArea关系
func universityWork() {
	universityList, universityAreaList := extractor.ExtractUniversityCSV(collector.CollectUniversityCSV1(), collector.CollectUniversityCSV2())
	err := storage.MInsertUniversity(universityList)
	if err != nil {
		panic(err)
	}
	err = storage.MInsertUniversityArea(universityAreaList)
	if err != nil {
		panic(err)
	}
	return
}

// 批量写入Position实体和PositionArea关系
func positionWork() {

}

// 批量写入Official实体和OfficialArea关系和OfficialPosition关系
func officialWork() {

}

// 执行整条流水线
func WorkStart() {
	fmt.Println("initDB start")
	initDB()
	fmt.Println("areaWork start")
	areaWork()
	fmt.Println("universityWork start")
	universityWork()
	fmt.Println("positionWork start")
	positionWork()
	fmt.Println("officialWork start")
	officialWork()
}
