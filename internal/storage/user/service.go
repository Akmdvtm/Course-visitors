package user

import (
	"context"
	"database/sql"
)

// storage implements contract Storage interface
type storage struct {
	// db database connection
	db *sql.DB
}

// NewStorage returns storage initiated struct
func NewStorage(db *sql.DB) *storage {
	return &storage{
		db: db,
	}
}

// CreateUser creates user in database
func (s *storage) CreateUser(ctx context.Context, user User) (int, error) {
	// define sql query
	query := `
		INSERT INTO users(name, discord_name, telegram_id)
		VALUES ($1, $2, $3) RETURNING id
	`

	var insertedID int

	// exec sql query, set placeholders and scan returned id
	err := s.db.QueryRowContext(ctx, query, user.Name, user.DiscordName, user.TelegramID).Scan(&insertedID)
	if err != nil {
		return 0, err
	}

	return insertedID, nil
}

func (s *storage) GetUser(id int) (*User, error) {
	// sql data request
	query := `SELECT id, name, discord_name, telegram_id, role FROM users WHERE id = ($1)`
	// exec sql query
	row := s.db.QueryRow(query, id)
	// Initialize a pointer to a structure
	user := &User{}
	// the value of the sql.Row value to the corresponding field in the User structure
	err := row.Scan(&user.ID, &user.Name, &user.DiscordName, &user.TelegramID, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *storage) GetUsers(id int, discord_name, telegram_name string) ([]*User, error) {
	query := `
	SELECT id, name, discord_name, telegram_id, role FROM users WHERE id = ($1) AND discord_name = ($2) AND telegram_id = ($3)
	`
	rows, err := s.db.Query(query, id, discord_name, telegram_name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.DiscordName, &user.TelegramID, &user.Role); err != nil {
			return users, err
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}
