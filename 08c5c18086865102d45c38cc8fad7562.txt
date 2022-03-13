package initRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasteTest/src/DBS"
	"pasteTest/src/handler"
)

func SetupRouter()*gin.Engine{
	router := gin.Default()
	DBS.Init()
	router.LoadHTMLGlob( "html/*" )
	router.StaticFS("/js", http.Dir("./js/"))
	router.StaticFS("/css", http.Dir("./css/"))
	router.GET("/",handler.ToBindHtml)
	router.POST("/getUrl",handler.GiveUrl)
	router.POST("/submit",handler.FormSubmit)
	router.GET("/:realUrl",handler.FormPerform)
	router.POST("/urlBind",handler.UrlBind)
	router.POST("/getPoster",handler.ContentPerform)
	return router
}
