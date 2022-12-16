package object

type DatabaseConnection struct {
	Conn interface{}
}

func (c *DatabaseConnection) Type() Type {
	return DB_CONNECTION
}

func (c *DatabaseConnection) Inspect() string {
	return "<DB_CONNECTION>"
}

func (c *DatabaseConnection) InvokeMethod(method string, args ...Object) Object {
	return NewError("type error: %s object has no method %s", c.Type(), method)
}

func (c *DatabaseConnection) ToInterface() interface{} {
	return c.Conn
}

func (c *DatabaseConnection) Equals(other Object) Object {
	value := other.Type() == DB_CONNECTION && c.Conn == other.(*DatabaseConnection).Conn
	return NewBool(value)
}

func (c *DatabaseConnection) GetAttr(name string) (Object, bool) {
	return nil, false
}
