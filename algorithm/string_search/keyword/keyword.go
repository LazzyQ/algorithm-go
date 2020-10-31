package keyword

type InvertedIndex struct {
	Index map[string][]string
}

func NewInvertedIndex() *InvertedIndex {
	return &InvertedIndex{
		Index: make(map[string][]string),
	}
}

func (i *InvertedIndex) Add(k, v string) {
	if iv, ok := i.Index[k]; ok {
		iv = append(iv, v)
		i.Index[k] = iv
		return
	}
	iv := make([]string, 0)
	iv = append(iv, v)
	i.Index[k] = iv
}

func (i *InvertedIndex) Search(keyword []string) string {
	return ""
}
