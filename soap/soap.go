package soap

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetSefazConfigSync(xml string) (envelop string) {
	// var rxml = strings.ReplaceAll(xml, "&", "&amp;")
	// rxml = strings.ReplaceAll(rxml, "<", "&lt;")
	// rxml = strings.ReplaceAll(rxml, ">", "&gt;")
	// rxml = strings.ReplaceAll(rxml, "'", "&apos;")
	// rxml = strings.ReplaceAll(rxml, `"`, "&quot;")

	envelop = `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
    xmlns:sped="http://sped.fazenda.gov.br/">
    <soapenv:Header />
    <soapenv:Body>
        <sped:ReceberLoteEventos>
            <sped:loteEventos>
                <Reinf xmlns="http://www.reinf.esocial.gov.br/schemas/envioLoteEventos/v1_05_01">
                    <loteEventos>` + xml + `</loteEventos>
                </Reinf>
            </sped:loteEventos>
        </sped:ReceberLoteEventos>
    </soapenv:Body>
</soapenv:Envelope>`
	return envelop
}

func SoapCallSync(cert tls.Certificate, baseURL, xml, soapAction string) string {

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
	req.Header.Add("Content-Type", "text/xml")

	if soapAction != "" {
		req.Header.Add("SOAPAction", soapAction)
	}

	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	s := strings.TrimSpace(string(content))

	return s
}

func GetSefazConfig(xml, tpInsc, nrInsc string) (envelop string) {
	// var rxml = strings.ReplaceAll(xml, "&", "&amp;")
	// rxml = strings.ReplaceAll(rxml, "<", "&lt;")
	// rxml = strings.ReplaceAll(rxml, ">", "&gt;")
	// rxml = strings.ReplaceAll(rxml, "'", "&apos;")
	// rxml = strings.ReplaceAll(rxml, `"`, "&quot;")

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

	fmt.Println("xml envelop ==================")
	fmt.Println(envelop)
	fmt.Println("==================")

	return envelop
}

func SoapCall(cert tls.Certificate, baseURL, xml string) string {

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
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("==================")
	fmt.Println(response.StatusCode)
	fmt.Println("==================")

	content, _ := io.ReadAll(response.Body)

	s := strings.TrimSpace(string(content))

	return s
}
