package service

import (
	"github.com/AugustoKlaic/golearningstack/pkg/service"
	mock "github.com/AugustoKlaic/golearningstack/test/mock"
	"github.com/golang/mock/gomock"
	"testing"
)

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

func TestGetAllMessages(t *testing.T) {
	// setup test here
}

// link for study: https://medium.com/@ninucium/mocking-integration-tests-in-go-nina-pakshina-d01eefe5251d
/*
setup para usar o mockgen:
 - run: go get -u github.com/golang/mock/mockgen
 - run:go install github.com/golang/mock/mockgen@latest
 - add this to the file that the mock will be generated:
		//go:generate mockgen -source=service.go -destination=mock/service.go
 - run: go generate -v ./... to generate mocks
*/
