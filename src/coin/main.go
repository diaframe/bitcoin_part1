package main

import "core"

func main() {
	BlockChain := core.GerenerateNewBlockChain()
	BlockChain.SendData("send 1 to jacky")
	BlockChain.SendData("send 2 to sam")
	BlockChain.SendData("send 3 to tom")
	BlockChain.PrintBlock()
}
