/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function (nums) {
    let result = [];
    nums.sort((a, b) => a - b);
    for (let i = 0; i < nums.length - 2; i++) {
        let a = nums[i];
        if (0 < a) return result;
        if (0 < i && nums[i] == nums[i - 1]) {
            continue;
        }
        let indexB = i + 1;
        let indexC = nums.length - 1;
        while (indexB < indexC) {
            let b = nums[indexB];
            let c = nums[indexC];
            if (a + b + c < 0) {
                indexB++;
            } else if (0 < a + b + c) {
                indexC--;
            } else {
                result.push([a, b, c]);
                while (indexB < indexC && nums[indexB] == nums[indexB + 1]) {
                    indexB++;
                }
                while (indexB < indexC && nums[indexC] == nums[indexC - 1]) {
                    indexC--;
                }
                indexB++;
                indexC--;
            }
        }
    }
    return result;
};