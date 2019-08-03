package questions

import "fmt"

//Sample code to run this

// keyTimes := [][]int{{0, 2}, {1, 5}, {0, 9}, {2, 15}}
// slowestKey(keyTimes)

func SlowestKey1(keyTimes [][]int32) string {
	prevTime := keyTimes[0][1]
	var currentTime int32

	longestTimeKey := keyTimes[0][0]
	longestTime := keyTimes[0][1]

	for i := 1; i < len(keyTimes); i++ {
		currentTime = keyTimes[i][1]

		if currentTime-prevTime > longestTime {
			longestTime = currentTime - prevTime
			longestTimeKey = keyTimes[i][0]
		}

		prevTime = currentTime
	}

	return getChar(longestTimeKey)
}

func getChar(longestTimeKey int32) string {
	m1 := make(map[int32]string)
	m1[0] = "a"
	m1[1] = "b"
	m1[2] = "c"
	m1[3] = "d"
	m1[4] = "e"
	m1[5] = "f"
	m1[6] = "g"
	m1[7] = "h"
	m1[8] = "i"
	m1[9] = "j"
	m1[10] = "k"
	m1[11] = "l"
	m1[12] = "m"
	m1[13] = "n"
	m1[14] = "o"
	m1[15] = "p"
	m1[16] = "q"
	m1[17] = "r"
	m1[18] = "s"
	m1[19] = "t"
	m1[20] = "u"
	m1[21] = "v"
	m1[22] = "w"
	m1[23] = "x"
	m1[24] = "y"
	m1[25] = "z"

	return m1[longestTimeKey]
}

func SlowestKey(keyTimes [][]int) {
	prevTime := keyTimes[0][1]
	currentTime := 0

	longestTimeKey := keyTimes[0][0]
	longestTime := keyTimes[0][1]

	for i := 1; i < len(keyTimes); i++ {
		currentTime = keyTimes[i][1]

		if currentTime-prevTime > longestTime {
			longestTime = currentTime - prevTime
			longestTimeKey = keyTimes[i][0]
		}

		prevTime = currentTime
	}

	fmt.Println("longestTimeKey", longestTimeKey)
	fmt.Println("longestTime", longestTime)
}
