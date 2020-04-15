# golang-upload-file-google-storage
Packages&nbsp
`
 go get google.golang.org/api/storage/v1
 `
 &nbsp
 `
 go get golang.org/x/oauth2
 `
 &nbsp
 setting acces to GCS
 &nbsp
```
authconf := &jwt.Config{
		Email:      "your email acces to GCS",
		PrivateKey: []byte("key"),
		Scopes:     []string{storage.DevstorageReadWriteScope},
		TokenURL:   "https://accounts.google.com/o/oauth2/token",
	}
```
