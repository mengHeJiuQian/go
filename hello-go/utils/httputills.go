package httputils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://www.baidu.com")
	fmt.Print(response)
	marshal, _ := json.Marshal(response)
	fmt.Println(marshal)

	//resp, _ := http.Get("http://localhost:8080/wechatmobile/api/v1/beneficiaryChg/queryLocation?policyNo=01263175")
	//data, _ := json.Marshal(resp)
	//fmt.Println()
}

func HttpGetHtml(url string) (htmlText string, error error) {
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
