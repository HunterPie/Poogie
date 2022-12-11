package main

import (
	"context"
	"fmt"
	"golang.org/x/exp/slices"
	"log"
	"strings"

	"github.com/Haato3o/poogie/core/config"
	"github.com/Haato3o/poogie/pkg/mongodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var apiConfig config.ApiConfiguration

type S3Wrapper struct {
	connection *s3.S3
	bucket     string
}

func NewS3Wrapper() *S3Wrapper {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			apiConfig.AwsAccountKey,
			apiConfig.AwsAccountSecret,
			"",
		),
		Endpoint: aws.String("https://sfo3.digitaloceanspaces.com"),
	})

	if err != nil {
		log.Fatal(err)
	}

	service := s3.New(session)

	return &S3Wrapper{
		connection: service,
		bucket:     apiConfig.AwsBucketName,
	}
}

func (s *S3Wrapper) DeleteBackups(ctx context.Context, userId string, files []string) int {
	objects := make([]*s3.ObjectIdentifier, 0)

	for _, file := range files {
		objects = append(objects, &s3.ObjectIdentifier{
			Key: aws.String(fmt.Sprintf("backups/%s/%s", userId, file)),
		})
	}

	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(s.bucket),
		Delete: &s3.Delete{
			Objects: objects,
			Quiet:   aws.Bool(false),
		},
	}

	_, err := s.connection.DeleteObjectsWithContext(ctx, input)

	if err != nil {
		log.Println(err)
		return 0
	}

	return len(objects)
}

func (s *S3Wrapper) ListUsers(ctx context.Context) map[string][]string {
	query := &s3.ListObjectsInput{
		Bucket: aws.String(s.bucket),
		Prefix: aws.String("backups/"),
	}

	resp, err := s.connection.ListObjectsWithContext(ctx, query)

	if err != nil {
		log.Fatalln(err)
	}

	users := make(map[string][]string, 0)

	for _, item := range resp.Contents {
		if *item.Key == "backups/" {
			continue
		}

		split := strings.Split(*item.Key, "/")
		userId := split[1]
		file := split[2]

		if _, ok := users[userId]; !ok {
			users[userId] = make([]string, 0)
		}

		users[userId] = append(users[userId], file)
	}

	return users
}

func main() {
	_ = godotenv.Load()
	_ = envconfig.Process("POOGIE", &apiConfig)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, _ := mongodb.New(
		apiConfig.DatabaseUri,
		apiConfig.DatabaseName,
		false,
	)

	repository := db.GetBackupsRepository()

	bucket := NewS3Wrapper()

	users := bucket.ListUsers(ctx)

	for key, value := range users {
		backupsQt := len(value)
		if backupsQt <= 2 {
			continue
		}

		remoteBackups := repository.FindAllByUserId(ctx, key)

		if len(remoteBackups) == backupsQt {
			continue
		}

		remoteBackupNames := make([]string, 0)
		for _, remote := range remoteBackups {
			remoteBackupNames = append(remoteBackupNames, fmt.Sprintf("%s.zip", remote.Id))
		}

		backupsToDelete := make([]string, 0)

		for _, bkp := range value {
			if !slices.Contains(remoteBackupNames, bkp) {
				backupsToDelete = append(backupsToDelete, bkp)
			}
		}

		//log.Printf("Users %s has %d backups, %d to delete", key, len(value), len(backupsToDelete))

		deleted := bucket.DeleteBackups(ctx, key, backupsToDelete)

		log.Printf("Deleted %d backups for user: %s", deleted, key)
	}
}
