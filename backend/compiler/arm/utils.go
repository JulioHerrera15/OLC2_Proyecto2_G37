package arm

func StringTo1ByteArray(str string) []byte {
	var resultado []byte
	var elementIndex int = 0

	for elementIndex < len(str) {
		resultado = append(resultado, str[elementIndex])
		elementIndex++		
	}

	resultado = append(resultado, 0) // Agregar el byte nulo al final

	return resultado
}