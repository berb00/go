package server

import (
	"encoding/json"
	"fmt"
	"golang/mysql"
	"io"
	"io/ioutil"
	"net/http"
	_ "strings"
)

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")                                                            // 允许访问所有域，可以换成具体url，注意仅具体url才能带cookie信息
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token") //header的类型
		w.Header().Add("Access-Control-Allow-Credentials", "true")                                                    //设置为true，允许ajax异步请求带cookie信息
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")                             //允许请求方法
		w.Header().Set("content-type", "application/json;charset=UTF-8")                                              //返回数据格式是json
		// if r.Method == "POST" {
		//     w.WriteHeader(http.StatusNoContent)
		//     return
		// }
		f(w, r)
	}
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Data struct {
	Name string
	Age  int
}
type Ret struct {
	Code  int
	Param string
	Msg   string
	Data  []Data
}

/*
type Request struct {
    Method string
    URL *url.URL
    Proto      string // "HTTP/1.0"
    ProtoMajor int    // 1
    ProtoMinor int    // 0
    Header Header
    Body io.ReadCloser
    GetBody func() (io.ReadCloser, error)
    ContentLength int64
    TransferEncoding []string
    Close bool
    Host string
    Form url.Values
    PostForm url.Values
    MultipartForm *multipart.Form
    Trailer Header
    RemoteAddr string
    RequestURI string
    TLS *tls.ConnectionState
    Cancel <-chan struct{}
    Response *Response
    ctx context.Context
}
*/

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json")

	fmt.Println("请求方法: 		", r.Method)
	fmt.Println("请求主机: 		", r.Host)
	fmt.Println("请求协议: 		", r.Proto)
	fmt.Println("请求远程地址: 	 ", r.RemoteAddr)
	fmt.Println("请求路由: 		", r.RequestURI)

	// var user User
	// decoder := json.NewDecoder(r.Body).Decode(&user)
	// fmt.Println("请求参数: ", decoder, r.Body)

	body, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	body_str := string(body)

	// var user User
	// json.Unmarshal([]byte(body_str), &user)
	// fmt.Printf("请求参数: %v", user.Password)

	var userInfo UserInfo
	json.Unmarshal([]byte(body_str), &userInfo)
	fmt.Printf("请求参数: %v-%v", userInfo.Age, userInfo.Name)

	mysql.InsertRowDemo(userInfo.Name, userInfo.Age) // 向数据库插入前端传如数据

	data := Data{Name: userInfo.Name, Age: userInfo.Age}
	ret := new(Ret)
	id := r.FormValue("id")
	ret.Code = 0
	ret.Param = id
	ret.Msg = "success"
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret.Data = append(ret.Data, data)
	ret_json, _ := json.Marshal(ret)
	io.WriteString(w, string(ret_json))

}
func HttpServer() {
	http.HandleFunc("/post", cors(index))

	http.ListenAndServe(":9090", nil)
}
