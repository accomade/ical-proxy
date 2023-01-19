package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func rateLimit(c *gin.Context) {
	ip := c.ClientIP()
	value := int(ips.Add(ip, 1))
	if value%50 == 0 {
		fmt.Printf("ip: %s, count: %d\n", ip, value)
	}
	if value >= 200 {
		if value%200 == 0 {
			fmt.Println("ip blocked")
		}
		c.Abort()
		c.String(http.StatusServiceUnavailable, "you were automatically banned :)")
	}
}

func ical(c *gin.Context) {
	url := c.Param("url")

	req, err := http.NewRequest(http.MethodGet, url, nil) //c.Request.Response.Body)
	if err != nil {
		log.Printf("Error fetching cal: %v", err)
		c.String(400, "Something's off on request creation")
	}

	client := http.Client{
		Timeout: 1000,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error fetching cal: %v", err)
		c.String(400, "Something's off while fetching")
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error fetching cal: %v", err)
		c.String(400, "Something's off with resp Body")
	}
	c.String(200, string(b))
}
