package main

import (
    _ "embed"
    "archive/zip"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    b64 "encoding/base64"
    "strings"
    "github.com/jpillora/opts"
)

//go:embed 0003-cluster-wide-machineconfigs.yaml.template 
var tmpl string

func main() {
    type config struct {
	File  string `opts:"help=file to load (required)"`
	Output string    `opts:"help=output format (optional)"`
    }
    c := config{}
    opts.Parse(&c)

    fmt.Println(BuildYML(c.File))
}

func BuildYML(file string) (string) {
    if _, err := os.Stat(file); err == nil {
        data := Extract(file)
        sEnc := b64.StdEncoding.EncodeToString([]byte(data))
        return strings.Replace(tmpl, "BASE64_ENCODED_PEM_FILE", sEnc, 2)
    } else {
        return string("Certificates zip file not found.")
    }
}

func Extract(src string) (string) {
    path := GetPath() + "/output-folder"
    files, err := Unzip(src, path)
    if err != nil {
        log.Fatal(err)
    }
    data := "" 
    for _, b := range files {
        if b == path + "/consumer_export.zip" {
            inner_files, err := Unzip(path + "/consumer_export.zip", "output-folder")
            if err != nil {
               log.Fatal(err)
            }
            for _, file := range inner_files {
                if strings.Contains(file, "pem") {
                     data = ReadPEM(file)
                }
            }
        }
    }
    DeleteWF()
    return string(data)
}

func ReadPEM(file string) (string) {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatal(err)
    }
    return string(data)
}

func GetPath() (string) {
    path, err := os.Getwd()
    if err != nil {
        log.Println(err)
    }
    return path
}

func DeleteWF() {
    err := os.RemoveAll(GetPath() + "/output-folder")
    if err != nil {
        log.Fatal(err)
    }
}

func Unzip(src string, dest string) ([]string, error) {

    var filenames []string

    r, err := zip.OpenReader(src)
    if err != nil {
        return filenames, err
    }
    defer r.Close()

    for _, f := range r.File {

        fpath := filepath.Join(dest, f.Name)

        if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
            return filenames, fmt.Errorf("%s: illegal file path", fpath)
        }

        filenames = append(filenames, fpath)

        if f.FileInfo().IsDir() {
            // Make Folder
            err := os.MkdirAll(fpath, os.ModePerm)
            if err != nil {
               return filenames, err
            }
            continue
        }

        err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm)
        if err != nil {
            return filenames, err
        }

        outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
        if err != nil {
            return filenames, err
        }

        rc, err := f.Open()
        if err != nil {
            return filenames, err
        }

        _, err = io.Copy(outFile, rc)

        outFile.Close()
        rc.Close()

        if err != nil {
            return filenames, err
        }
    }
    return filenames, nil
}
