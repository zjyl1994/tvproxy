package main

import (
	"bufio"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// from https://www.rthk.hk/feeds/dtt/rthktv31_https.m3u8
func rthk31Handler(c *gin.Context) {
	m3u8ProxyHandler("https://rthklive1-lh.akamaihd.net/i/rthk31_1@167495/index_2052_av-p.m3u8?sd=10&rebase=on", c)
}

// from https://www.rthk.hk/feeds/dtt/rthktv32_https.m3u8
func rthk32Handler(c *gin.Context) {
	m3u8ProxyHandler("https://rthklive2-lh.akamaihd.net/i/rthk32_1@168450/index_1080_av-p.m3u8?sd=10&rebase=on", c)
}

func m3u8ProxyHandler(m3u8url string, c *gin.Context) {
	resp, err := getHTTPClient().Get(m3u8url)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	bodyString := string(bodyBytes)
	processedBody := m3u8Proc(bodyString, baseURL+"i.ts?url=")
	c.Data(200, resp.Header.Get("Content-Type"), []byte(processedBody))
}

func m3u8Proc(data string, prefixURL string) string {
	var sb strings.Builder
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		l := scanner.Text()
		if strings.HasPrefix(l, "#") {
			sb.WriteString(l)
		} else {
			sb.WriteString(prefixURL)
			sb.WriteString(url.QueryEscape(l))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
