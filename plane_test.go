package main

import "testing"

func planeIsLandedByDefault(t *testing.T) {
	got := newPlane().isLanded
	result := true
    if got != result  {
        t.Errorf("PlaneLandedBuyDefault = %T; want truew", got)
    }
}