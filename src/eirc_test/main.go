package main

import (
	"eirc_test/structure"
	"encoding/json"
	"fmt"
)

func main() {
	//使用structure方法1
	student1 := structure.Students{
		Name: "amy",
		ID:   1,
	}

	//使用structure方法2
	student2 := structure.Students{}
	student2.Name = "Bin"
	student2.ID = 2

	//放入class陣列並印出
	class := []structure.Students{student1, student2}
	fmt.Println(class)

	//將class的東西放入class2
	//index是陣列中的位子
	//value是student1裡的值
	//append類似+=
	class2 := []structure.Students{}
	for index, value := range class {
		class2 = append(class2, value)
		fmt.Println(index, class2)
	}

	//序列化struct轉json
	//marshal是[]byte
	marshal, err := json.Marshal(class)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

	//反序列化json轉struct
	class3 := []structure.Students{}
	err = json.Unmarshal(marshal, &class3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(class3)
}
