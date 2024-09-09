package api

import (
	"log"
	"net/http"
	"sync"

	"github.com/ciehanski/libgen-cli/libgen"
	"github.com/gin-gonic/gin"
	"github.com/mojtabafarzaneh/free-lib/models"
)

type SearchControler struct {
}

func NewSearchControler() *SearchControler {
	return &SearchControler{}
}

func (sc *SearchControler) Search(c *gin.Context) {
	searchArgs := new(models.SearchQuery)
	if err := c.BindJSON(searchArgs); err != nil {
		log.Println("Error parsing search request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid search request"})
		return
	}

	if searchArgs.Query == "" {
		log.Println("Empty search query")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query cannot be empty"})
		return
	}

	searchMirror := libgen.GetWorkingMirror(libgen.SearchMirrors)
	if searchMirror.Host == "" {
		log.Println("No valid search mirror found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No valid search mirror found"})
		return
	}

	var wg sync.WaitGroup
	var searchResults []*libgen.Book
	resultCh := make(chan []*libgen.Book, 1) // Buffered channel to avoid blocking
	errorCh := make(chan error, 1)           // Channel to handle errors

	// Perform the search in a separate goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		books, err := libgen.Search(&libgen.SearchOptions{
			Query:        searchArgs.Query,
			SearchMirror: searchMirror,
			Results:      10,
			Print:        false,
		})
		if err != nil {
			errorCh <- err
			return
		}
		resultCh <- books
	}()

	go func() {
		wg.Wait()
		close(resultCh)
		close(errorCh)
	}()

	select {
	case books := <-resultCh:
		searchResults = books
	case err := <-errorCh:
		log.Println("Error during search:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, searchResults)

}
