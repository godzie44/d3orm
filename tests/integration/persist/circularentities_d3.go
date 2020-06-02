// Code generated by d3. DO NOT EDIT.

package persist

import "fmt"
import "d3/orm/entity"
import "database/sql/driver"

func (s *ShopCirc) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*ShopCirc)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			ExtractField:  s.__d3_makeFieldExtractor(),
			SetFieldVal:   s.__d3_makeFieldSetter(),
			CompareFields: s.__d3_makeComparator(),
			NewInstance:   s.__d3_makeInstantiator(),
			Copy:          s.__d3_makeCopier(),
		},
	}
}

func (s *ShopCirc) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*ShopCirc)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Name":
			return sTyped.Name, nil

		case "Profile":
			return sTyped.Profile, nil

		case "FriendShop":
			return sTyped.FriendShop, nil

		case "TopSeller":
			return sTyped.TopSeller, nil

		case "Sellers":
			return sTyped.Sellers, nil

		case "KnownSellers":
			return sTyped.KnownSellers, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *ShopCirc) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &ShopCirc{}
	}
}

func (s *ShopCirc) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*ShopCirc)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Name":
			eTyped.Name = val.(string)
			return nil
		case "Profile":
			eTyped.Profile = val.(entity.WrappedEntity)
			return nil
		case "FriendShop":
			eTyped.FriendShop = val.(entity.WrappedEntity)
			return nil
		case "TopSeller":
			eTyped.TopSeller = val.(entity.WrappedEntity)
			return nil
		case "Sellers":
			eTyped.Sellers = val.(entity.Collection)
			return nil
		case "KnownSellers":
			eTyped.KnownSellers = val.(entity.Collection)
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

func (s *ShopCirc) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*ShopCirc)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &ShopCirc{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		if srcTyped.Profile != nil {
			copy.Profile = srcTyped.Profile.(entity.Copiable).DeepCopy().(entity.WrappedEntity)
		}
		if srcTyped.FriendShop != nil {
			copy.FriendShop = srcTyped.FriendShop.(entity.Copiable).DeepCopy().(entity.WrappedEntity)
		}
		if srcTyped.TopSeller != nil {
			copy.TopSeller = srcTyped.TopSeller.(entity.Copiable).DeepCopy().(entity.WrappedEntity)
		}
		if srcTyped.Sellers != nil {
			copy.Sellers = srcTyped.Sellers.(entity.Copiable).DeepCopy().(entity.Collection)
		}
		if srcTyped.KnownSellers != nil {
			copy.KnownSellers = srcTyped.KnownSellers.(entity.Copiable).DeepCopy().(entity.Collection)
		}

		return copy
	}
}

func (s *ShopCirc) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*ShopCirc)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*ShopCirc)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Name":
			return e1Typed.Name == e2Typed.Name
		case "Profile":
			return e1Typed.Profile == e2Typed.Profile
		case "FriendShop":
			return e1Typed.FriendShop == e2Typed.FriendShop
		case "TopSeller":
			return e1Typed.TopSeller == e2Typed.TopSeller
		case "Sellers":
			return e1Typed.Sellers == e2Typed.Sellers
		case "KnownSellers":
			return e1Typed.KnownSellers == e2Typed.KnownSellers
		default:
			return false
		}
	}
}

func (s *ShopProfileCirc) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*ShopProfileCirc)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			ExtractField:  s.__d3_makeFieldExtractor(),
			SetFieldVal:   s.__d3_makeFieldSetter(),
			CompareFields: s.__d3_makeComparator(),
			NewInstance:   s.__d3_makeInstantiator(),
			Copy:          s.__d3_makeCopier(),
		},
	}
}

func (s *ShopProfileCirc) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*ShopProfileCirc)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Shop":
			return sTyped.Shop, nil

		case "Description":
			return sTyped.Description, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *ShopProfileCirc) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &ShopProfileCirc{}
	}
}

func (s *ShopProfileCirc) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*ShopProfileCirc)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Shop":
			eTyped.Shop = val.(entity.WrappedEntity)
			return nil
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

func (s *ShopProfileCirc) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*ShopProfileCirc)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &ShopProfileCirc{}

		copy.Id = srcTyped.Id
		copy.Description = srcTyped.Description

		if srcTyped.Shop != nil {
			copy.Shop = srcTyped.Shop.(entity.Copiable).DeepCopy().(entity.WrappedEntity)
		}

		return copy
	}
}

func (s *ShopProfileCirc) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*ShopProfileCirc)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*ShopProfileCirc)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Shop":
			return e1Typed.Shop == e2Typed.Shop
		case "Description":
			return e1Typed.Description == e2Typed.Description
		default:
			return false
		}
	}
}

func (s *SellerCirc) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*SellerCirc)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			ExtractField:  s.__d3_makeFieldExtractor(),
			SetFieldVal:   s.__d3_makeFieldSetter(),
			CompareFields: s.__d3_makeComparator(),
			NewInstance:   s.__d3_makeInstantiator(),
			Copy:          s.__d3_makeCopier(),
		},
	}
}

func (s *SellerCirc) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*SellerCirc)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Name":
			return sTyped.Name, nil

		case "CurrentShop":
			return sTyped.CurrentShop, nil

		case "KnownShops":
			return sTyped.KnownShops, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (s *SellerCirc) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &SellerCirc{}
	}
}

func (s *SellerCirc) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*SellerCirc)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Name":
			eTyped.Name = val.(string)
			return nil
		case "CurrentShop":
			eTyped.CurrentShop = val.(entity.WrappedEntity)
			return nil
		case "KnownShops":
			eTyped.KnownShops = val.(entity.Collection)
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

func (s *SellerCirc) __d3_makeCopier() entity.Copier {
	return func(src interface{}) interface{} {
		srcTyped, ok := src.(*SellerCirc)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		copy := &SellerCirc{}

		copy.Id = srcTyped.Id
		copy.Name = srcTyped.Name

		if srcTyped.CurrentShop != nil {
			copy.CurrentShop = srcTyped.CurrentShop.(entity.Copiable).DeepCopy().(entity.WrappedEntity)
		}
		if srcTyped.KnownShops != nil {
			copy.KnownShops = srcTyped.KnownShops.(entity.Copiable).DeepCopy().(entity.Collection)
		}

		return copy
	}
}

func (s *SellerCirc) __d3_makeComparator() entity.FieldComparator {
	return func(e1, e2 interface{}, fName string) bool {
		if e1 == nil || e2 == nil {
			return e1 == e2
		}

		e1Typed, ok := e1.(*SellerCirc)
		if !ok {
			return false
		}
		e2Typed, ok := e2.(*SellerCirc)
		if !ok {
			return false
		}

		switch fName {

		case "Id":
			return e1Typed.Id == e2Typed.Id
		case "Name":
			return e1Typed.Name == e2Typed.Name
		case "CurrentShop":
			return e1Typed.CurrentShop == e2Typed.CurrentShop
		case "KnownShops":
			return e1Typed.KnownShops == e2Typed.KnownShops
		default:
			return false
		}
	}
}
