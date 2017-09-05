package users

import (
	"time"

	"roshar/core/errors"
	"roshar/utils/collection"
)

// Interest defines intersts of users
type Interest struct {
	ID        int64
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// CreateInterests creates interests in database
func (core *Core) CreateInterests(names []string) ([]Interest, error) {
	names = collection.LowerCaseStr(names)
	interests, err := core.ds.CreateInterests(names)
	if err != nil {
		return nil, errors.NewDBError(err)
	}

	return interests, nil
}

// DeleteInterests delets all records in interests table
func (core *Core) DeleteInterests() error {
	err := core.ds.DeleteInterests()
	if err != nil {
		return errors.NewDBError(err)
	}
	return nil
}
