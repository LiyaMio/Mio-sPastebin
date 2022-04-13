package handler

import (
	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"net/http"
)
//func MD5(v string) string{
//	d := []byte(v)
//	m := md5.New()
//	m.Write(d)
//	return hex.EncodeToString(m.Sum(nil))
//}
//func Url() string {
//	conn, err := net.Dial("udp", "baidu.com:80")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//	defer conn.Close()
//	ip := strings.Split(conn.LocalAddr().String(), ":")[0]
//	nowtime :=time.Now().Unix()
//	formatTimeStr:=time.Unix(nowtime,0).Format("20060102150405")
//	//md5加密
//	codeUrl := ip + formatTimeStr
//	realUrl := MD5(codeUrl)
//	return realUrl
//}
func Url()string{
	id,_ :=gonanoid.Generate("abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",16)
	return id
}
func GiveUrl(c*gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"url": Url(),
	})
	//str := Url() + ".txt"
	//code := user.Content
	//content := []byte(code)
	//ioutil.WriteFile(str,content,0666)
}

