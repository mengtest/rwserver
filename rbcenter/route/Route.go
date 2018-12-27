package route

import (
	"github.com/gorilla/mux"
	"net/http"
	Ctrl "../controller"
	"time"
	tb "../../rbwork/base"
	R "../../rbstruct/base"
	"encoding/json"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name,route.LoginValidate)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
//请求拦截器
func logger(inner http.Handler, name string,loginValidate bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		r.ParseForm()
		strToken:=r.Header.Get("token")
		claims,err:=tb.DecodeToken(strToken)
		if loginValidate == true {
			if strToken =="" {
				json.NewEncoder(w).Encode(R.ErrorMsg("未登录").OutLog())
				return
			}
			if err != nil && claims["uid"] !=nil {
				json.NewEncoder(w).Encode(R.ErrorMsg("请重新登录").OutLog())
				return
			}
		}
		if claims["uid"] != nil {
			r.Form.Set("lLoginId",claims["uid"].(string)) //给每个请求设置登录人ID
		}

		tb.LogInfo("request==>",r.RequestURI," params==>",r.Form)
		inner.ServeHTTP(w, r)
		tb.LogInfo("completed:"+name,"  time:",time.Since(start))
	})
}

//定义路由
type Route struct {
	Name        string  //接口命名
	LoginValidate bool  //接口是否需要验证登录
	Method      string  //接口请求方式
	Pattern     string  //接口请求URL
	HandlerFunc http.HandlerFunc  //接口业务逻辑
}

type Routes []Route

var routes = Routes{
	Route{Name: "CheckVersion", LoginValidate:false, Method: "GET", Pattern: "/checkVersion", HandlerFunc: Ctrl.CheckVersion},   //检测更新接口
	Route{Name: "Login",LoginValidate:false, Method: "POST", Pattern: "/login", HandlerFunc: Ctrl.Login}, //登录
}


