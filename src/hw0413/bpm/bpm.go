package bpm

import (
	"log"

	"github.com/go-resty/resty/v2"
)

// 全域變數
var (
	//Public API網址
	server_addr = "https://jsonplaceholder.typicode.com/albums"
	//參數id
	id_num = "/1"
)

// Get
func Get_test() []byte {
	//新增一個第三方叫client
	client := resty.New()
	//向API伺服器發送一個Get的request並取得response跟error
	resp, err := client.R().
		//將預先設置的全域變數帶入(Public API網址)
		Get(server_addr)
	//如果err不是空值印出並退出程式
	if err != nil {
		//log 輸出內容
		//fatal 退出應用程式
		log.Fatal(err)
	}

	//回傳response的body
	return resp.Body()
}

// Post
func Post_test(body []byte) string {
	client := resty.New()
	//向API伺服器發送一個Post的request並取得response跟error
	resp, err := client.R().
		//設定request的資料格式為json
		SetHeader("Content-Type", "application/json").
		//傳送input body
		SetBody(body).
		//將預先設置的全域變數帶入(Public API網址)
		Post(server_addr)
	//如果err不是空值印出並退出程式
	if err != nil {
		//log 輸出內容
		//fatal 退出應用程式
		log.Fatal(err)
	}

	//將response的body轉為string後回傳
	return string(resp.Body())
}

// Patch
func Patch_test(body []byte) string {
	//新增一個第三方叫client
	client := resty.New()
	//向API伺服器發送一個Patch的request並取得response跟error
	resp, err := client.R().
		//設定request的資料格式為json
		SetHeader("Content-Type", "application/json").
		//傳送input body
		SetBody(body).
		//將預先設置的全域變數帶入(Public API網址+參數id)
		Patch(server_addr + id_num)
	//如果err不是空值印出並退出程式
	if err != nil {
		//log 輸出內容
		//fatal 退出應用程式
		log.Fatal(err)
	}

	//將response的body轉為string後回傳
	return string(resp.Body())
}

// Put
func Put_test(body []byte) string {
	//新增一個第三方叫client
	client := resty.New()
	//向API伺服器發送一個Put的request並取得response跟error
	resp, err := client.R().
		//設定request的資料格式為json
		SetHeader("Content-Type", "application/json").
		//傳送input body
		SetBody(body).
		//將預先設置的全域變數帶入(Public API網址+參數id)
		Put(server_addr + id_num)
	//如果err不是空值印出並退出程式
	if err != nil {
		//log 輸出內容
		//fatal 退出應用程式
		log.Fatal(err)
	}

	//將response的body轉為string後回傳
	return string(resp.Body())
}

// Delete
func Delete_test() []byte {
	//新增一個第三方叫client
	client := resty.New()
	//向API伺服器發送一個Delete的request並取得response跟error
	resp, err := client.R().
		//將預先設置的全域變數帶入(Public API網址+參數id)
		Delete(server_addr + id_num)
	//如果err不是空值印出並退出程式
	if err != nil {
		//log 輸出內容
		//fatal 退出應用程式
		log.Fatal(err)
	}

	//回傳response的body
	return resp.Body()
}
