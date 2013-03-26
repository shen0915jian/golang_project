package mymath

func Sqrt(x float32) float32 {
        z := 0.0
        for i := 0; i < 1000; i++ {
            z -= (z*z - x) / (2 * x)
        }
        return z
    }

