package hertzstudy

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type Args struct {
	Query      string   `query:"query"`
	QuerySlice []string `query:"q"`
	Path       string   `path:"path"`
	Header     string   `header:"header"`
	Form       string   `form:"form"`
	Json       string   `json:"json"`
	Vd         int      `query:"vd" vd:"$==0 || $==1"`
}

//func main() {
//	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
//	h.POST("v:path/bind", func(c context.Context, ctx *app.RequestContext) {
//		var arg Args
//		err := ctx.BindAndValidate(&arg)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(arg)
//	})
//	h.Spin()
//}

func MyMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		//pre-handle,每次执行前打印
		fmt.Println("pre-handle")
		//相当于洋葱模型，向下执行handler或者中间件，如果不需要可以不写
		ctx.Next(c)
		//post-handler，每次执行完打印
		fmt.Println("post-handle")
	}
}

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	h.Use(MyMiddleware())
	h.GET("/middleware", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "Hello hertz!")
	})
	h.Spin()
}
