package main

import (
	"errors"
	"fmt"
	"os"
)

func createDB(db_path string) error {
	// Create new db file if it doesn't exists
	if _, err := os.Stat(db_path); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(db_path)
		if err != nil {
			fmt.Printf("Error creating new database. %s. Error: %s", db_path, err)
			return err
		}
		defer file.Close()
	}

	return nil
}
