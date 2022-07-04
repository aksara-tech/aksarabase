package scanner

type reflectScanner struct {
	OutputScanner
	InputScanner
	PointerScanner
}

func NewReflectScanner(outputScanner OutputScanner, inputScanner InputScanner, pointerScanner PointerScanner) *reflectScanner {
	return &reflectScanner{OutputScanner: outputScanner, InputScanner: inputScanner, PointerScanner: pointerScanner}
}
