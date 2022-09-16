package handler

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
}
type Service interface {
	FindAll() ([]User, error)
}
type repository struct {
	db *gorm.DB
}

type service struct {
	repo Repository
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func NewService(repository Repository) *service {
	return &service{repository}
}
func (r *repository) FindAll() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error

	return user, err
}
func (s *service) FindAll() ([]User, error) {
	users, err := s.repo.FindAll()

	return users, err
}
