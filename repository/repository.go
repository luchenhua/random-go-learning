package repository

// Repository is the interface of all the DB functions regardless of the DB type.
type Repository interface {
	Delete(id int64) (bool, error)
}
