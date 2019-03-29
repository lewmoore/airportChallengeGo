package main

import "testing"

func TestPlaneIsLandedByDefault(t *testing.T) {
	got := NewPlane().isLanded
	result := true
    if got != result  {
        t.Errorf("PlaneLandedBuyDefault = %T; want truew", got)
    }
}