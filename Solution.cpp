
#include <vector>
#include <unordered_set>
using namespace std;

class Solution {

public:
    int minOperations(const vector<int>& input, const vector<int>& target) const {
        unordered_set<int> uniqueLettersMismatch;
        for (int i = 0; i < input.size(); ++i) {
            if (input[i] != target[i]) {
                uniqueLettersMismatch.insert(input[i]);
            }
        }
        return uniqueLettersMismatch.size();
    }
};
