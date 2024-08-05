func isValid(s string) bool {
	stck := new(Stack)

	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stck.Push(byte(v))
		}

		if v == ')' {
			if stck.Top() == '(' {
				stck.Pop()
			} else {
                return false
            }
		}

		if v == ']' {
			if stck.Top() == '[' {
				stck.Pop()
			} else {
				return false
			}
		}

		if v == '}' {
			if stck.Top() == '{' {
				stck.Pop()
			} else {
				return false
			}
		}
	}

	return stck.isEmpty()
}

type Stack []byte

func (st *Stack) Push(v byte) {
	*st = append(*st, v)
}

func (st *Stack) Pop() error {
	if st.isEmpty() {
		return errors.New("stack is empty")
	}

	*st = (*st)[:len(*st)-1]

	return nil
}

func (st *Stack) isEmpty() bool {
	if len(*st) == 0 {
		return true
	}

	return false
}

func (st *Stack) Top() byte {
	if st.isEmpty() {
		return 0
	}

	return (*st)[len(*st)-1]
}