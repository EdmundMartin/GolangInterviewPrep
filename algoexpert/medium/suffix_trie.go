package medium

type SuffixTrie map[byte]SuffixTrie

func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	for idx := range str {
		node := trie
		for j := idx; j < len(str); j++ {
			letter := str[j]
			if _, found := node[letter]; !found {
				node[letter] = NewSuffixTrie()
			}
			node = node[letter]
		}
		node['*'] = nil
	}
}

func (trie SuffixTrie) Contains(str string) bool {
	node := trie
	for idx := range str {
		letter := str[idx]
		if _, found := node[letter]; !found {
			return false
		}
		node = node[letter]
	}
	_, found := node['*']
	return found
}
