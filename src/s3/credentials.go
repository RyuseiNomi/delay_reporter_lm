package s3Worker

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func GetCredential() (*session.Session, error) {

	var sess *session.Session

	if os.Getenv("AWS_SAM_LOCAL") == "true" {
		/* Yield credential for local */
		log.Printf("Start process getting credential as a local")
		credential := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
		sess, _ = session.NewSession(&aws.Config{
			Credentials:      credential,
			Region:           aws.String("ap-northeast-1"),
			Endpoint:         aws.String("http://172.18.0.2:9000"),
			S3ForcePathStyle: aws.Bool(true),
		})
	} else {
		/* Yield credential for production */
		log.Printf("Start process getting credential as a production")
		sess, _ = session.NewSession(&aws.Config{
			Region:           aws.String("ap-northeast-1"),
			S3ForcePathStyle: aws.Bool(true),
		})
	}

	_, err := sess.Config.Credentials.Get()

	if err != nil {
		return nil, err
	}

	return sess, nil
}
