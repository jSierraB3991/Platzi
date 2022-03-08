package my_package

import "fmt"

// Pc is a struct for pc public
type Pc struct {
    Ram int
    Disk int
    Brand string
}


func (myPcG Pc) Ping(){
    fmt.Println(myPcG.Brand, "Pong")
}

func (myPcG *Pc) DuplicateRam() {
    myPcG.Ram = myPcG.Ram * 2
}
