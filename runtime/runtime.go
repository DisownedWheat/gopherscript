package runtime

type CsGoObj interface {
	GetType() string
	GetValue() interface{}
}

type CsGoValue struct {
	_type string
}

func (c *CsGoValue) GetType() string {
	return c._type
}

type CsGoString struct {
	value string
	_type string
}

func (c *CsGoString) GetType() string {
	return "string"
}

func (c *CsGoString) GetValue() string {
	return c.value
}

type CsGoInt struct {
	value int64
	_type string
}

func (c *CsGoInt) GetType() string {
	return "int"
}

func (c *CsGoInt) GetValue() int64 {
	return c.value
}

type CsGoMap struct {
	value map[string]interface{}
	_type string
}

func (c *CsGoMap) GetType() string {
	return "obj"
}
func (c *CsGoMap) GetValue() map[string]interface{} {
	return c.value
}
