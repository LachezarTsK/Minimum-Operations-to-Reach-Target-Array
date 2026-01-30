
function minOperations(input: number[], target: number[]): number {
    const uniqueLettersMismatch = new Set<number>();
    for (let i = 0; i < input.length; ++i) {
        if (input[i] !== target[i]) {
            uniqueLettersMismatch.add(input[i]);
        }
    }
    return uniqueLettersMismatch.size;
};
