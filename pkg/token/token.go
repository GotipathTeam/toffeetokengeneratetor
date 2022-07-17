package token

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
	"time"
)

//CDN secret key generated from gotipath portal.
var secretKey string = "npkj0qkaczlkapq5uuzr2yh1cftut4zdz8o6ifb0dff4xq4vh0comb82tdt506fh"

//CDN Baser Url
var baseUrl string = "https://tm.gpcdn.net"

func generate20CharHmacSha1(resource string) string {
	hash := hmac.New(sha1.New, []byte(secretKey))
	io.WriteString(hash, resource)
	token := fmt.Sprintf("%x", hash.Sum(nil))

	return token[0:20]
}

func GenerateFinalUrl(path string) string {
	fmt.Println(time.Now().UTC())

	stime := time.Now().UTC().Format("20060102150405") //stime Start time (not valid before this time) e.g. 20120424115300 datetime format UTC

	etime := time.Now().UTC().Add(time.Hour).Format("20060102150405") //etime End time (not valid after this time)

	resource := fmt.Sprintf("%s?stime=%s&etime=%s", path, stime, etime)

	token := generate20CharHmacSha1(resource)

	println(stime)
	println(etime)

	url := fmt.Sprintf("%s%s&encoded=0%s", baseUrl, resource, token)

	return url

}
