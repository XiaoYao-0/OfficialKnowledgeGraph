package workflow

import (
	"OfficialKnowledgeGraph/collector"
	"OfficialKnowledgeGraph/extractor"
)

// 批量写入University实体和UniversityArea关系
func universityWork() {

}

// 批量写入Area实体和AreaArea关系
func areaWork() {
	areaList, areaAreaList := extractor.ExtractArea(collector.CollectAreaJson())

}

// 批量写入Position实体和PositionArea关系
func positionWork() {

}

// 批量写入Official实体和OfficialArea关系和OfficialPosition关系
func officialWork() {

}

//
