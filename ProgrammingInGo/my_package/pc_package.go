package my_package

import "fmt"

// Pc is a struct for pc public
type Pc struct {
    Ram int
    Disk int
    Brand string
}

func (myPcG Pc) String() string {
    return fmt.Sprintf("I have %dGB of ram, %dGb of disk and is an %s", myPcG.Ram, myPcG.Disk, myPcG.Brand)
}


func (myPcG Pc) Ping(){
    fmt.Println(myPcG.Brand, "Pong")
}

func (myPcG *Pc) DuplicateRam() {
    myPcG.Ram = myPcG.Ram * 2
}
