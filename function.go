package main

import (
	"context"
	"log"
)

// GCSEvent is the payload of a GCS event. Please refer to the docs for
// additional information regarding GCS events.
type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// handler prints a message when a file is changed in a Cloud Storage bucket.
func handler(ctx context.Context, e GCSEvent) error {
	log.Printf("Processing file: %s in bucket: %s", e.Name, e.Bucket)
	return nil
}
