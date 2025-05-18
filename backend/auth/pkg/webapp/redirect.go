package webapp

import (
	"context"
	"net/http"

	"github.com/Doremi203/couply/backend/auth/pkg/errors"
	"google.golang.org/protobuf/proto"
)

func redirectHeaderMatcher(_ context.Context, w http.ResponseWriter, _ proto.Message) error {
	headers := w.Header()
	if location, ok := headers["Grpc-Metadata-Location"]; ok {
		if len(location) == 0 {
			return errors.Errorf("location header not found")
		}
		w.Header().Set("Location", location[0])
		w.WriteHeader(http.StatusFound)
	}

	return nil
}
