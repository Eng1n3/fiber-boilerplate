package main

import (
	"fiber-boilerplate/pkg/entities"
	"fmt"

	gs "github.com/randree/gormseeder"
	"gorm.io/gorm"
)

func init() {
	gs.Seed(gs.State{

		Tag: "insert_roles",

		Perform: func(db *gorm.DB) error {
			roles := &[]entities.Role{
				{
					Name:        "user",
					Description: "role user"},
				{
					Name:        "admin",
					Description: "role admin"},
			}

			fmt.Println("Inserting roles:", roles)
			// Use Generics API inside the transaction
			if err := db.Create(roles).Error; err != nil {
				// return any error will rollback
				return err
			}

			return nil
		},
	})
}
