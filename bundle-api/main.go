package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/nmnellis/opa-examples/bundle-api/download"
)

var (
	storageType      = flag.String("storage-type", "file", "Where are the bundles stored?<file,gcs>")
	storagePath      = flag.String("storage-path", "", "directory location or gcs bucket")
	port             = flag.String("port", "8080", "Server port")
	host             = flag.String("host", "0.0.0.0", "Server host")
	bundleDownloader download.BundleDownloader
)

func main() {
	flag.Parse()

	switch *storageType {
	case "file":
		bundleDownloader = &download.FileBundleDownloader{
			Directory: *storagePath,
		}
	case "gcs":
		client, err := storage.NewClient(context.Background())
		if err != nil {
			panic(err)
		}
		bundleDownloader = &download.GCSBundleDownloader{
			Bucket: *storagePath,
			Client: client,
		}
	default:
		panic("unknown storage type " + *storageType)
	}

	http.HandleFunc("/", bundlehandler)
	log.Printf("Listening on port %v:%v", *host, *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", *host, *port), nil))
}

//bundlehandler http handler
func bundlehandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request received %v", r.URL.Path)

	file := r.RequestURI
	exists, err := bundleDownloader.Exists(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, fmt.Sprintf("file does not exist %v", file), http.StatusBadRequest)
		return
	}

	revision := r.Header.Get("If-None-Match")
	etag, err := bundleDownloader.GetETag(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Etag", etag)

	if revision == etag {
		fmt.Printf("OPA already has latest version %v for file %v\n", revision, file)
		w.WriteHeader(http.StatusNotModified)
		return
	}

	bundle, err := bundleDownloader.Download(file)
	if err != nil {
		msg := fmt.Sprintf("error downloading file %v %v", file, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	fmt.Printf("Downloaded %v,latest ETag %v\n", file, etag)

	w.Header().Set("Content-Type", "application/gzip")
	//deliver bundle
	io.Copy(w, bytes.NewReader(bundle))
}
