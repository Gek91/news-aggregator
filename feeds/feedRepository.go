package feeds

import "time"

type FeedRepository interface {
	GetFeedsInfo() []*FeedInfo
}

// don't export type -> no literal initialization
type localDataFeedRepositoryImpl struct {
	urls []*FeedInfo
}

func NewLocalDataFeedRepository() FeedRepository {
	result := &localDataFeedRepositoryImpl{}
	result.urls = []*FeedInfo{
		{"Google Developer", "https://developers.googleblog.com/feeds/posts/default?alt=rss", time.RFC3339},
		{"Google Cloud Blog", "https://cloudblog.withgoogle.com/rss/", time.RFC1123Z},
	}

	return result
}

func (repository *localDataFeedRepositoryImpl) GetFeedsInfo() []*FeedInfo {
	return repository.urls
}
