package service

import (
	"errors"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	mock "github.com/AugustoKlaic/golearningstack/test/mock"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

/*
 - In unit testing when I need to mock I am using mockgen
 - setup to use mockgen:
 - run: go get -u github.com/golang/mock/mockgen
 - run:go install github.com/golang/mock/mockgen@latest
 - add this to the file that the mock will be generated:
		//go:generate mockgen -source=service.go -destination=mock/service.go
 - run: go generate -v ./... to generate mocks
*/

type TestSuite struct {
	mockController  *gomock.Controller
	mockRepo        *mock.MockLearningRepositoryInterface
	learningService *service.LearningService
}

func setupTestSuite(t *testing.T) *TestSuite {
	mockController := gomock.NewController(t)
	mockRepo := mock.NewMockLearningRepositoryInterface(mockController)
	learningService := service.NewLearningService(mockRepo)

	return &TestSuite{
		mockController:  mockController,
		mockRepo:        mockRepo,
		learningService: learningService,
	}
}

var firstMessage, secondMessage = entity.MessageEntity{
	Id:       1,
	Content:  "Test message 1",
	DateTime: time.Now(),
}, entity.MessageEntity{
	Id:       2,
	Content:  "Test message 1",
	DateTime: time.Now(),
}

var allMessages = []entity.MessageEntity{firstMessage, secondMessage}

func TestGetAllMessages(t *testing.T) {
	var suite = setupTestSuite(t)

	suite.mockRepo.EXPECT().FindAllMessages().Return(allMessages, nil).Times(1)

	var expected, err = suite.learningService.GetAllMessages()

	assert.Equal(t, nil, err)
	assert.Equal(t, len(expected), 2)
	assert.Equal(t, expected[0].Id, firstMessage.Id)
	assert.Equal(t, expected[1].Id, secondMessage.Id)
}

func TestGetAllMessagesEmpty(t *testing.T) {
	var suite = setupTestSuite(t)

	suite.mockRepo.EXPECT().FindAllMessages().Return([]entity.MessageEntity{}, nil).Times(1)

	var expected, err = suite.learningService.GetAllMessages()

	assert.Equal(t, nil, err)
	assert.Equal(t, len(expected), 0)
}

func TestGetAllMessagesError(t *testing.T) {
	var suite = setupTestSuite(t)

	suite.mockRepo.EXPECT().FindAllMessages().Return(nil, errors.New("")).Times(1)

	var expected, err = suite.learningService.GetAllMessages()

	assert.Equal(t, nil, expected)
	assert.Equal(t, "problem retrieving messages: ", err.Error())
}

func TestCreateMessage(t *testing.T) {
	var suite = setupTestSuite(t)

	suite.mockRepo.EXPECT().CreateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{})).Return(&firstMessage, nil).Times(1)
	var expected, err = suite.learningService.CreateMessage(&firstMessage)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected.Content, firstMessage.Content)
	assert.Equal(t, expected.Id, firstMessage.Id)
}
