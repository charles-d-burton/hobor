package encoder

type GetAPIMessage struct {
	Message string `cbor:"1,keyasint" json:"message"`
}
