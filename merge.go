package main

import mergo "github.com/imdario/mergo"

func merge(dst interface{}, src interface{}, overwrite bool, append bool) error {
	if overwrite {
		return mergo.Merge(dst, src, mergo.WithOverride)
	} else if append {
		return mergo.Merge(dst, src, mergo.WithAppendSlice)
	}
	return mergo.Merge(dst, src)
}
