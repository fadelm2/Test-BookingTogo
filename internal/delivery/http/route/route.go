package route

import (
	"github.com/bookingtogo/internal/delivery/http"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	App                *mux.Router
	CustomerController *http.CustomerController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {

	c.App.HandleFunc("/customer", c.CustomerController.Create).Methods("POST")
	c.App.HandleFunc("/customer", c.CustomerController.Update).Methods("PUT")
	c.App.HandleFunc("/customer/{id}", c.CustomerController.GetCustomer).Methods("GET")
	c.App.HandleFunc("/customer/{id}", c.CustomerController.Delete).Methods("DELETE")
	c.App.HandleFunc("/health", c.CustomerController.Check).Methods("GET")

}
