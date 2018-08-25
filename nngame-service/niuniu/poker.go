package niuniu

import (
	"errors"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	// 扑克
	kPokers = []string{
		// 黑桃 spade
		"SA", "S2", "S3", "S4", "S5", "S6", "S7", "S8", "S9", "S10", "SJ", "SQ", "SK",
		// 红心 heart
		"HA", "H2", "H3", "H4", "H5", "H6", "H7", "H8", "H9", "H10", "HJ", "HQ", "HK",
		// 梅花 club
		"CA", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9", "C10", "CJ", "CQ", "CK",
		// 方块 diamond
		"DA", "D2", "D3", "D4", "D5", "D6", "D7", "D8", "D9", "D10", "DJ", "DQ", "DK",
	}

	// 计算中 10 以* 代替
	// 点数
	kPokerValues = map[uint8]int32{
		'A': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, '*': 10, 'J': 11, 'Q': 12, 'K': 13,
	}

	// 权重
	kPokerWeights = map[uint8]int32{
		'A': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, '*': 10, 'J': 10, 'Q': 10, 'K': 10,
	}

	// ErrIllegalPoker
	ErrIllegalPoker = errors.New("illegal poker")
)

// 转换
func convert(poker string) string {
	if len(poker) != 3 {
		return poker
	}

	if !(poker[1] == '1' && poker[2] == '0') {
		return poker
	}

	return string([]uint8{poker[0], '*'})
}

// 合法性检查
func isLegalPoker(poker string) bool {
	poker = strings.ToUpper(convert(poker))

	if len(poker) != 2 {
		return false
	}
	if poker[0] != 'S' && poker[0] != 'H' && poker[0] != 'C' && poker[0] != 'D' {
		return false
	}

	_, being := kPokerValues[poker[1]]

	return being
}

// 获取扑克点数
func getPokerValue(poker string) (int32, error) {
	if !isLegalPoker(poker) {
		return -1, ErrIllegalPoker
	}
	poker = convert(poker)
	return kPokerValues[poker[1]], nil
}

// 获取扑克权重
func getPokerWeight(poker string) (int32, error) {
	if !isLegalPoker(poker) {
		return -1, ErrIllegalPoker
	}
	poker = convert(poker)
	return kPokerWeights[poker[1]], nil
}

// 获取扑克花色
func getPokerSuit(poker string) (int32, error) {
	if !isLegalPoker(poker) {
		return -1, ErrIllegalPoker
	}

	switch {
	case strings.HasPrefix(poker, "S"):
		return 4, nil
	case strings.HasPrefix(poker, "H"):
		return 3, nil
	case strings.HasPrefix(poker, "C"):
		return 2, nil
	case strings.HasPrefix(poker, "D"):
		return 1, nil
	default:
		panic("unknown error: illegal poker")
	}
}

var (
	ErrIllegalPokers = errors.New("illegal pokers")
)

// 计算牌面
func getPokerValues(pokers []string) ([]int32, error) {
	values := make([]int32, 5)
	for i, v := range pokers {
		w, err := getPokerValue(v)
		if err != nil {
			return nil, err
		}
		values[i] = w
	}
	return values, nil
}

// 计算权重
func getPokerWeights(pokers []string) ([]int32, error) {
	values := make([]int32, 5)
	for i, v := range pokers {
		w, err := getPokerWeight(v)
		if err != nil {
			return nil, err
		}
		values[i] = w
	}
	return values, nil
}

// 计算花色
func getPokerSuits(pokers []string) ([]int32, error) {
	values := make([]int32, 5)
	for i, v := range pokers {
		w, err := getPokerSuit(v)
		if err != nil {
			return nil, err
		}
		values[i] = w
	}
	return values, nil
}

