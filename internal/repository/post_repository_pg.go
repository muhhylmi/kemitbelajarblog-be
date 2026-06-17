package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kemitbelajar/kemitbelajarblog-backend/internal/model"
)

// postgresPostRepository implements PostRepository using PostgreSQL.
type postgresPostRepository struct {
	db *sql.DB
}

// NewPostgresPostRepository creates a new PostgreSQL-backed PostRepository.
func NewPostgresPostRepository(db *sql.DB) PostRepository {
	return &postgresPostRepository{db: db}
}

// selectColumns is the shared column list for SELECT queries.
const selectColumns = `
	p.id, p.title, p.summary, p.content, p.category,
	p.read_time, p.published_at, p.image, p.image_alt,
	p.shares, p.views, p.is_featured, p.status, p.created_at, p.updated_at,
	a.id, a.name, a.avatar, a.role
`

// scanPost scans a row into a Post struct.
func scanPost(scanner interface{ Scan(dest ...any) error }) (*model.Post, error) {
	var post model.Post
	err := scanner.Scan(
		&post.ID, &post.Title, &post.Summary, &post.Content, &post.Category,
		&post.ReadTime, &post.PublishedAt, &post.Image, &post.ImageAlt,
		&post.Shares, &post.Views, &post.IsFeatured, &post.Status, &post.CreatedAt, &post.UpdatedAt,
		&post.Author.ID, &post.Author.Name, &post.Author.Avatar, &post.Author.Role,
	)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postgresPostRepository) FindAll() ([]model.Post, error) {
	query := fmt.Sprintf(`
		SELECT %s
		FROM posts p
		JOIN authors a ON p.author_id = a.id
		ORDER BY p.created_at DESC
	`, selectColumns)

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts: %w", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		post, err := scanPost(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, *post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	// Return empty slice instead of nil
	if posts == nil {
		posts = []model.Post{}
	}

	return posts, nil
}

func (r *postgresPostRepository) FindByID(id string) (*model.Post, error) {
	query := fmt.Sprintf(`
		SELECT %s
		FROM posts p
		JOIN authors a ON p.author_id = a.id
		WHERE p.id = $1
	`, selectColumns)

	row := r.db.QueryRow(query, id)
	post, err := scanPost(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to query post by id: %w", err)
	}

	return post, nil
}

func (r *postgresPostRepository) FindByStatus(status string) ([]model.Post, error) {
	query := fmt.Sprintf(`
		SELECT %s
		FROM posts p
		JOIN authors a ON p.author_id = a.id
		WHERE p.status = $1
		ORDER BY p.created_at DESC
	`, selectColumns)

	rows, err := r.db.Query(query, status)
	if err != nil {
		return nil, fmt.Errorf("failed to query posts by status: %w", err)
	}
	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		post, err := scanPost(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan post: %w", err)
		}
		posts = append(posts, *post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	if posts == nil {
		posts = []model.Post{}
	}

	return posts, nil
}

func (r *postgresPostRepository) Create(post *model.Post) error {
	// Default author: look up first author or use the default Julian Thorne
	var authorID string
	err := r.db.QueryRow(`SELECT id FROM authors WHERE name = 'Julian Thorne' LIMIT 1`).Scan(&authorID)
	if err != nil {
		// Fallback: get any author
		err = r.db.QueryRow(`SELECT id FROM authors ORDER BY created_at ASC LIMIT 1`).Scan(&authorID)
		if err != nil {
			return fmt.Errorf("no authors found in database: %w", err)
		}
	}

	now := time.Now()
	_, err = r.db.Exec(`
		INSERT INTO posts (id, title, summary, content, category, author_id, read_time, published_at, image, image_alt, shares, views, is_featured, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
	`,
		post.ID, post.Title, post.Summary, post.Content, post.Category,
		authorID, post.ReadTime, post.PublishedAt, post.Image, post.ImageAlt,
		post.Shares, post.Views, post.IsFeatured, post.Status, now, now,
	)
	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}

	// Load the author data back into the post
	err = r.db.QueryRow(`SELECT id, name, avatar, role FROM authors WHERE id = $1`, authorID).Scan(
		&post.Author.ID, &post.Author.Name, &post.Author.Avatar, &post.Author.Role,
	)
	if err != nil {
		return fmt.Errorf("failed to load author for created post: %w", err)
	}

	post.CreatedAt = now
	post.UpdatedAt = now

	return nil
}

func (r *postgresPostRepository) Update(post *model.Post) error {
	now := time.Now()
	result, err := r.db.Exec(`
		UPDATE posts SET
			title = $1, summary = $2, content = $3, category = $4,
			read_time = $5, published_at = $6, image = $7, image_alt = $8,
			shares = $9, views = $10, is_featured = $11, status = $12, updated_at = $13
		WHERE id = $14
	`,
		post.Title, post.Summary, post.Content, post.Category,
		post.ReadTime, post.PublishedAt, post.Image, post.ImageAlt,
		post.Shares, post.Views, post.IsFeatured, post.Status, now,
		post.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("post not found: %s", post.ID)
	}

	post.UpdatedAt = now
	return nil
}

func (r *postgresPostRepository) Delete(id string) error {
	result, err := r.db.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("post not found: %s", id)
	}

	return nil
}

func (r *postgresPostRepository) IncrementShares(id string) error {
	result, err := r.db.Exec(`UPDATE posts SET shares = shares + 1, updated_at = NOW() WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to increment shares: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("post not found: %s", id)
	}

	return nil
}

func (r *postgresPostRepository) IncrementViews(id string) error {
	result, err := r.db.Exec(`UPDATE posts SET views = views + 1, updated_at = NOW() WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to increment views: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("post not found: %s", id)
	}

	return nil
}
