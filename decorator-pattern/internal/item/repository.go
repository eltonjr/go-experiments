package item

import (
	"context"
	"errors"
)

type Repository interface {
	Get(ctx context.Context, id int) (*Item, error)
	Set(ctx context.Context, item *Item) (err error)
}

var (
	ErrItemNotFound = errors.New("item not found")
	ErrMissingID    = errors.New("invalid item; missing ID")
)

type repo struct {
	db map[int]*Item
}

// NewRepository creates a local database just for
// the sake of the POC
func NewRepository() (Repository, error) {
	return &repo{
		db: make(map[int]*Item),
	}, nil
}

func (repo *repo) Get(_ context.Context, id int) (*Item, error) {
	i, ok := repo.db[id]
	if !ok {
		return nil, ErrItemNotFound
	}

	return i, nil
}

func (repo *repo) Set(_ context.Context, item *Item) error {
	if item.ID == 0 {
		return ErrMissingID
	}

	repo.db[item.ID] = item
	return nil
}
