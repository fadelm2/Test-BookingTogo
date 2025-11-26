package http

import (
	"github.com/bookingtogo/internal/helper"
	"github.com/bookingtogo/internal/model"
	"github.com/bookingtogo/internal/usecase"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type NationalityController struct {
	UseCase *usecase.NationalityUseCase
	Log     *logrus.Logger
}

func NewNationalityController(useCase *usecase.NationalityUseCase, log *logrus.Logger) *NationalityController {
	return &NationalityController{
		UseCase: useCase,
		Log:     log,
	}
}

func (h *NationalityController) GetNationality(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	request := new(model.GetNationalityRequest)

	request.ID, _ = params["id"]

	h.Log.Info("Request All get Nationality")
	result, err := h.UseCase.FindAll(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	helper.WriteJSON(w, result)
}
