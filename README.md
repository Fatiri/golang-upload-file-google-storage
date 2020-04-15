# golang-upload-file-google-storage
Packages\
`
go get google.golang.org/api/storage/v1
 `\
 `
 go get golang.org/x/oauth2
 `

 setting acces to GCS
```\
authconf := &jwt.Config{
		Email:      "your email acces to GCS",
		PrivateKey: []byte("key"),
		Scopes:     []string{storage.DevstorageReadWriteScope},
		TokenURL:   "https://accounts.google.com/o/oauth2/token",
	}
```
