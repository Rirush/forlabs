package tests

import (
	"fmt"
	"github.com/Rirush/forlabs"
	"testing"
)

func TestGrid(t *testing.T) {
	c, err := forlabs.Authenticate(ValidUsername, ValidPassword)
	if err != nil {
		t.Errorf("Authentication failed: %s\n", err)
		return
	}
	grid, err := c.GetGrid()
	if err != nil {
		t.Errorf("Get grid failed: %s\n", err)
		return
	}
	fmt.Println(grid)
}

func TestSchedule(t *testing.T) {
	c, err := forlabs.Authenticate(ValidUsername, ValidPassword)
	if err != nil {
		t.Errorf("Authentication failed: %s\n", err)
		return
	}
	sched, err := c.GetSchedule(0)
	if err != nil {
		t.Errorf("Get schedule failed: %s\n", err)
		return
	}
	fmt.Println(sched)
}

func TestStudies(t *testing.T) {
	c, err := forlabs.Authenticate(ValidUsername, ValidPassword)
	if err != nil {
		t.Errorf("Authentication failed: %s\n", err)
		return
	}
	st, err := c.GetStudies(133)
	if err != nil {
		t.Errorf("Get studies failed: %s\n", err)
		return
	}
	fmt.Println(st)
}