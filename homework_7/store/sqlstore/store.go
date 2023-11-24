package sqlstore

import (
	"database/sql"

	"github.com/UnderMountain96/ITEA_GO/student_testing/store"
)

type Store struct {
	db             *sql.DB
	testRepository *TestRepository
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Test() store.TestRepository {
	if s.testRepository != nil {
		return s.testRepository
	}

	s.testRepository = &TestRepository{
		store: s,
	}

	return s.testRepository
}
