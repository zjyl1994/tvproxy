package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func m3uHandler(c *gin.Context) {
	m3u := `#EXTM3U
#EXTINF:-1,無綫新聞台
${baseURL}tvb/inews.m3u8
#EXTINF:-1,無綫財經資訊台
${baseURL}tvb/finance.m3u8
#EXTINF:-1,RTHK 31
${baseURL}rthk/31.m3u8
#EXTINF:-1,RTHK 32
${baseURL}rthk/32.m3u8
`
	processedBody := strings.ReplaceAll(m3u, "${baseURL}", baseURL)
	c.Data(200, "application/vnd.apple.mpegurl", []byte(processedBody))
}
