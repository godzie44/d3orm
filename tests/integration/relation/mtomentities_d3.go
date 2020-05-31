// Code generated by d3. DO NOT EDIT.

package relation

import "fmt"
import "d3/orm/entity"

func (b *BookLL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*BookLL)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: b.__d3_makeFieldExtractor(),
			FieldSetter:    b.__d3_makeFieldSetter(),
			CompareFields:  b.__d3_makeComparator(),
			Instantiator:   b.__d3_makeInstantiator(),
			Copier:         b.__d3_makeCopier(),
		},
	}
}

func (b *BookLL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*BookLL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "ID":
			return sTyped.ID, nil

		case "Authors":
			return sTyped.Authors, nil

		case "Name":
			return sTyped.Name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (b *BookLL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &BookLL{}
	}
}

func (b *BookLL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*BookLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "ID":
			eTyped.ID = val.(int32)
			return nil
		case "Authors":
			eTyped.Authors = val.(entity.Collection)
			return nil
		case "Name":
			eTyped.Name = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (b *BookLL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*BookLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &BookLL{}

		copy.ID = srcTyped.ID
		copy.Name = srcTyped.Name

		if srcTyped.Authors != nil {
			copy.Authors = srcTyped.Authors.(entity.Copiable).DeepCopy().(entity.Collection)
		}

		return copy
	}
}

func (b *BookLL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*BookLL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*BookLL)
		if !ok {
			return false
		}

		switch fName {

		case "ID":
			return e1Typed.ID == e2Typed.ID
		case "Authors":
			return e1Typed.Authors == e2Typed.Authors
		case "Name":
			return e1Typed.Name == e2Typed.Name
		default:
			return false
		}
	}
}

func (a *AuthorLL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*AuthorLL)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: a.__d3_makeFieldExtractor(),
			FieldSetter:    a.__d3_makeFieldSetter(),
			CompareFields:  a.__d3_makeComparator(),
			Instantiator:   a.__d3_makeInstantiator(),
			Copier:         a.__d3_makeCopier(),
		},
	}
}

func (a *AuthorLL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*AuthorLL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "ID":
			return sTyped.ID, nil

		case "Name":
			return sTyped.Name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (a *AuthorLL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &AuthorLL{}
	}
}

func (a *AuthorLL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*AuthorLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "ID":
			eTyped.ID = val.(int32)
			return nil
		case "Name":
			eTyped.Name = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (a *AuthorLL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*AuthorLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &AuthorLL{}

		copy.ID = srcTyped.ID
		copy.Name = srcTyped.Name

		return copy
	}
}

func (a *AuthorLL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*AuthorLL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*AuthorLL)
		if !ok {
			return false
		}

		switch fName {

		case "ID":
			return e1Typed.ID == e2Typed.ID
		case "Name":
			return e1Typed.Name == e2Typed.Name
		default:
			return false
		}
	}
}

func (b *BookEL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*BookEL)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: b.__d3_makeFieldExtractor(),
			FieldSetter:    b.__d3_makeFieldSetter(),
			CompareFields:  b.__d3_makeComparator(),
			Instantiator:   b.__d3_makeInstantiator(),
			Copier:         b.__d3_makeCopier(),
		},
	}
}

func (b *BookEL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*BookEL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Rel":
			return sTyped.Rel, nil

		case "Name":
			return sTyped.Name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (b *BookEL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &BookEL{}
	}
}

func (b *BookEL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*BookEL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Id":
			eTyped.Id = val.(int32)
			return nil
		case "Rel":
			eTyped.Rel = val.(entity.Collection)
			return nil
		case "Name":
			eTyped.Name = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (b *BookEL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*BookEL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &BookEL{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		if srcTyped.Rel != nil {
			copy.Rel = srcTyped.Rel.(entity.Copiable).DeepCopy().(entity.Collection)
		}

		return copy
	}
}

func (b *BookEL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*BookEL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*BookEL)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Rel":
			return e1Typed.Rel == e2Typed.Rel
		case "Name":
			return e1Typed.Name == e2Typed.Name
		default:
			return false
		}
	}
}

func (a *AuthorEL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*AuthorEL)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: a.__d3_makeFieldExtractor(),
			FieldSetter:    a.__d3_makeFieldSetter(),
			CompareFields:  a.__d3_makeComparator(),
			Instantiator:   a.__d3_makeInstantiator(),
			Copier:         a.__d3_makeCopier(),
		},
	}
}

func (a *AuthorEL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*AuthorEL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Rel":
			return sTyped.Rel, nil

		case "Name":
			return sTyped.Name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (a *AuthorEL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &AuthorEL{}
	}
}

func (a *AuthorEL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*AuthorEL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Id":
			eTyped.Id = val.(int32)
			return nil
		case "Rel":
			eTyped.Rel = val.(entity.Collection)
			return nil
		case "Name":
			eTyped.Name = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (a *AuthorEL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*AuthorEL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &AuthorEL{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		if srcTyped.Rel != nil {
			copy.Rel = srcTyped.Rel.(entity.Copiable).DeepCopy().(entity.Collection)
		}

		return copy
	}
}

func (a *AuthorEL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*AuthorEL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*AuthorEL)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Rel":
			return e1Typed.Rel == e2Typed.Rel
		case "Name":
			return e1Typed.Name == e2Typed.Name
		default:
			return false
		}
	}
}

func (r *Redactor) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*Redactor)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: r.__d3_makeFieldExtractor(),
			FieldSetter:    r.__d3_makeFieldSetter(),
			CompareFields:  r.__d3_makeComparator(),
			Instantiator:   r.__d3_makeInstantiator(),
			Copier:         r.__d3_makeCopier(),
		},
	}
}

func (r *Redactor) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*Redactor)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Name":
			return sTyped.Name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (r *Redactor) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &Redactor{}
	}
}

func (r *Redactor) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*Redactor)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Id":
			eTyped.Id = val.(int32)
			return nil
		case "Name":
			eTyped.Name = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (r *Redactor) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*Redactor)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &Redactor{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		return copy
	}
}

func (r *Redactor) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*Redactor)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*Redactor)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Name":
			return e1Typed.Name == e2Typed.Name
		default:
			return false
		}
	}
}
