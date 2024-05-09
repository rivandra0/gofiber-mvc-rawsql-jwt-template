package queries

import (
	"errors"

	"github.com/rivandra0/fiber-mvc-template/src/database"
	"github.com/rivandra0/fiber-mvc-template/src/models"
)

type UserQuery struct{}

func (*UserQuery) GetOne(email string, password string) (models.User, error) {
	db := database.GetDatabase()

	user := models.User{}
	err := db.QueryRow("CALL SP_AppUser_GetOne(?,?)", email, password).Scan(&user.Id,
		&user.Email,
		&user.UserName,
		&user.GroupId,
		&user.Tier,
		&user.IsEmailConfirmed,
		&user.IsActive,
		&user.CreatedBy,
		&user.CreatedAt,
		&user.UpdatedBy,
		&user.UpdatedAt)
	if err != nil {
		return user, errors.New("wrong login credentials")
	}

	return user, nil
}
