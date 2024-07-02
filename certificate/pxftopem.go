package certificate

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"golang.org/x/crypto/pkcs12"
)

type AzureKeyVaultCert struct {
	// key represents the private key of the certificate
	Key []byte
	// cert represents the server certificate
	Cert []byte
}

func PfxToPem(pfx keyvault.SecretBundle) (*tls.Certificate, error) {

	pfxBytes, err := base64.StdEncoding.DecodeString(*pfx.Value)

	if err != nil {
		return nil, err
	}

	pemBlocks, err := pkcs12.ToPEM(pfxBytes, "")

	if err != nil {
		return nil, err
	}

	certs := &AzureKeyVaultCert{}

	for i, v := range pemBlocks {
		if strings.Contains(v.Type, "KEY") {
			var keyPEM bytes.Buffer
			err = pem.Encode(&keyPEM, pemBlocks[i])
			if err != nil {
				return nil, fmt.Errorf("error encoding key pem block: %v", err)
			}
			certs.Key = keyPEM.Bytes()
		}

		if strings.Contains(v.Type, "CERTIFICATE") {
			var certPEM bytes.Buffer
			err = pem.Encode(&certPEM, pemBlocks[1])
			if err != nil {
				return nil, fmt.Errorf("error encoding certificate pem block: %v", err)
			}

			if certs.Cert == nil {
				certs.Cert = certPEM.Bytes()
			} else {
				certs.Cert = append(certs.Cert, certPEM.Bytes()...)
			}
		}

	}

	cert, err := tls.X509KeyPair(certs.Cert, certs.Key)

	if err != nil {
		return nil, fmt.Errorf("error creating X509 Key Pair: %v", err)
	}

	return &cert, nil
}
