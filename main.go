package main

import (
    "os"
    "fmt"
    "encoding/pem"
    "crypto/x509"
    "io/ioutil"
    "errors"
)

func dieUsage(err error) {
    usageFmt := `
Usage: %s [keyfile]

Removes the password from an RSA key that was encrypted with an empty string
password.  keyfile should be in PEM format.  Unencrypted key will be written
to stdout in PEM format.`

    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n\n", err)
    }
    fmt.Fprintf(os.Stderr, usageFmt, os.Args[0])
    os.Exit(1)
}

func die(err error) {
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
}

func main() {
    if len(os.Args) != 2 {
        dieUsage(nil)
    }

    file := os.Args[1]

    if _, err := os.Stat(file); err != nil {
        dieUsage(err)
    }

    keyBytes, err := ioutil.ReadFile(file)
    if err != nil {
        die(err)
    }

    pemData, _ := pem.Decode(keyBytes)
    if pemData == nil {
        die(errors.New(fmt.Sprintf("PEM decode of %s failed", file)))
    }

    if pemData.Type != "RSA PRIVATE KEY" {
        die(errors.New(fmt.Sprintf("Expected 'RSA PRIVATE KEY' but got '%s'", pemData.Type)))
    }

    password := ""
    derBytes, err := x509.DecryptPEMBlock(pemData, []byte(password))
    if err != nil {
        die(err)
    }

    result := pem.Block{
        Type: "RSA PRIVATE KEY",
        Bytes: derBytes,
    }

    if err := pem.Encode(os.Stdout, &result); err != nil {
        die(err)
    }
}
