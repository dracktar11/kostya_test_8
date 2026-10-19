package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	client, err := selectClient(db, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	newID, err := insertClient(db, Client{
		FIO:      "Иванов Иван Иванович",
		Login:    "ivanov",
		Birthday: "20000101",
		Email:    "ivanov@mail.ru",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(newID)

	if err = deleteClient(db, newID); err != nil {
		fmt.Println(err)
		return
	}
}

func selectClient(db *sql.DB, id int) (Client, error) {
	cl := Client{}

	row := db.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&cl.ID, &cl.FIO, &cl.Login, &cl.Birthday, &cl.Email)
	if err != nil {
		return cl, err
	}

	return cl, nil
}

func insertClient(db *sql.DB, client Client) (int, error) {
	res, err := db.Exec("INSERT INTO clients (fio, login, birthday, email) VALUES (:fio, :login, :birthday, :email)",
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email))
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func deleteClient(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM clients WHERE id = :id", sql.Named("id", id))

	return err
}
