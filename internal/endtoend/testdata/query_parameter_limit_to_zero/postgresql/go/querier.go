// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0

package querytest

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)
	DeleteAuthor(ctx context.Context, arg DeleteAuthorParams) error
	GetAuthor(ctx context.Context, arg GetAuthorParams) (Author, error)
	ListAuthors(ctx context.Context) ([]Author, error)
}

var _ Querier = (*Queries)(nil)
