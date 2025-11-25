package config

import (
	"github.com/bookingtogo/internal/delivery/http"
	"github.com/bookingtogo/internal/delivery/http/route"
	"github.com/bookingtogo/internal/repository"
	"github.com/bookingtogo/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *mux.Route
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	postRepository := repository.NewPostsRepository(config.Log)

	// setup Usecase
	postUseCase := usecase.NewPostUseCase(config.DB, config.Log, config.Validate, postRepository)

	// setup Controlle

	postController := http.NewPostController(postUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		PostController: postController,
	}
	routeConfig.Setup()

}
