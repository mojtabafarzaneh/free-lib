package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/free-lib/api"
)

func main() {
	app := gin.Default()
	//conf := config.Get()
	search := app.Group("/search")
	download := app.Group("/download")
	sc := api.NewSearchControler()
	dc := api.NewDownloadController()
	download.POST("", dc.Download)
	search.POST("", sc.Search)
	app.Run()
}
