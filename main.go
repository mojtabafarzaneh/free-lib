package main

import (
	"fmt"

	"github.com/ciehanski/libgen-cli/libgen"
	"github.com/gin-gonic/gin"
)

func main() {
	working := libgen.GetWorkingMirror(libgen.SearchMirrors)
	fmt.Println("the working mirror is: ", working)
	app := gin.Default()
	app.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ok"})
	})
	app.Run()
}
