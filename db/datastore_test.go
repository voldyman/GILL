package db

import (
	"testing"
	"reflect"
	"time"
)

func TestDataStore(t *testing.T) {
	testDir := "/tmp/" + string(time.Now().UTC().UnixNano())

	db := GetDB(testDir)
	db.AddUser("testNick", "127.0.0.1", "test")

	recs, err := db.GetUserForNick("testNick")
	
	if err != nil {
		t.Fatal(err)
	}
	if len(recs) != 1 {
		t.Fatal("Number of Records isn't equal")
	}
}

func TestQuery(t *testing.T) {
	query := make(map[string]interface{}, 0)
	query["re"] = "test"
	query["in"] = []interface{} {"test"}

	testQuery := createQuery("test", "test")

	if !reflect.DeepEqual(query, testQuery) {
		t.Fatal("Query Doesn't match")
	}
}
