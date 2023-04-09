package gcp

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	bucketName := "cocktails-ai"
	objectName := "gpt4_prompt.txt"

	content, err := ReadFile(bucketName, objectName)
	if err != nil {
		t.Fatalf("Error reading file from GCS: %v", err)
	}

	fmt.Printf("Content of GCS file:\n%s\n", content)
}
