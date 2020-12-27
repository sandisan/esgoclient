package main

import (
	"io"
	"fmt"
	"log"
	"net/http"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
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
 cert, err := tls.LoadX509KeyPair("elk.crt", "elk.key")
       if err != nil {
               log.Fatal(err)
       }

        client := &http.Client{
                Transport: &http.Transport{
                        TLSClientConfig: &tls.Config{
                                RootCAs:      caCertPool,
                               Certificates: []tls.Certificate{cert},
                        },
                },
        }
    //client := &http.Client{Transport: tr}
    resp, err := client.Get("https://elasticsearch:9200/_cluster/health")
    if err != nil {
	fmt.Println(err)
	io.WriteString(w, err.Error())
    }else {
	    // read all response body
	    data, _ := ioutil.ReadAll( resp.Body )
	    // close response body
	    resp.Body.Close()
	    io.WriteString(w, string(data))
    }
	
}
