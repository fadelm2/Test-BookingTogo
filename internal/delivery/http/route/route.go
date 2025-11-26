package route

import (
	"github.com/bookingtogo/internal/delivery/http"
	"github.com/gorilla/mux"
)

type RouteConfig struct {
	App                     *mux.Router
	CustomerController      *http.CustomerController
	FamilyListController    *http.FamilyListController
	NationalitiesController *http.NationalityController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {

	c.App.HandleFunc("/api/customer", c.CustomerController.Create).Methods("POST")
	c.App.HandleFunc("/api/customer/{id}", c.CustomerController.Update).Methods("PUT")
	c.App.HandleFunc("/api/customer", c.CustomerController.FindAll).Methods("GET")
	c.App.HandleFunc("/api/customer/{id}", c.CustomerController.GetCustomer).Methods("GET")
	c.App.HandleFunc("/api/customer/{id}", c.CustomerController.Delete).Methods("DELETE")
	c.App.HandleFunc("/api/health", c.CustomerController.Check).Methods("GET")
	c.App.HandleFunc("/api/customer/{id}/family", c.FamilyListController.Create).Methods("POST")
	c.App.HandleFunc("/api/customer/{id}/family", c.FamilyListController.GetList).Methods("GET")

	c.App.HandleFunc("/api/nationalities", c.NationalitiesController.GetNationality).Methods("GET")

}
