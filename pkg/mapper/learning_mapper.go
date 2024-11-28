package mapper

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/request"
	"github.com/AugustoKlaic/golearningstack/pkg/api/response"
	"github.com/AugustoKlaic/golearningstack/pkg/domain"
)

func ToMessageEntity(request request.MessageRequest) *domain.MessageEntity {
	return &domain.MessageEntity{
		Content:  request.Content,
		DateTime: request.DateTime,
	}
}

func ToMessageResponses(entities ...domain.MessageEntity) []response.Message {
	responses := make([]response.Message, len(entities))
	for i, entity := range entities {
		responses[i] = ToMessageResponse(&entity)
	}
	return responses
}

func ToMessageResponse(entity *domain.MessageEntity) response.Message {
	return response.Message{
		Id:       entity.Id,
		Content:  entity.Content,
		DateTime: entity.DateTime,
	}
}
