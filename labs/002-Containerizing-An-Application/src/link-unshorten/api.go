package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

//API Returns Shortened Link Destination
func UnshortenLinkAPI(c *gin.Context) {
	origURL := c.Query("url")

	//add protocol
	if !isValidUrl(origURL) {
		origURL = "http://" + origURL
	}

	//Enforce dialer timeouts
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 8 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 8 * time.Second,
	}
	var netClient = &http.Client{
		Timeout:   time.Second * 8,
		Transport: netTransport,
	}
	resp, err := netClient.Head(origURL)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "not allowed",
		})
		log.Println("http.Get => %v", err.Error())
	}

	finalURL := resp.Request.URL.String()
	u, _ := url.Parse(finalURL)
	q, _ := url.ParseQuery(u.RawQuery)
	u.RawQuery = q.Encode()

	c.JSON(200, gin.H{
		"expanded": finalURL,
		"original": origURL,
	})
}
