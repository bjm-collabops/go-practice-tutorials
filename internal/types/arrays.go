// 같은 디렉토리(internal/types/) 안의 모든 .go 파일들은 같은 패키지명을 사용해야 함
package types

import "fmt"

func DemonstrateArrays() {
	// 배열 (크기 고정) - [크기]타입
	var ages [3]int = [3]int{20, 25, 30}
	fmt.Println("Array:", ages, len(ages))

	// 배열로 선언 - append 불가능
	var fruitArray [3]string
	fruitArray[0] = "apple"
	fruitArray[1] = "banana"
	fmt.Println("Array:", fruitArray)
	// fruitArray = append(fruitArray, "cherry") // ❌ 에러!

	// 슬라이스 (크기 동적) - []타입만
	// Slice 는 Go만의 독특한 개념.
	var scores = []int{100, 50, 60} // ✅ append 가능
	fmt.Println("Slice:", scores)
	scores = append(scores, 70)
	fmt.Println("Slice after append:", scores)

	// 슬라이스로 선언
	var fruitSlice []string
	fruitSlice = append(fruitSlice, "apple", "banana")
	fmt.Println("Slice:", fruitSlice)
	fruitSlice = append(fruitSlice, "cherry")
	fmt.Println("Slice after append:", fruitSlice)
}
