package main

import "fmt"

// this is how you define a struct
type Device struct {
    Name string
    Id int
}

func main() {
    fmt.Println("Here are your devices:")

    // option 1 of creating an instance 
    device1 := Device{Name:"Keyboard", Id: 1}

    // option 2 (no keywords)
    device2 := Device{"Mouse", 2}

    // option 3 - creates a Pointer to a struct
    device3 := new(Device)
    device3.Name = "Printer"
    device3.Id = 3

    fmt.Println(device1)
    fmt.Println(device2)
    fmt.Println(device3)
}
