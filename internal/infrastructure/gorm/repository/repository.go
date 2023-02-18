package repository

import (
	"gorm.io/gorm"
)

type UserGORM struct {
	gorm.Model
}

type repository struct {
	DB *gorm.DB
}

//// Set tablename (GORM)
//func (UserGORM) TableName() string {
//	return "users"
//}
//
//func (u UserGORM) toEntitiesUser() *entities.User {
//	return &entities.User{
//		ID:        u.ID,
//		Username:  u.Username,
//		Password:  u.Password,
//		CreatedAt: u.CreatedAt,
//	}
//}
//
//func NewUserGORM(entityUser *entities.User) *UserGORM {
//	u := UserGORM{}
//	u.ID = entityUser.ID
//	u.Username = entityUser.Username
//	u.Password = entityUser.Password
//	u.CreatedAt = entityUser.CreatedAt
//	return &u
//}
//

//
//func NewUserGORMRepository(db *gorm.DB) user.UserRepository {
//	return &repository{
//		DB: db,
//	}
//}
//
//func (r *repository) Create(entityUser *entities.User) error {
//	u := NewUserGORM(entityUser)
//
//	err := r.DB.Create(&u).Error
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (r *repository) List() ([]*entities.User, error) {
//	var users []UserGORM
//
//	err := r.DB.Find(&users).Error
//	if err != nil {
//		return nil, err
//	}
//
//	// TODO: Refactor. maybe inefficient.
//	result := make([]*entities.User, 0, len(users))
//	for _, user := range users {
//		result = append(result, user.toEntitiesUser())
//	}
//
//	return result, nil
//}
//
//func (r *repository) GetByUsername(username string) (*entities.User, error) {
//	var user UserGORM
//
//	r.DB.Find(&user, "username = ?", username)
//
//	// If no such user present return an error
//	if id.UUIDIsNil(user.ID) {
//		return nil, errors.New("User does not exists.")
//	}
//
//	return user.toEntitiesUser(), nil
//}
//
//func (r *repository) Get(userID string) (*entities.User, error) {
//	var user UserGORM
//
//	r.DB.Find(&user, "id = ?", userID)
//
//	// If no such user present return an error
//	if id.UUIDIsNil(user.ID) {
//		return nil, errors.New("User does not exists.")
//	}
//
//	return user.toEntitiesUser(), nil
//}
