package main
import (
    "fmt"
)


func UserInput() (float64, float64, float64) {
    /**
    Set user input
    :params:
        acceleration:float64
        velocity:float64
        displacement:float64
    :return:
        acceleration, velocity, displacement
    **/
    var acceleration float64
    var velocity float64
    var displacement float64

    fmt.Printf("Acceleration: ")
    _, err := fmt.Scanln(&acceleration)
    if err != nil {
        acceleration = 0.0
    }
    fmt.Printf("Initial Velocity: ")
    _, err = fmt.Scanln(&velocity)
    if err != nil {
        velocity = 0.0
    }

    fmt.Printf("Initial Displacement: ")
    _, err = fmt.Scanln(&displacement)
    if err != nil {
        displacement = 0.0
    }

    return acceleration, velocity, displacement
}


func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
    /**
    Given 3 float values, calculate the displacement over time
    :equation:
        s = ½ a t2 + vot + so
    :args:
        a:float64 - acceleration
        vo:float64 - initial velocity
        so:float64 - initial displacement
    :params:
        t:float64 - time
    :result:
        displacement over time
    */
    return func(t float64) float64 {
        return 0.5*a*t*t + vo*t + so
    }
}


func main() {
    /**
    main for module
    :params:
        acceleration:float64
        velocity:float64 - initial velocity
        displacement:float64 - initial displacement
        gdf:GenDisplaceFn - decleration of method

    */
    acceleration, velocity, displacement := UserInput()
    gdf := GenDisplaceFn(acceleration, velocity, displacement)
    var i int = 0
    for i < 10 {
        fmt.Printf("Time (seconds): ")
        var time float64
        _, err := fmt.Scanln(&time)
        if err != nil {
            fmt.Printf("Invalid value - Goodbye!")
            i = 11
        } else {
            var result float64 = gdf(time)
            fmt.Printf("With the following params: \n\tAcceleration: %f\n\tVelocity: %f\n\tInitial Displacement: %f\n\tTime: %f\nThe Displacement is: %f\n",
                  acceleration, velocity, displacement, time, result)
        }
        i++
    }
}

/**
Acceleration: 10
Initial Velocity: 2
Initial Displacement: 1
Time (seconds): 3
With the following params:
        Acceleration: 10.000000
        Velocity: 2.000000
        Initial Displacement: 1.000000
        Time: 3.000000
The Displacement is: 52.000000
Time (seconds): 5
With the following params:
        Acceleration: 10.000000
        Velocity: 2.000000
        Initial Displacement: 1.000000
        Time: 5.000000
The Displacement is: 136.000000
Acceleration: 10
Initial Velocity: 5
Initial Displacement: -15
Time (seconds): 3.2
With the following params:
        Acceleration: 10.000000
        Velocity: 5.000000
        Initial Displacement: -15.000000
        Time: 3.200000
The Displacement is: 52.200000
**/