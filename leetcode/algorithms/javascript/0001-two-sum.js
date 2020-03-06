/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
    let m = new Map();
    for (let i = 0; i < nums.length; i++) {
        let j = m.get(target - nums[i]);
        if (j != undefined && j != i) {
            return [j, i];
        }
        m.set(nums[i], i);
    }
};
