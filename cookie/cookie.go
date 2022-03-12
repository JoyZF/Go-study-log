/**
	golang 操作cookie 案例 ， https://github.com/gorilla/securecookie 使用securecookie 可以对cookie 进行aes加密 增加cookie的安全性
*/

package main

import (
	"fmt"
	"github.com/gorilla/securecookie"
	"net/http"
)

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/readcookie", ReadCookieServer)
	http.HandleFunc("/writecookie", WriteCookieServer)
	http.HandleFunc("/deletecookie", DeleteCookieServer)

	http.ListenAndServe(":8887", nil)
}


func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

func ReadCookieServer(w http.ResponseWriter, req *http.Request) {
	// read cookie
	//cookie, err := req.Cookie("testcookiename")
	//if err == nil {
	//	cookievalue := cookie.Value
	//	w.Write([]byte("<b>cookie的值是：" + cookievalue + "</b>\n"))
	//} else {
	//	w.Write([]byte("<b>读取出现错误：" + err.Error() + "</b>\n"))
	//}

	hashKey := []byte("zhouzhouzhouzhouzhouzhouzhouzhou")

	var blockKey = []byte("zhouzhouzhouzhouzhouzhouzhouzhou")
	var s2 = securecookie.New(hashKey, blockKey)

	cookie, err := req.Cookie("cookie-name")
	if  err == nil {
		value := make(map[string]string)
		if err = s2.Decode("cookie-name", cookie.Value, &value); err == nil {
			fmt.Fprintf(w, "The value of foo is %q", value["foo"])
		}
	}


}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
	//cookie := http.Cookie{Name: "testcookiename", Value: "testcookievalue", Path: "/", MaxAge: 86400}
	//http.SetCookie(w, &cookie)
	hashKey := []byte("zhouzhouzhouzhouzhouzhouzhouzhou")

	var blockKey = []byte("zhouzhouzhouzhouzhouzhouzhouzhou")
	var s = securecookie.New(hashKey, blockKey)
	value := map[string]string{
		"foo": "bar",
	}
	encode, err := s.Encode("cookie-name", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "cookie-name",
			Value: encode,
			Path:  "/",
			Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(err)
	w.Write([]byte("<b>设置cookie成功。</b>\n"))
}

func DeleteCookieServer(w http.ResponseWriter, req *http.Request) {
	cookie := http.Cookie{Name: "testcookiename", Path: "/", MaxAge: -1}
	http.SetCookie(w, &cookie)

	w.Write([]byte("<b>删除cookie成功。</b>\n"))
}