/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
    if (strs.length === 1) return strs[0];
    function LCP(s1, s2) {
        let len = s1.length > s2.length ? s2.length : s1.length;
        let res = "";
        for (let i = 0; i < len; i++) {
            if (s1[i] === s2[i]) res += s1[i];
            else {
                break;
            }
        }
        return res;
    }
    let res = LCP(strs[0], strs[1]);
    if (strs.length > 2) {
        for (let i = 2; i < strs.length; i++) {
            res = LCP(res, strs[i]);
        }
        return res;
    } else {
        return res;
    }
};  

console.log(longestCommonPrefix(["dog","racecar","car"]));