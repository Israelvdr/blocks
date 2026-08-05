package golfers

const (
	ELEMENT_ALWAYS_FIRST string = "element is always first"
	ELEMENT_NULL         string = "null element"
	ELEMENT_STANDARD     string = "standard element"
)

type elementCharacteristic interface {
	Characteristic() string
	Equals(elementCharacteristic) bool
}

type characteristicAlwaysFirst struct {
}

func (b characteristicAlwaysFirst) Characteristic() string {
	return ELEMENT_ALWAYS_FIRST
}
func (b characteristicAlwaysFirst) Equals(comp elementCharacteristic) bool {
	_, ok := comp.(characteristicAlwaysFirst)
	return ok
}

type characteristicNull struct {
}

func (b characteristicNull) Characteristic() string {
	return ELEMENT_NULL
}
func (b characteristicNull) Equals(comp elementCharacteristic) bool {
	_, ok := comp.(characteristicNull)
	return ok
}

type characteristicStandard struct {
}

func (b characteristicStandard) Characteristic() string {
	return ELEMENT_STANDARD
}
func (b characteristicStandard) Equals(comp elementCharacteristic) bool {
	_, ok := comp.(characteristicStandard)
	return ok
}
