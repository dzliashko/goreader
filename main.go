package main

import (
	"fmt"
	"goreader/models"
)

func main() {
	feed := models.NewFeed("Golang Weekly", "https://golangweekly.com/rss/", "Golang Weekly")
	articles := []models.Article{}
	articles = append(articles, models.Article{
		ID:      0,
		FeedID:  0,
		Title:   "Gin: 12 years, 88K stars, and zero broken APIs",
		URL:     "https://golangweekly.com/link/187511/rss",
		Content: "Building Gin: Simple Over Easy — Did you know Gin, the popular Go web framework, was built for a social network that never took off?",
		IsRead:  false,
	})
	articles = append(articles, models.Article{
		ID:      1,
		FeedID:  0,
		Title:   "Comparing six Go cache designs",
		URL:     "https://golangweekly.com/link/187208/rss",
		Content: "Awesome Go: ~3000 Categorized Go Resources — Most curated ‘awesome’ collections go stale, but I’ve been impressed that Go’s gets almost-daily updates! It’s a perennially useful resource and worth revisiting in a week short on big Go news (are we all enjoying the World Cup?",
		IsRead:  false,
	})

	for _, article := range articles {
		fmt.Printf("Feed Title: %s, Article Title: %s\n", feed.Title, article.Title)
	}
}
