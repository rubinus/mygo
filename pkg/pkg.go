package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"time"
)

func TestPkg() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	//dat, err := ioutil.ReadFile("test.go")
	//check(err)
	//fmt.Print(string(dat))

	f, err := os.Open("test.go")
	check(err)

	// Read some bytes from the beginning of the file.
	// Allow up to 5 to be read but also note how many
	// actually were read.
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	url := "https://api.xuebaclass.com/xuebaapi/v1/provinces"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println("======string body=====", string(body))

	type Province struct {
		Id       int    `json:"id"`
		Province string `json:"province"`
	}
	provinces := make([]Province, 0)
	err = json.Unmarshal([]byte(body), &provinces)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(provinces)

	fmt.Println("cpus:", runtime.NumCPU(), runtime.Version(), runtime.GOMAXPROCS(10))
	fmt.Println("goroot:", runtime.GOROOT())
	fmt.Println("os/platform:", runtime.GOOS)

	fmt.Println(time.Now().Unix(), time.Now().Minute())

	fmt.Println(time.Unix(time.Now().Unix(), 0).Format("2010-01-01 00:00:00 PM"))

	var timestamp = time.Now().Unix()
	tm2 := time.Unix(timestamp, 0)
	fmt.Println(tm2.Format("2006-01-02 03:04:05 PM"))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
