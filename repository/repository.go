//use cases, interface - for relaition layer
type User interface {
	CreateUser(models.User) error
}

type Account interface {
}
type Repository struct {
	User
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		User: NewUserRepository(db),
	}
}