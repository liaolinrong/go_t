package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//read data from file
	bCert, err := ioutil.ReadFile("E:\\acm\\k8s_cert5_gai2_wincerts_cmdparam\\k8s\\k8s-server-cert\\cert.pem")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	//fmt.Print(string(bCert))

	bKey, err := ioutil.ReadFile("E:\\acm\\k8s_cert5_gai2_wincerts_cmdparam\\k8s\\k8s-server-cert\\key.pem")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	//fmt.Print(string(bKey))

	//http post
	resp, err := http.PostForm("http://localhost:8080/api/v1/alterCfg/addCertAndKey",
		url.Values{"cert": {string(bCert)}, "key": {string(bKey)}})

	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println(string(body))
}
