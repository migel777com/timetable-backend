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
	TeamsMeetingJoinurl string `json:"teams_meeting_joinurl"`
	SubjectId string `json:"subject_id"`
}

type TimetableModel struct {
	DB *sql.DB
}

func (m *TimetableModel) GetByWeekDay(weekDay string) ([]*Timetable, error) {
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, room, room_id, lesson_type, room_type, classtime_day, classtime_time, elective_id, teams_meeting_joinurl, subject_id FROM schedule_timetable WHERE classtime_day = ?;`
	rows, err := m.DB.Query(stmt, weekDay)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeTime, &t.ClasstimeDay, &t.ElectiveId, &t.TeamsMeetingJoinurl, &t.SubjectId)
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

func (m *TimetableModel) GetByGroup(groupId int64) ([]*Timetable, error) {
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, room, room_id, lesson_type, room_type, classtime_day, classtime_time, elective_id, teams_meeting_joinurl, subject_id FROM schedule_timetable WHERE id IN (SELECT timetable_id from schedule_timetable_groups WHERE group_id = ?);`
	rows, err := m.DB.Query(stmt, groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeTime, &t.ClasstimeDay, &t.ElectiveId, &t.TeamsMeetingJoinurl, &t.SubjectId)
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
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, room, room_id, lesson_type, room_type, classtime_day, classtime_time, elective_id, teams_meeting_joinurl, subject_id FROM schedule_timetable WHERE tutor_id = ?;`
	rows, err := m.DB.Query(stmt, tutorId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeTime, &t.ClasstimeDay, &t.ElectiveId, &t.TeamsMeetingJoinurl, &t.SubjectId)
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
	stmt := `SELECT id, schedule_block_id, subject, tutor, tutor_id, room, room_id, lesson_type, room_type, classtime_day, classtime_time, elective_id, teams_meeting_joinurl, subject_id FROM schedule_timetable WHERE room_id = ?;`
	rows, err := m.DB.Query(stmt, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var timetable []*Timetable
	for rows.Next() {
		t := &Timetable{}
		err = rows.Scan(&t.Id, &t.ScheduleBlockId, &t.Subject, &t.Tutor, &t.TutorId, &t.Room, &t.RoomId, &t.LessonType, &t.RoomType, &t.ClasstimeTime, &t.ClasstimeDay, &t.ElectiveId, &t.TeamsMeetingJoinurl, &t.SubjectId)
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
