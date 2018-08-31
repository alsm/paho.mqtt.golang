package packets

import (
	"bytes"
	"fmt"
	"io"
)

// PropPayloadFormat, etc are the list of property codes for the
// MQTT packet properties
const (
	PropPayloadFormat          byte = 1
	PropMessageExpiry               = 2
	PropContentType                 = 3
	PropResponseTopic               = 8
	PropCorrelationData             = 9
	PropSubscriptionIdentifier      = 11
	PropSessionExpiryInterval       = 17
	PropAssignedClientID            = 18
	PropServerKeepAlive             = 19
	PropAuthMethod                  = 21
	PropAuthData                    = 22
	PropRequestProblemInfo          = 23
	PropWillDelayInterval           = 24
	PropRequestResponseInfo         = 25
	PropResponseInfo                = 26
	PropServerReference             = 28
	PropReasonString                = 31
	PropReceiveMaximum              = 33
	PropTopicAliasMaximum           = 34
	PropTopicAlias                  = 35
	PropMaximumQOS                  = 36
	PropRetainAvailable             = 37
	PropUser                        = 38
	PropMaximumPacketSize           = 39
	PropWildcardSubAvailable        = 40
	PropSubIDAvailable              = 41
	PropSharedSubAvailable          = 42
)

// Properties is a struct representing the all the described properties
// allowed by the MQTT protocol, determining the validity of a property
// relvative to the packettype it was received in is provided by the
// ValidateID function
type Properties struct {
	// PayloadFormat indicates the format of the payload of the message
	// 0 is unspecified bytes
	// 1 is UTF8 encoded character data
	PayloadFormat *byte
	// MessageExpiry is the lifetime of the message in seconds
	MessageExpiry *uint32
	// ContentType is a UTF8 string describing the content of the message
	// for example it could be a MIME type
	ContentType string
	// ResponseTopic is a UTF8 string indicating the topic name to which any
	// response to this message should be sent
	ResponseTopic string
	// CorrelationData is binary data used to associate future response
	// messages with the original request message
	CorrelationData []byte
	// SubscriptionIdentifier is an identifier of the subscription to which
	// the Publish matched
	SubscriptionIdentifier *uint32
	// SessionExpiryInterval is the time in seconds after a client disconnects
	// that the server should retain the session information (subscriptions etc)
	SessionExpiryInterval *uint32
	// AssignedClientID is the server assigned client identifier in the case
	// that a client connected without specifying a clientID the server
	// generates one and returns it in the Connack
	AssignedClientID string
	// ServerKeepAlive allows the server to specify in the Connack packet
	// the time in seconds to be used as the keep alive value
	ServerKeepAlive *uint16
	// AuthMethod is a UTF8 string containing the name of the authentication
	// method to be used for extended authentication
	AuthMethod string
	// AuthData is binary data containing authentication data
	AuthData []byte
	// RequestProblemInfo is used by the Client to indicate to the server to
	// include the Reason String and/or User Properties in case of failures
	RequestProblemInfo *byte
	// WillDelayInterval is the number of seconds the server waits after the
	// point at which it would otherwise send the will message before sending
	// it. The client reconnecting before that time expires causes the server
	// to cancel sending the will
	WillDelayInterval *uint32
	// RequestResponseInfo is used by the Client to request the Server provide
	// Response Information in the Connack
	RequestResponseInfo *byte
	// ResponseInfo is a UTF8 encoded string that can be used as the basis for
	// createing a Response Topic. The way in which the Client creates a
	// Response Topic from the Response Information is not defined. A common
	// use of this is to pass a globally unique portion of the topic tree which
	// is reserved for this Client for at least the lifetime of its Session. This
	// often cannot just be a random name as both the requesting Client and the
	// responding Client need to be authorized to use it. It is normal to use this
	// as the root of a topic tree for a particular Client. For the Server to
	// return this information, it normally needs to be correctly configured.
	// Using this mechanism allows this configuration to be done once in the
	// Server rather than in each Client
	ResponseInfo string
	// ServerReference is a UTF8 string indicating another server the client
	// can use
	ServerReference string
	// ReasonString is a UTF8 string representing the reason associated with
	// this response, intended to be human readable for diagnostic purposes
	ReasonString string
	// ReceiveMaximum is the maximum number of QOS1 & 2 messages allowed to be
	// 'inflight' (not having received a PUBACK/PUBCOMP response for)
	ReceiveMaximum *uint16
	// TopicAliasMaximum is the highest value permitted as a Topic Alias
	TopicAliasMaximum *uint16
	// TopicAlias is used in place of the topic string to reduce the size of
	// packets for repeated messages on a topic
	TopicAlias *uint16
	// MaximumQOS is the highest QOS level permitted for a Publish
	MaximumQOS *byte
	// RetainAvailable indicates whether the server supports messages with the
	// retain flag set
	RetainAvailable *byte
	// User is a map of user provided properties
	User map[string]string
	// MaximumPacketSize allows the client or server to specify the maximum packet
	// size in bytes that they support
	MaximumPacketSize *uint32
	// WildcardSubAvailable indicates whether wildcard subscriptions are permitted
	WildcardSubAvailable *byte
	// SubIDAvailable indicates whether subscription identifiers are supported
	SubIDAvailable *byte
	// SharedSubAvailable indicates whether shared subscriptions are supported
	SharedSubAvailable *byte
}

