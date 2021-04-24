package varuint

// Maximum number of bytes required to encode uint64.
const MaxBufSize = 9

// Encode encodes a value into buf and returns the number of bytes written.
// Buffer must have enough space to encode the value
func Encode(buf []byte, v uint64) (n int) {
	switch {
	case v <= 240:
		buf[0] = byte(v)
		n = 1
	case v <= 2287:
		v240 := v - 240
		buf[0] = byte(v240>>8) + 241
		buf[1] = byte(v240)
		n = 2
	case v <= 67823:
		buf[0] = 249
		v2288 := v - 2288
		buf[1] = byte(v2288 >> 8)
		buf[2] = byte(v2288)
		n = 3
	case v <= 1<<24-1:
		buf[0] = 250
		buf[3] = byte(v)
		shl := v >> 8
		buf[2] = byte(shl)
		buf[1] = byte(shl >> 8)
		n = 4
	case v <= 1<<32-1:
		buf[0] = 251
		buf[4] = byte(v)
		shl := v >> 8
		buf[3] = byte(shl)
		shl >>= 8
		buf[2] = byte(shl)
		buf[1] = byte(shl >> 8)
		n = 5
	case v <= 1<<40-1:
		buf[0] = 252
		buf[5] = byte(v)
		shl := v >> 8
		buf[4] = byte(shl)
		shl >>= 8
		buf[3] = byte(shl)
		shl >>= 8
		buf[2] = byte(shl)
		buf[1] = byte(shl >> 8)
		n = 6
	case v <= 1<<48-1:
		buf[0] = 253
		buf[6] = byte(v)
		shl := v >> 8
		buf[5] = byte(shl)
		shl >>= 8
		buf[4] = byte(shl)
		shl >>= 8
		buf[3] = byte(shl)
		shl >>= 8
		buf[2] = byte(shl)
		buf[1] = byte(shl >> 8)
		n = 7
	case v <= 1<<56-1:
		buf[0] = 254
		buf[7] = byte(v)
		shl := v >> 8
		buf[6] = byte(shl)
		shl >>= 8
		buf[5] = byte(shl)
		shl >>= 8
		buf[4] = byte(shl)
		shl >>= 8
		buf[3] = byte(shl)
		shl >>= 8
		buf[2] = byte(shl)
		buf[1] = byte(shl >> 8)
		n = 8
	default:
		buf[0] = 255
		buf[8] = byte(v)
		shl := v >> 8
		buf[7] = byte(shl)
		shl >>= 8
		buf[6] = byte(shl)
		shl >>= 8
		buf[5] = byte(shl)
		shl >>= 8
		buf[4] = byte(shl)
		shl >>= 8
		buf[3] = byte(shl)
		shl >>= 8
		buf[2] = byte(shl)
		buf[1] = byte(shl >> 8)
		n = 9
	}

	return
}

// EncodeSize returns size of bytes required to encode value
func EncodeSize(v uint64) (s int) {
	switch {
	case v <= 240:
		s = 1
	case v <= 2287:
		s = 2
	case v <= 67823:
		s = 3
	case v <= 1<<24-1:
		s = 4
	case v <= 1<<32-1:
		s = 5
	case v <= 1<<40-1:
		s = 6
	case v <= 1<<48-1:
		s = 7
	case v <= 1<<56-1:
		s = 8
	default:
		s = 9
	}

	return
}

// IsDecodable returns true if value can be decoded
func IsDecodable(buf []byte) (is bool) {
	l := len(buf)

	if l == 0 {
		return
	}

	v0 := buf[0]

	if v0 <= 240 {
		return true
	}
	if v0 <= 248 {
		return l >= 2
	}

	return l >= int(v0-246)
}

// Decode returns a decoded uint64 from buf and the number of bytes read.
// Buf must have enough space to decode the number
func Decode(buf []byte) (v uint64, n int) {
	v0 := buf[0]

	if v0 <= 240 {
		v = uint64(v0)
		n = 1
		return
	}

	if v0 <= 248 {
		v = ((uint64(v0) - 241) << 8) + uint64(buf[1]) + 240
		n = 2
		return
	}

	switch v0 {
	case 249:
		v = 2288 +
			uint64(buf[1])<<8 +
			uint64(buf[2])
		n = 3
	case 250:
		v = uint64(buf[1])<<16 |
			uint64(buf[2])<<8 |
			uint64(buf[3])
		n = 4
	case 251:
		v = uint64(buf[1])<<24 |
			uint64(buf[2])<<16 |
			uint64(buf[3])<<8 |
			uint64(buf[4])
		n = 5
	case 252:
		v = uint64(buf[1])<<32 |
			uint64(buf[2])<<24 |
			uint64(buf[3])<<16 |
			uint64(buf[4])<<8 |
			uint64(buf[5])
		n = 6
	case 253:
		v = uint64(buf[1])<<40 |
			uint64(buf[2])<<32 |
			uint64(buf[3])<<24 |
			uint64(buf[4])<<16 |
			uint64(buf[5])<<8 |
			uint64(buf[6])
		n = 7
	case 254:
		v = uint64(buf[1])<<48 |
			uint64(buf[2])<<40 |
			uint64(buf[3])<<32 |
			uint64(buf[4])<<24 |
			uint64(buf[5])<<16 |
			uint64(buf[6])<<8 |
			uint64(buf[7])
		n = 8
	case 255:
		v = uint64(buf[1])<<56 |
			uint64(buf[2])<<48 |
			uint64(buf[3])<<40 |
			uint64(buf[4])<<32 |
			uint64(buf[5])<<24 |
			uint64(buf[6])<<16 |
			uint64(buf[7])<<8 |
			uint64(buf[8])
		n = 9
	}

	return
}
