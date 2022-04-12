package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Index(c *gin.Context) {
	fmt.Println("=========")
	fmt.Println("=========")
	fmt.Println("=========")
	fmt.Println(c.FullPath())
	fmt.Println(c.Request.URL.Path)
	fmt.Println(c.Request.Host)
	fmt.Println(c.Request.Context().Err())

	c.String(200, "ok")
}

func BuildRouter() {
	fmt.Println("=========")
	fmt.Println("=========")
	router := gin.Default()
	for i := 0; i < 5000; i++ {
		router.GET(fmt.Sprintf("/hello%d/:name/world/:user/:id", i), nil)
		router.GET(fmt.Sprintf("/world%d/:name/:user/:id", i), nil)
	}
}

func BuildRouterWithRandStr(routers []string) {
	router := gin.Default()
	for index, item := range routers {
		router.GET(fmt.Sprintf("/%s/:name/%s%d/:user/:id", item, item, index), Index)
	}
}

func BuildNormalRouter(routers []string) {
	router := gin.Default()
	for index, item := range routers {
		router.GET(fmt.Sprintf("/%s/name/%s%d/user/id", item, item, index), Index)
	}
}

type Aer interface {
	Hello() bool
	World() bool
}

type temp struct {
	a string
}

func (t *temp) Hello(flag string) string {
	return fmt.Sprintf("%s%s", t.a, flag)
}

type Foo struct {
	A string
}

func (f *Foo) Hello() {
	fmt.Println("=====", f.A)
}

type Fooer interface {
	Hello()
}

var addr = flag.String("listen-address", ":8081", "The address to listen on for HTTP requests.")

type Student struct {
	name string
}

func (s Student) sayHello(o string) {
	fmt.Println(o, s.name)
}

func (s Student) sayHello2(o string) {
	fmt.Println(o, s.name)
}

type Person interface {
	sayHello(o string)
}

type Person2 interface {
	sayHello(o string)
	sayHello2(o string)
}

func checkPath(path string) bool {
	for index := 0; index < len(path); index++ {
		if string(path[index]) == "[" {
			for j := index; j < len(path); j++ {
				if string(path[j]) == "]" {
					fmt.Println(path[index+1 : j])
					index = j
				}
			}
		}
	}
	return false
}

func f() {
	defer catch("f")

	g()
}

func g() {
	defer m()

	panic("g panic")
}

func m() {
	defer catch("m")
}

func catch(funcname string) {
	if r := recover(); r != nil {
		fmt.Println(funcname, "recover:", r)
	}
}

func ToLoganPath(path string) string {
	strs := strings.Split(path, "/")
	for index, str := range strs {
		if strings.HasPrefix(str, ":") {
			strs[index] = "{" + str[1:] + "}"
		}
	}
	return strings.Join(strs, "/")
}

func contentTest(c context.Context) {
	v := c.Value("vv").(*Ct)
	v.Hello = "hello"
	return
}

type Ct struct {
	Hello string
}

type InspectCtxItem struct {
	Suggest  string      `json:"suggest"`
	Metadata interface{} `json:"metadata"`
}
type CodeKey string
type InspectCtx struct {
	Prox        map[CodeKey]InspectCtxItem `json:"prox"`
	Plugin      map[CodeKey]InspectCtxItem `json:"plugin"`
	Upstream    map[CodeKey]InspectCtxItem `json:"upstream"`
	RPCRequest  string                     `json:"rpc_request"`
	RPCResponse string                     `json:"rpc_response"`
}

func main() {
	c := context.Background()
	v := &Ct{
		Hello: "world",
	}
	c = context.WithValue(c, "vv", v)
	contentTest(c)
	v1 := c.Value("vv").(*Ct)
	fmt.Println("========", *v1)

	strs := strings.Split("edith.hello.xiaohongshu.com", ".")
	index := 2
	if strs[len(strs)-3] == "sit" || strs[len(strs)-3] == "beta" {
		index = 3
	}

	host := strings.Join(strs[0:len(strs)-index], ".")
	fmt.Println("=======", host)

	values := url.Values{}
	values.Add("h", "hello")
	values.Add("w", "hello")
	fmt.Println(values.Encode())
	var s []*Student
	if err := json.Unmarshal([]byte("null"), &s); err != nil {
		fmt.Println("===xxxxxx", err)
	}
	fmt.Println("xxxx", s)
	fmt.Println(ToLoganPath("/hello/world"))
	fmt.Println(ToLoganPath("/hello/world/"))
	fmt.Println(ToLoganPath("/"))
	fmt.Println(ToLoganPath("/hello/:world/asd"))
	fmt.Println(ToLoganPath("hello"))

	emails := []string{"kaito@xiaohongshu.com"}
	var userList string
	for index, email := range emails {
		userList = userList + fmt.Sprintf("<%s>", email)
		if index != len(emails)-1 {
			userList = userList + " , "
		}
	}
	fmt.Println("xxxxxxxxxx", userList)
	f()
	clientIP := ""
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	fmt.Println("======", clientIP)

	h := hmac.New(sha256.New, []byte("5614f80df5bde1479e77e49d903b90d4"))

	sign := fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println(sign)
	h.Write([]byte(sign))
	fmt.Printf("%x\n", h.Sum(nil))

	return

	l, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(l).Format("2006-01-02 15:04:05"))
	checkPath("/api/[0-9a-z]")
	router := gin.Default()
	router.GET("/api", Index)
	router.POST("/api/sns/v1/media/:music_id/use_filter", Index)
	router.POST("/api/sns/v1/media/cancel_audio_task", Index)
	router.NoRoute(Index)
	router.Run()
}
