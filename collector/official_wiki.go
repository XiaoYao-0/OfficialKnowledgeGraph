// 根据官员姓名，爬取对应wiki上的词条信息，传出map(id, text)

package collector

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"strings"
)

type WikiItem struct {
	ID    int64
	Name  string
	Text1 *html.Node
	Text2 *html.Node
	Text3 []*html.Node
}

func CollectOfficialWiki(url string, id int64) WikiItem {
	var wikiItem WikiItem
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		doc, err = htmlquery.LoadURL(url)
		if err != nil {
			fmt.Println("Can't load url: ", url)
			return WikiItem{}
		}
	}
	wikiItem.ID = id
	// name
	split := strings.Split(url, "/")
	flag := 0
	for _, s := range split {
		if flag == 1 {
			wikiItem.Name = s
			break
		}
		if s == "item" {
			flag = 1
		}
	}
	// text1
	list := htmlquery.Find(doc, "//meta[@name='description']/@content")
	if len(list) != 1 {
		fmt.Println("CollectOfficialWiki url: ", url, "error")
		return WikiItem{}
	}
	wikiItem.Text1 = list[0]
	// text2
	list = htmlquery.Find(doc, "//div[@class='basic-info cmn-clearfix']")
	if len(list) != 1 {
		fmt.Println("CollectOfficialWiki url: ", url, "error")
		return WikiItem{}
	}
	wikiItem.Text2 = list[0]
	// text3
	list = htmlquery.Find(doc, "//div[@class='para-title level-2']/h2[text()='人物履历']")
	index := -1
	for i, n := range list {
		if htmlquery.InnerText(n) == "人物履历" {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("CollectOfficialWiki url: ", url, "error")
		return WikiItem{}
	}
	list = htmlquery.Find(doc, "//div[@class='para-title level-2']/h2[text()='人物履历']/parent::div/following-sibling::div[count(.|//div[@class='para-title level-2']/h2[text()='人物履历']/parent::div/following-sibling::div[@class='para-title level-2' and 1]/preceding-sibling::div) = count(//div[@class='para-title level-2']/h2[text()='人物履历']/parent::div/following-sibling::div[@class='para-title level-2' and 1]/preceding-sibling::div)]")
	wikiItem.Text3 = list

	return wikiItem
}
