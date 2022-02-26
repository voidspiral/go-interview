package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
//cookie中的字段，不是实际使用

type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)
}
func GetCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Println(cookie)
	fmt.Fprint(w, cookie)
}
func main() {

	http.HandleFunc("/set", SetCookie) // 设置访问的路由
	http.HandleFunc("/get", GetCookie)
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口


	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}