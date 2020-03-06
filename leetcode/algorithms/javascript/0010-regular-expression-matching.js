/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
var isMatch = function (s, p) {
    if (p == '') {
        return s == '';
    }
    if (p.length > 1 && p[1] == '*') {
        return isMatch(s, p.substr(2)) || (s != '') && (s[0] == p[0] || p[0] == '.') && isMatch(s.substr(1), p)
    }
    return (s != '') && (s[0] == p[0] || p[0] == '.') && isMatch(s.substr(1), p.substr(1))
};

// var isMatch = function (s, p) {
//     if (p == '') {
//         return s == '';
//     }
//     if (p.length > 1 && p[1] == '*') {
//         return isMatch(s, p.slice(2)) || (s != '' && (s[0] == p[0] || p[0] == '.')) && isMatch(s.slice(1), p);
//     }
//     return s != '' && (s[0] == p[0] || p[0] == '.') && isMatch(s.slice(1), p.slice(1));
// };

// var isMatch = function(s, p) {
//     let reg = new RegExp('^'+p+'$');
//     return reg.test(s);
// };

// https://developer.mozilla.org/zh-TW/docs/Web/JavaScript/Guide/Regular_Expressions