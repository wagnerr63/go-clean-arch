package configs

import "go-clean-arch/utils/httpRouter"

func InitTools() (httpRouter.Router) {
	router:= httpRouter.NewMuxRouter()
	return router
}