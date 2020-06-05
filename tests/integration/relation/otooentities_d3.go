// Code generated by d3. DO NOT EDIT.

package relation

import "fmt"
import "d3/orm/entity"
import "database/sql/driver"

func (s *ShopLL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*ShopLL)(nil),
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

func (s *ShopLL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*ShopLL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "ID":
			return sTyped.ID, nil

		case "Profile":
			return sTyped.Profile, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *ShopLL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &ShopLL{}
	}
}

func (s *ShopLL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*ShopLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Profile":
			eTyped.Profile = val.(*entity.Cell)
			return nil
		case "Data":
			eTyped.Data = val.(string)
			return nil

		case "ID":
			if valuer, isValuer := val.(driver.Valuer); isValuer {
				v, err := valuer.Value()
				if err != nil {
					return eTyped.ID.Scan(nil)
				}
				return eTyped.ID.Scan(v)
			}
			return eTyped.ID.Scan(val)
		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *ShopLL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*ShopLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &ShopLL{}

		copy.ID = srcTyped.ID
		copy.Data = srcTyped.Data

		if srcTyped.Profile != nil {
			copy.Profile = srcTyped.Profile.DeepCopy().(*entity.Cell)
		}

		return copy
	}
}

func (s *ShopLL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*ShopLL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*ShopLL)
		if !ok {
			return false
		}

		switch fName {

		case "ID":
			return e1Typed.ID == e2Typed.ID
		case "Profile":
			return e1Typed.Profile == e2Typed.Profile
		case "Data":
			return e1Typed.Data == e2Typed.Data
		default:
			return false
		}
	}
}

func (p *ProfileLL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*ProfileLL)(nil),
		TableName: "profile",
		Tools: entity.InternalTools{
			ExtractField:  p.__d3_makeFieldExtractor(),
			SetFieldVal:   p.__d3_makeFieldSetter(),
			CompareFields: p.__d3_makeComparator(),
			NewInstance:   p.__d3_makeInstantiator(),
			Copy:          p.__d3_makeCopier(),
		},
	}
}

func (p *ProfileLL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*ProfileLL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "ID":
			return sTyped.ID, nil

		case "Photo":
			return sTyped.Photo, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (p *ProfileLL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &ProfileLL{}
	}
}

func (p *ProfileLL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*ProfileLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "ID":
			eTyped.ID = val.(int32)
			return nil
		case "Photo":
			eTyped.Photo = val.(*entity.Cell)
			return nil
		case "Data":
			eTyped.Data = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (p *ProfileLL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*ProfileLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &ProfileLL{}

		copy.ID = srcTyped.ID
		copy.Data = srcTyped.Data

		if srcTyped.Photo != nil {
			copy.Photo = srcTyped.Photo.DeepCopy().(*entity.Cell)
		}

		return copy
	}
}

func (p *ProfileLL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*ProfileLL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*ProfileLL)
		if !ok {
			return false
		}

		switch fName {

		case "ID":
			return e1Typed.ID == e2Typed.ID
		case "Photo":
			return e1Typed.Photo == e2Typed.Photo
		case "Data":
			return e1Typed.Data == e2Typed.Data
		default:
			return false
		}
	}
}

func (p *PhotoLL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*PhotoLL)(nil),
		TableName: "photo",
		Tools: entity.InternalTools{
			ExtractField:  p.__d3_makeFieldExtractor(),
			SetFieldVal:   p.__d3_makeFieldSetter(),
			CompareFields: p.__d3_makeComparator(),
			NewInstance:   p.__d3_makeInstantiator(),
			Copy:          p.__d3_makeCopier(),
		},
	}
}

func (p *PhotoLL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*PhotoLL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "ID":
			return sTyped.ID, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (p *PhotoLL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &PhotoLL{}
	}
}

func (p *PhotoLL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*PhotoLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "ID":
			eTyped.ID = val.(int32)
			return nil
		case "Data":
			eTyped.Data = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (p *PhotoLL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*PhotoLL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &PhotoLL{}

		copy.ID = srcTyped.ID
		copy.Data = srcTyped.Data

		return copy
	}
}

func (p *PhotoLL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*PhotoLL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*PhotoLL)
		if !ok {
			return false
		}

		switch fName {

		case "ID":
			return e1Typed.ID == e2Typed.ID
		case "Data":
			return e1Typed.Data == e2Typed.Data
		default:
			return false
		}
	}
}

func (s *ShopEL) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*ShopEL)(nil),
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

func (s *ShopEL) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*ShopEL)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Profile":
			return sTyped.Profile, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *ShopEL) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &ShopEL{}
	}
}

func (s *ShopEL) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*ShopEL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Id":
			eTyped.Id = val.(int32)
			return nil
		case "Profile":
			eTyped.Profile = val.(*entity.Cell)
			return nil
		case "Data":
			eTyped.Data = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *ShopEL) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*ShopEL)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &ShopEL{}

		copy.Id = srcTyped.Id
		copy.Data = srcTyped.Data

		if srcTyped.Profile != nil {
			copy.Profile = srcTyped.Profile.DeepCopy().(*entity.Cell)
		}

		return copy
	}
}

func (s *ShopEL) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*ShopEL)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*ShopEL)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Profile":
			return e1Typed.Profile == e2Typed.Profile
		case "Data":
			return e1Typed.Data == e2Typed.Data
		default:
			return false
		}
	}
}
