package core

type Spectrum struct {
	FirstFoot  int // 0: left 1: right
	*List[int]     // -1: blank 0: left up 1: right up 2: middle 3: left down 4: right down
}

func NewSpectrum(firstFoot int, notes ...int) *Spectrum {
	return &Spectrum{
		FirstFoot: firstFoot,
		List:      NewList[int](notes...),
	}
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
