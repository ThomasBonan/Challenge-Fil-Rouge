package main

import "fmt"

func main() {
	flag := Xorring()
	fmt.Println("Bienvenue à ce CTF! Trouvez le flag caché.")
	fmt.Println(flag)
	_ = flag // Pour éviter les erreurs de variable inutilisée
}

func Xorring() string {
	var byteList []byte
	a := byte(0x46)
	b := byte(0x4c)
	c := byte(0x41)
	d := byte(0x47)
	e := byte(0x7b)
	f := byte(0x72)
	g := byte(0x65)
	h := byte(0x76)
	i := byte(0x65)
	j := byte(0x72)
	k := byte(0x73)
	l := byte(0x65)
	m := byte(0x5f)
	n := byte(0x65)
	o := byte(0x6e)
	p := byte(0x67)
	q := byte(0x7d)
	key := byte(0x5a)
	byteList = append(byteList, a^key)
	byteList = append(byteList, b^key)
	byteList = append(byteList, c^key)
	byteList = append(byteList, d^key)
	byteList = append(byteList, e^key)
	byteList = append(byteList, f^key)
	byteList = append(byteList, g^key)
	byteList = append(byteList, h^key)
	byteList = append(byteList, i^key)
	byteList = append(byteList, j^key)
	byteList = append(byteList, k^key)
	byteList = append(byteList, l^key)
	byteList = append(byteList, m^key)
	byteList = append(byteList, n^key)
	byteList = append(byteList, o^key)
	byteList = append(byteList, p^key)
	byteList = append(byteList, q^key)
	return string(byteList)
}
