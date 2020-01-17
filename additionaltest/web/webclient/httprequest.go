package webclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ClientGetData 取資料
func ClientGetData() {
	resp, err := http.PostForm("http://localhost", url.Values{"name": {"aaa"}, "age": {"20"}})

	if err != nil {
		fmt.Println("Request fail err = ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read Body data err = ", err)
		return
	}

	fmt.Println("resp = ", resp)
	fmt.Println("body = ", string(body))
}

// ClientPostData Post Data
func ClientPostData() {

}
