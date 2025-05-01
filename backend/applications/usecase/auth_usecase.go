package usecase

import (
	"backend/domain/entity"
	"backend/domain/repository"
	authdto "backend/dtos/auth"
	"backend/pkg/jwt"
	"backend/pkg/password"
	"backend/pkg/response"
)

// authusecase interface
type AuthUsecase interface {
	SignIn(dtos authdto.LoginDto) (*string, error)
	Save(dtos authdto.CreateAuthDto) (*string, error)
}

// definig struct for auth usecase
type useAuthUsecase struct {
	AuthRepo repository.AuthRepository
}

// authusecase handler
func NewAuthUsecase(repo repository.AuthRepository) *useAuthUsecase {
	return &useAuthUsecase{
		AuthRepo: repo,
	}
}

func (u *useAuthUsecase) Save(dto authdto.CreateAuthDto) (*string, error) {

	// schema users for register
	var schema = entity.UserEntity{
		Username: dto.Username,
		Password: password.GeneratePassword(dto.Password),
		Name:     dto.Name,
	}

	//register user
	data, err := u.AuthRepo.Save(schema)
	if err != nil {
		return nil, err
	}

	// create token
	token, err := jwt.GenerateToken(data.ID)
	if err != nil {

		// failed created user
		return nil, response.NewErrorResponse(400, "gagal generate token", nil)
	}

	return &token, nil
}

func (u *useAuthUsecase) SignIn(dtos authdto.LoginDto) (*string, error) {
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
