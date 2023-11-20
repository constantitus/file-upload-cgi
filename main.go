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
    contents string
}

func new_file(name string, content string) (file file_t) {
    file.name = name
    file.contents = content
    return
}

var (
    files []file_t
    username string
    password string
    cookies struct {
        user string
        pass string
    }
    message strings.Builder
    remember bool
    overwrite bool
    logout bool
    creds_valid bool
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
        creds_valid = CheckCredentials()

        if creds_valid {
            WriteFiles()
        }
    }

    PrintHtml(creds_valid)
}
