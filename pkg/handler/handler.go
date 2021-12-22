package handler


import (
	"aws_rekognition/pkg/model"
	"aws_rekognition/pkg/view"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/gin-gonic/gin"

)

func Handler(c *gin.Context){

	imagePath := "openspace.jpg"

	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Println(err)
		view.ReturnErrorResponse(
			c,
			http.StatusInternalServerError,
			"Internal Server Error",
			"Unable to retrieve images",
		)
		return
	}

	image := &rekognition.Image{
		Bytes: imageBytes,
	}

	accessKeyId := ""
	secretAccessKey := ""
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
	}))

	reko := rekognition.New(sess)

	input := &rekognition.DetectFacesInput{}
	input.SetImage(image)
	output, err := reko.DetectFaces(input)
	if err != nil {
		log.Println(err)
		view.ReturnErrorResponse(
			c,
			http.StatusInternalServerError,
			"Internal Server Error",
			"Unable to send image to AWS",
		)
		return
	}


	var emotion model.Emotion

	for _, label := range output.FaceDetails {
		emotion.Value = *label.Smile.Value
		emotion.Confidence = *label.Smile.Confidence
		}
	}


	c.JSON(200,emotion)
}
