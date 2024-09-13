package db

import (
	"fmt"
	"log"
)

func (dbClient *DBClient) DeleteUser(id int) error {
	_, err := dbClient.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return fmt.Errorf("error deleting user: %w", err)
	}
	_, err = dbClient.AivenDB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		log.Println("Error deleting user:", err)
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}
