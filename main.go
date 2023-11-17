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
    message strings.Builder
    remember bool
    overwrite bool
    logout bool
    cred_valid bool
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

    ParseBuffer(buffer, boundary)
    if logout {
        DeleteCookies()
    } else {
        CheckCookies()
        cred_valid = CheckCredentials()

        if cred_valid {
            WriteFiles()
        }
    }

    PrintHtml(cred_valid)
}
