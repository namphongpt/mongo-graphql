package graph

import (
	"mongo-graph/services/post"
	"mongo-graph/services/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver struct
type Resolver struct {
	UserService user.Service
	PostService post.Service
}
