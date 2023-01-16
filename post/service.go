package post

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(post *Post) error {
	if err := s.repo.Save(*post); err != nil {
		return err
	}

	return nil

}
