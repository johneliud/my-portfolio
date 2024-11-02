package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/johneliud/my-portfolio/models"
)

// Reads all blog posts for preview
func (bs *BlogStore) LoadPosts() ([]models.BlogPost, error) {
	var posts []models.BlogPost

	err := filepath.Walk(bs.postsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) != ".json" {
			return nil
		}

		post, err := bs.loadPost(path)
		if err != nil {
			return fmt.Errorf("error loading post %s: %w", path, err)
		}

		// Create blog preview
		previewPost := models.BlogPost{
			ID:          post.ID,
			Title:       post.Title,
			Summary:     post.Summary,
			Date:        post.Date,
			ReadingTime: post.ReadingTime,
			Tags:        post.Tags,
			ExternalURL: post.ExternalURL,
		}
		posts = append(posts, previewPost)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking posts directory: %w", err)
	}

	// Sort posts by date, newest first
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	return posts, nil
}

// Reads and parses a JSON file
func (bs *BlogStore) loadPost(filename string) (*models.BlogPost, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read post file: %w", err)
	}

	var post models.BlogPost

	if err := json.Unmarshal(file, &post); err != nil {
		return nil, fmt.Errorf("failed to unmarshal post: %w", err)
	}

	// Ensure reading time is calculated
	if post.ReadingTime == 0 {
		post.ReadingTime = CalculateReadingTime()
	}
	return &post, nil
}
