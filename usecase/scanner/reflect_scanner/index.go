package reflect_scanner

import "aksarabase-v2/usecase/scanner"

type reflectScanner struct {
	scanner.OutputScanner
	scanner.InputScanner
	scanner.PointerScanner
}

func NewReflectScanner(outputScanner scanner.OutputScanner, inputScanner scanner.InputScanner, pointerScanner scanner.PointerScanner) *reflectScanner {
	return &reflectScanner{OutputScanner: outputScanner, InputScanner: inputScanner, PointerScanner: pointerScanner}
}