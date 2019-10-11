package Chapter14

import (
	"fmt"
	"math"
	"strconv"
)

type polar struct{
	radius float64
	degrees float64
}
type Cartesian struct{
	x float64
	y float64
}

func (c *Cartesian) ToPolar(polar){

}
func receivePolar(chP chan polar){
	var radius,degrees string
	for{
		_, _ = fmt.Scanf("%s %s", &radius, &degrees)
		rd,_:=strconv.ParseFloat(radius,64)
		deg,_:=strconv.ParseFloat(degrees,64)
		chP<-polar{radius:rd,degrees:deg}
	}
}
func ConverSion(chp chan polar) chan Cartesian{
	cht:=make(chan Cartesian)
	go func(){
			polar:=<-chp
			x:=polar.radius*math.Cos(polar.degrees)
			y:=polar.radius*math.Sin(polar.degrees)
			cht<-Cartesian{x:x,y:y}
	}()
	return cht
}
func E10polarToCartesian(){
	chP:=make(chan polar)
	go receivePolar(chP)
	for{
		Car,_:=<-ConverSion(chP)
		fmt.Printf("The Cartisan coordinate is :(%f,%f)\n",Car.x,Car.y)
	}
}
