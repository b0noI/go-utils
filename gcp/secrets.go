package gcp

import (
	"context"
	errors "errors"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

// AccessSecretVersion retrieves the payload data for a secret version from
// the Google Cloud Secret Manager API.
//
// Parameters:
//   secretName - the name of the secret version to retrieve.
//	 Example: "projects/991052578721/secrets/backend-cli-credentials/versions/latest"
//
// Returns:
//   the payload data for the secret version, or an error if the secret
//   could not be accessed
func AccessSecretVersion(secretName string) ([]byte, error) {
	// Create the client.
	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to create secretmanager client: %v", err))
	}

	// Build the request.
	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	// Call the API.
	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to access secret version: %v", err))
	}

	return result.Payload.Data, nil
}
