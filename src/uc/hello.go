package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
     _"github.com/go-sql-driver/mysql"
)


var(
	name string
	age int
)
func main() {

	var a int

	a = 19
	//这是编码成了utf-8的text 是个什么鬼 打印到制台是一个带？的方框
	s := string(a)

	r, size := utf8.DecodeRuneInString(s)




	fmt.Println("value is:", a)


	fmt.Println("the decode result is",r)
	fmt.Println("the decode size is",size)

	myPointer := getPointer()

	fmt.Println("the address is",myPointer)
	fmt.Println("the value is",*myPointer)


	slice2 := []byte {'h','e','l','l','o'}

	fmt.Println(slice2)

	entry:=[]string{"jack","bob"}

	for i,val:=range entry{

		fmt.Printf("At position %d, the character %s is present\n", i, val)

	}



	fmt.Println("-------------------------------------------|")

	//fmt.Println("enter your name: ")
	//fmt.Scanln(&name)
	//fmt.Println("enter your age: ")
	//fmt.Scanln(&age)
	//
	//fmt.Printf("hi %s %d \n",name,age)


	ch :=make(chan  string)

	go sendData(ch)

	go getData(ch)

	time.Sleep(2e9)


	//
	//http.HandleFunc("/hello", sayhelloName) //设置访问的路由
	//http.HandleFunc("/login",login)
	//err := http.ListenAndServe(":9090", nil) //设置监听的端口
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}



	//数据库操作
	db,err :=sql.Open("mysql","root:ai9oMsUan,D9@tcp(127.0.0.1:3306)/sale")
	checkErr(err)

	stmt, e := db.Prepare("insert into userinfo set username='czg',department='oppo',created='czg'")
	checkErr(e)

	fmt.Println("the result is ",stmt)


}

func getPointer()  ( myPointer *int)  {
	a :=234
	return &a
}

func sendData(ch chan string)  {

	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"

}

func getData(ch chan string)  {

	var input string

	time.Sleep(2e9)

	for   {

		input= <-ch

		fmt.Printf("%s ", input)

	}

}

func sayhelloName(w http.ResponseWriter ,r * http.Request)  {

	r.ParseForm()

	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的



}

func checkErr(err error)  {

	if err!=nil{

		panic(err)

	}



}

func login( w http.ResponseWriter,r *http.Request)  {
	fmt.Println("method:",r.Method)

	if(r.Method== "GET") {
		t,_:=template.ParseFiles("login.gtpl")

		log.Println(t.Execute(w,nil))
	}else{

		r.ParseForm()
		fmt.Println("username:",r.Form["username"])

		fmt.Println("password",r.Form["password"])


	}
}