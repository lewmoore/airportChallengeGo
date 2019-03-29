package main

type Plane struct {
	isLanded bool
}

func NewPlane() Plane{
	return Plane{isLanded: true}
}

func PlaneTakeOff(plane Plane) bool {
	plane.isLanded = false
	return plane.isLanded
}

func main(){
}