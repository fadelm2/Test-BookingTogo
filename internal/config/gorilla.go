package config

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func NewGorillaRouter(config *viper.Viper) *mux.Router {
	r := mux.NewRouter()

	return r
	
}
