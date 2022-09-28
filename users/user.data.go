package user

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/fontainecoutino/CelebsDontCare/database"
)

func getUser(userID int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := database.DB.QueryRowContext(ctx,
		`SELECT * FROM users WHERE id = $1`, userID)

	user := &User{}
	err := row.Scan(&user.UserID, &user.Name)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func getUserList() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DB.QueryContext(ctx, `SELECT * FROM users `)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer results.Close()
	users := make([]User, 0)
	for results.Next() {
		var user User
		results.Scan(
			&user.UserID,
			&user.Name)

		users = append(users, user)
	}
	return users, nil
}

func insertUser(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`INSERT INTO users (name) VALUES ($1)`, user.Name)

	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func removeUser(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := database.DB.ExecContext(ctx,
		`DELETE FROM users where id = $1`, productID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
