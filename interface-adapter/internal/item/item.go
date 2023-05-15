package item

import "context"

type Module interface {
	Get(ctx context.Context, id int) (*Item, error)
	Set(ctx context.Context, item *Item) (err error)
}

type module struct {
	repo Repository
}

// NewItemModule creates a "service" layer.
// It could add some business logic but
// after all calls the Repository layer.
func NewItemModule(r Repository) (Module, error) {
	return &module{
		repo: r,
	}, nil
}

func (m *module) Get(ctx context.Context, id int) (*Item, error) {
	return m.repo.Get(ctx, id)
}

func (m *module) Set(ctx context.Context, item *Item) error {
	return m.repo.Set(ctx, item)
}
