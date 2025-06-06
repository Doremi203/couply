package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/Doremi203/couply/cloud-functions/libs/token"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Request struct {
	HttpMethod                      string              `json:"httpMethod"`
	Path                            string              `json:"path"`
	Headers                         map[string]string   `json:"headers"`
	MultiValueHeaders               map[string][]string `json:"multiValueHeaders"`
	QueryStringParameters           map[string]string   `json:"queryStringParameters"`
	MultiValueQueryStringParameters map[string][]string `json:"multiValueQueryStringParameters"`
	RequestContext                  struct {
		Identity struct {
			SourceIp  string `json:"sourceIp"`
			UserAgent string `json:"userAgent"`
		} `json:"identity"`
		HttpMethod       string `json:"httpMethod"`
		RequestId        string `json:"requestId"`
		RequestTime      string `json:"requestTime"`
		RequestTimeEpoch int    `json:"requestTimeEpoch"`
	} `json:"requestContext"`
	Body            string `json:"body"`
	IsBase64Encoded bool   `json:"isBase64Encoded"`
}

type RequestBody struct {
	Token       string `json:"token"`
	Bucket      string `json:"bucket"`
	OrderNumber int    `json:"order_number"` // available only for profile photos
}

type Response struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded"`
}

type ResponseBody struct {
	URL string `json:"url"`
}

const (
	verificationPhotosBucket = "couply-verification-photos"
	profilePhotosBucket      = "testing-couply-profile-photos"
)

var AllowedBuckets = []string{
	verificationPhotosBucket,
	profilePhotosBucket,
}

func Handler(ctx context.Context, req Request) (Response, error) {
	fmt.Println("request: ", req)

	corsHeaders := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "*",
		"Access-Control-Allow-Headers": "*",
	}

	if req.HttpMethod == "OPTIONS" {
		return Response{
			StatusCode: 200,
			Headers:    corsHeaders,
			Body:       "",
		}, nil
	}

	if req.HttpMethod != "POST" {
		return Response{
			StatusCode: 405,
			Headers:    corsHeaders,
			Body:       errBody("method Not Allowed"),
		}, nil
	}
	var body RequestBody
	err := json.Unmarshal([]byte(req.Body), &body)
	if err != nil {
		fmt.Println("failed to unmarshal request body:", err)
		return Response{
			StatusCode: 400,
			Headers:    corsHeaders,
			Body:       errBody("bad request"),
		}, nil
	}
	if !slices.Contains(AllowedBuckets, body.Bucket) {
		fmt.Println("bucket not allowed:", body.Bucket)
		return Response{
			StatusCode: 403,
			Headers:    corsHeaders,
			Body:       errBody("forbidden: bucket not allowed"),
		}, nil
	}

	tokenProvider, err := configureTokenProvider()
	if err != nil {
		fmt.Println(err)
		return Response{}, errors.New("internal server error")
	}
	s3Client, err := configureMinioClient()
	if err != nil {
		fmt.Println(err)
		return Response{}, errors.New("internal server error")
	}

	userToken, err := tokenProvider.Parse(body.Token)
	if err != nil {
		fmt.Println(err)
		return Response{
			Headers:    corsHeaders,
			StatusCode: 401,
		}, nil
	}

	photoURLGenerator := newObjectStoragePhotoURLGenerator(s3Client)

	objectKey, err := generateKey(userToken.GetUserID(), body.Bucket, body.OrderNumber)
	if err != nil {
		fmt.Println(err)
		return Response{}, errors.New("internal server error")
	}

	uploadURL, err := photoURLGenerator.GenerateUpload(ctx, objectKey, body.Bucket)
	if err != nil {
		fmt.Println(err)
		return Response{}, errors.New("internal server error")
	}

	respBody, err := json.Marshal(ResponseBody{
		URL: uploadURL,
	})
	if err != nil {
		fmt.Println(err)
		return Response{}, errors.New("internal server error")
	}

	return Response{
		StatusCode: 200,
		Headers:    corsHeaders,
		Body:       string(respBody),
	}, nil
}

func configureTokenProvider() (token.Provider, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, fmt.Errorf("JWT_SECRET_KEY must be set")
	}

	return token.NewJWTProvider(token.Config{
		SecretKey: secretKey,
	}), nil
}

func configureMinioClient() (*minio.Client, error) {
	accessKey := os.Getenv("S3_ACCESS_KEY")
	secretKey := os.Getenv("S3_SECRET_KEY")
	if accessKey == "" || secretKey == "" {
		return nil, fmt.Errorf("S3_ACCESS_KEY and S3_SECRET_KEY must be set")
	}

	client, err := minio.New("storage.yandexcloud.net", &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	return client, nil
}

func generateKey(userID uuid.UUID, bucket string, orderNumber int) (string, error) {
	switch bucket {
	case profilePhotosBucket:
		photoID, err := uuid.NewRandom()
		if err != nil {
			return "", fmt.Errorf("failed to generate photo id: %w", err)
		}
		return fmt.Sprintf("users/%s/slot/%d/%s.jpg", userID.String(), orderNumber, photoID), nil
	default:
		return fmt.Sprintf("%s/%s.jpg", userID.String(), uuid.New().String()), nil
	}
}

func newObjectStoragePhotoURLGenerator(
	client *minio.Client,
) *objectStoragePhotoURLGenerator {
	return &objectStoragePhotoURLGenerator{
		client: client,
	}
}

type objectStoragePhotoURLGenerator struct {
	client *minio.Client
}

func (g *objectStoragePhotoURLGenerator) GenerateUpload(
	ctx context.Context,
	key string,
	bucket string,
) (string, error) {
	expires := time.Minute * 15
	uploadURL, err := g.client.PresignedPutObject(ctx, bucket, key, expires)
	if err != nil {
		return "", fmt.Errorf(
			"generate signed upload url bucket: %s, key: %s %w",
			bucket,
			key,
			err,
		)
	}

	return uploadURL.String(), nil
}

func errBody(msg string) string {
	return fmt.Sprintf(`{"error": "%s"}`, msg)
}
