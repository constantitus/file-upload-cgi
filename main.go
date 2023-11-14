package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Should end with "/" or left blank for relative path
const userDir = ""
const storageDir = ""

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
    // fmt.Printf("<p>env:[%s]", env)
    // fmt.Printf("<p>buf:[%s]", buffer)
    // fmt.Printf("<p>boundary:[%s]", boundary)

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

    /* fmt.Printf("<p>file_name:[%s]", file_name)
    fmt.Printf("<p>file_contents:[%s]", file_contents)
    fmt.Printf("<p>username:[%s]", username)
    fmt.Printf("<p>password:[%s]", password) */

    // Credentials checking
    cred_valid := false
    if username != "" {
        cred_file, err := os.ReadFile(userDir + username + ".txt")
        if err != nil {
            fmt.Printf("<p>User not found")
        } else {
            tmp := sha256.New()
            tmp.Write([]byte(password))
            pass_hashed := hex.EncodeToString(tmp.Sum(nil))
            file_string := strings.TrimSuffix(string(cred_file), "\n")
            if pass_hashed == file_string {
                cred_valid = true
            } else {
                fmt.Printf("<p>Wrong Password")
            }
        }
    }

    // Writing
    // TODO: check file before, maybe ask the user if they want to overwrite
    // if the file doesn't exist, create and write
    if cred_valid && file_name != "" {
        file, err := os.OpenFile(
            storageDir + username + "/" + file_name,
            os.O_RDWR|os.O_CREATE,
            0644)
        filestat, _ := file.Stat()
        if err != nil {
            fmt.Printf("<p>Error: %s", err)
        }
        if filestat.Size() > 0 {
            fmt.Println("<p>The file already exists")
        } else {
            filesize, err := file.WriteString(file_contents)
            if err != nil {
                fmt.Printf("<p> %s", err)
            } else {
                fmt.Printf("<p> written %d bytes", filesize)
            }
            file.Close()
        }
    }

    // Printing
    fmt.Printf("<form method=\"post\" enctype=\"multipart/form-data\">")
    if cred_valid {
        fmt.Printf("<input type=\"hidden\" name=\"user\" value=\"%s\">", username)
        fmt.Printf("<input type=\"hidden\" name=\"pass\" value=\"%s\">", password)

        fmt.Printf("<input type=\"file\" name=\"the_file\" class=\"form\"><p>")
        fmt.Printf("<input type=\"submit\" value=\"Upload file\" class=\"button\"></form></body></html>")
    } else {
        fmt.Printf("<p>Username <input type=\"text\" name=\"user\" value=\"%s\" class=\"field\"><p>", username)
        fmt.Printf("<p>Password <input type=\"password\" name=\"pass\" value=\"%s\" class=\"field\"><p>", password)
        fmt.Printf("<input type=\"submit\" value=\"Login\" class=\"button\"></form></body></html>")
    }

    fmt.Printf("</div>")
    fmt.Printf("</body></html>")
}
