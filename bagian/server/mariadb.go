package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addBagian         = `insert into tbBagian(IDBagian,NamaBagian, Deskripsi, Status, CreatedBy, CreatedOn, UpdatedBy, UpdateOn)values (?,?,?,?,?,?,?,?,?,?,?,?)`
	selectBagianAktif = `select IDBagian,NamaBagian, Deskripsi, Status, CreatedBy, CreatedOn, UpdatedBy, UpdateOn from tbBagian where status=1`
	selectBagian      = `select IDBagian,NamaBagian, Deskripsi, Status, CreatedBy, CreatedOn, UpdatedBy, UpdateOn from tbBagian`
	updateBagian      = `update tbBagian set NamaBagian=?, Deskripsi=?, Status=?, CreatedBy=?, CreatedOn=?, UpdatedBy=?, UpdateOn=? where IDBagian=?`
)

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddBagian(bagian Bagian) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addBagian, bagian.NamaBagian, bagian.Deskripsi, OnAdd, "1", time.Now(), "1", time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadBagianAktif() (Bagians, error) {
	bagian := Bagians{}
	rows, _ := rw.db.Query(selectBagianAktif)
	defer rows.Close()
	for rows.Next() {
		var k Bagian
		err := rows.Scan(&k.IDBagian, &k.NamaBagian, &k.Deskripsi, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdateOn)
		if err != nil {
			fmt.Println("error query:", err)
			return bagian, err
		}
		bagian = append(bagian, k)
	}
	//fmt.Println("db nya:", customer)
	return bagian, nil
}

func (rw *dbReadWriter) ReadBagian() (Bagians, error) {
	bagian := Bagians{}
	rows, _ := rw.db.Query(selectBagian)
	defer rows.Close()
	for rows.Next() {
		var k Bagian
		err := rows.Scan(&k.IDBagian, &k.NamaBagian, &k.Deskripsi, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdateOn)
		if err != nil {
			fmt.Println("error query:", err)
			return bagian, err
		}
		bagian = append(bagian, k)
	}
	//fmt.Println("db nya:", customer)
	return bagian, nil
}

func (rw *dbReadWriter) UpdateBagian(kar Bagian) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateBagian, kar.IDBagian, kar.NamaBagian, kar.Deskripsi, kar.Status, kar.CreatedBy, kar.CreatedOn, kar.UpdatedBy, kar.UpdateOn)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
