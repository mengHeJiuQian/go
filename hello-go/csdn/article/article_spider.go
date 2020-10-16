package article

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	html, _ := httpGetHtml(beginUrl)

}

func httpGetHtml(url string) (htmlText string, error error) {
	resp, respErr := http.Get(url)
	if respErr != nil {
		error = respErr
		fmt.Println("抓取网页出现问题，url=" + url)
		return
	}

	defer resp.Body.Close()

	bodyByte, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(bodyByte)
	return bodyStr, error
}
