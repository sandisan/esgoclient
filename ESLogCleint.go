package main

import (
	"io"
	"net/http"
	"crypto/tls"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":8080", nil)
}
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	
    client := &http.Client{Transport: tr}
    resp, err := client.Get("https://elasticsearch.ibm-common-services.svc:9200/_cluster/health")
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
