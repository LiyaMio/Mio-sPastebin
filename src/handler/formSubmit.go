package handler

import (
	"fmt"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"pasteTest/src/DBS"
	"pasteTest/src/model"
	"strings"
)
var user model.UserModel
var realUrl model.UrlModel
func ToBindHtml(c*gin.Context){
	c.HTML(http.StatusOK,"paste.html",nil)
}
func FormSubmit(c*gin.Context){
	err := c.ShouldBind(&user)
	if err !=nil {
		fmt.Println("绑定失败")
	}else {
		fmt.Println("绑定成功")
		fmt.Println("poster",user.Poster)
		fmt.Println("syntax",user.Syntax)
		fmt.Println("content",user.Content)
	}
}
func FormPerform(c*gin.Context){
	c.HTML(http.StatusOK,"code.html",nil)


}
func ContentPerform(c*gin.Context){//根据url查找需要展示的内容
	url := c.PostForm("url")
	name := url + ".txt"
	content,err := os.Open(name)
	decoder := mahonia.NewDecoder("utf-8")
	fd,_ := ioutil.ReadAll(decoder.NewReader(content))
	contentstr := strings.Split(string(fd),"\n")
	if err !=nil{
		fmt.Println("read fail",err)
	}
	fmt.Println(contentstr)
	c.JSON(http.StatusOK,gin.H{
		"content": contentstr,
	})
}
func UrlBind(c*gin.Context){
	err := c.ShouldBind(&realUrl)
	if err !=nil {
		fmt.Println("绑定失败")
	}else {
		fmt.Println("绑定成功")
		fmt.Println("realUrl",realUrl.Url)
	}

	fmt.Println("write", user.Content)
	str := realUrl.Url + ".txt"
	code := user.Content
	content := []byte(code)
	fmt.Println("write", user.Content)
	ioutil.WriteFile(str,content,0666)
	url := realUrl.Url
	name :=user.Syntax
	DBS.StructInsert(url,name,str)

}
//func Search(c*gin.Context){
//	name := c.PostForm("name")
//	str := DBS.QueryRowByName(name)
//
//}