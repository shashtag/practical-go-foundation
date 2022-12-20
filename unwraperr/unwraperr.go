package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(errors.Unwrap(fmt.Errorf("ihiii %w ", errors.New("sss"))))
}
