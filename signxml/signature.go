package signxml

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func SignXml(cnpj, xml, signLocation, action string) (string, error) {
	postBody, _ := json.Marshal(map[string]interface{}{
		"cnpj":         cnpj,
		"xml":          xml,
		"signTag":      signLocation,
		"signLocation": signLocation,
		"action":       action,
		"reference":    true,
	})

	body := bytes.NewBuffer(postBody)

	resp, err := http.Post(os.Getenv("SIGNATURE_BASE_URL")+"/reinf", "application/json", body)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	//Read the response body
	responseBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(responseBody), err
}
