package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	//设置环境变量VERSION为当前Golang版本
	os.Setenv("VERSION", "Golang version: "+runtime.Version())

	//添加http路由handler
	http.HandleFunc("/healthz", homeWorkModule2)

	//启动http本地80端口服务
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homeWorkModule2(w http.ResponseWriter, r *http.Request) {
	//1. 配置响应体Header和请求的Header一致
	for k, v := range r.Header {
		w.Header().Set(k, v[0])
	}

	//2. 获取本地环境变量VERSION，并在响应中打印
	_, err := fmt.Fprintf(w, os.Getenv("VERSION"))
	if err != nil {
		log.Fatal(err)
	}

	statusCode := http.StatusOK

	//3.打印访问日志，包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	defer func() {
		reqIP, reqClient := GetRequestInfo(r)
		log.Printf("request ip: %s\nrequest client: %s\nresponse http code: %d", reqIP, reqClient, statusCode)
	}()

	//4.设置响应状态码
	w.WriteHeader(statusCode)
}

func GetRequestInfo(r *http.Request) (requestIP, requestClient string) {
	requestIP = r.Header.Get("X-FORWARDED-FOR")
	if requestIP == "" {
		requestIP = r.RemoteAddr
	}
	requestClient = string(r.Header.Get("User-Agent"))
	return
}
