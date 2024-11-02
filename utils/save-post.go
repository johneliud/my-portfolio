package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/johneliud/my-portfolio/models"
)

// Writes a blog post to a the appropriate directory
func (bs *BlogStore) SavePost(post *models.BlogPost) error {
	if post.Title == "" {
		return errors.New("post title is required")
	}

	if post.Content == "" {
		return errors.New("post content is required")
	}

	if post.Date.IsZero() {
		post.Date = time.Now()
	}

	// Generate ID if not provided
	if post.ID == "" {
		sanitized := strings.ToLower(strings.ReplaceAll(post.Title, " ", "-"))
		post.ID = fmt.Sprintf("%s-%d", sanitized, post.Date.Unix())
	}

	// Calculate reading time
	post.ReadingTime = CalculateReadingTime(post.Content)

	// Generate summary if not provided
	if post.Summary == "" {
		words := strings.Fields(post.Content)
		if len(words) > 30 {
			post.Summary = strings.Join(words[:30], " ") + "..."
		} else {
			post.Summary = strings.Join(words, " ")
		}
	}

	// Get the appropriate file path
	filePath := bs.getPostPath(post.Date, post.ID)
	if filePath == "" {
		return errors.New("failed to create post directory")
	}

	// Marshal to JSON with pretty printing
	data, err := json.MarshalIndent(post, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal post: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write post file: %w", err)
	}
	return nil
}
