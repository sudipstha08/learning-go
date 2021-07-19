package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"google.golang.org/api/option"
)

// InitStorageClient init the storage client
func initStorageClient() *storage.Client {
	ctx := context.Background()
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		panic("Unable to load serviceAccountKey.json file")
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	client, err := storage.NewClient(ctx, opt)
	if err != nil {
		panic("Error creating GCP stroage client")
	}
	return client
}


type gcpStorageService struct {
	StorageClient *storage.Client
}

// GCPStorageService Interface
type GCPStorageService interface {
	UploadFile(req *http.Request,file multipart.File, fileHeader *multipart.FileHeader) (*url.URL, error)
}

// NewFirebaseAuthService constructor
func NewGCPStorageService() GCPStorageService {
	sc := initStorageClient()
	return &gcpStorageService{
		StorageClient: sc,
	}
}

const googleStorage = "https://storage.googleapis.com/"

func (sc *gcpStorageService) UploadFile(req *http.Request,file multipart.File, fileHeader *multipart.FileHeader) (*url.URL, error) {
	fileName := fileHeader.Filename

	directory := strings.Split(fileName,"/")
	nameAndExtension := strings.Split(directory[len(directory)-1],".")
	fileExtension := nameAndExtension[len(nameAndExtension)-1]
	uniqueId := uuid.New()

	fileName = strings.Replace(uniqueId.String(), "-", "", -1)
	fileName = fmt.Sprintf("%v.%v",fileName,fileExtension)
	directory[len(directory)-1] = fileName

	fileName = strings.Join(directory,"/")

	// storage bucket name
	bucketName := "learning-go-7da5a.appspot.com/files"
	bucket := sc.StorageClient.Bucket(bucketName)
	ctx := req.Context()
	obj := bucket.Object(fileName).NewWriter(ctx)

	obj.ContentType = fileHeader.Header.Get("Content-Type")
	if _, err := io.Copy(obj, file); err != nil {
		return nil,err
	}

	// Either
	// obj.Metadata = map[string]string{
	// 	"x-goog-acl":"public-read",
	// }

	// or use access control list for given object
	acl := bucket.Object(obj.Name).ACL() 
	
	// close the object
	if err := obj.Close(); err != nil {
		return nil,err
	}
	
	// ...or continued
	// set access control list to give any user the permission to view the uploaded file
	if err := acl.Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return nil, err
	}
	return url.Parse(googleStorage + bucketName + "/" + obj.Attrs().Name) // complete url to access the uploaded file
}