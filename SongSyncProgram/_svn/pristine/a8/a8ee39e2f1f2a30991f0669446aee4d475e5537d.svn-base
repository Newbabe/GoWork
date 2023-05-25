package util

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"io/ioutil"
	"os"
	"runtime/debug"
	"time"
)

var sess *session.Session
var sessSingapore *session.Session

func init() {
	sess, _ = session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")},
	)
	sessSingapore, _ = session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1")},
	)
}
func ExitS3File(filePath string) bool {
	service := s3.New(sessSingapore)
	a := s3.GetObjectInput{
		Bucket: aws.String("herook-file-us"),
		Key:    aws.String(filePath),
	}
	result, err := service.GetObject(&a)

	if err != nil {
		fmt.Println(err)
		return false
	}

	// Make sure to close the body when done with it for S3 GetObject APIs or
	// will leak connections.
	defer result.Body.Close()

	return true
}

const BucketUrl = "https://s3-ap-southeast-1.amazonaws.com/s3file.global-media.com.tw/"
const bucketSong = "s3file.global-media.com.tw/"

func ExistS3SongFile(filePath string) bool {
	svc := s3.New(sessSingapore)
	input := &s3.ListObjectsInput{
		Bucket:  aws.String(bucketSong),
		MaxKeys: aws.Int64(2),
		//Marker:  aws.String("youtube_struct_lyrics/"),
		Prefix: aws.String(filePath),
	}

	result, err := svc.ListObjects(input)
	if err != nil {
		fmt.Println("判断异常", err)
		//if aerr, ok := err.(awserr.Error); ok {
		//	switch aerr.Code() {
		//	case s3.ErrCodeNoSuchBucket:
		//		fmt.Println(s3.ErrCodeNoSuchBucket, aerr.Error())
		//	default:
		//		fmt.Println(aerr.Error())
		//	}
		//} else {
		//	// Print the error, cast err to awserr.Error to get the Code and
		//	// Message from an error.
		//	fmt.Println(err.Error())
		//}
		//fmt.Println("error")
		return false
	}
	//fmt.Println(result.Contents[0].Key , filePath)
	if len(result.Contents) > 0 && filePath == aws.StringValue(result.Contents[0].Key) {
		return true
	}

	return false
}

func CopyFile(orgPath, newPath string) bool {
	svc := s3.New(sess)
	bucket := "herook-file-us"
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		CopySource: aws.String("/" + bucket + "/" + orgPath),
		Key:        aws.String(newPath),
	}
	result, err := svc.CopyObject(input)
	if err != nil {
		fmt.Println(err)
		debug.PrintStack()
		return false
	}
	fmt.Println(result)
	return true
}

func Upload(filePath string, file *os.File) error {
	bucket := "herook-file-singapore"
	uploader := s3manager.NewUploader(sessSingapore)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
		Body:   file,
	})
	return err

}

func UploadUs(filePath string, file *os.File) error {
	bucket := "herook-file-us"
	uploader := s3manager.NewUploader(sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
		Body:   file,
	})
	return err

}

func UploadS3(filePath string, file *os.File) error {
	bucket := bucketSong
	uploader := s3manager.NewUploader(sessSingapore)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filePath),
		Body:   file,
	})
	return err

}

func DownLoadMp3(key string) []byte {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	service := s3.New(sessSingapore)

	out, _ := service.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String("herook-file-singapore"),
		Key:    aws.String(key),
	})
	defer out.Body.Close()
	bytes, _ := ioutil.ReadAll(out.Body)
	return bytes
}