// 计算最大牌面
func getMaxPokerValues(pokers []string) (int32, error) {
	values, err := getPokerValues(pokers)
	if err != nil {
		return -1, err
	}
	var max int32
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// 计算最大权重
func getMaxPokerWeight(pokers []string) (int32, error) {
	weights, err := getPokerWeights(pokers)
	if err != nil {
		return -1, err
	}
	var max int32
	for _, v := range weights {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// 计算最大花色
func getMaxPokerSuit(pokers []string) (int32, error) {
	var max int32
	for _, v := range pokers {
		w, err := getPokerSuit(v)
		if err != nil {
			return -1, err
		}
		if w > max {
			max = w
		}
	}
	return max, nil
}

// 合法性检查
func isLegalPokers(pokers []string) bool {
	if len(pokers) != 5 {
		return false
	}

	for _, v := range pokers {
		if !isLegalPoker(v) {
			return false
		}
	}

	return true
}

// 是否同花顺
func isStraightFlush(pokers []string) (bool, error) {
	if !isLegalPokers(pokers) {
		return false, ErrIllegalPokers
	}

	straight, err := isStraight(pokers)
	if err != nil {
		return false, err
	}
	flush, err := isFlush(pokers)
	if err != nil {
		return false, err
	}

	return straight && flush, nil
}

// 是否炸弹
func isBoom(pokers []string) (bool, error) {
	if !isLegalPokers(pokers) {
		return false, ErrIllegalPokers
	}

	values, err := getPokerValues(pokers)
	if err != nil {
		return false, err
	}

	first := values[0] == values[1] && values[0] == values[2] && values[0] == values[3]
	last := values[4] == values[3] && values[4] == values[2] && values[4] == values[1]

	return first || last, nil
}

// 是否五小牛
func isFiveSmallCow(pokers []string) (bool, error) {
	if !isLegalPokers(pokers) {
		return false, ErrIllegalPokers
	}

	values, err := getPokerWeights(pokers)
	if err != nil {
		return false, err
	}

	var sum int32
	for _, v := range values {
		if v >= 5 {
			return false, nil
		}
		sum += v
	}

	return sum < 10, nil
}

// 是否五花牛
func isFiveFlowerCow(pokers []string) (bool, error) {
	if !isLegalPokers(pokers) {
		return false, ErrIllegalPokers
	}

	values, err := getPokerValues(pokers)
	if err != nil {
		return false, err
	}

	for _, v := range values {
		if v < 11 {
			return false, nil
		}
	}

	return true, nil
}

// 是否同花
func isFlush(pokers []string) (bool, error) {
	if !isLegalPokers(pokers) {
		return false, ErrIllegalPokers
	}

	suits, err := getPokerSuits(pokers)
	if err != nil {
		return false, err
	}

	for _, v := range suits {
		if suits[0] != v {
			return false, nil
		}
	}
	return true, nil
}

// 是否顺子
func isStraight(pokers []string) (bool, error) {
	if !isLegalPokers(pokers) {
		return false, ErrIllegalPokers
	}

	values, err := getPokerValues(pokers)
	if err != nil {
		return false, err
	}

	for i := range values {
		if i == 0 {
			continue
		}
		if values[i]-values[i-1] != 1 {
			return false, nil
		}
	}

	return true, nil
}

// 是否葫芦
func isFullHouse(pokers []string) (bool, error) {
	if !isLegalPokers(pokers) {
		return false, ErrIllegalPokers
	}

	values, err := getPokerValues(pokers)
	if err != nil {
		return false, err
	}

	first := (values[0] == values[1] && values[0] == values[2]) && (values[3] == values[4])
	last := (values[2] == values[3] && values[2] == values[4]) && (values[0] == values[1])

	return first || last, nil
}

// 是否有牛
// 若有牛，则返回 true，并给出具体类型，牛牛 则为 0
// 否则返回 false
func isCow(pokers []string) (bool, int32, error) {
	if !isLegalPokers(pokers) {
		return false, -1, ErrIllegalPokers
	}

	values, err := getPokerWeights(pokers)
	if err != nil {
		return false, -1, err
	}

	first := values[0] + values[1] + values[2]
	last := values[3] + values[4]

	return first%10 == 0, last % 10, nil
}

var (
	devLock = sync.Mutex{}
	dev     = rand.New(rand.NewSource(time.Now().Unix()))
)

// Acquire5 获取指定数量的5 张牌
func Acquire5(group int) [][]string {
	pool := make([]string, len(kPokers))
	copy(pool, kPokers)

	devLock.Lock()
	dev.Shuffle(len(pool), func(i, j int) {
		pool[i], pool[j] = pool[j], pool[i]
	})
	devLock.Unlock()

	if group*5 > len(pool) {
		panic("acquire too more poker numbers")
	}

	var result [][]string
	for i := 0; i < group; i++ {
		copied := append([]string{}, pool[:5]...)
		result = append(result, copied)
		pool = pool[5:]
	}
	return result
}
