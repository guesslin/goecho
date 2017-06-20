package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func echo(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Data(http.StatusNoContent, "text/plain", []byte{})
		return
	}
	c.Data(http.StatusOK, http.DetectContentType(data), data)
}

type nullWriter struct{}

func (n *nullWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func main() {
	var ip string
	var port int
	var verbose bool
	flag.StringVar(&ip, "ip", "0.0.0.0", "ip to host")
	flag.IntVar(&port, "port", 8080, "port to listen")
	flag.BoolVar(&verbose, "verbose", false, "show requests")
	flag.Parse()

	var writer io.Writer

	server := gin.New()

	if verbose {
		writer = os.Stdout
	} else {
		writer = &nullWriter{}
	}

	server.Use(gin.LoggerWithWriter(writer), gin.Recovery())

	server.Any("/*path", echo)

	server.Run(fmt.Sprintf("%s:%d", ip, port))
}
