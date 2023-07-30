package mock

import (
	"time"

	"github.com/jahidhimon/snppetbox/pkg/models"
)

var mockSnippet = &models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond...",
	UserID:  2,
	Created: time.Now(),
	Expires: time.Now().Add(7 * 24 * time.Hour),
}

type SnippetModel struct{}

func (m *SnippetModel) Insert(title, content, expires string, id int) (int, error) {
	return 2, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippet, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}

func (m *SnippetModel) UserSnippets(userID int) ([]*models.Snippet, error) {
	return []*models.Snippet{mockSnippet}, nil
}
