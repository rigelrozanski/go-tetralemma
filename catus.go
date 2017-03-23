package catus

import "fmt"

type Catus struct {
	Is   bool
	Isnt bool
}

func New() *Catus {
	return &Catus{
		Is:   false,
		Isnt: false,
	}
}

func (c *Catus) TrueNotFalse() bool {
	if Is && !Isnt {
		return true
	} else {
		return false
	}
}

func (c *Catus) NotTrueFalse() bool {
	if !Is && Isnt {
		return true
	} else {
		return false
	}
}

func (c *Catus) TrueFalse() bool {
	if Is && Isnt {
		return true
	} else {
		return false
	}
}

func (c *Catus) NotTrueNotFalse() bool {
	if !Is && !Isnt {
		return true
	} else {
		return false
	}
}

func (c *Catus) EqualNotUnequal(compare *Catus) *Catus {
	c1 := getcatus(c)
	c2 := getcatus(compare)

	switch {
	case c1.value == c2.value:
		return getCatus(catusTNF)
	case c1.opposite == c2.value:
		return getCatus(catusNTF)
	default:
		return getCatus(catusTF)
	}
}

func (c *Catus) UnequalNotEqual(compare *Catus) *Catus {
	c1 := getcatus(c)
	c2 := getcatus(compare)

	switch {
	case c1.value == c2.value:
		return getCatus(catusNTF)
	case c1.opposite == c2.value:
		return getCatus(catusTNF)
	default:
		return getCatus(catusNTNF)
	}
}

//////////////////////////////////// unexposed
const (
	TNF  byte = 0x00 //true and not false
	NTF  byte = 0x01 //not true and false
	TF   byte = 0x02 //both true and false
	NTNF byte = 0x03 //neither true nor false
)

type catus struct {
	value    byte
	opposite byte
}

var (
	catusTNF  = catus{TNF, NTF}
	catusNTF  = catus{NTF, TNF}
	catusTF   = catus{TF, NTNF}
	catusNTNF = catus{NTNF, TF}
)

func (c *Catus) getcatus() *catus {
	if c.Is {
		if c.Isnt {
			return catusTF
		} else {
			return catusTNF
		}
	} else {
		if c.Isnt {
			return catusNTF
		} else {
			return catusNTNF
		}
	}
}

func (c catus) getCatus() *Catus {
	switch c.value {
	case catusTNF:
		return &Catus{true, false}
	case catusNTF:
		return &Catus{false, true}
	case catusTF:
		return &Catus{true, true}
	case catusNTNF:
		return &Catus{false, false}
	}

}
