package main

import (
	"os"
	"strconv"
)

func WriteFiles() {
    for _, file := range files {
        if file.name == "" {
            message += "<p>No file selected"
            continue
        }
        out, err := os.OpenFile(
            StorageDir + username + "/" + file.name,
            os.O_RDWR|os.O_CREATE,
            0644)
        if err != nil {
            message += "<p>" + err.Error()
            continue
        }
        out_stat, _ := out.Stat()
        exists := out_stat.Size() != 0
        message += "<p>"
        if exists {
            if overwrite {
                message += "over" // overwritten
                Write(file, *out)
            } else {
                message += "File already exists: " + file.name
            }
        } else {
            Write(file, *out)
        }
        out.Close()
    }
}

func Write(file file_t, out os.File) {
    out_size, err := out.WriteString(file.content)
    if err != nil {
        message += err.Error()
    } else {
        message += "written " + file.name + " (" + strconv.Itoa(out_size) + " bytes)"
    }
}
