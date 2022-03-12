package model
type UserModel struct {
	Poster string `form:"poster"`
	Syntax string `form:"syntax"`
	Content string `form:"content"`
}
type UrlModel struct {
	Url string `form:"url"`
}