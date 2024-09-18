package api

import (
	"log"
	"net/http"
	"regexp"

	"github.com/ciehanski/libgen-cli/libgen"
	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/free-lib/models"
)

type DownloadController struct {
}

func NewDownloadController() *DownloadController {
	return &DownloadController{}
}

func (dc *DownloadController) Download(c *gin.Context) {
	downloadHash := models.DownloadQuery{}

	if err := c.BindJSON(&downloadHash); err != nil {
		log.Println("Error parsing search request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search request"})
		return
	}

	re := regexp.MustCompile(libgen.SearchMD5)
	if !re.MatchString(downloadHash.Hash) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid MD5 hash"})
	}

	hashArray := []string{downloadHash.Hash}

	searchMirror := libgen.GetWorkingMirror(libgen.SearchMirrors)
	bookDetails, err := libgen.GetDetails(&libgen.GetDetailsOptions{
		Hashes:       hashArray,
		SearchMirror: searchMirror,
		Print:        false,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	for _, book := range bookDetails {
		if err := libgen.GetDownloadURL(book, false); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "error getting download URL"})
			return
		}
		if err := libgen.DownloadBook(book, "/home/fuck_linus/free-lib"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}

	}
}
