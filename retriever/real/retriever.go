package real

import (
	"io"
	"net/http"
	"os"
	"time"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r Retriever) Get(url string) string {
	//resp, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//result, err := httputil.DumpRequest(resp, true)
	//resp.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//return string(result)

	client := &http.Client{}

	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	//处理返回结果
	response, _ := client.Do(reqest)
	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)

	//返回的状态码
	status := response.StatusCode

	return string(status)
}
