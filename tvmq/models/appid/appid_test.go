package appid

import (
	"fmt"
	"testing"

	"github.com/json-iterator/go"
)

func TestFindById(t *testing.T) {
	u, err := FindById("5bee35269b9e33d2a84c68ff", nil)
	fmt.Println(u, err)
}

func TestGetAllAppids(t *testing.T) {
	appids, err := GetAllAppids()
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	b, err := jsonIterator.Marshal(appids) //encoding/json
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(len(appids), string(b))
}
