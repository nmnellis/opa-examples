package download

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"cloud.google.com/go/storage"
)

type BundleDownloader interface {
	Download(path string) ([]byte, error)
	GetETag(path string) (string, error)
	Exists(path string) (bool, error)
}

type FileBundleDownloader struct {
	Directory string
}

func (d *FileBundleDownloader) Download(fileName string) ([]byte, error) {
	filePath := strings.TrimSuffix(d.Directory, "/") + "/" + strings.TrimPrefix(fileName, "/")
	return ioutil.ReadFile(filePath)
}
func (d *FileBundleDownloader) Exists(fileName string) (bool, error) {
	filePath := strings.TrimSuffix(d.Directory, "/") + "/" + strings.TrimPrefix(fileName, "/")
	if _, err := os.Stat(filePath); err != nil {
		return false, nil
	}
	return true, nil
}
func (d *FileBundleDownloader) GetETag(fileName string) (string, error) {
	filePath := strings.TrimSuffix(d.Directory, "/") + "/" + strings.TrimPrefix(fileName, "/")
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

type GCSBundleDownloader struct {
	Bucket string
	Client *storage.Client
}

func (d *GCSBundleDownloader) Download(path string) ([]byte, error) {
	objHandle := d.Client.Bucket(d.Bucket).Object(path)
	reader, err := objHandle.NewReader(context.Background())
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return nil, fmt.Errorf("item not found gs://%v/%v", path, d.Bucket)
		}
		return nil, fmt.Errorf("error finding object gs://%v/%v %v", d.Bucket, path, err.Error())
	}

	return ioutil.ReadAll(reader)
}

func (d *GCSBundleDownloader) GetETag(path string) (string, error) {
	objHandle := d.Client.Bucket(d.Bucket).Object(path)
	// Get the attributes first
	attrs, err := objHandle.Attrs(context.Background())
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return "", fmt.Errorf("item not found gs://%v/%v", path, d.Bucket)
		}
		return "", fmt.Errorf("error getting object attributes gs://%v/%v %v", d.Bucket, path, err.Error())

	}

	return attrs.Etag, nil
}
func (d *GCSBundleDownloader) Exists(path string) (bool, error) {
	objHandle := d.Client.Bucket(d.Bucket).Object(path)
	// Get the attributes first
	_, err := objHandle.Attrs(context.Background())
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return false, nil
		}
		return false, fmt.Errorf("error getting object attributes gs://%v/%v %v", d.Bucket, path, err.Error())
	}

	return true, nil
}
