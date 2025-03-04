
func (x int) isPrime() bool {
    if x <= 1 {
        return false
    }
    for i := 2; i * i <= x; i++ {
        if x % i == 0 {
            return false
        }
    }
    return true
}

func (x int) factorial() int {
    if x <= 1 {
        return 1
    }
    return x * (x-1).factorial()
}

func (x int) binomialCoefficient(n int) int {
    if n == 0 || n == x {
        return 1
    }
    return (x-n+1).binomialCoefficient(n-1) + (x-n).binomialCoefficient(n-1)
}