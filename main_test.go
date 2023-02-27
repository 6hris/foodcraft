package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"foodcraft/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestListRecipesHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/recipes", ts.URL))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode) // check that the status code is 200
	data, _ := ioutil.ReadAll(resp.Body)

	var recipes []models.Recipe
	json.Unmarshal(data, &recipes)
	assert.Equal(t, len(recipes), 24) // check that there are 24 recipes in the database
}

// not working - need authorization + index out of range when no auth problem
func TestNewRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	recipe := models.Recipe{
		ID: 123456,
	}

	raw, _ := json.Marshal(recipe)
	resp, err := http.Post(fmt.Sprintf("%s/recipes", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode) // check that the status code is 200
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "")
}

func TestSignInHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/signin", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode) // check that the status code is 200
}

func TestSignInHandlerFail(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "abcd",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/signin", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode) // check that the status code is 401 (invalid user/password)
}

func TestSignUpHandlerFail(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/signup", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode) // check that the status code is 500 (user exists already)
}

func TestSignUpHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin10",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/signup", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode) // check that the status code is 200
	data, _ := ioutil.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "Account has been created")
}

func TestRefreshHandlerUnauthorized(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/refresh", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode) // check that the status code is 401 (no sign in yet)
}

// not working - need authorization beforehand
func TestRefreshHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	user := models.User{
		Username: "admin",
		Password: "password",
	}

	raw, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("%s/refresh", ts.URL), "application/json", bytes.NewBuffer(raw))
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusAccepted, resp.StatusCode) // check that the status code is 200
}
