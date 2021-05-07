// 解析地域json，解析出area列表以及area-area关系组

package extractor

import (
	"OfficialKnowledgeGraph/item"
	"encoding/json"
)

type AreaJsonStruct []struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Children []struct {
		Code     string `json:"code"`
		Name     string `json:"name"`
		Children []struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"children"`
	} `json:"children"`
}

func phase(areaJson string) AreaJsonStruct {
	var areaJsonStruct AreaJsonStruct
	err := json.Unmarshal([]byte(areaJson), &areaJsonStruct)
	if err != nil {
		panic(err)
	}
	return areaJsonStruct
}

func extractAreaStruct(areaJsonStruct AreaJsonStruct) ([]item.Area, []item.AreaArea) {
	var areaList []item.Area
	var areaAreaList []item.AreaArea
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	id := 0
	root := item.Area{
		ID:    0,
		Name:  "中国",
		Level: 0,
	}
	areaList = append(areaList, root)
	id++
	for _, i := range areaJsonStruct {
		areaList = append(areaList, item.Area{
			ID:    id,
			Name:  i.Name,
			Level: 1,
		})
		map1[i.Code] = id
		areaAreaList = append(areaAreaList, item.AreaArea{
			RootAreaID:  0,
			ChildAreaID: id,
		})
		id++
	}
	for _, i := range areaJsonStruct {
		for _, j := range i.Children {
			if j.Name == "市辖区" {
				for _, k := range j.Children {
					areaList = append(areaList, item.Area{
						ID:    id,
						Name:  k.Name,
						Level: 2,
					})
					map2[k.Code] = id
					areaAreaList = append(areaAreaList, item.AreaArea{
						RootAreaID:  map1[i.Code],
						ChildAreaID: id,
					})
					id++
				}
			} else {
				areaList = append(areaList, item.Area{
					ID:    id,
					Name:  j.Name,
					Level: 2,
				})
				map2[j.Code] = id
				areaAreaList = append(areaAreaList, item.AreaArea{
					RootAreaID:  map1[i.Code],
					ChildAreaID: id,
				})
				id++
			}
		}
	}
	for _, i := range areaJsonStruct {
		for _, j := range i.Children {
			if j.Name != "市辖区" {
				for _, k := range j.Children {
					areaList = append(areaList, item.Area{
						ID:    id,
						Name:  k.Name,
						Level: 3,
					})
					areaAreaList = append(areaAreaList, item.AreaArea{
						RootAreaID:  map2[j.Code],
						ChildAreaID: id,
					})
					id++
				}
			}
		}
	}
	return areaList, areaAreaList
}

func ExtractArea(areaJson string) ([]item.Area, []item.AreaArea) {
	return extractAreaStruct(phase(areaJson))
}
