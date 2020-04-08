package main

import "fmt"

/*
1.简单版
1.定义结构
	1.前区块哈希
	2.当前区块哈希
	3.数据
2.创建区块
3.生产哈希
4.引入区块链
5.添加区块
6.重构代码

2.升级版
1.补充区块字段
2.更新哈希计算函数
3.优化代码
*/

//1.定义结构
type Block struct {
	//1.前区块哈希
	PrevHash []byte
	//2.当前区块哈希
	Hash []byte
	//3.数据
	Data []byte
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	var (
		block *Block
	)

	block = &Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, //先填空 后面再计算 //TODO
		Data:     []byte(data),
	}
	return block
}

//引入哈希

func main() {
	block := NewBlock("测试", []byte{})
	fmt.Println(block.PrevHash)
	fmt.Println(block.Data)
	fmt.Printf("区块数据: %s\n", block.Data)
}
