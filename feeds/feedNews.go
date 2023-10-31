package feeds

import (
	"fmt"
	"time"
)

type FeedNews struct {
	feedId          string
	Title           string
	Link            string
	PublicationTime time.Time
	Categories      []string
}

func (f *FeedNews) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%v\n", f.feedId, f.Title, f.Link, f.PublicationTime.String(), f.Categories)
}
