func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

    var sol1 []int = nums1
    var sol2 []int = nums2

    pa := 0 
    pb := 0

    k := 0  

    length1 := len(sol1) 
    length2 := len(sol2)

    first := 0
    second := 0  
    max := (length1 + length2) / 2 

    for {
        if k > max {
            break
        }

        min := 0
        if pa < length1 && pb < length2 {
            item_sol1 := sol1[pa]
            item_sol2 := sol2[pb]

            if item_sol1 < item_sol2 {
              min = item_sol1
              pa ++
             k++
            }else{
                min = item_sol2
                pb++
                k++
            }
        }else if pa < length1 {
            min = sol1[pa]
            pa ++ 
            k++
        }else{
            min = sol2[pb]
            pb++
            k++
        }
         // Consider odd even case
      if (length1 + length2) % 2 == 0 {

            if (k - 1) == (max - 1) {
                first = min
            }

            if (k - 1) == max {
                second = min
            }

        }else{
            if (k - 1) == max {
                first = min
                second = 0
            }
        }    
    }

    if (length1 + length2) % 2 == 0{
        return (float64(first) + float64(second)) / 2
    }else{
        return float64(first)
    }
}
