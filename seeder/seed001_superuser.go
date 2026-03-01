package main

// import (
// 	"fiber-boilerplate/models"
// 	"os"

// 	"golang.org/x/crypto/bcrypt"

// 	"gorm.io/gorm"
// )

// func (s *Seeder) SeedUser(tx *gorm.DB) error {
// 	password, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SUPERUSER_PASSWORD")), 10) // Hash the password
// 	if err != nil {
// 		panic(err)
// 	}
// 	user := &models.User{
// 		Name:     "superuser",
// 		Email:    os.Getenv("SUPERUSER_EMAIL"),
// 		Password: string(password),
// 	}

// 	seed := &models.Seed{
// 		Name: "UserSeeder",
// 	}
// 	// Use Generics API inside the transaction
// 	if err := tx.Create(user).Error; err != nil {
// 		// return any error will rollback
// 		return err
// 	}

// 	if err := tx.Create(seed).Error; err != nil {
// 		// return any error will rollback
// 		return err
// 	}

// 	// return nil will commit the whole transaction
// 	return nil
// }
import (
	"fiber-boilerplate/pkg/entities"
	"fmt"
	"os"

	gs "github.com/randree/gormseeder"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func init() {
	gs.Seed(gs.State{

		Tag: "insert_superuser",

		Perform: func(db *gorm.DB) error {
			password, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SUPERUSER_PASSWORD")), bcrypt.DefaultCost) // Hash the password
			if err != nil {
				panic(err)
			}
			user := &entities.User{
				Username: "superuser",
				Email:    os.Getenv("SUPERUSER_EMAIL"),
				Password: string(password),
			}

			fmt.Println("Inserting superuser with email:", user.Email)
			// Use Generics API inside the transaction
			if err := db.Create(user).Error; err != nil {
				// return any error will rollback
				return err
			}

			return nil
		},
	})
}
