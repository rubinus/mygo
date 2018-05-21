package handle

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type userCustomErr string

func (uce userCustomErr) Error() string {
	return uce.Message()
}
func (uce userCustomErr) Message() string {
	return string(uce)
}

const prix = "/list/"

func DeferHandle(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prix) != 0 {
		return userCustomErr("必须得有一个前缀:" + prix)
	}
	path := request.URL.Path[len(prix):]
	file, err := os.Open(path) //打开文件
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
