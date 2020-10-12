package pjson



func (this *Schema) GetOtherTypes()[]Type {
	types := this.AllTypes()
	ret:=[]Type{}
	for _,t:=range types {
		if t==this.Type {
			continue
		}
		ret = append(ret, t)
	}
	return ret
}

func (this *Schema) AllTypes()[]Type {
	return []Type{Object,Array,String,Number,Boolean,Null}
}
