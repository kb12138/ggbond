package map1

import "fmt"

func main() {
	var map11 = make(map[string]string)

	//map["zhangsanan"] = "23";
	map11["lisi"] = "24"
	fmt.Println(map11)

	//fmt.Println(map11["zhangsan"])
	fmt.Println(map11["lisi"])

}
