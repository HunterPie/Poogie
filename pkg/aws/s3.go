package aws

import (
	"errors"
	"time"

	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/core/domain/cache"
	"github.com/Haato3o/poogie/core/domain/persistence/bucket"
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
		// TODO: Add logging
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

// FindBy implements bucket.FileBucket
func (b *S3Bucket) FindBy(name string) ([]byte, error) {
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

	_, err := downloader.Download(buffer, query, func(d *s3manager.Downloader) {
		d.Concurrency = 4
	})

	if err != nil {
		return nil, errors.New("no file with given name was found")
	}

	b.cache.Set(name, buffer.Bytes())

	return buffer.Bytes(), nil
}

func (b *S3Bucket) FindMostRecent() (string, error) {
	name, hasCachedName := b.cache.Get(MOST_RECENT_KEY)

	if hasCachedName {
		return name.(string), nil
	}

	query := &s3.ListObjectsInput{
		Bucket: aws.String(b.bucket),
		Prefix: &b.prefix,
	}

	resp, err := b.connection.ListObjects(query)

	if err != nil {
		return "", err
	}

	var mostRecent *s3.Object

	for _, item := range resp.Contents {
		if mostRecent == nil || item.LastModified.Unix() > mostRecent.LastModified.Unix() {
			mostRecent = item
		}
	}

	b.cache.Set(MOST_RECENT_KEY, *mostRecent.Key)

	return *mostRecent.Key, nil
}
