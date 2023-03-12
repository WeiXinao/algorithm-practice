/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function(a, b) {
    let carry = 0;
    let curr = 0;
    let temp = '';
    let res = '';
    let minLen = 0;
    if (a.length > b.length) {
        minLen = b.length;
        temp = a;
    } else if (a.length < b.length) {
        minLen = a.length;
        temp = b;
    } else {
        minLen = a.length;
    }
    for (let i = 0; i < minLen; i++) {
        let intA = parseInt(a[a.length - 1 - i]);
        let intB = parseInt(b[b.length - 1 - i]);
        curr = (intA + intB + carry) % 2;
        res = curr + res;
        carry = parseInt((intA + intB + carry) / 2);
    }
    if (temp.length === 0 && carry > 0) return carry + res;
    for (let i = minLen; i < temp.length; i++) {
        let intA = parseInt(temp[temp.length - 1 - i]);
        curr = (intA + carry) % 2;
        res = curr + res;
        carry = parseInt((intA + carry) / 2);
    }
    if (carry > 0) return carry + res;
    else return res;
};

let a = "100", b = "110010";
console.log(addBinary(a, b));