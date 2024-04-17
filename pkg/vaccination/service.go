package boilerplate

type Service interface {
	GetMessage() string
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

func (s *service) GetMessage() string {
	return s.repository.RetrieveMessage()
}
