package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
