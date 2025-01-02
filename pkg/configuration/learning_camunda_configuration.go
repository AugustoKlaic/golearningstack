package configuration

import (
	"context"
	"github.com/camunda-community-hub/zeebe-client-go/v8/pkg/zbc"
	"log"
	"os"
)

var (
	CamundaClient       zbc.Client
	CamundaProcessId    = "message-updater"
	camundaConfigLogger = log.New(os.Stdout, "CONFIGURATION: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func CreateClient() {
	var err error

	CamundaClient, err = zbc.NewClient(&zbc.ClientConfig{
		GatewayAddress:         Props.Camunda.ZeebeAddress,
		UsePlaintextConnection: true,
	})

	if err != nil {
		camundaConfigLogger.Fatalf("Error creating client: %s", err)
	}

	camundaConfigLogger.Println("Successfully created client")

	deployBpmnProcesses("camundabpmn/resources/message-updater.bpmn")
}

func deployBpmnProcesses(fileNames ...string) {
	ctx := context.Background()
	for _, filename := range fileNames {
		response, err := CamundaClient.NewDeployResourceCommand().AddResourceFile(filename).Send(ctx)
		if err != nil {
			camundaConfigLogger.Fatalf("Error deploying BPMN %s: %s", filename, err)
		}
		camundaConfigLogger.Printf("BPMN deployed: %s", response)
	}
}
