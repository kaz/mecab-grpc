package mecab

import (
	"sync"

	"github.com/shogo82148/go-mecab"
)

type (
	MecabPool struct {
		internalPool *sync.Pool
	}
	either struct {
		result mecab.MeCab
		err    error
	}
)

func NewPool(config Config) *MecabPool {
	return &MecabPool{
		internalPool: &sync.Pool{
			New: func() interface{} {
				result, err := mecab.New(config)
				if err != nil {
					return &either{err: err}
				}
				return &either{result: result}
			},
		},
	}
}

func (p *MecabPool) Get() (mecab.MeCab, error) {
	e := p.internalPool.Get().(*either)
	return e.result, e.err
}

func (p *MecabPool) Put(m mecab.MeCab) {
	p.internalPool.Put(&either{result: m})
}
