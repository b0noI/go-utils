package gcp

import "testing"

// Not yet a proper test, just to check if this is working
func TestAccessSecretVersion(t *testing.T) {
	_, err := AccessSecretVersion("projects/991052578721/secrets/backend-cli-credentials/versions/latest")
	if err != nil {
		t.Errorf("MyType.AppendToMessage did not update the message correctly. Error: %v", err)
	}
}
