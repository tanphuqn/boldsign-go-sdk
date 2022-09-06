# BoldSign Go SDK
A Go wrapper for the BoldSign API.

The unofficial library for using the BoldSign API for golang.

https://www.boldsign.com/help/api/general/preparing-your-application/

## Usage

### Get the boldsign module

Note that you need to include the **v** in the version tag.

```
$ go get github.com/tanphuqn/boldsign@v0.1.0
```


### Client

```go
client := boldsign.Client{ClientID: "CLIENT ID", Secret: "SECRET"}
```

### Embedded Signature Request

```go
	request := model.EmbeddedDocumentRequest{
		Title:              "Sent from API Curl",
		ShowToolbar:        true,
		RedirectUrl:        "https://boldsign.dev/sign/redirect",
		Message:            "Message",
		EnableSigningOrder: false,
		Signers:            signers,
		Files:              files,
		SendViewOption:     "FillingPage",
	}
	fmt.Printf("%+v\n", request)
	result, err := client.CreateEmbeddedRequestUrl(request)
	if err != nil {
		log.Fatal(err)
	}
    // type EmbeddedSendCreated
    fmt.Println(result.GetSendUrl())
```

## Testing

```
$ go test
```

## Tagging

```
$ git tag v0.1.0
$ git push origin --tags
```
