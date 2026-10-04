# 二分查找(Binary Search)
&emsp;&emsp;适用场景：从有序数组中找目标数


**二分查找框架如下：**

``` go
int binarySearch(int[] nums, int target) {
    int left = 0, right = ...;
    while(...) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            ...
        } else if (nums[mid] < target) {
            left = ...
        } else if (nums[mid] > target) {
            right = ...
        }
    }
    return ...;
}
```

**基于闭区间的二分查找：**
&emsp;&emsp;如下，由于right=nums.length-1，整个查找区间就变成了[left, right]，那整个查找的终止条件应该为left<=right，因为只有在[right+1, right]区间里面，才没有任何数可供查找。同理，每一次对left/right的重新计算都为mid的加1/减1，因为mid已经参与计算了，那下一次的查找就应该在[left, mid-1]或[mid+1, right]之间进行。
``` go
int binarySearch(int[] nums, int target) {
    int left = 0, right = nums.length-1;
    while(left <= right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            return mid
        } else if (nums[mid] < target) {
            left = mid + 1
        } else if (nums[mid] > target) {
            right = mid - 1
        }
    }
    return -1;
}
```

**基于开区间的二分查找：**
&emsp;&emsp;如下，由于right=nums.length，整个查找区间就变成了[left. right)，那整个查找的终止条件应该为left<right，因为只有在[right, right)区间里面，才没有任何数可供查找。同理，对left的重新计算为mid+1，对right的重新计算为mid，因为mid已经参与了计算，所以下次查找应该在[left, mid) 或 [mid+1, right)间进行
``` go
int binarySearch(int[] nums, int target) {
    int left = 0, right = nums.length;
    while(left < right) {
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            ...
        } else if (nums[mid] < target) {
            left = ...
        } else if (nums[mid] > target) {
            right = ...
        }
    }
    return ...;
}
```