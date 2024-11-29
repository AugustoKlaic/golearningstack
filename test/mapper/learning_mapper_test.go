package mapper

import (
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"github.com/AugustoKlaic/golearningstack/pkg/mapper"
	"testing"
	"time"
)

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
