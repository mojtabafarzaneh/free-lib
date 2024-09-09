package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/free-lib/api"
)

func main() {
	app := gin.Default()
	//conf := config.Get()
	search := app.Group("/search")
	sc := api.NewSearchControler()
	search.POST("", sc.Search)
	app.Run()
}
