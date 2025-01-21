package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoodOrderExample(t *testing.T) {
	strs := goodsOrder("agew")
	t.Log(strs)
	assert.Equal(t, 24, len(strs))
}

func goodsOrder(goods string) []string {
	strsMap := make(map[string]struct{})
	goodBytes := []byte(goods)
	for i := range len(goodBytes) {
		goodsOrderBytes(goodBytes, i, "", &strsMap)
	}
	strs := make([]string, 0, len(strsMap))
	for str := range strsMap {
		strs = append(strs, str)
	}
	return strs
}

func goodsOrderBytes(goods []byte, idx int, str string, strs *map[string]struct{}) {
	if goods[idx] == 0 {
		return
	}
	newGoods := make([]byte, len(goods))
	copy(newGoods, goods)
	str += string(newGoods[idx])
	newGoods[idx] = 0
	if len(str) == len(goods) {
		(*strs)[str] = struct{}{}
		return
	}
	for i := 0; i < len(newGoods); i++ {
		goodsOrderBytes(newGoods, i, str, strs)
	}
}
