package data

import (
	"database/sql"
)

type User struct {
	Id          int64         `json:"id"`
	Email       string        `json:"email"`
	Organization string 	  `json:"organization"`
}

type UserModel struct {
	DB *sql.DB
}


//func (m *UserModel) GetAll() ([]*User, error) {
//	stmt := `SELECT id, email, password, phone, name, surname, photo_id, country, city, address, createdTime FROM users ORDER BY id;`
//	rows, err := m.DB.Query(stmt)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	var users []*User
//	for rows.Next() {
//		u := &User{}
//		err = rows.Scan(&u.Id, &u.Email, &u.Password, &u.Phone, &u.Name, &u.Surname, &u.PhotoId, &u.Country, &u.City, &u.Address, &u.CreatedTime)
//		if err != nil {
//			return nil, err
//		}
//		users = append(users, u)
//	}
//
//	if len(users) == 0 {
//		return nil, sql.ErrNoRows
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return users, nil
//}
//
//func (m *UserModel) Get(id int) (*User, error) {
//	stmt := `SELECT id, email, password, phone, name, surname, photo_id, country, city, address, createdTime FROM users WHERE id = ?;`
//	rows := m.DB.QueryRow(stmt, id)
//
//	u := &User{}
//	err := rows.Scan(&u.Id, &u.Email, &u.Password, &u.Phone, &u.Name, &u.Surname, &u.PhotoId, &u.Country, &u.City, &u.Address, &u.CreatedTime)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return u, nil
//}
//
//func (m *UserModel) GetByEmail(email string) (*User, error) {
//	stmt := `SELECT id, email, password, phone, name, surname, photo_id, country, city, address, createdTime FROM users WHERE email = ?;`
//	rows := m.DB.QueryRow(stmt, email)
//
//	u := &User{}
//	err := rows.Scan(&u.Id, &u.Email, &u.Password, &u.Phone, &u.Name, &u.Surname, &u.PhotoId, &u.Country, &u.City, &u.Address, &u.CreatedTime)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = rows.Err(); err != nil {
//		return nil, err
//	}
//
//	return u, nil
//}
//
//func (m *UserModel) Insert(email, password, phone, name, surname, country, city, address string) (int, error) {
//	stmt := `INSERT INTO users (email, password, phone, name, surname, country, city, address) VALUES (?,?,?,?,?,?,?,?)`
//
//	result, err := m.DB.Exec(stmt, email, password, phone, name, surname, country, city, address)
//	if err != nil {
//		return 0, err
//	}
//
//	id, err := result.LastInsertId()
//	if err != nil {
//		return 0, err
//	}
//
//	return int(id), nil
//}
//
//func (m *UserModel) UpdateAll(id int, email, password, phone, name, surname, country, city, address string) error {
//	stmt := `UPDATE users SET email=?, password=?, phone=?, name=?, surname=?, country=?, city=?, address=? WHERE id = ?`
//
//	result, err := m.DB.Exec(stmt, email, password, phone, name, surname, country, city, address, id)
//
//	if err != nil {
//		return err
//	}
//
//	if temp, _ := result.RowsAffected(); temp == 0 {
//		return errors.New("no affected rows")
//	}
//
//	return nil
//}
//
//func (m *UserModel) UpdatePhotoId(id int, photoId int) error {
//	stmt := `UPDATE users SET photo_id=? WHERE id = ?`
//
//	result, err := m.DB.Exec(stmt, photoId, id)
//
//	if err != nil {
//		return err
//	}
//
//	if temp, _ := result.RowsAffected(); temp == 0 {
//		return errors.New("no affected rows")
//	}
//
//	return nil
//}
//
//func (m *UserModel) Delete(id int) error {
//	stmt := `DELETE FROM users WHERE id = ?;`
//
//	result, err := m.DB.Exec(stmt, id)
//
//	if err != nil {
//		return err
//	}
//
//	if temp, _ := result.RowsAffected(); temp == 0 {
//		return errors.New("no affected rows")
//	}
//
//	return nil
//}
