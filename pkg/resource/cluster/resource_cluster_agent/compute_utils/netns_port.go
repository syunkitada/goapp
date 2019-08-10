package compute_utils

func AssignNetnsPortId(assignedNetnsPortIds []bool) int {
	size := len(assignedNetnsPortIds)
	for i := 0; i < size; i++ {
		if !assignedNetnsPortIds[i] {
			assignedNetnsPortIds[i] = true
			return i
		}
	}
	return -1
}
