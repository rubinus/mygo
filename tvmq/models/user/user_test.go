package user

import (
	"fmt"
	"testing"
)

func TestFindByAppidOpenid(t *testing.T) {
	u, err := FindByAppidOpenid("123", "456")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(u.Id.Hex(), u.MinAppid, u.MinOpenid)
}

func TestFindById(t *testing.T) {
	u, err := FindById("5bee35269b9e33d2a84c68ff", nil)
	fmt.Println(u, err)
}

func TestGetAllUsers(t *testing.T) {
	u, err := GetAllUsers()
	fmt.Println(u, err)
}
