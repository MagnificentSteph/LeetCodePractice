func maxProfit(prices []int) int {
    n := len(prices)
    if n < 1{
        return 0
    }
    
    l := make([]int, n)
    r := make([]int, n)
    l_min := prices[0]
    r_max := prices[n-1]
    tmp := 0
    for i:=1;i<n;i++{
        if prices[i] < l_min{
            l_min = prices[i]
        }
        tmp = prices[i] - l_min
        if tmp > l[i-1]{
            l[i] = tmp
        }else{
            l[i] = l[i-1]
        }
    }
    
    for i:=n-2;i>-1;i--{
        if prices[i] > r_max{
            r_max = prices[i]
        }
        tmp = r_max -  prices[i]
        if tmp > r[i+1]{
            r[i] = tmp
        }else{
            r[i] = r[i+1]
        }
    }
    result := 0
    for i := 0; i < n;i++{
        tmp = l[i] + r[i]
        if tmp > result{
            result = tmp
        }
    }

    return result
}