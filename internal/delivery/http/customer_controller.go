package http

import (
	"encoding/json"
	"github.com/bookingtogo/internal/helper"
	"github.com/bookingtogo/internal/model"
	"github.com/bookingtogo/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CustomerController struct {
	UseCase *usecase.CustomerUseCase
	Log     *logrus.Logger
}

func NewCustomerController(useCase *usecase.CustomerUseCase, log *logrus.Logger) *CustomerController {
	return &CustomerController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *CustomerController) Create(w http.ResponseWriter, r *http.Request) {
	request := new(model.CreateCustomerRequest)

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		helper.WriteJSON(w, helper.NewBadRequest("Invalid JSON format", err))
		return
	}

	result, err := c.UseCase.Create(r.Context(), request)
	if err != nil {
		helper.WriteJSON(w, err)
		return
	}

	helper.WriteJSON(w, result)
}

func (c *CustomerController) Update(w http.ResponseWriter, r *http.Request) {
	request := new(model.UpdateCustomerRequest)

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		helper.WriteJSON(w, helper.NewBadRequest("Invalid JSON format", err))
		return
	}

	result, err := c.UseCase.Update(r.Context(), request)
	if err != nil {
		helper.WriteJSON(w, err)
		return
	}

	helper.WriteJSON(w, result)
}

func (h *CustomerController) GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := new(model.GetCustomerRequest)

	request.ID, _ = helper.StringToInt(params["id"])

	c, err := h.UseCase.Get(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	helper.WriteJSON(w, c)
}

func (c *CustomerController) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := new(model.DeleteCustomerRequest)

	request.ID, _ = helper.StringToInt(params["id"])

	if err := c.UseCase.Delete(r.Context(), request); err != nil {
		helper.WriteJSON(w, err)
		return
	}

	helper.WriteJSON(w, map[string]string{
		"message": "Customer deleted successfully",
	})
}

func (c *CustomerController) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	c.Log.Info("gateway is running")

	w.Write([]byte(`{"status":"ok"}`))
}

func (h *CustomerController) FindAll(w http.ResponseWriter, r *http.Request) {
	request := new(model.AllCustomerRequest)

	h.Log.Info("Request All get Customers")
	result, err := h.UseCase.FindAll(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	helper.WriteJSON(w, result)
}
