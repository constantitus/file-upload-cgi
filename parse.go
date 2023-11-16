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

        _, tmp_file, found_file := strings.Cut(v, "filename=\"")
        if found_file {
            tmp_name, _, _ := strings.Cut(tmp_file, "\"")
            _, tmp_file, _ = strings.Cut(v, "\r\n\r\n")
            tmp_content, _ := strings.CutSuffix(tmp_file, "\r\n--")
            files = append(files, new_file(tmp_name, tmp_content))
            continue
        }

        tmp_user, found_user := strings.CutPrefix(v, "name=\"user\"\r\n\r\n")
        if found_user {
            username, _ = strings.CutSuffix(tmp_user, "\r\n--")
            continue
        }

        tmp_pass, found_pass := strings.CutPrefix(v, "name=\"pass\"\r\n\r\n")
        if found_pass {
            password, _ = strings.CutSuffix(tmp_pass, "\r\n--")
            continue
        }

        _, remember = strings.CutPrefix(v, "name=\"remember\"\r\n\r\non")
        _, overwrite = strings.CutPrefix(v, "name=\"overwrite\"\r\n\r\non")
    }
}

