package feeds

import (
	"slices"
	"testing"
)

var feedsService = NewFeedService(NewLocalDataFeedRepository()).(*feedServiceImpl)

func TestGetUrlsToParse(t *testing.T) {

	var tests = []struct {
		input     []string
		count     int
		feedNames []string
	}{
		{nil, 2, []string{"Google Developer", "Google Cloud Blog"}},
		{[]string{"ciao"}, 0, []string{}},
		{[]string{"Google Developer"}, 1, []string{"Google Developer"}},
		{[]string{"Google Cloud Blog"}, 1, []string{"Google Cloud Blog"}},
		{[]string{"Google Developer", "Google Cloud Blog"}, 2, []string{"Google Developer", "Google Cloud Blog"}},
	}

	for _, test := range tests {
		result := feedsService.getUrlsToParse(test.input)

		if len(result) != test.count {
			t.Errorf("getUrlsToParse(%q) result count = %v instead of %v", test.input, len(result), test.count)
		}

		nameSlice := make([]string, 0, len(result))
		for _, value := range result {
			nameSlice = append(nameSlice, value.name)
		}

		if !slices.Equal(nameSlice, test.feedNames) {
			t.Errorf("getUrlsToParse(%q) feed names not equal", test.input)
		}

	}
}
