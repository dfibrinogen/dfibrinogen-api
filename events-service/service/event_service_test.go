package service

import (
	"github.com/dfibrinogen/dfibrinogen-api/events-service/model"
	"github.com/dfibrinogen/dfibrinogen-api/events-service/repository"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEventService_GetDataAll(t *testing.T) {

	mockRepo := repository.InitMockEventRepo()

	mockRepo.On("FetchEventAll").
		Return([]*model.Event{{"EVENT.01", "Test 01", "Location 01", time.Now(), time.Now(), nil}}, nil).
		Once()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/events")

	h := &eventService{repo: mockRepo}

	// Assertions
	if assert.NoError(t, h.GetDataAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
