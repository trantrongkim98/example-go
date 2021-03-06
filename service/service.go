package service

import (
	"github.com/trantrongkim98/example-go/service/book"
	"github.com/trantrongkim98/example-go/service/category"
	"github.com/trantrongkim98/example-go/service/lendbook"
	"github.com/trantrongkim98/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
	BookService     book.Service
	LendBookService lendbook.Service
}
