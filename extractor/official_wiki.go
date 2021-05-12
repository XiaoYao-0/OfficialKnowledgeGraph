// 根据官员姓名，爬取对应wiki上的词条信息，传出map(id, text)

package extractor

import (
	"OfficialKnowledgeGraph/collector"
	"OfficialKnowledgeGraph/item"
	"fmt"
	"github.com/antchfx/htmlquery"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func extractOfficialList(urlList []string) []collector.WikiItem {
	var wikiItemList []collector.WikiItem
	var id int64
	id = 0
	// var wg sync.WaitGroup
	// var mutex sync.Mutex
	for _, url := range urlList {
		id++
		// wg.Add(1)
		// defer wg.Done()
		wikiItem := collector.CollectOfficialWiki(url, id)
		if wikiItem.ID != 0 {
			wikiItemList = append(wikiItemList, wikiItem)
		}
		// mutex.Lock()
		// mutex.Unlock()
	}
	// wg.Wait()
	return wikiItemList
}

func ExtractOfficial(urlList []string, areaMap, universityMap map[string]int64) ([]item.Official, []item.OfficialArea, []item.OfficialUniversity, []item.OfficialPosition, []item.Position, []item.PositionArea) {
	var positionID int64
	positionID = 0
	wikiItemList := extractOfficialList(urlList)
	dict1, dict2 := collector.CollectPositionLevel()
	nations := []string{"汉", "满", "蒙古", "回", "藏", "维吾尔", "苗", "彝", "壮", "布依", "侗", "瑶", "白", "土家", "哈尼", "哈萨克", "傣", "黎", "傈僳", "佤", "畲", "高山", "拉祜", "水", "东乡", "纳西", "景颇", "柯尔克孜", "土", "达斡尔", "仫佬", "羌", "布朗", "撒拉", "毛南", "仡佬", "锡伯", "阿昌", "普米", "朝鲜", "塔吉克", "怒", "乌孜别克", "俄罗斯", "鄂温克", "德昂", "保安", "裕固", "京", "塔塔尔", "独龙", "鄂伦春", "赫哲", "门巴", "珞巴", "基诺"}
	var officialList []item.Official
	var officialAreaList []item.OfficialArea
	var officialUniversityList []item.OfficialUniversity
	var officialPositionList []item.OfficialPosition
	var positionList []item.Position
	var positionAreaList []item.PositionArea
	rYear, _ := regexp.Compile("(19|20)[0-9][0-9]")
	rBirthYear, _ := regexp.Compile("19[0-9][0-9]")
	logID := 0
	for _, wikiItem := range wikiItemList {
		logID++
		if logID%100 == 0 {
			fmt.Println(100*float64(logID)/float64(len(wikiItemList)), "%100")
		}
		// wg.Add(1)
		if wikiItem.ID == 0 {
			continue
		}
		official := item.Official{
			ID:   wikiItem.ID,
			Name: wikiItem.Name,
		}
		officialArea := item.OfficialArea{
			OfficialID: wikiItem.ID,
			AreaID:     0,
		}
		officialUniversity := item.OfficialUniversity{
			OfficialID: wikiItem.ID,
		}
		// gender
		if strings.Index(htmlquery.InnerText(wikiItem.Text1), "男") == -1 {
			if strings.Index(htmlquery.InnerText(wikiItem.Text1), "女") == -1 {
				official.Gender = 2
			} else {
				official.Gender = 1
			}
		} else {
			if strings.Index(htmlquery.InnerText(wikiItem.Text1), "女") == -1 {
				official.Gender = 0
			} else {
				if strings.Index(htmlquery.InnerText(wikiItem.Text1), "男") < strings.Index(htmlquery.InnerText(wikiItem.Text1), "女") {
					official.Gender = 0
				} else {
					official.Gender = 1
				}

			}
		}
		// birth year, nationality, area, university
		list1 := htmlquery.Find(wikiItem.Text2, "//dt[@class='basicInfo-item name']")
		list2 := htmlquery.Find(wikiItem.Text2, "//dd[@class='basicInfo-item value']")
		for i := 0; i < len(list1); i++ {
			k, v := htmlquery.InnerText(list1[i]), htmlquery.InnerText(list2[i])
			if strings.Contains(k, "民族") {
				official.Nationality = v
			}
			if strings.Contains(k, "出生日期") {
				if rBirthYear.FindString(v) != "" {
					birthYear, _ := strconv.Atoi(rBirthYear.FindString(v))
					official.BirthYear = birthYear
				}
			}
			if strings.Contains(k, "出生地") || strings.Contains(k, "籍贯") {
				var areaID int64
				for areaName, id := range areaMap {
					if id > areaID && strings.Contains(v, areaName) {
						areaID = id
					}
				}
				officialArea.AreaID = areaID
			}
			if strings.Contains(k, "校") {
				universityID := universityMap[v]
				if universityID != 0 {
					officialUniversity.UniversityID = universityID
				}
			}
		}
		text1 := htmlquery.InnerText(wikiItem.Text1)
		if official.Nationality == "" {
			var indexes []int
			for _, n := range nations {
				if i := strings.Index(text1, n); i != -1 {
					indexes = append(indexes, i)
				}
			}
			if len(indexes) != 0 {
				official.Nationality = nations[min(indexes)] + "族"
			}
		}
		if official.BirthYear == 0 {
			if rBirthYear.FindString(text1) != "" {
				birthYear, _ := strconv.Atoi(rBirthYear.FindString(text1))
				official.BirthYear = birthYear
			}
		}
		if officialArea.AreaID == 0 {
			var areaID int64
			for areaName, id := range areaMap {
				if id > areaID && strings.Contains(text1, areaName) {
					areaID = id
				}
			}
			officialArea.AreaID = areaID
		}
		if officialUniversity.UniversityID == 0 {
			for universityName, id := range universityMap {
				if strings.Contains(text1, universityName) {
					officialUniversity.UniversityID = id
					break
				}
			}
		}
		// position
		for _, n := range wikiItem.Text3 {
			officialPosition := item.OfficialPosition{
				OfficialID: wikiItem.ID,
			}
			startYear, endYear := 0, 0
			text := htmlquery.InnerText(n)
			years := rYear.FindAllString(text, 2)
			if len(years) == 0 {
				continue
			}
			startYear, _ = strconv.Atoi(years[0])
			if len(years) >= 2 {
				endYear, _ = strconv.Atoi(years[1])
			}
			positionID++
			officialPosition.StartYear = startYear
			officialPosition.EndYear = endYear
			officialPosition.PositionID = positionID
			positionName := ""
			for i, r := range text {
				if r == 12288 || r == ',' || r == '，' || r == ' ' {
					positionName = text[i+len(string(r)):]
					break
				}
			}
			runes := []rune(positionName)
			for i := len(runes) - 1; i >= 0; i-- {
				if unicode.Is(unicode.Han, runes[i]) {
					positionName = string(runes[:i+1])
					break
				}
			}
			if positionName == "" {
				continue
			}
			position, positionAreas := extractPosition(positionName, positionID, dict1, dict2, areaMap)
			officialPositionList = append(officialPositionList, officialPosition)
			positionList = append(positionList, position)
			positionAreaList = append(positionAreaList, positionAreas...)
		}
		officialList = append(officialList, official)
		officialAreaList = append(officialAreaList, officialArea)
		officialUniversityList = append(officialUniversityList, officialUniversity)
		//var wg sync.WaitGroup
		//var mutex1, mutex2 sync.Mutex
		//rBirthYear, _ := regexp.Compile("19[0-9][0-9]")
		//for _, wikiItem := range wikiItemList {
		//	wg.Add(1)
		//	go func(wikiItem collector.WikiItem) {
		//		defer wg.Done()
		//		if wikiItem.ID == 0 {
		//			return
		//		}
		//		official := item.Official{
		//			ID:   wikiItem.ID,
		//			Name: wikiItem.Name,
		//		}
		//		officialArea := item.OfficialArea{
		//			OfficialID: wikiItem.ID,
		//			AreaID:     0,
		//		}
		//		officialUniversity := item.OfficialUniversity{
		//			OfficialID: wikiItem.ID,
		//		}
		//		// gender
		//		if strings.Index(htmlquery.InnerText(wikiItem.Text1), "男") == -1 {
		//			if strings.Index(htmlquery.InnerText(wikiItem.Text1), "女") == -1 {
		//				official.Gender = 2
		//			} else {
		//				official.Gender = 1
		//			}
		//		} else {
		//			if strings.Index(htmlquery.InnerText(wikiItem.Text1), "女") == -1 {
		//				official.Gender = 0
		//			} else {
		//				if strings.Index(htmlquery.InnerText(wikiItem.Text1), "男") < strings.Index(htmlquery.InnerText(wikiItem.Text1), "女") {
		//					official.Gender = 0
		//				} else {
		//					official.Gender = 1
		//				}
		//
		//			}
		//		}
		//		// birth year, nationality, area, university
		//		list1 := htmlquery.Find(wikiItem.Text2, "//dt[@class=basicInfo-item name]")
		//		list2 := htmlquery.Find(wikiItem.Text2, "//dd[@class=basicInfo-item value]")
		//		for i := 0; i < len(list1); i++ {
		//			k, v := htmlquery.InnerText(list1[i]), htmlquery.InnerText(list2[i])
		//			if strings.Contains(k, "民族") {
		//				official.Nationality = v
		//			}
		//			if strings.Contains(k, "出生日期") {
		//				if rBirthYear.FindString(v) != "" {
		//					birthYear, _ := strconv.Atoi(rBirthYear.FindString(v))
		//					official.BirthYear = birthYear
		//				}
		//			}
		//			if strings.Contains(k, "出生地") || strings.Contains(k, "籍贯") {
		//				var areaID int64
		//				for areaName, id := range areaMap {
		//					if id > areaID && strings.Contains(v, areaName) {
		//						areaID = id
		//					}
		//				}
		//				officialArea.AreaID = areaID
		//			}
		//			if strings.Contains(k, "校") {
		//				universityID := universityMap[v]
		//				if universityID != 0 {
		//					officialUniversity.UniversityID = universityID
		//				}
		//			}
		//		}
		//		text1 := htmlquery.InnerText(wikiItem.Text1)
		//		if official.Nationality == "" {
		//			var indexes []int
		//			for _, n := range nations {
		//				if i := strings.Index(text1, n); i != -1 {
		//					indexes = append(indexes, i)
		//				}
		//			}
		//			if len(indexes) != 0 {
		//				official.Nationality = nations[min(indexes)] + "族"
		//			}
		//		}
		//		if official.BirthYear == 0 {
		//			if rBirthYear.FindString(text1) != "" {
		//				birthYear, _ := strconv.Atoi(rBirthYear.FindString(text1))
		//				official.BirthYear = birthYear
		//			}
		//		}
		//		if officialArea.AreaID == 0 {
		//			var areaID int64
		//			for areaName, id := range areaMap {
		//				if id > areaID && strings.Contains(text1, areaName) {
		//					areaID = id
		//				}
		//			}
		//			officialArea.AreaID = areaID
		//		}
		//		if officialUniversity.UniversityID == 0 {
		//			for universityName, id := range universityMap {
		//				if strings.Contains(text1, universityName) {
		//					officialUniversity.UniversityID = id
		//					break
		//				}
		//			}
		//		}
		//		// position
		//		for _, n := range wikiItem.Text3 {
		//			officialPosition := item.OfficialPosition{
		//				OfficialID: wikiItem.ID,
		//			}
		//			startYear, endYear := 0, 0
		//			text := htmlquery.InnerText(n)
		//			years := rYear.FindAllString(text, 2)
		//			if len(years) == 0 {
		//				continue
		//			}
		//			startYear, _ = strconv.Atoi(years[0])
		//			if len(years) >= 2 {
		//				endYear, _ = strconv.Atoi(years[1])
		//			}
		//			positionID++
		//			officialPosition.StartYear = startYear
		//			officialPosition.EndYear = endYear
		//			officialPosition.PositionID = positionID
		//			positionName := ""
		//			for i, r := range text {
		//				if r == 12288 || r == ',' || r == '，' || r == ' ' {
		//					positionName = text[i+len(string(r)):]
		//					break
		//				}
		//			}
		//			runes := []rune(positionName)
		//			for i := len(positionName) - 1; i >= 0; i-- {
		//				if unicode.Is(unicode.Han, runes[i]) {
		//					positionName = string(runes[:i+1])
		//					break
		//				}
		//			}
		//			if positionName == "" {
		//				return
		//			}
		//			position, positionAreas := extractPosition(positionName, positionID, dict1, dict2, areaMap)
		//			mutex2.Lock()
		//			officialPositionList = append(officialPositionList, officialPosition)
		//			positionList = append(positionList, position)
		//			positionAreaList = append(positionAreaList, positionAreas...)
		//			mutex2.Unlock()
		//		}
		//		mutex1.Lock()
		//		officialList = append(officialList, official)
		//		officialAreaList = append(officialAreaList, officialArea)
		//		officialUniversityList = append(officialUniversityList, officialUniversity)
		//		mutex1.Unlock()
		//	}(wikiItem)

	}

	return officialList, officialAreaList, officialUniversityList, officialPositionList, positionList, positionAreaList
}

func min(arr []int) int {
	min := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}
