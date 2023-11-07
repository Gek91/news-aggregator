package main

import (
	"fmt"
	"net/http"
	"newsAggregator/feeds"
	"newsAggregator/mail"
	"strings"

	"github.com/gin-gonic/gin"
)

var feedsService feeds.FeedService
var mailService mail.MailService

func main() {

	feedsService = feeds.NewFeedService(feeds.NewLocalDataFeedRepository())
	mailService = mail.NewMailService(&mail.GoogleGmailService{})

	router := gin.Default()

	router.GET("/feeds", getFeed)

	router.GET("/mail", getMail)

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

func getMail(context *gin.Context) {
	result := mailService.GetNews(context)

	fmt.Println("ctx.Err()", context.Err())

	context.IndentedJSON(http.StatusOK, result)
}
