// 爬取相应页面，

package extractor

import (
	"OfficialKnowledgeGraph/item"
	"OfficialKnowledgeGraph/storage"
	"strings"
)

func delimiter() []string {
	return []string{"国", "区", "县", "旗", "道", "院", "省", "盟", "州", "划", "会", "岛", "城", "镇", "湖", "域", "市", "园", "港"}
}

func extractPosition(name string, id int64, dict1 map[string]int, dict2 map[string]int, areaMap map[string]int64) (item.Position, []item.PositionArea) {
	position := item.Position{
		ID:   id,
		Name: name,
	}
	var positionAreaList []item.PositionArea
	length := len(name)
	areaLevelIndexes := make([]int, length)
	for k, v := range areaMap {
		if i := strings.Index(name, k); i != -1 {
			areaLevelIndexes[i] = storage.QueryAreaLevelByID(v)
			positionAreaList = append(positionAreaList, item.PositionArea{
				PositionID: id,
				AreaID:     v,
			})
		}
	}
	levelSum := 0
	levelCount := 0
	for k, v := range dict1 {
		if i := strings.Index(name, k); i != -1 {
			levelSum += v
			levelCount++
		}
	}
	for k, v := range dict2 {
		if i := strings.Index(name, k); i != -1 {
			for j := i; j >= 0; j-- {
				if al := areaLevelIndexes[j]; al != 0 {
					levelSum += v + al*2
					break
				}
			}
			levelSum += v
			levelCount++
		}
	}
	if levelCount == 0 {
		if len(positionAreaList) == 0 {
			// fmt.Println("ExtractPosition error, name:", name)
			return position, positionAreaList
		}
		positionAreaList = append(positionAreaList, item.PositionArea{
			PositionID: id,
			AreaID:     0,
		})
		position.Level = 10
	} else {
		position.Level = levelSum / levelCount
	}
	return position, positionAreaList
}
