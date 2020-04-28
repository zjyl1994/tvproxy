package main

import (
	"github.com/gin-gonic/gin"
)

func tsProxyHandler(c *gin.Context) {
	remoteUrl := c.Query("url")
	resp, err := getHttpClient().Get(remoteUrl)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer resp.Body.Close()
	c.DataFromReader(200, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
