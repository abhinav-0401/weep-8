package weep

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	MEM_INIT_ADDRESS = 0x200
)

type System struct {
	Regs       [16]uint8
	Mem        [4096]uint8
	Index      uint16
	Pc         uint16
	Stack      [16]uint16
	StackPtr   uint8
	DelayTimer uint8
	SoundTimer uint8
	Display    [64 * 32]uint32
	Opcode     uint16
	rngSeed    rand.Source
}

func New() *System {
	var sys = &System{}
	sys.Pc = MEM_INIT_ADDRESS
	sys.rngSeed = rand.NewSource(time.Now().Unix())
	return sys
}

// Load ROM into Weep's memory
func (sys *System) LoadRom(path string) {
	// open the file
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	// get the length of the file
	var finalOffset int64
	finalOffset, err = file.Seek(0, 2)
	if err != nil {
		log.Fatal("error calculating the length of rom file", err)
	}

	// read the file
	file.Seek(0, 0)
	var rom = make([]byte, finalOffset)
	_, err = file.Read(rom)
	if err != nil {
		log.Fatal("error reading the file", err)
	}

	// load the contents into weep's mem, starting from 0x200
	for i, b := range rom {
		sys.Mem[MEM_INIT_ADDRESS+i] = b
	}
	sys.PrintMem()
}

func (sys *System) GenRand() int16 {
	var rand = rand.New(sys.rngSeed)
	return int16(math.Abs(rand.Float64() * 256))
}

func (sys *System) PrintMem() {
	for i := MEM_INIT_ADDRESS; i < 4096; i++ {
		if sys.Mem[i] == 0 {
			break
		}
		fmt.Println(sys.Mem[i])
	}
}
