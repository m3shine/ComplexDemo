package entity

import (
	"io/ioutil"
)

type Page struct {
	name string
	body []byte
}

func (p *Page) genHTML(){
	filename := p.name + ".html"
	return ioutil.WriteFile(filename,p.body,0600)
}


func init(){
}
