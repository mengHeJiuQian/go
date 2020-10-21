package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	// "github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
	"hello-go/csdn/model"
	httputils "hello-go/utils"
)

/**
 * 通过文章链接抓取文章得一些信息，如发布日期，标题等
 */

const mongoServer = "localhost"
const mongoPort = "27017"
const mongoDatabase = "csdn_article"
const mongoCollection = "article_base"

func main() {

	beginUrl := "https://blog.csdn.net/wei242425445/article/details/88417407"
	if !strings.Contains(beginUrl, "/article/details/") {
		fmt.Println(beginUrl + "不是一个文章链接，跳过")
		return
	}

	// 初始化colly
	c := colly.NewCollector()

	htmlText, _ := httputils.HttpGetHtml(beginUrl)
	cArticleInfo := filterArticleInfo(htmlText)
	fmt.Print(cArticleInfo)
}

/*compile := regexp.MustCompile(csdnUrlRe)
allSubmatch := compile.FindAllStringSubmatch(contents, -1)

urlList := list.New()

for _, submatch := range allSubmatch {
	var urlWithShit = submatch[0]
	index := strings.Index(urlWithShit, "\"")
	index++
	urlWithShit = urlWithShit[index:]

	url := urlWithShit[0 : len(urlWithShit)-1]
	urlList.PushBack(url)
}*/
// 提取出csdn文件的一些内容
func filterArticleInfo(htmlText string) *model.CArticleInfo {
	fmt.Println(htmlText)

	return &model.CArticleInfo{}
}
