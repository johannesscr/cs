package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/upload/attachment", handler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

/*
The pipeline we want.
Virus Scanner
Attachment Scalar
Mime Type -> Preview Factory -> Preview Generator
Key -> Encryption -> Storage Factory -> Storage
*/

// handler handles the upload attachment request
func handler(w http.ResponseWriter, r *http.Request) {
	// suppose the company ID is in the url values
	companyId := r.URL.Query().Get("company-id")
	// get the attachment from the multipart form
	file, fh, err := r.FormFile("attachment")
	if err != nil {
		w.WriteHeader(400)
		_, _ = w.Write([]byte(`Unable to read attachment`))
		return
	}
	xb := make([]byte, 0)
	_, _ = file.Read(xb)
	// now that we have the company ID we can get the storage using the factory
	// and since the factory does all the storage logic we know we can always
	// call the upload method from the interface.
	//
	// now we are free to inject (dependency injection here) any storage
	// implementation or swap them out.
	enc := newEncryption(getEncryptionKey())
	storage := storageFactory(companyId, enc)
	storage.upload(xb, fh.Filename)
}

type Storage interface {
	upload(xb []byte, filename string)
	download()
}

func getEncryptionKey() string {
	return ""
}

type Encryption interface {
	encrypt([]byte) ([]byte, error)
}

func newEncryption(key string) Encryption {
	return nil
}

// storageFactory returns the storage location and type based on the company ID.
func storageFactory(companyId string, encryption Encryption) Storage {
	if companyId == "0" {
		// Constructor Injection
		return SFTP{
			Encryption: encryption,
		}
	}
	return S3{
		Encryption: encryption,
	}
}

type S3 struct {
	Encryption Encryption
}

func (s S3) upload(xb []byte, filename string) {
	xb, _ = s.Encryption.encrypt(xb)
}
func (s S3) download() {}

type SFTP struct {
	Encryption Encryption
}

func (s SFTP) upload(xb []byte, filename string) {
	xb, _ = s.Encryption.encrypt(xb)
}
func (s SFTP) download() {}
