package core

type Position []int // -1: blank 0: left up 1: right up 2: middle 3: left down 4: right down

func (p Position) At(positions ...int) bool {
	if len(positions) != len(p) {
		return false
	}
	for i := 0; i < len(positions); i++ {
		if positions[i] != p[i] {
			return false
		}
	}
	return true
}

type Note struct {
	Foot     int // 0: left 1: right 2: both
	Position Position
}

func NewNote(foot int, positions ...int) Note {
	if len(positions) > 1 {
		foot = 2
	}
	return Note{
		Foot:     foot,
		Position: positions,
	}
}

func (n Note) Single() bool {
	return len(n.Position) == 1
}

type Spectrum struct {
	Notes *List[Note]
}

func NewSpectrum(notes ...Note) Spectrum {
	if len(notes) == 0 {
		notes = append(notes, NewNote(0, 0))
	}
	return Spectrum{
		Notes: NewList(notes...),
	}
}

func (s *Spectrum) Head() *Node[Note] {
	return s.Notes.Head
}

func (s *Spectrum) Tail() *Node[Note] {
	return s.Notes.Tail
}

func (s *Spectrum) FirstNote() Note {
	return s.Head().Data
}

func (s *Spectrum) LastNote() Note {
	return s.Tail().Data
}

func (s *Spectrum) FirstFoot() int {
	return s.FirstNote().Foot
}

func (s *Spectrum) LastFoot() int {
	return s.LastNote().Foot
}

func (s *Spectrum) Positions() []Position {
	positions := make([]Position, s.Notes.Len)
	s.Notes.ForRange(func(data Note) {
		positions = append(positions, data.Position)
	})
	return positions
}

var TwoNotesMap map[int]*List[int]
var FourNotesRunMap map[int]*List[int]
var FourNotesStairMap map[int]*List[int]
var FourNotesDiagonalMap map[int]*List[int]
var LeftSideMap map[int]*List[int]
var RightSideMap map[int]*List[int]

func Add(m map[int]*List[int], data *List[int]) {
	m[len(m)] = data
}

func init() {
	if TwoNotesMap == nil {
		TwoNotesMap = make(map[int]*List[int])
		Add(TwoNotesMap, NewList[int](0, 1))
		Add(TwoNotesMap, NewList[int](0, 2))
		Add(TwoNotesMap, NewList[int](0, 3))
		Add(TwoNotesMap, NewList[int](0, 4))
		Add(TwoNotesMap, NewList[int](1, 4))
		Add(TwoNotesMap, NewList[int](2, 1))
		Add(TwoNotesMap, NewList[int](2, 4))
		Add(TwoNotesMap, NewList[int](3, 0))
		Add(TwoNotesMap, NewList[int](3, 1))
		Add(TwoNotesMap, NewList[int](3, 2))
		Add(TwoNotesMap, NewList[int](3, 4))
		Add(TwoNotesMap, NewList[int](4, 1))
	}
	if FourNotesRunMap == nil {
		FourNotesRunMap = make(map[int]*List[int])
		Add(FourNotesRunMap, NewList[int](0, 1, 0, 1))
		Add(FourNotesRunMap, NewList[int](0, 2, 0, 2))
		Add(FourNotesRunMap, NewList[int](0, 3, 0, 3))
		Add(FourNotesRunMap, NewList[int](0, 4, 0, 4))
		Add(FourNotesRunMap, NewList[int](1, 4, 1, 4))
		Add(FourNotesRunMap, NewList[int](2, 1, 2, 1))
		Add(FourNotesRunMap, NewList[int](2, 4, 2, 4))
		Add(FourNotesRunMap, NewList[int](3, 0, 3, 0))
		Add(FourNotesRunMap, NewList[int](3, 1, 3, 1))
		Add(FourNotesRunMap, NewList[int](3, 2, 3, 2))
		Add(FourNotesRunMap, NewList[int](3, 4, 3, 4))
		Add(FourNotesRunMap, NewList[int](4, 1, 4, 1))
	}
	if FourNotesStairMap == nil {
		FourNotesStairMap = make(map[int]*List[int])
		Add(FourNotesStairMap, NewList[int](0, 2, 1, 4))
		Add(FourNotesStairMap, NewList[int](0, 3, 1, 4))
		Add(FourNotesStairMap, NewList[int](3, 0, 2, 1))
		Add(FourNotesStairMap, NewList[int](3, 2, 4, 1))
	}
	if FourNotesDiagonalMap == nil {
		FourNotesDiagonalMap = make(map[int]*List[int])
		Add(FourNotesDiagonalMap, NewList[int](0, 2, 4, 1))
		Add(FourNotesDiagonalMap, NewList[int](3, 2, 1, 4))
	}
	if LeftSideMap == nil {
		LeftSideMap = make(map[int]*List[int])
		Add(LeftSideMap, NewList[int](0, 3, 0, 2))
		Add(LeftSideMap, NewList[int](0, 3, 2, 3, 0, 2))
		Add(LeftSideMap, NewList[int](3, 0, 3, 2))
		Add(LeftSideMap, NewList[int](3, 0, 2, 0, 3, 2))
	}
	if RightSideMap == nil {
		RightSideMap = make(map[int]*List[int])
		Add(RightSideMap, NewList[int](2, 1, 4, 1))
		Add(RightSideMap, NewList[int](2, 4, 1, 4))
	}
}
