package main

import (
	"fmt"
	"riakao/pkg/evaluate"
	"riakao/pkg/parser"
	"riakao/pkg/scanner"
	"strings"
)

func main() {
	input := `{platform} in ["android", "ios"]`
	s := scanner.NewScanner(strings.NewReader(input))
	bs := scanner.NewBufferScanner(s)
	p := parser.NewParser(bs)
	expr, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", expr)

	result, err := evaluate.Evaluate(expr, map[interface{}]interface{}{
		"platform": "ios",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
