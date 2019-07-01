// Package model contains the types for schema 'public'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"gopkg.in/guregu/null.v3"
)

// Month represents a row from 'public.months'.
type Month struct {
	ID         int         `db:"id" json:"id"`                   // id
	Display    null.String `db:"display" json:"display"`         // display
	CreateDate null.Time   `db:"create_date" json:"create_date"` // create_date
	UpdateDate null.Time   `db:"update_date" json:"update_date"` // update_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Month exists in the database.
func (m *Month) Exists() bool {
	return m._exists
}

// Deleted provides information if the Month has been deleted from the database.
func (m *Month) Deleted() bool {
	return m._deleted
}

// Insert inserts the Month to the database.
func (m *Month) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.months (` +
		`display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, m.Display, m.CreateDate, m.UpdateDate)
	err = db.QueryRow(sqlstr, m.Display, m.CreateDate, m.UpdateDate).Scan(&m.ID)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Update updates the Month in the database.
func (m *Month) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if m._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.months SET (` +
		`display, create_date, update_date` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, m.Display, m.CreateDate, m.UpdateDate, m.ID)
	_, err = db.Exec(sqlstr, m.Display, m.CreateDate, m.UpdateDate, m.ID)
	return err
}

// Save saves the Month to the database.
func (m *Month) Save(db XODB) error {
	if m.Exists() {
		return m.Update(db)
	}

	return m.Insert(db)
}

// Upsert performs an upsert for Month.
//
// NOTE: PostgreSQL 9.5+ only
func (m *Month) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if m._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.months (` +
		`id, display, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, display, create_date, update_date` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.display, EXCLUDED.create_date, EXCLUDED.update_date` +
		`)`

	// run query
	XOLog(sqlstr, m.ID, m.Display, m.CreateDate, m.UpdateDate)
	_, err = db.Exec(sqlstr, m.ID, m.Display, m.CreateDate, m.UpdateDate)
	if err != nil {
		return err
	}

	// set existence
	m._exists = true

	return nil
}

// Delete deletes the Month from the database.
func (m *Month) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !m._exists {
		return nil
	}

	// if deleted, bail
	if m._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.months WHERE id = $1`

	// run query
	XOLog(sqlstr, m.ID)
	_, err = db.Exec(sqlstr, m.ID)
	if err != nil {
		return err
	}

	// set deleted
	m._deleted = true

	return nil
}

// MonthByID retrieves a row from 'public.months' as a Month.
//
// Generated from index 'months_pkey'.
func MonthByID(db XODB, id int) (*Month, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, display, create_date, update_date ` +
		`FROM public.months ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	m := Month{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&m.ID, &m.Display, &m.CreateDate, &m.UpdateDate)
	if err != nil {
		return nil, err
	}

	return &m, nil
}
