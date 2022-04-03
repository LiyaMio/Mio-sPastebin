package initRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pasteTest/src/DBS"
	"pasteTest/src/Middle"
	"pasteTest/src/handler"
)

func SetupRouter()*gin.Engine{
	router := gin.Default()
	DBS.Init()
	go handler.CleanData()
	router.LoadHTMLGlob( "html/*" )
	router.StaticFS("/js", http.Dir("./js/"))
	router.StaticFS("/css", http.Dir("./css/"))
	router.GET("/",Middle.Logger(),handler.ToBindHtml)
	router.POST("/getUrl",Middle.Logger(),handler.GiveUrl)
	router.POST("/submit",Middle.Logger(),handler.FormSubmit)
	router.GET("/:realUrl",Middle.Logger(),handler.FormPerform)
	router.POST("/urlBind",Middle.Logger(),handler.UrlBind)
	router.POST("/getPoster",Middle.Logger(),handler.ContentPerform)
	//router.Use(Middle.Logger(),gin.Recovery())
	return router
}
