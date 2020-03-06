/**
 * @param {string} s
 * @return {string}
 */
/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function (s) {
    let start = 0, length = 0;
    let helper = function (i, j, start, length) {
        while (-1 < i && j < s.length && s[i] == s[j]) {
            i--;
            j++;
        }
        if (length < j - i - 1) {
            start = i + 1;
            length = j - i - 1;
        }
        return [start, length];
    };
    for (let i = 0; i < s.length; i++) {
        [start, length] = helper(i, i, start, length);
        [start, length] = helper(i, i + 1, start, length);
    }
    return s.slice(start, start + length);
};

// var longestPalindrome = function (s) {
//     let start = 0;
//     let length = 0;
//     let helper = function (i, j, start, length) {
//         while (i > -1 && j < s.length && s[i] == s[j]) {
//             i--;
//             j++;
//         }
//         if (length < j - i - 1) {
//             start = i + 1;
//             length = j - i - 1;
//         }
//         return [start, length];
//     };
//     for (let i = 0; i < s.length; i++) {
//         [start, length] = helper(i, i, start, length);
//         [start, length] = helper(i, i + 1, start, length);
//     }
//     return s.substr(start, length);
// };