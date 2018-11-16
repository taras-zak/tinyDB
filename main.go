package tinydb

import (
	"fmt"
	"github.com/serverhorror/rog-go/reverse"
	"log"
	"os"
	"strings"
	"errors"
)

var ErrorNoSuchKey = errors.New("No such key")

type DB struct {
	source string
	f      *os.File
}

func NewDB(filename string) *DB {
	d := &DB{source: filename}
	d.setup()
	return d
}

func (d *DB) setup() {
	fd, err := os.OpenFile(d.source, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	d.f = fd
}

func (d *DB) Set(key string, value string) (err error) {
	entry := fmt.Sprintf("%s:%s\n", key, value)
	_, err = d.f.Write([]byte(entry))
	return
}

func (d *DB) Get(key string) (res string, err error) {
	scan := reverse.NewScanner(d.f)
	for scan.Scan() {
		entry := scan.Text()
		s := strings.Split(entry, ":")
		k, v := s[0], s[1]
		if key == k {
			return v, nil
		}
	}
	return res, ErrorNoSuchKey
}

func (d *DB) Close() {
	if err := d.f.Truncate(0); err != nil {
		log.Fatal(err)
	}
	d.f.Close()
}
