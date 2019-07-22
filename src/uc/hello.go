package main

import (
	"fmt"
	"math"
	"net/http"
	"reflect"
	"sync"
	"time"
	"unsafe"
)

var (
	counter = 0
	//互斥锁 使按顺序访问代码
	lock sync.Mutex
)

func main() {

	now := time.Now().UTC()

	fmt.Println(now)

	format := now.Format("2006-01-02 15:04:05")

	fmt.Println(format)

	//exercise slice
	//todo slice的容量能超过相关数组的长度吗？
	var arr1 [6] int
	var slice1 = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	var number = make([]int, 3, 5)

	fmt.Println(number)
	fmt.Println(cap(number))
	fmt.Println(len(number))

	var slice2 = make([]int, 4)

	slice2[0] = 1;
	slice2[1] = 2;
	slice2[2] = 3;
	fmt.Println("slice2 value", slice2)

	for ix, value := range slice2 {
		fmt.Println(ix, value)
	}

	//map
	var mapdemo = make(map[int]string)

	mapdemo[1] = "1"
	mapdemo[3] = "3"

	fmt.Println("origin", mapdemo)
	//删除map元素
	delete(mapdemo, 1)

	//judge key——value is exist  amaze grammar
	_, ok := mapdemo[1]

	fmt.Println(ok)

	fmt.Println("result", mapdemo)

	//for-range 奇怪的语法

	//map类型的切片 想获取map类型的切片 必须使用两次make函数 第一次分配切片 第二次分配slice中的每个map元素

	fmt.Println(unsafe.Sizeof(int8(127)))

	fmt.Println(unsafe.Sizeof(int64(127)))

	fmt.Println(math.Pi)
	//not necessary all elements in array have to be assigned a value during short hand declaration
	a := [3] int{12}
	fmt.Println(a)

	a1 := [] int{2, 4, 5}

	//这样申明无意义呀
	// var b []int = a[1:2]
	fmt.Println(a1)

	//method one to create slice
	// var b = a1[0:2]
	//method two to create slice  returns a slice reference
	c := []int{1, 2}

	d := [] int{2, 3, 4}

	of := reflect.TypeOf(d)
	ofa1 := reflect.TypeOf(a1)

	fmt.Println(c)

	fmt.Println(of)

	fmt.Println(ofa1)

	fmt.Println("start a xiecheng")

	var arr4 [5] int

	arr4 = [5] int{1, 2, 4, 5, 6}

	arr4[0] = 3

	fmt.Println(arr4)

	/**
	go协程类似一个线程，但是go协程是由go自己调度，而不是os，在协程中的代码可以和其他代码并发执行，go协程
	 */
	go process()

	fmt.Println("done")

	fmt.Println("--------------------------------------------")

	for i := 0; i < 2; i++ {
		go incr()

	}

	p := person{"czg", 12, "female"}

	pp := &p

	pp.name = "pp"
	fmt.Println(pp)

	fmt.Println(p.name)
	p.name = "xx"

	fmt.Println(p)

	fmt.Println(reflect.TypeOf(p))

	fmt.Println("++++++++++++++++++++++")

	var ani animal
	ani = Snake{Poisonous: true}
	fmt.Println(ani.description())

	ani = cat{Sound: "miaomiao"}
	fmt.Println(ani.description())

	resp, err := http.Get("https://www.baidu.com/")

	if err != nil {

		fmt.Println(err)

	}

	fmt.Println(resp)

	//chanel
	cha := make(chan string)

	go func() { cha <- "hello" }()

	msg := <-cha

	fmt.Println(msg)

	ch := make(chan string)

	go sc(ch)

}

func sc(ch chan<- string) {

	ch <- "hello"

}

func process() {
	fmt.Println("process")
}

func incr() {

	lock.Lock()
	defer lock.Unlock()
	counter ++
	fmt.Println(counter)
}

type person struct {
	name   string
	age    int
	gender string
}

type animal interface {
	description() string
}

type cat struct {
	Type  string
	Sound string
}
type Snake struct {
	Type      string
	Poisonous bool
}

func (s Snake) description() string {

	return fmt.Sprintf("poisonous: %v", s.Poisonous)

}

func (c cat) description() string {

	return fmt.Sprint("sound:", c.Sound)

}
