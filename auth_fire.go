package main

import (
	"context"
	"errors"
	"fmt"
	// "io/ioutil"
	"strings"
	"log"
	"time"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/gin-gonic/gin"
	// dynamodbclient iface
)

const (
	privKeyPath = "./keys/privkey_apr2020.pem"
	pubKeyPath  = "./keys/pubkey_apr2020.pub"
)

// used in authenticated POST endpoints
// https://firebase.google.com/docs/auth/admin/verify-id-tokens
// https://godoc.org/firebase.google.com/go/auth
// Added new error handling functions to be used in conjunction with VerifyIDToken() and VerifySessionCookie() APIs:
// auth.IsIDTokenInvalid()
// auth.IsIDTokenExpired()
func VerifyFireToken(tok string) (interface{}, error) {
	log.Println("TOKEN ", tok)

	// Access auth service from the default app
	// https://firebase.google.com/docs/auth/admin/verify-id-tokens
	fclient, err := firapp.Auth(context.Background())
	if err != nil {

		log.Println("TOKEN ERROR", err)
		// log.Fatalf("error getting Auth client: %v\n", err)
	}
	session, err := fclient.VerifyIDToken(context.Background(), tok)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", session)
	return session, nil
}

// used in authenticated GET endpoints
// protected endpoints while session is valid
// return new jwt
// https://godoc.org/github.com/dgrijalva/jwt-go#ParseRSAPublicKeyFromPEM
// https://github.com/serverless/examples/blob/master/aws-golang-auth-examples/functions/auth/main.go
// https://www.npmjs.com/package/serverless-offline#token-authorizers
// https://gist.github.com/ctalladen78/753395c5bc49de019c55f6495e901398
// https://github.com/sohamkamani/jwt-go-example
func VerifyJWTToken(inputToken string) (string, jwt.MapClaims, error) {
	// privKey, err := ioutil.ReadFile(privKeyPath)
	// if err != nil {
	// 	return nil, err
	// }
	privKey := os.Getenv("LOCAL_PRIVATE_KEY")	
	inputToken = strings.Replace(inputToken, "Bearer ", "", -1)
	// stackoverflow.com/questions/556987770/decode-jwt-without-validation-and-find-scope
	// token, _, err := new(jwt.Parser).ParseUnverified(inputToken, jwt.MapClaims{})
	// if err != nil {
	// 	return nil, err
	// }
	// if claims, ok := token.Claims.(jwt.MapClaims); ok {
	// 	fmt.Println("PARSED CLAIMS", claims["userid"])
	// }

	// TODO check token.Header["alg"]
	// TODO check token.Claims["exp"]
	// TODO check token.Claims["userid"]
	// claims := jwt.MapClaims{}
	// tokenDetails, err := jwt.ParseWithClaims(inputToken, claims, func(token *jwt.Token) (interface{}, error){
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}
	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return privKey, nil
	// })
	// fmt.Println("INPUT TOKEN DETAILS", tokenDetails)
	// fmt.Println("IS VALID TOKEN", tokenDetails.Valid)
	// for key, val := range claims {
	// 	fmt.Printf("PARSED CLAIMS %v %v \n", key, val)
	tokenDetails, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(privKey), nil
	})

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("Invalid token malformed")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("JWT expired")
		} else {
			fmt.Println("Validation error:", ve, err)
		}
	} else if tokenDetails.Valid {
		fmt.Println("Valid token")
	} else {
		fmt.Println("Invalid token details:", err)
	}

	// // return jwt.claim info
	if claims, ok := tokenDetails.Claims.(jwt.MapClaims); ok && tokenDetails.Valid {
		fmt.Println("TOKEN VALID ", tokenDetails.Valid)
		// fmt.Println("TOKEN DETAILS ", tokenDetails)
		fmt.Println("PARSED CLAIMS ", claims)

		// check if expired
		// if time.Unix(claims["exp"], 0).Sub(time.Now()) > 30*time.Second {
		// 	return nil, errors.New("EXPIRED TOKEN")
		// }

		get_tok, err := CreateGetToken("2394772470822913")
		if err != nil {
			return "", nil, err
		}
		return get_tok, claims, nil
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("SIGNATURE INVALID")
			return "",nil, err
		}
		return "",nil, err
	}

	/**
	if claims, ok := token.Claims.(*jwtCustomClaims); ok && token.Valid {
		fmt.Printf("CLAIMS %v %v", claims, claims.StandardClaims)
	}

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			fmt.Println("SIGNATURE INVALID")
			return false, err
		}
		return false, err
	}
	// check if expired
	if time.Unix(claims.StandardClaims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return false, errors.New("EXPIRED TOKEN")
	}

	return token.Valid, nil
	**/
	// return empty claims object
	return "",nil, errors.New("Invalid token")
}

// generate locally verfied token 
func CreateGetToken(userid string) (string, error) {
	// privKey, err := ioutil.ReadFile(privKeyPath)
	// if err != nil {
	// 	fmt.Println("READ ERROR ", err)
	// 	return "", nil
	// }
	// fmt.Println("PRIV KEY ", privKey)
	privKey := os.Getenv("LOCAL_PRIVATE_KEY")	
	// create jwt token
	newToken := jwt.New(jwt.SigningMethodHS256)
	aclaims := newToken.Claims.(jwt.MapClaims) // make new custom claims map
	aclaims["userid"] = userid
	// match fbToken expiration
	aclaims["exp"] = time.Now().Add(time.Hour * 168).Unix()
	signed, err := newToken.SignedString([]byte(privKey))
	if err != nil {
		return "", err
	}
	fmt.Println("NEW TOKEN ", signed)
	return signed, nil
}
