package item

import (
	"context"

	"github.com/eltonjr/go-experiments/decorator-pattern/pkg/log"
)

type itemLogger struct {
	impl   Module
	logger log.Logger
}

// NewItemLogger receives a Module and also implements
// the Module interface.
// When called the methods of the Module, it adapts
// with a log and calls the real implementation.
func NewItemLogger(l log.Logger, m Module) Module {
	l.Debug("[Module] New Module created\n")
	return &itemLogger{
		logger: l,
		impl:   m,
	}
}

func (m *itemLogger) Get(ctx context.Context, id int) (*Item, error) {
	defer m.logger.Debugf("[Module] Get item %d\n", id)

	return m.impl.Get(ctx, id)
}

func (m *itemLogger) Set(ctx context.Context, item *Item) error {
	defer m.logger.Debugf("[Module] Set item %d\n", item.ID)

	return m.impl.Set(ctx, item)
}
