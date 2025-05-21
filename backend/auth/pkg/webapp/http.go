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
		w.Header().Del("Grpc-Metadata-Location")
		w.WriteHeader(http.StatusFound)
	}

	return nil
}

func setCookieHeaderMatcher(_ context.Context, w http.ResponseWriter, _ proto.Message) error {
	headers := w.Header()
	if cookies, ok := headers["Grpc-Metadata-Set-Cookie"]; ok {
		if len(cookies) == 0 {
			return errors.Errorf("set cookie header value not found")
		}
		for _, cookie := range cookies {
			w.Header().Add("Set-Cookie", cookie)
		}
	}
	w.Header().Del("Grpc-Metadata-Set-Cookie")

	return nil
}
