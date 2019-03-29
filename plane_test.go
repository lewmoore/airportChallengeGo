package main

import "testing"

func TestPlaneIsLandedByDefault(t *testing.T) {
	got := NewPlane().isLanded
	result := true
    if got != result  {
        t.Errorf("PlaneLandedBuyDefault = %v; want True", got)
    }
}

func TestPlaneTakeOff(t *testing.T) {
    plane := NewPlane()

    expected := PlaneTakeOff(plane)
    result := plane.isLanded == false

    if expected != result {
        t.Errorf("PlaneTakeOff = %v; want False", expected)
    }
}