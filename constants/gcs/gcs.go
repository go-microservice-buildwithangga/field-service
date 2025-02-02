package constants

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

type ServiceAccountKeyJSON struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

type GCSClient struct {
	ServiceAccountKeyJSON ServiceAccountKeyJSON
	BucketName            string
}

type IGCSClient interface {
	UploadFileToGCS(context.Context, string, []byte) (string, error)
}

func NewGCSClient(serviceAccountKeyJSON ServiceAccountKeyJSON, bucketName string) IGCSClient {
	return &GCSClient{
		ServiceAccountKeyJSON: serviceAccountKeyJSON,
		BucketName:            bucketName,
	}
}

func (g *GCSClient) createClient(ctx context.Context) (*storage.Client, error) {
	reqBodyBytes := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBytes).Encode(g.ServiceAccountKeyJSON)
	if err != nil {
		logrus.Errorf("Failed to create client : %v", err)
		return nil, err
	}

	jsonByte := reqBodyBytes.Bytes()
	client, err := storage.NewClient(ctx, option.WithCredentialsJSON(jsonByte))
	if err != nil {
		logrus.Errorf("Failed to create client : %v", err)
		return nil, err
	}
	return client, nil

}

func (g *GCSClient) UploadFileToGCS(ctx context.Context, fileName string, data []byte) (url string, err error) {
	const (
		contentType   = "application/octet-stream"
		timeInSeconds = 60
	)

	// Membuat client GCS
	client, err := g.createClient(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to create GCS client: %w", err)
	}
	defer client.Close()

	// Mengatur timeout
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeInSeconds)*time.Second)
	defer cancel()

	// Upload file ke GCS
	if err := g.uploadFile(ctx, client, fileName, data); err != nil {
		return "", err
	}

	// Update metadata file (content type)
	if err := g.updateMetadata(ctx, client, fileName, contentType); err != nil {
		return "", err
	}

	// Mengembalikan URL file yang telah diunggah
	url = fmt.Sprintf("https://storage.googleapis.com/%s/%s", g.BucketName, fileName)
	return url, nil
}

// Fungsi untuk mengunggah file ke GCS
func (g *GCSClient) uploadFile(ctx context.Context, client *storage.Client, fileName string, data []byte) error {
	obj := client.Bucket(g.BucketName).Object(fileName)
	writer := obj.NewWriter(ctx)
	writer.ChunkSize = 0

	defer writer.Close() // Pastikan writer ditutup setelah digunakan

	if _, err := io.Copy(writer, bytes.NewBuffer(data)); err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}
	return nil
}

// Fungsi untuk memperbarui metadata objek
func (g *GCSClient) updateMetadata(ctx context.Context, client *storage.Client, fileName, contentType string) error {
	obj := client.Bucket(g.BucketName).Object(fileName)
	_, err := obj.Update(ctx, storage.ObjectAttrsToUpdate{ContentType: contentType})
	if err != nil {
		return fmt.Errorf("failed to update metadata: %w", err)
	}
	return nil
}
