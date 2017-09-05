package pg

import (
	"roshar/services/users-svc/pkg/users"
)

// CreateInterests inserts list of intersts
func (ds *Datastore) CreateInterests(names []string) ([]users.Interest, error) {
	var interests []users.Interest
	now := ds.Timestamp()

	builder := ds.SQLBuilder.
		Insert("interests").
		Columns("name", "created_at", "updated_at").
		Suffix(ReturnAll)

	for _, name := range names {
		builder = builder.Values(name, now, now)
	}

	stmt, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	err = ds.Select(&interests, stmt, args...)
	if err != nil {
		return nil, err
	}

	return interests, nil
}

// DeleteInterests deletes all records in interests table
func (ds *Datastore) DeleteInterests() error {
	stmt, args, err := ds.SQLBuilder.Delete("interests").ToSql()
	if err != nil {
		return err
	}

	_, err = ds.Exec(stmt, args...)
	if err != nil {
		return err
	}

	return nil
}
