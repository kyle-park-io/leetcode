func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    price := make([]int, n)
    for i := 0; i < n; i++ {
        price[i] = -1
    }

    price[src] = 0

    for stop := 0; stop <= k; stop++ {
        temp := make([]int, n)
        copy(temp, price)

        for _, flight := range flights {
            from := flight[0]
            to := flight[1]
            p := flight[2]

            if price[from] == -1 {
                continue
            }

            total := price[from] + p

            if temp[to] == -1 || temp[to] > total {
                temp[to] = total
            }
        }

        price = temp
    }

    return price[dst]
}