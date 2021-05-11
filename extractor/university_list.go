// 数据来源http://www.moe.gov.cn/jyb_xxgk/s5743/s5744/202007/t20200709_470937.html

package extractor

import (
	"OfficialKnowledgeGraph/item"
	"OfficialKnowledgeGraph/storage"
	"bufio"
	"strings"
)

var id int64

func extractUniversityCSV1(universityCSV1 string) ([]item.University, []item.UniversityArea) {
	var universityList []item.University
	var universityAreaList []item.UniversityArea
	reader := strings.NewReader(universityCSV1)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "序号") || strings.HasSuffix(s, ",,,,,") {
			continue
		}
		split := strings.Split(s, ",")
		name, areaName := split[1], split[4]
		universityList = append(universityList, item.University{
			ID:   id,
			Name: name,
		})
		areaID := storage.QueryAreaIDByName(areaName)
		universityAreaList = append(universityAreaList, item.UniversityArea{
			UniversityID: id,
			AreaID:       areaID,
		})
		id++
	}
	return universityList, universityAreaList
}

func extractUniversityCSV2(universityCSV2 string) ([]item.University, []item.UniversityArea) {
	var universityList []item.University
	var universityAreaList []item.UniversityArea
	reader := strings.NewReader(universityCSV2)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		s := scanner.Text()
		if strings.HasPrefix(s, "序号") || strings.HasSuffix(s, ",,,,,") {
			continue
		}
		split := strings.Split(s, ",")
		name, areaName := split[1], split[3]
		universityList = append(universityList, item.University{
			ID:   id,
			Name: name,
		})
		areaID := storage.QueryAreaIDByName(areaName)
		universityAreaList = append(universityAreaList, item.UniversityArea{
			UniversityID: id,
			AreaID:       areaID,
		})
		id++
	}
	return universityList, universityAreaList
}

func ExtractUniversityCSV(universityCSV1, universityCSV2 string) ([]item.University, []item.UniversityArea) {
	id = 1
	uL1, uAL1 := extractUniversityCSV1(universityCSV1)
	uL2, uAL2 := extractUniversityCSV2(universityCSV2)
	return append(uL1, uL2...), append(uAL1, uAL2...)
}

func UniversityMap(universityList []item.University) map[string]int64 {
	universityMap := make(map[string]int64)
	for _, university := range universityList {
		universityMap[university.Name] = university.ID
	}
	return universityMap
}
