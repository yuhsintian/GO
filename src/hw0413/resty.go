package main

import (
	"HW0413/bpm"
	"HW0413/structure"
	"encoding/json"
	"fmt"
)

func main() {
	get()
	post()
	patch()
	put()
	delete()
}

func get() {
	read := []structure.Read{}
	err := json.Unmarshal(bpm.Get_test(), &read)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(read)
}

//以下指令為json格式，看不懂，所以要用上面那段轉成struct格式成看得懂的
// func get() {
// 	read := []structure.Read{}
// 	err := json.Unmarshal(bpm.Get_test(), &read)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(bpm.Get_test())
// }

func post() {
	read := []structure.Read{}
	marshal, err := json.Marshal(read)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bpm.Post_test(marshal))
}

// patch
func patch() {
	//宣告read調用struct並賦值
	read := []structure.Read{}
	//將read序列化為json格式並取得error
	marshal, err := json.Marshal(read)
	//如果err不是空值印出並退出程式
	if err != nil {
		//印出error
		fmt.Println(err)
	}
	//將json格式的read傳入bpm.Patch_test()，並將response.body印出
	fmt.Println(bpm.Patch_test(marshal))
}

// put
func put() {
	//宣告read調用struct並賦值
	read := []structure.Read{}
	//將read序列化為json格式並取得error
	marshal, err := json.Marshal(read)
	//如果err不是空值印出並退出程式
	if err != nil {
		//印出error
		fmt.Println(err)
	}
	//將json格式的read傳入bpm.Put_test()，並將response.body印出
	fmt.Println(bpm.Put_test(marshal))
}

// delete
func delete() {
	//將bpm.Delete_test()的response.body轉為string並印出
	fmt.Println(string(bpm.Delete_test()))
}
