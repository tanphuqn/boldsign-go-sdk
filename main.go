package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	url := "https://api-eu.boldsign.com/v1/document/createEmbeddedRequestUrl"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("Title", "Sent from API Curl 2222")
	_ = writer.WriteField("Signers[0][Name]", "Signer Name 1")
	_ = writer.WriteField("Signers[0][EmailAddress]", "tanphuqn@gmail.com")
	file, _ := os.Open("test.pdf")
	defer file.Close()
	part8, errFile8 := writer.CreateFormFile("Files", filepath.Base("test.pdf"))
	if errFile8 != nil {
		fmt.Println(errFile8)
		return
	}
	_, errFile8 = io.Copy(part8, file)
	if errFile8 != nil {
		fmt.Println(errFile8)
		return
	}
	_ = writer.WriteField("Signers[1][Name]", "Signer Name 2")
	_ = writer.WriteField("Signers[1][EmailAddress]", "tanphuqn+12@gmail.com")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IkEwRTNCNkYwQUEyRjFEQTZDODFGQ0U2RDE1MEQ5QjM4MjkxNjQ1Q0NSUzI1NiIsInR5cCI6ImF0K2p3dCIsIng1dCI6Im9PTzI4S292SGFiSUg4NXRGUTJiT0NrV1JjdyJ9.eyJuYmYiOjE2NjI0ODI3MjUsImV4cCI6MTY2MjQ4NjMyNSwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50LWV1LmJvbGRzaWduLmNvbSIsImF1ZCI6Imh0dHBzOi8vYWNjb3VudC1ldS5ib2xkc2lnbi5jb20vcmVzb3VyY2VzIiwiY2xpZW50X2lkIjoiOTNmYWEwYmUtOTMzOC00ZGZmLTg2ZDQtZTk5M2E5NzQ3YjhlIiwibmFtZSI6Ik1ldmx1ZGluIER6aWhpYyIsIkVtYWlsSWQiOiJtZEBsYXd0ZWNoMzY1LmNvLnVrIiwicHJlZmVycmVkX3VzZXJuYW1lIjoibWRAbGF3dGVjaDM2NS5jby51ayIsImlkIjoiMjg3MGVhZTAtN2E3NS00YzExLWI5ZTUtYTFlZDNmNTY1ZWZiIiwib3JnYW5pemF0aW9uaWQiOiIxZWQwNDJhMy0yYjRmLTQyNWItOTg4Ni0yYmFkOTlkYjI3NzUiLCJvcmdhbml6YXRpb25uYW1lIjoiTGF3dGVjaCAzNjUgTGltaXRlZCIsImZpcnN0TmFtZSI6Ik1ldmx1ZGluIiwibGFzdE5hbWUiOiJEemloaWMiLCJ0ZWFtaWQiOiJiOGFkM2I3MC0zYzkyLTQxZjEtYWI0NS1lZWNjNjU1OGE2NmYiLCJyb2xlaWQiOiIxIiwic3ViIjoiMjg3MGVhZTAtN2E3NS00YzExLWI5ZTUtYTFlZDNmNTY1ZWZiIiwidGltZXpvbmUiOiJFdXJvcGUvTG9uZG9uIiwic2Vzc2lvblRpbWVvdXQiOiIzMCIsInRlbmFudHN0YXR1cyI6IkFjdGl2ZSIsInN1YnNjcmlwdGlvbmlkIjoiYTc4ZTczMDYtYzdiYS00MjBjLTg3OTQtMWMxOTYwMDEzODhlIiwiZW52aXJvbm1lbnQiOiJUZXN0IiwianRpIjoiOUE2REU1MDYzNTkxRjBCRDdENkFFRDA4RTRGQUZFODkiLCJpYXQiOjE2NjI0ODI3MjUsInNjb3BlIjpbIkJvbGRTaWduLkRvY3VtZW50cy5BbGwiLCJCb2xkU2lnbi5UZW1wbGF0ZXMuQWxsIl19.Fowz6namff9O5J0EOxe03NzbPFKOcpqx2XPhFHEdonccPAAnc01RMBwG7x5BBGWt_6kr7Pz_wGX4l59bqLLor3nNNssjyG_1kjyeMhTj_K0YXzWoTEtGItMB88bDpNqebvnQ7SmKqkmgiydH_xMPYwxSMABoCOOrJZUrOp11amO5UocCxp9Scxs64gDGuUdU373qbicqNcMnhnY1kUaIS4b6wjGXaE7AAAVM-4jxtExWNluNREwSdog3vniSkS3BcSzlW2sPtDv42XCTVfZggk2Ilhs94SHGI1Z-bF1-qSAIY29Gq5RbgekuiTQaTWNfWYkxTNSh9xh8O7xLy7zJvw")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.StatusCode)
	fmt.Println(string(body))
}
