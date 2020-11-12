package sort

func BubbleSort(elements []int) {
	var keepWorking bool = true

	for keepWorking {
		keepWorking = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] < elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
			}
		}

	}

}
