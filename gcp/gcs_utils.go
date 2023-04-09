package gcp

import (
	"context"
	"fmt"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func ReadFile(bucketName, objectName string) (string, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return "", fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	object := bucket.Object(objectName)

	r, err := object.NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("object.NewReader: %v", err)
	}
	defer r.Close()

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	return string(data), nil
}
