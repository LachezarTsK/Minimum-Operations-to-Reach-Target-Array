
package main

func minOperations(input []int, target []int) int {
    uniqueLettersMismatch := NewHashSet[int]()
    for i := range input {
        if input[i] != target[i] {
            uniqueLettersMismatch.Add(input[i])
        }
    }
    return uniqueLettersMismatch.Size()
}

type HashSet[T comparable] struct {
    conainer map[T]bool
}

func NewHashSet[T comparable]() HashSet[T] {
    return HashSet[T]{conainer: map[T]bool{}}
}

func (this *HashSet[T]) Contains(value T) bool {
    return this.conainer[value]
}

func (this *HashSet[T]) Add(value T) {
    this.conainer[value] = true
}

func (this *HashSet[T]) Remove(value T) {
    delete(this.conainer, value)
}

func (this *HashSet[T]) Size() int {
    return len(this.conainer)
}
