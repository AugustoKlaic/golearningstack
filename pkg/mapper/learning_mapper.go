package mapper

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/message/request"
	"github.com/AugustoKlaic/golearningstack/pkg/api/message/response"
	. "github.com/AugustoKlaic/golearningstack/pkg/api/security/request"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
)

func ToMessageEntity(request request.MessageRequest) *entity.MessageEntity {
	return &entity.MessageEntity{
		Content:  request.Content,
		DateTime: request.DateTime,
	}
}

func ToMessageResponses(entities ...entity.MessageEntity) []response.Message {
	responses := make([]response.Message, len(entities))
	for i, entity := range entities {
		responses[i] = ToMessageResponse(&entity)
	}
	return responses
}

func ToMessageResponse(entity *entity.MessageEntity) response.Message {
	return response.Message{
		Id:       entity.Id,
		Content:  entity.Content,
		DateTime: entity.DateTime,
	}
}

func ToUserCredentialsEntity(userCredentials *LoginRequest) *entity.UserCredentials {
	return &entity.UserCredentials{
		Username: userCredentials.UserName,
		Password: userCredentials.Password,
	}
}
