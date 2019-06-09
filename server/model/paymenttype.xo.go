// Package model contains the types for schema 'public'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// PaymentType represents a row from 'public.payment_types'.
type PaymentType struct {
	ID         int            `json:"id"`          // id
	Name       sql.NullString `json:"name"`        // name
	CreateDate pq.NullTime    `json:"create_date"` // create_date
	UpdateDate pq.NullTime    `json:"update_date"` // update_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the PaymentType exists in the database.
func (pt *PaymentType) Exists() bool {
	return pt._exists
}

// Deleted provides information if the PaymentType has been deleted from the database.
func (pt *PaymentType) Deleted() bool {
	return pt._deleted
}

// Insert inserts the PaymentType to the database.
func (pt *PaymentType) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if pt._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.payment_types (` +
		`name, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, pt.Name, pt.CreateDate, pt.UpdateDate)
	err = db.QueryRow(sqlstr, pt.Name, pt.CreateDate, pt.UpdateDate).Scan(&pt.ID)
	if err != nil {
		return err
	}

	// set existence
	pt._exists = true

	return nil
}

// Update updates the PaymentType in the database.
func (pt *PaymentType) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pt._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if pt._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.payment_types SET (` +
		`name, create_date, update_date` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, pt.Name, pt.CreateDate, pt.UpdateDate, pt.ID)
	_, err = db.Exec(sqlstr, pt.Name, pt.CreateDate, pt.UpdateDate, pt.ID)
	return err
}

// Save saves the PaymentType to the database.
func (pt *PaymentType) Save(db XODB) error {
	if pt.Exists() {
		return pt.Update(db)
	}

	return pt.Insert(db)
}

// Upsert performs an upsert for PaymentType.
//
// NOTE: PostgreSQL 9.5+ only
func (pt *PaymentType) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if pt._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.payment_types (` +
		`id, name, create_date, update_date` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, name, create_date, update_date` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.name, EXCLUDED.create_date, EXCLUDED.update_date` +
		`)`

	// run query
	XOLog(sqlstr, pt.ID, pt.Name, pt.CreateDate, pt.UpdateDate)
	_, err = db.Exec(sqlstr, pt.ID, pt.Name, pt.CreateDate, pt.UpdateDate)
	if err != nil {
		return err
	}

	// set existence
	pt._exists = true

	return nil
}

// Delete deletes the PaymentType from the database.
func (pt *PaymentType) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !pt._exists {
		return nil
	}

	// if deleted, bail
	if pt._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.payment_types WHERE id = $1`

	// run query
	XOLog(sqlstr, pt.ID)
	_, err = db.Exec(sqlstr, pt.ID)
	if err != nil {
		return err
	}

	// set deleted
	pt._deleted = true

	return nil
}

// PaymentTypeByID retrieves a row from 'public.payment_types' as a PaymentType.
//
// Generated from index 'payment_types_pkey'.
func PaymentTypeByID(db XODB, id int) (*PaymentType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, create_date, update_date ` +
		`FROM public.payment_types ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	pt := PaymentType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&pt.ID, &pt.Name, &pt.CreateDate, &pt.UpdateDate)
	if err != nil {
		return nil, err
	}

	return &pt, nil
}
