package service

import (
	"context"
	"online_store/internal/middleware"
	"online_store/internal/presentation"
	"online_store/util"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

func (s *service) CreateUser(ctx context.Context, req *presentation.CreateUserRequest) (*presentation.CreateUserResponse, error) {

	ctx, cancel := context.WithTimeout(ctx, s.timout)

	defer cancel()
	hashedPassword, err := util.HashedPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &presentation.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)

	if err != nil {
		return nil, err
	}

	res := &presentation.CreateUserResponse{
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

func (s *service) Login(ctx context.Context, req *presentation.LoginUserRequest) (*presentation.LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timout)

	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &presentation.LoginUserResponse{}, err
	}

	// Replace with actual user data
	token, err := s.Middleware.GenerateToken(middleware.User{ID: int(u.ID), Username: u.Username})

	return &presentation.LoginUserResponse{
			AccesToken: token,
			Username:   u.Username,
			ID:         strconv.Itoa(int(u.ID))},
		nil
}
