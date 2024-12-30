package handlers

import (
	"context"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/entities"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/worker"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/zbc"
	"log"
	"os"
)

var camundaMessageServiceHandlerLogger = log.New(os.Stdout, "CAMUNDA_HANDLER: ", log.Ldate|log.Ltime|log.Lshortfile)

type CamundaMessageServiceHandler struct {
	camundaClient zbc.Client
}

func NewCamundaMessageServiceHandler() *CamundaMessageServiceHandler {
	return &CamundaMessageServiceHandler{}
}

func (handler *CamundaMessageServiceHandler) HandleMessageChange(client worker.JobClient, job entities.Job) {
	var message entity.MessageEntity
	_ = job.GetVariablesAs(message)

	message.Content = message.Content + " Message edited by camunda service task!"

	ctx := context.Background()
	response, _ := client.NewCompleteJobCommand().JobKey(job.Key).VariablesFromObject(message)
	_, _ = response.Send(ctx)

	camundaMessageServiceHandlerLogger.Println("Camunda service task ended")
}
