package caesar_cipher_encryptor

const (
	asciiLeft  = 97
	asciiRight = 122
	asciiLen   = 26
)

func CaesarCipherEncryptor(str string, key int) string {
	key = key % asciiLen
	p := []byte(str)
	buf := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		sh := p[i] + byte(key)
		if sh <= asciiRight {
			buf[i] = sh
		} else {
			h := asciiRight - p[i] + 1
			buf[i] = asciiLeft + byte(key) - h
		}
	}
	return string(buf)
}
