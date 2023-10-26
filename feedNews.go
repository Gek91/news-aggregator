package main

import (
	"fmt"
	"time"
)

type FeedNews struct {
	feedId          string
	Title           string
	Link            string
	publicationTime time.Time
	Categories      []string
}

func (f *FeedNews) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%v\n", f.feedId, f.Title, f.Link, f.publicationTime.String(), f.Categories)
}
