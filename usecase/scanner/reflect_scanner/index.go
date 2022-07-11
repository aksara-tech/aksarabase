package reflect_scanner

import "gitlab.com/aksaratech/aksarabase-go/v3/usecase/scanner"

type reflectScanner struct {
	scanner.OutputScanner
	scanner.InputScanner
	scanner.PointerScanner
}

func NewReflectScanner(outputScanner scanner.OutputScanner, inputScanner scanner.InputScanner, pointerScanner scanner.PointerScanner) *reflectScanner {
	return &reflectScanner{OutputScanner: outputScanner, InputScanner: inputScanner, PointerScanner: pointerScanner}
}
