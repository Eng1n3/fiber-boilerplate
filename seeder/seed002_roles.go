package main

import (
	"fiber-boilerplate/pkg/entities"
	"fiber-boilerplate/pkg/enums"
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
					Name:        string(enums.User),
					Description: "role user"},
				{
					Name:        string(enums.Admin),
					Description: "role admin"},
				{
					Name:        string(enums.SuperUser),
					Description: "role superuser"},
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
