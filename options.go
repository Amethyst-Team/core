package main

// import (
// 	"os"
// )

type IOptions struct {
	LISTEN_ADDRESS string
	LISTEN_PORT    string
}

var Options = IOptions{
	LISTEN_ADDRESS: "0.0.0.0", //os.Getenv("LISTEN_ADDRESS"),
	LISTEN_PORT:    "8000",    //os.Getenv("LISTEN_PORT"),
}
