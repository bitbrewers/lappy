package lappy

import (
	"database/sql"
	"errors"
	"time"

	"github.com/bitbrewers/tranx2"
	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	dbMap *gorp.DbMap
}

var (
	ErrNoOngoingRace = errors.New("no ongoing race found")
	ErrRaceOngoing   = errors.New("race already ongoing")
)

func NewSqliteStorage(dsn string) (*SqliteStorage, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(Record{}, "records").SetKeys(true, "ID")
	dbmap.AddTableWithName(Race{}, "races").SetKeys(true, "ID")
	if err := dbmap.CreateTablesIfNotExists(); err != nil {
		return nil, err
	}

	return &SqliteStorage{dbMap: dbmap}, nil
}

func (s *SqliteStorage) Save(passing tranx2.Passing) (rec Record, err error) {
	tx, err := s.dbMap.Begin()
	if err != nil {
		return
	}

	var raceID int64
	if err = tx.SelectOne(&raceID, "SELECT id FROM races WHERE ended IS NULL"); err == sql.ErrNoRows {
		tx.Commit()
		return rec, ErrNoOngoingRace
	}

	if err != nil {
		tx.Rollback()
		return
	}

	rec.RaceID = raceID
	rec.Passing = passing

	if err = tx.Insert(&rec); err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (s *SqliteStorage) StartRace(started time.Time) (race Race, err error) {
	tx, err := s.dbMap.Begin()
	if err != nil {
		return
	}

	// No error means that there is already race on going
	var raceID int64
	if err = tx.SelectOne(&raceID, "SELECT id FROM races WHERE ended IS NULL"); err == nil {
		tx.Commit()
		return race, ErrRaceOngoing
	}

	// sql.ErrNoRows is the error we should have otherwise something broke
	if err != sql.ErrNoRows {
		tx.Rollback()
		return
	}

	race.Started = &started
	if err = tx.Insert(&race); err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}

func (s *SqliteStorage) EndRace(ended time.Time) (err error) {
	tx, err := s.dbMap.Begin()
	if err != nil {
		return
	}

	race := &Race{}
	if err = tx.SelectOne(&race, "SELECT * FROM races WHERE ended IS NULL"); err == sql.ErrNoRows {
		tx.Commit()
		return ErrNoOngoingRace
	}

	if err != nil {
		tx.Rollback()
		return
	}

	race.Ended = &ended
	if _, err = tx.Update(race); err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}
