package mapper

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/response"
	"github.com/AugustoKlaic/golearningstack/pkg/domain"
)

func ToMessageEntity(request response.Message) *domain.MessageEntity {
	return &domain.MessageEntity{
		Content:  request.Content,
		DateTime: request.DateTime,
	}
}

func ToMessageResponse(entities ...domain.MessageEntity) []response.Message {
	responses := make([]response.Message, len(entities))
	for i, entity := range entities {
		responses[i] = response.Message{
			Id:       entity.Id,
			Content:  entity.Content,
			DateTime: entity.DateTime,
		}
	}
	return responses
}
