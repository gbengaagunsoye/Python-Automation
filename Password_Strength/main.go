package main

import "fmt"

func validatePassword(password string) bool {
	requiredLength := 8
	passLength := len(password)
	if passLength >= requiredLength {
		specialChar, lowerCase, upperCase, numbers := 0, 0, 0, 0
		for i := 0; i < passLength; i++ {
			asciiCode := int(password[i])
			fmt.Println(password[i])

			// Special Characters (32–47 / 58–64 / 91–96 / 123–126): Special characters include all printable characters that are neither letters nor numbers.
			// These include punctuation or technical, mathematical characters. ASCII also includes the space (a non-visible but printable character),
			// and, therefore, does not belong to the control characters category, as one might suspect.

			// Numbers (30–39): These numbers include the ten Arabic numerals from 0-9.
			// Letters (65–90 / 97–122): Letters are divided into two blocks, with the first
			// group containing the uppercase letters and the second group containing the lowercase

			if asciiCode >= 32 && asciiCode <= 47 || asciiCode >= 58 && asciiCode <= 64 || asciiCode >= 91 && asciiCode <= 96 || asciiCode >= 123 && asciiCode <= 47 {
				specialChar++
			} else if asciiCode >= 48 && asciiCode <= 57 {
				numbers++
			} else if asciiCode >= 65 && asciiCode <= 69 {
				upperCase++
			} else if asciiCode >= 97 && asciiCode <= 122 {
				lowerCase++
			}

			fmt.Println(specialChar, upperCase, lowerCase, numbers)

			if specialChar >= 1 && lowerCase >= 1 && upperCase >= 1 && numbers >= 1 {
				return true
			}

		}

	} else {
		return false
	}

	return false
}

func main() {

	myPassword := "Adegbeste@1"
	fmt.Println(myPassword)
	fmt.Println(validatePassword(myPassword))

}
