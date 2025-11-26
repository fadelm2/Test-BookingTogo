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
	App      *mux.Router // ✔ benar
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	customerRepository := repository.NewCustomerRepository(config.Log)
	familyListRepository := repository.NewFamilyListRepository(config.Log)

	// setup Usecase
	customerUseCase := usecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, customerRepository)
	familyListUseCase := usecase.NewFamilyListUseCase(config.DB, config.Log, config.Validate, familyListRepository, customerRepository)

	// setup Controlle

	customerController := http.NewCustomerController(customerUseCase, config.Log)
	familyListController := http.NewFamilyListController(familyListUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:                  config.App,
		CustomerController:   customerController,
		FamilyListController: familyListController,
	}
	routeConfig.Setup()

}
