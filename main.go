package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/88250/gulu"
	"github.com/parnurzeal/gorequest"
)

var logger = gulu.Log.NewLogger(os.Stdout)

func main() {
	result := map[string]interface{}{}
	response, data, errors := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Get("https://hacpai.com/api/v2/user/88250/events?size=8").Timeout(7*time.Second).
		Set("User-Agent", "Profile Bot; +https://github.com/88250/88250").EndStruct(&result)
	if nil != errors || http.StatusOK != response.StatusCode {
		logger.Fatalf("fetch events failed: %+v, %s", errors, data)
	}
	if 0 != result["code"].(float64) {
		logger.Fatalf("fetch events failed: %s", data)
	}

	buf := &bytes.Buffer{}
	buf.WriteString("\n\n")
	for _, event := range result["data"].([]interface{}) {
		evt := event.(map[string]interface{})
		title := evt["title"].(string)
		url := evt["url"].(string)
		content := evt["content"].(string)
		buf.WriteString("* [" + title + "](" + url + ")：" + content + "\n")
	}
	buf.WriteString("\n")

	fmt.Println(buf.String())

	readme, err := ioutil.ReadFile("README.md")
	if nil != err {
		logger.Fatalf("read README.md failed: %s", data)
	}

	startFlag := []byte("<!--events start -->")
	start := bytes.Index(readme, startFlag)
	beforeStart := readme[:start+len(startFlag)]
	endFlag := []byte("<!--events end -->")
	end := bytes.Index(readme, endFlag)
	afterEnd := readme[end:]
	readme = append(beforeStart, buf.Bytes()...)
	readme = append(readme, afterEnd...)
	readme = bytes.ReplaceAll(readme, []byte("${events}"), buf.Bytes())
	if err := ioutil.WriteFile("README.md", readme, 0644); nil != err {
		logger.Fatalf("write README.md failed: %s", data)
	}
}