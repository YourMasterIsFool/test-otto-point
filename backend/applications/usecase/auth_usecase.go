package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	Authdto "backend/dtos/Auth"
	authdto "backend/dtos/Auth"
	"backend/pkg/jwt"
	"backend/pkg/password"
	"backend/pkg/response"
	"fmt"
)

type AuthUsecase interface {
	SignIn(dtos Authdto.LoginDto) (*string, error)
	Save(dtos authdto.CreateAuthDto) (*string, error)
}

type useAuthUsecase struct {
	AuthRepo repository.AuthRepository
}

func NewAuthUsecase(repo repository.AuthRepository) *useAuthUsecase {
	return &useAuthUsecase{
		AuthRepo: repo,
	}
}

func (u *useAuthUsecase) Save(dto authdto.CreateAuthDto) (*string, error) {

	var schema = entity.UserEntity{
		Username: dto.Username,
		Password: password.GeneratePassword(dto.Password),
		Name:     dto.Name,
	}
	data, err := u.AuthRepo.Save(schema)
	if err != nil {
		return nil, err
	}

	// create token
	token, err := jwt.GenerateToken(data.ID)
	if err != nil {
		fmt.Println(err)
		return nil, response.NewErrorResponse(400, "gagal generate token", nil)
	}

	return &token, nil
}

func (u *useAuthUsecase) SignIn(dtos Authdto.LoginDto) (*string, error) {
	user, err := u.AuthRepo.FindByUsername(dtos.Username)
	if err != nil {
		return nil, response.NewErrorResponse(400, "username tidak ada", nil)
	}

	if !password.ComparePassword(user.Password, dtos.Password) {
		return nil, response.NewErrorResponse(400, "password salah", nil)
	}

	// create token

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {

		return nil, response.NewErrorResponse(400, "gagal generate token", nil)
	}

	return &token, nil

}
