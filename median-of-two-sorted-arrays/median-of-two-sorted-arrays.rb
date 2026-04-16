def find_median_sorted_arrays(nums1, nums2)
   median = 0.0
   nums1.concat nums2
   nums1.sort!
   len = nums1.length
   if (len % 2) == 0
        median = (nums1[len/2] + nums1[len/2 - 1]) / 2.0
   else
        median = nums1[(len-1)/2.0]
   end
end


# def find_median_sorted_arrays_logarithmic(nums1, nums2)
     
# end