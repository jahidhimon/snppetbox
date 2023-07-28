package mysql

import (
	"database/sql"
	
	"github.com/jahidhimon/snppetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string, userID int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires, userID)
VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY), ?);`

	result, err := m.DB.Exec(stmt, title, content, expires, userID)
	if err != nil {
		return 0, err
	}
	
	// Use LastInsertId() method on the result object to get the ID of our newly
	// inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	
	// ID returned has the type int64, so we convert it to an int type
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `
SELECT id, title, content, created, expires, userID FROM snippets
WHERE expires > UTC_TIMESTAMP() AND id = ?
`
	s := &models.Snippet{}

	err := m.DB.QueryRow(stmt, id).
		Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.UserID)
	
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}
	
	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	// ADD userID to home screen
	stmt := `
SELECT id, title, content, created, expires, userID FROM snippets
WHERE expires > UTC_TIMESTAMP() ORDER BY created LIMIT 10
`
	snippets := []*models.Snippet{}
	
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires, &s.UserID)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// rows.Err() returns any error happening during iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return snippets, nil
}

func (m *SnippetModel) UserSnippets(userID int) ([]*models.Snippet, error){
	stmt := `
SELECT id, title, content, created, expires FROM snippets
WHERE expires > UTC_TIMESTAMP() AND userID = ? ORDER BY created
`
	snippets := []*models.Snippet{}
	rows, err := m.DB.Query(stmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	
	return snippets, nil
}
