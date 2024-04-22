package boldsign_test

import (
	"fmt"
	"log"
	"testing"

	boldsign "github.com/tanphuqn/boldsign-go-sdk/api"
	"github.com/tanphuqn/boldsign-go-sdk/model"
)

var clientID = "baa3163b-93bb-47dc-9e16-1697d090197f"
var secret = "bb72eb4d-9aca-45cf-a16e-0774ec0c046f"
var brandId = "99fe1fa8-3883-4286-a903-297ada4407f5"

// var documentId = ""

// func TestCreateEmbeddedRequestUrl(t *testing.T) {

// 	var signers []model.DocumentSigner
// 	signers = append(signers, model.DocumentSigner{Name: "SignerName1", EmailAddress: "tanphuqn@gmail.com", SignerOrder: 1})
// 	var files []string
// 	// files = append(files, "./test.pdf", "./download.png")
// 	files = append(files, "./test.pdf")
// 	// files = append(files, "./download.png")
// 	// files = append(files, "./Product Road Map.docx")

// 	client := Client{ClientID: clientID, Secret: secret}
// 	request := model.EmbeddedDocumentRequest{
// 		BrandId:               brandId,
// 		Title:                 "Sent from API Curl 3",
// 		Message:               "This is document message sent from API Curl",
// 		Signers:               signers,
// 		Files:                 files,
// 		EnableSigningOrder:    true,
// 		ShowToolbar:           true,
// 		ShowSaveButton:        true,
// 		ShowNavigationButtons: true,
// 		ShowSendButton:        true,
// 		ShowPreviewButton:     true,
// 		DisableExpiryAlert:    true,
// 		DisableEmails:         true,
// 		OnBehalfOf:            "minhthy01011991@gmail.com",
// 	}
// 	// fmt.Printf("%+v\n", request)
// 	result, err := client.CreateEmbeddedRequestUrl(request)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(result)
// }

// func TestCreateSenderIdentities(t *testing.T) {
// 	client := Client{ClientID: clientID, Secret: secret}
// 	request := model.SenderCreateRequest{
// 		Name:  "Yahoo sender",
// 		Email: "minhthy01011991@gmail.com",
// 	}
// 	result, err := client.CreateSenderIdentity(request)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(result)
// }

// func TestUpdateSenderIdentity(t *testing.T) {
// 	client := Client{ClientID: clientID, Secret: secret}
// 	request := model.SenderUpdateRequest{
// 		Name: "Yahoo sender Update",
// 	}
// 	err := client.UpdateSenderIdentity("minhthy01011991@gmail.com", request)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(err)
// }

// func TestDeleteSenderIdentity(t *testing.T) {
// 	client := Client{ClientID: clientID, Secret: secret}
// 	err := client.DeleteSenderIdentity("minhthy01011991@gmail.com")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(err)
// }

// func TestVerifySenderIdentity(t *testing.T) {
// 	client := Client{ClientID: clientID, Secret: secret}
// 	isVerified, err := client.VerifySenderIdentity("minhthy01011991@gmail.com")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(isVerified)
// }

// func TestGetEmbeddedSignLink(t *testing.T) {
// 	client := Client{ClientID: clientID, Secret: secret}
// 	result, err := client.GetEmbeddedSignLink(documentId, "tanphuqn@gmail.com", "")
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(result)
// }

// func TestDownloadDocument(t *testing.T) {
// 	client := Client{ClientID: clientID, Secret: secret}
// 	response, err := client.DownloadDocument(documentId)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	fileName := "signed_" + documentId + ".pdf"
// 	err = ioutil.WriteFile(fileName, response, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	fmt.Println(fileName)
// }

// func TestGetProperties(t *testing.T) {
// 	client := Client{ClientID: clientID, Secret: secret}
// 	response, err := client.GetProperties(documentId)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	fmt.Println(response)
// }

func TestCreateEmbeddedTemplateRequestUrl(t *testing.T) {

	var roles []model.TemplateRole
	roles = append(roles, model.TemplateRole{Name: "Manager", Index: 1})
	var files []string
	// files = append(files, "./test.pdf", "./download.png")
	files = append(files, "./test.pdf")
	// files = append(files, "./download.png")
	// files = append(files, "./Product Road Map.docx")

	client := boldsign.Client{ClientID: clientID, Secret: secret}
	request := model.EmbeddedDocumentRequest{
		BrandId:               brandId,
		Title:                 "API template",
		Description:           "API template description",
		DocumentTitle:         "API document title",
		DocumentMessage:       "API document message description",
		AllowMessageEditing:   true,
		Roles:                 roles,
		Files:                 files,
		ShowToolbar:           true,
		ShowSaveButton:        true,
		ShowSendButton:        true,
		ShowPreviewButton:     true,
		ShowNavigationButtons: true,
		ShowTooltip:           false,

		AllowNewFiles:    true,
		AllowModifyFiles: true,
		ViewOption:       "PreparePage",
	}
	// fmt.Printf("%+v\n", request)
	result, err := client.CreateEmbeddedTemplateRequestUrl(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", result)
}

func TestGetEmbeddedTemplateEditUrl(t *testing.T) {
	//&{TemplateId:2d2115fe-e10d-461a-a71d-9196657afba9
	//CreateUrl:https://app-eu.boldsign.com/document/embed/?templateId=2d2115fe-e10d-461a-a71d-9196657afba9e_oPy5tgDL;1ed042a3-2b4f-425b-9886-2bad99db2775 EditUrl:}
	//EditUrl:https://app-eu.boldsign.com/document/embed/?templateId=2d2115fe-e10d-461a-a71d-9196657afba9e_VQbxDNxT;1ed042a3-2b4f-425b-9886-2bad99db2775&isEditTemplate=true
	client := boldsign.Client{ClientID: clientID, Secret: secret}
	request := model.EmbeddedDocumentRequest{
		ShowToolbar:           true,
		ViewOption:            "PreparePage",
		ShowSaveButton:        true,
		ShowSendButton:        true,
		ShowPreviewButton:     true,
		ShowNavigationButtons: true,
		ShowTooltip:           false,
	}
	templateId := "2d2115fe-e10d-461a-a71d-9196657afba9"
	// fmt.Printf("%+v\n", request)
	result, err := client.GetEmbeddedTemplateEditUrl(templateId, request)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%+v\n", result)
}
