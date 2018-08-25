package niuniu

import "fmt"

// GetPokersPattern 计算牌型
// boom 炸弹
// full_house 葫芦
// nn 牛牛
// n[1-9] 牛1 - 牛9
// n0 无牛
// straight_flush 同花顺
// flush 同花
// straight 顺子
// 计算模式
func GetPokersPattern(pokers []string, mode int32) (weight int32, pattern string, rate int32, e error) {
	if !isLegalPokers(pokers) {
		return -1, "", -1, ErrIllegalPokers
	}

	maxValue, err := getMaxPokerValues(pokers)
	if err != nil {
		return -1, "", -1, err
	}

	maxSuit, err := getMaxPokerSuit(pokers)
	if err != nil {
		return -1, "", -1, err
	}

	if mode == 0 {
		boom, err := isBoom(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if boom {
			return 60000 + 0 + maxValue*10 + maxSuit, "boom", 6, nil
		}

		full, err := isFullHouse(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if full {
			return 50000 + 0 + maxValue*10 + maxSuit, "full_house", 5, nil
		}

		cow, cowNumber, err := isCow(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if cow {
			if cowNumber == 0 {
				return 40000 + 0 + maxValue*10 + maxSuit, "nn", 4, nil
			} else if cowNumber == 9 {
				return 30000 + cowNumber*1000 + maxValue*10 + maxSuit, "n9", 3, nil
			} else if cowNumber == 7 || cowNumber == 8 {
				return 30000 + cowNumber*1000 + maxValue*10 + maxSuit, fmt.Sprintf("n%v", cowNumber), 2, nil
			} else {
				return 30000 + cowNumber*1000 + maxValue*10 + maxSuit, fmt.Sprintf("n%v", cowNumber), 1, nil
			}
		} else {
			return 20000 + 0 + maxValue*10 + maxSuit, "n0", 1, nil
		}
	} else {
		straightFlush, err := isStraightFlush(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if straightFlush {
			return 90000 + 0 + maxValue*10 + maxSuit, "straight_flush", 17, nil
		}

		boom, err := isBoom(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if boom {
			return 80000 + 0 + maxValue*10 + maxSuit, "boom", 16, nil
		}

		fiveSmallCow, err := isFiveSmallCow(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if fiveSmallCow {
			return 70000 + 0 + maxValue*10 + maxSuit, "5small", 15, nil
		}

		fiveFlowerCow, err := isFiveFlowerCow(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if fiveFlowerCow {
			return 60000 + 0 + maxValue*10 + maxSuit, "5flower", 14, nil
		}

		flush, err := isFlush(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if flush {
			return 50000 + 0 + maxValue*10 + maxSuit, "flush", 13, nil
		}

		straight, err := isStraight(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if straight {
			return 40000 + 0 + maxValue*10 + maxSuit, "straight", 12, nil
		}

		fullHouse, err := isFullHouse(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if fullHouse {
			return 30000 + 0 + maxValue*10 + maxSuit, "full_house", 11, nil
		}

		cow, cowNumber, err := isCow(pokers)
		if err != nil {
			return -1, "", -1, err
		}
		if cow {
			if cowNumber == 0 {
				return 20000 + 0 + maxValue*10 + maxSuit, "nn", 10, nil
			} else {
				return 10000 + cowNumber*1000 + maxValue*10 + maxSuit, fmt.Sprintf("n%v", cowNumber), cowNumber, nil
			}
		} else {
			return 0 + 0 + maxValue*10 + maxSuit, "n0", 1, nil
		}
	}
}

// SearchBestPokerPattern 搜索最佳模式
func SearchBestPokerPattern(pokers []string, mode int32) (best []string, weight int32, pattern string, rate int32, e error) {
	if !isLegalPokers(pokers) {
		return nil, -1, "", -1, ErrIllegalPokers
	}

	for _, v := range permutations(pokers) {
		w, p, r, err := GetPokersPattern(v, mode)
		if err != nil {
			return nil, -1, "", -1, err
		}
		if w > weight {
			best, weight, pattern, rate = v, w, p, r
		}
	}
	return
}

func permutations(slice []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(slice []string, n int) {
		if n == 1 {
			tmp := make([]string, len(slice))
			copy(tmp, slice)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(slice, n-1)
				if n%2 == 1 {
					tmp := slice[i]
					slice[i] = slice[n-1]
					slice[n-1] = tmp
				} else {
					tmp := slice[0]
					slice[0] = slice[n-1]
					slice[n-1] = tmp
				}
			}
		}
	}
	helper(slice, len(slice))
	return res
}
