package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// Get client from account nickname.
func GetClientFromName(name string, config *oauth2.Config) (*http.Client, error) {
	tokFile := name + ".json"
	fptr, err := os.Open(tokFile)
	if err != nil {
		return nil, err
	}
	defer fptr.Close()

	token := &oauth2.Token{}
	if err := json.NewDecoder(fptr).Decode(token); err != nil {
		return nil, err
	}

	return config.Client(context.Background(), token), nil
}

func DriveServiceFactory(name string, config *oauth2.Config) (*DriveService, error) {
	client, err := GetClientFromName(name, config)
	if err != nil {
		return nil, err
	}
	svc, err := drive.NewService(context.TODO(), option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}
	return NewDriveService(svc), nil
}

// Manages state and used token/client for the Drive API client.
type DriveService struct {
	// Httpclient associated with the token.
	service *drive.Service
}

func NewDriveService(service *drive.Service) *DriveService {
	return &DriveService{
		service: service,
	}
}

func (s DriveService) ListFolder(currentParent string) (*drive.FileList, error) {
	result, err := s.service.
		Files.List().
		Q(fmt.Sprintf("'%s' in parents", currentParent)).Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}
