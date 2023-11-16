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
        return Check(false)
    }
    if cookies.user != "" {
        username = cookies.user
        password = cookies.pass
        return Check(true)
    }
    return false
}

func Check(cookie bool) bool {
    cred_file, err := os.ReadFile(UserDir + username + ".txt")
    if err == nil {
        file_string := strings.TrimSuffix(string(cred_file), "\n")
        if cookie && password == file_string {
            return true
        } else {
            tmp := sha256.New()
            tmp.Write([]byte(password))
            hashed_pass := hex.EncodeToString(tmp.Sum(nil))
            if hashed_pass == file_string {
                if remember {
                    fmt.Printf("Set-Cookie: username=%s\n", username)
                    fmt.Printf("Set-Cookie: passhash=%s\n", hashed_pass)
                }
                return true
            }
        }
    }
    return false
}
