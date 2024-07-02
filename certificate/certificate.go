package certificate

import (
	"crypto/tls"
)

type Certificate struct {
	cert tls.Certificate
}

func (c *Certificate) SetCertificate(cnpj string) (*tls.Certificate, error) {
	res, err := GetSecret(cnpj)

	if err != nil {
		return nil, err
	}

	cert, err := PfxToPem(res)

	if err != nil {
		return nil, err
	}

	c.cert = *cert

	return cert, nil
}
