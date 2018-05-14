package ginsrv

import (
	"net/http"

	"git.zhuzi.me/zzjz/zhuzi-payment/lib/config"

	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/gin-gonic/gin.v1"
)

type lRoute struct {
	*gin.RouterGroup
}

func (i *lRoute) GET(relativePath string, handler HandlerFunc) {
	i.RouterGroup.GET(relativePath, func(ctx *gin.Context) {
		rsp := handler(ctx)
		responseGin(ctx, rsp)
	})
}

func (i *lRoute) PUT(relativePath string, handler HandlerFunc) {
	i.RouterGroup.PUT(relativePath, func(ctx *gin.Context) {
		rsp := handler(ctx)
		responseGin(ctx, rsp)
	})
}

func (i *lRoute) POST(relativePath string, handler HandlerFunc) {
	i.RouterGroup.POST(relativePath, func(ctx *gin.Context) {
		rsp := handler(ctx)
		responseGin(ctx, rsp)
	})
}

func (i *lRoute) DELETE(relativePath string, handler HandlerFunc) {
	i.RouterGroup.DELETE(relativePath, func(ctx *gin.Context) {
		rsp := handler(ctx)
		responseGin(ctx, rsp)
	})
}

func responseGin(ctx *gin.Context, rsp interface{}) {
	switch r := rsp.(type) {
	case ErrorCoder:
		ctx.JSON(r.Code(), map[string]interface{}{
			"code": r.Code(),
			"msg":  r.Error(),
		})
	case error:
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": http.StatusBadRequest,
			"msg":  r.Error(),
		})
	default:
		ctx.JSON(http.StatusOK, r)
	}
	return
}

type ErrorCoder interface {
	Code() int
	Error() string
}

type lEngine struct {
	*gin.Engine
}

func (i *lEngine) Group(relativePath string, handlers ...gin.HandlerFunc) *lRoute {
	return &lRoute{i.Engine.Group(relativePath, handlers...)}
}

type HandlerFunc func(*gin.Context) interface{}

var httpServer *http.Server

func Listen(addr string) {
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		GracefulStop()
	}()

	httpServer.Addr = addr
	err := httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("http.ListenAndServe err:%+v", err)
	}
}

func GracefulStop() {
	err := httpServer.Shutdown(nil)
	if err != nil {
		log.Print(err)
	}
}

var Engine *lEngine

func init() {
	if !config.C.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	Engine = &lEngine{gin.New()}

	// 跨域
	Engine.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Headers", "Token,Content-Type")
		ctx.Header("Access-Control-Allow-Methods", "OPTIONS,PUT,POST,GET,DELETE")
	})

	// OPTIONS
	Engine.NoRoute(func(ctx *gin.Context) {
		if ctx.Request.Method == "OPTIONS" {
			ctx.JSON(200, nil)
		}
	})

	httpServer = &http.Server{
		Handler: Engine,
	}
}
