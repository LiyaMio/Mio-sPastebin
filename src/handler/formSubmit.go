package handler

import (
	"fmt"
	"pasteTest/src/Encryption"
	"strings"
	//"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	//"os"
	"pasteTest/src/DBS"
	"pasteTest/src/model"
)
var user model.UserModel
var realUrl model.UrlModel
func ToBindHtml(c*gin.Context){
	c.HTML(http.StatusOK,"paste.html",nil)
	fmt.Println("paste页面加载成功")
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
	fmt.Println("相关信息获取成功")
}
func FormPerform(c*gin.Context){
	c.HTML(http.StatusOK,"code.html",nil)
	fmt.Println("code页面加载成功")
}
func ContentPerform(c*gin.Context){//根据url查找需要展示的内容
	url := c.PostForm("url")
	key := c.PostForm("priK")
	name := url + ".txt"
	fd,_ := ioutil.ReadFile(name)
	fmt.Println(string(fd))
	contentstr,_ := Encryption.AesCBCDncrypt(fd,[]byte(key))
	DBS.UpdateTime(url)
	fmt.Println("传输成功",contentstr)
	c.JSON(http.StatusOK,gin.H{
		"content": string(contentstr),
	})
	fmt.Println("url获取查询成功")
}
func UrlBind(c*gin.Context){
	err := c.ShouldBind(&realUrl)
	if err !=nil {
		fmt.Println("绑定失败")
	}else {
		fmt.Println("绑定成功")
		fmt.Println("realUrl",realUrl.Url)
	}
	str := realUrl.Url + ".txt"
	fmt.Println("写入名称",str)
	tmp := user.Content
	code := strings.Replace(tmp,"<","&lt;",-1)
	code = strings.Replace(code,">","&gt;",-1)
	key:=Encryption.AesKey(code)
	content,_:=Encryption.AesCBCEncrypt([]byte(code),key)
	fmt.Println("写入成功", content)
	ioutil.WriteFile(str,content,0666)
	url := realUrl.Url
	name :=user.Syntax
	DBS.StructInsert(url,name,str)
	fmt.Println("数据绑定成功")
	c.JSON(http.StatusOK,gin.H{
		"key":string(key),
	})
	//DBS.QueryMutiRowTime()

}