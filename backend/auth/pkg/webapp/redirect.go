package webapp

import (
	"context"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func redirectHeaderMatcher(_ context.Context, w http.ResponseWriter, _ proto.Message) error {
	headers := w.Header()
	if location, ok := headers["Grpc-Metadata-Location"]; ok {
		w.Header().Set("Location", location[0])
		w.WriteHeader(http.StatusFound)
	}

	return nil
}
