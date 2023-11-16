package main

import (
	"fmt"
	"os"
)

func PrintHtml(cred_valid bool) {
    // web header
    fmt.Printf("Content-Type: text/html; charset=utf-8\r\n\r\n")
    fmt.Printf(`<!DOCTYPE html><html><head><title>upload</title>`)
    style, err := os.ReadFile(css)
    if err == nil {
        fmt.Printf(`<style type="text/css">%s</style>`, string(style))
    }

    fmt.Printf(`<body><div class="box">`)
    fmt.Print(`<form method="post" enctype="multipart/form-data">`)
    if cred_valid {
        fmt.Printf(`<input type="hidden" name="user" value="%s">`, username)
        fmt.Printf(`<input type="hidden" name="pass" value="%s">`, password)

        fmt.Print(`<h1>File Upload</h1>`)
        fmt.Print(`<input type="file" name="the_file" multiple="multiple" class="form"><p>`)
        overwrite_box := `<p><div class="check-box"><input type="checkbox" name="overwrite" id="overwrite" class="checkbox" `
        if overwrite { overwrite_box += "checked " }
        overwrite_box += `><label for="overwrite">Overwrite</label></div>`
        fmt.Print(overwrite_box)
        fmt.Print(`<input type="submit" value="Upload file" class="button"></form>`)
        // fmt.Print(`<input type="submit" name="logout" value="Log out" class="small-button">`)
    } else {
        fmt.Printf(`<h1>Login</h1>`)
        fmt.Printf(`<p><input type="text" placeholder="Username" name="user" value="%s" class="field">`, username)
        fmt.Printf(`<p><input type="password" placeholder="Password" name="pass" value="%s" class="field">`, password)
        remember_box := `<p><div class="check-box"><input type="checkbox" name="remember" id="remember" class="checkbox" `
        if remember { remember_box += "checked " }
        remember_box += `><label for="remember">Remember me</label></div>`
        fmt.Print(remember_box)
        fmt.Print(`<p><input type="submit" value="Login" class="button"></form>`)
    }
    fmt.Print(message)
    fmt.Print(`</div>`)

    fmt.Print(`</body></html>`)
}
