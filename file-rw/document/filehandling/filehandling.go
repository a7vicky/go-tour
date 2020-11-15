package main

import (  
    "fmt"
    "flag"
    "io/ioutil"
)

func main() {
    // file path passed from the command line 
    fptr := flag.String("fpath", "test.txt", "file path to read from")
    flag.Parse()
    //reads the file and returns a byte slice which is stored in data. 
    data, err := ioutil.ReadFile(*fptr)
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    //convert data to a string and display the contents of the file.
    fmt.Println("Contents of file:", string(data))
}
