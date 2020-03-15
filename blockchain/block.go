package blockchain

// BlockChain structure
type BlockChain struct {
	Blocks []*Block
}

// Block structure
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// AddBlock method that will append new block to our blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis function that will generate the genesis block (1st block of the blockchain)
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain function that initializes the blockchain by calling the constructor
// with the Genesis function thatgenerates the 1st block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// CreateBlock function that will actually create next Block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
