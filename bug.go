func processData(data []int) (int, error) {
    if len(data) == 0 {
        return 0, errors.New("data slice is empty")
    }

    sum := 0
    for _, num := range data {
        sum += num
    }

    avg := float64(sum) / float64(len(data))

    //This line is the culprit. Go will panic here if the result of average is a very large number
    return int(avg), nil
}