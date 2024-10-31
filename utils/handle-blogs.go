package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

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
