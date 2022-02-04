package lib

import (
    "crypto/md5"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path"
    "strconv"
)

var staticFiles []string

func getStaticFiles() []string {
    dir := "static"
    if len(staticFiles) == 0 {
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            log.Fatal(err)
        }

        for _, file := range files {
            if !file.IsDir() {
                dat, err := os.ReadFile(path.Join(dir, file.Name()))
                if err == nil {
                    staticFiles = append(staticFiles, string(dat))
                } else {
                    fmt.Println("error reading file: ", file.Name(), err)
                }
            }
        }
    }
    return staticFiles
}

func intFromString(str string) int64 {
    data := []byte(str)
    hex := fmt.Sprintf("%x", md5.Sum(data))
    value, err := strconv.ParseInt(hex[:12], 16, 64)
    if err != nil {
        fmt.Printf ("Conversion failed: %s\n", err)
    } else {
        return value
    }
    return 0
}

func getStaticFileForHost(host string) string {
    intForHost :=  intFromString(host)
    staticFiles := getStaticFiles()

    staticIndex := intForHost % int64(len(staticFiles))
    return staticFiles[staticIndex]
}
