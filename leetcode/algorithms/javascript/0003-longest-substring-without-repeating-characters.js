/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
    let length = 0, beforeSubstring = -1;
    let m = new Map();
    for (let i = 0; i < s.length; i++) {
        let j = m.get(s[i]);
        if (j != undefined && beforeSubstring < j) {
            beforeSubstring = j;
        }
        m.set(s[i], i);
        if (length < i - beforeSubstring) {
            length = i - beforeSubstring;
        }
    }
    return length;
};