package tinydb

import (
	"log"
	"strconv"
	"testing"
)

func Test_basic_StoreGetKey(t *testing.T) {
	myDB := NewDB("/tmp/testdb")
	defer myDB.Close()

	key := "foo"
	value := "bar"

	if err := myDB.Set(key, value); err != nil {
		t.Errorf("Unexpected error while setting key %s", err.Error())
	}

	data, err := myDB.Get(key)
	if err != nil {
		t.Errorf("Unexpected error while getting key %s", err.Error())
	}

	if data != value {
		t.Errorf("Unexpected value, need: %s, got: %s", value, data)
	}
}

func Test_basic_massUpload(t *testing.T) {
	myDB := NewDB("/tmp/testdb")
	defer myDB.Close()

	for i := 0; i < 201; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		if err := myDB.Set(key, value); err != nil {
			log.Fatal(err)
		}
	}

	data, err := myDB.Get("key200")
	if err != nil {
		log.Fatal(err)
	}
	if data != "value200" {
		t.Errorf("Unexpected value, need: %s, got: %s", "value200", data)
	}

}

func Test_basic_overrideKey(t *testing.T) {
	myDB := NewDB("/tmp/testdb")
	defer myDB.Close()

	key := "foo"
	value := "bar"

	if err := myDB.Set(key, value); err != nil {
		log.Fatal(err)
	}

	newValue := "fuzzy"
	if err := myDB.Set(key, newValue); err != nil {
		log.Fatal(err)
	}
	data, err := myDB.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	if data != newValue {
		log.Fatal(err)
	}

}
