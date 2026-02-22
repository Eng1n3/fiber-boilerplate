package main

import (
	app "fiber-boilerplate"
	"fiber-boilerplate/pkg/entities"
	"fmt"

	gs "github.com/randree/gormseeder"
	"gorm.io/gorm"
)

func init() {
	gs.Seed(gs.State{

		Tag: "insert_permissions",

		Perform: func(db *gorm.DB) error {
			routes := app.App().GetRoutes()

			permissions := []entities.Permission{}
			for _, route := range routes {
				permission := entities.Permission{
					Name:        route.Name,
					Description: fmt.Sprintf("Permission for %s %s", route.Method, route.Path),
					Path:        route.Path,
					Method:      route.Method,
				}
				if route.Name != "" {
					permissions = append(permissions, permission)
				}
			}
			// dataR, _ := json.MarshalIndent(permissions, "", "  ")
			// fmt.Println(string(dataR), 333444)
			// panic("err")
			if err := db.Create(permissions).Error; err != nil {
				// return any error will rollback
				return err
			}

			return nil
		},
	})
}
