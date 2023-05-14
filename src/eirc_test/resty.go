package main

import (
	"eirc_test/bpm"
	"eirc_test/structure"
	"encoding/json"
	"fmt"
)

func main() {
	get()
	post()
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
	read := structure.Read{
		Title:  "hello world",
		Userid: 5,
	}
	marshal, err := json.Marshal(read)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bpm.Post_test(marshal))
}
