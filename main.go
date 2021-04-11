
package main

import (
	"context"
	// "fmt"
	"log"
	"os"
	firebase "firebase.google.com/go"
	// "github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	// validator "gopkg.in/go-playground/validator.v9"
	// "net/http"
	// "strconv"
	// "path/filepath"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
)

var ginLambda *ginadapter.GinLambda
// https://godoc.org/firebase.google.com/go
var firapp *firebase.App
var ctrl *DbController


func init(){
	firebaseCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")	
	firebaseKey := os.Getenv("FIREBASE_KEY")	
	log.Println("FIREBASE URL", firebaseCreds)
	log.Println("FIREBASE TOKEN", firebaseKey)
	firapp, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Printf("error initializing app: %v\n", err)
	}
	log.Println("FIREBASE APP", firapp)
	dbURI := os.Getenv("DATABASE_LOCAL_URL")
	log.Println("DATABASE URI", dbURI)
	ctrl.InitDbConnection(dbURI)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	router := gin.New()
	if ginLambda == nil {
		router.Use(gin.Logger())
		// router.POST("/user/edit", UpdateUserBioHandler)
		router.GET("/citylist", GetAllCitiesHandler) // verify post_token
		router.GET("/profile", GetUserProfileHandler) // verify post_token
		router.GET("/check_session", CheckSessionHandler) // verify get_token 
		router.POST("/signin", SignInHandler)
		router.GET("/signup", SignUpHandler)
		// router.POST("/signout", SignOutHandler)
	}
	ginLambda = ginadapter.New(router)
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}