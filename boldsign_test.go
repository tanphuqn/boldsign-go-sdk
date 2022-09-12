package boldsign

import (
	"fmt"
	"log"
	"testing"

	"github.com/tanphuqn/boldsign-go-sdk/model"
)

func TestCreateEmbeddedRequestUrl(t *testing.T) {

	var signers []model.DocumentSigner
	signers = append(signers, model.DocumentSigner{Name: "SignerName1", EmailAddress: "tanphuqn@gmail.com"})
	// signers = append(signers, model.DocumentSigner{Name: "Signer Name 2", SignerOrder: 2, EmailAddress: "tanphuqn+2@gmail.com"})
	// reminderSettings := model.ReminderSettings{
	// 	ReminderDays:       1,
	// 	ReminderCount:      1,
	// 	EnableAutoReminder: false,
	// }
	var files []string
	files = append(files, "./test.pdf", "./download.png")
	// files = append(files, "./test.pdf")
	// files = append(files, "./download.png")
	clientID := "93faa0be-9338-4dff-86d4-e993a9747b8e"
	secret := "7b426406-b20a-4342-9122-8952f1a0e9ce"

	client := Client{ClientID: clientID, Secret: secret}
	request := model.EmbeddedDocumentRequest{
		Title:              "Sent from API Curl",
		Message:            "This is document message sent from API Curl",
		RedirectUrl:        "https://boldsign.dev/sign/redirect",
		Signers:            signers,
		Files:              files,
		EnableSigningOrder: false,
		ShowToolbar:        true,
		// DisableExpiryAlert:    false,
		// SendViewOption:        "FillingPage",
		// ReminderSettings:      reminderSettings,
		// BrandId:               "",
		// EnableReassign:        false,
		// ExpiryDays: 180,
		// EnablePrintAndSign:    false,
		// ShowSaveButton:        false,
		// OnBehalfOf:            "",
		// UseTextTags:           false,
		// SendLinkValidTill:     "",
		// ShowNavigationButtons: false,
		// ShowSendButton:        false,
		// HideDocumentId: false,
		// EnableEmbeddedSigning: false,
		// ShowPreviewButton:     false,
		// DisableEmails:         false,
	}
	// fmt.Printf("%+v\n", request)
	result, err := client.CreateEmbeddedRequestUrl(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(result)
}
