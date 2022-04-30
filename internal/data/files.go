package data

import (
	"database/sql"
	"time"
)

type File struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	CreatedTime time.Time `json:"createdTime"`
}

type FileModel struct {
	DB *sql.DB
}

func (m *FileModel) GetAll() ([]*File, error){
	stmt := `SELECT id, name, type, path, createdTime FROM files ORDER BY id;`
	rows, err := m.DB.Query(stmt)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var files []*File
	for rows.Next() {
		f := &File{}
		err = rows.Scan(&f.Id, &f.Name, &f.Type, &f.Path, &f.CreatedTime)
		if err!=nil{
			return nil, err
		}
		files = append(files, f)
	}

	if len(files)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return files, nil
}

func (m *FileModel) Get(id int) (*File, error){
	stmt := `SELECT id, name, type, path, createdTime FROM files WHERE id = ?;`
	rows:= m.DB.QueryRow(stmt, id)


	f := &File{}
	err := rows.Scan(&f.Id, &f.Name, &f.Type, &f.Path, &f.CreatedTime)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return f, nil
}

func (m *FileModel) Insert(name, typ, path string) (int, error){
	stmt := `INSERT INTO files (name, type, path) VALUES (?,?,?)`

	result, err := m.DB.Exec(stmt, name, typ, path)
	if err!=nil{
		return 0, err
	}

	id, err := result.LastInsertId()
	if err!=nil{
		return 0, err
	}

	return int(id), nil
}
