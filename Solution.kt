
class Solution {

    fun minOperations(input: IntArray, target: IntArray): Int {
        val uniqueLettersMismatch = mutableSetOf<Int>()
        for (i in input.indices) {
            if (input[i] != target[i]) {
                uniqueLettersMismatch.add(input[i])
            }
        }
        return uniqueLettersMismatch.size
    }
}
