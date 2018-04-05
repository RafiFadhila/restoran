package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addJabatan         = `insert into tbJabatan(IDJabatan, Jabatan, Gaji, Status, CreatedBy, CreatedOn, UpdatedBy, UpdateOn)values (?,?,?,?,?,?,?,?,?,?,?,?)`
	selectJabatanAktif = `select IDJabatan, Jabatan, Gaji, Status, CreatedBy, CreatedOn, UpdatedBy, UpdateOn from tbJabatan where status=1`
	selectJabatan      = `select IDJabatan, Jabatan, Gaji, Status, CreatedBy, CreatedOn, UpdatedBy, UpdateOn from tbJabatan`
	updateJabatan      = `update tbJabatan set Jabatan=?, Gaji=?, Status=?, CreatedBy=?, CreatedOn=?, UpdatedBy=?, UpdateOn=? where IDJabatan=?`
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

func (rw *dbReadWriter) AddJabatan(jabatan Jabatan) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addJabatan, jabatan.IDJabatan, jabatan.Jabatan, jabatan.Gaji, jabatan.Status, "1", time.Now(), "1", time.Now())
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadJabatanAktif() (Jabatans, error) {
	jabatan := Jabatans{}
	rows, _ := rw.db.Query(selectJabatanAktif)
	defer rows.Close()
	for rows.Next() {
		var k Jabatan
		err := rows.Scan(&k.IDJabatan, &k.Jabatan, &k.Gaji, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdateOn)
		if err != nil {
			fmt.Println("error query:", err)
			return jabatan, err
		}
		jabatan = append(jabatan, k)
	}
	//fmt.Println("db nya:", customer)
	return jabatan, nil
}

func (rw *dbReadWriter) ReadJabatan() (Jabatans, error) {
	jabatan := Jabatans{}
	rows, _ := rw.db.Query(selectJabatan)
	defer rows.Close()
	for rows.Next() {
		var k Jabatan
		err := rows.Scan(&k.IDJabatan, &k.Jabatan, &k.Gaji, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdateOn)
		if err != nil {
			fmt.Println("error query:", err)
			return jabatan, err
		}
		jabatan = append(jabatan, k)
	}
	//fmt.Println("db nya:", customer)
	return jabatan, nil
}

func (rw *dbReadWriter) UpdateJabatan(kar Jabatan) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateJabatan, kar.IDJabatan, kar.Jabatan, kar.Gaji, kar.Status, kar.CreatedBy, kar.CreatedOn, kar.UpdatedBy, kar.UpdateOn)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
