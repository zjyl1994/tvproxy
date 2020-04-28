package main

import (
	"errors"
	"io/ioutil"
	"regexp"

	"github.com/gin-gonic/gin"
)

func parseTVB(liveName string) (string, error) {
	resp, err := getHttpClient().Get("http://news.tvb.com/live/" + liveName)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	m3u8Regexp := regexp.MustCompile(`<source src="(.*?)" type="application/x-mpegURL">`)
	match := m3u8Regexp.FindStringSubmatch(string(bBody))
	if len(match) > 1 {
		return match[1], nil
	} else {
		return "", nil
	}
}

func tvbHandler(liveName string, c *gin.Context) {
	m3u8, err := parseTVB(liveName)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	if m3u8 == "" {
		c.AbortWithError(404, errors.New("video not found"))
	} else {
		c.Redirect(302, m3u8)
	}
}

func iNewsHandler(c *gin.Context) {
	tvbHandler("inews", c)
}

func financeHandler(c *gin.Context) {
	tvbHandler("j5_ch85", c)
}
