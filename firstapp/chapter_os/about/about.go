package about

import (
	"fmt"
)

const (
	About   = "This is a simple web application written in Go."
	Version = "0.0.1"
	Author  = "Kornilov LN"
	Email   = "ln.kornilovstar@gmail.com"
	Github  = "https://github.com/KornilovLN"
	Date    = "2024-10-02"
)

func AboutInfo() {
	fmt.Println("--- Info about program ----------------------------------")
	fmt.Println("About:   ", About)
	fmt.Println("Version: ", Version)
	fmt.Println("Author:  ", Author)
	fmt.Println("Email:   ", Email)
	fmt.Println("Github:  ", Github)
	fmt.Println("Date:    ", Date)
	fmt.Println("---------------------------------------------------------")
}
