package client

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Client struct {
	client *sql.DB
}

// My Name
type Name struct {
	Id    int    `json:"id"`
	First string `json:"firstname"`
	Last  string `json:"lastname"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"first"`
	Password string `json:"password"`
	Admin    int    `json:"admin"`
}

func New() *Client {

	c, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}

	conn := &Client{c}

	return conn
}

func init() {

	db := New()

	CreateTables(db)

	myName := &Name{Id: 1, First: "Happy", Last: "Gilmore"}
	user := &User{1, "jbot", "secretJB", 1}
	userStandard := &User{2, "regs", "password", 0}

	WriteName(db, myName)
	AddUser(db, user)
	AddUser(db, userStandard)

	db.client.Close()

}

func CreateTables(db *Client) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS names(
		Id INTEGER NOT NULL PRIMARY KEY,
		First TEXT,
		Last TEXT,
		InsertedDatetime DATETIME
	);
	`

	usr_table := `CREATE TABLE IF NOT EXISTS users(
		Id INTEGER NOT NULL PRIMARY KEY,
		Username TEXT,
		Password TEXT,
		Admin INTEGER,
		InsertedDatetime DATETIME
	);
	`

	_, err := db.client.Exec(sql_table)
	if err != nil {
		panic(err)
	}

	_, err = db.client.Exec(usr_table)
	if err != nil {
		panic(err)
	}

}

// Add a user to the db
func AddUser(db *Client, user *User) {

	sql_newUser := `
	INSERT OR REPLACE INTO users(
		Id,
		Username,
		Password,
		Admin,
		InsertedDatetime
	) values(?, ?, ?, ?, CURRENT_TIMESTAMP)`

	stmt, err := db.client.Prepare(sql_newUser)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(user.Id, user.Username, user.Password, user.Admin)
	if err2 != nil {
		panic(err2)
	}

}

func GetUser(db *Client, username string) User {
	sql_readall := `
	SELECT Id, Username, Password, Admin FROM users WHERE username = ?
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.client.Query(sql_readall, username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []User
	for rows.Next() {
		item := User{}
		err2 := rows.Scan(&item.Id, &item.Username, &item.Password, &item.Admin)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result[0]
}




// This creates a new row or updates an existing one
func WriteName(db *Client, name *Name) {
	sql_additem := `
	INSERT OR REPLACE INTO names(
		Id,
		First,
		Last,
		InsertedDatetime
	) values(?, ?, ?, CURRENT_TIMESTAMP)
	`

	stmt, err := db.client.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(name.Id, name.First, name.Last)
	if err2 != nil {
		panic(err2)
	}

}

// This deletes a given name id
func Delete(db *Client, id int) {
	sql_additem := `
	DELETE FROM names WHERE id = ?`

	stmt, err := db.client.Prepare(sql_additem)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(id)
	if err2 != nil {
		panic(err2)
	}
}

func ReadName(db *Client) []Name {
	sql_readall := `
	SELECT Id, First, Last FROM names
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.client.Query(sql_readall)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var result []Name
	for rows.Next() {
		item := Name{}
		err2 := rows.Scan(&item.Id, &item.First, &item.Last)
		if err2 != nil {
			panic(err2)
		}
		result = append(result, item)
	}
	return result
}
