package caproto

import "fmt"

func NewEmptyExtensions() *Extensions {
	return &Extensions{
		IsEmpty:    true,
		Extensions: nil,
	}
}

func (e *Extensions) SetExtensionsRoot() {
	e.IsEmpty = false
	e.Extensions = make(map[string]*Extension)
}

func (e *Extensions) Get(key string) (interface{}, bool, error) {
	if e.IsEmpty {
		return nil, false, fmt.Errorf("Extensions is empty")
	}

	if val, ok := e.Extensions[key]; ok {
		return val, true, nil
	}

	return nil, false, nil
}

func (e *Extensions) Set(key string, value *Extension) {
	if e.IsEmpty {
		e.SetExtensionsRoot()
	}

	e.Extensions[key] = value
}

func (e *Extension) Get(key string) (interface{}, bool, error) {
	if e.Type != ExtensionType_EXTENSION {
		return nil, false, fmt.Errorf("Extension is not a map")
	}

	if e.ExtensionValue == nil {
		return nil, false, fmt.Errorf("Extension type is Extension but Extension is nil")
	}

	extMap := e.GetExtensionValue().GetExtensions()

	if val, ok := extMap[key]; ok {
		return val, true, nil
	}

	return nil, false, nil
}

func MakeNewExtension(value interface{}) (*Extension, error) {
	switch v := value.(type) {
	case map[string]*Extension:
		return &Extension{
			Type: ExtensionType_EXTENSION,
			ExtensionValue: &Extensions{
				IsEmpty:    false,
				Extensions: v,
			},
		}, nil
	case []*Extension:
		return &Extension{
			Type:       ExtensionType_ARRAY,
			ArrayValue: &ExtensionArray{Values: v},
		}, nil
	case []interface{}:
		exts := make([]*Extension, len(v))
		for i, val := range v {
			ext, err := MakeNewExtension(val)
			if err != nil {
				return nil, err
			}
			exts[i] = ext
		}
		return &Extension{
			Type:       ExtensionType_ARRAY,
			ArrayValue: &ExtensionArray{Values: exts},
		}, nil
	case string:
		return &Extension{
			Type:        ExtensionType_STRING,
			StringValue: StringValue(v),
		}, nil
	case int64:
		return &Extension{
			Type:         ExtensionType_INTEGER,
			IntegerValue: &v,
		}, nil
	case bool:
		return &Extension{
			Type:         ExtensionType_BOOLEAN,
			BooleanValue: &v,
		}, nil
	case []byte:
		return &Extension{
			Type:       ExtensionType_BYTES,
			BytesValue: v,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported type for Extension")
	}
}
