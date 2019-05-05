package HttpModule

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Root(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()   //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path:", r.URL.Path)
	fmt.Println("scheme:", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println(r.URL.User)
	_, _ = w.Write([]byte("hello"))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal("r.ParseForm(): ", err)
	}
}

func Get(url string, header map[string]string) (resp *http.Response, err error) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	return client.Do(request)
}
