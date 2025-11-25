package test

import (
	"article/internal/model"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	requestBody := model.CreatePostRequest{
		Title:    "Cara Membuat REST API dengan Fiber dan GORM di Golang",
		Content:  " 'REST API menjadi dasar untuk komunikasi antar aplikasi modern. Fiber adalah framework Go yang cepat dan ringan, sementara GORM memudahkan kita berinteraksi dengan database. Pada artikel ini, kita membuat API lengkap mulai dari routing, service hingga repository.",
		Status:   "publish",
		Category: "programming",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/article", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.PostResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Title, responseBody.Data.Title)
	assert.Equal(t, requestBody.Content, responseBody.Data.Content)
	assert.NotNil(t, responseBody.Data.CreatedDate)
	assert.NotNil(t, responseBody.Data.UpdatedDate)
}
