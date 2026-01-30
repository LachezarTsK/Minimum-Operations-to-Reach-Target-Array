
/**
 * @param {number[]} input
 * @param {number[]} target
 * @return {number}
 */
var minOperations = function (input, target) {
    const uniqueLettersMismatch = new Set();
    for (let i = 0; i < input.length; ++i) {
        if (input[i] !== target[i]) {
            uniqueLettersMismatch.add(input[i]);
        }
    }
    return uniqueLettersMismatch.size;
};
