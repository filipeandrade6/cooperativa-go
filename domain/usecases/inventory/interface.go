package inventory

import "github.com/filipeandrade6/cooperagro/domain/entities"

// TODO porquer duas interfaces e pq interface Reader e Writer separada?

// Driver Adapter (preciso que o repositorio consiga fazer isso)

type Reader interface {
	GetInventoryByID(id entities.ID) (*entities.Inventory, error)
	SearchInventory(query string) ([]*entities.Inventory, error)
	ListInventory() ([]*entities.Inventory, error)
}

type Writer interface {
	CreateInventory(e *entities.Inventory) entities.ID
	UpdateInventory(e *entities.Inventory) error
	DeleteInventory(id entities.ID) error
}

type Repository struct {
	Reader
	Writer
}

// Driven Adapter (o que o use case comanda)

type UseCase interface {
	GetByID(id entities.ID) (*entities.Inventory, error)
	Search(query string) ([]*entities.Inventory, error)
	List() ([]*entities.Inventory, error)
	Create(
		customerID,
		productID entities.ID,
		quantity int,
		unitOfMeasureID entities.ID,
	) entities.ID
	Update(e *entities.Inventory) error
	Delete(id entities.ID) error
}