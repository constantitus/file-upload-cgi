package main

import "strings"

func ParseBuffer(buffer []byte, boundary string) {
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
}

