package serialization

import msgpack "gopkg.in/vmihailenco/msgpack.v2"

type MsgpackProvider struct {
	SerializationFactory
}

func (provider *MsgpackProvider) Marshal(v ...interface{}) ([]byte, error) {
	bytes, err := msgpack.Marshal(v...)
	return bytes, err
}

func (provider *MsgpackProvider) Unmarshal(data []byte, v ...interface{}) error {
	err := msgpack.Unmarshal(data, v...)
	return err
}
