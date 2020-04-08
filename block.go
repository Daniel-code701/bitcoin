package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

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
	//1.版本号
	Version uint64
	//2.前区块哈希
	PrevHash []byte
	//3.Merkel根 这就是一个哈希值
	MerkelRoot []byte
	//4.时间戳
	TimeStamp uint64
	//5.难度值
	Difficulty uint64
	//6.随机数 也就是挖坑要找的数据
	Nonce uint64

	//a.当前区块哈希 正常比特币没有当前区块哈希 我们为了实现方便 做了简化
	Hash []byte
	//b.数据
	Data []byte
}

//实现一个辅助函数 功能是将uint64转成byte
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}

//2.创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	var (
		block *Block
	)

	block = &Block{
		Version:    00,
		PrevHash:   prevBlockHash,
		MerkelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		Hash:       []byte{},
		Data:       []byte(data),
	}

	block.SetHash()

	return block
}

//3.引入哈希
func (block *Block) SetHash() {
	//blockInfo := make([]byte,10,20)
	//hash := [32]byte{}
	var (
		blockInfo []byte
	)
	//1.拼装数据
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.Data...)
	//2.sha256
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}

//4.引入区块链
type BlockChain struct {
	//定义一个区块链数组
	blocks []*Block
}

//5.定义一个区块链
func NewBlockChain() *BlockChain {
	//创建一个创世块 并作为第一个区块添加到区块中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

//6.添加区块
func (bc *BlockChain) AddBlock(data string) {
	//如何获取 前区块哈希?
	//获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	prevHash := lastBlock.Hash
	//a.创建新的区块
	block := NewBlock(data, prevHash)
	//b.添加到区块链数组中
	bc.blocks = append(bc.blocks, block)
}

//创世块
func GenesisBlock() *Block {
	//block := NewBlock("创世块",[]byte{})
	//return block
	return NewBlock("创世块", []byte{})
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("增加一个区块")
	bc.AddBlock("增加二个区块")
	for i, block := range bc.blocks {
		fmt.Printf("=====当前区块高度: %d\n", i)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据: %s\n", block.Data)
	}

}
