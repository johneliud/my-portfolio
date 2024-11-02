package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/johneliud/my-portfolio/models"
)

type BlogStore struct {
	postsDir string
}

func NewBlogStore(postsDirectory string) (*BlogStore, error) {
	postsDir := filepath.Join(postsDirectory, "posts")
	if err := os.MkdirAll(postsDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create posts directory: %w", err)
	}
	return &BlogStore{postsDir: postsDir}, nil
}

func (bs *BlogStore) getPostPath(_ time.Time, id string) string {
	return filepath.Join(bs.postsDir, id+".json")
}

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

func (bs *BlogStore) GetPost(id string) (*models.BlogPost, error) {
	if id == "" {
		return nil, errors.New("post ID cannot be empty")
	}

	// Direct file path
	path := filepath.Join(bs.postsDir, id+".json")
	return bs.loadPost(path)
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
		post.ReadingTime = CalculateReadingTime(post.Content)
	}
	return &post, nil
}
