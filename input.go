package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"schedsim/internal/scheduler"
)

type prompter struct{ r *bufio.Reader }

func (p *prompter) line(msg string) (string, error) {
	fmt.Print(msg)
	s, err := p.r.ReadString('\n')
	if err != nil && strings.TrimSpace(s) == "" {
		return "", err
	}
	return strings.TrimSpace(s), nil
}

// intVal asks until a whole number >= min is entered. Empty input returns def (if def >= min).
func (p *prompter) intVal(msg string, def, min int) (int, error) {
	for {
		s, err := p.line(msg)
		if err != nil {
			return 0, err
		}
		if s == "" && def >= min {
			return def, nil
		}
		n, convErr := strconv.Atoi(s)
		if convErr != nil || n < min {
			fmt.Printf("  please enter a whole number >= %d\n", min)
			continue
		}
		return n, nil
	}
}

// interactive collects the workload, algorithm and quantum from the user.
func interactive(p *prompter) ([]scheduler.Process, string, int, error) {
	fmt.Println("=== SchedSim interactive mode ===")
	n, err := p.intVal("Number of processes: ", -1, 1)
	if err != nil {
		return nil, "", 0, err
	}
	var procs []scheduler.Process
	seen := map[int]bool{}
	for i := 1; i <= n; i++ {
		fmt.Printf("\nProcess %d of %d\n", i, n)
		var pid int
		for {
			if pid, err = p.intVal(fmt.Sprintf("  PID [%d]: ", i), i, 0); err != nil {
				return nil, "", 0, err
			}
			if !seen[pid] {
				break
			}
			fmt.Println("  that PID is already used")
		}
		seen[pid] = true
		arr, err := p.intVal("  Arrival time [0]: ", 0, 0)
		if err != nil {
			return nil, "", 0, err
		}
		bur, err := p.intVal("  Burst time (>=1): ", -1, 1)
		if err != nil {
			return nil, "", 0, err
		}
		pri, err := p.intVal("  Priority (lower = higher) [1]: ", 1, 0)
		if err != nil {
			return nil, "", 0, err
		}
		procs = append(procs, scheduler.Process{PID: pid, Arrival: arr, Burst: bur, Priority: pri})
	}

	var algo string
	for {
		s, err := p.line(fmt.Sprintf("\nAlgorithm %v or 'all' [all]: ", scheduler.Algorithms))
		if err != nil {
			return nil, "", 0, err
		}
		algo = strings.ToLower(s)
		if algo == "" {
			algo = "all"
		}
		if algo == "all" || contains(scheduler.Algorithms, algo) {
			break
		}
		fmt.Println("  unknown algorithm")
	}
	quantum := 2
	if algo == "all" || algo == "rr" || algo == "mlfq" {
		if quantum, err = p.intVal("Time quantum [2]: ", 2, 1); err != nil {
			return nil, "", 0, err
		}
	}
	return procs, algo, quantum, nil
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}
