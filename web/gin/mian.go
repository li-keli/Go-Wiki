package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// docker 部署
// docker run -p 8081:80 --rm -d -v "$(pwd)/conf.d":/etc/nginx/conf.d -v "$(pwd)":/var/work -w /var/work nginx sh -c "nginx & ./gin"
func main() {
	fmt.Println("hello")
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "ok"})
	})
	r.Run()
}
