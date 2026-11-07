
package main

import "testing"

func TestCheckHealth(t *testing.T) {
    tests := []struct {
        name     string
        cpu      float64
        memory   float64
        expected string
    }{
        {"High CPU", 95, 60, "CRITICAL"},
        {"High Memory", 60, 95, "CRITICAL"},
        {"CPU Warning", 80, 60, "WARNING"},
        {"Memory Warning", 60, 80, "WARNING"},
        {"Healthy Service", 50, 50, "HEALTHY"},
        {"Boundary Values", 90, 90, "WARNING"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            service := Service{
                CPU:             tt.cpu,
                Memory:          tt.memory,
                CPUThreshold:    90,
                MemoryThreshold: 90,
            }

            got := service.checkHealth()

            if got != tt.expected {
                t.Errorf(
                    "expected %s, got %s",
                    tt.expected,
                    got,
                )
            }
        })
    }
}
