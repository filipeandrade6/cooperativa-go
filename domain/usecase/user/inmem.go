package user

import (
	"strings"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type inmem struct {
	m map[entity.ID]*entity.User
}

func newInmem() *inmem {
	return &inmem{
		m: map[entity.ID]*entity.User{},
	}
}

func (i *inmem) GetUserByID(id entity.ID) (*entity.User, error) {
	bp, ok := i.m[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	return bp, nil
}

func (i *inmem) SearchUser(query string) ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range i.m {
		if strings.Contains(strings.ToLower(j.FirstName), query) {
			d = append(d, j)
		}
	}

	return d, nil
}

func (i *inmem) ListUser() ([]*entity.User, error) {
	var d []*entity.User
	for _, j := range i.m {
		d = append(d, j)
	}

	return d, nil
}

func (i *inmem) CreateUser(e *entity.User) (entity.ID, error) {
	i.m[e.ID] = e

	return e.ID, nil
}

func (i *inmem) UpdateUser(e *entity.User) error {
	_, err := i.GetUserByID(e.ID)
	if err != nil {
		return err
	}

	i.m[e.ID] = e

	return nil
}

func (i *inmem) DeleteUser(id entity.ID) error {
	if i.m[id] == nil {
		return entity.ErrNotFound
	}

	delete(i.m, id)

	return nil
}