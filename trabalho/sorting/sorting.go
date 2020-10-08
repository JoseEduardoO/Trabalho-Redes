package sorting

// BubbleSort -> Implementação do método de ordenação Bubble sort
func BubbleSort(vetor ...int) []int {
	if len(vetor) < 2 {
		return vetor
	}

	vetorLength := len(vetor)

	for i := 0; i < vetorLength; i++ {
		for j := 0; j <= vetorLength-1; j++ {
			if vetor[i] < vetor[j] {
				vetor[i], vetor[j] = vetor[j], vetor[i]
			}
		}
	}

	return vetor
}

// InsertionSort -> Implementação do Insertion Sort
func InsertionSort(vetor ...int) []int {
	if len(vetor) < 2 {
		return vetor
	}

	for i, value := range vetor {
		j := i - 1

		for ; j >= 0 && value < vetor[j]; j-- {
			vetor[j+1] = vetor[j]
		}

		vetor[j+1] = value
	}

	return vetor
}

// SelectionSort -> Implementação do Selection Sort
func SelectionSort(vetor ...int) []int {
	if len(vetor) < 2 {
		return vetor
	}

	vetorLength := len(vetor)

	for i := 0; i < vetorLength-1; i++ {
		min := i

		for j := i + 1; j < vetorLength; j++ {
			if vetor[j] < vetor[min] {
				min = j
			}
		}

		vetor[i], vetor[min] = vetor[min], vetor[i]
	}

	return vetor
}
