package user

import (
	"context"
	"local-notice-hex-go/internal/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Create(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error)
	GetByID(ctx context.Context, id int64) (*user.UserResponse, error)
	GetByUsername(ctx context.Context, username string) (*user.UserResponse, error)
	GetAll(ctx context.Context) ([]*user.UserResponse, error)
	Update(ctx context.Context, id int64, req *user.UserRequest) (*user.UserResponse, error)
	Delete(ctx context.Context, id int64) error
}

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) Service {
	return &UserService{repo: repo}
}

func (s *UserService) Create(ctx context.Context, req *user.UserRequest) (*user.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	status := true
	if req.Status != nil {
		status = *req.Status
	}

	u := &user.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Status:   status,
	}

	created, err := s.repo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	return s.toResponse(created), nil
}

func (s *UserService) GetByID(ctx context.Context, id int64) (*user.UserResponse, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.toResponse(u), nil
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (*user.UserResponse, error) {
	u, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return s.toResponse(u), nil
}

func (s *UserService) GetAll(ctx context.Context) ([]*user.UserResponse, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*user.UserResponse, len(users))
	for i, u := range users {
		responses[i] = s.toResponse(u)
	}

	return responses, nil
}

func (s *UserService) Update(ctx context.Context, id int64, req *user.UserRequest) (*user.UserResponse, error) {
	existing, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Username != "" {
		existing.Username = req.Username
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		existing.Password = string(hashedPassword)
	}

	if req.Status != nil {
		existing.Status = *req.Status
	}

	updated, err := s.repo.Update(ctx, id, existing)
	if err != nil {
		return nil, err
	}

	return s.toResponse(updated), nil
}

func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserService) toResponse(u *user.User) *user.UserResponse {
	return &user.UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
