package router

import (
	"github.com/gin-gonic/gin"
	"ginscaffold/controller"
	"ginscaffold/global"
	"ginscaffold/middleware"
	"ginscaffold/pkg/result"
	"net/http"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(middleware.AccessLog())
	router.Use(Recover)

	//static
	router.StaticFS("/static", http.Dir(global.StaticSetting.StaticDir))
	//article
	articlec:=controller.NewArticleController()
	router.GET("/article/getone/:id", articlec.GetOne);
	router.GET("/article/list", articlec.GetList);
	//user
	userc:=controller.NewUserController()
	router.POST("/user/login", userc.Login);
	router.GET("/user/info",middleware.JWTAuthMiddleware(), userc.Info);
	router.GET("/user/pass", userc.Pass);
	return router
}

//404
func HandleNotFound(c *gin.Context) {
	global.Logger.Errorf("handle not found: %v", c.Request.RequestURI)
	//global.Logger.Errorf("stack: %v",string(debug.Stack()))
	result.NewResult(c).ErrorCode(404,"资源未找到")
	return
}

//500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			global.Logger.Errorf("panic: %v", r)
			//log stack
			global.Logger.Errorf("stack: %v",string(debug.Stack()))
			//print stack
			debug.PrintStack()
			//return
			result.NewResult(c).ErrorCode(500,"服务器内部错误")
		}
	}()
	//继续后续接口调用
	c.Next()
}


/*
	f, _ := os.Create("/data/gologs/logs/test.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式

		return fmt.Sprintf("%s - [%s] \"%s %s\" %s %d %d %s \"%s\" %s\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.BodySize,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)

		//global.AccessLogger.Infof(logCont)

		//return logCont
	}))
*/

