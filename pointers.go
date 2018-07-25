package main

import (
  "fmt"
)

func main(){
  car := Car{
    Door : 2,
    Transmistion : "automatic",
    Odometer : 36000,
  }
  
  describe(Car)
}

func describe(c car){
  fmt.Println("this car has ",c.Door," door")
  fmt.Println("this car's Transmition is ", c.Transmistion)
  fmt.Println("and the Odomater is about", c.Odomater)
}