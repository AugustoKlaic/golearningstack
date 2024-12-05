package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/AugustoKlaic/golearningstack/pkg/api/controller"
	"github.com/AugustoKlaic/golearningstack/pkg/api/request"
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

var messageRequest = request.MessageRequest{
	Content:  "Message Request",
	DateTime: time.Date(2024, 3, 3, 3, 0, 0, 0, time.UTC),
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

func TestCreateMessage(t *testing.T) {
	var suite = setupTestSuite(t)
	gin.SetMode(gin.TestMode)
	router := SetupRouter(suite.learningController)

	t.Run("should create a message successfully", func(t *testing.T) {
		suite.mockService.EXPECT().CreateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{})).
			DoAndReturn(func(mappedRequest *entity.MessageEntity) (*entity.MessageEntity, error) {
				return &entity.MessageEntity{
					Id:       1,
					Content:  mappedRequest.Content,
					DateTime: mappedRequest.DateTime,
				}, nil
			}).Times(1)

		rec := httptest.NewRecorder()
		body := jsonEncoder(t, messageRequest)

		if req, err := http.NewRequest("POST", "/learning", bytes.NewReader(body)); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusCreated, rec.Code)

			var expected response.Message
			jsonDecoder(t, rec.Body.String(), &expected)

			assert.Equal(t, messageRequest.Content, expected.Content)
			assert.Equal(t, messageRequest.DateTime, expected.DateTime)
			assert.Equal(t, 1, expected.Id)
		}
	})

	t.Run("should occur error mapping request body 400 bad request", func(t *testing.T) {
		rec := httptest.NewRecorder()
		if req, err := http.NewRequest("POST", "/learning", bytes.NewReader([]byte("{}"))); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			var actualResponse map[string]interface{}
			jsonDecoder(t, rec.Body.String(), &actualResponse)
			expectedMessage := "invalid request JSON. Error: Key: 'MessageRequest.Content' Error:Field validation for 'Content' failed on the 'required' tag\n" +
				"Key: 'MessageRequest.DateTime' Error:Field validation for 'DateTime' failed on the 'required' tag"

			assert.Equal(t, expectedMessage, actualResponse["message"])
		}
	})

	t.Run("should occur error creating message 500 internal server error", func(t *testing.T) {
		suite.mockService.EXPECT().CreateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{})).
			DoAndReturn(func(mappedRequest *entity.MessageEntity) (*entity.MessageEntity, error) {
				return nil, errors.New("problem creating message")
			}).Times(1)

		rec := httptest.NewRecorder()
		body := jsonEncoder(t, messageRequest)
		if req, err := http.NewRequest("POST", "/learning", bytes.NewReader(body)); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			var actualResponse map[string]interface{}
			jsonDecoder(t, rec.Body.String(), &actualResponse)
			assert.Equal(t, "problem creating message", actualResponse["message"])
		}
	})
}

func TestDeleteMessage(t *testing.T) {
	suite := setupTestSuite(t)
	gin.SetMode(gin.TestMode)
	router := SetupRouter(suite.learningController)

	t.Run("should delete a message successfully 204 NoContent", func(t *testing.T) {
		suite.mockService.EXPECT().DeleteMessage(gomock.AssignableToTypeOf(int(0))).Return(nil).Times(1)
		rec := httptest.NewRecorder()
		if req, err := http.NewRequest("DELETE", "/learning/1", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusNoContent, rec.Code)
		}
	})

	t.Run("should occur error when deleting message 500 internal server error", func(t *testing.T) {
		suite.mockService.EXPECT().DeleteMessage(gomock.AssignableToTypeOf(int(0))).
			Return(errors.New("problem deleting message")).Times(1)
		rec := httptest.NewRecorder()
		if req, err := http.NewRequest("DELETE", "/learning/1", nil); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			var actualResponse map[string]interface{}
			jsonDecoder(t, rec.Body.String(), &actualResponse)
			expectedMessage := "problem deleting message"
			assert.Equal(t, expectedMessage, actualResponse["message"])
		}
	})
}

func TestUpdateMessage(t *testing.T) {
	var suite = setupTestSuite(t)
	gin.SetMode(gin.TestMode)
	router := SetupRouter(suite.learningController)

	suite.mockService.EXPECT().UpdateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{}),
		gomock.AssignableToTypeOf(int(0))).
		DoAndReturn(func(updatedMessage *entity.MessageEntity, id int) (*entity.MessageEntity, error) {
			for _, message := range allMessages {
				if message.Id == id {
					message.Content = messageRequest.Content
					message.DateTime = messageRequest.DateTime
					return &message, nil
				}
			}
			return nil, &MessageNotFoundError{Id: id}
		}).Times(2)

	t.Run("should update a message successfully 200 ok", func(t *testing.T) {
		rec := httptest.NewRecorder()
		body := jsonEncoder(t, messageRequest)
		if req, err := http.NewRequest("PATCH", "/learning/1", bytes.NewReader(body)); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusOK, rec.Code)

			var expected response.Message
			jsonDecoder(t, rec.Body.String(), &expected)

			assert.Equal(t, messageRequest.Content, expected.Content)
			assert.Equal(t, messageRequest.DateTime, expected.DateTime)
			assert.Equal(t, 1, expected.Id)
		}
	})

	t.Run("should occur error mapping request body 400 bad request", func(t *testing.T) {
		rec := httptest.NewRecorder()
		if req, err := http.NewRequest("PATCH", "/learning/1", bytes.NewReader([]byte("{}"))); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			var actualResponse map[string]interface{}
			jsonDecoder(t, rec.Body.String(), &actualResponse)
			expectedMessage := "invalid request JSON. Error: Key: 'MessageRequest.Content' Error:Field validation for 'Content' failed on the 'required' tag\n" +
				"Key: 'MessageRequest.DateTime' Error:Field validation for 'DateTime' failed on the 'required' tag"

			assert.Equal(t, expectedMessage, actualResponse["message"])
		}
	})

	t.Run("should not found message to update 404 not found", func(t *testing.T) {
		rec := httptest.NewRecorder()
		body := jsonEncoder(t, messageRequest)
		if req, err := http.NewRequest("PATCH", "/learning/3", bytes.NewReader(body)); err != nil {
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

	t.Run("should occur error updating message 500 internal server error", func(t *testing.T) {
		rec := httptest.NewRecorder()
		suite.mockService.EXPECT().UpdateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{}), gomock.AssignableToTypeOf(int(0))).
			Return(nil, errors.New("problem updating message")).
			Times(1)

		body := jsonEncoder(t, messageRequest)
		if req, err := http.NewRequest("PATCH", "/learning/1", bytes.NewReader(body)); err != nil {
			t.Fatalf("Erro ao criar requisição: %v", err)
		} else {
			router.ServeHTTP(rec, req)
			assert.Equal(t, http.StatusInternalServerError, rec.Code)

			var actualResponse map[string]interface{}
			jsonDecoder(t, rec.Body.String(), &actualResponse)
			assert.Equal(t, "problem updating message", actualResponse["message"])
		}
	})
}

func jsonEncoder[T any](t *testing.T, body T) []byte {
	encodedJson, err := json.Marshal(body)

	if err != nil {
		t.Fatalf("Erro ao codificar JSON: %v", err)
	}

	return encodedJson
}

func jsonDecoder[T any](t *testing.T, body string, target *T) {
	if err := json.Unmarshal([]byte(body), target); err != nil {
		t.Fatalf("Erro ao decodificar JSON: %v", err)
	}
}
