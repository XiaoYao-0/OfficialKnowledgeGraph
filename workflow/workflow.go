package workflow

import (
	"OfficialKnowledgeGraph/collector"
	"OfficialKnowledgeGraph/extractor"
	"OfficialKnowledgeGraph/storage"
	"fmt"
	"sync"
	"time"
)

var areaMap map[string]int64
var universityMap map[string]int64

// 初始化neo4j数据库
func initDB() {
	err := storage.InitDB()
	if err != nil {
		panic(err)
	}
}

// 批量写入Area实体和AreaArea关系
func areaWork() {
	areaList, areaAreaList := extractor.ExtractArea(collector.CollectAreaJson())
	areaMap = extractor.AreaMap(areaList)
	fmt.Println("areaWork: ExtractArea() and AreaMap()")

	storage.MInsertArea(areaList)
	fmt.Printf("areaWork: MInsertArea(), %d Areas\n", len(areaList))

	storage.MInsertAreaArea(areaAreaList)
	fmt.Printf("areaWork: MInsertAreaArea(), %d AreaAreas\n", len(areaAreaList))
	return
}

// 批量写入University实体和UniversityArea关系
func universityWork() {
	universityList, universityAreaList := extractor.ExtractUniversityCSV(collector.CollectUniversityCSV1(), collector.CollectUniversityCSV2())
	universityMap = extractor.UniversityMap(universityList)
	fmt.Println("universityWork: ExtractUniversityCSV() and UniversityMap()")

	storage.MInsertUniversity(universityList)
	fmt.Printf("universityWork: MInsertUniversity(), %d Universities\n", len(universityList))

	storage.MInsertUniversityArea(universityAreaList)
	fmt.Printf("universityWork: MInsertUniversityArea(), %d UniversityAreas\n", len(universityAreaList))
	return
}

// 批量写入Official实体Position实体和OfficialArea关系和OfficialPosition关系和PositionArea关系和OfficialUniversity关系
func officialPositionWork() {
	var urlList []string
	var wg sync.WaitGroup
	var mutex sync.Mutex
	files := collector.SplitFile()
	fmt.Println("officialPositionWork: SplitFile()")
	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			urlList0 := collector.CollectOfficialURLList(file)
			mutex.Lock()
			urlList = append(urlList, urlList0...)
			mutex.Unlock()
			fmt.Printf("%v done\n", file)
		}(file)
	}
	wg.Wait()
	fmt.Printf("officialPositionWork: CollectOfficialURLList(), %d urls\n", len(urlList))

	officialList, officialAreaList, officialUniversityList, officialPositionList, positionList, positionAreaList := extractor.ExtractOfficial(urlList, areaMap, universityMap)
	fmt.Println("officialPositionWork: ExtractOfficial()")

	storage.MInsertOfficial(officialList)
	fmt.Printf("officialPositionWork: MInsertOfficial(), %d Officials\n", len(officialList))

	storage.MInsertOfficialArea(officialAreaList)
	fmt.Printf("officialPositionWork: MInsertOfficialArea(), %d OfficialAreas\n", len(officialAreaList))

	storage.MInsertOfficialUniversity(officialUniversityList)
	fmt.Printf("officialPositionWork: MInsertOfficialUniversity(), %d OfficialUniversitys\n", len(officialUniversityList))

	storage.MInsertPosition(positionList)
	fmt.Printf("officialPositionWork: MInsertPosition(), %d Positions\n", len(positionList))

	storage.MInsertOfficialPosition(officialPositionList)
	fmt.Printf("officialPositionWork: MInsertOfficialPosition(), %d OfficialPositions\n", len(officialPositionList))

	storage.MInsertPositionArea(positionAreaList)
	fmt.Printf("officialPositionWork: MInsertPositionArea(), %d PositionAreas\n", len(positionAreaList))

	count := storage.MDeleteOfficial()
	fmt.Printf("officialPositionWork: MDeleteOfficial(), %d Officials\n", count)
}

// 执行整条流水线
func WorkStart() {
	t0 := time.Now()
	fmt.Println("initDB start")
	initDB()
	t1 := time.Now()
	fmt.Printf("initDB end, consumed %d ms\n", t1.Sub(t0).Milliseconds())
	fmt.Println("areaWork start")
	areaWork()
	t2 := time.Now()
	fmt.Printf("areaWork end, consumed %d ms\n", t2.Sub(t1).Milliseconds())
	fmt.Println("universityWork start")
	universityWork()
	t3 := time.Now()
	fmt.Printf("universityWork end, consumed %d ms\n", t3.Sub(t2).Milliseconds())
	fmt.Println("officialPositionWork start")

	officialPositionWork()
	t4 := time.Now()
	fmt.Printf("officialPositionWork end, consumed %d ms\n", t4.Sub(t3).Milliseconds())

	fmt.Printf("Successful, consumed %d ms\n", t4.Sub(t0).Milliseconds())
}
