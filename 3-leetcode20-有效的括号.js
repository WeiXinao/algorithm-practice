/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    const stack = [];
    const left = '([{';
    const right = ')]}';
    let i;
    for (i = 0; i < s.length; i++) {
        if (left.includes(s[i])) {
            stack.push(s[i]);
        } else if (right.includes(s[i]) && stack.length > 0) {
            if (s[i] === ')' && stack[stack.length - 1] === '(') {
                stack.pop();
            } else if (s[i] === ']' && stack[stack.length - 1] === '[') {
                stack.pop();
            } else if (s[i] === '}' && stack[stack.length - 1] === '{') {
                stack.pop();
            } else return false;
        } else return false;
    }
    if (stack.length !== 0) {
        return false;
    } else {
        return true;
    }
};

let s = "(]";
console.log(isValid(s));