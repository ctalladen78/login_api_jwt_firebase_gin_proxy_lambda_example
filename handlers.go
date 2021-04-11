package main

import (
	// "context"
	"fmt"
	// "log"
	// "os"
	"net/http"

	// firebase "firebase.google.com/go"
	// "github.com/joho/godotenv"
	"github.com/lithammer/shortuuid"
	"github.com/gin-gonic/gin"
	// validator "gopkg.in/go-playground/validator.v9"
)

type SessionInfo struct {
	FireToken string `json:"fire_token"`
	UserId    string `json:"userid"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ObjectId string `json:"object_id"` // hash key "USER-abc123"
	ObjectSk string `json:"object_sk"`
	// all global secondary index
	CreatedAt      string `json:"created_at" binding:"exists"`         // private
	Name           string `json:"display_name" binding:"exists,min=5"` // public
	AvatarLink     string `json:"avatar_link" binding:"exists,min=5"`  // public
	Bio            string `json:"bio" binding:"exists,min=5"`          // public
	Email          string `json:"email" binding:"exists,min=5"`        // private
	BgColor        string `json:"color"`                               // public
	CurrentCityPID string `json:"current_city_pid"`
	// SocialLink string `json:"social_link"` // public
	// WebLink    string `json:"web_link"`    // public
	// BitcoinAddress  string `json:"bitcoin"`        // public
	// EthereumAddress string `json:"ethereum"`       // public
	// Phone      string `json:"phone"`       // public whatsapp
	// IsEmailVerified bool   `json:"email_verified"` // private
	// CurrentGeo map[string]float64 `json:"current_geo"` // public query index
	CurrentGeo GeoPoint `json:"current_geo"` // public query index
	// LastLogin       string `json:"last_login"`     // public
	// LastUpdated     string `json:"last_updated"`   // private
}

type GeoPoint struct {
	Lat  float32 `json:"Lat"`
	Long float32 `json:"Long"`
}

func SignInHandler(c *gin.Context){
	tok := c.GetHeader("Authorization")
	firebaseSession := &SessionInfo{}
	err := c.BindJSON(firebaseSession)
	if err != nil {
		c.AbortWithError(501, err)
	}
	fmt.Println("SIGNIN INFO ", firebaseSession)
	// post_token, err := VerifyFireToken(firebaseSession.FireToken)
	get_tok, claims,err := VerifyJWTToken(tok)
	fmt.Println("VERIFIED CLAIMS ", claims["exp"], claims["userid"])
	if err != nil {
		fmt.Println("SIGNIN ERROR ", err)
		c.AbortWithError(501, err)
	}
	// user, err := checkUserExists(firebaseSession.UserId)
	// if err != nil || u == nil {
	// 	c.AbortWithError(501, err)
	// }
	// c.JSONP(http.StatusOK, gin.H{"post_tok": sess, "get_tok": tok, "user": user})
	c.JSONP(http.StatusOK, gin.H{"get_token": get_tok})
}

func SignUpHandler(c *gin.Context){
	// firtok := c.GetHeader("Authorization")
	// sess, err := VerifyFireToken(firtok)
	// if err != nil {
	// 	c.AbortWithError(501, err)
	// }
	// fmt.Println("SESSION ", sess)
	// u := &User{}
	// err = c.BindJSON(u)
	// if err != nil {
	// 	c.AbortWithError(501, err)
	// 	return
	// }
	id := shortuuid.New()
	fmt.Println("SESSION ", id)
	// user  := &User{}
	// var user *User
	// uchan := make(chan *User)
	// checkUserExists(user.objectId)
	// go ctrl.SaveUserToDBGoroutine(id, u.Email, u.Name, u.Bio, u.AvatarLink, uchan)
	// user = <-uchan
	// close(uchan)
	// if err != nil {
	// 	c.AbortWithError(501, err)
	// }
	uid := c.Query("user_id")
	// "2394772470822913"
	get_tok, err := CreateGetToken(uid)
	if err != nil {
		c.AbortWithError(501, err)
	}
	c.JSONP(200, gin.H{"token": get_tok})
}

func CheckSessionHandler(c *gin.Context){
	token := c.GetHeader("Authorization")
	fmt.Println("TOKEN ", token)
	// TODO Validate token
	tok, claims, _ := VerifyJWTToken(token)
	fmt.Println("CLAIM.USERID ", claims)
	fmt.Println("NEW TOKEN ", tok)
	// sess, err := VerifyFireToken(token)
	// if err != nil {
	// 	fmt.Println("TOKEN ERROR", err)
	// }
	// fmt.Println("SESSION ", sess)

	c.JSONP(200, gin.H{"data": nil})
}

func checkUserExists(uid string) (interface{}, error) {
	uid = fmt.Sprintf("USER-%s", uid)
	// u, err := ctrl.GetUser(uid)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("CHECK USER EXIST ", u)
	// if u.(User).ObjectId == uid {
	// 	return u, nil // user exists continue signin
	// }
	// return nil, err // user not exists prompt sign up
	return nil,nil
}

// func UpdateUserBioHandler(c *gin.Context){}

func GetAllCitiesHandler(c *gin.Context){

	// c.String(200, "test")
	// res, err := ctrl.Scan("citytable2")
	// if err != nil {
	// 	c.AbortWithError(501, err)
	// }
	// fmt.Println("SCAN ALL ",res)
	c.JSONP(200, gin.H{"data": "all_cities"})

	// c.JSON(200, "TESTTEST")
}

func GetUserProfileHandler(c *gin.Context){}
