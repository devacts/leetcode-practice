class Test
    def binary_search(a, target)
        left = 0
        right = a.length - 1

        while right >= left do
            mid = left + (right - left) / 2
            case target <=> a[mid]
            when 1
                left = mid + 1
            when -1
                right = mid - 1
            when 0
                return mid
            end
        end
        return -1
    end
end

arr = [2, 5, 8, 12, 16, 23, 38, 56, 72, 91]
#target = 91
#target = 72
#target = 100

#target = 2
#target = 0
#target = 16
target = 13

test = Test.new.binary_search(arr, target)
p test

# 1st iteration
# left = 0
# right = 9
# mid = 4

# true, 91 > a[4] or 16
# left = 5
# mid = 5 + (9-5)/2 = 5 + 2 = 7

# true, 91 > a[7] or 56
# left = 7 + 1 = 8
# mid = 8 + (9-8)/2 = 8 + 0 = 8

# true, 91 > a[8] or 72
# left = 9
# mid = 9 + (9-9)/2 = 9

# true, 91 == 91 # return 9
# incase target was 100
# true, 100 > a[9] or 91
# left = 10
# mid = 10 + (9-10)/2 = 10 + 0
