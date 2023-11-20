package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintHtml(cred_valid bool) {
    var page strings.Builder

    page.WriteString("Content-Type: text/html; charset=utf-8\r\n\r\n")
    page.WriteString(`<!DOCTYPE html>
<html>
    <head>
        <title>upload</title>`)
    style, err := os.ReadFile(css)
    if err == nil {
        page.WriteString(`
        <style type="text/css">`)
        page.Write(style)
        page.WriteString(`
        </style>`)
    }
    page.WriteString(`
    </head>
    <body>
        <div class="box">
            <form method="post" enctype="multipart/form-data">`,)
    if cred_valid {
        page.WriteString(`
                <input type="hidden" name="user" value="`)
        page.WriteString(username)
        page.WriteString(`">
                <input type="hidden" name="pass" value="`)
        page.WriteString(password)
        page.WriteString(`">
                <h1>File Upload</h1>
                <p><input type="file" name="the_file" multiple="multiple" class="form">
                <p><div class="check-box">
                <input type="checkbox" name="overwrite" id="overwrite" class="checkbox" `)
        if overwrite { page.WriteString("checked ") }
        page.WriteString(`>
                <label for="overwrite">Overwrite</label></div>
                <input type="submit" value="Upload file" class="button">
            </form>
            <form method="post" enctype="multipart/form-data">
                <input type="hidden" name="logout">
                <input type="submit" value="Log out" class="small-button">
            </form>`)
    } else {
        page.WriteString(`
                <h1>Login</h1>
                <p><input type="text" placeholder="Username" name="user" class="field" value="`)
        page.WriteString(username)
        page.WriteString(`">
                <p><input type="password" placeholder="Password" name="pass" class="field" value="`)
        page.WriteString(password)
        page.WriteString(`">
                <p><div class="check-box"><input type="checkbox" name="remember" id="remember" class="checkbox" `)
        if remember { page.WriteString("checked ") }
        page.WriteString(`>
                <label for="remember">Remember me</label></div>
                <p><input type="submit" value="Login" class="button">
            </form>`)
    }
    page.WriteString(message.String())
    page.WriteString(`
        </div>
    </body>
</html>
`)
    fmt.Print(page.String())
}
