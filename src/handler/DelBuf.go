package handler

import (
	"fmt"
	"pasteTest/src/DBS"
	"time"
)

func CleanData(){
	ticker := time.NewTicker(time.Minute*5)
	fmt.Println("current ",time.Now())
	for {
			tm:=<-ticker.C
			fmt.Println("nowtime ",tm)
			DBS.QueryMutiRowTime()
		}

}