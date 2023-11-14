package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    // web header
    fmt.Printf("Content-Type: text/html; charset=utf-8\r\n\r\n")
    // TODO: css
    fmt.Printf("<body><div class=\"frame box\">")

    var buffer []byte
    env := os.Getenv("CONTENT_TYPE")
    boundary, uploading := strings.CutPrefix(env, "multipart/form-data; boundary=")
    if uploading {
        reader := bufio.NewReader(os.Stdin)
        var err error = nil
        for err == nil {
            var tmp byte
            tmp, err = reader.ReadByte()
            buffer = append(buffer, tmp)
        }
    }
    fmt.Printf("<p>env:[%s]", env)
    fmt.Printf("<p>buf:[%s]", buffer)
    //fmt.Printf("<p>boundary:[%s]", boundary)

    var file_name string
    var file_contents string
    var username string
    var password string

    // Formatting
    input := strings.Split(string(buffer), boundary)
    for _, v := range input {
        cut_status := false
        v, cut_status = strings.CutPrefix(v, "\r\nContent-Disposition: form-data; ")
        if !cut_status {
            continue
        }

        _, tmp, found_name := strings.Cut(v, "filename=\"")
        if found_name {
            file_name, _, _ = strings.Cut(tmp, "\"")
            _, tmp, _ = strings.Cut(v, "\r\n\r\n")
            file_contents, _ = strings.CutSuffix(tmp, "\r\n--")
            continue
        }

        tmp_user, user_status := strings.CutPrefix(v, "name=\"user\"\r\n\r\n")
        if user_status {
            username, _ = strings.CutSuffix(tmp_user, "\r\n--")
            continue
        }
        tmp_pass, pass_status := strings.CutPrefix(v, "name=\"pass\"\r\n\r\n")
        if pass_status {
            password, _ = strings.CutSuffix(tmp_pass, "\r\n--")
            continue
        }
    }

    fmt.Printf("<p>file_name:[%s]", file_name)
    fmt.Printf("<p>file_contents:[%s]", file_contents)
    fmt.Printf("<p>username:[%s]", username)
    fmt.Printf("<p>password:[%s]", password)

    // Writing
    // TODO: check file before, maybe ask the user if they want to overwrite
    // if the file doesn't exist, create and write
    if uploading {
        file, err := os.OpenFile(file_name, os.O_RDWR|os.O_CREATE, 0644)
        filestat, _ := file.Stat()
        if err != nil {
            fmt.Printf("<p>Error: %s", err)
        }
        if filestat.Size() > 0 {
            fmt.Println("<p>The file already exists")
        } else {
            filesize, err := file.WriteString(username + "\n" + password)
            if err != nil {
                fmt.Printf("<p> %s", err)
            } else {
                fmt.Printf("<p> written %d bytes", filesize)
            }
            file.Close()
        }
    }

    // Printing
    fmt.Printf("<form method=\"post\" enctype=\"multipart/form-data\"><input type=\"file\" name=\"the_file\" class=\"form\"><p>")
    // check for password
    // TODO: make users and store both usernames and passwords
    if file_contents != "" { // TODO: check the proper thing
        fmt.Printf("<input type=\"hidden\" name=\"user\" value=\"%s\">", password)
        fmt.Printf("<input type=\"hidden\" name=\"pass\" value=\"%s\">", password)
    } else {
        fmt.Printf("Username <input type=\"text\" name=\"user\" class=\"field\"><p>")
        fmt.Printf("Password <input type=\"password\" name=\"pass\" class=\"field\"><p>")
    }
    fmt.Printf("<input type=\"submit\" value=\"Upload file\" class=\"upload\"></form></body></html>")

    fmt.Printf("</div>")
    fmt.Printf("</body></html>")
}
