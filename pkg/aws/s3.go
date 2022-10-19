package aws

import (
	"bytes"
	"context"
	"errors"
	"io"
	"strings"
	"time"

	"github.com/Haato3o/poogie/core/cache"
	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/persistence/bucket"
	"github.com/Haato3o/poogie/core/tracing"
	"github.com/Haato3o/poogie/pkg/log"
	"github.com/Haato3o/poogie/pkg/memcache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	MOST_RECENT_KEY = "most-recent-version"
	AWS_REGION      = "us-east-1"
	AWS_ENDPOINT    = "https://sfo3.digitaloceanspaces.com"
)

type S3Bucket struct {
	connection *s3.S3
	cache      cache.ICache
	bucket     string
	prefix     string
	fileType   string
}

// Delete implements bucket.IBucket
func (b *S3Bucket) Delete(ctx context.Context, name string) {
	txn := tracing.FromContext(ctx)
	segment := txn.StartSegment("S3Bucket.Delete")
	defer segment.End()

	input := s3.DeleteObjectInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(b.prefix + name + b.fileType),
	}

	b.connection.DeleteObjectWithContext(ctx, &input)
}

// DownloadToStream implements bucket.IBucket
func (b *S3Bucket) DownloadToStream(ctx context.Context, name string) (bucket.StreamedFile, error) {
	txn := tracing.FromContext(ctx)
	segment := txn.StartSegment("S3Bucket.DownloadToStream")
	defer segment.End()

	input := s3.GetObjectInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(b.prefix + name + b.fileType),
	}

	response, err := b.connection.GetObjectWithContext(ctx, &input)

	return bucket.StreamedFile{
		Reader: response.Body,
		Size:   *response.ContentLength,
		Type:   *response.ContentType,
	}, err
}

// UploadFromStream implements bucket.IBucket
func (b *S3Bucket) UploadFromStream(ctx context.Context, name string, file io.Reader) (bool, error) {
	txn := tracing.FromContext(ctx)
	segment := txn.StartSegment("S3Bucket.UploadFromStream")
	defer segment.End()

	uploader := s3manager.NewUploaderWithClient(b.connection)
	input := s3manager.UploadInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(b.prefix + name + b.fileType),
		Body:   file,
	}

	_, err := uploader.UploadWithContext(ctx, &input, func(u *s3manager.Uploader) {
		u.Concurrency = 10
	})

	log.Error("failed to upload to bucket", err)

	return err == nil, err
}

// Upload implements bucket.IBucket
func (b *S3Bucket) Upload(ctx context.Context, name string, data []byte) (bool, error) {
	txn := tracing.FromContext(ctx)
	segment := txn.StartSegment("S3Bucket.Upload")
	defer segment.End()

	uploader := s3manager.NewUploaderWithClient(b.connection)
	buffer := bytes.NewReader(data)

	input := s3manager.UploadInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(b.prefix + name + b.fileType),
		Body:   buffer,
		ACL:    aws.String("public-read"),
	}
	_, err := uploader.UploadWithContext(ctx, &input)

	log.Error("failed to upload to bucket", err)

	return err == nil, err
}

// FindBy implements bucket.IBucket
func (b *S3Bucket) FindBy(ctx context.Context, name string) ([]byte, error) {
	txn := tracing.FromContext(ctx)
	segment := txn.StartSegment("S3Bucket.FindBy")
	defer segment.End()

	file, hasCachedFile := b.cache.Get(name)

	if hasCachedFile {
		return file.([]byte), nil
	}

	query := &s3.GetObjectInput{
		Bucket: aws.String(b.bucket),
		Key:    aws.String(b.prefix + name + b.fileType),
	}

	downloader := s3manager.NewDownloaderWithClient(b.connection)

	buffer := aws.NewWriteAtBuffer([]byte{})

	_, err := downloader.DownloadWithContext(ctx, buffer, query, func(d *s3manager.Downloader) {
		d.Concurrency = 10
	})

	if err != nil {
		return nil, errors.New("no file with given name was found")
	}

	b.cache.Set(name, buffer.Bytes())

	return buffer.Bytes(), nil
}

// FindMostRecent implements bucket.IBucket
func (b *S3Bucket) FindMostRecent(ctx context.Context) (string, error) {
	txn := tracing.FromContext(ctx)
	segment := txn.StartSegment("S3Bucket.FindMostRecent")
	defer segment.End()

	name, hasCachedName := b.cache.Get(MOST_RECENT_KEY)

	if hasCachedName {
		return name.(string), nil
	}

	query := &s3.ListObjectsInput{
		Bucket: aws.String(b.bucket),
		Prefix: &b.prefix,
	}

	resp, err := b.connection.ListObjectsWithContext(ctx, query)

	if err != nil {
		return "", err
	}

	var mostRecent *s3.Object

	for _, item := range resp.Contents {
		if mostRecent == nil || item.LastModified.Unix() > mostRecent.LastModified.Unix() {
			mostRecent = item
		}
	}

	if mostRecent == nil {
		return "", errors.New("no files found")
	}

	fileName := removeSuffixAndPrefix(*mostRecent.Key, b.fileType, b.prefix)

	b.cache.Set(MOST_RECENT_KEY, fileName)

	return fileName, nil
}

func New(configuration *config.ApiConfiguration, prefix, fileType string) bucket.IBucket {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(AWS_REGION),
		Credentials: credentials.NewStaticCredentials(
			configuration.AwsAccountKey,
			configuration.AwsAccountSecret,
			"",
		),
		Endpoint: aws.String(AWS_ENDPOINT),
	})

	if err != nil {
		log.Error("failed to create S3 instance", err)
		return nil
	}

	cache := memcache.New(5 * time.Hour)
	service := s3.New(session)

	return &S3Bucket{
		connection: service,
		cache:      cache,
		bucket:     configuration.AwsBucketName,
		prefix:     prefix,
		fileType:   fileType,
	}
}

func removeSuffixAndPrefix(str, suffix, prefix string) string {
	str = strings.Replace(str, suffix, "", 1)
	return strings.Replace(str, prefix, "", 1)
}
