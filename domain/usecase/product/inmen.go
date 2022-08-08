package product

import (
	"strings"

	"github.com/filipeandrade6/cooperagro/domain/entity"
)

type inmem struct {
	m map[entity.ID]*entity.Product
}

func newInmem() *inmem {
	return &inmem{
		m: map[entity.ID]*entity.Product{},
	}
}

func (i *inmem) GetProductByID(id entity.ID) (*entity.Product, error) {
	bp, ok := i.m[id]
	if !ok {
		return nil, entity.ErrNotFound
	}

	return bp, nil
}

func (i *inmem) SearchProduct(query string) ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range i.m {
		if strings.Contains(strings.ToLower(j.Name), query) {
			d = append(d, j)
		}
	}

	return d, nil
}

func (i *inmem) ListProduct() ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range i.m {
		d = append(d, j)
	}

	return d, nil
}

func (i *inmem) CreateProduct(e *entity.Product) (entity.ID, error) {
	i.m[e.ID] = e

	return e.ID, nil
}

func (i *inmem) UpdateProduct(e *entity.Product) error {
	_, err := i.GetProductByID(e.ID)
	if err != nil {
		return err
	}

	i.m[e.ID] = e

	return nil
}

func (i *inmem) DeleteProduct(id entity.ID) error {
	if i.m[id] == nil {
		return entity.ErrNotFound
	}

	delete(i.m, id)

	return nil
}
