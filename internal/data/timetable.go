package data

import "database/sql"

type Timetable struct {
	Id int64 `json:"id"`
	ScheduleBlockId int64 `json:"schedule_block_id"`
	Subject string `json:"subject"`
	Tutor string `json:"tutor"`
	TutorId int64 `json:"tutor_id"`
	Room string `json:"room"`
	RoomId int64 `json:"room_id"`
	LessonType string `json:"lesson_type"`
	RoomType string `json:"room_type"`
	ClasstimeDay string `json:"classtime_day"`
	ClasstimeTime string `json:"classtime_time"`
	ElectiveId int64 `json:"elective_id"`
	SubjectId string `json:"subject_id"`
}

type TimetableModel struct {
	DB *sql.DB
}

func (m *TimetableModel) GetByWeekDay(weekDay string) ([]*Timetable, error) {
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, IFNULL(room, ''), IFNULL(room_id, 0), IFNULL(lesson_type, ''), IFNULL(room_type, ''), classtime_day, classtime_time, elective_id, subject_id FROM schedule_timetable WHERE classtime_day = ? AND classtime_time IS NOT NULL AND schedule_block_id IN (SELECT id FROM schedule_block WHERE active = 1);`
	rows, err := m.DB.Query(stmt, weekDay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeDay, &t.ClasstimeTime, &t.ElectiveId, &t.SubjectId)
		if err != nil {
			return nil, err
		}
		timetable = append(timetable, t)
	}

	if len(timetable) == 0 {
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return timetable, nil
}

func (m *TimetableModel) GetByGroup(group string) ([]*Timetable, error) {
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, IFNULL(room, ''), IFNULL(room_id, 0), IFNULL(lesson_type, ''), IFNULL(room_type, ''), classtime_day, classtime_time, elective_id, subject_id FROM schedule_timetable WHERE id IN (SELECT timetable_id from schedule_timetable_groups WHERE name = ?) AND schedule_block_id IN (SELECT id FROM schedule_block WHERE active = 1) AND classtime_day IS NOT NULL AND classtime_time IS NOT NULL;`
	rows, err := m.DB.Query(stmt, group)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeDay, &t.ClasstimeTime, &t.ElectiveId, &t.SubjectId)
		if err != nil {
			return nil, err
		}
		timetable = append(timetable, t)
	}

	if len(timetable) == 0 {
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return timetable, nil
}

func (m *TimetableModel) GetByTutor(tutorId int64) ([]*Timetable, error) {
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, IFNULL(room, ''), IFNULL(room_id, 0), IFNULL(lesson_type, ''), IFNULL(room_type, ''), classtime_day, classtime_time, elective_id, subject_id FROM schedule_timetable WHERE tutor_id = ? AND schedule_block_id IN (SELECT id FROM schedule_block WHERE active = 1) AND classtime_day IS NOT NULL AND classtime_time IS NOT NULL;`
	rows, err := m.DB.Query(stmt, tutorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeDay, &t.ClasstimeTime, &t.ElectiveId, &t.SubjectId)
		if err != nil {
			return nil, err
		}
		timetable = append(timetable, t)
	}

	if len(timetable) == 0 {
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return timetable, nil
}

func (m *TimetableModel) GetByRoom(roomId int64) ([]*Timetable, error) {
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, IFNULL(room, ''), IFNULL(room_id, 0), IFNULL(lesson_type, ''), IFNULL(room_type, ''), classtime_day, classtime_time, elective_id, subject_id FROM schedule_timetable WHERE room_id = ? AND schedule_block_id IN (SELECT id FROM schedule_block WHERE active = 1) AND classtime_day IS NOT NULL AND classtime_time IS NOT NULL;`
	rows, err := m.DB.Query(stmt, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeDay, &t.ClasstimeTime, &t.ElectiveId, &t.SubjectId)
		if err != nil {
			return nil, err
		}
		timetable = append(timetable, t)
	}

	if len(timetable) == 0 {
		return nil, sql.ErrNoRows
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return timetable, nil
}
