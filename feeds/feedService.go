package feeds

import (
	"fmt"
	"slices"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedService interface {
	GetNews(filter []string) []*FeedNews
}

type feedServiceImpl struct {
	feedRepository FeedRepository
}

func NewFeedService(feedRepository FeedRepository) FeedService {
	result := &feedServiceImpl{feedRepository}

	return result
}

func (service *feedServiceImpl) GetNews(filter []string) []*FeedNews {

	var newsSlice = make([]*FeedNews, 0)
	newsChannel := make(chan []*FeedNews)

	urlsToParse := service.getUrlsToParse(filter)

	for _, url := range urlsToParse {
		go parseFeed(url, newsChannel)
	}

	for _ = range urlsToParse {
		newsSlice = append(newsSlice, <-newsChannel...)
	}

	sort.Slice(
		newsSlice,
		func(i, j int) bool {
			return newsSlice[i].PublicationTime.UTC().Before(newsSlice[j].PublicationTime.UTC())
		})

	return newsSlice
}

func (service *feedServiceImpl) getUrlsToParse(filter []string) []*FeedInfo {

	if filter == nil {
		return service.feedRepository.GetFeedsInfo()
	}

	var result = make([]*FeedInfo, 0)

	for _, url := range service.feedRepository.GetFeedsInfo() {
		if slices.Contains(filter, url.name) {
			result = append(result, url)
		}
	}

	return result
}

func parseFeed(feedInfo *FeedInfo, channel chan []*FeedNews) {

	var newsSlice = make([]*FeedNews, 0)

	feedParser := gofeed.NewParser()
	if feed, err := feedParser.ParseURL(feedInfo.url); err == nil {

		for _, feedItem := range feed.Items {

			var publishedTime time.Time

			if publishedTime, err = time.Parse(feedInfo.publicationFormat, feedItem.Published); err != nil {
				fmt.Println(err.Error())
			}

			newsSlice = append(newsSlice, &FeedNews{
				feed.Title,
				feedItem.Title,
				feedItem.Link,
				publishedTime,
				feedItem.Categories,
			})
		}
	} else {
		panic(err.Error())
	}

	channel <- newsSlice
}
