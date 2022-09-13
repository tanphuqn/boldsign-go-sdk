# BoldSign Go SDK
A Go wrapper for the BoldSign API.

The unofficial library for using the BoldSign API for golang.

https://www.boldsign.com/help/api/general/preparing-your-application/

## Usage

### Get the boldsign module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/tanphuqn/boldsign-go-sdk
```


### Client

```go
client := boldsign.Client{ClientID: "CLIENT ID", Secret: "SECRET"}
```

### Embedded Signature Request

```go
	var signers []model.DocumentSigner
	signers = append(signers, model.DocumentSigner{Name: "SignerName1", EmailAddress: "tanphuqn@gmail.com"})
	var files []string
	files = append(files, "./test.pdf")

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
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v1.0.0
$ git push origin --tags
```
