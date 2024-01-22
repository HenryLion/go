package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func RandomGenerator1() uint64 {
	rnds := make([]byte, 8)
	_, err := rand.Read(rnds)
	if err != nil {
		panic(err)
	}
	return binary.BigEndian.Uint64(rnds)
}

// RandomGenerator 可以在不阻塞主程序的情况下持续生成随机数
func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()
	return c
}

func main() {
	randomNumbers := RandomGenerator()

	// 接收 10 个随机数
	for i := 0; i < 10; i++ {
		randomNumber := <-randomNumbers
		fmt.Println(randomNumber)
	}
}

// RandomGenerator 比RandomGenerator1函数的优势

// 如果使用RandomGenerator1函数，那么在生成随机数的过程中，主程序会被阻塞，
// 直到生成随机数的过程结束，主程序才能继续执行。
func gameLoop() {
	for {
		// 游戏的主循环逻辑
		randomNumber := RandomGenerator1()
		if randomNumber%100 == 0 {
			// spawnMonster() 打怪
		}
		if randomNumber%500 == 0 {
			// dropItem() 掉装备
		}
		// 更多的随机事件
	}
}

func gameLoop1() {
	randomNumbers := RandomGenerator()
	for {
		// 游戏的主循环逻辑
		select {
		case randomNumber, ok := <-randomNumbers:
			if !ok {
				return // 随机数生成器出错，结束游戏循环
			}
			if randomNumber%100 == 0 {
				// spawnMonster() 打怪
			}
			if randomNumber%500 == 0 {
				// dropItem() 掉装备
			}
			// 更多的随机事件
		default:
			// 无新的随机数，继续游戏循环
		}
	}
}
