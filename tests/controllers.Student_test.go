package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/richardsang2008/MathTestsGenerator/controllers"
	"github.com/richardsang2008/MathTestsGenerator/models/response"
	"github.com/stretchr/testify/assert"

	"net/http"

	"testing"
)

func TestGetStudentByStudentId(t *testing.T) {
	//Build expected body
	body := gin.H{
		"Student":   "Joy",
		"StudnetId": 1,
	}
	// Get a  router
	routes := controllers.Routes{}
	r := routes.InitializeRoutes()
	// Create a request to send to the above route
	w := performRequest(r, "GET", "/api/Student/byStudentId?studnetId=1")
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	response := response.Student{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.Equal(t, body["StudnetId"], response.StudnetId)
}
