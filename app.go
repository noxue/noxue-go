package main

import (
	"fmt"
	"noxue/dao"
)

func main() {

	fmt.Println(dao.UserDao.Select(nil,[]string{"-name"}, 2,3))
}
