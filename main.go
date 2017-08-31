package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()
	router.GET("/", helloworld)
	router.GET("/pass", pass)
	router.GET("/find", find)

	router.LoadHTMLGlob("html/*.html")
	router.Run(":8080")
}

func helloworld(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func pass(c *gin.Context) {
	var buf struct {
		Code string `form:"code"`
	}
	c.Bind(&buf)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"code":   buf.Code,
		"answer": solve(buf.Code),
	})
}

func find(c *gin.Context) {
	var buf struct {
		Code string `form:"code"`
	}
	c.Bind(&buf)

	c.String(200, solve(buf.Code))

}
func solve(code string) string {
	var ret string
	if strings.Count(code, "{") > 0 || strings.Count(code, "include") > 0 {
		ret = "C/C++"
	} else {
		ret = "Not defined yet"
	}
	return ret
}
