package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	tvb := r.Group("/tvb")
	{
		tvb.GET("/inews.m3u8", iNewsHandler)
		tvb.GET("/finance.m3u8", financeHandler)
	}
	r.Run()
}
