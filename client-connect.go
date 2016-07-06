package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var ln = fmt.Println

type Person struct {
	Name string
	Age  int
}

func main() {
	client := http.Client{}
	url := "http://wx.qlogo.cn/mmopen/ajNVdqHZLLCxbcG9yjficibCwqJ5bNhZlRbxB0zYbXkFtEjZxPZR5Na8ZZ0RLrmom6JbNUxhuj8tdsXyGPwIlDOw/0"
	resp, err := client.Get(url)
	defer resp.Body.Close()
	check(err)

	fout, err := os.OpenFile("head.jpg", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	check(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	check(err)

	fout.Write(respBody)
	fmt.Println("done.")
}

//====================================================================================================
func check(err error) {
	if err != nil {
		panic(err)
	}
}
