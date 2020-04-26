package main

// import "encoding/json"
// import "fmt"

// type Person struct{
// 	Name     string
// 	Email    string
// }
// type Person2 struct{
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// }

// type Person1 struct{
// 	Name     *string
// 	Email    *string
// }

// type int int64

// func(a Person2) test(){
// 	marshal, _ := json.Marshal(a)
// 	fmt.Println("json ",string(marshal))
// }

// //func test(a Person1){
// //	marshal, _ := json.Marshal(a)
// //	fmt.Println("json ",string(marshal))
// //}
// func test(a *Person1){
// 	marshal, _ := json.Marshal(a)
// 	fmt.Println("json ",string(marshal))
// }
// func(a *Person1) test(){
// 	w:="张易"
// 	a.Name=&w
// 	marshal, _ := json.Marshal(a)
// 	fmt.Println("json ",string(marshal))
// }
// func(a Person1) test2(){
// 	marshal, _ := json.Marshal(a)
// 	fmt.Println("json ",string(marshal))
// }
// func(per Person2) String() string{
// 	str:=fmt.Sprintf("Name=[%v],Email=[%v",per.Name,per.Email)
// 	return str
// }

// func main() {
// 	var p1 Person
// 	p1.Name="小明"
// 	p1.Email="zy84338719@hotmail.com"
// 	var p2 Person1
// 	p2.Name = &p1.Name
// 	zy:="12312"
// 	p2.Name =&zy
// 	var n  int64= 32
// 	var w int
// 	w=int(n)
// 	//var www Person1 = Person1(p1)
// 	fmt.Println(w)

// 	fmt.Println(p1,p2,*p2.Name,p2.Email)

// 	var p3 Person2 = Person2{"1","2"}
// 	//marshal, _ := json.Marshal(p3)
// 	//fmt.Println("json ",string(marshal))
// 	p3.test()
// 	var p4 Person2 = Person2(p1)
// 	fmt.Println(p4)

// 	//test(p2)
// 	test(&p2)
// 	p2.test()
// 	p2.test2()
// 	(&p2).test()
// 	(&p2).test2()

// }
import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port string
	flag.StringVar(&user, "u", "", "用户名默认为空")
	flag.StringVar(&pwd, "p", "", "密码默认为空")
	flag.StringVar(&host, "h", "localhost", "地址默认为空")
	flag.StringVar(&port, "port", "3306", "端口默认为3306")
	flag.Parse()
	fmt.Printf("user=%v pwd=%v host=%v post=%v", user, pwd, host, port)
}
