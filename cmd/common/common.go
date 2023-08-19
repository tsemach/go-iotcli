package common

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	// "net/http"
	"os"
	// "time"

	"github.com/tsemach/go-iotcli/config"
)

func GetEnv(env string) string {
	if env == "" {
		env = os.Getenv("IOTCLI_ENV")
	}

	switch env {
	case "dev":
		return env
	case "qa":
		return env
	case "stage":
		return env
	case "prod":
		return env
	}

	panic("env is not given, use --env flag or IOTCLI_ENV environment variable")
}

func GetRootCAs(cacertPath string) *x509.CertPool {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// read in the cert file
	caCert, err := os.ReadFile(cacertPath)
	if err != nil {
		log.Fatalf("failed to append %q to RootCAs: %v", "ca.crt", err)

		return rootCAs
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(caCert); !ok {
		log.Println("no certs appended, using system certs only")
	}

	return rootCAs
}

func GetClientPair(env string) (*tls.Certificate, error) {
	clientCrtPath, clientKeyPath := config.GetClientCert(env)

	cert, err := tls.LoadX509KeyPair(clientCrtPath, clientKeyPath)
	if err != nil {
		log.Fatalf("Error opening cert file %s and key %s, Error: %s", clientCrtPath, clientKeyPath, err)
	}

	return &cert, err
}

// func GetClient(env string) *http.Client {
// 	rootCAs := GetRootCAs(config.GetCAPath(env))
// 	cert, err := GetClientPair(env)

// 	if err != nil {
// 		log.Fatalf("error on getClient, err: %s", err)
// 	}

// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{
// 			ServerName:         "localhost",
// 			InsecureSkipVerify: true,
// 			Certificates:       []tls.Certificate{*cert},
// 			RootCAs:            rootCAs,
// 		},
// 	}

// 	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

// 	return client
// }

// func SendPost[REQ any, RES any](env string, path string, body *REQ) (*http.Response, *RES) {

// 	client := GetClient(env)

// 	postBody, _ := json.Marshal(*body)
// 	responseBody := bytes.NewBuffer(postBody)
// 	resp, err := client.Post(fmt.Sprintf("%s%s", config.GetEnvDomain(env), path), "application/json", responseBody)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()

// 	var ir RES
// 	err = json.NewDecoder(resp.Body).Decode(&ir)

// 	return resp, &ir
// }

func JsonPrettyEncode(data interface{}, out io.Writer) error {
	enc := json.NewEncoder(out)
	enc.SetIndent("", "    ")
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func JsonPrettyPrint(j any) {
	var buffer bytes.Buffer

	err := JsonPrettyEncode(j, &buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buffer.String())
}

func First[T, U any](val T, _ U) T {
	return val
}

// exists returns whether the given file or directory exists
func Exist(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { 
			return true, nil 
		}
    if os.IsNotExist(err) { 
			return false, nil 
		}
    return false, err
}