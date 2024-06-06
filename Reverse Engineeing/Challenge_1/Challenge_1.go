package main

import (
	"fmt"
)

func main() {
	hiddenFlag := []byte{0x46, 0x4c, 0x41, 0x47, 0x7b, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x65, 0x6e, 0x67, 0x7d} // FLAG{reverse_eng}
	flag := decodeFlag(hiddenFlag)
	fmt.Println("Bienvenue à ce CTF! Trouvez le flag caché.")
	_ = flag // Pour éviter les erreurs de variable inutilisée
}

func decodeFlag(encoded []byte) string {
	for i := range encoded {
		encoded[i] ^= 0x5a // XOR avec une valeur arbitraire pour obscurcir
	}
	return string(encoded)
}
