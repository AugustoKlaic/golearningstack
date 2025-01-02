package camundabpmn

import (
	"github.com/AugustoKlaic/golearningstack/camundabpmn/handlers"
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
)

type CamundaAdmin struct {
	messageCreateHandler         *handlers.CamundaMessageCreatedHandler
	camundaMessageServiceHandler *handlers.CamundaMessageServiceHandler
}

func NewCamundaAdmin() *CamundaAdmin {
	return &CamundaAdmin{
		messageCreateHandler:         handlers.NewCamundaMessageCreateHandler(),
		camundaMessageServiceHandler: handlers.NewCamundaMessageServiceHandler(),
	}
}

func (admin *CamundaAdmin) ExecuteProcess(message *entity.MessageEntity) {

	go configuration.CamundaClient.NewJobWorker().
		JobType("message-updater").
		Handler(admin.camundaMessageServiceHandler.HandleMessageChange).
		Open()

	admin.messageCreateHandler.CreateMessageHandler(message)
}

// https://github.com/jwulf/camunda-cloud-getting-started-go/blob/2d600ce751ec634f82391b1ffe83d6ce31e637c5/main.go#L26
// https://github.com/camunda/camunda-platform-get-started/blob/main/go/main.go#L154
// https://docs.camunda.io/docs/apis-tools/community-clients/go-client/job-worker/
