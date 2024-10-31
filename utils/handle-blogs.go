package utils

import (
	"fmt"
	"os"
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
