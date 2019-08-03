package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
)



const difficulty=4

func (block *Block)CaculateHash(data string,nonce *big.Int) string{
	Hash :=string(block.Index)+string(block.Timestamp)+nonce.String()+block.data+block.PreBlockHash
	ByteInHashNonce :=sha256.Sum256([]byte(Hash))
	StringInHashNonce :=hex.EncodeToString(ByteInHashNonce[:])
	return StringInHashNonce
}

func (block *Block)GetBitCoin(data string){
	block.Nonce = big.NewInt(0)
	prefix :=strings.Repeat("0",difficulty)
	fmt.Printf(" Mining the block containing %s :\n",block.data)
	for {
		block.Hash =block.CaculateHash(data,block.Nonce)
		fmt.Printf("\r 挖矿中: %s", block.Hash)
		if strings.HasPrefix(block.Hash,prefix){
			if block.isVaildHash(data) {
				fmt.Println()
				fmt.Printf(" 挖矿成功 %s 一共: %d个",block.data,block.Nonce)
				fmt.Println()
				block.Hash=block.Hash
				break
		}
		}else {
			block.Nonce.Add(block.Nonce,big.NewInt(1))
		}
	}
}

func (block *Block)isVaildHash(data string) bool{
		if block.Hash != block.CaculateHash(data,block.Nonce) {
			return  false
		}
		return true
}
