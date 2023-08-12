package common

import (
	"crypto/x509"
	"log"
	"os"
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