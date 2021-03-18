package main

import (
	"context"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	c "gin-example/controller"
	_ "gin-example/docs"
	"github.com/gin-gonic/gin"
)

// 文档标题
// @title Swagger Example API
// 版本
// @version 1.0

//下面两行都是一些声明，可选
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// 下面三行联系人信息，可选
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// 必填
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// 如果你想直接swagger调试API，下面两项需要填写正确。前者为服务文档的端口，ip。后者为基础路径
// @host 127.0.0.1:3000
// @BasePath /
func main() {
	userHandler := c.NewUserIns()

	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DEV"))
	//url := ginSwagger.URL("http://localhost:3000/swagger/doc.json") // The url pointing to API definition
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	v1 := r.Group("/v1")
	v1.Use(gin.Logger())
	userGroup := v1.Group("/user")
	userGroup.GET("/:name", userHandler.Get)
	userGroup.POST("/name", userHandler.Post)
	userGroup.POST("/upload", userHandler.UpLoad)
	v2 := r.Group("v2")
	userGroupV2 := v2.Group("user")
	userGroupV2.GET("/:name", userHandler.GetV2)
	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
