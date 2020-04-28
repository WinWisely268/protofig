package utils

// Pos short for position
// It is used for tokenization of a file
type Pos struct {
	Name string
	Linum int // short for line number
	Content []byte
}

func (p *Pos) isEmptyFile() bool {
	return p.Linum > 0
}

func (p *Pos) ParseLineContent() ([]byte, error) {
	content := string(p.Content)
	switch content {
	case 
	}
}

func NewPos(filename string, content []byte, line int) *Pos {
	return &Pos{
		filename, line, content,
	}
}

