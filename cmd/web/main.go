package main

import (
	"fmt"
	"net/http"
	
)

const port = ":8080"


func main()  {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/contact", Contact)


	fmt.Println("(http://localhost:8080) - Server started on port", port)
	http.ListenAndServe(port, nil)

	mygoprogram.Config()
}
