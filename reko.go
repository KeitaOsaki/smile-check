package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"io/ioutil"
	"log"
)

func main() {
	imagePath := ""
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Println(err)
	}

	accessKeyId := ""
	secretAccessKey := ""
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
	}))

	reko := rekognition.New(sess)

	params := &rekognition.DetectFacesInput{
		Image: &rekognition.Image{
			Bytes: imageBytes,
		},
		Attributes: []*string{
			aws.String("ALL"),
		},
	}

	res, err := reko.DetectFaces(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, face := range res.FaceDetails {

		if *face.Smile.Value == true {
			fmt.Printf("ナイスSmile!!%f\n",*face.Smile.Confidence)
		} else {
			fmt.Printf("どしたん？話聞こか？%f\n",*face.Smile.Confidence)
		}
		fmt.Printf("信頼度%f\n",*face.Smile.Confidence)
	}

}
