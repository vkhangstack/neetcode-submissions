func countStudents(students []int, sandwiches []int) int {
    n := len(students)
	q := make([]int, n)
	copy(q, students)

	res := n
	for _, sandwich := range sandwiches {
		cnt := 0
		for cnt < n && q[0] != sandwich {
			q = append(q[1:], q[0])
			cnt++
		}
		if q[0] == sandwich {
			q = q[1:]
			res--
		} else {
			break
		}
	}
	return res
}