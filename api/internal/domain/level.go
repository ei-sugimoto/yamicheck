package domain

import "fmt"

type Level int

func (l Level) String() (string, error) {
	switch l {
	case 1:
		return "この求人は、安全です。特に応募しても問題外ありませんが、telegram等に誘導された場合は闇バイトの可能性があるので、周りに相談しましょう。:)", nil
	case 2:
		return "この求人は、注意すべき求人です。闇バイトの可能性があるので、周りや然るべき場所に相談しましょう。すぐに応募するのは危険です。:(", nil
	case 3:
		return "この求人は、危険な求人です。闇バイトの可能性が非常に高いので、基本的に応募はしないようにしましょう。", nil
	case 4:
		return "絶対的に闇バイトです。ヤミバイト、ダメ、ゼッタイ。", nil
	default:
		return "", fmt.Errorf("invalid level: %d", l)
	}
}

func (l *Level) Validate() error {
	if *l < 1 || *l > 4 {
		return fmt.Errorf("invalid level: %d", *l)
	}
	return nil
}

func (l *Level) Integer() int {
	return int(*l)
}
