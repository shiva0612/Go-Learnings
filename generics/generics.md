```
func slices.All[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]

Slice ~[]E 
why not 
Slice []E 

# this will throw error without ~[]E bcz this is named slice
type MyInts []int
var m MyInts
slices.All(m)  // ❌ compile error

# why even use named slices or anything like this 
type Scores []int
func (s Scores) Average() float64 {
    sum := 0
    for _, v := range s {
        sum += v
    }
    return float64(sum) / float64(len(s))
}
```

