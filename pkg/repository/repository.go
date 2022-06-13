package repository

type Authorization interface {

}

type Payment interface {

}

type Repository struct {
	Authorization
	Payment
}

func NewRepository() *Repository {
	return &Repository{}
}
