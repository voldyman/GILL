package db

import (
        "github.com/HouzuoGuo/tiedot/db"
        "os"
        "errors"
)

type DataStore struct {
        url string
        db  *db.DB
        users *db.Col
}

type Record struct {
        Nick string    `json: nick`
        IP string      `json: ip`
        Data string    `json: data`
}

func GetDB(path string) *DataStore {
        db, err := db.OpenDB(path)
        if err != nil {
                panic(err)
        }

	db.Create("Users", 2)
        users := db.Use("Users")

        users.Index([]string{"nick","ip"})

        return &DataStore{
                url: path,
                db: db,
                users: users,
        }
}

func (ds *DataStore) Close() {
        os.RemoveAll(ds.url)
        ds.db.Close()
}

func (ds *DataStore) AddUser(nick, ip , data string) {
        ds.users.Insert(map[string]interface{} {
                "nick": nick,
                "ip": ip,
                "data": data,
        })
}

func (ds *DataStore) getUserForID(id uint64) Record{
        var result *Record
        ds.users.Read(id, &result)
        return *result
}

func (ds *DataStore) getUsersForIDs(ids map[uint64]struct{}) []Record{
        var results []Record
        for id := range ids {
                results = append(results, ds.getUserForID(id))
        }
        return results
}

func (ds *DataStore) GetUserForNick(nick string) (recs []Record, err error){
	query := createQuery(nick, "nick")

        queryResult := make(map[uint64]struct{})
        err = db.EvalQuery(query, ds.users, &queryResult)
        if err != nil {
                return		
        }
	
        if len(queryResult) == 0 {
                err = errors.New("No User Found")
		return
        }

        recs = ds.getUsersForIDs(queryResult)
	return


}

func createQuery(needle, col string) map[string]interface{}{
	return map[string]interface{}{
		"re":needle,
		"in":[]interface{} {col},
	}
}
