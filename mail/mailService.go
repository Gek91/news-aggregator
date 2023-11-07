package mail

import (
	"context"
	"time"
)

type MailService interface {
	GetNews(context context.Context) []*MailNews
}

type mailServiceImpl struct {
	googleGmailService *GoogleGmailService
}

func NewMailService(googleGmailService *GoogleGmailService) MailService {
	result := &mailServiceImpl{googleGmailService}
	return result
}

func (service *mailServiceImpl) GetNews(context context.Context) []*MailNews {

	var result = make([]*MailNews, 0, 0)

	listThreadsResponse := service.googleGmailService.GetThreadsList(context, "Label_1853585252419458839")

	for _, thread := range listThreadsResponse.Threads {

		mail := &MailNews{
			"",
			thread.Snippet,
			"",
			time.Now(),
		}

		result = append(result, mail)
	}

	return result
}
