package main

import (
	"fmt"
	"sync"
	"time"
)

type Article struct {
	Title   string
	Content string
}

type ArticleStore struct {
	articles map[string]Article
	rwMutex  sync.RWMutex
}

func (as *ArticleStore) ReadArticle(title string) (Article, bool) {
	as.rwMutex.RLock()
	defer as.rwMutex.RUnlock()

	article, exists := as.articles[title]
	return article, exists
}

func (as *ArticleStore) WriteArticle(title, content string) {
	as.rwMutex.Lock()
	defer as.rwMutex.Unlock()

	// Simulate processing time for writing/updating an article
	time.Sleep(100 * time.Millisecond)

	as.articles[title] = Article{title, content}
	fmt.Printf("Article '%s' updated\n", title)
}

func runArticleStore() {
	articleStore := ArticleStore{
		articles: make(map[string]Article),
	}

	request := 5

	// concurrent reades
	for i := 0; i < request; i++ {
		go func(userID int) {
			title := "Introduction to Go"
			article, exists := articleStore.ReadArticle(title)
			if exists {
				fmt.Printf("User %d is reading '%s': %s\n", userID, title, article.Content)
				return
			}

			fmt.Printf("User %d couldn't find article '%s'\n", userID, title)
		}(i + 1)
	}

	// exlusive writer
	time.Sleep(500 * time.Millisecond)

	// simulationg more readers after the update
	for i := 0; i < request; i++ {
		go func(userID int) {
			title := "Introduction to Go"
			article, exists := articleStore.ReadArticle(title)
			if exists {
				fmt.Printf("User %d is reading '%s': %s\n", userID, title, article.Content)
				return
			}

			fmt.Printf("User %d couldn't find article '%s'\n", userID, title)
		}(i + 1)
	}

	time.Sleep(500 * time.Millisecond)

	article, exists := articleStore.ReadArticle("Introduction to Go")
	if exists {
		fmt.Printf("Final state of the article: %s\n", article.Content)
		return
	}

	fmt.Println("Article 'Introduction to Go' not found.")
}
