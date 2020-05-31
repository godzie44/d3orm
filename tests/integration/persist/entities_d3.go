// Code generated by d3. DO NOT EDIT.

package persist

import "database/sql/driver"
import "fmt"
import "d3/orm/entity"

func (s *Shop) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*Shop)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: s.__d3_makeFieldExtractor(),
			FieldSetter:    s.__d3_makeFieldSetter(),
			CompareFields:  s.__d3_makeComparator(),
			Instantiator:   s.__d3_makeInstantiator(),
			Copier:         s.__d3_makeCopier(),
		},
	}
}

func (s *Shop) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*Shop)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Books":
			return sTyped.Books, nil

		case "Profile":
			return sTyped.Profile, nil

		case "Name":
			return sTyped.Name, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *Shop) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &Shop{}
	}
}

func (s *Shop) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*Shop)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Books":
			eTyped.Books = val.(entity.Collection)
			return nil
		case "Profile":
			eTyped.Profile = val.(entity.WrappedEntity)
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

func (s *Shop) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*Shop)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &Shop{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		if srcTyped.Books != nil {
			copy.Books = srcTyped.Books.(entity.Copiable).DeepCopy().(entity.Collection)
		}
		if srcTyped.Profile != nil {
			copy.Profile = srcTyped.Profile.(entity.Copiable).DeepCopy().(entity.WrappedEntity)
		}

		return copy
	}
}

func (s *Shop) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*Shop)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*Shop)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Books":
			return e1Typed.Books == e2Typed.Books
		case "Profile":
			return e1Typed.Profile == e2Typed.Profile
		case "Name":
			return e1Typed.Name == e2Typed.Name
		default:
			return false
		}
	}
}

func (s *ShopProfile) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*ShopProfile)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: s.__d3_makeFieldExtractor(),
			FieldSetter:    s.__d3_makeFieldSetter(),
			CompareFields:  s.__d3_makeComparator(),
			Instantiator:   s.__d3_makeInstantiator(),
			Copier:         s.__d3_makeCopier(),
		},
	}
}

func (s *ShopProfile) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*ShopProfile)
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

func (s *ShopProfile) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &ShopProfile{}
	}
}

func (s *ShopProfile) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*ShopProfile)
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

func (s *ShopProfile) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*ShopProfile)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &ShopProfile{}

		copy.Id = srcTyped.Id
		copy.Description = srcTyped.Description

		return copy
	}
}

func (s *ShopProfile) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*ShopProfile)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*ShopProfile)
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

func (b *Book) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*Book)(nil),
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

func (b *Book) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*Book)
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

func (b *Book) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &Book{}
	}
}

func (b *Book) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*Book)
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

func (b *Book) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*Book)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &Book{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		if srcTyped.Authors != nil {
			copy.Authors = srcTyped.Authors.(entity.Copiable).DeepCopy().(entity.Collection)
		}

		return copy
	}
}

func (b *Book) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*Book)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*Book)
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

func (a *Author) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*Author)(nil),
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

func (a *Author) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*Author)
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

func (a *Author) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &Author{}
	}
}

func (a *Author) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*Author)
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

func (a *Author) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*Author)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &Author{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		return copy
	}
}

func (a *Author) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*Author)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*Author)
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
