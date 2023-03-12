/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function(s) {
    let sTrim = s.trim();
    if (sTrim.length > 0) {
        let slicedStr = sTrim.split(' ');
        return slicedStr[slicedStr.length - 1].length;
    } else {
        return 0;
    }
};

let s = "luffy is still joyboy";
console.log(lengthOfLastWord(s));