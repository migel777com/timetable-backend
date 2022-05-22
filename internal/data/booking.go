package data

import (
	"database/sql"
	"errors"
	"time"
)

type Booking struct {
	Id int64 `json:"id"`
	Room string `json:"room"`
	RoomId int64 `json:"room_id"`
	Reserver string `json:"reserver"`
	ReserverId string `json:"reserver_id"`
	ReserverInfo string `json:"reserver_info"`
	Day string `json:"day"`
	Date time.Time `json:"date"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
	Reason string `json:"reason"`
	Confirmed bool `json:"confirmed"`
	CreatedTime time.Time `json:"created_time"`
}

type BookingModel struct {
	DB *sql.DB
}

func (m *BookingModel) GetAll() ([]*Booking, error){
	stmt := `SELECT id, room, room_id, reserver, reserver_id, reserver_info, day, date, start_time, end_time, reason, confirmed, created_time FROM booking ORDER BY date;`
	rows, err := m.DB.Query(stmt)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		b := &Booking{}
		err = rows.Scan(&b.Id, &b.Room, &b.RoomId, &b.Reserver, &b.ReserverId, &b.ReserverInfo, &b.Day, &b.Date, &b.StartTime, &b.EndTime, &b.Reason, &b.Confirmed, &b.CreatedTime)
		if err!=nil{
			return nil, err
		}
		bookings = append(bookings, b)
	}

	if len(bookings)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (m *BookingModel) GetAllByRoom(roomId int64) ([]*Booking, error){
	stmt := `SELECT id, room, room_id, reserver, reserver_id, reserver_info, day, date, start_time, end_time, reason, confirmed, created_time FROM booking WHERE room_id = ? ORDER BY date;`
	rows, err := m.DB.Query(stmt, roomId)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		b := &Booking{}
		err = rows.Scan(&b.Id, &b.Room, &b.RoomId, &b.Reserver, &b.ReserverId, &b.ReserverInfo, &b.Day, &b.Date, &b.StartTime, &b.EndTime, &b.Reason, &b.Confirmed, &b.CreatedTime)
		if err!=nil{
			return nil, err
		}
		bookings = append(bookings, b)
	}

	if len(bookings)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (m *BookingModel) GetAllByReserver(reserverId string) ([]*Booking, error){
	stmt := `SELECT id, room, room_id, reserver, reserver_id, reserver_info, day, date, start_time, end_time, reason, confirmed, created_time FROM booking WHERE reserver_id = ? ORDER BY date;`
	rows, err := m.DB.Query(stmt, reserverId)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		b := &Booking{}
		err = rows.Scan(&b.Id, &b.Room, &b.RoomId, &b.Reserver, &b.ReserverId, &b.ReserverInfo, &b.Day, &b.Date, &b.StartTime, &b.EndTime, &b.Reason, &b.Confirmed, &b.CreatedTime)
		if err!=nil{
			return nil, err
		}
		bookings = append(bookings, b)
	}

	if len(bookings)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (m *BookingModel) GetAllByDate(date time.Time) ([]*Booking, error){
	stmt := `SELECT id, room, room_id, reserver, reserver_id, reserver_info, day, date, start_time, end_time, reason, confirmed, created_time FROM booking WHERE DATE(date) = DATE(?) ORDER BY date;`
	rows, err := m.DB.Query(stmt, date)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		b := &Booking{}
		err = rows.Scan(&b.Id, &b.Room, &b.RoomId, &b.Reserver, &b.ReserverId, &b.ReserverInfo, &b.Day, &b.Date, &b.StartTime, &b.EndTime, &b.Reason, &b.Confirmed, &b.CreatedTime)
		if err!=nil{
			return nil, err
		}
		bookings = append(bookings, b)
	}

	if len(bookings)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (m *BookingModel) GetAllConfirmed() ([]*Booking, error){
	stmt := `SELECT id, room, room_id, reserver, reserver_id, reserver_info, day, date, start_time, end_time, reason, created_time FROM booking WHERE confirmed = 1 ORDER BY date;`
	rows, err := m.DB.Query(stmt)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		b := &Booking{}
		err = rows.Scan(&b.Id, &b.Room, &b.RoomId, &b.Reserver, &b.ReserverId, &b.ReserverInfo, &b.Day, &b.Date, &b.StartTime, &b.Reason, &b.EndTime, &b.CreatedTime)
		if err!=nil{
			return nil, err
		}
		bookings = append(bookings, b)
	}

	if len(bookings)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (m *BookingModel) GetAllRequests() ([]*Booking, error){
	stmt := `SELECT id, room, room_id, reserver, reserver_id, reserver_info, day, date, start_time, end_time, reason, created_time FROM booking WHERE confirmed = 0 ORDER BY date;`
	rows, err := m.DB.Query(stmt)
	if err!=nil{
		return nil, err
	}
	defer rows.Close()

	var bookings []*Booking
	for rows.Next() {
		b := &Booking{}
		err = rows.Scan(&b.Id, &b.Room, &b.RoomId, &b.Reserver, &b.ReserverId, &b.ReserverInfo, &b.Day, &b.Date, &b.StartTime, &b.Reason, &b.EndTime, &b.CreatedTime)
		if err!=nil{
			return nil, err
		}
		bookings = append(bookings, b)
	}

	if len(bookings)==0{
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}

func (m *BookingModel) Insert(room, reserver, reserverInfo, day, startTime, endTime, reason string, roomId int64, reserverId string, date time.Time) (int, error) {
	stmt := `INSERT INTO booking (room, room_id, reserver, reserver_id, reserver_info, day, date, start_time, end_time, reason) VALUES (?,?,?,?,?,?,?,?,?,?)`

	result, err := m.DB.Exec(stmt, room, roomId, reserver, reserverId, reserverInfo, day, date, startTime, endTime, reason)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *BookingModel) Confirm(id int64) error {
	stmt := `UPDATE booking SET confirmed = 1 WHERE id = ?`

	result, err := m.DB.Exec(stmt, id)

	if err != nil {
		return err
	}

	if temp, _ := result.RowsAffected(); temp == 0 {
		return errors.New("no affected rows")
	}

	return nil
}

func (m *BookingModel) Delete(id int64) error {
	stmt := `DELETE FROM booking WHERE id = ?;`

	result, err := m.DB.Exec(stmt, id)

	if err != nil {
		return err
	}

	if temp, _ := result.RowsAffected(); temp == 0 {
		return errors.New("no affected rows")
	}

	return nil
}
