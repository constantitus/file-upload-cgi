package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func CheckCredentials() bool {
    if username != "" {
        return Check(username, password, false)
    }
    if cookies.user != "" {
        return Check(cookies.user, cookies.pass, true)
    }
    return false
}

func Check(user string, pass string, hashed bool) bool {
    cred_file, err := os.ReadFile(UserDir + user + ".txt")
    if err == nil {
        file_string := strings.TrimSuffix(string(cred_file), "\n")
        if hashed && pass == file_string {
            return true
        } else {
            tmp := sha256.New()
            tmp.Write([]byte(pass))
            pass = hex.EncodeToString(tmp.Sum(nil))
            if pass == file_string {
                if remember {
                    fmt.Printf("Set-Cookie: username=%s\n", user)
                    fmt.Printf("Set-Cookie: passhash=%s\n", pass)
                }
                return true
            }
        }
    }
    return false
}
