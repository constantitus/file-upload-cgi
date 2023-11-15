package upload

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const css = "styles.css"

// Should end with "/" or left blank for relative path
const userDir = ""
const storageDir = ""

type file_t struct {
    name string
    content string
}

func new_file(name string, content string) (file file_t) {
    file.name = name
    file.content = content
    return
}

func main() {
    // web header
    fmt.Printf("Content-Type: text/html; charset=utf-8\r\n\r\n")
    fmt.Printf("<!DOCTYPE html><html><head><title>upload</title>")
    style, err := os.ReadFile(css)
    if err == nil {
        fmt.Printf("<style type=\"text/css\">%s</style>", string(style))
    }


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

    var files []file_t
    var username string
    var password string

    // Formatting
    input := strings.Split(string(buffer), boundary)
    for _, v := range input {
        var cut_status bool
        v, cut_status = strings.CutPrefix(v, "\r\nContent-Disposition: form-data; ")
        if !cut_status {
            continue
        }

        _, tmp, found_name := strings.Cut(v, "filename=\"")
        if found_name {
            tmp_name, _, _ := strings.Cut(tmp, "\"")
            _, tmp, _ = strings.Cut(v, "\r\n\r\n")
            tmp_content, _ := strings.CutSuffix(tmp, "\r\n--")
            files = append(files, new_file(tmp_name, tmp_content))
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

    fmt.Printf("<body><div class=\"box\">")
    var message string

    // Credentials checking
    cred_valid := false
    if username != "" {
        cred_file, err := os.ReadFile(userDir + username + ".txt")
        if err != nil {
            // Username not found
            message += "<p>Wrong username/password"
        } else {
            tmp := sha256.New()
            tmp.Write([]byte(password))
            pass_hashed := hex.EncodeToString(tmp.Sum(nil))
            file_string := strings.TrimSuffix(string(cred_file), "\n")
            if pass_hashed == file_string {
                cred_valid = true
            } else {
                // Wrong password
                message += "<p>Wrong username/password"
            }
        }
    } else {
        //blank the password when there's no username
        password = ""
    }

    // Writing
    if cred_valid {
        for _, file := range files {
            if file.name == "" {
                message += "<p>No file selected"
                continue
            }
            out, err := os.OpenFile(
                storageDir + username + "/" + file.name,
                os.O_RDWR|os.O_CREATE,
                0644)
            out_stat, _ := out.Stat()
            if err != nil {
                message += fmt.Sprintf("<p>Error: %s", err)
                continue
            }
            if out_stat.Size() > 0 {
                message += fmt.Sprintf("<p>File already exists: %s", file.name)
                // TODO: maybe ask the user if they want to overwrite
            } else {
                out_size, err := out.WriteString(file.content)
                if err != nil {
                    message += fmt.Sprintf("<p>%s", err)
                } else {
                    message += fmt.Sprintf("<p>Written %s (%d bytes)", file.name, out_size)
                }
                out.Close()
            }
        }
    }

    // Printing
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

    // DEBUG
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
