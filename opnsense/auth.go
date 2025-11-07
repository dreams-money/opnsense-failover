package opnsense

import "encoding/base64"

var Authorization string

func SetAuthorization(key, secret string) {
	auth := key + ":" + secret
	auth = base64.StdEncoding.EncodeToString([]byte(auth))

	Authorization = "Basic " + auth
}
