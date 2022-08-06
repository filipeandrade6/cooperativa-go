package unitofmeasure

import (
	"strings"
	"time"

	"github.com/filipeandrade6/cooperagro/src/domain/entities"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetByID(id entities.ID) (*entities.UnitOfMeasure, error) {
	u, err := s.repo.GetByID(id)
	if u == nil {
		return nil, entities.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Service) Search(query string) ([]*entities.UnitOfMeasure, error) {
	unitsOfMeasure, err := s.repo.Search(strings.ToLower(query))
	if err != nil {
		return nil, err
	}
	if len(unitsOfMeasure) == 0 {
		return nil, entities.ErrNotFound
	}

	return unitsOfMeasure, nil
}

func (s *Service) List() ([]*entities.UnitOfMeasure, error) {
	unitsOfMeasure, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(unitsOfMeasure) == 0 {
		return nil, entities.ErrNotFound
	}

	return unitsOfMeasure, nil
}

func (s *Service) Create(name string) (entities.ID, error) {
	u, err := entities.NewUnitOfMeasure(name)
	if err != nil {
		return entities.NewID(), err
	}

	return s.repo.Create(u)
}

func (s *Service) Update(e *entities.UnitOfMeasure) error {
	if err := e.Validate(); err != nil {
		return err
	}

	e.UpdatedAt = time.Now()

	return s.repo.Update(e)
}

func (s *Service) Delete(id entities.ID) error {
	if _, err := s.GetByID(id); err != nil {
		return err
	}

	return s.repo.Delete(id)
}
