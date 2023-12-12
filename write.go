package main

import (
	"os"
	"strconv"
	"strings"
)

func WriteFiles() {
    for _, file := range files {
        if file.name == "" {
            message.WriteString("<p>No file selected")
            continue
        }
        tmp := strings.Split(file.name, "/")
        file_name := tmp[len(tmp)-1]
        if file_name == "" {
            message.WriteString("<p>Invalid file name")
            continue
        }
        out, err := os.OpenFile(
            StorageDir + username + "/" + file_name,
            os.O_RDWR|os.O_CREATE,
            0644)
        if err != nil {
            message.WriteString("<p>")
            message.WriteString(err.Error())
            continue
        }
        out_stat, _ := out.Stat()
        exists := out_stat.Size() != 0
        message.WriteString("<p>")
        if exists {
            if overwrite {
                message.WriteString("over") // overwritten
                Write(&file, *out)
            } else {
                message.WriteString("File already exists: ")
                message.WriteString(file.name)
            }
        } else {
            Write(&file, *out)
        }
        out.Close()
    }
}

func Write(file *file_t, out os.File) {
    out_size, err := out.WriteString(file.contents)
    if err != nil {
        message.WriteString(err.Error())
    } else {
        message.WriteString("written ")
        message.WriteString(file.name)
        message.WriteString(" (")
        message.WriteString(strconv.Itoa(out_size))
        message.WriteString(" bytes)")
    }
}
