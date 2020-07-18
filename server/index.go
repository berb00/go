package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	_ "strings"
	"encoding/json"
)


func HttpServer() {
	http.HandleFunc("/post", receiveClientRequest)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}

func receiveClientRequest(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")             //返回数据格式是json

    // r.ParseForm()
	// fmt.Println("收到客户端请求: ", r.Form)
	fmt.Fprintln(w, "Hello 沙河！")
}


func cors(f http.HandlerFunc) http.HandlerFunc {
	fmt.Printf("dsada")
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")  // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
        w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
        w.Header().Add("Access-Control-Allow-Credentials", "true") //设置为true，允许ajax异步请求带cookie信息
        w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") //允许请求方法
        w.Header().Set("content-type", "application/json;charset=UTF-8")             //返回数据格式是json
        if r.Method == "POST" {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        f(w, r)
    }
}
type User struct {
    Username string `json:"username"`
    Password  string  `json:"password"`
}

func index(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("AAA %v", r)
    fmt.Println(string(body))
    var user User
    if err := json.Unmarshal(body, &user); err == nil {
        fmt.Println(user)
    } else {
        fmt.Println(err)
    }
	w.Write([]byte("Hello Golang"))
	fmt.Fprintln(w, "Hello 沙河！")
	
}
func ZTestPost() {
    http.HandleFunc("/post", cors(index))
    http.ListenAndServe(":9090", nil)
}