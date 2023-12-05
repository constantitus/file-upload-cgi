package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func CheckCredentials() {
    if username != "" {
        creds_valid = Check(username, password, false)
    }
    if cookies.user != "" {
        creds_valid = Check(cookies.user, cookies.pass, true)
    }
}

func Check(user string, pass string, cookie bool) bool {
    cred_file, err := os.ReadFile(UserDir + user + ".txt")
    if err == nil {
        file_string := strings.TrimSuffix(string(cred_file), "\n")
        if cookie && pass == file_string {
            return true
        } else {
            tmp := sha256.New()
            tmp.Write([]byte(pass))
            hashed_pass := hex.EncodeToString(tmp.Sum(nil))
            if hashed_pass == file_string {
                if remember {
                    fmt.Printf("Set-Cookie: username=%s\n", user)
                    fmt.Printf("Set-Cookie: passhash=%s\n", hashed_pass)
                }
                return true
            }
        }
    }
    password = ""
    return false
}
