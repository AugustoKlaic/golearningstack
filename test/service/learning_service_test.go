package service

import (
	"errors"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	error2 "github.com/AugustoKlaic/golearningstack/pkg/domain/error"
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	mock "github.com/AugustoKlaic/golearningstack/test/mock/repository"
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

	t.Run("should get all messages successfully", func(t *testing.T) {
		suite.mockRepo.EXPECT().FindAllMessages().Return(allMessages, nil).Times(1)

		var expected, err = suite.learningService.GetAllMessages()

		assert.Equal(t, nil, err)
		assert.Equal(t, len(expected), 2)
		assert.Equal(t, expected[0].Id, firstMessage.Id)
		assert.Equal(t, expected[1].Id, secondMessage.Id)
	})

	t.Run("should get all messages empty", func(t *testing.T) {
		suite.mockRepo.EXPECT().FindAllMessages().Return([]entity.MessageEntity{}, nil).Times(1)

		var expected, err = suite.learningService.GetAllMessages()

		assert.Equal(t, nil, err)
		assert.Equal(t, len(expected), 0)
	})

	t.Run("should get all messages error", func(t *testing.T) {
		suite.mockRepo.EXPECT().FindAllMessages().Return(nil, errors.New("")).Times(1)

		var expected, err = suite.learningService.GetAllMessages()

		assert.Equal(t, nil, expected)
		assert.Equal(t, "problem retrieving messages: ", err.Error())
	})
}

func TestCreateMessage(t *testing.T) {
	var suite = setupTestSuite(t)

	t.Run("should create a message successfully", func(t *testing.T) {
		suite.mockRepo.EXPECT().CreateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{})).Return(&firstMessage, nil).Times(1)

		var expected, err = suite.learningService.CreateMessage(&firstMessage)

		assert.Equal(t, nil, err)
		assert.Equal(t, expected.Content, firstMessage.Content)
		assert.Equal(t, expected.Id, firstMessage.Id)
	})

	t.Run("should create a message error", func(t *testing.T) {
		suite.mockRepo.EXPECT().CreateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{})).Return(nil, errors.New("")).Times(1)

		var expected, err = suite.learningService.CreateMessage(&firstMessage)

		assert.Equal(t, nil, expected)
		assert.Equal(t, "problem creating message: ", err.Error())
	})
}

func TestGetMessage(t *testing.T) {
	var suite = setupTestSuite(t)

	suite.mockRepo.EXPECT().GetMessage(gomock.AssignableToTypeOf(int(0))).
		DoAndReturn(
			func(id int) (*entity.MessageEntity, error) {
				for _, message := range allMessages {
					if message.Id == id {
						return &message, nil
					}
				}
				return nil, errors.New("")
			}).Times(3)

	t.Run("should get first message successfully", func(t *testing.T) {
		var expectedFirst, errFirst = suite.learningService.GetMessage(1)

		assert.Equal(t, nil, errFirst)
		assert.Equal(t, expectedFirst.Content, firstMessage.Content)
		assert.Equal(t, expectedFirst.Id, firstMessage.Id)
	})

	t.Run("should get second message successfully", func(t *testing.T) {
		var expectedSecond, errSecond = suite.learningService.GetMessage(2)

		assert.Equal(t, nil, errSecond)
		assert.Equal(t, expectedSecond.Content, secondMessage.Content)
		assert.Equal(t, expectedSecond.Id, secondMessage.Id)
	})

	t.Run("should not find message", func(t *testing.T) {
		var expectedNil, err = suite.learningService.GetMessage(0)

		assert.Equal(t, error2.MessageNotFoundError{Id: 0}, err)
		assert.Equal(t, nil, expectedNil)
	})
}

func TestUpdateMessage(t *testing.T) {
	var suite = setupTestSuite(t)
	var capturedMessage *entity.MessageEntity
	var update = &entity.MessageEntity{
		Content:  "updated",
		DateTime: time.Date(2000, 1, 1, 1, 0, 0, 0, time.UTC),
	}

	suite.mockRepo.EXPECT().GetMessage(gomock.AssignableToTypeOf(int(0))).
		DoAndReturn(
			func(id int) (*entity.MessageEntity, error) {
				for _, message := range allMessages {
					if message.Id == id {
						capturedMessage = &message
						return capturedMessage, nil
					}
				}
				return nil, errors.New("")
			}).Times(3)

	suite.mockRepo.EXPECT().UpdateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{})).
		DoAndReturn(func(updatedMessage *entity.MessageEntity) (*entity.MessageEntity, error) {
			assert.Equal(t, capturedMessage, updatedMessage)
			return updatedMessage, nil
		}).Times(1)

	t.Run("should update message successfully", func(t *testing.T) {
		var expectedUpdate, nilErr = suite.learningService.UpdateMessage(update, 1)

		assert.Equal(t, nil, nilErr)
		assert.Equal(t, expectedUpdate.Content, update.Content)
		assert.Equal(t, expectedUpdate.DateTime, update.DateTime)
	})

	t.Run("should not find message to update", func(t *testing.T) {
		var expectedNil, NotFoundErr = suite.learningService.UpdateMessage(update, 0)

		assert.Equal(t, nil, expectedNil)
		assert.Equal(t, NotFoundErr, error2.MessageNotFoundError{Id: 0})
	})

	t.Run("should occur error when updating", func(t *testing.T) {
		suite.mockRepo.EXPECT().UpdateMessage(gomock.AssignableToTypeOf(&entity.MessageEntity{})).
			Return(nil, errors.New("")).Times(1)

		var _, err = suite.learningService.UpdateMessage(update, 1)
		assert.Equal(t, "problem updating message with id: 1. Error: ", err.Error())
	})
}

func TestDeleteMessage(t *testing.T) {
	var suite = setupTestSuite(t)

	t.Run("should delete message successfully", func(t *testing.T) {
		suite.mockRepo.EXPECT().DeleteMessage(gomock.AssignableToTypeOf(int(0))).Return(nil).Times(1)
		var err = suite.learningService.DeleteMessage(1)
		assert.Equal(t, nil, err)
	})

	t.Run("should occur error when deleting", func(t *testing.T) {
		suite.mockRepo.EXPECT().DeleteMessage(gomock.AssignableToTypeOf(int(0))).Return(errors.New("")).Times(1)
		var err = suite.learningService.DeleteMessage(1)
		assert.Equal(t, "problem deleting message with id: 1. Error: ", err.Error())
	})
}
