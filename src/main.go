package main

import "pasteTest/src/initRouter"

func main(){
	router := initRouter.SetupRouter()
	_=router.Run()
}
