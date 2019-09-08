package main

import "fmt"

func main() {

	//map 是无序的
	m := map[string]string{
		"name":     "ccmouse",
		"course":   "golang",
		"site":     "imooc",
		"qualisty": "notbad",
	}
	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("getting values")
	courseName := m["course"]
	fmt.Println(courseName)
	caurseName := m["caurse"] //假如key写错了返回是一个空串
	fmt.Println(caurseName)

	//判断一个key 是否存在

	if caurseNamea, ok := m["caurse"]; ok {
		fmt.Println("key", caurseNamea)
	} else {
		fmt.Println("key不存在")
	}

	if courseNamea, ok := m["course"]; ok {
		fmt.Println("key", courseNamea)
	} else {
		fmt.Println("key不存在")
	}

	//删除一个元素
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")

	name, ok = m["name"]
	fmt.Println(name, ok)

}

// 创建 make(map[string]int)
// map是无序字典
