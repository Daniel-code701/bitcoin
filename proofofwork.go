package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//1.定义一个工作量证明的结构
type ProofOfWork struct {
	block *Block
	//一个非常大的数 有很多丰富的方法
	target *big.Int
}

//2.提供创建pow的函数
func NewProofOWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	//我们指定的难度值 现在是一个string类型 需要进行转换
	targetStr := "00000000000000111111"
	//引入的辅助变量 目的是将上面的难度值转成big.Int
	tmpInt := big.Int{}
	//将难度值赋值给Int 指定16进制的格式
	tmpInt.SetString(targetStr, 16)
	pow.target = &tmpInt
	return &pow
}

//3.提供不断计算的hash函数
//Run()
func (pow *ProofOfWork) Run() ([]byte, uint64) {

	var nonce uint64
	block := pow.block
	var hash [32]byte
	for {
		//1.拼装数据(区块的数据 不断变化的随机数)
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.PrevHash,
			block.MerkelRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		//将二维的切片数组连接起来返回一个一维的切片
		blockInfo := bytes.Join(tmp, []byte{})

		//2.作哈希运算
		hash = sha256.Sum256(blockInfo)
		//3.与pow中的target作比较
		tmpInt := big.Int{}
		//将我们得到的hash数组 转换为一个bigInt
		tmpInt.SetBytes(hash[:])

		//比较当前的哈希值 如果当前的哈希值小于当前的哈希值 就说明找到了 否则继续找
		if tmpInt.Cmp(pow.target) == -1 {
			//a.找到了 退出返回
			fmt.Printf("挖坑成功!hash: %s\n, nonce: %d\n", hash)
			break
		} else {
			//b.没找到 继续找 随机数+1
			nonce++
		}
	}
	return hash[:], nonce
}

//4.提供一个校验函数
func main() {

}
