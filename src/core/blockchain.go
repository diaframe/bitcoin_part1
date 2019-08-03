package core

import (
	"fmt"
	"log"
	"strconv"
)

type BlockChain struct {
	Blocks [] *Block
}

func (bc *BlockChain) SendData(data string){
	PreBlock :=bc.Blocks[len(bc.Blocks)-1]
	NewBlock :=GenerateNewBlock(data,*PreBlock)
	NewBlock.GetBitCoin(data)
	bc.AppendBlock(&NewBlock)
}

func (bc *BlockChain)AppendBlock (NewBlock *Block){
	if len(bc.Blocks) == 0{
		bc.Blocks=append(bc.Blocks,NewBlock)
		return
	}
	if isVaild(*NewBlock,*bc.Blocks[len(bc.Blocks)-1]){
		bc.Blocks=append(bc.Blocks,NewBlock)
	}else{
		log.Fatal("InValid block")
	}

}

func (bc *BlockChain) PrintBlock(){
	for _,block := range bc.Blocks{
		fmt.Printf("Index: %d \n",block.Index)
		fmt.Printf("Timestamp: %d \n",block.Timestamp)
		fmt.Printf("PreBlockHash: %s \n",block.PreBlockHash)
		fmt.Printf("Hash: %s \n",block.Hash)
		fmt.Printf("Data: %s \n",block.data)
		fmt.Println("Pow:"+strconv.FormatBool(block.isVaildHash(block.data)))
		fmt.Println()
	}
}

func GerenerateNewBlockChain()BlockChain{
	Chain :=BlockChain{[]*Block{}}
	Chain.GenerateGenesisBlock()
	return Chain
}

func isVaild(NewBlock Block,OldBlock Block)bool{
	if NewBlock.PreBlockHash!=OldBlock.Hash{
		return false
	}
	if NewBlock.Index!=OldBlock.Index+1{
		return false
	}
	if NewBlock.CaculateHash(NewBlock.data,NewBlock.Nonce)!=NewBlock.Hash{
		return  false
	}
	return true
}
