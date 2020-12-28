package main

import (
	"io"
	"fmt"
	"flag"
	"log"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	_ "github.com/nvn1729/badimportdemo/dontimportme"
)

var (
	certFile = flag.String("cert", "elk-crt.pem", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", "elk-key.pem", "A PEM encoded private key file.")
	caFile   = flag.String("CA", "elk-ca.pem", "A PEM eoncoded CA's certificate file.")
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
		io.WriteString(w, err.Error())
		return
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
		io.WriteString(w, err.Error())
		return
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	resp, err := client.Get("https://elasticsearch:9200/_cluster/health")
	if err != nil {
		fmt.Println(err)
		io.WriteString(w, err.Error())
		return
	}else {
	    // read all response body
	    data, err := ioutil.ReadAll( resp.Body )
	    if err != nil {
		fmt.Println(err)
		io.WriteString(w, err.Error())
		return
	    } else {
		 // close response body
		resp.Body.Close()
		io.WriteString(w, string(data))  
	    }
	}
	
}
