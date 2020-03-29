package todo

type Service interface {
	List() ([]*TODO, error)
	Insert(title, text string) error
}

type Repository interface {
	List() ([]*TODO, error)
	Insert(title, text string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) List() ([]*TODO, error) {
	list, err := s.repository.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *service) Insert(title, text string) error {
	if err := s.repository.Insert(title, text); err != nil {
		return err
	}
	return nil
}
