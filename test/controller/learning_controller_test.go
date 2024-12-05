package controller

import (
	"encoding/json"
	"errors"
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	"github.com/AugustoKlaic/golearningstack/pkg/api/response"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/router"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	. "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	mock "github.com/AugustoKlaic/golearningstack/test/mock/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type TestSuite struct {
	mockController     *gomock.Controller
	mockService        *mock.MockLearningServiceInterface
	learningController *controller.LearningController
}

func setupTestSuite(t *testing.T) *TestSuite {
	mockController := gomock.NewController(t)
	mockService := mock.NewMockLearningServiceInterface(mockController)
	learningController := controller.NewLearningController(mockService)

	return &TestSuite{
		mockController:     mockController,
		mockService:        mockService,
		learningController: learningController,
	}
}

var firstMessage, secondMessage = entity.MessageEntity{
	Id:       1,
	Content:  "Test message 1",
	DateTime: time.Date(2024, 1, 1, 1, 0, 0, 0, time.UTC),
}, entity.MessageEntity{
	Id:       2,
	Content:  "Test message 1",
	DateTime: time.Date(2024, 2, 2, 2, 0, 0, 0, time.UTC),
}
var allMessages = []entity.MessageEntity{firstMessage, secondMessage}

func TestGetAllMessages(t *testing.T) {
	var suite = setupTestSuite(t)
	gin.SetMode(gin.TestMode)
	router := SetupRouter(suite.learningController)

	t.Run("should get all messages 200 ok", func(t *testing.T) {
		rec := httptest.NewRecorder()
		suite.mockService.EXPECT().GetAllMessages().Return(allMessages, nil).Times(1)
		if req, err := http.NewRequest("GET", "/learning", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)

			var allResponses []response.Message

			jsonDecoder(t, rec.Body.String(), &allResponses)

			assert.Equal(t, allResponses[0].Content, allMessages[0].Content)
			assert.Equal(t, allResponses[0].DateTime, allMessages[0].DateTime)
			assert.Equal(t, allResponses[0].Id, allMessages[0].Id)
			assert.Equal(t, allResponses[1].Content, allMessages[1].Content)
			assert.Equal(t, allResponses[1].DateTime, allMessages[1].DateTime)
			assert.Equal(t, allResponses[1].Id, allMessages[1].Id)
		}
	})

	t.Run("should get empty messages 200 ok", func(t *testing.T) {
		rec := httptest.NewRecorder()
		suite.mockService.EXPECT().GetAllMessages().Return([]entity.MessageEntity{}, nil).Times(1)
		if req, err := http.NewRequest("GET", "/learning", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)

			var allResponses []response.Message

			jsonDecoder(t, rec.Body.String(), &allResponses)
			assert.Equal(t, len(allResponses), 0)
		}
	})

	t.Run("should get error 500 internal server error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		suite.mockService.EXPECT().GetAllMessages().Return(nil, errors.New("")).Times(1)
		if req, err := http.NewRequest("GET", "/learning", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestGetMessage(t *testing.T) {
	var suite = setupTestSuite(t)
	gin.SetMode(gin.TestMode)
	router := SetupRouter(suite.learningController)

	suite.mockService.EXPECT().GetMessage(gomock.AssignableToTypeOf(int(0))).
		DoAndReturn(
			func(id int) (*entity.MessageEntity, error) {
				for _, message := range allMessages {
					if message.Id == id {
						return &message, nil
					}
				}
				return nil, &MessageNotFoundError{Id: id}
			}).Times(3)

	t.Run("should get first message by id 200 ok", func(t *testing.T) {
		rec := httptest.NewRecorder()
		if req, err := http.NewRequest("GET", "/learning/1", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)

			var expected response.Message

			jsonDecoder(t, rec.Body.String(), &expected)

			assert.Equal(t, firstMessage.Content, expected.Content)
			assert.Equal(t, firstMessage.DateTime, expected.DateTime)
			assert.Equal(t, firstMessage.Id, expected.Id)
		}
	})

	t.Run("should get second message by id 200 ok", func(t *testing.T) {
		rec := httptest.NewRecorder()
		if req, err := http.NewRequest("GET", "/learning/2", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)

			var expected response.Message

			jsonDecoder(t, rec.Body.String(), &expected)

			assert.Equal(t, secondMessage.Content, expected.Content)
			assert.Equal(t, secondMessage.DateTime, expected.DateTime)
			assert.Equal(t, secondMessage.Id, expected.Id)
		}
	})

	t.Run("should not find message by id 404 not found", func(t *testing.T) {
		rec := httptest.NewRecorder()
		if req, err := http.NewRequest("GET", "/learning/3", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusNotFound, rec.Code)

			var actualResponse map[string]interface{}
			jsonDecoder(t, rec.Body.String(), &actualResponse)
			expectedMessage := "Message not found. Error: message with Id 3 not found"

			assert.Equal(t, expectedMessage, actualResponse["message"])
		}
	})
}

func jsonDecoder[T any](t *testing.T, body string, target *T) {
	if err := json.Unmarshal([]byte(body), target); err != nil {
		t.Fatalf("Erro ao decodificar JSON: %v", err)
	}
}
