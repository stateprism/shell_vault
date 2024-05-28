package protocol

import "fmt"

func NewEmptyExtensions() *Extensions {
	return &Extensions{
		IsEmpty:    true,
		Extensions: nil,
	}
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
