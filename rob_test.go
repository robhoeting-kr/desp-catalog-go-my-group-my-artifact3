// Code generated by github.com/actgardner/gogen-avro/v8. DO NOT EDIT.
/*
 * SOURCES:
 *     EventHeader.avsc
 *     RobTest.avsc
 */
package my_group_my_artifact3

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/actgardner/gogen-avro/v9/compiler"
	"github.com/actgardner/gogen-avro/v9/vm"
	"github.com/actgardner/gogen-avro/v9/vm/types"
)

var _ = fmt.Printf

type RobTest struct {
	// Example: This is a sample attribute, please edit.
	Sample string `json:"sample"`

	Sample2 *UnionNullString `json:"sample2"`

	Sample3 *UnionNullString `json:"sample3"`

	EventHeader *UnionNullEventHeader `json:"eventHeader"`
}

const RobTestAvroCRC64Fingerprint = "j\xe3\xb5q\xacT\xbb\v"

func NewRobTest() RobTest {
	r := RobTest{}
	r.Sample2 = NewUnionNullString()

	r.Sample2 = nil
	r.Sample3 = NewUnionNullString()

	r.Sample3 = nil
	r.EventHeader = NewUnionNullEventHeader()

	r.EventHeader = nil
	return r
}

func DeserializeRobTest(r io.Reader) (RobTest, error) {
	t := NewRobTest()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func DeserializeRobTestFromSchema(r io.Reader, schema string) (RobTest, error) {
	t := NewRobTest()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return t, err
	}

	err = vm.Eval(r, deser, &t)
	return t, err
}

func writeRobTest(r RobTest, w io.Writer) error {
	var err error
	err = vm.WriteString(r.Sample, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sample2, w)
	if err != nil {
		return err
	}
	err = writeUnionNullString(r.Sample3, w)
	if err != nil {
		return err
	}
	err = writeUnionNullEventHeader(r.EventHeader, w)
	if err != nil {
		return err
	}
	return err
}

func (r RobTest) Serialize(w io.Writer) error {
	return writeRobTest(r, w)
}

func (r RobTest) Schema() string {
	return "{\"fields\":[{\"doc\":\"Example: This is a sample attribute, please edit.\",\"name\":\"sample\",\"type\":\"string\"},{\"default\":null,\"doc\":\"\",\"name\":\"sample2\",\"type\":[\"null\",{\"avro.java.string\":\"String\",\"type\":\"string\"}]},{\"default\":null,\"doc\":\"\",\"name\":\"sample3\",\"type\":[\"null\",{\"avro.java.string\":\"String\",\"type\":\"string\"}]},{\"default\":null,\"doc\":\"\",\"name\":\"eventHeader\",\"type\":[\"null\",{\"doc\":\"The below fields include header information and should be included on every event in the DESP. Inspired by: https://github.com/cloudevents/spec/blob/v0.2/spec.md\",\"fields\":[{\"doc\":\"A unique identifier of the event - for example, a randomly generated GUID\",\"name\":\"id\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Time the event occurred in milliseconds since epoch, UTC timezone.\",\"name\":\"time\",\"type\":\"long\"},{\"doc\":\"Type of occurrence which has happened. Reference the domain.event registered in schema-registry.\",\"name\":\"type\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Service that produced the event. Future: reference to producer registry.\",\"name\":\"source\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}},{\"doc\":\"Service that produced the event. Future: reference to producer registry.\",\"name\":\"source3\",\"type\":{\"avro.java.string\":\"String\",\"type\":\"string\"}}],\"name\":\"EventHeader\",\"namespace\":\"com.kroger.desp.commons.desp.healthcheck.rph\",\"type\":\"record\"}]}],\"name\":\"com.kroger.desp.events.desp.healthcheck.rph.RobTest\",\"type\":\"record\"}"
}

func (r RobTest) SchemaName() string {
	return "com.kroger.desp.events.desp.healthcheck.rph.RobTest"
}

func (_ RobTest) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ RobTest) SetInt(v int32)       { panic("Unsupported operation") }
func (_ RobTest) SetLong(v int64)      { panic("Unsupported operation") }
func (_ RobTest) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ RobTest) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ RobTest) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ RobTest) SetString(v string)   { panic("Unsupported operation") }
func (_ RobTest) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RobTest) Get(i int) types.Field {
	switch i {
	case 0:
		return &types.String{Target: &r.Sample}
	case 1:
		r.Sample2 = NewUnionNullString()

		return r.Sample2
	case 2:
		r.Sample3 = NewUnionNullString()

		return r.Sample3
	case 3:
		r.EventHeader = NewUnionNullEventHeader()

		return r.EventHeader
	}
	panic("Unknown field index")
}

func (r *RobTest) SetDefault(i int) {
	switch i {
	case 1:
		r.Sample2 = nil
		return
	case 2:
		r.Sample3 = nil
		return
	case 3:
		r.EventHeader = nil
		return
	}
	panic("Unknown field index")
}

func (r *RobTest) NullField(i int) {
	switch i {
	case 1:
		r.Sample2 = nil
		return
	case 2:
		r.Sample3 = nil
		return
	case 3:
		r.EventHeader = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ RobTest) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ RobTest) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ RobTest) Finalize()                        {}

func (_ RobTest) AvroCRC64Fingerprint() []byte {
	return []byte(RobTestAvroCRC64Fingerprint)
}

func (r RobTest) MarshalJSON() ([]byte, error) {
	var err error
	output := make(map[string]json.RawMessage)
	output["sample"], err = json.Marshal(r.Sample)
	if err != nil {
		return nil, err
	}
	output["sample2"], err = json.Marshal(r.Sample2)
	if err != nil {
		return nil, err
	}
	output["sample3"], err = json.Marshal(r.Sample3)
	if err != nil {
		return nil, err
	}
	output["eventHeader"], err = json.Marshal(r.EventHeader)
	if err != nil {
		return nil, err
	}
	return json.Marshal(output)
}

func (r *RobTest) UnmarshalJSON(data []byte) error {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}

	var val json.RawMessage
	val = func() json.RawMessage {
		if v, ok := fields["sample"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sample); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("no value specified for sample")
	}
	val = func() json.RawMessage {
		if v, ok := fields["sample2"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sample2); err != nil {
			return err
		}
	} else {
		r.Sample2 = NewUnionNullString()

		r.Sample2 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["sample3"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.Sample3); err != nil {
			return err
		}
	} else {
		r.Sample3 = NewUnionNullString()

		r.Sample3 = nil
	}
	val = func() json.RawMessage {
		if v, ok := fields["eventHeader"]; ok {
			return v
		}
		return nil
	}()

	if val != nil {
		if err := json.Unmarshal([]byte(val), &r.EventHeader); err != nil {
			return err
		}
	} else {
		r.EventHeader = NewUnionNullEventHeader()

		r.EventHeader = nil
	}
	return nil
}
