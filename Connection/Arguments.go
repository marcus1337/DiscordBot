package Connection;

import(
	"flag"
)

var (
	token string
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

func getToken() string{
	return token
}