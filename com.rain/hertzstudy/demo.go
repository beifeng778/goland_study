package hertzstudy

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})
	h.Spin()
}

func RegisterRoute(h *server.Hertz) {
	h.GET("/get", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "get")
	})
	h.POST("/post", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "post")
	})
	h.PUT("/put", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "put")
	})
	h.DELETE("/delete", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "delete")
	})
	h.PATCH("/patch", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "patch")
	})
	h.HEAD("/head", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "head")
	})
	h.OPTIONS("/options", func(c context.Context, ctx *app.RequestContext) {
		ctx.String(consts.StatusOK, "options")
	})
	//simple group:v1
	v1 := h.Group("/v1")
	{
		//loginEndpoint is a handler func
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/streaming_read", readEndpoint)
	}
	//simple group:v2
	v2 := h.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/streaming_read", readEndpoint)
	}

	// This handler will match: "/hertz/version", but will not match : "/hertz/" or "/hertz"
	h.GET("/hertz/:version", func(ctx context.Context, c *app.RequestContext) {
		version := c.Param("version")
		c.String(consts.StatusOK, "Hello %s", version)
	})
	h.Spin()

	h.GET("/hertz/:version/*action", func(ctx context.Context, c *app.RequestContext) {
		version := c.Param("version")
		action := c.Param("action")
		message := version + " is " + action
		c.String(consts.StatusOK, message)
	})
	h.Spin()
}

func readEndpoint(ctx context.Context, ctx2 *app.RequestContext) {

}

func submitEndpoint(ctx context.Context, ctx2 *app.RequestContext) {

}

func loginEndpoint(ctx context.Context, ctx2 *app.RequestContext) {

}
