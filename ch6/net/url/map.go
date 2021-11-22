package url

type Value map[string][]string // A Pointer

func (value Value) Get(k string) string {
	if sli := value[k]; sli != nil {
		return sli[0]
	}
	return ""
}

func (value Value) Add(k, v string) {
	if value == nil {
		return
	}
	value[k] = append(value[k], v)
}
