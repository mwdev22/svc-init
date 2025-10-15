package store

import (
	"context"

	"github.com/mwdev22/rest-boilerplate-standard/internal/models"
)

// store contains interfaces for the repository layer
// it defines the methods that the repository should implement, allowing for easy swapping of different data sources

type UserRepository interface {
	Create(context.Context, *models.User) error
	// ... and so on for other methods and repos
}
