package sendxml

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
)

func GetSefazConfig(xml, tpInsc, nrInsc string) (envelop string) {
	envelop = `<?xml version="1.0" encoding="utf-8"?>
<Reinf xmlns="http://www.reinf.esocial.gov.br/schemas/envioLoteEventosAssincrono/v1_00_00">
  <envioLoteEventos>
    <ideContribuinte>
      <tpInsc>` + tpInsc + `</tpInsc>
      <nrInsc>` + nrInsc + `</nrInsc>
    </ideContribuinte>
    <eventos>` + xml + `</eventos>
  </envioLoteEventos>
</Reinf>`

	return envelop
}

func SendXML(cert tls.Certificate, baseURL, xml string) (xmlReturn string, statusCode int, err error) {

	body := string(xml)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Renegotiation:      tls.RenegotiateOnceAsClient,
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{cert},
			},
		},
	}

	req, _ := http.NewRequest("POST", baseURL, bytes.NewBufferString(body))
	req.Header.Add("Content-Type", "application/xml; charset=utf-8")

	response, err := client.Do(req)

	if err != nil {
		return "", 0, err
	}

	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	statusCode = response.StatusCode

	if err != nil {
		return "", 0, err
	}

	xmlReturn = strings.TrimSpace(string(content))

	return xmlReturn, statusCode, nil
}
