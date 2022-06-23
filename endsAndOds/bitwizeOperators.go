package main

import "fmt"

func main() {
	a, b := uint16(0xffff), uint16(180)
	fmt.Printf("%016b %04[1]x\n", a)         //1111111111111111 ffff
	fmt.Printf("%016b %04[1]x\n", a&^0b1111) //AND NOT  1111111111110000 fff0
	fmt.Printf("%016b %04[1]x\n", a&0b1111)  //AND 0000000000001111 000f
	fmt.Println()
	fmt.Printf("%016b %04[1]x\n", b)        //0000000010110100 00b4
	fmt.Printf("%016b %04[1]x\n", ^b)       //NOT 1111111101001011 ff4b
	fmt.Printf("%016b %04[1]x\n", b|0b1111) // OR 0000000010111111 00bf
	fmt.Printf("%016b %04[1]x\n", a^0b1111) //XOR  1111111111110000 fff0

}
