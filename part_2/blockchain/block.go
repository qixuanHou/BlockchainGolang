package blockchain

type BlockChain struct {
	blocks []*Block
}
type Block struct {
	Hash 		[] byte
	Data 		 []byte
	PrevHash []byte
	Nonce 		int
}

func CreateBlock(data string, PrevHash []byte) * Block {
	block := &Block{[]byte{}, []byte(data), PrevHash, 0}
	pow := NewProof(block)
	nounce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nounce
	return block
}

func (chain * BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func (chain * BlockChain) GetBlocks() []*Block {
	return chain.blocks;
}