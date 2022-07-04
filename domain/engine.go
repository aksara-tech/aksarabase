package domain

import "aksarabase-v2/usecase/scanner"

type Engine struct {
	//Scanner default scanner is scanner.reflectScanner
	Scanner scanner.Scanner
}
