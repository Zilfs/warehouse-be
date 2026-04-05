package usecase

import (
	"context"
	"warehouse/internal/core/domain/entity"
	"warehouse/internal/core/ports"
)

type userUsecase struct {
	repo ports.UserRepository
}

// NewUserUsecase adalah constructor untuk menginisialisasi usecase.
func NewUserUsecase(r ports.UserRepository) ports.UserUsecase {
	return &userUsecase{
		repo: r,
	}
}

// CreateUser mengimplementasikan logika pembuatan user baru.
func (u *userUsecase) CreateUser(ctx context.Context, user *entity.User) error {
	return u.repo.Save(ctx, user)
}

// GetUserByID mengambil satu user berdasarkan ID uniknya.
func (u *userUsecase) GetUserByID(ctx context.Context, id int) (*entity.User, error) {
	return u.repo.FindByID(ctx, id)
}

// GetAllUsers mengambil seluruh daftar user yang tersedia.
func (u *userUsecase) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return u.repo.FindAll(ctx)
}

// UpdateUser memperbarui data user yang sudah ada.
func (u *userUsecase) UpdateUser(ctx context.Context, user *entity.User) error {
	return u.repo.Update(ctx, user)
}

// DeleteUser menghapus user dari sistem (Metode yang sebelumnya missing).
func (u *userUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.repo.Delete(ctx, id)
}
