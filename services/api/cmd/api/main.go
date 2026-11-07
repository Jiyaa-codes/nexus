package main

import "fmt"

type Service struct {
    Name             string
    CPU              float64
    Memory           float64
    RPS              int
    CPUThreshold     float64
    MemoryThreshold  float64
}

/*func calculateHealth(CPU float64,Memory float64) string{
    if CPU>90 || Memory>90{
        return "CRITICAL"
    }
    if CPU>75 || Memory>75{
        return "WARNING"
    }
    return "HEALTHY"
}*/
//instead of passing two sep values:calhealth(cpu,memory) we can pass the whole struct checkservice(service)
//For now, we're using 83% of the critical threshold as the warning threshold. Later, we'll make warning thresholds configurable too.
/*func checkService(service Service) string {
    if service.CPU > service.CPUThreshold || service.Memory > service.MemoryThreshold {
        return "CRITICAL"
    }

    if service.CPU > service.CPUThreshold*0.83 || service.Memory > service.MemoryThreshold*0.83 {
        return "WARNING"
    }

    return "HEALTHY"
} day2*/
func (s Service) checkHealth() string {
    if s.CPU > s.CPUThreshold || s.Memory > s.MemoryThreshold {
        return "CRITICAL"
    }

    if s.CPU > 75 || s.Memory > 75 {
        return "WARNING"
    }

    return "HEALTHY"
}
// Pointer receiver: modifies original service
func (s *Service) updateCPU(value float64) {
    s.CPU = value
}
func main() {
    /*payment := Service{
        Name:      "payment-service",
        CPU:       82.5,
        Memory:    68.2,
    }
    auth :=Service{
        Name: "auth-service",
        CPU: 45.0,
        Memory: 40.0,
    }
    db :=Service{
        Name: "db-service",
        CPU: 92.0,
        Memory: 65.0,
    }*/
    // Demonstrating methods and pointers
    payment := Service{
        Name: "payment-service",
        CPU: 82.5,
        Memory: 68.2,
        RPS: 1200,
        CPUThreshold: 90,
        MemoryThreshold: 90,
    }

    fmt.Println("Before update: ",payment.checkHealth())

    payment.updateCPU(95)

    fmt.Println("AFter update:",payment.checkHealth())
    fmt.Println("Updated CPU: ",payment.CPU)

     services := []Service{
        {
            Name:   "payment-service",
            CPU:    82.5,
            Memory: 68.2,
            RPS: 1200,
            CPUThreshold: 85,
            MemoryThreshold: 90,
        },
        {
            Name:   "auth-service",
            CPU:    45.0,
            Memory: 40.0,
            RPS: 500,
            CPUThreshold: 95,
            MemoryThreshold: 90,
        },
    }
        /*{
            Name:   "db-service",
            CPU:    92.0,
            Memory: 65.0,
            RPS: 2500,
        }*/
    totalServices := len(services)
    criticalCount := 0
    warningCount := 0
    healthyCount := 0
    fmt.Println("NEXUS SERVICE MONITOR")
    fmt.Println("---------------------")
//What does `for _, service := range services` mean?
//Ignore the index and access each service
   
    for _, service := range services {

        status := service.checkHealth()

        fmt.Println("\n" + service.Name)
        fmt.Printf("CPU: %.1f%%\n", service.CPU)
        fmt.Printf("Memory: %.1f%%\n", service.Memory)
        fmt.Printf("RPS: %d\n",service.RPS)
        fmt.Println("Status:", status)
        switch status {
            case "CRITICAL":
                criticalCount++
            case "WARNING":
                warningCount++
            default:
                healthyCount++
        }
    }
    fmt.Println("---------------------")
    
    fmt.Println("Total Services:", totalServices)
    fmt.Println("Critical Services:", criticalCount)
    fmt.Println("Warning Services:", warningCount)
    fmt.Println("Healthy Services:", healthyCount)

    /*status := cpuhealth(payment.CPU,payment.Memory)
    status1 := cpuhealth(auth.CPU,auth.Memory)
    status2 := cpuhealth(db.CPU,db.Memory)
    fmt.Println(payment.Name)
    fmt.Printf("CPU: %.1f%%\n", payment.CPU)
    fmt.Printf("Memory: %.1f%%\n", payment.Memory)
    fmt.Println(status)

    fmt.Println(auth.Name)
    fmt.Printf("CPU: %.1f%%\n",auth.CPU)
    fmt.Printf("Memory: %.1f%%\n",auth.Memory)
    fmt.Println(status1)

    fmt.Println(db.Name)
    fmt.Printf("CPU: %.1f%%\n",db.CPU)
    fmt.Printf("Memory: %.1f%%\n",db.Memory)
    fmt.Println(status2)*/
}