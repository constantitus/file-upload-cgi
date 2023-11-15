package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
)

func CheckCredentials() (cred_valid bool) {
    if username != "" {
        cred_file, err := os.ReadFile(UserDir + username + ".txt")
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
    return
}

