package infra

import (
	"database/sql"
	"fmt"
	"phone/domain"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteDb struct {
	db *sql.DB
}

func (s sqliteDb) GetAll() ([]domain.Entry, error) {
	const tp = "SELECT * FROM entry;"

	rows, err := s.db.Query(tp)
	if err != nil {
		return nil, fmt.Errorf("infra: データ取得に失敗しました %w", err)
	}

	var result []domain.Entry
	for rows.Next() {
		var e domain.Entry
		if err := rows.Scan(&e.Id, &e.Name, &e.Phone); err != nil {
			return nil, fmt.Errorf("infra: Scan()に失敗しました %w", err)
		}
		result = append(result, e)
	}

	return result, nil
}

func (s sqliteDb) GetById(target domain.Id) (domain.Entry, error) {
	const tp = "SELECT * FROM entry WHERE id = ?;"
	row := s.db.QueryRow(tp, target)

	if row.Err() != nil {
		return domain.Entry{}, fmt.Errorf("infra: レコードが存在しません %w", row.Err())
	}

	id := domain.Id(0)
	name := domain.Name("")
	phone := domain.Phone("")

	err := row.Scan(&id, &name, &phone)

	if err != nil {
		return domain.Entry{}, fmt.Errorf("infra: Scan()に失敗しました %w", row.Err())
	}

	return domain.Entry{
		Id: id, Name: name, Phone: phone,
	}, nil
}

func (s sqliteDb) Update(entry domain.Entry) error {
	const tp = "UPDATE entry SET name = ?, phone = ? WHERE id = ?;"
	_, err := s.db.Exec(tp, entry.Name, entry.Phone, entry.Id)

	if err != nil {
		return fmt.Errorf("infra: Updateに失敗しました %w", err)
	}

	return nil
}

func (s sqliteDb) Save(entry domain.Entry) error {
	res, err := s.GetById(entry.Id)
	if res.Id != 0 {
		return s.Update(entry)
	}

	const tp = "INSERT INTO entry(name, phone) values (?,?);"
	_, err = s.db.Exec(tp, entry.Name, entry.Phone)

	if err != nil {
		return fmt.Errorf("infra: DBへのセーブに失敗しました %w", err)
	}

	return nil
}

func NewSQLiteRepository() domain.EntryRepository {
	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS entry (
    id   INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    phone TEXT NOT NULL
)`)

	if err != nil {
		panic(err)
	}

	return sqliteDb{
		db: db,
	}
}
