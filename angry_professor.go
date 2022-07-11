package hackerrank

//https://www.hackerrank.com/challenges/angry-professor/problem?isFullScreen=true
//input: k, a
//output: YES or NO
//k is the number of students that must be on time to not be late
//a is the list of arrival times
//if the number of students on time is less than k, return NO
func angryProfessor(k int32, a []int32) string {
	var onTime []int32
	for i := 0; i < len(a); i++ {
		if a[i] <= 0 {
			onTime = append(onTime, a[i])
		}
	}
	if len(onTime) < int(k) {
		return "YES"
	}
	return "NO"
}
