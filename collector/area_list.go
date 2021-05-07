// 数据来源https://github.com/modood/Administrative-divisions-of-China/blob/master/dist/pca-code.json
// 透传json的string

package collector

import "io/ioutil"

const AREA_LIST = "data/pca-code.json"

func CollectAreaJson() string {
	areaJson := ""
	bytes, err := ioutil.ReadFile(AREA_LIST)
	if err != nil {
		panic(err)
	}
	areaJson = string(bytes)
	return areaJson
}
