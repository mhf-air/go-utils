package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func connect() {
	client := http.Client{}
	url := "http://wx.qlogo.cn/mmopen/ajNVdqHZLLCxbcG9yjficibCwqJ5bNhZlRbxB0zYbXkFtEjZxPZR5Na8ZZ0RLrmom6JbNUxhuj8tdsXyGPwIlDOw/0"
	resp, err := client.Get(url)
	defer resp.Body.Close()
	Check(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	Check(err)

	fout, err := os.OpenFile("head.jpg", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	Check(err)

	fout.Write(respBody)
	fmt.Println("done.")
}

func clientMain() {
	client := http.Client{}
	reqUrl := "http://192.168.1.122:9000/utils/importfile/import"
	postBody := url.Values{}
	postBody.Add("super", "yes")
	postBody.Add("macaddr", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	postBody.Add("type", "commodity")

	file, err := os.Open("commodity.xlsx")
	Check(err)
	stat, err := file.Stat()
	Check(err)
	b := make([]byte, stat.Size())
	_, err = file.Read(b)
	Check(err)
	postBody.Add("file", string(b))
	postData := []byte(postBody.Encode())

	req, err := http.NewRequest("POST", reqUrl, bytes.NewBuffer(postData))
	Check(err)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Api-Shopid", "f20ea308ea9189489ee6c0400e2d28e4")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	Check(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	Check(err)

	m := map[string]interface{}{}
	json.Unmarshal(respBody, &m)
	Ln(m)
	fmt.Println("done.")

}
