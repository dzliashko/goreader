package models

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
		ID:          0,
		Title:       title,
		URL:         url,
		Description: description,
	}
}
