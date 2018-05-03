package util

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	// Go MySQL Driver is a MySQL driver for Go's (golang) database/sql package
	_ "github.com/go-sql-driver/mysql"
)

// A minimal version of Http response
type HttpResponse struct {
	StatusCode int
	Header     http.Header
	Body       string
}

// Getenv retrieves the value of the environment variable named by the key
//
func Getenv(key string) string {
	return os.Getenv(key)
}

// JSONDecode parses the JSON-encoded data and stores the result in v
//
func JSONDecode(str string, v interface{}) error {
	return json.Unmarshal([]byte(str), &v)
}

// B64EncodeToString returns the base64 encoding of src.
//
func B64EncodeToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// B64DecodeString returns the bytes represented by the base64 string s.
//
func B64DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// KmsEncrypt uses KMS to Encrypt a string
//
func KmsEncrypt(text string) (string, error) {
	// // Initialize a session in us-west-2 that the SDK will use to load
	// // credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)
	if err != nil {
		fmt.Println("Got error setting up the session: ", err)
		return "", err
	}
	// Create KMS service client
	svc := kms.New(sess)

	// Encrypt data key
	keyID := "ada987df-3c2f-4305-bd9f-eaff3e773e09"
	// Encrypt the data key
	result, err := svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(keyID),
		Plaintext: []byte(text),
	})

	if err != nil {
		fmt.Println("Got error encrypting data: ", err)
		return "", err
	}

	return B64EncodeToString(result.CiphertextBlob), nil
}

// KmsDecrypt uses KMS to Decrypt a token
//
func KmsDecrypt(s string) (string, error) {
	blob, err := B64DecodeString(s)
	if err != nil {
		fmt.Println("Got error decoding the input String: ", err)
		return "", err
	}
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)
	if err != nil {
		fmt.Println("Got error setting up the session: ", err)
		return "", err
	}

	// Create KMS service client
	kmsClient := kms.New(sess)
	// Decrypt the data
	result, err := kmsClient.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})
	if err != nil {
		fmt.Println("Got error decrypting data: ", err)
		return "", err
	}

	return string(result.Plaintext), nil
}

// LoadConfig gets app configouration
//
func LoadConfig(path string, v interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&v)
	if err != nil {
		return err
	}

	file.Close()

	return nil
}

// OpenDB opens a DB connection
//
func OpenDB(user string, password string, host string, name string) (*sql.DB, error) {
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, name)
	db, err := sql.Open("mysql", source)

	// defer the close till after the main function has finished
	// executing
	//defer db.Close()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// HttpGet issues a GET request to the specified URL. If the response is a
// redirect codes, HttpGet follows the redirect, up to a maximum of 10 redirects
//
func HttpGet(url string) (*HttpResponse, error) {
	// Issue a GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// Load the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Compose the httpResponse
	httpResponse := HttpResponse{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(data),
	}

	return &httpResponse, nil
}

// HttpPost issues a POST to the specified URL.
//
func HttpPost(url string, body interface{}) (*HttpResponse, error) {
	jsonValue, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Issue a Post request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	// Load the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Compose the httpResponse
	httpResponse := HttpResponse{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       string(data),
	}

	return &httpResponse, nil
}
