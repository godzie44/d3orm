// Code generated by d3. DO NOT EDIT.

package benchmark

import "database/sql/driver"
import "fmt"
import "d3/orm/entity"

func (s *shop) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*shop)(nil),
		TableName: "shop",
		Tools: entity.InternalTools{
			ExtractField:  s.__d3_makeFieldExtractor(),
			SetFieldVal:   s.__d3_makeFieldSetter(),
			CompareFields: s.__d3_makeComparator(),
			NewInstance:   s.__d3_makeInstantiator(),
			Copy:          s.__d3_makeCopier(),
		},
	}
}

func (s *shop) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*shop)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "id":
			return sTyped.id, nil

		case "books":
			return sTyped.books, nil

		case "profile":
			return sTyped.profile, nil

		case "name":
			return sTyped.name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *shop) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &shop{}
	}
}

func (s *shop) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*shop)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "books":
			eTyped.books = val.(entity.Collection)
			return nil
		case "profile":
			eTyped.profile = val.(entity.WrappedEntity)
			return nil
		case "name":
			eTyped.name = val.(string)
			return nil

		case "id":
			if valuer, isValuer := val.(driver.Valuer); isValuer {
				v, err := valuer.Value()
				if err != nil {
					return eTyped.id.Scan(nil)
				}
				return eTyped.id.Scan(v)
			}
			return eTyped.id.Scan(val)
		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *shop) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*shop)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &shop{}

		copy.id = srcTyped.id
		copy.name = srcTyped.name

		if srcTyped.books != nil {
			copy.books = srcTyped.books.(entity.Copiable).DeepCopy().(entity.Collection)
		}
		if srcTyped.profile != nil {
			copy.profile = srcTyped.profile.(entity.Copiable).DeepCopy().(entity.WrappedEntity)
		}

		return copy
	}
}

func (s *shop) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*shop)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*shop)
		if !ok {
			return false
		}

		switch fName {

		case "id":
			return e1Typed.id == e2Typed.id
		case "books":
			return e1Typed.books == e2Typed.books
		case "profile":
			return e1Typed.profile == e2Typed.profile
		case "name":
			return e1Typed.name == e2Typed.name
		default:
			return false
		}
	}
}

func (p *profile) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*profile)(nil),
		TableName: "prof",
		Tools: entity.InternalTools{
			ExtractField:  p.__d3_makeFieldExtractor(),
			SetFieldVal:   p.__d3_makeFieldSetter(),
			CompareFields: p.__d3_makeComparator(),
			NewInstance:   p.__d3_makeInstantiator(),
			Copy:          p.__d3_makeCopier(),
		},
	}
}

func (p *profile) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*profile)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Description":
			return sTyped.Description, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (p *profile) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &profile{}
	}
}

func (p *profile) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*profile)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Description":
			eTyped.Description = val.(string)
			return nil

		case "Id":
			if valuer, isValuer := val.(driver.Valuer); isValuer {
				v, err := valuer.Value()
				if err != nil {
					return eTyped.Id.Scan(nil)
				}
				return eTyped.Id.Scan(v)
			}
			return eTyped.Id.Scan(val)
		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (p *profile) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*profile)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &profile{}

		copy.Id = srcTyped.Id
		copy.Description = srcTyped.Description

		return copy
	}
}

func (p *profile) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*profile)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*profile)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Description":
			return e1Typed.Description == e2Typed.Description
		default:
			return false
		}
	}
}

func (b *book) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*book)(nil),
		TableName: "book",
		Tools: entity.InternalTools{
			ExtractField:  b.__d3_makeFieldExtractor(),
			SetFieldVal:   b.__d3_makeFieldSetter(),
			CompareFields: b.__d3_makeComparator(),
			NewInstance:   b.__d3_makeInstantiator(),
			Copy:          b.__d3_makeCopier(),
		},
	}
}

func (b *book) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*book)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Authors":
			return sTyped.Authors, nil

		case "Name":
			return sTyped.Name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (b *book) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &book{}
	}
}

func (b *book) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*book)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Authors":
			eTyped.Authors = val.(entity.Collection)
			return nil
		case "Name":
			eTyped.Name = val.(string)
			return nil

		case "Id":
			if valuer, isValuer := val.(driver.Valuer); isValuer {
				v, err := valuer.Value()
				if err != nil {
					return eTyped.Id.Scan(nil)
				}
				return eTyped.Id.Scan(v)
			}
			return eTyped.Id.Scan(val)
		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (b *book) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*book)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &book{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		if srcTyped.Authors != nil {
			copy.Authors = srcTyped.Authors.(entity.Copiable).DeepCopy().(entity.Collection)
		}

		return copy
	}
}

func (b *book) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*book)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*book)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Authors":
			return e1Typed.Authors == e2Typed.Authors
		case "Name":
			return e1Typed.Name == e2Typed.Name
		default:
			return false
		}
	}
}

func (a *author) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*author)(nil),
		TableName: "author",
		Tools: entity.InternalTools{
			ExtractField:  a.__d3_makeFieldExtractor(),
			SetFieldVal:   a.__d3_makeFieldSetter(),
			CompareFields: a.__d3_makeComparator(),
			NewInstance:   a.__d3_makeInstantiator(),
			Copy:          a.__d3_makeCopier(),
		},
	}
}

func (a *author) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*author)
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

func (a *author) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &author{}
	}
}

func (a *author) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*author)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Name":
			eTyped.Name = val.(string)
			return nil

		case "Id":
			if valuer, isValuer := val.(driver.Valuer); isValuer {
				v, err := valuer.Value()
				if err != nil {
					return eTyped.Id.Scan(nil)
				}
				return eTyped.Id.Scan(v)
			}
			return eTyped.Id.Scan(val)
		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (a *author) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*author)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &author{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		return copy
	}
}

func (a *author) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*author)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*author)
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
