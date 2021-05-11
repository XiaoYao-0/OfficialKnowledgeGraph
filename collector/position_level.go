// 爬取相应页面，职位级别对应表https://zhuanlan.zhihu.com/p/47577228，传出map(职位, 等级)

package collector

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

const POSITION_LEVEL_LIST1 = "data/position-level1.txt"
const POSITION_LEVEL_LIST2 = "data/position-level2.txt"

func collectPositionLevel1() map[string]int {
	dict := make(map[string]int)
	bytes, err := ioutil.ReadFile(POSITION_LEVEL_LIST1)
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader(string(bytes))
	scanner := bufio.NewScanner(reader)
	level := -1
	for scanner.Scan() {
		s := scanner.Text()
		switch {
		case s == "":
			{
				continue
			}
		case strings.HasPrefix(s, "#"):
			{
				level++
			}
		case strings.HasPrefix(s, "!"):
			{
				split := strings.Split(s, "：")
				level, _ := strconv.Atoi(split[1])
				dict[split[0][1:]] = level
			}
		default:
			split := strings.Split(s, "、")
			for _, i := range split {
				dict[i] = level
			}
		}
	}
	return dict
}

func collectPositionLevel2() map[string]int {
	dict := make(map[string]int)
	bytes, err := ioutil.ReadFile(POSITION_LEVEL_LIST2)
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader(string(bytes))
	scanner := bufio.NewScanner(reader)
	level := 1
	for scanner.Scan() {
		s := scanner.Text()
		switch {
		case strings.HasPrefix(s, "#"):
			{
				level--
			}
		default:
			split := strings.Split(s, "、")
			for _, i := range split {
				dict[i] = level
			}
		}
	}
	return dict
}

// 第一个是直接匹配的，如国家主席对应正国级；第二个为已知Area时，相对于Area的级别
func CollectPositionLevel() (map[string]int, map[string]int) {
	return collectPositionLevel1(), collectPositionLevel2()
}
