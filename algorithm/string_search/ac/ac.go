package ac

import "container/list"

type node struct {
	Str   string
	Table map[int32]*node
	Fail  *node
}

func newNode() *node {
	return &node{
		Table: make(map[int32]*node),
	}
}

func (n node) IsWord() bool {
	return n.Str != ""
}

type AhoCorasickAutomation struct {
	Root    *node
	Targets []string
}

func NewAC(targets []string) *AhoCorasickAutomation {
	ac := &AhoCorasickAutomation{
		Root:    newNode(),
		Targets: targets,
	}
	ac.buildTrieTree()
	ac.build()
	return ac
}

func (ac *AhoCorasickAutomation) buildTrieTree() {
	for _, target := range ac.Targets {
		cur := ac.Root
		for _, char := range target {
			if n, ok := cur.Table[char]; !ok {
				nn := newNode()
				cur.Table[char] = nn
				cur = nn
			} else {
				cur = n
			}
		}
		cur.Str = target
	}
}

func (ac *AhoCorasickAutomation) build() {
	queue := list.New()

	for _, node := range ac.Root.Table {
		if node != nil {
			node.Fail = ac.Root
			queue.PushBack(node)
		}
	}

	for ; queue.Len() != 0; {
		e := queue.Front()
		queue.Remove(e)
		p := e.Value.(*node)
		for k, pv := range p.Table {
			if pv != nil {
				// 孩子节点入队列
				queue.PushBack(pv)
				failTo := p.Fail
				for {
					// 说明找到根节点还没找到
					if failTo == nil {
						pv.Fail = ac.Root
						break
					}

					if fv, ok := failTo.Table[k]; ok {
						pv.Fail = fv
						break
					} else {
						failTo = failTo.Fail
					}
				}
			}
		}
	}
}

func (ac *AhoCorasickAutomation) FindAll(text string) []string {
	cur := ac.Root
	result := make([]string, 0)

	str := []rune(text)
	for i := 0; i < len(str);  {
		char := str[i]
		if v, ok := cur.Table[char]; ok {
			cur = v
			if cur.IsWord() {
				result = append(result, cur.Str)
			}

			fail := cur.Fail
			for fail != nil && fail.IsWord() {
				result = append(result, fail.Str)
				fail = fail.Fail
			}
			i++
		} else {
			cur = cur.Fail
			if cur == nil {
				cur = ac.Root
				i++
			}
		}
	}
	return result
}
