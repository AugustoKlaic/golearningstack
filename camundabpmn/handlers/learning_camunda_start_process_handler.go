package handlers

import (
	"context"
	"github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/zbc"
	"log"
	"os"
)

var camundaMessageCreateHandlerLogger = log.New(os.Stdout, "CAMUNDA_HANDLER: ", log.Ldate|log.Ltime|log.Lshortfile)

type CamundaMessageCreatedHandler struct {
	camundaClient zbc.Client
}

func NewCamundaMessageCreateHandler() *CamundaMessageCreatedHandler {
	return &CamundaMessageCreatedHandler{
		camundaClient: configuration.CamundaClient,
	}
}

func (handler *CamundaMessageCreatedHandler) CreateMessageHandler(message *entity.MessageEntity) {
	ctx := context.Background()

	request, err := handler.
		camundaClient.
		NewCreateInstanceCommand().
		BPMNProcessId(configuration.CamundaProcessId).
		LatestVersion().
		VariablesFromObject(message)

	if err != nil {
		camundaMessageCreateHandlerLogger.Fatalf("error creating instance: %s", err)
	}

	response, _ := request.WithResult().Send(ctx)

	camundaMessageCreateHandlerLogger.Printf("Create camunda instanceId: %v", response)
}
