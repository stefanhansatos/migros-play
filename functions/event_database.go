package functions

import (
	"cloud.google.com/go/functions/metadata"
	"context"
	"fmt"
	"log"
)

// RTDBEvent is the payload of a RTDB event.
type RTDBEvent struct {
	Data  interface{} `json:"data"`
	Delta interface{} `json:"delta"`
}

// HelloRTDB handles changes to a Firebase RTDB.
func HelloRTDB(ctx context.Context, e RTDBEvent) error {
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	log.Printf("Function triggered by change to: %v", meta.Resource)
	log.Printf("%+v", e)
	return nil
}
