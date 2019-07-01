// Package model contains the types for schema 'public'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"gopkg.in/guregu/null.v3"
)

// Payment represents a row from 'public.payments'.
type Payment struct {
	ID              int       `db:"id" json:"id"`                               // id
	PaymentTypeID   null.Int  `db:"payment_type_id" json:"payment_type_id"`     // payment_type_id
	PaymentStatusID null.Int  `db:"payment_status_id" json:"payment_status_id"` // payment_status_id
	Amount          null.Int  `db:"amount" json:"amount"`                       // amount
	CreateDate      null.Time `db:"create_date" json:"create_date"`             // create_date
	UpdateDate      null.Time `db:"update_date" json:"update_date"`             // update_date
	MonthID         null.Int  `db:"month_id" json:"month_id"`                   // month_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Payment exists in the database.
func (p *Payment) Exists() bool {
	return p._exists
}

// Deleted provides information if the Payment has been deleted from the database.
func (p *Payment) Deleted() bool {
	return p._deleted
}

// Insert inserts the Payment to the database.
func (p *Payment) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.payments (` +
		`payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID)
	err = db.QueryRow(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID).Scan(&p.ID)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Update updates the Payment in the database.
func (p *Payment) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.payments SET (` +
		`payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6` +
		`) WHERE id = $7`

	// run query
	XOLog(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID, p.ID)
	_, err = db.Exec(sqlstr, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID, p.ID)
	return err
}

// Save saves the Payment to the database.
func (p *Payment) Save(db XODB) error {
	if p.Exists() {
		return p.Update(db)
	}

	return p.Insert(db)
}

// Upsert performs an upsert for Payment.
//
// NOTE: PostgreSQL 9.5+ only
func (p *Payment) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.payments (` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.payment_type_id, EXCLUDED.payment_status_id, EXCLUDED.amount, EXCLUDED.create_date, EXCLUDED.update_date, EXCLUDED.month_id` +
		`)`

	// run query
	XOLog(sqlstr, p.ID, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID)
	_, err = db.Exec(sqlstr, p.ID, p.PaymentTypeID, p.PaymentStatusID, p.Amount, p.CreateDate, p.UpdateDate, p.MonthID)
	if err != nil {
		return err
	}

	// set existence
	p._exists = true

	return nil
}

// Delete deletes the Payment from the database.
func (p *Payment) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !p._exists {
		return nil
	}

	// if deleted, bail
	if p._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.payments WHERE id = $1`

	// run query
	XOLog(sqlstr, p.ID)
	_, err = db.Exec(sqlstr, p.ID)
	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// Month returns the Month associated with the Payment's MonthID (month_id).
//
// Generated from foreign key 'month_id'.
func (p *Payment) Month(db XODB) (*Month, error) {
	return MonthByID(db, int(p.MonthID.Int64))
}

// PaymentStatus returns the PaymentStatus associated with the Payment's PaymentStatusID (payment_status_id).
//
// Generated from foreign key 'payments_payment_status_id_fkey'.
func (p *Payment) PaymentStatus(db XODB) (*PaymentStatus, error) {
	return PaymentStatusByID(db, int(p.PaymentStatusID.Int64))
}

// PaymentType returns the PaymentType associated with the Payment's PaymentTypeID (payment_type_id).
//
// Generated from foreign key 'payments_payment_type_id_fkey'.
func (p *Payment) PaymentType(db XODB) (*PaymentType, error) {
	return PaymentTypeByID(db, int(p.PaymentTypeID.Int64))
}

// PaymentByID retrieves a row from 'public.payments' as a Payment.
//
// Generated from index 'payments_pkey'.
func PaymentByID(db XODB, id int) (*Payment, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, payment_type_id, payment_status_id, amount, create_date, update_date, month_id ` +
		`FROM public.payments ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	p := Payment{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&p.ID, &p.PaymentTypeID, &p.PaymentStatusID, &p.Amount, &p.CreateDate, &p.UpdateDate, &p.MonthID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
