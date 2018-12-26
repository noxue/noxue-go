package main

import (
	"fmt"
	"noxue/model"
)

func main() {
	a := &model.Tag{
		Name:"asdfasdf",
	}
	a.SetDoc(a)
	err:=a.Save()

	var b model.Tag

	err = a.FindByPk("5c22bd18311f62802e8aa332",&b)

	fmt.Println(b,err,b.Id.Hex())
}
