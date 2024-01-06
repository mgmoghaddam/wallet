package member

import (
	"fmt"
	"time"
	"wallet/internal/serr"
)

const memberColumns = "id,first_name,last_name,email,phone,created_at,updated_at"

type Member struct {
	ID        int64     `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Phone     string    `db:"phone"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (s Storage) Create(u *Member) error {
	sqlStmt := `
	INSERT INTO member (first_name, last_name, email, phone) VALUES ($1, $2, $3, $4) 
	                     RETURNING id, created_at, updated_at`
	err := s.db.QueryRow(sqlStmt, u.FirstName, u.LastName, u.Email, u.Phone).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s Storage) Update(u *Member) error {
	sqlStmt := `
	UPDATE member SET first_name = $1, last_name = $2, email = $3, phone = $4, updated_at = now() WHERE id = $5
	                     RETURNING updated_at`
	err := s.db.QueryRow(sqlStmt, u.FirstName, u.LastName, u.Email, u.Phone, u.ID).Scan(&u.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s Storage) getAllByPage(limit, offset int, count bool) ([]*Member, int, error) {
	var total int
	if count {
		err := s.db.QueryRow("SELECT count(*) FROM member").Scan(&total)
		if err != nil {
			return nil, 0, serr.DBError("List", "member", err)
		}
	}
	pagination := fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)
	order := " ORDER BY created_at DESC"
	rows, err := s.db.Query("SELECT " + memberColumns + " FROM member" + order + pagination)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	members := make([]*Member, 0)
	for rows.Next() {
		u, err := s.ScanMember(rows)
		if err != nil {
			return nil, 0, serr.DBError("List", "member", err)
		}
		members = append(members, u)
	}
	return members, total, nil
}

func (s Storage) GetById(id int64) (*Member, error) {
	query := `
	SELECT ` + memberColumns + ` FROM member WHERE id = $1`
	u := &Member{}
	row := s.db.QueryRow(query, id)
	u, err := s.ScanMember(row)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// get member by phone
func (s Storage) GetByPhone(phone string) (*Member, error) {
	query := `
	SELECT ` + memberColumns + ` FROM member WHERE phone = $1`
	u := &Member{}
	row := s.db.QueryRow(query, phone)
	u, err := s.ScanMember(row)
	if err != nil {
		return nil, err
	}
	return u, nil
}
