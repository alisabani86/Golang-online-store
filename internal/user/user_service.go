package user

import (
	"context"

	"server/internal/middleware"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type service struct {
	Repository
	timout time.Duration
	middleware.Middleware
}

func NewService(repository Repository, middleware middleware.Middleware) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
		middleware,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, s.timout)

	defer cancel()
	hashedPassword, err := util.HashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)

	if err != nil {
		return nil, err
	}

	res := &CreateUserResponse{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) Login(ctx context.Context, req *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timout)

	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	// Replace with actual user data
	token, err := s.Middleware.GenerateToken(middleware.User{ID: int(u.ID), Username: u.Username})

	return &LoginUserResponse{
			AccesToken: token,
			Username:   u.Username,
			ID:         strconv.Itoa(int(u.ID))},
		nil
}
