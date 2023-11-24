package main

import (
	"RobinStock/login"
	"fmt"
)

func main() {

	options := &login.LoginOptions{
		Username: "ianinweze@gmail.com",
		Password: "Nigeria1998",
	}

	resp, err := login.Login(*options)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}
