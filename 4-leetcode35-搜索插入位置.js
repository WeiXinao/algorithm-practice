/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function (nums, target) {
    let lo = 0;
    let hi = nums.length - 1;
    while (lo <= hi) {
        let mid = Math.floor((lo + hi) / 2);
        if (target > nums[mid]) {
            lo = mid + 1;
        } else if (target < nums[mid]) {
            hi = mid - 1;
        } else {
            return mid;
        }
    }
    nums.splice(lo, 0, target);
    return lo;
};

/*
0  1  3 
2  2  3
3  3  3
*/

let nums = [1, 3, 5, 6], target = 7;
console.log(searchInsert(nums, target));
console.log(nums);