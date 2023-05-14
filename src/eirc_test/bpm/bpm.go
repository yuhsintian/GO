package bpm

import (
	"log"

	"github.com/go-resty/resty/v2"
)

func Get_test() []byte {
	client := resty.New()
	resp, err := client.R().
		Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
		//log print
		//fatal 1.打印輸出內容 2.退出應用程序
	}

	return resp.Body()
}

func Post_test(body []byte) string {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(string(body)).
		Post("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	return string(resp.Body())
}
