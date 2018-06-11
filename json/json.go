package json

import (
	"fmt"
	"strings"

	"os"

	"github.com/bitly/go-simplejson"
	"github.com/json-iterator/go"
)

func TestJson() {
	//简单应用NewDecoder
	fmt.Println("=========testjson============\n\n")
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	str := jsoniter.Get(val, "Colors").ToString()
	fmt.Println(str)

	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reader := strings.NewReader(`{"branch":"beta","change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`)
	decoder := json.NewDecoder(reader)
	params := make(map[string]interface{})
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", params)
	}

	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err = json.Unmarshal(jsonBlob, &animals) //encoding/json
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	var animals2 []Animal
	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	fmt.Println("jsonBlob []byte===", string(jsonBlob))
	jsonIterator.Unmarshal(jsonBlob, &animals2) //json_iterator
	fmt.Printf("%+v\n", animals2)
	fmt.Println("\n")
	//------
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group) //encoding_json
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("====%s====\n", b)

	fmt.Println("------------simplejson----------")

	js, err := simplejson.NewJson([]byte(`{
		"test": {
			"string_array": ["asdf", "ghjk", "zxcv"],
			"array": [1, "2", 3],
			"arraywithsubs": [{"subkeyone": 1},
			{"subkeytwo": 2, "subkeythree": 3}],
			"int": 10,
			"float": 5.150,
			"bignum": 9223372036854775807,
			"string": "simplejson",
			"bool": true
		}
	}`))
	if err != nil {
		panic("json format error")
	}
	s, err := js.Get("test").Get("arraywithsubs").Array()
	if err != nil {
		fmt.Println("decode error: get int failed!")
		return
	}
	fmt.Println(json.MarshalToString(s))
	fmt.Println("------------simplejson----------")

	os.Stdout.Write(b)
	fmt.Println("\n")
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	b, err = json_iterator.Marshal(group) //json_iterator
	os.Stdout.Write(b)

	fmt.Println("\n========")
	fns := functions()
	for f := range fns {
		fmt.Println(f, fns[f])
		fns[f]()
	}
	fmt.Println("=========testjson====================")

	//个人理解：
	//闭包就是能够读取其他函数内部变量的函数。
	//只有函数内部的子函数才能读取局部变量，因此可以把闭包简单理解成”定义在一个函数内部的函数”。
	// ctr, incr and ctr1, incr1 are different
	ctr, incr := counter(100)
	ctr1, incr1 := counter(100)
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1 - ", ctr1())
	// incr by 1
	incr()
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1- ", ctr1())
	// incr1 by 2
	incr1()
	incr1()
	fmt.Println("counter - ", ctr())
	fmt.Println("counter1- ", ctr1())
}

func functions() []func() {
	// pitfall of using loop variables
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	for i, v := range arr {
		fmt.Println(v, "----v-----", i)

		result = append(result, func() { fmt.Printf("index - %d, value - %d\n", i, arr[i]) })
	}

	return result
}

func counter(start int) (func() int, func()) {
	// if the value gets mutated, the same is reflected in closure
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	// both ctr and incr have same reference to start
	// closures are created, but are not called
	return ctr, incr
}
