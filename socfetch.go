package socfetch

import (
	"time"
)

type Media interface {
	Type() string
	Text() string
	Created() time.Time
}

func Merge(med []Media, add ...[]Media) []Media {
	for _, v := range add {
		for _, j := range v {
			med = append(med, j)
		}
	}
	return med
}

type ByDate []Media

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Created().Sub(a[j].Created()) < 0 }
