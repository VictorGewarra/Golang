func searchRange(nums []int, target int) []int {
    f:=-1
    l:=-1
    for i,v :=range nums{
        if v==target {
            l=i;
            if f<0 {
                f=i
            }
        }
        
    }
      return []int{f,l}
}