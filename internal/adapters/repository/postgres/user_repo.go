package postgres

import (
	"context"
	"warehouse/internal/core/domain/entity"
	"warehouse/internal/core/ports"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

// UserDB adalah Persistent Model dengan tag GORM
type UserDB struct {
	gorm.Model
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	db.AutoMigrate(&UserDB{}) // Auto-migrate tabel
	return &userRepo{db}
}

func (r *userRepo) Save(ctx context.Context, u *entity.User) error {
	dbModel := UserDB{Username: u.Username, Email: u.Email, Password: u.Password}
	return r.db.WithContext(ctx).Create(&dbModel).Error
}

func (r *userRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	var usersDB []UserDB
	r.db.WithContext(ctx).Find(&usersDB)

	var users []entity.User
	for _, val := range usersDB {
		users = append(users, entity.User{ID: val.ID, Username: val.Username, Email: val.Email})
	}
	return users, nil
}

func (r *userRepo) FindByID(ctx context.Context, id int) (*entity.User, error) {
	var userDB UserDB
	if err := r.db.WithContext(ctx).First(&userDB, id).Error; err != nil {
		return nil, err
	}
	return &entity.User{ID: userDB.ID, Username: userDB.Username, Email: userDB.Email, Password: userDB.Password}, nil
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var userDB UserDB
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&userDB).Error; err != nil {
		return nil, err
	}
	return &entity.User{ID: userDB.ID, Username: userDB.Username, Email: userDB.Email, Password: userDB.Password}, nil
}

func (r *userRepo) Update(ctx context.Context, user *entity.User) error {
	var userDB UserDB
	if err := r.db.WithContext(ctx).First(&userDB, user.ID).Error; err != nil {
		return err
	}
	userDB.Username = user.Username
	userDB.Email = user.Email
	if user.Password != "" {
		userDB.Password = user.Password
	}
	return r.db.WithContext(ctx).Save(&userDB).Error
}

func (r *userRepo) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&UserDB{}, id).Error
}
