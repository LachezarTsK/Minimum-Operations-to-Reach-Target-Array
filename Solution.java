
import java.util.HashSet;
import java.util.Set;

public class Solution {

    public int minOperations(int[] input, int[] target) {
        Set<Integer> uniqueLettersMismatch = new HashSet<>();
        for (int i = 0; i < input.length; ++i) {
            if (input[i] != target[i]) {
                uniqueLettersMismatch.add(input[i]);
            }
        }
        return uniqueLettersMismatch.size();
    }
}
