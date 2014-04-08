package main

import (
        "github.com/voldyman/GILL/webserver"
        "github.com/voldyman/GILL/db"
        "fmt"
)

func main() {
        datastore := db.GetDB("/tmp/MyDB")

        datastore.AddUser("voldyman", "127.0.0.1", "woah")
        datastore.AddUser("ares", "127.0.0.1", "woah")
        datastore.AddUser("user", "127.0.0.1", "woah")

        users, err := datastore.GetUserForNick("a")
        if err != nil {
                panic(err)
        }

        for id, user := range users {
                fmt.Printf("id: %d\n", id)
                fmt.Println(user.Nick)
        }
        datastore.Close()
	webserver.StartServer()
}
