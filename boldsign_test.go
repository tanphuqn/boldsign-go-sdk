package boldsign

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

var clientID = ""
var secret = ""
var brandId = ""
var documentId = ""

func TestCreateEmbeddedRequestUrl(t *testing.T) {

	var signers []model.DocumentSigner
	signers = append(signers, model.DocumentSigner{Name: "SignerName1", EmailAddress: "tanphuqn@gmail.com"})
	var files []string
	// files = append(files, "./test.pdf", "./download.png")
	files = append(files, "./test.pdf")
	// files = append(files, "./download.png")

	client := Client{ClientID: clientID, Secret: secret}
	request := model.EmbeddedDocumentRequest{
		BrandId:               brandId,
		Title:                 "Sent from API Curl",
		Message:               "This is document message sent from API Curl",
		Signers:               signers,
		Files:                 files,
		EnableSigningOrder:    true,
		ShowToolbar:           true,
		ShowSaveButton:        true,
		ShowNavigationButtons: true,
		ShowSendButton:        true,
		ShowPreviewButton:     true,
	}
	// fmt.Printf("%+v\n", request)
	result, err := client.CreateEmbeddedRequestUrl(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}

func TestGetEmbeddedSignLink(t *testing.T) {
	client := Client{ClientID: clientID, Secret: secret}
	result, err := client.GetEmbeddedSignLink(documentId, "tanphuqn@gmail.com", "")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}

func TestDownloadDocument(t *testing.T) {
	client := Client{ClientID: clientID, Secret: secret}
	response, err := client.DownloadDocument(documentId)
	if err != nil {
		log.Fatal(err)
		return
	}

	fileName := "signed_" + documentId + ".pdf"
	err = ioutil.WriteFile(fileName, response, 0644)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(fileName)
}

func TestGetProperties(t *testing.T) {
	client := Client{ClientID: clientID, Secret: secret}
	response, err := client.GetProperties(documentId)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(response)
}
