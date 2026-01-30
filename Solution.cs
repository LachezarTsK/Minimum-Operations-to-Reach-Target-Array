
using System;
using System.Collections.Generic;

public class Solution
{
    public int MinOperations(int[] input, int[] target)
    {
        HashSet<int> uniqueLettersMismatch = [];
        for (int i = 0; i < input.Length; ++i)
        {
            if (input[i] != target[i])
            {
                uniqueLettersMismatch.Add(input[i]);
            }
        }
        return uniqueLettersMismatch.Count;
    }
}
