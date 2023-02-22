package ports

import "context"

type Repository[T interface{}] interface {
	GetAll(ctx context.Context,
		filter map[string]interface{}) []T
	GetFirst(ctx context.Context,
		filter map[string]interface{}) T
	Insert(ctx context.Context,
		entity T)
	Replace(ctx context.Context,
		filter map[string]interface{},
		entity T)
	DeleteAll(ctx context.Context,
		filter map[string]interface{})
}
