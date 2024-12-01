package mapper

import (
	"github.com/AugustoKlaic/golearningstack/pkg/api/request"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"github.com/AugustoKlaic/golearningstack/pkg/mapper"
	"testing"
	"time"
)

func TestMapToEntity(t *testing.T) {
	var messageRequest = request.MessageRequest{
		Content:  "Hello World!",
		DateTime: time.Date(2024, 12, 1, 18, 0, 0, 0, time.UTC),
	}

	var want = mapper.ToMessageEntity(messageRequest)

	if want.Content != messageRequest.Content {
		t.Errorf("expected Content: %v, got: %v", messageRequest.Content, want.Content)
	}
	if want.DateTime != messageRequest.DateTime {
		t.Errorf("expected DateTime: %v, got: %v", messageRequest.DateTime, want.DateTime)
	}
}

func TestMapToResponse(t *testing.T) {
	var messageEntity = entity.MessageEntity{
		Content:  "Hello World!",
		DateTime: time.Date(2024, 12, 1, 18, 0, 0, 0, time.UTC),
	}

	var want = mapper.ToMessageResponse(&messageEntity)

	if want.Content != messageEntity.Content {
		t.Errorf("expected Content: %v, got: %v", messageEntity.Content, want.Content)
	}
	if want.DateTime != messageEntity.DateTime {
		t.Errorf("expected DateTime: %v, got: %v", messageEntity.DateTime, want.DateTime)
	}
}

func TestMapToResponses(t *testing.T) {
	var messageEntities = []entity.MessageEntity{
		{
			Content:  "Hello World!",
			DateTime: time.Date(2024, 12, 1, 18, 0, 0, 0, time.UTC),
		},
		{
			Content:  "Hello World 2!",
			DateTime: time.Date(2024, 10, 12, 4, 0, 0, 0, time.UTC),
		},
	}

	var want = mapper.ToMessageResponses(messageEntities...)

	for i := 0; i < len(messageEntities); i++ {
		if want[i].Content != messageEntities[i].Content {
			t.Errorf("expected Content: %v, got: %v", messageEntities[i].Content, want[i].Content)
		}
		if want[i].DateTime != messageEntities[i].DateTime {
			t.Errorf("expected DateTime: %v, got: %v", messageEntities[i].DateTime, want[i].DateTime)
		}
	}
}
