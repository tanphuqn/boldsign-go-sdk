package boldsign

import (
	"fmt"
	"log"
	"testing"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

func TestCreateEmbeddedRequestUrl(t *testing.T) {

	var signers []model.DocumentSigner
	signers = append(signers, model.DocumentSigner{Name: "SignerName1", EmailAddress: "tanphuqn@gmail.com", SignerOrder: 1})
	// signers = append(signers, model.DocumentSigner{Name: "Signer Name 2", SignerOrder: 2, EmailAddress: "tanphuqn+2@gmail.com"})

	var files []string
	files = append(files, "./test.pdf")

	clientID := "93faa0be-9338-4dff-86d4-e993a9747b8e"
	secret := "31dc6637-8a6f-4e2a-9827-579a8d0d87a4"

	client := Client{ClientID: clientID, Secret: secret}
	request := model.EmbeddedDocumentRequest{
		Title:              "Sent from API Curl",
		Message:            "This is document message sent from API Curl",
		EnableSigningOrder: true,
		Signers:            signers,
		ShowToolbar:        true,
		RedirectUrl:        "https://boldsign.dev/sign/redirect",

		Files: files,
	}
	fmt.Printf("%+v\n", request)
	result, err := client.CreateEmbeddedRequestUrl(request)
	if err != nil {
		println("dest")
		log.Fatal(err)
		return
	}
	// https://scriptable.com/blog/how-to-create-a-go-package-golang
	fmt.Println(result)

}
