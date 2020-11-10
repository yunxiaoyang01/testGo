package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type User struct {
	Name string  `json:"name"`
	Age  int     `json:"age"`
}

func main()  {
	user1 :=&User{
		Name : "路斌",
		Age:  27,
	}
	user2 := &User{}
	err := deepCopy(user2,user1)
	if err !=nil{
		fmt.Println(err.Error())
	}
	user2.Name="孙颖"
	print(user1.Name)
	print(user2.Name)

}
func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
