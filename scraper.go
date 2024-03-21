package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ghost-vk/rssagg/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	ctx := context.Background()
	ticker := time.NewTicker(timeBetweenRequest)
	// Если сделать через range: for range ticker.C {
	// то сначала будет ожидание, а нужно первое выполнение сразу
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(ctx, int32(concurrency))
		if err != nil {
			log.Println("worker: error get next feeds", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, f := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, f)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	log.Println("Start scrape feed", feed.Name)

	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("worker: error mark feed as fetched", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("worker: error fetch rss feed", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Found post", item.Title, "on feed", feed.Name)
	}

	log.Printf("Feed %s collected, %v post found\n", feed.Name, len(rssFeed.Channel.Item))
}
