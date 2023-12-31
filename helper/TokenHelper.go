package helper

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhanarrizky/go-blog/database"
	"github.com/dhanarrizky/go-blog/models"
)

var SECRET_KEY string = os.Getenv("SECRET_KEY")

type SignedDetailes struct {
	UserName  string
	Email     string
	FirstName string
	LastName  string
	Uid       string
	Role      string
	jwt.StandardClaims
}

func GenerateJwtToken(username, email, firstname, lastname, uid, role string) (string, string, error) {
	claims := &SignedDetailes{
		UserName:  username,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
		Uid:       uid,
		Role:      role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(3)).Unix(),
		},
	}

	refreshClaims := &SignedDetailes{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("error here")
		log.Println(err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Println("and error here")
		log.Println(err)
	}

	return token, refreshToken, err
}

var DB = database.ConDB()

func UpdateJwtToken(signedToken, refreshSignedToken, uid string) {
	db := DB.Begin()

	var users models.User

	uidInt, _ := strconv.Atoi(uid)
	db.Find(&users, uidInt)

	users.Token = &signedToken
	users.RefreshToken = &refreshSignedToken

	DB.Save(&users)

	if db.Error != nil {
		log.Println(db.Error.Error())
		return
	}

	defer db.Rollback()

	if db.RowsAffected > 0 {
		log.Println("updated jwt token has been successfully")
		return
	}
	return
}

func ValidateJwtToken(tokens string) (claims *SignedDetailes, msg string) {
	token, err := jwt.ParseWithClaims(
		tokens,
		&SignedDetailes{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetailes)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}

	return claims, msg
}
