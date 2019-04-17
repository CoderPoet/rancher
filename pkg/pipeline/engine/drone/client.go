package drone

import "net/http"

type Client struct {
	API   string
	User  string
	Token string

	HTTPClient *http.Client
}
