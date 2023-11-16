package main

import (
	"bufio"
	"os"
	"strings"
)

const css = "styles.css"

// Should end with "/" or left blank for relative path
const UserDir = ""
const StorageDir = ""

type file_t struct {
    name string
    content string
}

func new_file(name string, content string) (file file_t) {
    file.name = name
    file.content = content
    return
}

var (
    files []file_t
    username string
    password string
    message string
    remember bool
    overwrite bool
    cookies struct {
        user string
        pass string
    }
)

func main() {
    var buffer []byte
    content_type := os.Getenv("CONTENT_TYPE")
    boundary, is_reading := strings.CutPrefix(content_type, "multipart/form-data; boundary=")
    if is_reading {
        reader := bufio.NewReader(os.Stdin)
        var err error = nil
        for err == nil {
            var tmp byte
            tmp, err = reader.ReadByte()
            buffer = append(buffer, tmp)
        }
    }

    http_cookie := strings.Split(os.Getenv("HTTP_COOKIE"), "; ")
    for _, v := range http_cookie {
        tmp_user, found_user := strings.CutPrefix(v, "username=")
        if found_user { cookies.user = tmp_user }
        tmp_pass, found_pass := strings.CutPrefix(v, "passhash=")
        if found_pass { cookies.pass = tmp_pass }
    }

    ParseBuffer(buffer, boundary)
    cred_valid := CheckCredentials()

    if cred_valid {
        WriteFiles()
    }

    PrintHtml(cred_valid)
}
