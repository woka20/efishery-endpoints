package tokenization

import "efishery-endpoints/auth-app/repository/user"

type TokenLogiCInterface interface {
}

type TokenLogic struct {
	UserRepo user.UserRepoInterface
}

func NewTokenLogic() TokenLogiCInterface {
	return &TokenLogic{
		UserRepo: user.NewRepository(),
	}
}
