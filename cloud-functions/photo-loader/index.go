package main

import (
	"context"
	"errors"
	"fmt"
	"mime"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

type Request struct {
	Messages []struct {
		EventMetadata struct {
			EventId        string    `json:"event_id"`
			EventType      string    `json:"event_type"`
			CreatedAt      time.Time `json:"created_at"`
			TracingContext struct {
				TraceId      string `json:"trace_id"`
				SpanId       string `json:"span_id"`
				ParentSpanId string `json:"parent_span_id"`
			} `json:"tracing_context"`
			CloudId  string `json:"cloud_id"`
			FolderId string `json:"folder_id"`
		} `json:"event_metadata"`
		Details struct {
			BucketId string `json:"bucket_id"`
			ObjectId string `json:"object_id"`
		} `json:"details"`
	} `json:"messages"`
}

type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded"`
}

const (
	host   = "rc1a-1ns3p57pcpp20u6s.mdb.yandexcloud.net"
	port   = 6432
	user   = "doremi"
	dbname = "matcher"
)

func Handler(ctx context.Context, req Request) (Response, error) {
	pass := os.Getenv("PG_PASSWORD")
	if pass == "" {
		return Response{}, errors.New("PG_PASSWORD environment variable is not set")
	}
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=verify-full&target_session_attrs=read-write",
		user, pass, host, port, dbname)
	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		return Response{}, fmt.Errorf("pgx connect to db: %w", err)
	}
	defer conn.Close(ctx)

	parts := strings.Split(req.Messages[0].Details.ObjectId, "/")
	userID := parts[1]
	photoParts := strings.Split(parts[4], ".")
	photoExtension := photoParts[1]
	mimeType := mime.TypeByExtension("." + photoExtension)
	orderNumber, err := strconv.Atoi(parts[3])
	if err != nil {
		return Response{}, fmt.Errorf("parse order number: %w", err)
	}

	const query = `
		INSERT INTO photos(user_id, order_number, object_key, mime_type, uploaded_at)
		VALUES ($1, $2, $3, $4, now())
		ON CONFLICT (user_id, order_number) DO UPDATE
		SET object_key = $3, mime_type = $4, uploaded_at = now()
	`

	fmt.Println("args: ", userID, orderNumber, req.Messages[0].Details.ObjectId, mimeType)
	_, err = conn.Exec(ctx, query,
		userID,
		orderNumber,
		req.Messages[0].Details.ObjectId,
		mimeType,
	)
	if err != nil {
		return Response{}, fmt.Errorf("upsert photo record: %w", err)
	}

	return Response{
		StatusCode: 200,
	}, nil
}
