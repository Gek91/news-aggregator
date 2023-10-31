package main

import (
	"net/http"
	"newsAggregator/feeds"
	"strings"

	"github.com/gin-gonic/gin"
)

var feedsService feeds.FeedService

func main() {

	feedsService = feeds.NewFeedService(feeds.NewLocalDataFeedRepository())

	router := gin.Default()

	router.GET("/feeds", getFeed)

	router.Run("localhost:8080")
}

func getFeed(context *gin.Context) {

	filter := retrieveFilterFromString(context.Query("names"))

	context.IndentedJSON(http.StatusOK, feedsService.GetNews(filter))
}

func retrieveFilterFromString(queryString string) []string {

	var filter []string

	queryString = strings.Trim(queryString, "")

	if len(queryString) != 0 {
		filter = strings.Split(queryString, ",")
	}

	return filter
}
