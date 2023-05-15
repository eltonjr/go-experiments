package item

import (
	"context"

	"github.com/eltonjr/go-experiments/interface-adapter/pkg/log"
)

type repoLogger struct {
	impl   Repository
	logger log.Logger
}

// NewRepoLogger receives a Repository and also implements
// the Repository interface.
// When called the methods of the Repository, it adapts
// with a log and calls the real implementation.
func NewRepoLogger(l log.Logger, r Repository) Repository {
	l.Debug("[Repository] New Repository created\n")
	return &repoLogger{
		logger: l,
		impl:   r,
	}
}

func (l *repoLogger) Get(ctx context.Context, id int) (*Item, error) {
	defer l.logger.Debugf("[Repository] Get item %d\n", id)

	return l.impl.Get(ctx, id)
}

func (l *repoLogger) Set(ctx context.Context, item *Item) error {
	defer l.logger.Debugf("[Repository] Set item %d\n", item.ID)

	return l.impl.Set(ctx, item)
}
