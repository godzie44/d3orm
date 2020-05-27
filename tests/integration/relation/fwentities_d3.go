// Code generated by d3. DO NOT EDIT.

package relation

import "fmt"
import "d3/orm/entity"

func (f *fwTestEntity1) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*fwTestEntity1)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: f.__d3_makeFieldExtractor(),
			FieldSetter:    f.__d3_makeFieldSetter(),
			Instantiator:   f.__d3_makeInstantiator(),
		},
	}
}

func (f *fwTestEntity1) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*fwTestEntity1)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Rel":
			return sTyped.Rel, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (f *fwTestEntity1) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &fwTestEntity1{}
	}
}

func (f *fwTestEntity1) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*fwTestEntity1)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Id":
			eTyped.Id = val.(int32)
			return nil
		case "Rel":
			eTyped.Rel = val.(entity.WrappedEntity)
			return nil
		case "Data":
			eTyped.Data = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (f *fwTestEntity2) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*fwTestEntity2)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: f.__d3_makeFieldExtractor(),
			FieldSetter:    f.__d3_makeFieldSetter(),
			Instantiator:   f.__d3_makeInstantiator(),
		},
	}
}

func (f *fwTestEntity2) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*fwTestEntity2)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Rel":
			return sTyped.Rel, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (f *fwTestEntity2) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &fwTestEntity2{}
	}
}

func (f *fwTestEntity2) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*fwTestEntity2)
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
		case "Data":
			eTyped.Data = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (f *fwTestEntity3) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*fwTestEntity3)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: f.__d3_makeFieldExtractor(),
			FieldSetter:    f.__d3_makeFieldSetter(),
			Instantiator:   f.__d3_makeInstantiator(),
		},
	}
}

func (f *fwTestEntity3) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*fwTestEntity3)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Rel":
			return sTyped.Rel, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (f *fwTestEntity3) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &fwTestEntity3{}
	}
}

func (f *fwTestEntity3) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*fwTestEntity3)
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
		case "Data":
			eTyped.Data = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}

func (f *fwTestEntity4) D3Token() entity.MetaToken {
	return entity.MetaToken{
		Tpl:       (*fwTestEntity4)(nil),
		TableName: "",
		Tools: entity.InternalTools{
			FieldExtractor: f.__d3_makeFieldExtractor(),
			FieldSetter:    f.__d3_makeFieldSetter(),
			Instantiator:   f.__d3_makeInstantiator(),
		},
	}
}

func (f *fwTestEntity4) __d3_makeFieldExtractor() entity.FieldExtractor {
	return func(s interface{}, name string) (interface{}, error) {
		sTyped, ok := s.(*fwTestEntity4)
		if !ok {
			return nil, fmt.Errorf("invalid entity type")
		}

		switch name {

		case "Id":
			return sTyped.Id, nil

		case "Data":
			return sTyped.Data, nil

		default:
			return nil, fmt.Errorf("field %s not found", name)
		}
	}
}

func (f *fwTestEntity4) __d3_makeInstantiator() entity.Instantiator {
	return func() interface{} {
		return &fwTestEntity4{}
	}
}

func (f *fwTestEntity4) __d3_makeFieldSetter() entity.FieldSetter {
	return func(s interface{}, name string, val interface{}) error {
		eTyped, ok := s.(*fwTestEntity4)
		if !ok {
			return fmt.Errorf("invalid entity type")
		}

		switch name {
		case "Id":
			eTyped.Id = val.(int32)
			return nil
		case "Data":
			eTyped.Data = val.(string)
			return nil

		default:
			return fmt.Errorf("field %s not found", name)
		}
	}
}
