function insert(intervals: number[][], newInterval: number[]): number[][] {
    
    if (intervals.length == 0) {
        return [newInterval]
    }
    
    const result :number[][] = []
    let start = 0
    
    while (start < intervals.length) {
        if(intervals[start][1] < newInterval[0]) {
            result.push(intervals[start])
        }else {
            break
        }
        start++
    }
    
    while(start < intervals.length) {
        // 1，5 -> 3,8
        if(intervals[start][0] < newInterval[0]) {
            newInterval[0] = intervals[start][0]
        }
        if(newInterval[1] <= intervals[start][1]) {
            // 2,5 vs [3,6] true
            // 2,5 vs [6,7] true 
            if(intervals[start][0] <= newInterval[1]) {
                newInterval[1] = intervals[start][1]
                start++
            }
            break
        }
        start++
    }
    
    result.push(newInterval, ...intervals.slice(start))
    
    return result
};