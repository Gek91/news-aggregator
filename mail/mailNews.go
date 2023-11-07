package mail

import (
	"fmt"
	"time"
)

type MailNews struct {
	senderId        string
	Snippet         string
	Link            string
	PublicationTime time.Time
}

func (f *MailNews) String() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n", f.senderId, f.Snippet, f.Link, f.PublicationTime.String())

}
