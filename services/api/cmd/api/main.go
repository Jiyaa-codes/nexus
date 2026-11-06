package main

import "fmt"

type Service struct {
    Name       string
    CPU        float64
    Memory     float64
    RPS        int
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
func checkService(service Service) string {
    if service.CPU > 90 || service.Memory > 90 {
        return "CRITICAL"
    }

    if service.CPU > 75 || service.Memory > 75 {
        return "WARNING"
    }

    return "HEALTHY"
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
     services := []Service{
        {
            Name:   "payment-service",
            CPU:    82.5,
            Memory: 68.2,
            RPS: 1200,
        },
        {
            Name:   "auth-service",
            CPU:    45.0,
            Memory: 40.0,
            RPS: 500,
        },
        {
            Name:   "db-service",
            CPU:    92.0,
            Memory: 65.0,
            RPS: 2500,
        },
    }

    fmt.Println("NEXUS SERVICE MONITOR")
    fmt.Println("---------------------")
//What does `for _, service := range services` mean?
//Ignore the index and access each service
    for _, service := range services {

        status := checkService(service)

        fmt.Println("\n" + service.Name)
        fmt.Printf("CPU: %.1f%%\n", service.CPU)
        fmt.Printf("Memory: %.1f%%\n", service.Memory)
        fmt.Printf("RPS: %d\n",service.RPS)
        fmt.Println("Status:", status)
    }
    fmt.Println("---------------------")
    totalServices := len(services)
    fmt.Println("Total Services:", totalServices)
    criticalCount := 0
    warningCount := 0
    healthyCount := 0
    for _, service := range services {

        status := checkService(service)

        if status == "CRITICAL" {
            criticalCount++
        }else if status =="WARNING"{
            warningCount++
        }else{
            healthyCount++
        }
    }

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