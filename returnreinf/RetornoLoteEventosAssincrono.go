package returnreinf

import (
	"encoding/json"
	"encoding/xml"
	"errors"
)

type Reinf struct {
	XMLName                      xml.Name `xml:"Reinf"`
	Text                         string   `xml:",chardata"`
	Xmlns                        string   `xml:"xmlns,attr"`
	RetornoLoteEventosAssincrono struct {
		Text            string `xml:",chardata"`
		IdeContribuinte struct {
			Text   string `xml:",chardata"`
			TpInsc string `xml:"tpInsc"`
			NrInsc string `xml:"nrInsc"`
		} `xml:"ideContribuinte"`
		Status struct {
			Text         string `xml:",chardata"`
			CdResposta   string `xml:"cdResposta"`
			DescResposta string `xml:"descResposta"`
			Ocorrencias  struct {
				Text       string `xml:",chardata"`
				Ocorrencia []struct {
					Text        string `xml:",chardata"`
					Codigo      string `xml:"codigo"`
					Descricao   string `xml:"descricao"`
					Tipo        string `xml:"tipo"`
					Localizacao string `xml:"localizacao"`
				} `xml:"ocorrencia"`
			} `xml:"ocorrencias"`
		} `xml:"status"`
		DadosRecepcaoLote struct {
			Text                     string `xml:",chardata"`
			DhRecepcao               string `xml:"dhRecepcao"`
			VersaoAplicativoRecepcao string `xml:"versaoAplicativoRecepcao"`
			ProtocoloEnvio           string `xml:"protocoloEnvio"`
		} `xml:"dadosRecepcaoLote"`
		DadosProcessamentoLote struct {
			Text                              string `xml:",chardata"`
			VersaoAplicativoProcessamentoLote string `xml:"versaoAplicativoProcessamentoLote"`
		} `xml:"dadosProcessamentoLote"`
		RetornoEventos struct {
			Text   string `xml:",chardata"`
			Evento []struct {
				Text          string `xml:",chardata"`
				ID            string `xml:"Id,attr"`
				EvtDupl       string `xml:"evtDupl,attr"`
				RetornoEvento struct {
					Text       string `xml:",chardata"`
					AnyElement string `xml:"any_element"`
				} `xml:"retornoEvento"`
			} `xml:"evento"`
		} `xml:"retornoEventos"`
	} `xml:"retornoLoteEventosAssincrono"`
}

type StatusReturn struct {
	Status   string `json:"status"`
	Protocol string `json:"protocol"`
}

func LoadXML(xml string) (*Reinf, error) {
	var returnReinf Reinf

	if err := json.Unmarshal([]byte(xml), &returnReinf); err != nil {
		return nil, errors.New("não foi possível converter o xml RetornoLoteEventosAssincrono para struct")
	}

	return &returnReinf, nil
}

func (r *Reinf) GetStatus() StatusReturn {
	return StatusReturn{
		Status:   r.RetornoLoteEventosAssincrono.Status.DescResposta,
		Protocol: r.RetornoLoteEventosAssincrono.DadosRecepcaoLote.ProtocoloEnvio,
	}
}
