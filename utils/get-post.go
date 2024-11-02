package utils

import (
	"errors"
	"path/filepath"
	"time"

	"github.com/johneliud/my-portfolio/models"
)

func (bs *BlogStore) getPostPath(_ time.Time, id string) string {
	return filepath.Join(bs.postsDir, id+".json")
}

func (bs *BlogStore) GetPost(id string) (*models.BlogPost, error) {
	if id == "" {
		return nil, errors.New("post ID cannot be empty")
	}

	// Direct file path
	path := filepath.Join(bs.postsDir, id+".json")
	return bs.loadPost(path)
}
