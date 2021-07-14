package usecase

import (
	"phone/domain"
	"phone/infra"
)

type Entry interface {
	AddEntry(id domain.Id, name domain.Name, number domain.Phone) error
	GetAllEntries() ([]domain.Entry, error)
}

type entryImpl struct {
	db domain.EntryRepository
}

func NewEntry() Entry {
	db := infra.NewSQLiteRepository()
	return &entryImpl{db: db}
}

func (e entryImpl) AddEntry(id domain.Id, name domain.Name, number domain.Phone) error {
	entry := domain.Entry{
		Id:    id,
		Name:  name,
		Phone: number,
	}

	return e.db.Save(entry)
}

func (e entryImpl) GetAllEntries() ([]domain.Entry, error) {
	return e.db.GetAll()
}
