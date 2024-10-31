package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
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
	if err := os.MkdirAll(postsDirectory, 0755); err != nil {
		return nil, fmt.Errorf("failed to create posts directory: %w", err)
	}
	return &BlogStore{postsDir: postsDirectory}, nil
}

// Reads all blog posts for preview
func (bs *BlogStore) LoadPosts() ([]models.BlogPost, error) {
	var posts []models.BlogPost

	err := filepath.WalkDir(bs.postsDir, func(path string, d fs.DirEntry, err error) error {
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

// Load a single blog post by ID
func (bs *BlogStore) GetPost(id string) (*models.BlogPost, error) {
	if id == "" {
		return nil, errors.New("post ID cannot be empty")
	}
	filename := filepath.Join(bs.postsDir, id+".json")
	return bs.loadPost(filename)
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

	// Validate required fields
	if post.Title == "" {
		return nil, errors.New("post title is required")
	}

	if post.Content == "" {
		return nil, errors.New("post content is required")
	}

	if post.Date.IsZero() {
		return nil, errors.New("post date is required")
	}

	// Ensure reading time is calculated
	if post.ReadingTime == 0 {
		post.ReadingTime = CalculateReadingTime(post.Content)
	}
	return &post, nil
}

// Writes a blog post to a JSON file
func (bs *BlogStore) SavePost(post *models.BlogPost) error {
	// Validate post
	if post.Title == "" {
		return errors.New("post title is required")
	}

	if post.Content == "" {
		return errors.New("post content is required")
	}

	if post.Date.IsZero() {
		post.Date = time.Now()
	}

	if post.ID == "" {
		// Generate ID from title and date
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

	// Marshal to JSON with pretty printing
	data, err := json.MarshalIndent(post, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal post: %w", err)
	}

	// Write to file
	filename := filepath.Join(bs.postsDir, post.ID+".json")
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write post file: %w", err)
	}
	return nil
}
