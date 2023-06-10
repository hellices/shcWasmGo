package main

// syscall/js 패키지는 타깃 아키텍처로 wasm을 사용할 때
// Web Assembly 호스트 환경에 Javascript API에 대한 엑세스를 제공
import "syscall/js"

func main() {
	// 런타임에 표준 출력은 웹 브라우저 콘솔에 바인딩 되므로,
	// WASM 로딩이 완료되면 웹 브라우저 콘솔에 출력
	println("Wasm loaded.")

	// js.Global() 함수는 Javascript Global 네임스페이스에 엑세스
	alert := js.Global().Get("alert")

	// alert 함수를 호출합니다.
	alert.Invoke("Wasm 테스트 해보자~!")
}

// "export add" 주석을 이용해 Javascript에 add 함수를 exports
// 두 개의 32bit 정숫값을 받아서 하나의 32bit 정숫값을 반환하는 함수

//export add
func add(x int, y int) int {
	return x + y
}