// NewProperties creates a new Properties and applies all the
// provided/listed option functions to configure them
func NewProperties(opts ...func(*Properties)) *Properties {
	p := &Properties{
		User: make(map[string]string),
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// PayloadFormat is a Properties option function that sets the
// payload format for a Properties struct
func PayloadFormat(x byte) func(*Properties) {
	return func(i *Properties) {
		i.PayloadFormat = &x
	}
}

// MessageExpiry is a Properties option function that sets the
// pub expiry for a Properties struct
func MessageExpiry(x uint32) func(*Properties) {
	return func(i *Properties) {
		i.MessageExpiry = &x
	}
}

// ContentType is a Properties option function that sets the
// content type for a Properties struct
func ContentType(x string) func(*Properties) {
	return func(i *Properties) {
		i.ContentType = x
	}
}

// ResponseTopic is a Properties option function that sets the
// Response topic for a Properties struct
func ResponseTopic(x string) func(*Properties) {
	return func(i *Properties) {
		i.ResponseTopic = x
	}
}

// CorrelationData is a Properties option function that sets the
// correlation data for a Properties struct
func CorrelationData(x []byte) func(*Properties) {
	return func(i *Properties) {
		i.CorrelationData = x
	}
}

// SubscriptionIdentifier is a Properties option function that sets the
// subscription identifier for a Properties struct
func SubscriptionIdentifier(x *uint32) func(*Properties) {
	return func(i *Properties) {
		i.SubscriptionIdentifier = x
	}
}

// SessionExpiryInterval is a Properties option function that sets the
// session expiry interval for a Properties struct
func SessionExpiryInterval(x *uint32) func(*Properties) {
	return func(i *Properties) {
		i.SessionExpiryInterval = x
	}
}

// AssignedClientID is a Properties option function that sets the
// assigned client id for a Properties struct
func AssignedClientID(x string) func(*Properties) {
	return func(i *Properties) {
		i.AssignedClientID = x
	}
}

// ServerKeepAlive is a Properties option function that sets the
// server keep alive for a Properties struct
func ServerKeepAlive(x *uint16) func(*Properties) {
	return func(i *Properties) {
		i.ServerKeepAlive = x
	}
}

// AuthMethod is a Properties option function that sets the
// auth method for a Properties struct
func AuthMethod(x string) func(*Properties) {
	return func(i *Properties) {
		i.AuthMethod = x
	}
}

// AuthData is a Properties option function that sets the
// auth data for a Properties struct
func AuthData(x []byte) func(*Properties) {
	return func(i *Properties) {
		i.AuthData = x
	}
}

// RequestProblemInfo is a Properties option function that sets the
// request problem info for a Properties struct
func RequestProblemInfo(x *byte) func(*Properties) {
	return func(i *Properties) {
		i.RequestProblemInfo = x
	}
}

// WillDelayInterval is a Properties option function that sets the
// will delay interval for a Properties struct
func WillDelayInterval(x *uint32) func(*Properties) {
	return func(i *Properties) {
		i.WillDelayInterval = x
	}
}

// RequestResponseInfo is a Properties option function that sets the
// request response info for a Properties struct
func RequestResponseInfo(x *byte) func(*Properties) {
	return func(i *Properties) {
		i.RequestResponseInfo = x
	}
}

// ResponseInfo is a Properties option function that sets the
// response info for a Properties struct
func ResponseInfo(x string) func(*Properties) {
	return func(i *Properties) {
		i.ResponseInfo = x
	}
}

// ServerReference is a Properties option function that sets the
// server reference for a Properties struct
func ServerReference(x string) func(*Properties) {
	return func(i *Properties) {
		i.ServerReference = x
	}
}

// ReasonString is a Properties option function that sets the
// reason string for a Properties struct
func ReasonString(x string) func(*Properties) {
	return func(i *Properties) {
		i.ReasonString = x
	}
}

// ReceiveMaximum is a Properties option function that sets the
// receive maximum for a Properties struct
func ReceiveMaximum(x *uint16) func(*Properties) {
	return func(i *Properties) {
		i.ReceiveMaximum = x
	}
}

// TopicAliasMaximum is a Properties option function that sets the
// topic alias maximum for a Properties struct
func TopicAliasMaximum(x *uint16) func(*Properties) {
	return func(i *Properties) {
		i.TopicAliasMaximum = x
	}
}

// TopicAlias is a Properties option function that sets the
// topic alias for a Properties struct
func TopicAlias(x *uint16) func(*Properties) {
	return func(i *Properties) {
		i.TopicAlias = x
	}
}

// MaximumQOS is a Properties option function that sets the
// maximum qos for a Properties struct
func MaximumQOS(x *byte) func(*Properties) {
	return func(i *Properties) {
		i.MaximumQOS = x
	}
}

// RetainAvailable is a Properties option function that sets the
// retain available for a Properties struct
func RetainAvailable(x *byte) func(*Properties) {
	return func(i *Properties) {
		i.RetainAvailable = x
	}
}

// UserMap is a Properties option function that sets the
// user properties to be the values in the provided map
func UserMap(x map[string]string) func(*Properties) {
	return func(i *Properties) {
		for k, v := range x {
			i.User[k] = v
		}
	}
}

// UserSingle is a Properties option function that sets the
// a single key/value property in the user properties
func UserSingle(k, v string) func(*Properties) {
	return func(i *Properties) {
		i.User[k] = v
	}
}

// MaximumPacketSize is a Properties option function that sets the
// maximum packet size for a Properties struct
func MaximumPacketSize(x *uint32) func(*Properties) {
	return func(i *Properties) {
		i.MaximumPacketSize = x
	}
}

// WildcardSubAvailable is a Properties option function that sets the
// wildcard sub available for a Properties struct
func WildcardSubAvailable(x *byte) func(*Properties) {
	return func(i *Properties) {
		i.WildcardSubAvailable = x
	}
}

// SubIDAvailable is a Properties option function that sets the
// sub id available for a Properties struct
func SubIDAvailable(x *byte) func(*Properties) {
	return func(i *Properties) {
		i.SubIDAvailable = x
	}
}

// SharedSubAvailable is a Properties option function that sets the
// shared sub available for a Properties struct
func SharedSubAvailable(x *byte) func(*Properties) {
	return func(i *Properties) {
		i.SharedSubAvailable = x
	}
}

// Pack takes all the defined properties for an Properties and produces
// a slice of bytes representing the wire format for the information
func (i *Properties) Pack(p PacketType) []byte {
	var b bytes.Buffer

	if i == nil {
		return nil
	}

	if p == PUBLISH {
		if i.PayloadFormat != nil {
			b.WriteByte(PropPayloadFormat)
			b.WriteByte(*i.PayloadFormat)
		}

		if i.MessageExpiry != nil {
			b.WriteByte(PropMessageExpiry)
			writeUint32(*i.MessageExpiry, &b)
		}

		if i.ContentType != "" {
			b.WriteByte(PropContentType)
			writeString(i.ContentType, &b)
		}

		if i.ResponseTopic != "" {
			b.WriteByte(PropResponseTopic)
			writeString(i.ResponseTopic, &b)
		}

		if i.CorrelationData != nil && len(i.CorrelationData) > 0 {
			b.WriteByte(PropCorrelationData)
			b.Write(i.CorrelationData)
		}

		if i.TopicAlias != nil {
			b.WriteByte(PropTopicAlias)
			writeUint16(*i.TopicAlias, &b)
		}
	}

	if p == PUBLISH || p == SUBSCRIBE {
		if i.SubscriptionIdentifier != nil {
			b.WriteByte(PropSubscriptionIdentifier)
			writeUint32(*i.SubscriptionIdentifier, &b)
		}
	}

	if p == CONNECT || p == CONNACK {
		if i.ReceiveMaximum != nil {
			b.WriteByte(PropReceiveMaximum)
			writeUint16(*i.ReceiveMaximum, &b)
		}

		if i.TopicAliasMaximum != nil {
			b.WriteByte(PropTopicAliasMaximum)
			writeUint16(*i.TopicAliasMaximum, &b)
		}

		if i.MaximumQOS != nil {
			b.WriteByte(PropMaximumQOS)
			b.WriteByte(*i.MaximumQOS)
		}

		if i.MaximumPacketSize != nil {
			b.WriteByte(PropMaximumPacketSize)
			writeUint32(*i.MaximumPacketSize, &b)
		}
	}

	if p == CONNACK {
		if i.AssignedClientID != "" {
			b.WriteByte(PropAssignedClientID)
			writeString(i.AssignedClientID, &b)
		}

		if i.ServerKeepAlive != nil {
			b.WriteByte(PropServerKeepAlive)
			writeUint16(*i.ServerKeepAlive, &b)
		}

		if i.WildcardSubAvailable != nil {
			b.WriteByte(PropWildcardSubAvailable)
			b.WriteByte(*i.WildcardSubAvailable)
		}

		if i.SubIDAvailable != nil {
			b.WriteByte(PropSubIDAvailable)
			b.WriteByte(*i.SubIDAvailable)
		}

		if i.SharedSubAvailable != nil {
			b.WriteByte(PropSharedSubAvailable)
			b.WriteByte(*i.SharedSubAvailable)
		}

		if i.RetainAvailable != nil {
			b.WriteByte(PropRetainAvailable)
			b.WriteByte(*i.RetainAvailable)
		}

		if i.ResponseInfo != "" {
			b.WriteByte(PropResponseInfo)
			writeString(i.ResponseInfo, &b)
		}
	}

	if p == CONNECT {
		if i.RequestProblemInfo != nil {
			b.WriteByte(PropRequestProblemInfo)
			b.WriteByte(*i.RequestProblemInfo)
		}

		if i.WillDelayInterval != nil {
			b.WriteByte(PropWillDelayInterval)
			writeUint32(*i.WillDelayInterval, &b)
		}

		if i.RequestResponseInfo != nil {
			b.WriteByte(PropRequestResponseInfo)
			b.WriteByte(*i.RequestResponseInfo)
		}
	}

	if p == CONNECT || p == DISCONNECT {
		if i.SessionExpiryInterval != nil {
			b.WriteByte(PropSessionExpiryInterval)
			writeUint32(*i.SessionExpiryInterval, &b)
		}
	}

	if p == CONNECT || p == CONNACK || p == AUTH {
		if i.AuthMethod != "" {
			b.WriteByte(PropAuthMethod)
			writeString(i.AuthMethod, &b)
		}

		if i.AuthData != nil && len(i.AuthData) > 0 {
			b.WriteByte(PropAuthData)
			b.Write(i.AuthData)
		}
	}

	if p == CONNACK || p == DISCONNECT {
		if i.ServerReference != "" {
			b.WriteByte(PropServerReference)
			writeString(i.ServerReference, &b)
		}
	}

	if p != CONNECT {
		if i.ReasonString != "" {
			b.WriteByte(PropReasonString)
			writeString(i.ReasonString, &b)
		}
	}

	for k, v := range i.User {
		b.WriteByte(PropUser)
		writeString(k, &b)
		writeString(v, &b)
	}

	return b.Bytes()
}

// Unpack takes a buffer of bytes and reads out the defined properties
// filling in the appropriate entries in the struct, it returns the number
// of bytes used to store the Prop data and any error in decoding them
func (i *Properties) Unpack(r *bytes.Buffer, p PacketType) error {
	vbi, err := getVBI(r)
	if err != nil {
		return err
	}
	size, err := decodeVBI(vbi)
	if err != nil {
		return err
	}
	if size == 0 {
		return nil
	}

	buf := bytes.NewBuffer(r.Next(size))
	for {
		PropType, err := buf.ReadByte()
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		if !ValidateID(p, PropType) {
			return fmt.Errorf("Invalid Prop type %d for packet %d", PropType, p)
		}
		switch PropType {
		case PropPayloadFormat:
			pf, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.PayloadFormat = &pf
		case PropMessageExpiry:
			pe, err := readUint32(buf)
			if err != nil {
				return err
			}
			i.MessageExpiry = &pe
		case PropContentType:
			ct, err := readString(buf)
			if err != nil {
				return err
			}
			i.ContentType = ct
		case PropResponseTopic:
			tr, err := readString(buf)
			if err != nil {
				return err
			}
			i.ResponseTopic = tr
		case PropCorrelationData:
			cd, err := readBinary(buf)
			if err != nil {
				return err
			}
			i.CorrelationData = cd
		case PropSubscriptionIdentifier:
			si, err := readUint32(buf)
			if err != nil {
				return err
			}
			i.SubscriptionIdentifier = &si
		case PropSessionExpiryInterval:
			se, err := readUint32(buf)
			if err != nil {
				return err
			}
			i.SessionExpiryInterval = &se
		case PropAssignedClientID:
			ac, err := readString(buf)
			if err != nil {
				return err
			}
			i.AssignedClientID = ac
		case PropServerKeepAlive:
			sk, err := readUint16(buf)
			if err != nil {
				return err
			}
			i.ServerKeepAlive = &sk
		case PropAuthMethod:
			am, err := readString(buf)
			if err != nil {
				return err
			}
			i.AuthMethod = am
		case PropAuthData:
			ad, err := readBinary(buf)
			if err != nil {
				return err
			}
			i.AuthData = ad
		case PropRequestProblemInfo:
			rp, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.RequestProblemInfo = &rp
		case PropWillDelayInterval:
			wd, err := readUint32(buf)
			if err != nil {
				return err
			}
			i.WillDelayInterval = &wd
		case PropRequestResponseInfo:
			rp, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.RequestResponseInfo = &rp
		case PropResponseInfo:
			ri, err := readString(buf)
			if err != nil {
				return err
			}
			i.ResponseInfo = ri
		case PropServerReference:
			sr, err := readString(buf)
			if err != nil {
				return err
			}
			i.ServerReference = sr
		case PropReasonString:
			rs, err := readString(buf)
			if err != nil {
				return err
			}
			i.ReasonString = rs
		case PropReceiveMaximum:
			rm, err := readUint16(buf)
			if err != nil {
				return err
			}
			i.ReceiveMaximum = &rm
		case PropTopicAliasMaximum:
			ta, err := readUint16(buf)
			if err != nil {
				return err
			}
			i.TopicAliasMaximum = &ta
		case PropTopicAlias:
			ta, err := readUint16(buf)
			if err != nil {
				return err
			}
			i.TopicAlias = &ta
		case PropMaximumQOS:
			mq, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.MaximumQOS = &mq
		case PropRetainAvailable:
			ra, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.RetainAvailable = &ra
		case PropUser:
			k, err := readString(buf)
			if err != nil {
				return err
			}
			v, err := readString(buf)
			if err != nil {
				return err
			}
			i.User[k] = v
		case PropMaximumPacketSize:
			mp, err := readUint32(buf)
			if err != nil {
				return err
			}
			i.MaximumPacketSize = &mp
		case PropWildcardSubAvailable:
			ws, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.WildcardSubAvailable = &ws
		case PropSubIDAvailable:
			si, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.SubIDAvailable = &si
		case PropSharedSubAvailable:
			ss, err := buf.ReadByte()
			if err != nil {
				return err
			}
			i.SharedSubAvailable = &ss
		default:
			return fmt.Errorf("Unknown Prop type %d", PropType)
		}
	}

	return nil
}

// ValidProperties is a map of the various properties and the
// PacketTypes that property is valid for.
var ValidProperties = map[byte]map[PacketType]struct{}{
	PropPayloadFormat:          {PUBLISH: {}},
	PropMessageExpiry:          {PUBLISH: {}},
	PropContentType:            {PUBLISH: {}},
	PropResponseTopic:          {PUBLISH: {}},
	PropCorrelationData:        {PUBLISH: {}},
	PropTopicAlias:             {PUBLISH: {}},
	PropSubscriptionIdentifier: {PUBLISH: {}, SUBSCRIBE: {}},
	PropSessionExpiryInterval:  {CONNECT: {}, DISCONNECT: {}},
	PropAssignedClientID:       {CONNACK: {}},
	PropServerKeepAlive:        {CONNACK: {}},
	PropWildcardSubAvailable:   {CONNACK: {}},
	PropSubIDAvailable:         {CONNACK: {}},
	PropSharedSubAvailable:     {CONNACK: {}},
	PropRetainAvailable:        {CONNACK: {}},
	PropResponseInfo:           {CONNACK: {}},
	PropAuthMethod:             {CONNECT: {}, CONNACK: {}, AUTH: {}},
	PropAuthData:               {CONNECT: {}, CONNACK: {}, AUTH: {}},
	PropRequestProblemInfo:     {CONNECT: {}},
	PropWillDelayInterval:      {CONNECT: {}},
	PropRequestResponseInfo:    {CONNECT: {}},
	PropServerReference:        {CONNACK: {}, DISCONNECT: {}},
	PropReasonString:           {CONNACK: {}, PUBACK: {}, PUBREC: {}, PUBREL: {}, PUBCOMP: {}, SUBACK: {}, UNSUBACK: {}, DISCONNECT: {}, AUTH: {}},
	PropReceiveMaximum:         {CONNECT: {}, CONNACK: {}},
	PropTopicAliasMaximum:      {CONNECT: {}, CONNACK: {}},
	PropMaximumQOS:             {CONNECT: {}, CONNACK: {}},
	PropMaximumPacketSize:      {CONNECT: {}, CONNACK: {}},
	PropUser:                   {CONNECT: {}, CONNACK: {}, PUBLISH: {}, PUBACK: {}, PUBREC: {}, PUBREL: {}, PUBCOMP: {}, SUBSCRIBE: {}, UNSUBSCRIBE: {}, SUBACK: {}, UNSUBACK: {}, DISCONNECT: {}, AUTH: {}},
}

// ValidateID takes a PacketType and a property name and returns
// a boolean indicating if that property is valid for that
// PacketType
func ValidateID(p PacketType, i byte) bool {
	_, ok := ValidProperties[i][p]
	return ok
}
