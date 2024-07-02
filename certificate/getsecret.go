package certificate

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getKeyVaultClient() (client keyvault.BaseClient) {
	keyvaultClient := keyvault.New()
	clientCredentialConfig := auth.NewClientCredentialsConfig(os.Getenv("AZURE_CLIENT_ID"),
		os.Getenv("AZURE_CLIENT_SECRET"), os.Getenv("AZURE_TENANT_ID"))

	clientCredentialConfig.Resource = "https://vault.azure.net"
	authorizer, err := clientCredentialConfig.Authorizer()

	if err != nil {
		fmt.Printf("Error occured while creating azure KV authroizer %v ", err)

	}
	keyvaultClient.Authorizer = authorizer

	return keyvaultClient
}

func GetSecret(cnpj string) (keyvault.SecretBundle, error) {
	keyvaultClient := getKeyVaultClient()
	vaultUri := fmt.Sprintf("https://%s.vault.azure.net", os.Getenv("KEY_VAULT_NAME"))

	res, err := keyvaultClient.GetSecret(context.Background(), vaultUri, cnpj, "")

	if err != nil {
		fmt.Printf("Error occured Get Secret %s , %v", cnpj, err)
		return res, err
	}

	return res, nil
}
