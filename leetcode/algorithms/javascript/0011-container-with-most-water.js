/**
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
    let result = 0;
    let i = 0;
    let j = height.length - 1;
    while (i < j) {
        let h = height[i];
        if (height[j] < h) {
            h = height[j];
        }
        if (result < h * (j - i)) {
            result = h * (j - i);
        }
        if (h == height[i]) {
            i++;
        } else {
            j--;
        }
    }
    return result;
};

// var maxArea = function (height) {
//     let result = 0;
//     let i = 0, j = height.length - 1;
//     while (i < j) {
//         result = Math.max(result, Math.min(height[i], height[j]) * (j - i));
//         if (height[i] < height[j]) {
//             i++;
//         } else {
//             j--;
//         }
//     }
//     return result;
// };
