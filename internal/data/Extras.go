package data

import "database/sql"

type Group struct {
	Id int64 `json:"group_id"`
	Name string `json:"name"`
}

type Teacher struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Active bool `json:"active"`
	Email string `json:"email"`
}

type Room struct {
	Id int64 `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Floor int64 `json:"floor"`
	Seats int64 `json:"seats"`
	Type string `json:"type"`
}

type ExtrasModel struct {
	DB *sql.DB
}

func (m *ExtrasModel) GetAllGroups() ([]*Group, error){
	stmt := `SELECT DISTINCT(group_id), name FROM schedule_timetable_groups ORDER BY group_id;`
	rows, err := m.DB.Query(stmt)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var groups []*Group
	for rows.Next() {
		g := &Group{}
		err = rows.Scan(&g.Id, &g.Name)
		if err!=nil{
			return nil, err
		}
		groups = append(groups, g)
	}

	if len(groups)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func (m *ExtrasModel) GetAllTeachers() ([]*Teacher, error){
	stmt := `SELECT id, name, active, email FROM teacher ORDER BY name;`
	rows, err := m.DB.Query(stmt)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var teachers []*Teacher
	for rows.Next() {
		t := &Teacher{}
		err = rows.Scan(&t.Id, &t.Name, &t.Active, &t.Email)
		if err!=nil{
			return nil, err
		}
		teachers = append(teachers, t)
	}

	if len(teachers)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return teachers, nil
}

func (m *ExtrasModel) GetAllRooms() ([]*Room, error){
	stmt := `SELECT id, code, name, floor, seats, type FROM schedule_room ORDER BY id;`
	rows, err := m.DB.Query(stmt)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var rooms []*Room
	for rows.Next() {
		r := &Room{}
		err = rows.Scan(&r.Id, &r.Code, &r.Name, &r.Floor, &r.Seats, &r.Type)
		if err!=nil{
			return nil, err
		}
		rooms = append(rooms, r)
	}

	if len(rooms)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (m *ExtrasModel) GetRoom(id int64) (*Room, error) {
	stmt := `SELECT id, code, name, floor, seats, type FROM schedule_room WHERE id = ?;`
	rows := m.DB.QueryRow(stmt, id)

	r := &Room{}
	err := rows.Scan(&r.Id, &r.Code, &r.Name, &r.Floor, &r.Seats, &r.Type)
	if err != nil {
		return nil, err
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return r, nil
}
