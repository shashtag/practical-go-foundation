package main

import (
	"fmt"
	"math/big"
)

func main() {

	count := big.NewInt(0)
	m := big.NewInt(1)
	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {
				for d := 0; d <= 9; d++ {
					for e := 0; e <= 9; e++ {
						for f := 0; f <= 9; f++ {
							for g := 0; g <= 9; g++ {
								for h := 0; h <= 9; h++ {
									for i := 0; i <= 9; i++ {
										for j := 0; j <= 9; j++ {
											if a+b+c+d+e+f+g+h+i+j == 20 {
												count.Add(count, m)
											}
										}
									}
								}
							}
						}

					}
				}
			}
		}
	}
	fmt.Println(count)

}
