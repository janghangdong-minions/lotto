package main

import "html/template"

var (
	// golang 은 대문자 변수는 위에 변수와 같이 주석을 써야한다.

	// TEMPLATES 는 메모리의 할당되는 기본 템플릿 변수이다.
	TEMPLATES = template.New("")
)
