func processData(data []int) (int, error) {
    if len(data) == 0 {
        return 0, errors.New("data slice is empty")
    }

    sum := 0
    for _, num := range data {
        sum += num
    }

    avg := float64(sum) / float64(len(data))

    //Check for potential overflow before casting to int
    if avg > float64(math.MaxInt) || avg < float64(math.MinInt){
        return 0, fmt.Errorf("average value is too large or too small for int type: %f", avg)
    }

    return int(avg), nil
} 
import "math"
import "fmt"
import "errors"