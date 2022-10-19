package main

import "testing"

func CheckName(t *testing.T) {
	var result string = PickName()
	switch result {
	case
	    "Hrant",
	    "Anna",
	    "James":
    }
    t.Errorf("Unexpected name: %s", result)
}

