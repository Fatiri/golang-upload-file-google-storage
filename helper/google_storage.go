package helper

import (
	"errors"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/jwt"
	storage "google.golang.org/api/storage/v1"
)

// Authenticate to Google Cloud Storage and return handler
func getGCS() (service *storage.Service, err error) {

	// OAuth2 info from GCS console
	authconf := &jwt.Config{
		Email:      "ilhamfatiri@email.com",
		PrivateKey: []byte("key"),
		Scopes:     []string{storage.DevstorageReadWriteScope},
		TokenURL:   "https://accounts.google.com/o/oauth2/token",
	}

	client := authconf.Client(oauth2.NoContext)

	service, err = storage.New(client)
	if err != nil {
		return nil, errors.New("problem saving file to gcs")
	}

	return

}

// Upload a file to Google Cloud Storage
func UploadGCS(filepath, filename string) (url *storage.Object, err error) {

	service, err := getGCS()
	if err != nil {
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, errors.New("problem opening file for gcs")
	}
	defer file.Close()

	object := &storage.Object{
		Name:         filename,
		CacheControl: "public, max-age=31536000",
	}

	linkUrl, err := service.Objects.Insert("thisbucketgcs", object).Media(file).Do()
	if err != nil {
		return nil, err
	}

	return linkUrl, nil

}

// Delete a file from Google Cloud Storage
func DeleteGCS(object string) (err error) {

	service, err := getGCS()
	if err != nil {
		return
	}

	err = service.Objects.Delete("thisbucketgcs", object).Do()
	if err != nil {
		return errors.New("problem deleting gcs file")
	}

	return

}
