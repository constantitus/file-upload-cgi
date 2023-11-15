package main

import (
	"fmt"
	"os"
)

func PrintHeader() {
    // web header
    fmt.Printf("Content-Type: text/html; charset=utf-8\r\n\r\n")
    fmt.Printf("<!DOCTYPE html><html><head><title>upload</title>")
    style, err := os.ReadFile(css)
    if err == nil {
        fmt.Printf("<style type=\"text/css\">%s</style>", string(style))
    }

    fmt.Printf("<body><div class=\"box\">")
}

func PrintFooter() {
    fmt.Print("<form method=\"post\" enctype=\"multipart/form-data\">")
    if cred_valid {
        fmt.Printf("<input type=\"hidden\" name=\"user\" value=\"%s\">", username)
        fmt.Printf("<input type=\"hidden\" name=\"pass\" value=\"%s\">", password)

        fmt.Print("<h1>File Upload</h1>")
        fmt.Print("<input type=\"file\" name=\"the_file\" multiple=\"multiple\" class=\"form\"><p>")
        fmt.Print("<input type=\"submit\" value=\"Upload file\" class=\"button\"></form></body></html>")
    } else {
        fmt.Printf("<h1>Login</h1>")
        fmt.Printf("<input type=\"text\" placeholder=\"Username\" name=\"user\" value=\"%s\" class=\"field\"><p>", username)
        fmt.Printf("<input type=\"password\" placeholder=\"Password\" name=\"pass\" value=\"%s\" class=\"field\"><p>", password)
        fmt.Print("<input type=\"submit\" value=\"Login\" class=\"button\"></form></body></html>")
    }
    fmt.Print(message)
    fmt.Print("</div>")

    // Testing
    // fmt.Printf("<p>env:[%s]", env)
    // fmt.Printf("<p>boundary:[%s]", boundary)
    // fmt.Printf("<p>buf:[%s]", buffer)
    /* for _, file := range files {
        fmt.Printf("<p>file_name:[%s]", file.name)
        fmt.Printf("<p>file_contents:[%s]", file.content)
    } */
    // fmt.Printf("<p>username:[%s]", username)
    // fmt.Printf("<p>password:[%s]", password)

    fmt.Print("</body></html>")
}
