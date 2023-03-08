/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function(x) {
    if (x < 0) return false;
    let xStr = x.toString();
    for (let i = 0, j = xStr.length - 1; i++, j--; i < j) {
        if (xStr[i] != xStr[j]) return false;
    }
    return true;
};

console.log(isPalindrome(10));