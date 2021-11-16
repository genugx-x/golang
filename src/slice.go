package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	// fmt.Println(a)

	// s := make([]int, 5, 10)
	// println(len(s), cap(s))
	// fmt.Println(s)

	var e []int
	if e == nil {
		// println("Nil Slice")
	}
	// println(len(s), cap(s))

	b := []int{0, 1, 2, 3, 4, 5}
	b = b[2:4]
	// fmt.Println(b)

	c := []int{0, 1, 2, 3, 4, 5}
	c = c[2:5]
	fmt.Println(c)
	c = c[0:]
	fmt.Println(c)

	d := []int{0, 1}
	//println(d)
	//fmt.Println(d)
	d = append(d, 2)
	//println(d)
	//fmt.Println(d)
	d = append(d, 3, 4, 5)
	//println(d)
	//fmt.Println(d)

	sliceA := make([]int, 0, 3)
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// fmt.Println(len(sliceA), cap(sliceA))
	}
	// fmt.Println(sliceA)

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	// fmt.Println(target)
	// println(len(target), cap(target))
}

/**
 슬라이스는 크게 3개의 필드로 구성
 첫째 필드 - 배열의 포인터 정보
 두번째 필드 - 배열 요소로부터 길이
 세번째 필드 - 전체 배열의 크기

s := int {0, 1, 2, 3, 4, 5} len: 6, cap: 6
	=>	0 | 1 | 2 | 3 | 4 | 5

l := s[2:5] len: 3, cap: 4
	=> 2 | 3 | 4
	=> l[2:5]를 통해 2번 인덱스의 값의 주소값을 새로 가리키고 4번 인덱스까지
		 총 len 3을 가리킨다. 여기서 마지막 인덱스(5번)는 고정으로 가리켜
		 cap이 3이 아닌 4가 된다.
		 즉 [:]을 통해 가져오는 것은 사용하고 있는 길이 len 값만 가져온다. (l값을 출력해보면 3,4,5만 나옴)
		 마지막 인덱스(5번)의 값은 그 주소에 그대로 존재하기 때문에
		 l[0:]으로 출력을 하면 2, 3, 4만 가진것으로 보이지만
		 마지막 cap 인덱스까지 1[0:4]를 출력해보면 2, 3, 4, 5의 값을 확인할 수 있다.
**/
