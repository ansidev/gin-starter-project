package author

type IAuthorRepository interface {
	GetByID(id int64) (Author, error)
}
