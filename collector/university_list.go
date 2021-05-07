// 数据来源http://www.moe.gov.cn/jyb_xxgk/s5743/s5744/202007/t20200709_470937.html

package collector

import (
	"io/ioutil"
)

const UNIVERSITY_LIST1 = "data/university1.csv"
const UNIVERSITY_LIST2 = "data/university2.csv"

func CollectUniversityCSV1() string {
	f1, err := ioutil.ReadFile(UNIVERSITY_LIST1)
	if err != nil {
		panic(err)
	}
	return string(f1)
}

func CollectUniversityCSV2() string {
	f2, err := ioutil.ReadFile(UNIVERSITY_LIST2)
	if err != nil {
		panic(err)
	}
	return string(f2)
}
