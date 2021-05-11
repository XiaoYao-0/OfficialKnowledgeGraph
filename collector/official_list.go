// 读取官员姓名列表（未消重）

package collector

import (
	"bufio"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

const OFFCIAL_LIST = "data/official-list.csv"
const MULTI_ITEM = "这是一个<a href=\"[/view/10812277.htm](https://baike.baidu.com/view/10812277.htm)\" target=\"_blank\">多义词</a>，请在下列<a href=\"[/view/340519.htm](https://baike.baidu.com/view/340519.htm)\" target=\"_blank\">义项</a>上选择浏览"
const BAIDU_WIKI = "https://baike.baidu.com/item/"
const OFFICIAL_WIKI = "人物履历"

func CollectOfficialURLList() []string {
	f, err := ioutil.ReadFile(OFFCIAL_LIST)
	if err != nil {
		panic(err)
	}
	var urlList []string
	set := make(map[string]bool)
	reader := strings.NewReader(string(f))
	scanner := bufio.NewScanner(reader)
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for scanner.Scan() {
		s := scanner.Text()
		name := strings.Split(s, ",")[5]
		if set[name] == true {
			continue
		}
		set[name] = true
		url := BAIDU_WIKI + name
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				resp, err = http.Get(url)
				if err != nil {
					fmt.Println("Can't get url: ", url)
					return
				}
			}
			bytes, _ := ioutil.ReadAll(resp.Body)
			text := string(bytes)
			if !strings.Contains(text, MULTI_ITEM) {
				mutex.Lock()
				urlList = append(urlList, url)
				mutex.Unlock()
				return
			}
			doc, err := htmlquery.LoadURL(url)
			if err != nil {
				doc, err = htmlquery.LoadURL(url)
				if err != nil {
					fmt.Println("Can't load url: ", url)
					return
				}
			}
			list := htmlquery.Find(doc, "//div[@class='para']/a/@href")
			for _, n := range list {
				url = htmlquery.SelectAttr(n, "href")
				wg.Add(1)
				go func(url string) {
					defer wg.Done()
					resp, err := http.Get(url)
					if err != nil {
						resp, err = http.Get(url)
						if err != nil {
							fmt.Println("Can't get url: ", url)
							return
						}
					}
					bytes, _ := ioutil.ReadAll(resp.Body)
					text := string(bytes)
					if strings.Contains(text, OFFICIAL_WIKI) {
						mutex.Lock()
						urlList = append(urlList, url)
						mutex.Unlock()
					}
				}(url)
			}
		}(url)
	}
	wg.Wait()
	return urlList
}
