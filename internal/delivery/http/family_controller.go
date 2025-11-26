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

type FamilyListController struct {
	UseCase *usecase.FamilyListUseCase
	Log     *logrus.Logger
}

func NewFamilyListController(useCase *usecase.FamilyListUseCase, log *logrus.Logger) *FamilyListController {
	return &FamilyListController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *FamilyListController) Create(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := new(model.CreateFamilyListRequest)

	// set cst_id
	request.CustomerID, _ = helper.StringToInt(params["id"])

	// decode body
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		helper.WriteJSON(w, helper.NewBadRequest("Invalid JSON format", err))
		return
	}

	// log setelah request terisi
	jsonData, _ := json.MarshalIndent(request, "", "  ")
	c.Log.Info("Request Create Family: " + string(jsonData))

	// execute usecase
	result, err := c.UseCase.Create(r.Context(), request)
	if err != nil {
		helper.WriteJSON(w, err)
		return
	}

	helper.WriteJSON(w, result)
}

func (h *FamilyListController) GetFamilyList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := new(model.GetFamilyListRequest)

	request.ID, _ = params["id"]

	result, err := h.UseCase.FindAll(r.Context(), *request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	helper.WriteJSON(w, result)
}
