package main

import (
    "fmt"
    "math"
)

func main() {
    //variables
    var calculate int
    
    //ask the user what number they would like to get the sqrt
    fmt.Print("Enter Number for NEwtons Method: ")
    fmt.Scanf("%d",&calculate)
   
    //returns the result
     fmt.Println("The Result is: ",Newt(calculate))
    
}

func Newt(newt int) float64 {

    newNewt := float64(newt)

    if newNewt == 0 { return 0 }
    
    z := 1.0
    for i := 0; i < int(newNewt); i++ {
        z = z - ((math.Pow(z, 2) - newNewt) / (2 * z))
    }

    return z
}
