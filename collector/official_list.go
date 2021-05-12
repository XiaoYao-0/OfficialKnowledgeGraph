// 读取官员姓名列表（未消重）

package collector

import (
	"bufio"
	"fmt"
	"github.com/antchfx/htmlquery"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const OFFiCIAL_LIST = "data/official-list.csv"
const MULTI_ITEM = "这是一个<a href=\"/view/10812277.htm\" target=\"_blank\">多义词</a>，请在下列<a href=\"/view/340519.htm\" target=\"_blank\">义项</a>上选择浏览"
const BAIDU_WIKI_ITEM = "https://baike.baidu.com/item/"
const BAIDU_WIKI = "https://baike.baidu.com"
const OFFICIAL_WIKI = "人物履历"
const COUNT_OFFICIAL = 1000

func SplitFile() []string {
	var files []string
	f, err := ioutil.ReadFile(OFFiCIAL_LIST)
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader(string(f))
	scanner := bufio.NewScanner(reader)
	csvText := ""
	i := 0
	for scanner.Scan() {
		i++
		if i%COUNT_OFFICIAL == 0 {
			filename := fmt.Sprintf("data/official-list-%d", (i-1)/COUNT_OFFICIAL)
			_, err := os.Create(filename)
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(filename, []byte(csvText), os.ModePerm)
			if err != nil {
				panic(err)
			}
			csvText = ""
			files = append(files, filename)
		}
		s := scanner.Text()
		csvText += s + "\n"
	}
	if csvText != "" {
		filename := fmt.Sprintf("data/official-list-%d", (i-1)/COUNT_OFFICIAL)
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(filename, []byte(csvText), os.ModePerm)
		if err != nil {
			panic(err)
		}
		csvText = ""
		files = append(files, filename)
	}
	return files
}

func CollectOfficialURLList(filepath string) []string {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var urlList []string
	set := make(map[string]bool)
	reader := strings.NewReader(string(f))
	scanner := bufio.NewScanner(reader)
	// var wg sync.WaitGroup
	// var mutex sync.Mutex
	logID := 0
	for scanner.Scan() {
		logID++
		if logID%100 == 0 {
			fmt.Printf("%s %d%100", filepath, int(100*float64(logID)/float64(COUNT_OFFICIAL)))
			fmt.Println("")
		}
		s := scanner.Text()
		name := strings.Split(s, ",")[5]
		for strings.Contains(name, "（") {
			name = strings.Split(name, "（")[0]
		}
		name = strings.ReplaceAll(name, " ", "")
		name = strings.ReplaceAll(name, " ", "")
		if set[name] == true {
			continue
		}
		set[name] = true
		url := BAIDU_WIKI_ITEM + name + "?force=1"

		resp, err := http.Get(url)
		if err != nil {
			resp, err = http.Get(url)
			if err != nil {
				fmt.Println("！！！Can't get url: ", url, err)
				continue
			}
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		text := string(bytes)
		if !strings.Contains(text, MULTI_ITEM) {
			urlList = append(urlList, url)
			continue
		}
		doc, err := htmlquery.LoadURL(url)
		if err != nil {
			doc, err = htmlquery.LoadURL(url)
			if err != nil {
				fmt.Println("Can't load url: ", url)
				continue
			}
		}
		list := htmlquery.Find(doc, "//div[@class='para']/a/@href")
		for _, n := range list {
			url = BAIDU_WIKI + htmlquery.SelectAttr(n, "href")
			resp, err := http.Get(url)
			if err != nil {
				resp, err = http.Get(url)
				if err != nil {
					fmt.Println("Can't get url: ", url)
					continue
				}
			}
			bytes, _ := ioutil.ReadAll(resp.Body)
			text := string(bytes)
			if strings.Contains(text, OFFICIAL_WIKI) {
				urlList = append(urlList, url)
			}
		}
		// wg.Add(1)
		//		defer wg.Done()
		//go func(url string) {
		//	defer wg.Done()
		//	resp, err := http.Get(url)
		//	if err != nil {
		//		resp, err = http.Get(url)
		//		if err != nil {
		//			fmt.Println("！！！Can't get url: ", url, err)
		//			return
		//		}
		//	}
		//	bytes, _ := ioutil.ReadAll(resp.Body)
		//	text := string(bytes)
		//	if !strings.Contains(text, MULTI_ITEM) {
		//		mutex.Lock()
		//		urlList = append(urlList, url)
		//		mutex.Unlock()
		//		return
		//	}
		//	doc, err := htmlquery.LoadURL(url)
		//	if err != nil {
		//		doc, err = htmlquery.LoadURL(url)
		//		if err != nil {
		//			fmt.Println("Can't load url: ", url)
		//			return
		//		}
		//	}
		//	list := htmlquery.Find(doc, "//div[@class='para']/a/@href")
		//	for _, n := range list {
		//		url = htmlquery.SelectAttr(n, "href")
		//		wg.Add(1)
		//		go func(url string) {
		//			defer wg.Done()
		//			resp, err := http.Get(url)
		//			if err != nil {
		//				resp, err = http.Get(url)
		//				if err != nil {
		//					fmt.Println("Can't get url: ", url)
		//					return
		//				}
		//			}
		//			bytes, _ := ioutil.ReadAll(resp.Body)
		//			text := string(bytes)
		//			if strings.Contains(text, OFFICIAL_WIKI) {
		//				mutex.Lock()
		//				urlList = append(urlList, url)
		//				mutex.Unlock()
		//			}
		//		}(url)
		//	}
		//}(url)
	}
	// wg.Wait()
	return urlList
}
