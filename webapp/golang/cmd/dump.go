package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type Img struct {
	ID      int    `db:"id"`
	Imgdata []byte `db:"imgdata"`
	Mime    string `db:"mime"`
}

func main() {
	host := os.Getenv("ISUCONP_DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("ISUCONP_DB_PORT")
	if port == "" {
		port = "3306"
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("Failed to read DB port number from an environment variable ISUCONP_DB_PORT.\nError: %s", err.Error())
	}
	user := os.Getenv("ISUCONP_DB_USER")
	if user == "" {
		user = "root"
	}
	password := "root"
	dbname := os.Getenv("ISUCONP_DB_NAME")
	if dbname == "" {
		dbname = "isuconp"
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s.", err.Error())
	}
	// imgs := make([]Img, 0, 10056)
	var imgs []Img
	err = db.Select(&imgs, "SELECT id, imgdata, mime FROM `posts`")
	if err != nil {
		log.Fatalf("db select err %v", err)
	}
	fmt.Printf("cnt: %d", len(imgs))
	for _, img := range imgs {
		var filename string
		switch img.Mime {
		case "image/jpeg":
			filename = fmt.Sprintf("%d.jpg", img.ID)
		case "image/gif":
			filename = fmt.Sprintf("%d.gif", img.ID)
		case "image/png":
			filename = fmt.Sprintf("%d.png", img.ID)
		}

		f, err := os.Create(filename)
		if err != nil {
			log.Fatalf("file create error")
		}
		_, err = f.Write(img.Imgdata)
		if err != nil {
			log.Fatalf("file write error")
		}
	}

	defer db.Close()
}
