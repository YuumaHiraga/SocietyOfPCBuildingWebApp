package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/rs/cors"
)

// Most of the code for this package is from here https://medium.com/@pkbhowmick007/user-registration-and-login-template-using-golang-mongodb-and-jwt-d85f09f1295e

var SECRET_KEY = []byte("gosecretkey")

type User struct {
	gorm.Model
	FirstName string `json:"firstname" gorm:"firstname"`
	LastName  string `json:"lastname" gorm:"lastname"`
	Email     string `json:"email" gorm:"primaryKey" gorm:"uniqueIndex"`
	Password  string `json:"password" gorm:"password"`
}

// Takes in password, returns a hash
func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// This does something???
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Println("Error in JWT token generation")
		return "", err
	}
	return tokenString, nil
}

// This function registers a new user
func userRegister(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User
	var hold User //Just need an empty instance of a user struct
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//Check to see if there is already a user associated with the given email address
	if err := db.Where("Email = ?", user.Email).First(&hold).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&user)
			response.Write([]byte(`User registered`))
			return
		} else {
			panic("something terrible has happened")
		}
	} else {
		response.Write([]byte(`Email is already in use`))
		return
	}
}

// This function logs a user in
func userLogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user User
	var dbUser User
	json.NewDecoder(request.Body).Decode(&user)
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Check to see if the user exists
	if err := db.Where("Email = ?", user.Email).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Write([]byte(`No user with that email exists`))
			return
		} else {
			panic("something terrible has happened")
		}
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)

	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)

	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	jwtToken, err := GenerateJWT()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))
	response.Write([]byte(`{Successful}`))
}

// A simple little api that just exists for testing purposes
func test(response http.ResponseWriter, request *http.Request) {
	fmt.Print("Test success\n")
	response.Write([]byte(`Test success`))
}

func main() {
	db, err := gorm.Open(sqlite.Open("SPCB.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	router := mux.NewRouter()
	router.HandleFunc("/api/signup", userRegister).Methods("POST")
	router.HandleFunc("/api/login", userLogin).Methods("POST")
	router.HandleFunc("/api/test", test).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe("localhost:4200", handler))

	err = http.ListenAndServe("localhost:4200", handler)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}
