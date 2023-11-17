package main

import (
	"fmt"
	"os"
	"strings"
)

func CheckCookies() {
    cookie := strings.Split(os.Getenv("HTTP_COOKIE"), "; ")
    for _, v := range cookie {
        tmp_user, found_user := strings.CutPrefix(v, "username=")
        if found_user { cookies.user = tmp_user }
        tmp_pass, found_pass := strings.CutPrefix(v, "passhash=")
        if found_pass { cookies.pass = tmp_pass }
    }
}

func DeleteCookies() {
    fmt.Print("Set-Cookie: username=; expires=Thu, 01 Jan 1970 00:00:00 GMT\n")
    fmt.Print("Set-Cookie: passhash=; expires=Thu, 01 Jan 1970 00:00:00 GMT\n")
}
