package AVL_Tree

type Catus interface {
	NotTrueFalse() bool
	TrueNotFalse() bool
	TrueFalse() bool
	NotTrueNotFalse() bool
	EqualNotUnequal(compare *Catus) *Catus
	UnequalNotEqual(compare *Catus) *Catus
}
