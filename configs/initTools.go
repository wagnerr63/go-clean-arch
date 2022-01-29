package configs

import (
	"go-clean-arch/utils/httpRouter"
	"go-clean-arch/utils/token"
)

func InitTools() (httpRouter.Router, token.TokenHash) {
	router := httpRouter.NewMuxRouter()
	tokenHasher := token.NewJWT()

	return router, tokenHasher
}
