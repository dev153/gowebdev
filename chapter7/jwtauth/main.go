package main

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// using asymmetric crypto/RSA keys
// location of the files used for signing and verification

const (
	privateKeyPath = "keys/app.rsa"     // private RSA key
	publicKeyPath  = "keys/app.rsa.pub" // public RSA key
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AppClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token"`
}

type Response struct {
	Text string `json:"text"`
}

// init() method is always run before entering the main function.
func init() {
	var err error
	// loading the private key
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("[init] %s\n", err.Error())
		return
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}

	// loading the public key
	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatalf("[init] %s\n", err.Error())
		return
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

func GenerateJWT(username, role string) (string, error) {
	claims := AppClaims{
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	fmt.Println(token)

	ss, err := token.SignedString(signKey)

	if err != nil {
		return "", err
	}
	return ss, nil
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func TokenFromAuthHeader(r *http.Request) (string, error) {
	tmp := r.Header.Get("Authorization")
	var token string
	if tmp != "" {
		if len(tmp) > 6 && true == strings.HasPrefix(tmp, "Bearer") {
			token = tmp[7:]
			return token, nil
		}
	}
	return token, errors.New("No token in the HTTP request")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error in request body.")
		return
	}

	if user.UserName != "nasos.thomos" && user.Password != "Intralot1" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Username and/or Password are incorrect.")
		return
	}

	tokenString, err := GenerateJWT(user.UserName, "Member")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while signing token!")
		log.Printf("Token Signing error: %v\n", err)
		return
	}
	response := Token{tokenString}
	jsonResponse(response, w)
}

func authHandler(w http.ResponseWriter, r *http.Request) {

	tokenInHeader, err := TokenFromAuthHeader(r)

	token, err := jwt.Parse(tokenInHeader, func(*jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)

			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Fprintln(w, "Token expired. Get a new one.")
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, "Error while parsing token.")
				log.Printf("ValidationError error: %+v\n", vErr.Errors)
				return
			}
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Error while parsing token.")
			log.Printf("Token parse error: %v\n", err)
		}
	}
	var response Response
	if token.Valid {
		response = Response{"Authorized to the system"}
	} else {
		response = Response{"Invalid token"}
	}
	jsonResponse(response, w)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/auth", authHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":9099",
		Handler: router,
	}
	log.Println("Listening")
	server.ListenAndServe()
}
