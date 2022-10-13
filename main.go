package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "time"
)

func main() {
    fmt.Printf("Track Selection Randomizer\n\n")

    time.Sleep(time.Second)
    
    file, err := os.ReadFile("./tracks.txt")
    if err != nil {
        println("tracks.txt not exists\nPlease decode Common.bmg from CTGPRC, only take track names out and put it into the same directory with track-selection-randomizer.exe")
        return
    }

    s := strings.Split(string(file), "\n")

    rand.Seed(time.Now().Unix())

    fmt.Printf("Picked track: %s\n", s[rand.Intn(len(s))])

    fmt.Println("Press any key to exit the program.")
    scan := bufio.NewScanner(os.Stdin)
    scan.Scan()
}
