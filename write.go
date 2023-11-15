package main

import (
	"fmt"
	"os"
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
            out_stat, _ := out.Stat()
            if err != nil {
                message += fmt.Sprintf("<p>Error: %s", err)
                continue
            }
            if out_stat.Size() > 0 {
                message += fmt.Sprintf("<p>File already exists: %s", file.name)
                // TODO: maybe ask the user if they want to overwrite
            } else {
                out_size, err := out.WriteString(file.content)
                if err != nil {
                    message += fmt.Sprintf("<p>%s", err)
                } else {
                    message += fmt.Sprintf("<p>Written %s (%d bytes)", file.name, out_size)
                }
                out.Close()
            }
        }
}
