package main

import (
	"fmt"
	"os"
	"strings"
)

func PrintHtml(cred_valid bool) {
    var s strings.Builder

    s.WriteString("Content-Type: text/html; charset=utf-8\r\n\r\n")
    s.WriteString(`<!DOCTYPE html>
<html>
    <head>
        <title>upload</title>`)
    style, err := os.ReadFile(css)
    if err == nil {
        s.WriteString(`
        <style type="text/css">`)
        s.Write(style)
        s.WriteString(`
        </style>`)
    }
    s.WriteString(`
    </head>
    <body>
        <div class="box">
            <form method="post" enctype="multipart/form-data">`,)
    if cred_valid {
        s.WriteString(`
                <input type="hidden" name="user" value="`)
        s.WriteString(username)
        s.WriteString(`">
                <input type="hidden" name="pass" value="`)
        s.WriteString(password)
        s.WriteString(`">
                <h1>File Upload</h1>
                <p><input type="file" name="the_file" multiple="multiple" class="form">
                <p><div class="check-box">
                <input type="checkbox" name="overwrite" id="overwrite" class="checkbox" `)
        if overwrite { s.WriteString("checked ") }
        s.WriteString(`>
                <label for="overwrite">Overwrite</label></div>
                <input type="submit" value="Upload file" class="button">
            </form>
            <form method="post" enctype="multipart/form-data">
                <input type="hidden" name="logout">
                <input type="submit" value="Log out" class="small-button">
            </form>`)
    } else {
        s.WriteString(`
                <h1>Login</h1>
                <p><input type="text" placeholder="Username" name="user" class="field" value="`)
        s.WriteString(username)
        s.WriteString(`">
                <p><input type="password" placeholder="Password" name="pass" class="field" value="`)
        s.WriteString(password)
        s.WriteString(`">
                <p><div class="check-box"><input type="checkbox" name="remember" id="remember" class="checkbox" `)
        if remember { s.WriteString("checked ") }
        s.WriteString(`>
                <label for="remember">Remember me</label></div>
                <p><input type="submit" value="Login" class="button">
            </form>`)
    }
    s.WriteString(message.String())
    s.WriteString(`
        </div>
    </body>
</html>
`)
    fmt.Print(s.String())
}
