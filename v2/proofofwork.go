package main

import "math/big"

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
	//TODO
	return []byte{}, 0
}

//4.提供一个校验函数
func main() {

}
