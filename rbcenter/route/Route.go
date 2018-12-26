package route

import (
	"github.com/gorilla/mux"
	"net/http"
	Ctrl "../controller"
	"time"
	tb "../../rbwork/base"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
//输出请求执行时间
func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		r.ParseForm()
		tb.LogInfo("request==>",r.RequestURI," params==>",r.Form)
		inner.ServeHTTP(w, r)
		tb.LogInfo("completed:"+name,"  time:",time.Since(start))
	})
}

//定义路由
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Name: "Index", Method: "GET", Pattern: "/", HandlerFunc: Ctrl.Index},                //主页
	Route{Name: "Download", Method: "GET", Pattern: "/download", HandlerFunc: Ctrl.Index},     //下载页
	Route{Name: "Patch", Method: "GET", Pattern: "/patch", HandlerFunc: Ctrl.Index},           //补丁页
	Route{Name: "CheckVersion", Method: "GET", Pattern: "/checkVersion", HandlerFunc: Ctrl.CheckVersion},   //检测更新接口
	Route{Name: "Login", Method: "GET", Pattern: "/login", HandlerFunc: Ctrl.Login}, //登录
}


