package mctype


type Module struct {
	Block  *Block
	Entity *Entity
	Point  Position
}

type Block struct {
	Name *string
	Data int16
}

type ConstBlock struct {
	Name string
	Data int16
}

type DoubleModule struct {
	Begin           Position
	End             Position
	Block, OldBlock *Block
	Entity          *Entity
}

var takenBlocks map[*ConstBlock]*Block = make(map[*ConstBlock]*Block)

func (req *ConstBlock) Take() *Block {
	block, ok := takenBlocks[req]
	if ok {
		return block
	}
	block=&Block {
		Name:&req.Name, //ConstBlock shouldn't be destroyed
		Data:req.Data,
	}
	takenBlocks[req]=block
	return block
}