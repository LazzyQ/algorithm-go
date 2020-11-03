package main

import (
	"errors"
)

type Matcher struct {
	index [][]*AskWay
}

type AskWay struct {
	Keywords []string
	Repeat   []int8
	Question string
}

func NewMatcher() *Matcher {
	matcher := &Matcher{
		index: make([][]*AskWay, 4, 4),
	}
	return matcher
}

func (m *Matcher) Append(question string, keywords []string) error {
	keywordCountMap := make(map[string]int8, len(keywords))
	for _, keyword := range keywords {
		if count, ok := keywordCountMap[keyword]; ok {
			keywordCountMap[keyword] = count + 1
		} else {
			keywordCountMap[keyword] = 1
		}
	}

	repeat := make([]int8, 0, len(keywordCountMap))
	for _, w := range keywords {
		if c, ok := keywordCountMap[w]; ok {
			repeat = append(repeat, c)
		} else {
			return errors.New("关键词不匹配")
		}
	}

	askWay := &AskWay{
		Question: question,
		Keywords: keywords,
		Repeat:   repeat,
	}

	wc := len(askWay.Keywords)
	if len(m.index) < wc {
		newIndex := make([][]*AskWay, wc, wc)
		for i, askWays := range m.index {
			newIndex[i] = askWays
		}
		newIndex[wc-1] = make([]*AskWay, 0)
		newIndex[wc-1] = append(newIndex[wc-1], askWay)
		m.index = newIndex
	}
	askWays := m.index[wc-1]
	askWays = append(askWays, askWay)
	m.index[wc-1] = askWays
	return nil
}

func (m *Matcher) Search(keywords []string) string {
	if len(keywords) == 0 {
		return ""
	}

	keywordCountMap := make(map[string]int8, len(keywords))
	for _, keyword := range keywords {
		if count, ok := keywordCountMap[keyword]; ok {
			keywordCountMap[keyword] = count + 1
		} else {
			keywordCountMap[keyword] = 1
		}
	}

	for wc := len(keywordCountMap); wc > 0; wc-- {
		if matched := m.searchWithWC(wc, keywordCountMap); matched != nil {
			return matched.Question
		}
	}
	return ""
}

func (m *Matcher) searchWithWC(wc int, wcMap map[string]int8) *AskWay {
	minwc := len(m.index)
	if wc < minwc {
		minwc = wc
	}

	askWays := m.index[minwc-1]
	// 可以改成多个协程一起
	for _, askWay := range askWays {
		for i, keyword := range askWay.Keywords {
			if c, ok := wcMap[keyword]; !ok {
				break
			} else if askWay.Repeat[i] > c {
				break
			}

			if i == len(askWay.Keywords)-1 {
				return askWay
			}
		}
	}
	return nil

}
