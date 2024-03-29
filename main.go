package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"mygo/mongo"
	"mygo/mysql"
	"mygo/redis"
	"net/http"
	"os"
	"strings"
)

func main() {

	//rand.Seed(2)
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(3))
	}

	var sb strings.Builder
	sb.WriteString("bbbbbb")
	sb.WriteRune('a')
	fmt.Println(sb.String())
	bb := bytes.Buffer{}
	bb.WriteString("中国人")
	fmt.Println("bb:", bb.String())
	/*fmt.Println("1234")
	urlString := "%E6%82%A8%E7%9A%84%E6%89%8B%E6%9C%BA%E9%AA%8C%E8%AF%81%E7%A0%81%EF%BC%9A595386%EF%BC%8C3%E5%88%86%E9%92%9F%E4%B9%8B%E5%86%85%E8%BE%93%E5%85%A5%E6%9C%89%E6%95%88%EF%BC%8C%E8%B0%A2%E8%B0%A2%E3%80%82"
	fmt.Println(url.QueryUnescape(urlString))

	fmt.Println(mystd.MaxLenStringNoRepSubstr("abcabcbb"))
	fmt.Println(mystd.MaxLenStringNoRepSubstr("我人人人人我"))
	fmt.Println(mystd.MaxLenStringNoRepSubstr("abbbba"))
	//fmt.Println("IntToString=",mystd.IntToString(65))
	fmt.Println("isPrime=", mystd.IsPrime(13))
	arr := []int{1, 2, 3, 4, 5, 6, 20, 30, 45, 201, 1000, 822}
	fmt.Println("BinarySearch=", mystd.BinarySearch(arr, 45))
	s := strconv.Itoa(65)
	fmt.Printf("Itoa=%T,%v\n", s, s)

	arr3 := []int{1, 0, 0, 0, 12}
	fmt.Println("MoveZero", mystd.MoveZero(arr3))

	arr4 := []int{2, 7, 11, 15}
	fmt.Println("TwoSum", mystd.TwoSum(arr4, 9))

	arr1 := mystd.CreateRandArr(1000)
	startTime := time.Now()
	mystd.SelectSort(arr1)
	stopTime := time.Now()
	fmt.Println("SelectSort time:", stopTime.Sub(startTime))

	arrInsert := mystd.CreateRandArr(1000)
	startTime = time.Now()
	mystd.InsertSort(arrInsert)
	stopTime = time.Now()
	fmt.Println("InsertSort time:", stopTime.Sub(startTime))

	arr5 := mystd.CreateRandArr(1000)
	startTime = time.Now()
	mystd.MergeSort(arr5, 0, len(arr5)-1)
	stopTime = time.Now()
	fmt.Println("MergeSort time:", stopTime.Sub(startTime))

	arr6 := mystd.CreateRandArr(1000)
	startTime = time.Now()
	mystd.QuickSort(arr6, 0, len(arr6)-1)
	stopTime = time.Now()
	fmt.Println("QuickSort time:", stopTime.Sub(startTime))

	arr7 := mystd.CreateRandArr(1000)
	startTime = time.Now()
	mystd.QuickSort2(arr7, 0, len(arr7)-1)
	stopTime = time.Now()
	fmt.Println("QuickSort2 time:", stopTime.Sub(startTime))

	arr8 := mystd.CreateRandArr(1000)
	startTime = time.Now()
	mystd.QuickSort3(arr8, 0, len(arr8)-1)
	stopTime = time.Now()
	fmt.Println("QuickSort3 time:", stopTime.Sub(startTime))

	var newPeople = inter.People{
		"rubinus",
		36,
	}
	newPeople.Call(19919995183)
	newPeople.Say()

	var other inter.PAS
	var abc inter.ABC
	other = inter.People{
		"5班",
		23,
	}
	abc = inter.Student{
		"5班",
	}
	other.Say()
	abc.WhichClass()

	var def = inter.DEF{
		Name:  "zhu",
		Age:   30,
		Class: "6班",
	}
	def.WhichClass()
	def.Sing()
	def.Say()

	fmt.Printf("肥波拿起：%d\n", mystd.Fib(6))*/

	//var j = 1
	//for ; j > 0; j-- {
	//
	//}
	//fmt.Println(j, "=====j===")

	//pkg.TestPkg()

	//ip.TestIP()

	//pdf.TestPDF() 有问题

	//json.TestJson()

	//crawler.GetCity()
	//mongo.TestMongo()

	//http.HandleFunc("/mysqldb", mysqldb)
	//http.HandleFunc("/mongodb", mongodb)
	//http.HandleFunc("/redis", One)
	//err := http.ListenAndServe(":8088", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	//mysql.TestMysql()
	//mongo.TestMongo()
	//redis.TestRedis()

	//fmt.Println("\n\n======\n\n")
	//kafka.TestKafka()

	//for {
	//	time.Sleep(10 * time.Second)
	//	fmt.Println(10, "让它在docker中跑....=====")
	//}

	//zero := []int{0, 3, 4, 1, 0, 2, 0, 1, 0}
	//fmt.Println(mystd.MoveZero(zero))

	//nsq.Consumer("test1", "test-channel", "localhost:4150", 2)
	//nsq.Producer("test1", "localhost:4150")
	fmt.Println(0x7fffffff)
}

func One(w http.ResponseWriter, r *http.Request) {
	str := redis.TestRedis()
	io.WriteString(w, "<h1>hello redis by docker </h1>\n<h2>"+str+" </h2>")
}
func mongodb(w http.ResponseWriter, r *http.Request) {
	str := mongo.TestMongo()
	hostname, _ := os.Hostname()
	io.WriteString(w, fmt.Sprintf("<h1>docker-compose scale: %s ,"+
		"<<<<<<<终于可以本地开发并且在docker中运行了!!!! </h1>\n<h2> %s </h2>", hostname, str))
}

func mysqldb(w http.ResponseWriter, r *http.Request) {
	str := mysql.TestMysql()
	io.WriteString(w, "<h1>hello mysqldb </h1>\n<h2>"+str+" </h2>")
}
