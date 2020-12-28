package main

import (
	"io"
	"fmt"
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

// 	cert, err := ioutil.ReadFile("elk.crt")
// 	if err != nil {
// 		log.Fatalf("Couldn't load file", err)
// 	}
// 	certPool := x509.NewCertPool()
// 	certPool.AppendCertsFromPEM(cert)

// 	conf := &tls.Config{
// 		RootCAs: certPool,
// 		//InsecureSkipVerify: true,
// 	}
// //     tr := &http.Transport{
// //         TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// //     }
//   tr := &http.Transport{
//         TLSClientConfig: conf,
//     }
//      caCert, err := ioutil.ReadFile("elk-ca.pem")
// 	caCert, err := ioutil.ReadFile("elk.crt")
//       if err != nil {
//               log.Fatal(err)
// 	      io.WriteString(w, err.Error())
// 	      return
//        }
//        caCertPool := x509.NewCertPool()
//        caCertPool.AppendCertsFromPEM(caCert)
// //       cert, err := tls.LoadX509KeyPair("elk-crt.pem", "elk-key.pem")
// //        if err != nil {
// //                log.Fatal(err)
// // 	       io.WriteString(w, err.Error())
// // 	      return
// //        }

// //         client := &http.Client{
// //                 Transport: &http.Transport{
// //                         TLSClientConfig: &tls.Config{
// //                                 RootCAs:      caCertPool,
// //                                Certificates: []tls.Certificate{cert},
// //                         },
// //                 },
// //         }
//     //client := &http.Client{Transport: tr}
// 	client := &http.Client{
//         Transport: &http.Transport{
//             TLSClientConfig: &tls.Config{
//                 RootCAs:      caCertPool,
//             },
//         },
//     }
	
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
