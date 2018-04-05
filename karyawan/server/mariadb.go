package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addKaryawan                = `insert into tbKaryawan(Nama, Alamat, Telepon, Email, IDJabatan, IDBagian, Status, CreatedBy, CreatedOn, UpdatedBy, UpdatedOn, Keterangan)values (?,?,?,?,?,?,?,?,?,?,?)`
	selectKaryawanByJabatan    = `select IDKaryawan, Nama, Alamat, Telepon, Email, IDJabatan, IDBagian, Status, CreatedBy, CreatedOn, UpdatedBy, UpdatedOn, Keterangan from tbKaryawan where IDJabatan = ?`
	selectKaryawanByBagian     = `select IDKaryawan, Nama, Alamat, Telepon, Email, IDJabatan, IDBagian, Status, CreatedBy, CreatedOn, UpdatedBy, UpdatedOn, Keterangan from tbKaryawan where IDBagian=?`
	selectKaryawanByKeterangan = `select IDKaryawan, Nama, Alamat, Telepon, Email, IDJabatan, IDBagian, Status, CreatedBy, CreatedOn, UpdatedBy, UpdatedOn, Keterangan from tbKaryawan where Keterangan LIKE ?`
	selectKaryawan             = `select IDKaryawan, Nama, Alamat, Telepon, Email, IDJabatan, IDBagian, Status, CreatedBy, CreatedOn, UpdatedBy, UpdatedOn, Keterangan from tbKaryawan where status=1`
	updateKaryawan             = `update tbKaryawan set Nama=?, Alamat=?, Telepon=?, Email=?, IDJabatan=?, IDBagian=?, Status=?, CreatedBy=?, CreatedOn=?, UpdatedBy=?, UpdatedOn=?, Keterangan=? where IDKaryawan=?`
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

func (rw *dbReadWriter) AddKaryawan(karyawan Karyawan) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addKaryawan, karyawan.Nama, karyawan.Alamat, karyawan.Telepon, karyawan.Email, karyawan.IDJabatan, karyawan.IDBagian, OnAdd, "1", time.Now(), "1", time.Now(), karyawan.Keterangan)
	//fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadKaryawanByJabatan(IDJabatan string) (Karyawans, error) {
	karyawan := Karyawans{}
	rows, _ := rw.db.Query(selectKaryawanByJabatan, IDJabatan)
	defer rows.Close()
	for rows.Next() {
		var k Karyawan
		err := rows.Scan(&k.IDKaryawan, &k.Nama, &k.Alamat, &k.Telepon, &k.Email, &k.IDJabatan, &k.IDBagian, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdatedOn, &k.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return karyawan, err
		}
		karyawan = append(karyawan, k)
	}
	//fmt.Println("db nya:", customer)
	return karyawan, nil
}

func (rw *dbReadWriter) ReadKaryawanByBagian(IDBagian string) (Karyawans, error) {
	karyawan := Karyawans{}
	rows, _ := rw.db.Query(selectKaryawanByBagian, IDBagian)
	defer rows.Close()
	for rows.Next() {
		var k Karyawan
		err := rows.Scan(&k.IDKaryawan, &k.Nama, &k.Alamat, &k.Telepon, &k.Email, &k.IDJabatan, &k.IDBagian, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdatedOn, &k.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return karyawan, err
		}
		karyawan = append(karyawan, k)
	}
	//fmt.Println("db nya:", customer)
	return karyawan, nil
}

func (rw *dbReadWriter) ReadKaryawanByKeterangan(ket string) (Karyawans, error) {
	karyawan := Karyawans{}
	rows, _ := rw.db.Query(selectKaryawanByKeterangan, ket)
	defer rows.Close()
	for rows.Next() {
		var k Karyawan
		err := rows.Scan(&k.IDKaryawan, &k.Nama, &k.Alamat, &k.Telepon, &k.Email, &k.IDJabatan, &k.IDBagian, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdatedOn, &k.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return karyawan, err
		}
		karyawan = append(karyawan, k)
	}
	//fmt.Println("db nya:", customer)
	return karyawan, nil
}

func (rw *dbReadWriter) ReadKaryawan() (Karyawans, error) {
	karyawan := Karyawans{}
	rows, _ := rw.db.Query(selectKaryawan)
	defer rows.Close()
	for rows.Next() {
		var k Karyawan
		err := rows.Scan(&k.IDKaryawan, &k.Nama, &k.Alamat, &k.Telepon, &k.Email, &k.IDJabatan, &k.IDBagian, &k.Status, &k.CreatedBy, &k.CreatedOn, &k.UpdatedBy, &k.UpdatedOn, &k.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return karyawan, err
		}
		karyawan = append(karyawan, k)
	}
	//fmt.Println("db nya:", customer)
	return karyawan, nil
}

func (rw *dbReadWriter) UpdateKaryawan(kar Karyawan) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateKaryawan, kar.Nama, kar.Alamat, kar.Telepon, kar.Email, kar.IDJabatan, kar.IDBagian, kar.Status, kar.CreatedBy, kar.CreatedOn, kar.UpdatedBy, kar.UpdatedOn, kar.Keterangan, kar.IDKaryawan)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
