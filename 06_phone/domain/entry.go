package domain

type Id int
type Name string
type Phone string

type Entry struct {
	Id    Id
	Name  Name
	Phone Phone
}

type EntryRepository interface {
	GetAll() ([]Entry, error)
	Save(Entry) error
}
