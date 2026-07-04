package models

import "fmt"

type Feed struct {
	ID          int
	Title       string
	URL         string
	Description string
}

type Article struct {
	ID      int
	FeedID  int
	Title   string
	URL     string
	Content string
	IsRead  bool
}

func NewFeed(title, url, description string) Feed {
	return Feed{
		Title:       title,
		URL:         url,
		Description: description,
	}
}

func (f Feed) String() string {
	return fmt.Sprintf("[Feed] %s — %s", f.Title, f.URL)
}

func (a Article) String() string {
	return fmt.Sprintf("[Article] %s... (read: %t)", a.Title, a.IsRead)
}

func (a *Article) MarkAsRead() {
	a.IsRead = true
}
