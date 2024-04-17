package boilerplate

type Repository interface {
	RetrieveMessage() string
}

func NewRepository() Repository {
	return &repository{}
}

type repository struct {
}

func (r *repository) RetrieveMessage() string {
	return "message from repository"
}
