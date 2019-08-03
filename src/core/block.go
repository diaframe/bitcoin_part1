package core

import (
	"math/big"
	"time"
)

type Block struct {
	Index        int64
	Timestamp    int64
	data         string
	PreBlockHash string
	Hash         string
	Nonce        *big.Int
}


func GenerateNewBlock (data string,PreBlock Block) Block{
	NewBlock :=Block{}
	NewBlock.Index=PreBlock.Index+1
	NewBlock.Timestamp=time.Now().Unix()
	NewBlock.data=data
	NewBlock.PreBlockHash=PreBlock.Hash
	return NewBlock
}

func (bc *BlockChain)GenerateGenesisBlock(){
	data :="Genesis Block"
	PreBlock :=Block{}
	PreBlock.Index= -1
	PreBlock.Hash=""
	GenesisBlock :=GenerateNewBlock(data,PreBlock)
	GenesisBlock.GetBitCoin(data)
	bc.AppendBlock(&GenesisBlock)
}





