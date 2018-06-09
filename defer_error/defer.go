package main

import (
	"mygo/defer_error/handle"
	"net/http"
	"os"

	"fmt"

	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func wrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	//函数做为参数，返回值也是函数
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Error("panic oo : ", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			fmt.Println("有错", err.Error())
			//判断是不是自己定义的err
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/", wrapper(handle.DeferHandle))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
