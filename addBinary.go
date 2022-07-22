func AddBinary(a string, b string) string {
	var res string

	var minLength, maxLength, i int
	carry := "0"

	if len(a) > len(b) || len(a) == len(b) {
		minLength = len(b) - 1
		maxLength = len(a) - 1
	} else {
		minLength = len(a) - 1
		maxLength = len(b) - 1
	}
	i = 0
	lenA := len(a) - 1
	lenB := len(b) - 1
	for i <= minLength {
		last1 := a[lenA-i]
		last2 := b[lenB-i]

		if last1 == last2 && string(last1) == "1" {
			if carry == "1" {
				res = string("1") + res
				carry = "1"
			} else {
				carry = "1"
				res = string("0") + res
			}
		}
		if last1 == last2 && string(last1) == "0" {
			if carry == "1" {
				res = string("1") + res
				carry = "0"
			} else {
				res = string("0") + res
			}
		}
		if last1 != last2 {
			if carry == "1" {
				res = string("0") + res
				carry = "1"
			} else {
				res = string("1") + res
			}
		}
		i++
	}
	if lenA > lenB {
		for i <= maxLength {
			symbol := a[lenA-i]
			if carry == "0" {
				res = string(symbol) + res
			}
			if string(symbol) == "1" && carry == "1" {
				res = string("0") + res
				carry = "1"
			}
			if string(symbol) == "0" && carry == "1" {
				res = string("1") + res
				carry = "0"
			}
			i++
		}
	}
	if lenA < lenB {
		for i <= maxLength {
			symbol := b[lenB-i]
			if carry == "0" {
				res = string(symbol) + res
			}
			if string(symbol) == "1" && carry == "1" {
				res = string("0") + res
				carry = "1"
			}
			if string(symbol) == "0" && carry == "1" {
				res = string("1") + res
				carry = "0"
			}
			i++
		}
	}
	if carry == "1" {
		res = string("1") + res
	}

	return res
}