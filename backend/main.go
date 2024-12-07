package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Options struct {
	Address         string
	Port            int
	PrivateKeyPath  string
	CertificatePath string
}

func parseFlags() Options {
	var o Options

	flag.StringVar(&o.Address, "a", "127.0.0.1", "Address to listen on")
	flag.IntVar(&o.Port, "p", 8443, "Port to listen on")
	flag.StringVar(&o.PrivateKeyPath, "k", "privatekey.pem", "Path of private-key file")
	flag.StringVar(&o.CertificatePath, "c", "certificate.pem", "Path of certificate")

	flag.Parse()

	return o
}

func main() {
	o := parseFlags()

	router := gin.Default()

	router.Static("/home", "./webroot")

	router.RunTLS(fmt.Sprintf("%s:%d", o.Address, o.Port), o.CertificatePath, o.PrivateKeyPath)
}
