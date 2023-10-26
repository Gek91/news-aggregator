package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/mmcdole/gofeed"
)

type FeedInfo struct {
	url               string
	publicationFormat string
}

func main() {

	var urls = []*FeedInfo{
		{"https://developers.googleblog.com/feeds/posts/default?alt=rss", time.RFC3339},
		{"https://cloudblog.withgoogle.com/rss/", time.RFC1123Z},
	}

	var newsSlice = make([]*FeedNews, 0)

	newsChannel := make(chan []*FeedNews)

	for _, url := range urls {
		go parseFeed(url, newsChannel)
	}

	for _ = range urls {
		newsSlice = append(newsSlice, <-newsChannel...)
	}

	sort.Slice(
		newsSlice,
		func(i, j int) bool {
			return newsSlice[i].publicationTime.UTC().Before(newsSlice[j].publicationTime.UTC())
		})

	for _, elem := range newsSlice {
		fmt.Println(elem)
	}

}

func parseFeed(feedInfo *FeedInfo, channel chan []*FeedNews) {

	var newsSlice = make([]*FeedNews, 0)

	fp := gofeed.NewParser()
	if feed, err := fp.ParseURL(feedInfo.url); err == nil {

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
		fmt.Println(err.Error())
	}

	channel <- newsSlice
}
