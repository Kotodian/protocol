package dtc

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType, v)
	}
	*j = CostKindEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType(plain)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateFirmwareStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UpdateFirmwareStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UpdateFirmwareStatusEnumType_1, v)
	}
	*j = UpdateFirmwareStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateFirmwareStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UpdateFirmwareStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UpdateFirmwareStatusEnumType, v)
	}
	*j = UpdateFirmwareStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType, v)
	}
	*j = HashAlgorithmEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_41) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_41
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_41(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_127) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_127
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_127(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateFirmwareRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["firmware"]; !ok || v == nil {
		return fmt.Errorf("field firmware: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain UpdateFirmwareRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UpdateFirmwareRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FirmwareType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["location"]; !ok || v == nil {
		return fmt.Errorf("field location: required")
	}
	if v, ok := raw["retrieveDateTime"]; !ok || v == nil {
		return fmt.Errorf("field retrieveDateTime: required")
	}
	type Plain FirmwareType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FirmwareType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_126) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_126
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_126(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType, v)
	}
	*j = IdTokenEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnpublishFirmwareResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain UnpublishFirmwareResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UnpublishFirmwareResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnpublishFirmwareStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UnpublishFirmwareStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UnpublishFirmwareStatusEnumType_1, v)
	}
	*j = UnpublishFirmwareStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnpublishFirmwareStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UnpublishFirmwareStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UnpublishFirmwareStatusEnumType, v)
	}
	*j = UnpublishFirmwareStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_125) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_125
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_125(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnpublishFirmwareRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["checksum"]; !ok || v == nil {
		return fmt.Errorf("field checksum: required")
	}
	type Plain UnpublishFirmwareRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UnpublishFirmwareRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_124) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_124
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_124(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnlockConnectorResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain UnlockConnectorResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UnlockConnectorResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnlockStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UnlockStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UnlockStatusEnumType_1, v)
	}
	*j = UnlockStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnlockStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UnlockStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UnlockStatusEnumType, v)
	}
	*j = UnlockStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_40) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_40
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_40(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_1, v)
	}
	*j = IdTokenEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_123) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_123
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_123(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnlockConnectorRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["connectorId"]; !ok || v == nil {
		return fmt.Errorf("field connectorId: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	type Plain UnlockConnectorRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UnlockConnectorRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_122) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_122
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_122(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerMessageResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain TriggerMessageResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TriggerMessageResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerMessageStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TriggerMessageStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TriggerMessageStatusEnumType_1, v)
	}
	*j = TriggerMessageStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerMessageStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TriggerMessageStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TriggerMessageStatusEnumType, v)
	}
	*j = TriggerMessageStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_39) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_39
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_39(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_121) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_121
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_121(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerMessageRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["requestedMessage"]; !ok || v == nil {
		return fmt.Errorf("field requestedMessage: required")
	}
	type Plain TriggerMessageRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TriggerMessageRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageTriggerEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageTriggerEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageTriggerEnumType_1, v)
	}
	*j = MessageTriggerEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageTriggerEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageTriggerEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageTriggerEnumType, v)
	}
	*j = MessageTriggerEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_1, v)
	}
	*j = HashAlgorithmEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_15) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_15
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_15(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_120) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_120
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_120(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_9, v)
	}
	*j = MessageFormatEnumType_9(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenInfoType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain IdTokenInfoType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenInfoType_2(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCSPRequestDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hashAlgorithm"]; !ok || v == nil {
		return fmt.Errorf("field hashAlgorithm: required")
	}
	if v, ok := raw["issuerKeyHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerKeyHash: required")
	}
	if v, ok := raw["issuerNameHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerNameHash: required")
	}
	if v, ok := raw["responderURL"]; !ok || v == nil {
		return fmt.Errorf("field responderURL: required")
	}
	if v, ok := raw["serialNumber"]; !ok || v == nil {
		return fmt.Errorf("field serialNumber: required")
	}
	type Plain OCSPRequestDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = OCSPRequestDataType(plain)
	return nil
}

type AuthorizeRequestJson struct {
	// The X.509 certificated presented by EV and encoded in PEM format.
	//
	Certificate *string `json:"certificate,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken IdTokenType `json:"idToken"`

	// Iso15118CertificateHashData corresponds to the JSON schema field
	// "iso15118CertificateHashData".
	Iso15118CertificateHashData []OCSPRequestDataType `json:"iso15118CertificateHashData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizeRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	type Plain AuthorizeRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuthorizeRequestJson(plain)
	return nil
}

const AuthorizationStatusEnumType_5_Unknown AuthorizationStatusEnumType_5 = "Unknown"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_1(plain)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType_1 struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_1 `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType_1(plain)
	return nil
}

type AuthorizationStatusEnumType string

const AuthorizationStatusEnumType_5_NotAtThisTime AuthorizationStatusEnumType_5 = "NotAtThisTime"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType, v)
	}
	*j = AuthorizationStatusEnumType(v)
	return nil
}

const AuthorizationStatusEnumTypeAccepted AuthorizationStatusEnumType = "Accepted"
const AuthorizationStatusEnumTypeBlocked AuthorizationStatusEnumType = "Blocked"
const AuthorizationStatusEnumTypeConcurrentTx AuthorizationStatusEnumType = "ConcurrentTx"
const AuthorizationStatusEnumTypeExpired AuthorizationStatusEnumType = "Expired"
const AuthorizationStatusEnumTypeInvalid AuthorizationStatusEnumType = "Invalid"
const AuthorizationStatusEnumTypeNoCredit AuthorizationStatusEnumType = "NoCredit"
const AuthorizationStatusEnumTypeNotAllowedTypeEVSE AuthorizationStatusEnumType = "NotAllowedTypeEVSE"
const AuthorizationStatusEnumTypeNotAtThisLocation AuthorizationStatusEnumType = "NotAtThisLocation"
const AuthorizationStatusEnumTypeNotAtThisTime AuthorizationStatusEnumType = "NotAtThisTime"
const AuthorizationStatusEnumTypeUnknown AuthorizationStatusEnumType = "Unknown"

type AuthorizeCertificateStatusEnumType string

const AuthorizationStatusEnumType_5_NotAtThisLocation AuthorizationStatusEnumType_5 = "NotAtThisLocation"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizeCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizeCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizeCertificateStatusEnumType, v)
	}
	*j = AuthorizeCertificateStatusEnumType(v)
	return nil
}

const AuthorizeCertificateStatusEnumTypeAccepted AuthorizeCertificateStatusEnumType = "Accepted"
const AuthorizeCertificateStatusEnumTypeSignatureError AuthorizeCertificateStatusEnumType = "SignatureError"
const AuthorizeCertificateStatusEnumTypeCertificateExpired AuthorizeCertificateStatusEnumType = "CertificateExpired"
const AuthorizeCertificateStatusEnumTypeCertificateRevoked AuthorizeCertificateStatusEnumType = "CertificateRevoked"
const AuthorizeCertificateStatusEnumTypeNoCertificateAvailable AuthorizeCertificateStatusEnumType = "NoCertificateAvailable"
const AuthorizeCertificateStatusEnumTypeCertChainError AuthorizeCertificateStatusEnumType = "CertChainError"
const AuthorizeCertificateStatusEnumTypeContractCancelled AuthorizeCertificateStatusEnumType = "ContractCancelled"
const AuthorizationStatusEnumType_5_NotAllowedTypeEVSE AuthorizationStatusEnumType_5 = "NotAllowedTypeEVSE"
const AuthorizationStatusEnumType_5_NoCredit AuthorizationStatusEnumType_5 = "NoCredit"

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_2, v)
	}
	*j = IdTokenEnumType_2(v)
	return nil
}

const AuthorizationStatusEnumType_5_Invalid AuthorizationStatusEnumType_5 = "Invalid"
const AuthorizationStatusEnumType_5_Expired AuthorizationStatusEnumType_5 = "Expired"
const AuthorizationStatusEnumType_5_ConcurrentTx AuthorizationStatusEnumType_5 = "ConcurrentTx"
const AuthorizationStatusEnumType_5_Blocked AuthorizationStatusEnumType_5 = "Blocked"
const AuthorizationStatusEnumType_5_Accepted AuthorizationStatusEnumType_5 = "Accepted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType_5, v)
	}
	*j = AuthorizationStatusEnumType_5(v)
	return nil
}

type AuthorizationStatusEnumType_5 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageContentType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["content"]; !ok || v == nil {
		return fmt.Errorf("field content: required")
	}
	if v, ok := raw["format"]; !ok || v == nil {
		return fmt.Errorf("field format: required")
	}
	type Plain MessageContentType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageContentType_4(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_8, v)
	}
	*j = MessageFormatEnumType_8(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType_7) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType_7
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType_7(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_3, v)
	}
	*j = IdTokenEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_15) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_15 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_15, v)
	}
	*j = IdTokenEnumType_15(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_14) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_14 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_14, v)
	}
	*j = IdTokenEnumType_14(v)
	return nil
}

const AuthorizationStatusEnumType_4_Unknown AuthorizationStatusEnumType_4 = "Unknown"
const AuthorizationStatusEnumType_4_NotAtThisTime AuthorizationStatusEnumType_4 = "NotAtThisTime"
const AuthorizationStatusEnumType_4_NotAtThisLocation AuthorizationStatusEnumType_4 = "NotAtThisLocation"
const AuthorizationStatusEnumType_4_NotAllowedTypeEVSE AuthorizationStatusEnumType_4 = "NotAllowedTypeEVSE"
const AuthorizationStatusEnumType_4_NoCredit AuthorizationStatusEnumType_4 = "NoCredit"
const AuthorizationStatusEnumType_4_Invalid AuthorizationStatusEnumType_4 = "Invalid"
const AuthorizationStatusEnumType_4_Expired AuthorizationStatusEnumType_4 = "Expired"

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType_1(plain)
	return nil
}

const AuthorizationStatusEnumType_4_ConcurrentTx AuthorizationStatusEnumType_4 = "ConcurrentTx"
const AuthorizationStatusEnumType_4_Blocked AuthorizationStatusEnumType_4 = "Blocked"

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType, v)
	}
	*j = MessageFormatEnumType(v)
	return nil
}

const AuthorizationStatusEnumType_4_Accepted AuthorizationStatusEnumType_4 = "Accepted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType_4, v)
	}
	*j = AuthorizationStatusEnumType_4(v)
	return nil
}

type AuthorizationStatusEnumType_4 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType_7) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType_7
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType_7(plain)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType_7 struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_119 `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageContentType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["content"]; !ok || v == nil {
		return fmt.Errorf("field content: required")
	}
	if v, ok := raw["format"]; !ok || v == nil {
		return fmt.Errorf("field format: required")
	}
	type Plain MessageContentType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageContentType(plain)
	return nil
}

type AuthorizationStatusEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_119) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_119
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_119(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType_1, v)
	}
	*j = AuthorizationStatusEnumType_1(v)
	return nil
}

const AuthorizationStatusEnumType_1_Accepted AuthorizationStatusEnumType_1 = "Accepted"
const AuthorizationStatusEnumType_1_Blocked AuthorizationStatusEnumType_1 = "Blocked"
const AuthorizationStatusEnumType_1_ConcurrentTx AuthorizationStatusEnumType_1 = "ConcurrentTx"
const AuthorizationStatusEnumType_1_Expired AuthorizationStatusEnumType_1 = "Expired"
const AuthorizationStatusEnumType_1_Invalid AuthorizationStatusEnumType_1 = "Invalid"
const AuthorizationStatusEnumType_1_NoCredit AuthorizationStatusEnumType_1 = "NoCredit"
const AuthorizationStatusEnumType_1_NotAllowedTypeEVSE AuthorizationStatusEnumType_1 = "NotAllowedTypeEVSE"
const AuthorizationStatusEnumType_1_NotAtThisLocation AuthorizationStatusEnumType_1 = "NotAtThisLocation"
const AuthorizationStatusEnumType_1_NotAtThisTime AuthorizationStatusEnumType_1 = "NotAtThisTime"
const AuthorizationStatusEnumType_1_Unknown AuthorizationStatusEnumType_1 = "Unknown"

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransactionEventRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["eventType"]; !ok || v == nil {
		return fmt.Errorf("field eventType: required")
	}
	if v, ok := raw["seqNo"]; !ok || v == nil {
		return fmt.Errorf("field seqNo: required")
	}
	if v, ok := raw["timestamp"]; !ok || v == nil {
		return fmt.Errorf("field timestamp: required")
	}
	if v, ok := raw["transactionInfo"]; !ok || v == nil {
		return fmt.Errorf("field transactionInfo: required")
	}
	if v, ok := raw["triggerReason"]; !ok || v == nil {
		return fmt.Errorf("field triggerReason: required")
	}
	type Plain TransactionEventRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["offline"]; !ok || v == nil {
		plain.Offline = false
	}
	*j = TransactionEventRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain IdTokenInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenInfoType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerReasonEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TriggerReasonEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TriggerReasonEnumType_1, v)
	}
	*j = TriggerReasonEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransactionEventEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransactionEventEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransactionEventEnumType_1, v)
	}
	*j = TransactionEventEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_1, v)
	}
	*j = MessageFormatEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TriggerReasonEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TriggerReasonEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TriggerReasonEnumType, v)
	}
	*j = TriggerReasonEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransactionType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["transactionId"]; !ok || v == nil {
		return fmt.Errorf("field transactionId: required")
	}
	type Plain TransactionType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TransactionType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReasonEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReasonEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReasonEnumType_1, v)
	}
	*j = ReasonEnumType_1(v)
	return nil
}

const ChargingStateEnumType_1_Idle ChargingStateEnumType_1 = "Idle"

type AuthorizeCertificateStatusEnumType_1 string

const ChargingStateEnumType_1_SuspendedEVSE ChargingStateEnumType_1 = "SuspendedEVSE"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizeCertificateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizeCertificateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizeCertificateStatusEnumType_1, v)
	}
	*j = AuthorizeCertificateStatusEnumType_1(v)
	return nil
}

const AuthorizeCertificateStatusEnumType_1_Accepted AuthorizeCertificateStatusEnumType_1 = "Accepted"
const AuthorizeCertificateStatusEnumType_1_SignatureError AuthorizeCertificateStatusEnumType_1 = "SignatureError"
const AuthorizeCertificateStatusEnumType_1_CertificateExpired AuthorizeCertificateStatusEnumType_1 = "CertificateExpired"
const AuthorizeCertificateStatusEnumType_1_CertificateRevoked AuthorizeCertificateStatusEnumType_1 = "CertificateRevoked"
const AuthorizeCertificateStatusEnumType_1_NoCertificateAvailable AuthorizeCertificateStatusEnumType_1 = "NoCertificateAvailable"
const AuthorizeCertificateStatusEnumType_1_CertChainError AuthorizeCertificateStatusEnumType_1 = "CertChainError"
const AuthorizeCertificateStatusEnumType_1_ContractCancelled AuthorizeCertificateStatusEnumType_1 = "ContractCancelled"

type AuthorizeResponseJson struct {
	// CertificateStatus corresponds to the JSON schema field "certificateStatus".
	CertificateStatus *AuthorizeCertificateStatusEnumType_1 `json:"certificateStatus,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_1 `json:"customData,omitempty"`

	// IdTokenInfo corresponds to the JSON schema field "idTokenInfo".
	IdTokenInfo IdTokenInfoType `json:"idTokenInfo"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizeResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idTokenInfo"]; !ok || v == nil {
		return fmt.Errorf("field idTokenInfo: required")
	}
	type Plain AuthorizeResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuthorizeResponseJson(plain)
	return nil
}

type BootReasonEnumType string

const ChargingStateEnumType_1_SuspendedEV ChargingStateEnumType_1 = "SuspendedEV"

// UnmarshalJSON implements json.Unmarshaler.
func (j *BootReasonEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_BootReasonEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_BootReasonEnumType, v)
	}
	*j = BootReasonEnumType(v)
	return nil
}

const BootReasonEnumTypeApplicationReset BootReasonEnumType = "ApplicationReset"
const BootReasonEnumTypeFirmwareUpdate BootReasonEnumType = "FirmwareUpdate"
const BootReasonEnumTypeLocalReset BootReasonEnumType = "LocalReset"
const BootReasonEnumTypePowerUp BootReasonEnumType = "PowerUp"
const BootReasonEnumTypeRemoteReset BootReasonEnumType = "RemoteReset"
const BootReasonEnumTypeScheduledReset BootReasonEnumType = "ScheduledReset"
const BootReasonEnumTypeTriggered BootReasonEnumType = "Triggered"
const BootReasonEnumTypeUnknown BootReasonEnumType = "Unknown"
const BootReasonEnumTypeWatchdog BootReasonEnumType = "Watchdog"
const ChargingStateEnumType_1_EVConnected ChargingStateEnumType_1 = "EVConnected"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_2(plain)
	return nil
}

const ChargingStateEnumType_1_Charging ChargingStateEnumType_1 = "Charging"

// Charge_ Point
// urn:x-oca:ocpp:uid:2:233122
// The physical system where an Electrical Vehicle (EV) can be charged.
//
type ChargingStationType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_2 `json:"customData,omitempty"`

	// This contains the firmware version of the Charging Station.
	//
	//
	FirmwareVersion *string `json:"firmwareVersion,omitempty"`

	// Device. Model. CI20_ Text
	// urn:x-oca:ocpp:uid:1:569325
	// Defines the model of the device.
	//
	Model string `json:"model"`

	// Modem corresponds to the JSON schema field "modem".
	Modem *ModemType `json:"modem,omitempty"`

	// Device. Serial_ Number. Serial_ Number
	// urn:x-oca:ocpp:uid:1:569324
	// Vendor-specific device identifier.
	//
	SerialNumber *string `json:"serialNumber,omitempty"`

	// Identifies the vendor (not necessarily in a unique manner).
	//
	VendorName string `json:"vendorName"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingStationType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["model"]; !ok || v == nil {
		return fmt.Errorf("field model: required")
	}
	if v, ok := raw["vendorName"]; !ok || v == nil {
		return fmt.Errorf("field vendorName: required")
	}
	type Plain ChargingStationType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingStationType(plain)
	return nil
}

type BootReasonEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingStateEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingStateEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingStateEnumType_1, v)
	}
	*j = ChargingStateEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BootReasonEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_BootReasonEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_BootReasonEnumType_1, v)
	}
	*j = BootReasonEnumType_1(v)
	return nil
}

const BootReasonEnumType_1_ApplicationReset BootReasonEnumType_1 = "ApplicationReset"
const BootReasonEnumType_1_FirmwareUpdate BootReasonEnumType_1 = "FirmwareUpdate"
const BootReasonEnumType_1_LocalReset BootReasonEnumType_1 = "LocalReset"
const BootReasonEnumType_1_PowerUp BootReasonEnumType_1 = "PowerUp"
const BootReasonEnumType_1_RemoteReset BootReasonEnumType_1 = "RemoteReset"
const BootReasonEnumType_1_ScheduledReset BootReasonEnumType_1 = "ScheduledReset"
const BootReasonEnumType_1_Triggered BootReasonEnumType_1 = "Triggered"
const BootReasonEnumType_1_Unknown BootReasonEnumType_1 = "Unknown"
const BootReasonEnumType_1_Watchdog BootReasonEnumType_1 = "Watchdog"

type BootNotificationRequestJson struct {
	// ChargingStation corresponds to the JSON schema field "chargingStation".
	ChargingStation ChargingStationType `json:"chargingStation"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_2 `json:"customData,omitempty"`

	// Reason corresponds to the JSON schema field "reason".
	Reason BootReasonEnumType_1 `json:"reason"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BootNotificationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingStation"]; !ok || v == nil {
		return fmt.Errorf("field chargingStation: required")
	}
	if v, ok := raw["reason"]; !ok || v == nil {
		return fmt.Errorf("field reason: required")
	}
	type Plain BootNotificationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BootNotificationRequestJson(plain)
	return nil
}

type ChargingStateEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TransactionEventEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TransactionEventEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TransactionEventEnumType, v)
	}
	*j = TransactionEventEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReasonEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReasonEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReasonEnumType, v)
	}
	*j = ReasonEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RegistrationStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RegistrationStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RegistrationStatusEnumType, v)
	}
	*j = RegistrationStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReadingContextEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReadingContextEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReadingContextEnumType_3, v)
	}
	*j = ReadingContextEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PhaseEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PhaseEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PhaseEnumType_3, v)
	}
	*j = PhaseEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeterValueType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["sampledValue"]; !ok || v == nil {
		return fmt.Errorf("field sampledValue: required")
	}
	if v, ok := raw["timestamp"]; !ok || v == nil {
		return fmt.Errorf("field timestamp: required")
	}
	type Plain MeterValueType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MeterValueType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SampledValueType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value: required")
	}
	type Plain SampledValueType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SampledValueType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnitOfMeasureType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain UnitOfMeasureType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["multiplier"]; !ok || v == nil {
		plain.Multiplier = 0
	}
	if v, ok := raw["unit"]; !ok || v == nil {
		plain.Unit = "Wh"
	}
	*j = UnitOfMeasureType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignedMeterValueType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["encodingMethod"]; !ok || v == nil {
		return fmt.Errorf("field encodingMethod: required")
	}
	if v, ok := raw["publicKey"]; !ok || v == nil {
		return fmt.Errorf("field publicKey: required")
	}
	if v, ok := raw["signedMeterData"]; !ok || v == nil {
		return fmt.Errorf("field signedMeterData: required")
	}
	if v, ok := raw["signingMethod"]; !ok || v == nil {
		return fmt.Errorf("field signingMethod: required")
	}
	type Plain SignedMeterValueType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SignedMeterValueType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RegistrationStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RegistrationStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RegistrationStatusEnumType_1, v)
	}
	*j = RegistrationStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PhaseEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PhaseEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PhaseEnumType_2, v)
	}
	*j = PhaseEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeasurandEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MeasurandEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MeasurandEnumType_3, v)
	}
	*j = MeasurandEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LocationEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LocationEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LocationEnumType_3, v)
	}
	*j = LocationEnumType_3(v)
	return nil
}

type BootNotificationResponseJson struct {
	// This contains the CSMS’s current time.
	//
	CurrentTime string `json:"currentTime"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_3 `json:"customData,omitempty"`

	// When &lt;&lt;cmn_registrationstatusenumtype,Status&gt;&gt; is Accepted, this
	// contains the heartbeat interval in seconds. If the CSMS returns something other
	// than Accepted, the value of the interval field indicates the minimum wait time
	// before sending a next BootNotification request.
	//
	Interval int `json:"interval"`

	// Status corresponds to the JSON schema field "status".
	Status RegistrationStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *BootNotificationResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["currentTime"]; !ok || v == nil {
		return fmt.Errorf("field currentTime: required")
	}
	if v, ok := raw["interval"]; !ok || v == nil {
		return fmt.Errorf("field interval: required")
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain BootNotificationResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = BootNotificationResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReadingContextEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReadingContextEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReadingContextEnumType_2, v)
	}
	*j = ReadingContextEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_4(plain)
	return nil
}

type CancelReservationRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_4 `json:"customData,omitempty"`

	// Id of the reservation to cancel.
	//
	ReservationId int `json:"reservationId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CancelReservationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reservationId"]; !ok || v == nil {
		return fmt.Errorf("field reservationId: required")
	}
	type Plain CancelReservationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CancelReservationRequestJson(plain)
	return nil
}

type CancelReservationStatusEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeasurandEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MeasurandEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MeasurandEnumType_2, v)
	}
	*j = MeasurandEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CancelReservationStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CancelReservationStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CancelReservationStatusEnumType, v)
	}
	*j = CancelReservationStatusEnumType(v)
	return nil
}

const CancelReservationStatusEnumTypeAccepted CancelReservationStatusEnumType = "Accepted"
const CancelReservationStatusEnumTypeRejected CancelReservationStatusEnumType = "Rejected"

// UnmarshalJSON implements json.Unmarshaler.
func (j *LocationEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LocationEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LocationEnumType_2, v)
	}
	*j = LocationEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_5(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType_6) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType_6
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType_6(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_1(plain)
	return nil
}

type CancelReservationStatusEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_13) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_13 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_13, v)
	}
	*j = IdTokenEnumType_13(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CancelReservationStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CancelReservationStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CancelReservationStatusEnumType_1, v)
	}
	*j = CancelReservationStatusEnumType_1(v)
	return nil
}

const CancelReservationStatusEnumType_1_Accepted CancelReservationStatusEnumType_1 = "Accepted"
const CancelReservationStatusEnumType_1_Rejected CancelReservationStatusEnumType_1 = "Rejected"

type CancelReservationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_5 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status CancelReservationStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_1 `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CancelReservationResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain CancelReservationResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CancelReservationResponseJson(plain)
	return nil
}

type CertificateSigningUseEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_12) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_12 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_12, v)
	}
	*j = IdTokenEnumType_12(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSigningUseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateSigningUseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateSigningUseEnumType, v)
	}
	*j = CertificateSigningUseEnumType(v)
	return nil
}

const CertificateSigningUseEnumTypeChargingStationCertificate CertificateSigningUseEnumType = "ChargingStationCertificate"
const CertificateSigningUseEnumTypeV2GCertificate CertificateSigningUseEnumType = "V2GCertificate"

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_14) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_14
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_14(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_6) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_6
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_6(plain)
	return nil
}

type CertificateSigningUseEnumType_1 string

const ChargingStateEnumTypeIdle ChargingStateEnumType = "Idle"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSigningUseEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateSigningUseEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateSigningUseEnumType_1, v)
	}
	*j = CertificateSigningUseEnumType_1(v)
	return nil
}

const CertificateSigningUseEnumType_1_ChargingStationCertificate CertificateSigningUseEnumType_1 = "ChargingStationCertificate"
const CertificateSigningUseEnumType_1_V2GCertificate CertificateSigningUseEnumType_1 = "V2GCertificate"

type CertificateSignedRequestJson struct {
	// The signed PEM encoded X.509 certificate. This can also contain the necessary
	// sub CA certificates. In that case, the order of the bundle should follow the
	// certificate chain, starting from the leaf certificate.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-max-certificate-chain-size,MaxCertificateChainSize&gt;&gt;
	// can be used to limit the maximum size of this field.
	//
	CertificateChain string `json:"certificateChain"`

	// CertificateType corresponds to the JSON schema field "certificateType".
	CertificateType *CertificateSigningUseEnumType_1 `json:"certificateType,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_6 `json:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSignedRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["certificateChain"]; !ok || v == nil {
		return fmt.Errorf("field certificateChain: required")
	}
	type Plain CertificateSignedRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CertificateSignedRequestJson(plain)
	return nil
}

type CertificateSignedStatusEnumType string

const ChargingStateEnumTypeSuspendedEVSE ChargingStateEnumType = "SuspendedEVSE"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSignedStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateSignedStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateSignedStatusEnumType, v)
	}
	*j = CertificateSignedStatusEnumType(v)
	return nil
}

const CertificateSignedStatusEnumTypeAccepted CertificateSignedStatusEnumType = "Accepted"
const CertificateSignedStatusEnumTypeRejected CertificateSignedStatusEnumType = "Rejected"
const ChargingStateEnumTypeSuspendedEV ChargingStateEnumType = "SuspendedEV"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_7) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_7
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_7(plain)
	return nil
}

const ChargingStateEnumTypeEVConnected ChargingStateEnumType = "EVConnected"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_2(plain)
	return nil
}

type CertificateSignedStatusEnumType_1 string

const ChargingStateEnumTypeCharging ChargingStateEnumType = "Charging"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSignedStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateSignedStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateSignedStatusEnumType_1, v)
	}
	*j = CertificateSignedStatusEnumType_1(v)
	return nil
}

const CertificateSignedStatusEnumType_1_Accepted CertificateSignedStatusEnumType_1 = "Accepted"
const CertificateSignedStatusEnumType_1_Rejected CertificateSignedStatusEnumType_1 = "Rejected"

type CertificateSignedResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_7 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status CertificateSignedStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_2 `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSignedResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain CertificateSignedResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CertificateSignedResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingStateEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingStateEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingStateEnumType, v)
	}
	*j = ChargingStateEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_8) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_8
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_8(plain)
	return nil
}

type ChargingStateEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType_6) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType_6
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType_6(plain)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType_6 struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OperationalStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OperationalStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OperationalStatusEnumType, v)
	}
	*j = OperationalStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_118) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_118
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_118(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_117) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_117
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_117(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusNotificationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["connectorId"]; !ok || v == nil {
		return fmt.Errorf("field connectorId: required")
	}
	if v, ok := raw["connectorStatus"]; !ok || v == nil {
		return fmt.Errorf("field connectorStatus: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	if v, ok := raw["timestamp"]; !ok || v == nil {
		return fmt.Errorf("field timestamp: required")
	}
	type Plain StatusNotificationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusNotificationRequestJson(plain)
	return nil
}

const ConnectorStatusEnumType_1_Faulted ConnectorStatusEnumType_1 = "Faulted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *OperationalStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OperationalStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OperationalStatusEnumType_1, v)
	}
	*j = OperationalStatusEnumType_1(v)
	return nil
}

const ConnectorStatusEnumType_1_Unavailable ConnectorStatusEnumType_1 = "Unavailable"
const ConnectorStatusEnumType_1_Reserved ConnectorStatusEnumType_1 = "Reserved"

type ChangeAvailabilityRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_8 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType `json:"evse,omitempty"`

	// OperationalStatus corresponds to the JSON schema field "operationalStatus".
	OperationalStatus OperationalStatusEnumType_1 `json:"operationalStatus"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChangeAvailabilityRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["operationalStatus"]; !ok || v == nil {
		return fmt.Errorf("field operationalStatus: required")
	}
	type Plain ChangeAvailabilityRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChangeAvailabilityRequestJson(plain)
	return nil
}

type ChangeAvailabilityStatusEnumType string

const ConnectorStatusEnumType_1_Occupied ConnectorStatusEnumType_1 = "Occupied"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChangeAvailabilityStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChangeAvailabilityStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChangeAvailabilityStatusEnumType, v)
	}
	*j = ChangeAvailabilityStatusEnumType(v)
	return nil
}

const ChangeAvailabilityStatusEnumTypeAccepted ChangeAvailabilityStatusEnumType = "Accepted"
const ChangeAvailabilityStatusEnumTypeRejected ChangeAvailabilityStatusEnumType = "Rejected"
const ChangeAvailabilityStatusEnumTypeScheduled ChangeAvailabilityStatusEnumType = "Scheduled"
const ConnectorStatusEnumType_1_Available ConnectorStatusEnumType_1 = "Available"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_9) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_9
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_9(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConnectorStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConnectorStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConnectorStatusEnumType_1, v)
	}
	*j = ConnectorStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_3(plain)
	return nil
}

type ChangeAvailabilityStatusEnumType_1 string

type ConnectorStatusEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChangeAvailabilityStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChangeAvailabilityStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChangeAvailabilityStatusEnumType_1, v)
	}
	*j = ChangeAvailabilityStatusEnumType_1(v)
	return nil
}

const ChangeAvailabilityStatusEnumType_1_Accepted ChangeAvailabilityStatusEnumType_1 = "Accepted"
const ChangeAvailabilityStatusEnumType_1_Rejected ChangeAvailabilityStatusEnumType_1 = "Rejected"
const ChangeAvailabilityStatusEnumType_1_Scheduled ChangeAvailabilityStatusEnumType_1 = "Scheduled"

type ChangeAvailabilityResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_9 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ChangeAvailabilityStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_3 `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChangeAvailabilityResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain ChangeAvailabilityResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChangeAvailabilityResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_116) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_116
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_116(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_10) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_10
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_10(plain)
	return nil
}

type ClearCacheRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_10 `json:"customData,omitempty"`
}

type ClearCacheStatusEnumType string

const ConnectorStatusEnumTypeFaulted ConnectorStatusEnumType = "Faulted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearCacheStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearCacheStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearCacheStatusEnumType, v)
	}
	*j = ClearCacheStatusEnumType(v)
	return nil
}

const ClearCacheStatusEnumTypeAccepted ClearCacheStatusEnumType = "Accepted"
const ClearCacheStatusEnumTypeRejected ClearCacheStatusEnumType = "Rejected"
const ConnectorStatusEnumTypeUnavailable ConnectorStatusEnumType = "Unavailable"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_11) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_11
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_11(plain)
	return nil
}

const ConnectorStatusEnumTypeReserved ConnectorStatusEnumType = "Reserved"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_4(plain)
	return nil
}

type ClearCacheStatusEnumType_1 string

const ConnectorStatusEnumTypeOccupied ConnectorStatusEnumType = "Occupied"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearCacheStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearCacheStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearCacheStatusEnumType_1, v)
	}
	*j = ClearCacheStatusEnumType_1(v)
	return nil
}

const ClearCacheStatusEnumType_1_Accepted ClearCacheStatusEnumType_1 = "Accepted"
const ClearCacheStatusEnumType_1_Rejected ClearCacheStatusEnumType_1 = "Rejected"

type ClearCacheResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_11 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ClearCacheStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_4 `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearCacheResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain ClearCacheResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearCacheResponseJson(plain)
	return nil
}

type ChargingProfilePurposeEnumType string

const ConnectorStatusEnumTypeAvailable ConnectorStatusEnumType = "Available"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType, v)
	}
	*j = ChargingProfilePurposeEnumType(v)
	return nil
}

const ChargingProfilePurposeEnumTypeChargingStationExternalConstraints ChargingProfilePurposeEnumType = "ChargingStationExternalConstraints"
const ChargingProfilePurposeEnumTypeChargingStationMaxProfile ChargingProfilePurposeEnumType = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumTypeTxDefaultProfile ChargingProfilePurposeEnumType = "TxDefaultProfile"
const ChargingProfilePurposeEnumTypeTxProfile ChargingProfilePurposeEnumType = "TxProfile"

type ChargingProfilePurposeEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConnectorStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConnectorStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConnectorStatusEnumType, v)
	}
	*j = ConnectorStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_1, v)
	}
	*j = ChargingProfilePurposeEnumType_1(v)
	return nil
}

const ChargingProfilePurposeEnumType_1_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_1 = "ChargingStationExternalConstraints"
const ChargingProfilePurposeEnumType_1_ChargingStationMaxProfile ChargingProfilePurposeEnumType_1 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_1_TxDefaultProfile ChargingProfilePurposeEnumType_1 = "TxDefaultProfile"
const ChargingProfilePurposeEnumType_1_TxProfile ChargingProfilePurposeEnumType_1 = "TxProfile"

type ConnectorStatusEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_12) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_12
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_12(plain)
	return nil
}

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of a ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
//
type ClearChargingProfileType struct {
	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose *ChargingProfilePurposeEnumType_1 `json:"chargingProfilePurpose,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_12 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Specifies the id of the EVSE for which to clear charging profiles. An evseId of
	// zero (0) specifies the charging profile for the overall Charging Station.
	// Absence of this parameter means the clearing applies to all charging profiles
	// that match the other criteria in the request.
	//
	//
	EvseId *int `json:"evseId,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Specifies the stackLevel for which charging profiles will be cleared, if they
	// meet the other criteria in the request.
	//
	StackLevel *int `json:"stackLevel,omitempty"`
}

type ClearChargingProfileRequestJson struct {
	// ChargingProfileCriteria corresponds to the JSON schema field
	// "chargingProfileCriteria".
	ChargingProfileCriteria *ClearChargingProfileType `json:"chargingProfileCriteria,omitempty"`

	// The Id of the charging profile to clear.
	//
	ChargingProfileId *int `json:"chargingProfileId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_12 `json:"customData,omitempty"`
}

type ClearChargingProfileStatusEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignCertificateResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain SignCertificateResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SignCertificateResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearChargingProfileStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearChargingProfileStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearChargingProfileStatusEnumType, v)
	}
	*j = ClearChargingProfileStatusEnumType(v)
	return nil
}

const ClearChargingProfileStatusEnumTypeAccepted ClearChargingProfileStatusEnumType = "Accepted"
const ClearChargingProfileStatusEnumTypeUnknown ClearChargingProfileStatusEnumType = "Unknown"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_9, v)
	}
	*j = GenericStatusEnumType_9(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_13) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_13
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_13(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_38) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_38
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_38(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_5(plain)
	return nil
}

type ClearChargingProfileStatusEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_8, v)
	}
	*j = GenericStatusEnumType_8(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearChargingProfileStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearChargingProfileStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearChargingProfileStatusEnumType_1, v)
	}
	*j = ClearChargingProfileStatusEnumType_1(v)
	return nil
}

const ClearChargingProfileStatusEnumType_1_Accepted ClearChargingProfileStatusEnumType_1 = "Accepted"
const ClearChargingProfileStatusEnumType_1_Unknown ClearChargingProfileStatusEnumType_1 = "Unknown"

type ClearChargingProfileResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_13 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ClearChargingProfileStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_5 `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearChargingProfileResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain ClearChargingProfileResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearChargingProfileResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_115) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_115
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_115(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_14) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_14
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_14(plain)
	return nil
}

type ClearDisplayMessageRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_14 `json:"customData,omitempty"`

	// Id of the message that SHALL be removed from the Charging Station.
	//
	Id int `json:"id"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearDisplayMessageRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain ClearDisplayMessageRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearDisplayMessageRequestJson(plain)
	return nil
}

type ClearMessageStatusEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignCertificateRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["csr"]; !ok || v == nil {
		return fmt.Errorf("field csr: required")
	}
	type Plain SignCertificateRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SignCertificateRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearMessageStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearMessageStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearMessageStatusEnumType, v)
	}
	*j = ClearMessageStatusEnumType(v)
	return nil
}

const ClearMessageStatusEnumTypeAccepted ClearMessageStatusEnumType = "Accepted"
const ClearMessageStatusEnumTypeUnknown ClearMessageStatusEnumType = "Unknown"
const CertificateSigningUseEnumType_3_V2GCertificate CertificateSigningUseEnumType_3 = "V2GCertificate"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_15) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_15
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_15(plain)
	return nil
}

const CertificateSigningUseEnumType_3_ChargingStationCertificate CertificateSigningUseEnumType_3 = "ChargingStationCertificate"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_6) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_6
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_6(plain)
	return nil
}

type ClearMessageStatusEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSigningUseEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateSigningUseEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateSigningUseEnumType_3, v)
	}
	*j = CertificateSigningUseEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearMessageStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearMessageStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearMessageStatusEnumType_1, v)
	}
	*j = ClearMessageStatusEnumType_1(v)
	return nil
}

const ClearMessageStatusEnumType_1_Accepted ClearMessageStatusEnumType_1 = "Accepted"
const ClearMessageStatusEnumType_1_Unknown ClearMessageStatusEnumType_1 = "Unknown"

type ClearDisplayMessageResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_15 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ClearMessageStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_6 `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearDisplayMessageResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain ClearDisplayMessageResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearDisplayMessageResponseJson(plain)
	return nil
}

type CertificateSigningUseEnumType_3 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_16) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_16
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_16(plain)
	return nil
}

type ClearVariableMonitoringRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_16 `json:"customData,omitempty"`

	// List of the monitors to be cleared, identified by there Id.
	//
	Id []int `json:"id"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearVariableMonitoringRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain ClearVariableMonitoringRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearVariableMonitoringRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_114) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_114
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_114(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_17) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_17
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_17(plain)
	return nil
}

type ClearMonitoringStatusEnumType string

const CertificateSigningUseEnumType_2_V2GCertificate CertificateSigningUseEnumType_2 = "V2GCertificate"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearMonitoringStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearMonitoringStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearMonitoringStatusEnumType, v)
	}
	*j = ClearMonitoringStatusEnumType(v)
	return nil
}

const ClearMonitoringStatusEnumTypeAccepted ClearMonitoringStatusEnumType = "Accepted"
const ClearMonitoringStatusEnumTypeRejected ClearMonitoringStatusEnumType = "Rejected"
const ClearMonitoringStatusEnumTypeNotFound ClearMonitoringStatusEnumType = "NotFound"
const CertificateSigningUseEnumType_2_ChargingStationCertificate CertificateSigningUseEnumType_2 = "ChargingStationCertificate"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_7) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_7
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_7(plain)
	return nil
}

type ClearMonitoringResultType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_17 `json:"customData,omitempty"`

	// Id of the monitor of which a clear was requested.
	//
	//
	Id int `json:"id"`

	// Status corresponds to the JSON schema field "status".
	Status ClearMonitoringStatusEnumType `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_7 `json:"statusInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearMonitoringResultType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain ClearMonitoringResultType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearMonitoringResultType(plain)
	return nil
}

type ClearMonitoringStatusEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateSigningUseEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateSigningUseEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateSigningUseEnumType_2, v)
	}
	*j = CertificateSigningUseEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearMonitoringStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ClearMonitoringStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ClearMonitoringStatusEnumType_1, v)
	}
	*j = ClearMonitoringStatusEnumType_1(v)
	return nil
}

const ClearMonitoringStatusEnumType_1_Accepted ClearMonitoringStatusEnumType_1 = "Accepted"
const ClearMonitoringStatusEnumType_1_Rejected ClearMonitoringStatusEnumType_1 = "Rejected"
const ClearMonitoringStatusEnumType_1_NotFound ClearMonitoringStatusEnumType_1 = "NotFound"

type ClearVariableMonitoringResponseJson struct {
	// ClearMonitoringResult corresponds to the JSON schema field
	// "clearMonitoringResult".
	ClearMonitoringResult []ClearMonitoringResultType `json:"clearMonitoringResult"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_17 `json:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearVariableMonitoringResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["clearMonitoringResult"]; !ok || v == nil {
		return fmt.Errorf("field clearMonitoringResult: required")
	}
	type Plain ClearVariableMonitoringResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearVariableMonitoringResponseJson(plain)
	return nil
}

type ChargingLimitSourceEnumType string

type CertificateSigningUseEnumType_2 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType, v)
	}
	*j = ChargingLimitSourceEnumType(v)
	return nil
}

const ChargingLimitSourceEnumTypeEMS ChargingLimitSourceEnumType = "EMS"
const ChargingLimitSourceEnumTypeOther ChargingLimitSourceEnumType = "Other"
const ChargingLimitSourceEnumTypeSO ChargingLimitSourceEnumType = "SO"
const ChargingLimitSourceEnumTypeCSO ChargingLimitSourceEnumType = "CSO"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariablesResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["setVariableResult"]; !ok || v == nil {
		return fmt.Errorf("field setVariableResult: required")
	}
	type Plain SetVariablesResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetVariablesResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_18) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_18
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_18(plain)
	return nil
}

type ChargingLimitSourceEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetVariableStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetVariableStatusEnumType_1, v)
	}
	*j = SetVariableStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType_1, v)
	}
	*j = ChargingLimitSourceEnumType_1(v)
	return nil
}

const ChargingLimitSourceEnumType_1_EMS ChargingLimitSourceEnumType_1 = "EMS"
const ChargingLimitSourceEnumType_1_Other ChargingLimitSourceEnumType_1 = "Other"
const ChargingLimitSourceEnumType_1_SO ChargingLimitSourceEnumType_1 = "SO"
const ChargingLimitSourceEnumType_1_CSO ChargingLimitSourceEnumType_1 = "CSO"

type ClearedChargingLimitRequestJson struct {
	// ChargingLimitSource corresponds to the JSON schema field "chargingLimitSource".
	ChargingLimitSource ChargingLimitSourceEnumType_1 `json:"chargingLimitSource"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_18 `json:"customData,omitempty"`

	// EVSE Identifier.
	//
	EvseId *int `json:"evseId,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ClearedChargingLimitRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingLimitSource"]; !ok || v == nil {
		return fmt.Errorf("field chargingLimitSource: required")
	}
	type Plain ClearedChargingLimitRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ClearedChargingLimitRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableResultType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["attributeStatus"]; !ok || v == nil {
		return fmt.Errorf("field attributeStatus: required")
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	type Plain SetVariableResultType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetVariableResultType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_19) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_19
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_19(plain)
	return nil
}

type ClearedChargingLimitResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_19 `json:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_10) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_10
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_10(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_20) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_20
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_20(plain)
	return nil
}

type CostUpdatedRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_20 `json:"customData,omitempty"`

	// Current total cost, based on the information known by the CSMS, of the
	// transaction including taxes. In the currency configured with the configuration
	// Variable: [&lt;&lt;configkey-currency, Currency&gt;&gt;]
	//
	//
	TotalCost float64 `json:"totalCost"`

	// Transaction Id of the transaction the current cost are asked for.
	//
	//
	TransactionId string `json:"transactionId"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostUpdatedRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["totalCost"]; !ok || v == nil {
		return fmt.Errorf("field totalCost: required")
	}
	if v, ok := raw["transactionId"]; !ok || v == nil {
		return fmt.Errorf("field transactionId: required")
	}
	type Plain CostUpdatedRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostUpdatedRequestJson(plain)
	return nil
}

const AttributeEnumType_9_MaxSet AttributeEnumType_9 = "MaxSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_21) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_21
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_21(plain)
	return nil
}

type CostUpdatedResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_21 `json:"customData,omitempty"`
}

const AttributeEnumType_9_MinSet AttributeEnumType_9 = "MinSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_22) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_22
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_22(plain)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType_2 struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_22 `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType_2(plain)
	return nil
}

const AttributeEnumType_9_Target AttributeEnumType_9 = "Target"
const AttributeEnumType_9_Actual AttributeEnumType_9 = "Actual"

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_2, v)
	}
	*j = HashAlgorithmEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_9, v)
	}
	*j = AttributeEnumType_9(v)
	return nil
}

type AttributeEnumType_9 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_37) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_37
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_37(plain)
	return nil
}

type CertificateHashDataType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_22 `json:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType_2 `json:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateHashDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hashAlgorithm"]; !ok || v == nil {
		return fmt.Errorf("field hashAlgorithm: required")
	}
	if v, ok := raw["issuerKeyHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerKeyHash: required")
	}
	if v, ok := raw["issuerNameHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerNameHash: required")
	}
	if v, ok := raw["serialNumber"]; !ok || v == nil {
		return fmt.Errorf("field serialNumber: required")
	}
	type Plain CertificateHashDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CertificateHashDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetVariableStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetVariableStatusEnumType, v)
	}
	*j = SetVariableStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_12) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_12
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_12(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_3, v)
	}
	*j = HashAlgorithmEnumType_3(v)
	return nil
}

// A physical or logical component
//
type ComponentType_12 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_113 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_13 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_13) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_13
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_13(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_113) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_113
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_113(plain)
	return nil
}

const AttributeEnumType_8_MaxSet AttributeEnumType_8 = "MaxSet"
const AttributeEnumType_8_MinSet AttributeEnumType_8 = "MinSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_4, v)
	}
	*j = IdTokenEnumType_4(v)
	return nil
}

const AttributeEnumType_8_Target AttributeEnumType_8 = "Target"
const AttributeEnumType_8_Actual AttributeEnumType_8 = "Actual"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_8, v)
	}
	*j = AttributeEnumType_8(v)
	return nil
}

type AttributeEnumType_8 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariablesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["setVariableData"]; !ok || v == nil {
		return fmt.Errorf("field setVariableData: required")
	}
	type Plain SetVariablesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetVariablesRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["attributeValue"]; !ok || v == nil {
		return fmt.Errorf("field attributeValue: required")
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	type Plain SetVariableDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetVariableDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_9) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_9
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_9(plain)
	return nil
}

const AttributeEnumType_7_MaxSet AttributeEnumType_7 = "MaxSet"
const AttributeEnumType_7_MinSet AttributeEnumType_7 = "MinSet"
const AttributeEnumType_7_Target AttributeEnumType_7 = "Target"

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_5, v)
	}
	*j = IdTokenEnumType_5(v)
	return nil
}

const AttributeEnumType_7_Actual AttributeEnumType_7 = "Actual"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_7, v)
	}
	*j = AttributeEnumType_7(v)
	return nil
}

type AttributeEnumType_7 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_11) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_11
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_11(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_11 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_112 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_12 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_12) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_12
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_12(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_112) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_112
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_112(plain)
	return nil
}

const AttributeEnumType_6_MaxSet AttributeEnumType_6 = "MaxSet"
const AttributeEnumType_6_MinSet AttributeEnumType_6 = "MinSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType_2(plain)
	return nil
}

const AttributeEnumType_6_Target AttributeEnumType_6 = "Target"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomerInformationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["clear"]; !ok || v == nil {
		return fmt.Errorf("field clear: required")
	}
	if v, ok := raw["report"]; !ok || v == nil {
		return fmt.Errorf("field report: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain CustomerInformationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomerInformationRequestJson(plain)
	return nil
}

const AttributeEnumType_6_Actual AttributeEnumType_6 = "Actual"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_23) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_23
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_23(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_6, v)
	}
	*j = AttributeEnumType_6(v)
	return nil
}

type AttributeEnumType_6 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomerInformationStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CustomerInformationStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CustomerInformationStatusEnumType, v)
	}
	*j = CustomerInformationStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableMonitoringResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["setMonitoringResult"]; !ok || v == nil {
		return fmt.Errorf("field setMonitoringResult: required")
	}
	type Plain SetVariableMonitoringResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetVariableMonitoringResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetMonitoringStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetMonitoringStatusEnumType_1, v)
	}
	*j = SetMonitoringStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringResultType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["severity"]; !ok || v == nil {
		return fmt.Errorf("field severity: required")
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	type Plain SetMonitoringResultType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetMonitoringResultType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_8) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_8
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_8(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_8) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_8
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_8(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType_5, v)
	}
	*j = MonitorEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_36) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_36
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_36(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomerInformationStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CustomerInformationStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CustomerInformationStatusEnumType_1, v)
	}
	*j = CustomerInformationStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetMonitoringStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetMonitoringStatusEnumType, v)
	}
	*j = SetMonitoringStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType_4, v)
	}
	*j = MonitorEnumType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_10) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_10
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_10(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_10 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_111 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_11 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomerInformationResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain CustomerInformationResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomerInformationResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_11) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_11
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_11(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_24) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_24
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_24(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_111) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_111
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_111(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataTransferRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain DataTransferRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DataTransferRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetVariableMonitoringRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["setMonitoringData"]; !ok || v == nil {
		return fmt.Errorf("field setMonitoringData: required")
	}
	type Plain SetVariableMonitoringRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetVariableMonitoringRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_25) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_25
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_25(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["severity"]; !ok || v == nil {
		return fmt.Errorf("field severity: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	type Plain SetMonitoringDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["transaction"]; !ok || v == nil {
		plain.Transaction = false
	}
	*j = SetMonitoringDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_7) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_7
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_7(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataTransferStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DataTransferStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DataTransferStatusEnumType, v)
	}
	*j = DataTransferStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType_3, v)
	}
	*j = MonitorEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType_2, v)
	}
	*j = MonitorEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_9) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_9
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_9(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_9 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_110 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_10 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_10) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_10
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_10(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_9) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_9
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_9(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_110) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_110
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_110(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetNetworkProfileResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain SetNetworkProfileResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetNetworkProfileResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataTransferStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DataTransferStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DataTransferStatusEnumType_1, v)
	}
	*j = DataTransferStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetNetworkProfileStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetNetworkProfileStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetNetworkProfileStatusEnumType_1, v)
	}
	*j = SetNetworkProfileStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_35) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_35
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_35(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetNetworkProfileStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SetNetworkProfileStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SetNetworkProfileStatusEnumType, v)
	}
	*j = SetNetworkProfileStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_109) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_109
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_109(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetNetworkProfileRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["configurationSlot"]; !ok || v == nil {
		return fmt.Errorf("field configurationSlot: required")
	}
	if v, ok := raw["connectionData"]; !ok || v == nil {
		return fmt.Errorf("field connectionData: required")
	}
	type Plain SetNetworkProfileRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetNetworkProfileRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataTransferResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain DataTransferResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DataTransferResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VPNEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VPNEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VPNEnumType_1, v)
	}
	*j = VPNEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_26) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_26
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_26(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCPPVersionEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OCPPVersionEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OCPPVersionEnumType_1, v)
	}
	*j = OCPPVersionEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCPPTransportEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OCPPTransportEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OCPPTransportEnumType_1, v)
	}
	*j = OCPPTransportEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_4, v)
	}
	*j = HashAlgorithmEnumType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCPPInterfaceEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OCPPInterfaceEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OCPPInterfaceEnumType_1, v)
	}
	*j = OCPPInterfaceEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NetworkConnectionProfileType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["messageTimeout"]; !ok || v == nil {
		return fmt.Errorf("field messageTimeout: required")
	}
	if v, ok := raw["ocppCsmsUrl"]; !ok || v == nil {
		return fmt.Errorf("field ocppCsmsUrl: required")
	}
	if v, ok := raw["ocppInterface"]; !ok || v == nil {
		return fmt.Errorf("field ocppInterface: required")
	}
	if v, ok := raw["ocppTransport"]; !ok || v == nil {
		return fmt.Errorf("field ocppTransport: required")
	}
	if v, ok := raw["ocppVersion"]; !ok || v == nil {
		return fmt.Errorf("field ocppVersion: required")
	}
	if v, ok := raw["securityProfile"]; !ok || v == nil {
		return fmt.Errorf("field securityProfile: required")
	}
	type Plain NetworkConnectionProfileType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NetworkConnectionProfileType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VPNType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["key"]; !ok || v == nil {
		return fmt.Errorf("field key: required")
	}
	if v, ok := raw["password"]; !ok || v == nil {
		return fmt.Errorf("field password: required")
	}
	if v, ok := raw["server"]; !ok || v == nil {
		return fmt.Errorf("field server: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	if v, ok := raw["user"]; !ok || v == nil {
		return fmt.Errorf("field user: required")
	}
	type Plain VPNType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VPNType(plain)
	return nil
}

type CertificateHashDataType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_26 `json:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType_4 `json:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateHashDataType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hashAlgorithm"]; !ok || v == nil {
		return fmt.Errorf("field hashAlgorithm: required")
	}
	if v, ok := raw["issuerKeyHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerKeyHash: required")
	}
	if v, ok := raw["issuerNameHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerNameHash: required")
	}
	if v, ok := raw["serialNumber"]; !ok || v == nil {
		return fmt.Errorf("field serialNumber: required")
	}
	type Plain CertificateHashDataType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CertificateHashDataType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VPNEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_VPNEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_VPNEnumType, v)
	}
	*j = VPNEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCPPVersionEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OCPPVersionEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OCPPVersionEnumType, v)
	}
	*j = OCPPVersionEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_5, v)
	}
	*j = HashAlgorithmEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCPPTransportEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OCPPTransportEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OCPPTransportEnumType, v)
	}
	*j = OCPPTransportEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCPPInterfaceEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_OCPPInterfaceEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_OCPPInterfaceEnumType, v)
	}
	*j = OCPPInterfaceEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *APNType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["apn"]; !ok || v == nil {
		return fmt.Errorf("field apn: required")
	}
	if v, ok := raw["apnAuthentication"]; !ok || v == nil {
		return fmt.Errorf("field apnAuthentication: required")
	}
	type Plain APNType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["useOnlyPreferredNetwork"]; !ok || v == nil {
		plain.UseOnlyPreferredNetwork = false
	}
	*j = APNType(plain)
	return nil
}

// APN
// urn:x-oca:ocpp:uid:2:233134
// Collection of configuration data needed to make a data-connection over a
// cellular network.
//
// NOTE: When asking a GSM modem to dial in, it is possible to specify which mobile
// operator should be used. This can be done with the mobile country code (MCC) in
// combination with a mobile network code (MNC). Example: If your preferred network
// is Vodafone Netherlands, the MCC=204 and the MNC=04 which means the key
// PreferredNetwork = 20404 Some modems allows to specify a preferred network,
// which means, if this network is not available, a different network is used. If
// you specify UseOnlyPreferredNetwork and this network is not available, the modem
// will not dial in.
//
type APNType struct {
	// APN. APN. URI
	// urn:x-oca:ocpp:uid:1:568814
	// The Access Point Name as an URL.
	//
	Apn string `json:"apn"`

	// ApnAuthentication corresponds to the JSON schema field "apnAuthentication".
	ApnAuthentication APNAuthenticationEnumType_1 `json:"apnAuthentication"`

	// APN. APN. Password
	// urn:x-oca:ocpp:uid:1:568819
	// APN Password.
	//
	ApnPassword *string `json:"apnPassword,omitempty"`

	// APN. APN. User_ Name
	// urn:x-oca:ocpp:uid:1:568818
	// APN username.
	//
	ApnUserName *string `json:"apnUserName,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_108 `json:"customData,omitempty"`

	// APN. Preferred_ Network. Mobile_ Network_ ID
	// urn:x-oca:ocpp:uid:1:568822
	// Preferred network, written as MCC and MNC concatenated. See note.
	//
	PreferredNetwork *string `json:"preferredNetwork,omitempty"`

	// APN. SIMPIN. PIN_ Code
	// urn:x-oca:ocpp:uid:1:568821
	// SIM card pin code.
	//
	SimPin *int `json:"simPin,omitempty"`

	// APN. Use_ Only_ Preferred_ Network. Indicator
	// urn:x-oca:ocpp:uid:1:568824
	// Default: false. Use only the preferred Network, do
	// not dial in when not available. See Note.
	//
	UseOnlyPreferredNetwork bool `json:"useOnlyPreferredNetwork,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeleteCertificateRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["certificateHashData"]; !ok || v == nil {
		return fmt.Errorf("field certificateHashData: required")
	}
	type Plain DeleteCertificateRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DeleteCertificateRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_108) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_108
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_108(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_27) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_27
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_27(plain)
	return nil
}

const APNAuthenticationEnumType_1_AUTO APNAuthenticationEnumType_1 = "AUTO"
const APNAuthenticationEnumType_1_PAP APNAuthenticationEnumType_1 = "PAP"

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeleteCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DeleteCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DeleteCertificateStatusEnumType, v)
	}
	*j = DeleteCertificateStatusEnumType(v)
	return nil
}

const APNAuthenticationEnumType_1_NONE APNAuthenticationEnumType_1 = "NONE"
const APNAuthenticationEnumType_1_CHAP APNAuthenticationEnumType_1 = "CHAP"

// UnmarshalJSON implements json.Unmarshaler.
func (j *APNAuthenticationEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_APNAuthenticationEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_APNAuthenticationEnumType_1, v)
	}
	*j = APNAuthenticationEnumType_1(v)
	return nil
}

type APNAuthenticationEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_10) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_10
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_10(plain)
	return nil
}

const APNAuthenticationEnumTypeAUTO APNAuthenticationEnumType = "AUTO"
const APNAuthenticationEnumTypePAP APNAuthenticationEnumType = "PAP"

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeleteCertificateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DeleteCertificateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DeleteCertificateStatusEnumType_1, v)
	}
	*j = DeleteCertificateStatusEnumType_1(v)
	return nil
}

const APNAuthenticationEnumTypeNONE APNAuthenticationEnumType = "NONE"
const APNAuthenticationEnumTypeCHAP APNAuthenticationEnumType = "CHAP"

// UnmarshalJSON implements json.Unmarshaler.
func (j *APNAuthenticationEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_APNAuthenticationEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_APNAuthenticationEnumType, v)
	}
	*j = APNAuthenticationEnumType(v)
	return nil
}

type APNAuthenticationEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *DeleteCertificateResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain DeleteCertificateResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DeleteCertificateResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringLevelResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain SetMonitoringLevelResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetMonitoringLevelResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_28) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_28
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_28(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_7, v)
	}
	*j = GenericStatusEnumType_7(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_34) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_34
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_34(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FirmwareStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_FirmwareStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_FirmwareStatusEnumType, v)
	}
	*j = FirmwareStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_6, v)
	}
	*j = GenericStatusEnumType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_107) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_107
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_107(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringLevelRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["severity"]; !ok || v == nil {
		return fmt.Errorf("field severity: required")
	}
	type Plain SetMonitoringLevelRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetMonitoringLevelRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_106) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_106
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_106(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringBaseResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain SetMonitoringBaseResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetMonitoringBaseResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType_7, v)
	}
	*j = GenericDeviceModelStatusEnumType_7(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_33) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_33
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_33(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType_6, v)
	}
	*j = GenericDeviceModelStatusEnumType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_105) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_105
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_105(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetMonitoringBaseRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["monitoringBase"]; !ok || v == nil {
		return fmt.Errorf("field monitoringBase: required")
	}
	type Plain SetMonitoringBaseRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetMonitoringBaseRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitoringBaseEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitoringBaseEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitoringBaseEnumType_1, v)
	}
	*j = MonitoringBaseEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitoringBaseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitoringBaseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitoringBaseEnumType, v)
	}
	*j = MonitoringBaseEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_104) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_104
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_104(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetDisplayMessageResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain SetDisplayMessageResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetDisplayMessageResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DisplayMessageStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DisplayMessageStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DisplayMessageStatusEnumType_1, v)
	}
	*j = DisplayMessageStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_32) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_32
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_32(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FirmwareStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_FirmwareStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_FirmwareStatusEnumType_1, v)
	}
	*j = FirmwareStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DisplayMessageStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DisplayMessageStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DisplayMessageStatusEnumType, v)
	}
	*j = DisplayMessageStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_103) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_103
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_103(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetDisplayMessageRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["message"]; !ok || v == nil {
		return fmt.Errorf("field message: required")
	}
	type Plain SetDisplayMessageRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetDisplayMessageRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType_5, v)
	}
	*j = MessageStateEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType_5, v)
	}
	*j = MessagePriorityEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageInfoType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["message"]; !ok || v == nil {
		return fmt.Errorf("field message: required")
	}
	if v, ok := raw["priority"]; !ok || v == nil {
		return fmt.Errorf("field priority: required")
	}
	type Plain MessageInfoType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageInfoType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType_4, v)
	}
	*j = MessageStateEnumType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType_4, v)
	}
	*j = MessagePriorityEnumType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_7, v)
	}
	*j = MessageFormatEnumType_7(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageContentType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["content"]; !ok || v == nil {
		return fmt.Errorf("field content: required")
	}
	if v, ok := raw["format"]; !ok || v == nil {
		return fmt.Errorf("field format: required")
	}
	type Plain MessageContentType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageContentType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_6, v)
	}
	*j = MessageFormatEnumType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_8) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_8
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_8(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_8 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_102 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_9 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_9) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_9
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_9(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_102) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_102
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_102(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FirmwareStatusNotificationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain FirmwareStatusNotificationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FirmwareStatusNotificationRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetChargingProfileResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain SetChargingProfileResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetChargingProfileResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_29) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_29
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_29(plain)
	return nil
}

const ChargingProfileStatusEnumType_1_Rejected ChargingProfileStatusEnumType_1 = "Rejected"

type CertificateActionEnumType string

const ChargingProfileStatusEnumType_1_Accepted ChargingProfileStatusEnumType_1 = "Accepted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateActionEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateActionEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateActionEnumType, v)
	}
	*j = CertificateActionEnumType(v)
	return nil
}

const CertificateActionEnumTypeInstall CertificateActionEnumType = "Install"
const CertificateActionEnumTypeUpdate CertificateActionEnumType = "Update"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileStatusEnumType_1, v)
	}
	*j = ChargingProfileStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_30) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_30
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_30(plain)
	return nil
}

type CertificateActionEnumType_1 string

type ChargingProfileStatusEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateActionEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CertificateActionEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CertificateActionEnumType_1, v)
	}
	*j = CertificateActionEnumType_1(v)
	return nil
}

const CertificateActionEnumType_1_Install CertificateActionEnumType_1 = "Install"
const CertificateActionEnumType_1_Update CertificateActionEnumType_1 = "Update"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_31) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_31
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_31(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Get15118EVCertificateRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["action"]; !ok || v == nil {
		return fmt.Errorf("field action: required")
	}
	if v, ok := raw["exiRequest"]; !ok || v == nil {
		return fmt.Errorf("field exiRequest: required")
	}
	if v, ok := raw["iso15118SchemaVersion"]; !ok || v == nil {
		return fmt.Errorf("field iso15118SchemaVersion: required")
	}
	type Plain Get15118EVCertificateRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Get15118EVCertificateRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_101) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_101
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_101(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_31) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_31
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_31(plain)
	return nil
}

const ChargingProfileStatusEnumTypeRejected ChargingProfileStatusEnumType = "Rejected"
const ChargingProfileStatusEnumTypeAccepted ChargingProfileStatusEnumType = "Accepted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *Iso15118EVCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Iso15118EVCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Iso15118EVCertificateStatusEnumType, v)
	}
	*j = Iso15118EVCertificateStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileStatusEnumType, v)
	}
	*j = ChargingProfileStatusEnumType(v)
	return nil
}

type ChargingProfileStatusEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *SetChargingProfileRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingProfile"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfile: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	type Plain SetChargingProfileRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SetChargingProfileRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_11) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_11
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_11(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType_5, v)
	}
	*j = RecurrencyKindEnumType_5(v)
	return nil
}

const CostKindEnumType_9_RenewableGenerationPercentage CostKindEnumType_9 = "RenewableGenerationPercentage"

// UnmarshalJSON implements json.Unmarshaler.
func (j *Iso15118EVCertificateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_Iso15118EVCertificateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_Iso15118EVCertificateStatusEnumType_1, v)
	}
	*j = Iso15118EVCertificateStatusEnumType_1(v)
	return nil
}

const CostKindEnumType_9_RelativePricePercentage CostKindEnumType_9 = "RelativePricePercentage"
const CostKindEnumType_9_CarbonDioxideEmission CostKindEnumType_9 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_9, v)
	}
	*j = CostKindEnumType_9(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Get15118EVCertificateResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["exiResponse"]; !ok || v == nil {
		return fmt.Errorf("field exiResponse: required")
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain Get15118EVCertificateResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Get15118EVCertificateResponseJson(plain)
	return nil
}

type CostKindEnumType_9 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_32) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_32
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_32(plain)
	return nil
}

const ChargingRateUnitEnumType_13_A ChargingRateUnitEnumType_13 = "A"
const ChargingRateUnitEnumType_13_W ChargingRateUnitEnumType_13 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportBaseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReportBaseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReportBaseEnumType, v)
	}
	*j = ReportBaseEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_13) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_13 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_13, v)
	}
	*j = ChargingRateUnitEnumType_13(v)
	return nil
}

type ChargingRateUnitEnumType_13 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingProfileKind"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfileKind: required")
	}
	if v, ok := raw["chargingProfilePurpose"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfilePurpose: required")
	}
	if v, ok := raw["chargingSchedule"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedule: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["stackLevel"]; !ok || v == nil {
		return fmt.Errorf("field stackLevel: required")
	}
	type Plain ChargingProfileType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingProfileType_2(plain)
	return nil
}

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
//
type ChargingProfileType_2 struct {
	// ChargingProfileKind corresponds to the JSON schema field "chargingProfileKind".
	ChargingProfileKind ChargingProfileKindEnumType_5 `json:"chargingProfileKind"`

	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose ChargingProfilePurposeEnumType_9 `json:"chargingProfilePurpose"`

	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule []ChargingScheduleType_4 `json:"chargingSchedule"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Id of ChargingProfile.
	//
	Id int `json:"id"`

	// RecurrencyKind corresponds to the JSON schema field "recurrencyKind".
	RecurrencyKind *RecurrencyKindEnumType_4 `json:"recurrencyKind,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Value determining level in hierarchy stack of profiles. Higher values have
	// precedence over lower values. Lowest level is 0.
	//
	StackLevel int `json:"stackLevel"`

	// SHALL only be included if ChargingProfilePurpose is set to TxProfile. The
	// transactionId is used to match the profile to a specific transaction.
	//
	TransactionId *string `json:"transactionId,omitempty"`

	// Charging_ Profile. Valid_ From. Date_ Time
	// urn:x-oca:ocpp:uid:1:569234
	// Point in time at which the profile starts to be valid. If absent, the profile
	// is valid as soon as it is received by the Charging Station.
	//
	ValidFrom *string `json:"validFrom,omitempty"`

	// Charging_ Profile. Valid_ To. Date_ Time
	// urn:x-oca:ocpp:uid:1:569235
	// Point in time at which the profile stops to be valid. If absent, the profile is
	// valid until it is replaced by another profile.
	//
	ValidTo *string `json:"validTo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType_4, v)
	}
	*j = RecurrencyKindEnumType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportBaseEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReportBaseEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReportBaseEnumType_1, v)
	}
	*j = ReportBaseEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingScheduleType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingRateUnit"]; !ok || v == nil {
		return fmt.Errorf("field chargingRateUnit: required")
	}
	if v, ok := raw["chargingSchedulePeriod"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedulePeriod: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain ChargingScheduleType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingScheduleType_4(plain)
	return nil
}

// Charging_ Schedule
// urn:x-oca:ocpp:uid:2:233256
// Charging schedule structure defines a list of charging periods, as used in:
// GetCompositeSchedule.conf and ChargingProfile.
//
type ChargingScheduleType_4 struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType_12 `json:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType_5 `json:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// Charging_ Schedule. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569236
	// Duration of the charging schedule in seconds. If the duration is left empty,
	// the last period will continue indefinitely or until end of the transaction if
	// chargingProfilePurpose = TxProfile.
	//
	Duration *int `json:"duration,omitempty"`

	// Identifies the ChargingSchedule.
	//
	Id int `json:"id"`

	// Charging_ Schedule. Min_ Charging_ Rate. Numeric
	// urn:x-oca:ocpp:uid:1:569239
	// Minimum charging rate supported by the EV. The unit of measure is defined by
	// the chargingRateUnit. This parameter is intended to be used by a local smart
	// charging algorithm to optimize the power allocation for in the case a charging
	// process is inefficient at lower charging rates. Accepts at most one digit
	// fraction (e.g. 8.1)
	//
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`

	// SalesTariff corresponds to the JSON schema field "salesTariff".
	SalesTariff *SalesTariffType_4 `json:"salesTariff,omitempty"`

	// Charging_ Schedule. Start_ Schedule. Date_ Time
	// urn:x-oca:ocpp:uid:1:569237
	// Starting point of an absolute schedule. If absent the schedule will be relative
	// to start of charging.
	//
	StartSchedule *string `json:"startSchedule,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["salesTariffEntry"]; !ok || v == nil {
		return fmt.Errorf("field salesTariffEntry: required")
	}
	type Plain SalesTariffType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffType_4(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffEntryType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["relativeTimeInterval"]; !ok || v == nil {
		return fmt.Errorf("field relativeTimeInterval: required")
	}
	type Plain SalesTariffEntryType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffEntryType_4(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetBaseReportRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reportBase"]; !ok || v == nil {
		return fmt.Errorf("field reportBase: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain GetBaseReportRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetBaseReportRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RelativeTimeIntervalType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["start"]; !ok || v == nil {
		return fmt.Errorf("field start: required")
	}
	type Plain RelativeTimeIntervalType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RelativeTimeIntervalType_4(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_33) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_33
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_33(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsumptionCostType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cost"]; !ok || v == nil {
		return fmt.Errorf("field cost: required")
	}
	if v, ok := raw["startValue"]; !ok || v == nil {
		return fmt.Errorf("field startValue: required")
	}
	type Plain ConsumptionCostType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ConsumptionCostType_4(plain)
	return nil
}

// Consumption_ Cost
// urn:x-oca:ocpp:uid:2:233259
//
type ConsumptionCostType_4 struct {
	// Cost corresponds to the JSON schema field "cost".
	Cost []CostType_4 `json:"cost"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// Consumption_ Cost. Start_ Value. Numeric
	// urn:x-oca:ocpp:uid:1:569246
	// The lowest level of consumption that defines the starting point of this
	// consumption block. The block interval extends to the start of the next
	// interval.
	//
	StartValue float64 `json:"startValue"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType, v)
	}
	*j = GenericDeviceModelStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["amount"]; !ok || v == nil {
		return fmt.Errorf("field amount: required")
	}
	if v, ok := raw["costKind"]; !ok || v == nil {
		return fmt.Errorf("field costKind: required")
	}
	type Plain CostType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostType_4(plain)
	return nil
}

// Cost
// urn:x-oca:ocpp:uid:2:233258
//
type CostType_4 struct {
	// Cost. Amount. Amount
	// urn:x-oca:ocpp:uid:1:569244
	// The estimated or actual cost per kWh
	//
	Amount int `json:"amount"`

	// Cost. Amount_ Multiplier. Integer
	// urn:x-oca:ocpp:uid:1:569245
	// Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The
	// final value is determined by: amount * 10 ^ amountMultiplier
	//
	AmountMultiplier *int `json:"amountMultiplier,omitempty"`

	// CostKind corresponds to the JSON schema field "costKind".
	CostKind CostKindEnumType_8 `json:"costKind"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`
}

const CostKindEnumType_8_RenewableGenerationPercentage CostKindEnumType_8 = "RenewableGenerationPercentage"
const CostKindEnumType_8_RelativePricePercentage CostKindEnumType_8 = "RelativePricePercentage"
const CostKindEnumType_8_CarbonDioxideEmission CostKindEnumType_8 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_12) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_12
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_12(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_8, v)
	}
	*j = CostKindEnumType_8(v)
	return nil
}

type CostKindEnumType_8 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType_1, v)
	}
	*j = GenericDeviceModelStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit: required")
	}
	if v, ok := raw["startPeriod"]; !ok || v == nil {
		return fmt.Errorf("field startPeriod: required")
	}
	type Plain ChargingSchedulePeriodType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType_5(plain)
	return nil
}

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
//
type ChargingSchedulePeriodType_5 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_100) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_100
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_100(plain)
	return nil
}

const ChargingRateUnitEnumType_12_A ChargingRateUnitEnumType_12 = "A"
const ChargingRateUnitEnumType_12_W ChargingRateUnitEnumType_12 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetBaseReportResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetBaseReportResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetBaseReportResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_12) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_12 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_12, v)
	}
	*j = ChargingRateUnitEnumType_12(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_34) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_34
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_34(plain)
	return nil
}

type ChargingRateUnitEnumType_12 string

const ChargingProfilePurposeEnumType_9_TxProfile ChargingProfilePurposeEnumType_9 = "TxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_6, v)
	}
	*j = HashAlgorithmEnumType_6(v)
	return nil
}

const ChargingProfilePurposeEnumType_9_TxDefaultProfile ChargingProfilePurposeEnumType_9 = "TxDefaultProfile"
const ChargingProfilePurposeEnumType_9_ChargingStationMaxProfile ChargingProfilePurposeEnumType_9 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_9_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_9 = "ChargingStationExternalConstraints"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_9, v)
	}
	*j = ChargingProfilePurposeEnumType_9(v)
	return nil
}

type ChargingProfilePurposeEnumType_9 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_7, v)
	}
	*j = HashAlgorithmEnumType_7(v)
	return nil
}

const ChargingProfileKindEnumType_5_Relative ChargingProfileKindEnumType_5 = "Relative"
const ChargingProfileKindEnumType_5_Recurring ChargingProfileKindEnumType_5 = "Recurring"
const ChargingProfileKindEnumType_5_Absolute ChargingProfileKindEnumType_5 = "Absolute"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType_5, v)
	}
	*j = ChargingProfileKindEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *OCSPRequestDataType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hashAlgorithm"]; !ok || v == nil {
		return fmt.Errorf("field hashAlgorithm: required")
	}
	if v, ok := raw["issuerKeyHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerKeyHash: required")
	}
	if v, ok := raw["issuerNameHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerNameHash: required")
	}
	if v, ok := raw["responderURL"]; !ok || v == nil {
		return fmt.Errorf("field responderURL: required")
	}
	if v, ok := raw["serialNumber"]; !ok || v == nil {
		return fmt.Errorf("field serialNumber: required")
	}
	type Plain OCSPRequestDataType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = OCSPRequestDataType_1(plain)
	return nil
}

type ChargingProfileKindEnumType_5 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateStatusRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["ocspRequestData"]; !ok || v == nil {
		return fmt.Errorf("field ocspRequestData: required")
	}
	type Plain GetCertificateStatusRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetCertificateStatusRequestJson(plain)
	return nil
}

const ChargingProfilePurposeEnumType_8_TxProfile ChargingProfilePurposeEnumType_8 = "TxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_35) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_35
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_35(plain)
	return nil
}

const ChargingProfilePurposeEnumType_8_TxDefaultProfile ChargingProfilePurposeEnumType_8 = "TxDefaultProfile"
const ChargingProfilePurposeEnumType_8_ChargingStationMaxProfile ChargingProfilePurposeEnumType_8 = "ChargingStationMaxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetCertificateStatusEnumType, v)
	}
	*j = GetCertificateStatusEnumType(v)
	return nil
}

const ChargingProfilePurposeEnumType_8_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_8 = "ChargingStationExternalConstraints"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_8, v)
	}
	*j = ChargingProfilePurposeEnumType_8(v)
	return nil
}

type ChargingProfilePurposeEnumType_8 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_13) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_13
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_13(plain)
	return nil
}

const ChargingProfileKindEnumType_4_Relative ChargingProfileKindEnumType_4 = "Relative"
const ChargingProfileKindEnumType_4_Recurring ChargingProfileKindEnumType_4 = "Recurring"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetCertificateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetCertificateStatusEnumType_1, v)
	}
	*j = GetCertificateStatusEnumType_1(v)
	return nil
}

const ChargingProfileKindEnumType_4_Absolute ChargingProfileKindEnumType_4 = "Absolute"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType_4, v)
	}
	*j = ChargingProfileKindEnumType_4(v)
	return nil
}

type ChargingProfileKindEnumType_4 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateStatusResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetCertificateStatusResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetCertificateStatusResponseJson(plain)
	return nil
}

type ChargingLimitSourceEnumType_2 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *SendLocalListResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain SendLocalListResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SendLocalListResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType_2, v)
	}
	*j = ChargingLimitSourceEnumType_2(v)
	return nil
}

const ChargingLimitSourceEnumType_2_EMS ChargingLimitSourceEnumType_2 = "EMS"
const ChargingLimitSourceEnumType_2_Other ChargingLimitSourceEnumType_2 = "Other"
const ChargingLimitSourceEnumType_2_SO ChargingLimitSourceEnumType_2 = "SO"
const ChargingLimitSourceEnumType_2_CSO ChargingLimitSourceEnumType_2 = "CSO"

type ChargingLimitSourceEnumType_3 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *SendLocalListStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SendLocalListStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SendLocalListStatusEnumType_1, v)
	}
	*j = SendLocalListStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType_3, v)
	}
	*j = ChargingLimitSourceEnumType_3(v)
	return nil
}

const ChargingLimitSourceEnumType_3_EMS ChargingLimitSourceEnumType_3 = "EMS"
const ChargingLimitSourceEnumType_3_Other ChargingLimitSourceEnumType_3 = "Other"
const ChargingLimitSourceEnumType_3_SO ChargingLimitSourceEnumType_3 = "SO"
const ChargingLimitSourceEnumType_3_CSO ChargingLimitSourceEnumType_3 = "CSO"

type ChargingProfilePurposeEnumType_2 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_30) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_30
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_30(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_2, v)
	}
	*j = ChargingProfilePurposeEnumType_2(v)
	return nil
}

const ChargingProfilePurposeEnumType_2_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_2 = "ChargingStationExternalConstraints"
const ChargingProfilePurposeEnumType_2_ChargingStationMaxProfile ChargingProfilePurposeEnumType_2 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_2_TxDefaultProfile ChargingProfilePurposeEnumType_2 = "TxDefaultProfile"
const ChargingProfilePurposeEnumType_2_TxProfile ChargingProfilePurposeEnumType_2 = "TxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SendLocalListStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_SendLocalListStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_SendLocalListStatusEnumType, v)
	}
	*j = SendLocalListStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_36) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_36
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_36(plain)
	return nil
}

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
//
type ChargingProfileCriterionType struct {
	// For which charging limit sources, charging profiles SHALL be reported. If
	// omitted, the Charging Station SHALL not filter on chargingLimitSource.
	//
	ChargingLimitSource []ChargingLimitSourceEnumType_3 `json:"chargingLimitSource,omitempty"`

	// List of all the chargingProfileIds requested. Any ChargingProfile that matches
	// one of these profiles will be reported. If omitted, the Charging Station SHALL
	// not filter on chargingProfileId. This field SHALL NOT contain more ids than set
	// in
	// &lt;&lt;configkey-charging-profile-entries,ChargingProfileEntries.maxLimit&gt;&gt;
	//
	//
	ChargingProfileId []int `json:"chargingProfileId,omitempty"`

	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose *ChargingProfilePurposeEnumType_2 `json:"chargingProfilePurpose,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_36 `json:"customData,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Value determining level in hierarchy stack of profiles. Higher values have
	// precedence over lower values. Lowest level is 0.
	//
	StackLevel *int `json:"stackLevel,omitempty"`
}

type ChargingProfilePurposeEnumType_3 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_99) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_99
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_99(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_3, v)
	}
	*j = ChargingProfilePurposeEnumType_3(v)
	return nil
}

const ChargingProfilePurposeEnumType_3_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_3 = "ChargingStationExternalConstraints"
const ChargingProfilePurposeEnumType_3_ChargingStationMaxProfile ChargingProfilePurposeEnumType_3 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_3_TxDefaultProfile ChargingProfilePurposeEnumType_3 = "TxDefaultProfile"
const ChargingProfilePurposeEnumType_3_TxProfile ChargingProfilePurposeEnumType_3 = "TxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *SendLocalListRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["updateType"]; !ok || v == nil {
		return fmt.Errorf("field updateType: required")
	}
	if v, ok := raw["versionNumber"]; !ok || v == nil {
		return fmt.Errorf("field versionNumber: required")
	}
	type Plain SendLocalListRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SendLocalListRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetChargingProfilesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingProfile"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfile: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain GetChargingProfilesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetChargingProfilesRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UpdateEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UpdateEnumType_1, v)
	}
	*j = UpdateEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_37) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_37
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_37(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UpdateEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UpdateEnumType, v)
	}
	*j = UpdateEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_5, v)
	}
	*j = MessageFormatEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetChargingProfileStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetChargingProfileStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetChargingProfileStatusEnumType, v)
	}
	*j = GetChargingProfileStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_11) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_11 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_11, v)
	}
	*j = IdTokenEnumType_11(v)
	return nil
}

const AuthorizationStatusEnumType_3_Unknown AuthorizationStatusEnumType_3 = "Unknown"
const AuthorizationStatusEnumType_3_NotAtThisTime AuthorizationStatusEnumType_3 = "NotAtThisTime"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_14) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_14
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_14(plain)
	return nil
}

const AuthorizationStatusEnumType_3_NotAtThisLocation AuthorizationStatusEnumType_3 = "NotAtThisLocation"
const AuthorizationStatusEnumType_3_NotAllowedTypeEVSE AuthorizationStatusEnumType_3 = "NotAllowedTypeEVSE"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetChargingProfileStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetChargingProfileStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetChargingProfileStatusEnumType_1, v)
	}
	*j = GetChargingProfileStatusEnumType_1(v)
	return nil
}

const AuthorizationStatusEnumType_3_NoCredit AuthorizationStatusEnumType_3 = "NoCredit"
const AuthorizationStatusEnumType_3_Invalid AuthorizationStatusEnumType_3 = "Invalid"
const AuthorizationStatusEnumType_3_Expired AuthorizationStatusEnumType_3 = "Expired"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetChargingProfilesResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetChargingProfilesResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetChargingProfilesResponseJson(plain)
	return nil
}

type ChargingRateUnitEnumType string

const AuthorizationStatusEnumType_3_ConcurrentTx AuthorizationStatusEnumType_3 = "ConcurrentTx"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType, v)
	}
	*j = ChargingRateUnitEnumType(v)
	return nil
}

const ChargingRateUnitEnumTypeW ChargingRateUnitEnumType = "W"
const ChargingRateUnitEnumTypeA ChargingRateUnitEnumType = "A"
const AuthorizationStatusEnumType_3_Blocked AuthorizationStatusEnumType_3 = "Blocked"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_38) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_38
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_38(plain)
	return nil
}

type ChargingRateUnitEnumType_1 string

const AuthorizationStatusEnumType_3_Accepted AuthorizationStatusEnumType_3 = "Accepted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_1, v)
	}
	*j = ChargingRateUnitEnumType_1(v)
	return nil
}

const ChargingRateUnitEnumType_1_W ChargingRateUnitEnumType_1 = "W"
const ChargingRateUnitEnumType_1_A ChargingRateUnitEnumType_1 = "A"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType_3, v)
	}
	*j = AuthorizationStatusEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCompositeScheduleRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["duration"]; !ok || v == nil {
		return fmt.Errorf("field duration: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	type Plain GetCompositeScheduleRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetCompositeScheduleRequestJson(plain)
	return nil
}

type ChargingRateUnitEnumType_2 string

type AuthorizationStatusEnumType_3 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_2, v)
	}
	*j = ChargingRateUnitEnumType_2(v)
	return nil
}

const ChargingRateUnitEnumType_2_W ChargingRateUnitEnumType_2 = "W"
const ChargingRateUnitEnumType_2_A ChargingRateUnitEnumType_2 = "A"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationData) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	type Plain AuthorizationData
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuthorizationData(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_39) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_39
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_39(plain)
	return nil
}

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
//
type ChargingSchedulePeriodType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_39 `json:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit: required")
	}
	if v, ok := raw["startPeriod"]; !ok || v == nil {
		return fmt.Errorf("field startPeriod: required")
	}
	type Plain ChargingSchedulePeriodType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType(plain)
	return nil
}

type ChargingRateUnitEnumType_3 string

// Contains the identifier to use for authorization.
//
type AuthorizationData struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_98 `json:"customData,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken IdTokenType_5 `json:"idToken"`

	// IdTokenInfo corresponds to the JSON schema field "idTokenInfo".
	IdTokenInfo *IdTokenInfoType_1 `json:"idTokenInfo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_3, v)
	}
	*j = ChargingRateUnitEnumType_3(v)
	return nil
}

const ChargingRateUnitEnumType_3_W ChargingRateUnitEnumType_3 = "W"
const ChargingRateUnitEnumType_3_A ChargingRateUnitEnumType_3 = "A"

// Composite_ Schedule
// urn:x-oca:ocpp:uid:2:233362
//
type CompositeScheduleType struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType_3 `json:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType `json:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_39 `json:"customData,omitempty"`

	// Duration of the schedule in seconds.
	//
	Duration int `json:"duration"`

	// The ID of the EVSE for which the
	// schedule is requested. When evseid=0, the
	// Charging Station calculated the expected
	// consumption for the grid connection.
	//
	EvseId int `json:"evseId"`

	// Composite_ Schedule. Start. Date_ Time
	// urn:x-oca:ocpp:uid:1:569456
	// Date and time at which the schedule becomes active. All time measurements
	// within the schedule are relative to this timestamp.
	//
	ScheduleStart string `json:"scheduleStart"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CompositeScheduleType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingRateUnit"]; !ok || v == nil {
		return fmt.Errorf("field chargingRateUnit: required")
	}
	if v, ok := raw["chargingSchedulePeriod"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedulePeriod: required")
	}
	if v, ok := raw["duration"]; !ok || v == nil {
		return fmt.Errorf("field duration: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	if v, ok := raw["scheduleStart"]; !ok || v == nil {
		return fmt.Errorf("field scheduleStart: required")
	}
	type Plain CompositeScheduleType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CompositeScheduleType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenInfoType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain IdTokenInfoType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenInfoType_1(plain)
	return nil
}

const AuthorizationStatusEnumType_2_Unknown AuthorizationStatusEnumType_2 = "Unknown"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType, v)
	}
	*j = GenericStatusEnumType(v)
	return nil
}

const AuthorizationStatusEnumType_2_NotAtThisTime AuthorizationStatusEnumType_2 = "NotAtThisTime"
const AuthorizationStatusEnumType_2_NotAtThisLocation AuthorizationStatusEnumType_2 = "NotAtThisLocation"
const AuthorizationStatusEnumType_2_NotAllowedTypeEVSE AuthorizationStatusEnumType_2 = "NotAllowedTypeEVSE"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_15) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_15
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_15(plain)
	return nil
}

const AuthorizationStatusEnumType_2_NoCredit AuthorizationStatusEnumType_2 = "NoCredit"
const AuthorizationStatusEnumType_2_Invalid AuthorizationStatusEnumType_2 = "Invalid"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_1, v)
	}
	*j = GenericStatusEnumType_1(v)
	return nil
}

const AuthorizationStatusEnumType_2_Expired AuthorizationStatusEnumType_2 = "Expired"
const AuthorizationStatusEnumType_2_ConcurrentTx AuthorizationStatusEnumType_2 = "ConcurrentTx"
const AuthorizationStatusEnumType_2_Blocked AuthorizationStatusEnumType_2 = "Blocked"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCompositeScheduleResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetCompositeScheduleResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetCompositeScheduleResponseJson(plain)
	return nil
}

const AuthorizationStatusEnumType_2_Accepted AuthorizationStatusEnumType_2 = "Accepted"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_40) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_40
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_40(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuthorizationStatusEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuthorizationStatusEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuthorizationStatusEnumType_2, v)
	}
	*j = AuthorizationStatusEnumType_2(v)
	return nil
}

type AuthorizationStatusEnumType_2 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType, v)
	}
	*j = MessagePriorityEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageContentType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["content"]; !ok || v == nil {
		return fmt.Errorf("field content: required")
	}
	if v, ok := raw["format"]; !ok || v == nil {
		return fmt.Errorf("field format: required")
	}
	type Plain MessageContentType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageContentType_2(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_4, v)
	}
	*j = MessageFormatEnumType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType_5(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_10) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_10 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_10, v)
	}
	*j = IdTokenEnumType_10(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType_5(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType, v)
	}
	*j = MessageStateEnumType(v)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType_5 struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_98 `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_98) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_98
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_98(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_97) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_97
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_97(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SecurityEventNotificationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["timestamp"]; !ok || v == nil {
		return fmt.Errorf("field timestamp: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain SecurityEventNotificationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SecurityEventNotificationRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_96) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_96
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_96(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ResetResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain ResetResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ResetResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType_1, v)
	}
	*j = MessagePriorityEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ResetStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ResetStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ResetStatusEnumType_1, v)
	}
	*j = ResetStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_29) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_29
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_29(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ResetStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ResetStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ResetStatusEnumType, v)
	}
	*j = ResetStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_95) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_95
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_95(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ResetRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain ResetRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ResetRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType_1, v)
	}
	*j = MessageStateEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ResetEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ResetEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ResetEnumType_1, v)
	}
	*j = ResetEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ResetEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ResetEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ResetEnumType, v)
	}
	*j = ResetEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_94) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_94
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_94(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReserveNowResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain ReserveNowResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ReserveNowResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReserveNowStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReserveNowStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReserveNowStatusEnumType_1, v)
	}
	*j = ReserveNowStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetDisplayMessagesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain GetDisplayMessagesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetDisplayMessagesRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_28) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_28
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_28(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_41) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_41
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_41(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReserveNowStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReserveNowStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReserveNowStatusEnumType, v)
	}
	*j = ReserveNowStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_93) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_93
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_93(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetDisplayMessagesStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetDisplayMessagesStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetDisplayMessagesStatusEnumType, v)
	}
	*j = GetDisplayMessagesStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReserveNowRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["expiryDateTime"]; !ok || v == nil {
		return fmt.Errorf("field expiryDateTime: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	type Plain ReserveNowRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ReserveNowRequestJson(plain)
	return nil
}

const ConnectorEnumType_1_Unknown ConnectorEnumType_1 = "Unknown"
const ConnectorEnumType_1_Undetermined ConnectorEnumType_1 = "Undetermined"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_16) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_16
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_16(plain)
	return nil
}

const ConnectorEnumType_1_WResonant ConnectorEnumType_1 = "wResonant"
const ConnectorEnumType_1_WInductive ConnectorEnumType_1 = "wInductive"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetDisplayMessagesStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetDisplayMessagesStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetDisplayMessagesStatusEnumType_1, v)
	}
	*j = GetDisplayMessagesStatusEnumType_1(v)
	return nil
}

const ConnectorEnumType_1_Pan ConnectorEnumType_1 = "Pan"
const ConnectorEnumType_1_Other3Ph ConnectorEnumType_1 = "Other3Ph"
const ConnectorEnumType_1_Other1PhOver16A ConnectorEnumType_1 = "Other1PhOver16A"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetDisplayMessagesResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetDisplayMessagesResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetDisplayMessagesResponseJson(plain)
	return nil
}

const ConnectorEnumType_1_Other1PhMax16A ConnectorEnumType_1 = "Other1PhMax16A"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_42) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_42
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_42(plain)
	return nil
}

const ConnectorEnumType_1_SType3 ConnectorEnumType_1 = "sType3"
const ConnectorEnumType_1_SType2 ConnectorEnumType_1 = "sType2"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateIdUseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetCertificateIdUseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetCertificateIdUseEnumType, v)
	}
	*j = GetCertificateIdUseEnumType(v)
	return nil
}

const ConnectorEnumType_1_SCEE77 ConnectorEnumType_1 = "sCEE-7-7"
const ConnectorEnumType_1_SBS1361 ConnectorEnumType_1 = "sBS1361"
const ConnectorEnumType_1_S3093P32A ConnectorEnumType_1 = "s309-3P-32A"
const ConnectorEnumType_1_S3093P16A ConnectorEnumType_1 = "s309-3P-16A"
const ConnectorEnumType_1_S3091P32A ConnectorEnumType_1 = "s309-1P-32A"
const ConnectorEnumType_1_S3091P16A ConnectorEnumType_1 = "s309-1P-16A"
const ConnectorEnumType_1_CType2 ConnectorEnumType_1 = "cType2"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateIdUseEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetCertificateIdUseEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetCertificateIdUseEnumType_1, v)
	}
	*j = GetCertificateIdUseEnumType_1(v)
	return nil
}

const ConnectorEnumType_1_CType1 ConnectorEnumType_1 = "cType1"
const ConnectorEnumType_1_CTesla ConnectorEnumType_1 = "cTesla"
const ConnectorEnumType_1_CG105 ConnectorEnumType_1 = "cG105"
const ConnectorEnumType_1_CCCS2 ConnectorEnumType_1 = "cCCS2"
const ConnectorEnumType_1_CCCS1 ConnectorEnumType_1 = "cCCS1"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConnectorEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConnectorEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConnectorEnumType_1, v)
	}
	*j = ConnectorEnumType_1(v)
	return nil
}

type ConnectorEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_43) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_43
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_43(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType_4(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_9, v)
	}
	*j = IdTokenEnumType_9(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_8, v)
	}
	*j = HashAlgorithmEnumType_8(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_8, v)
	}
	*j = IdTokenEnumType_8(v)
	return nil
}

const ConnectorEnumTypeUnknown ConnectorEnumType = "Unknown"
const ConnectorEnumTypeUndetermined ConnectorEnumType = "Undetermined"

type CertificateHashDataType_2 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_43 `json:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType_8 `json:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateHashDataType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["hashAlgorithm"]; !ok || v == nil {
		return fmt.Errorf("field hashAlgorithm: required")
	}
	if v, ok := raw["issuerKeyHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerKeyHash: required")
	}
	if v, ok := raw["issuerNameHash"]; !ok || v == nil {
		return fmt.Errorf("field issuerNameHash: required")
	}
	if v, ok := raw["serialNumber"]; !ok || v == nil {
		return fmt.Errorf("field serialNumber: required")
	}
	type Plain CertificateHashDataType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CertificateHashDataType_2(plain)
	return nil
}

const ConnectorEnumTypeWResonant ConnectorEnumType = "wResonant"
const ConnectorEnumTypeWInductive ConnectorEnumType = "wInductive"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateIdUseEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetCertificateIdUseEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetCertificateIdUseEnumType_2, v)
	}
	*j = GetCertificateIdUseEnumType_2(v)
	return nil
}

const ConnectorEnumTypePan ConnectorEnumType = "Pan"
const ConnectorEnumTypeOther3Ph ConnectorEnumType = "Other3Ph"
const ConnectorEnumTypeOther1PhOver16A ConnectorEnumType = "Other1PhOver16A"
const ConnectorEnumTypeOther1PhMax16A ConnectorEnumType = "Other1PhMax16A"
const ConnectorEnumTypeSType3 ConnectorEnumType = "sType3"

type CertificateHashDataChainType struct {
	// CertificateHashData corresponds to the JSON schema field "certificateHashData".
	CertificateHashData CertificateHashDataType_2 `json:"certificateHashData"`

	// CertificateType corresponds to the JSON schema field "certificateType".
	CertificateType GetCertificateIdUseEnumType_2 `json:"certificateType"`

	// ChildCertificateHashData corresponds to the JSON schema field
	// "childCertificateHashData".
	ChildCertificateHashData []CertificateHashDataType_2 `json:"childCertificateHashData,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_43 `json:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CertificateHashDataChainType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["certificateHashData"]; !ok || v == nil {
		return fmt.Errorf("field certificateHashData: required")
	}
	if v, ok := raw["certificateType"]; !ok || v == nil {
		return fmt.Errorf("field certificateType: required")
	}
	type Plain CertificateHashDataChainType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CertificateHashDataChainType(plain)
	return nil
}

const ConnectorEnumTypeSType2 ConnectorEnumType = "sType2"
const ConnectorEnumTypeSCEE77 ConnectorEnumType = "sCEE-7-7"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetCertificateIdUseEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetCertificateIdUseEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetCertificateIdUseEnumType_3, v)
	}
	*j = GetCertificateIdUseEnumType_3(v)
	return nil
}

const ConnectorEnumTypeSBS1361 ConnectorEnumType = "sBS1361"
const ConnectorEnumTypeS3093P32A ConnectorEnumType = "s309-3P-32A"
const ConnectorEnumTypeS3093P16A ConnectorEnumType = "s309-3P-16A"
const ConnectorEnumTypeS3091P32A ConnectorEnumType = "s309-1P-32A"
const ConnectorEnumTypeS3091P16A ConnectorEnumType = "s309-1P-16A"
const ConnectorEnumTypeCType2 ConnectorEnumType = "cType2"
const ConnectorEnumTypeCType1 ConnectorEnumType = "cType1"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetInstalledCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetInstalledCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetInstalledCertificateStatusEnumType, v)
	}
	*j = GetInstalledCertificateStatusEnumType(v)
	return nil
}

const ConnectorEnumTypeCTesla ConnectorEnumType = "cTesla"
const ConnectorEnumTypeCG105 ConnectorEnumType = "cG105"
const ConnectorEnumTypeCCCS2 ConnectorEnumType = "cCCS2"
const ConnectorEnumTypeCCCS1 ConnectorEnumType = "cCCS1"

// UnmarshalJSON implements json.Unmarshaler.
func (j *HashAlgorithmEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_HashAlgorithmEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_HashAlgorithmEnumType_9, v)
	}
	*j = HashAlgorithmEnumType_9(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConnectorEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ConnectorEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ConnectorEnumType, v)
	}
	*j = ConnectorEnumType(v)
	return nil
}

type ConnectorEnumType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType_4(plain)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType_4 struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_92 `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_17) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_17
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_17(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_92) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_92
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_92(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_91) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_91
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_91(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetInstalledCertificateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetInstalledCertificateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetInstalledCertificateStatusEnumType_1, v)
	}
	*j = GetInstalledCertificateStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReservationStatusUpdateRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reservationId"]; !ok || v == nil {
		return fmt.Errorf("field reservationId: required")
	}
	if v, ok := raw["reservationUpdateStatus"]; !ok || v == nil {
		return fmt.Errorf("field reservationUpdateStatus: required")
	}
	type Plain ReservationStatusUpdateRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ReservationStatusUpdateRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReservationUpdateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReservationUpdateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReservationUpdateStatusEnumType_1, v)
	}
	*j = ReservationUpdateStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReservationUpdateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReservationUpdateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReservationUpdateStatusEnumType, v)
	}
	*j = ReservationUpdateStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetInstalledCertificateIdsResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetInstalledCertificateIdsResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetInstalledCertificateIdsResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_90) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_90
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_90(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_44) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_44
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_44(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStopTransactionResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain RequestStopTransactionResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequestStopTransactionResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStartStopStatusEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RequestStartStopStatusEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RequestStartStopStatusEnumType_3, v)
	}
	*j = RequestStartStopStatusEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_45) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_45
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_45(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_27) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_27
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_27(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetLocalListVersionResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["versionNumber"]; !ok || v == nil {
		return fmt.Errorf("field versionNumber: required")
	}
	type Plain GetLocalListVersionResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetLocalListVersionResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStartStopStatusEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RequestStartStopStatusEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RequestStartStopStatusEnumType_2, v)
	}
	*j = RequestStartStopStatusEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_46) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_46
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_46(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_89) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_89
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_89(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStopTransactionRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["transactionId"]; !ok || v == nil {
		return fmt.Errorf("field transactionId: required")
	}
	type Plain RequestStopTransactionRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequestStopTransactionRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LogEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LogEnumType, v)
	}
	*j = LogEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_88) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_88
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_88(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStartTransactionResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain RequestStartTransactionResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequestStartTransactionResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStartStopStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RequestStartStopStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RequestStartStopStatusEnumType_1, v)
	}
	*j = RequestStartStopStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogParametersType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["remoteLocation"]; !ok || v == nil {
		return fmt.Errorf("field remoteLocation: required")
	}
	type Plain LogParametersType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = LogParametersType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_26) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_26
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_26(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStartStopStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RequestStartStopStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RequestStartStopStatusEnumType, v)
	}
	*j = RequestStartStopStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LogEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LogEnumType_1, v)
	}
	*j = LogEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_87) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_87
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_87(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RequestStartTransactionRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["remoteStartId"]; !ok || v == nil {
		return fmt.Errorf("field remoteStartId: required")
	}
	type Plain RequestStartTransactionRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RequestStartTransactionRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType_3, v)
	}
	*j = RecurrencyKindEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetLogRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["log"]; !ok || v == nil {
		return fmt.Errorf("field log: required")
	}
	if v, ok := raw["logType"]; !ok || v == nil {
		return fmt.Errorf("field logType: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain GetLogRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetLogRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["idToken"]; !ok || v == nil {
		return fmt.Errorf("field idToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain IdTokenType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = IdTokenType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_47) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_47
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_47(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_7, v)
	}
	*j = IdTokenEnumType_7(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *IdTokenEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_IdTokenEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_IdTokenEnumType_6, v)
	}
	*j = IdTokenEnumType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LogStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LogStatusEnumType, v)
	}
	*j = LogStatusEnumType(v)
	return nil
}

const CostKindEnumType_7_RenewableGenerationPercentage CostKindEnumType_7 = "RenewableGenerationPercentage"
const CostKindEnumType_7_RelativePricePercentage CostKindEnumType_7 = "RelativePricePercentage"
const CostKindEnumType_7_CarbonDioxideEmission CostKindEnumType_7 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_7, v)
	}
	*j = CostKindEnumType_7(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_18) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_18
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_18(plain)
	return nil
}

type CostKindEnumType_7 string

const ChargingRateUnitEnumType_11_A ChargingRateUnitEnumType_11 = "A"

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LogStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LogStatusEnumType_1, v)
	}
	*j = LogStatusEnumType_1(v)
	return nil
}

const ChargingRateUnitEnumType_11_W ChargingRateUnitEnumType_11 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_11) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_11 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_11, v)
	}
	*j = ChargingRateUnitEnumType_11(v)
	return nil
}

type ChargingRateUnitEnumType_11 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingProfileKind"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfileKind: required")
	}
	if v, ok := raw["chargingProfilePurpose"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfilePurpose: required")
	}
	if v, ok := raw["chargingSchedule"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedule: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["stackLevel"]; !ok || v == nil {
		return fmt.Errorf("field stackLevel: required")
	}
	type Plain ChargingProfileType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingProfileType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetLogResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetLogResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetLogResponseJson(plain)
	return nil
}

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
//
type ChargingProfileType_1 struct {
	// ChargingProfileKind corresponds to the JSON schema field "chargingProfileKind".
	ChargingProfileKind ChargingProfileKindEnumType_3 `json:"chargingProfileKind"`

	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose ChargingProfilePurposeEnumType_7 `json:"chargingProfilePurpose"`

	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule []ChargingScheduleType_3 `json:"chargingSchedule"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Id of ChargingProfile.
	//
	Id int `json:"id"`

	// RecurrencyKind corresponds to the JSON schema field "recurrencyKind".
	RecurrencyKind *RecurrencyKindEnumType_2 `json:"recurrencyKind,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Value determining level in hierarchy stack of profiles. Higher values have
	// precedence over lower values. Lowest level is 0.
	//
	StackLevel int `json:"stackLevel"`

	// SHALL only be included if ChargingProfilePurpose is set to TxProfile. The
	// transactionId is used to match the profile to a specific transaction.
	//
	TransactionId *string `json:"transactionId,omitempty"`

	// Charging_ Profile. Valid_ From. Date_ Time
	// urn:x-oca:ocpp:uid:1:569234
	// Point in time at which the profile starts to be valid. If absent, the profile
	// is valid as soon as it is received by the Charging Station.
	//
	ValidFrom *string `json:"validFrom,omitempty"`

	// Charging_ Profile. Valid_ To. Date_ Time
	// urn:x-oca:ocpp:uid:1:569235
	// Point in time at which the profile stops to be valid. If absent, the profile is
	// valid until it is replaced by another profile.
	//
	ValidTo *string `json:"validTo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_48) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_48
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_48(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType_2, v)
	}
	*j = RecurrencyKindEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_1(plain)
	return nil
}

// A physical or logical component
//
type ComponentType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_48 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_1 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingScheduleType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingRateUnit"]; !ok || v == nil {
		return fmt.Errorf("field chargingRateUnit: required")
	}
	if v, ok := raw["chargingSchedulePeriod"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedulePeriod: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain ChargingScheduleType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingScheduleType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType(plain)
	return nil
}

// Class to report components, variables and variable attributes and
// characteristics.
//
type ComponentVariableType struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_48 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable *VariableType `json:"variable,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentVariableType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	type Plain ComponentVariableType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentVariableType(plain)
	return nil
}

// Charging_ Schedule
// urn:x-oca:ocpp:uid:2:233256
// Charging schedule structure defines a list of charging periods, as used in:
// GetCompositeSchedule.conf and ChargingProfile.
//
type ChargingScheduleType_3 struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType_10 `json:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType_4 `json:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Charging_ Schedule. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569236
	// Duration of the charging schedule in seconds. If the duration is left empty,
	// the last period will continue indefinitely or until end of the transaction if
	// chargingProfilePurpose = TxProfile.
	//
	Duration *int `json:"duration,omitempty"`

	// Identifies the ChargingSchedule.
	//
	Id int `json:"id"`

	// Charging_ Schedule. Min_ Charging_ Rate. Numeric
	// urn:x-oca:ocpp:uid:1:569239
	// Minimum charging rate supported by the EV. The unit of measure is defined by
	// the chargingRateUnit. This parameter is intended to be used by a local smart
	// charging algorithm to optimize the power allocation for in the case a charging
	// process is inefficient at lower charging rates. Accepts at most one digit
	// fraction (e.g. 8.1)
	//
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`

	// SalesTariff corresponds to the JSON schema field "salesTariff".
	SalesTariff *SalesTariffType_3 `json:"salesTariff,omitempty"`

	// Charging_ Schedule. Start_ Schedule. Date_ Time
	// urn:x-oca:ocpp:uid:1:569237
	// Starting point of an absolute schedule. If absent the schedule will be relative
	// to start of charging.
	//
	StartSchedule *string `json:"startSchedule,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["salesTariffEntry"]; !ok || v == nil {
		return fmt.Errorf("field salesTariffEntry: required")
	}
	type Plain SalesTariffType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitoringCriterionEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitoringCriterionEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitoringCriterionEnumType, v)
	}
	*j = MonitoringCriterionEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffEntryType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["relativeTimeInterval"]; !ok || v == nil {
		return fmt.Errorf("field relativeTimeInterval: required")
	}
	type Plain SalesTariffEntryType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffEntryType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RelativeTimeIntervalType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["start"]; !ok || v == nil {
		return fmt.Errorf("field start: required")
	}
	type Plain RelativeTimeIntervalType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RelativeTimeIntervalType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsumptionCostType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cost"]; !ok || v == nil {
		return fmt.Errorf("field cost: required")
	}
	if v, ok := raw["startValue"]; !ok || v == nil {
		return fmt.Errorf("field startValue: required")
	}
	type Plain ConsumptionCostType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ConsumptionCostType_3(plain)
	return nil
}

// Consumption_ Cost
// urn:x-oca:ocpp:uid:2:233259
//
type ConsumptionCostType_3 struct {
	// Cost corresponds to the JSON schema field "cost".
	Cost []CostType_3 `json:"cost"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Consumption_ Cost. Start_ Value. Numeric
	// urn:x-oca:ocpp:uid:1:569246
	// The lowest level of consumption that defines the starting point of this
	// consumption block. The block interval extends to the start of the next
	// interval.
	//
	StartValue float64 `json:"startValue"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["amount"]; !ok || v == nil {
		return fmt.Errorf("field amount: required")
	}
	if v, ok := raw["costKind"]; !ok || v == nil {
		return fmt.Errorf("field costKind: required")
	}
	type Plain CostType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitoringCriterionEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitoringCriterionEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitoringCriterionEnumType_1, v)
	}
	*j = MonitoringCriterionEnumType_1(v)
	return nil
}

// Cost
// urn:x-oca:ocpp:uid:2:233258
//
type CostType_3 struct {
	// Cost. Amount. Amount
	// urn:x-oca:ocpp:uid:1:569244
	// The estimated or actual cost per kWh
	//
	Amount int `json:"amount"`

	// Cost. Amount_ Multiplier. Integer
	// urn:x-oca:ocpp:uid:1:569245
	// Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The
	// final value is determined by: amount * 10 ^ amountMultiplier
	//
	AmountMultiplier *int `json:"amountMultiplier,omitempty"`

	// CostKind corresponds to the JSON schema field "costKind".
	CostKind CostKindEnumType_6 `json:"costKind"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`
}

const CostKindEnumType_6_RenewableGenerationPercentage CostKindEnumType_6 = "RenewableGenerationPercentage"
const CostKindEnumType_6_RelativePricePercentage CostKindEnumType_6 = "RelativePricePercentage"
const CostKindEnumType_6_CarbonDioxideEmission CostKindEnumType_6 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetMonitoringReportRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain GetMonitoringReportRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetMonitoringReportRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_6, v)
	}
	*j = CostKindEnumType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_49) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_49
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_49(plain)
	return nil
}

type CostKindEnumType_6 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit: required")
	}
	if v, ok := raw["startPeriod"]; !ok || v == nil {
		return fmt.Errorf("field startPeriod: required")
	}
	type Plain ChargingSchedulePeriodType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType_4(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType_2, v)
	}
	*j = GenericDeviceModelStatusEnumType_2(v)
	return nil
}

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
//
type ChargingSchedulePeriodType_4 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod"`
}

const ChargingRateUnitEnumType_10_A ChargingRateUnitEnumType_10 = "A"
const ChargingRateUnitEnumType_10_W ChargingRateUnitEnumType_10 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_10) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_10 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_10, v)
	}
	*j = ChargingRateUnitEnumType_10(v)
	return nil
}

type ChargingRateUnitEnumType_10 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_19) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_19
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_19(plain)
	return nil
}

const ChargingProfilePurposeEnumType_7_TxProfile ChargingProfilePurposeEnumType_7 = "TxProfile"
const ChargingProfilePurposeEnumType_7_TxDefaultProfile ChargingProfilePurposeEnumType_7 = "TxDefaultProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType_3, v)
	}
	*j = GenericDeviceModelStatusEnumType_3(v)
	return nil
}

const ChargingProfilePurposeEnumType_7_ChargingStationMaxProfile ChargingProfilePurposeEnumType_7 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_7_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_7 = "ChargingStationExternalConstraints"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_7, v)
	}
	*j = ChargingProfilePurposeEnumType_7(v)
	return nil
}

type ChargingProfilePurposeEnumType_7 string

const ChargingProfileKindEnumType_3_Relative ChargingProfileKindEnumType_3 = "Relative"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetMonitoringReportResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetMonitoringReportResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetMonitoringReportResponseJson(plain)
	return nil
}

type ComponentCriterionEnumType string

const ChargingProfileKindEnumType_3_Recurring ChargingProfileKindEnumType_3 = "Recurring"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentCriterionEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ComponentCriterionEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ComponentCriterionEnumType, v)
	}
	*j = ComponentCriterionEnumType(v)
	return nil
}

const ComponentCriterionEnumTypeActive ComponentCriterionEnumType = "Active"
const ComponentCriterionEnumTypeAvailable ComponentCriterionEnumType = "Available"
const ComponentCriterionEnumTypeEnabled ComponentCriterionEnumType = "Enabled"
const ComponentCriterionEnumTypeProblem ComponentCriterionEnumType = "Problem"
const ChargingProfileKindEnumType_3_Absolute ChargingProfileKindEnumType_3 = "Absolute"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_50) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_50
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_50(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType_3, v)
	}
	*j = ChargingProfileKindEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_2(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_50 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_2 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_1(plain)
	return nil
}

type ChargingProfileKindEnumType_3 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_1(plain)
	return nil
}

// Class to report components, variables and variable attributes and
// characteristics.
//
type ComponentVariableType_1 struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType_1 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_50 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable *VariableType_1 `json:"variable,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentVariableType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	type Plain ComponentVariableType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentVariableType_1(plain)
	return nil
}

type ComponentCriterionEnumType_1 string

const ChargingProfilePurposeEnumType_6_TxProfile ChargingProfilePurposeEnumType_6 = "TxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentCriterionEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ComponentCriterionEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ComponentCriterionEnumType_1, v)
	}
	*j = ComponentCriterionEnumType_1(v)
	return nil
}

const ComponentCriterionEnumType_1_Active ComponentCriterionEnumType_1 = "Active"
const ComponentCriterionEnumType_1_Available ComponentCriterionEnumType_1 = "Available"
const ComponentCriterionEnumType_1_Enabled ComponentCriterionEnumType_1 = "Enabled"
const ComponentCriterionEnumType_1_Problem ComponentCriterionEnumType_1 = "Problem"
const ChargingProfilePurposeEnumType_6_TxDefaultProfile ChargingProfilePurposeEnumType_6 = "TxDefaultProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetReportRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain GetReportRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetReportRequestJson(plain)
	return nil
}

const ChargingProfilePurposeEnumType_6_ChargingStationMaxProfile ChargingProfilePurposeEnumType_6 = "ChargingStationMaxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_51) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_51
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_51(plain)
	return nil
}

const ChargingProfilePurposeEnumType_6_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_6 = "ChargingStationExternalConstraints"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_6, v)
	}
	*j = ChargingProfilePurposeEnumType_6(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType_4, v)
	}
	*j = GenericDeviceModelStatusEnumType_4(v)
	return nil
}

type ChargingProfilePurposeEnumType_6 string

const ChargingProfileKindEnumType_2_Relative ChargingProfileKindEnumType_2 = "Relative"
const CostKindEnumTypeCarbonDioxideEmission CostKindEnumType = "CarbonDioxideEmission"
const ChargingProfileKindEnumType_2_Absolute ChargingProfileKindEnumType_2 = "Absolute"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType_2, v)
	}
	*j = ChargingProfileKindEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_20) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_20
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_20(plain)
	return nil
}

type ChargingProfileKindEnumType_2 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *AdditionalInfoType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["additionalIdToken"]; !ok || v == nil {
		return fmt.Errorf("field additionalIdToken: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain AdditionalInfoType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AdditionalInfoType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericDeviceModelStatusEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericDeviceModelStatusEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericDeviceModelStatusEnumType_5, v)
	}
	*j = GenericDeviceModelStatusEnumType_5(v)
	return nil
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type AdditionalInfoType_3 struct {
	// This field specifies the additional IdToken.
	//
	AdditionalIdToken string `json:"additionalIdToken"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// This defines the type of the additionalIdToken. This is a custom type, so the
	// implementation needs to be agreed upon by all involved parties.
	//
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_86) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_86
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_86(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_85) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_85
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_85(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportChargingProfilesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingLimitSource"]; !ok || v == nil {
		return fmt.Errorf("field chargingLimitSource: required")
	}
	if v, ok := raw["chargingProfile"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfile: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain ReportChargingProfilesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = ReportChargingProfilesRequestJson(plain)
	return nil
}

const ChargingLimitSourceEnumType_7_CSO ChargingLimitSourceEnumType_7 = "CSO"

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetReportResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain GetReportResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetReportResponseJson(plain)
	return nil
}

const ChargingLimitSourceEnumType_7_SO ChargingLimitSourceEnumType_7 = "SO"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_52) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_52
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_52(plain)
	return nil
}

const ChargingLimitSourceEnumType_7_Other ChargingLimitSourceEnumType_7 = "Other"
const ChargingLimitSourceEnumType_7_EMS ChargingLimitSourceEnumType_7 = "EMS"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_53) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_53
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_53(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType_7, v)
	}
	*j = ChargingLimitSourceEnumType_7(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetTransactionStatusResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["messagesInQueue"]; !ok || v == nil {
		return fmt.Errorf("field messagesInQueue: required")
	}
	type Plain GetTransactionStatusResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetTransactionStatusResponseJson(plain)
	return nil
}

type AttributeEnumType string

type ChargingLimitSourceEnumType_7 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType, v)
	}
	*j = AttributeEnumType(v)
	return nil
}

const AttributeEnumTypeActual AttributeEnumType = "Actual"
const AttributeEnumTypeTarget AttributeEnumType = "Target"
const AttributeEnumTypeMinSet AttributeEnumType = "MinSet"
const AttributeEnumTypeMaxSet AttributeEnumType = "MaxSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType_1, v)
	}
	*j = RecurrencyKindEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_54) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_54
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_54(plain)
	return nil
}

const CostKindEnumType_5_RenewableGenerationPercentage CostKindEnumType_5 = "RenewableGenerationPercentage"

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_3(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_2 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_54 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_3 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_2(plain)
	return nil
}

type AttributeEnumType_1 string

const CostKindEnumType_5_RelativePricePercentage CostKindEnumType_5 = "RelativePricePercentage"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_1, v)
	}
	*j = AttributeEnumType_1(v)
	return nil
}

const AttributeEnumType_1_Actual AttributeEnumType_1 = "Actual"
const AttributeEnumType_1_Target AttributeEnumType_1 = "Target"
const AttributeEnumType_1_MinSet AttributeEnumType_1 = "MinSet"
const AttributeEnumType_1_MaxSet AttributeEnumType_1 = "MaxSet"
const CostKindEnumType_5_CarbonDioxideEmission CostKindEnumType_5 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_2(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_5, v)
	}
	*j = CostKindEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariableDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	type Plain GetVariableDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetVariableDataType(plain)
	return nil
}

type CostKindEnumType_5 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariablesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["getVariableData"]; !ok || v == nil {
		return fmt.Errorf("field getVariableData: required")
	}
	type Plain GetVariablesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetVariablesRequestJson(plain)
	return nil
}

type AttributeEnumType_2 string

const ChargingRateUnitEnumType_9_A ChargingRateUnitEnumType_9 = "A"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_2, v)
	}
	*j = AttributeEnumType_2(v)
	return nil
}

const AttributeEnumType_2_Actual AttributeEnumType_2 = "Actual"
const AttributeEnumType_2_Target AttributeEnumType_2 = "Target"
const AttributeEnumType_2_MinSet AttributeEnumType_2 = "MinSet"
const AttributeEnumType_2_MaxSet AttributeEnumType_2 = "MaxSet"
const ChargingRateUnitEnumType_9_W ChargingRateUnitEnumType_9 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_55) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_55
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_55(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_9) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_9 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_9, v)
	}
	*j = ChargingRateUnitEnumType_9(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_4(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_3 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_55 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_4 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_3(plain)
	return nil
}

type ChargingRateUnitEnumType_9 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingProfileKind"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfileKind: required")
	}
	if v, ok := raw["chargingProfilePurpose"]; !ok || v == nil {
		return fmt.Errorf("field chargingProfilePurpose: required")
	}
	if v, ok := raw["chargingSchedule"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedule: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["stackLevel"]; !ok || v == nil {
		return fmt.Errorf("field stackLevel: required")
	}
	type Plain ChargingProfileType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingProfileType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariableStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetVariableStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetVariableStatusEnumType, v)
	}
	*j = GetVariableStatusEnumType(v)
	return nil
}

// Charging_ Profile
// urn:x-oca:ocpp:uid:2:233255
// A ChargingProfile consists of ChargingSchedule, describing the amount of power
// or current that can be delivered per time interval.
//
type ChargingProfileType struct {
	// ChargingProfileKind corresponds to the JSON schema field "chargingProfileKind".
	ChargingProfileKind ChargingProfileKindEnumType_1 `json:"chargingProfileKind"`

	// ChargingProfilePurpose corresponds to the JSON schema field
	// "chargingProfilePurpose".
	ChargingProfilePurpose ChargingProfilePurposeEnumType_5 `json:"chargingProfilePurpose"`

	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule []ChargingScheduleType_2 `json:"chargingSchedule"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Id of ChargingProfile.
	//
	Id int `json:"id"`

	// RecurrencyKind corresponds to the JSON schema field "recurrencyKind".
	RecurrencyKind *RecurrencyKindEnumType `json:"recurrencyKind,omitempty"`

	// Charging_ Profile. Stack_ Level. Counter
	// urn:x-oca:ocpp:uid:1:569230
	// Value determining level in hierarchy stack of profiles. Higher values have
	// precedence over lower values. Lowest level is 0.
	//
	StackLevel int `json:"stackLevel"`

	// SHALL only be included if ChargingProfilePurpose is set to TxProfile. The
	// transactionId is used to match the profile to a specific transaction.
	//
	TransactionId *string `json:"transactionId,omitempty"`

	// Charging_ Profile. Valid_ From. Date_ Time
	// urn:x-oca:ocpp:uid:1:569234
	// Point in time at which the profile starts to be valid. If absent, the profile
	// is valid as soon as it is received by the Charging Station.
	//
	ValidFrom *string `json:"validFrom,omitempty"`

	// Charging_ Profile. Valid_ To. Date_ Time
	// urn:x-oca:ocpp:uid:1:569235
	// Point in time at which the profile stops to be valid. If absent, the profile is
	// valid until it is replaced by another profile.
	//
	ValidTo *string `json:"validTo,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RecurrencyKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_RecurrencyKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_RecurrencyKindEnumType, v)
	}
	*j = RecurrencyKindEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingScheduleType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingRateUnit"]; !ok || v == nil {
		return fmt.Errorf("field chargingRateUnit: required")
	}
	if v, ok := raw["chargingSchedulePeriod"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedulePeriod: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain ChargingScheduleType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingScheduleType_2(plain)
	return nil
}

// Charging_ Schedule
// urn:x-oca:ocpp:uid:2:233256
// Charging schedule structure defines a list of charging periods, as used in:
// GetCompositeSchedule.conf and ChargingProfile.
//
type ChargingScheduleType_2 struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType_8 `json:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType_3 `json:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// Charging_ Schedule. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569236
	// Duration of the charging schedule in seconds. If the duration is left empty,
	// the last period will continue indefinitely or until end of the transaction if
	// chargingProfilePurpose = TxProfile.
	//
	Duration *int `json:"duration,omitempty"`

	// Identifies the ChargingSchedule.
	//
	Id int `json:"id"`

	// Charging_ Schedule. Min_ Charging_ Rate. Numeric
	// urn:x-oca:ocpp:uid:1:569239
	// Minimum charging rate supported by the EV. The unit of measure is defined by
	// the chargingRateUnit. This parameter is intended to be used by a local smart
	// charging algorithm to optimize the power allocation for in the case a charging
	// process is inefficient at lower charging rates. Accepts at most one digit
	// fraction (e.g. 8.1)
	//
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`

	// SalesTariff corresponds to the JSON schema field "salesTariff".
	SalesTariff *SalesTariffType_2 `json:"salesTariff,omitempty"`

	// Charging_ Schedule. Start_ Schedule. Date_ Time
	// urn:x-oca:ocpp:uid:1:569237
	// Starting point of an absolute schedule. If absent the schedule will be relative
	// to start of charging.
	//
	StartSchedule *string `json:"startSchedule,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["salesTariffEntry"]; !ok || v == nil {
		return fmt.Errorf("field salesTariffEntry: required")
	}
	type Plain SalesTariffType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffType_2(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffEntryType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["relativeTimeInterval"]; !ok || v == nil {
		return fmt.Errorf("field relativeTimeInterval: required")
	}
	type Plain SalesTariffEntryType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffEntryType_2(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_21) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_21
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_21(plain)
	return nil
}

type AttributeEnumType_3 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *RelativeTimeIntervalType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["start"]; !ok || v == nil {
		return fmt.Errorf("field start: required")
	}
	type Plain RelativeTimeIntervalType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RelativeTimeIntervalType_2(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_3, v)
	}
	*j = AttributeEnumType_3(v)
	return nil
}

const AttributeEnumType_3_Actual AttributeEnumType_3 = "Actual"
const AttributeEnumType_3_Target AttributeEnumType_3 = "Target"
const AttributeEnumType_3_MinSet AttributeEnumType_3 = "MinSet"
const AttributeEnumType_3_MaxSet AttributeEnumType_3 = "MaxSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsumptionCostType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cost"]; !ok || v == nil {
		return fmt.Errorf("field cost: required")
	}
	if v, ok := raw["startValue"]; !ok || v == nil {
		return fmt.Errorf("field startValue: required")
	}
	type Plain ConsumptionCostType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ConsumptionCostType_2(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_3(plain)
	return nil
}

// Consumption_ Cost
// urn:x-oca:ocpp:uid:2:233259
//
type ConsumptionCostType_2 struct {
	// Cost corresponds to the JSON schema field "cost".
	Cost []CostType_2 `json:"cost"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// Consumption_ Cost. Start_ Value. Numeric
	// urn:x-oca:ocpp:uid:1:569246
	// The lowest level of consumption that defines the starting point of this
	// consumption block. The block interval extends to the start of the next
	// interval.
	//
	StartValue float64 `json:"startValue"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariableResultType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["attributeStatus"]; !ok || v == nil {
		return fmt.Errorf("field attributeStatus: required")
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	type Plain GetVariableResultType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetVariableResultType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["amount"]; !ok || v == nil {
		return fmt.Errorf("field amount: required")
	}
	if v, ok := raw["costKind"]; !ok || v == nil {
		return fmt.Errorf("field costKind: required")
	}
	type Plain CostType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostType_2(plain)
	return nil
}

// Cost
// urn:x-oca:ocpp:uid:2:233258
//
type CostType_2 struct {
	// Cost. Amount. Amount
	// urn:x-oca:ocpp:uid:1:569244
	// The estimated or actual cost per kWh
	//
	Amount int `json:"amount"`

	// Cost. Amount_ Multiplier. Integer
	// urn:x-oca:ocpp:uid:1:569245
	// Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The
	// final value is determined by: amount * 10 ^ amountMultiplier
	//
	AmountMultiplier *int `json:"amountMultiplier,omitempty"`

	// CostKind corresponds to the JSON schema field "costKind".
	CostKind CostKindEnumType_4 `json:"costKind"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariableStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GetVariableStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GetVariableStatusEnumType_1, v)
	}
	*j = GetVariableStatusEnumType_1(v)
	return nil
}

const CostKindEnumType_4_RenewableGenerationPercentage CostKindEnumType_4 = "RenewableGenerationPercentage"
const CostKindEnumType_4_RelativePricePercentage CostKindEnumType_4 = "RelativePricePercentage"
const CostKindEnumType_4_CarbonDioxideEmission CostKindEnumType_4 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_4, v)
	}
	*j = CostKindEnumType_4(v)
	return nil
}

type CostKindEnumType_4 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType_3) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit: required")
	}
	if v, ok := raw["startPeriod"]; !ok || v == nil {
		return fmt.Errorf("field startPeriod: required")
	}
	type Plain ChargingSchedulePeriodType_3
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType_3(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GetVariablesResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["getVariableResult"]; !ok || v == nil {
		return fmt.Errorf("field getVariableResult: required")
	}
	type Plain GetVariablesResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = GetVariablesResponseJson(plain)
	return nil
}

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
//
type ChargingSchedulePeriodType_3 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_56) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_56
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_56(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_84) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_84
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_84(plain)
	return nil
}

const ChargingRateUnitEnumType_8_A ChargingRateUnitEnumType_8 = "A"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_57) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_57
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_57(plain)
	return nil
}

const ChargingRateUnitEnumType_8_W ChargingRateUnitEnumType_8 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *HeartbeatResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["currentTime"]; !ok || v == nil {
		return fmt.Errorf("field currentTime: required")
	}
	type Plain HeartbeatResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = HeartbeatResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_8) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_8 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_8, v)
	}
	*j = ChargingRateUnitEnumType_8(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_58) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_58
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_58(plain)
	return nil
}

type ChargingRateUnitEnumType_8 string

const ChargingProfilePurposeEnumType_5_TxProfile ChargingProfilePurposeEnumType_5 = "TxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *InstallCertificateUseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_InstallCertificateUseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_InstallCertificateUseEnumType, v)
	}
	*j = InstallCertificateUseEnumType(v)
	return nil
}

const ChargingProfilePurposeEnumType_5_TxDefaultProfile ChargingProfilePurposeEnumType_5 = "TxDefaultProfile"
const ChargingProfilePurposeEnumType_5_ChargingStationMaxProfile ChargingProfilePurposeEnumType_5 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_5_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_5 = "ChargingStationExternalConstraints"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_5, v)
	}
	*j = ChargingProfilePurposeEnumType_5(v)
	return nil
}

type ChargingProfilePurposeEnumType_5 string

const ChargingProfileKindEnumType_1_Relative ChargingProfileKindEnumType_1 = "Relative"

// UnmarshalJSON implements json.Unmarshaler.
func (j *InstallCertificateUseEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_InstallCertificateUseEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_InstallCertificateUseEnumType_1, v)
	}
	*j = InstallCertificateUseEnumType_1(v)
	return nil
}

const ChargingProfileKindEnumType_1_Recurring ChargingProfileKindEnumType_1 = "Recurring"
const ChargingProfileKindEnumType_1_Absolute ChargingProfileKindEnumType_1 = "Absolute"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType_1, v)
	}
	*j = ChargingProfileKindEnumType_1(v)
	return nil
}

type ChargingProfileKindEnumType_1 string

const ChargingProfilePurposeEnumType_4_TxProfile ChargingProfilePurposeEnumType_4 = "TxProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *InstallCertificateRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["certificate"]; !ok || v == nil {
		return fmt.Errorf("field certificate: required")
	}
	if v, ok := raw["certificateType"]; !ok || v == nil {
		return fmt.Errorf("field certificateType: required")
	}
	type Plain InstallCertificateRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = InstallCertificateRequestJson(plain)
	return nil
}

const ChargingProfilePurposeEnumType_4_TxDefaultProfile ChargingProfilePurposeEnumType_4 = "TxDefaultProfile"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_59) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_59
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_59(plain)
	return nil
}

const ChargingProfilePurposeEnumType_4_ChargingStationMaxProfile ChargingProfilePurposeEnumType_4 = "ChargingStationMaxProfile"
const ChargingProfilePurposeEnumType_4_ChargingStationExternalConstraints ChargingProfilePurposeEnumType_4 = "ChargingStationExternalConstraints"

// UnmarshalJSON implements json.Unmarshaler.
func (j *InstallCertificateStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_InstallCertificateStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_InstallCertificateStatusEnumType, v)
	}
	*j = InstallCertificateStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfilePurposeEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfilePurposeEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfilePurposeEnumType_4, v)
	}
	*j = ChargingProfilePurposeEnumType_4(v)
	return nil
}

type ChargingProfilePurposeEnumType_4 string

const ChargingProfileKindEnumTypeRelative ChargingProfileKindEnumType = "Relative"
const ChargingProfileKindEnumTypeRecurring ChargingProfileKindEnumType = "Recurring"

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_22) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_22
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_22(plain)
	return nil
}

const ChargingProfileKindEnumTypeAbsolute ChargingProfileKindEnumType = "Absolute"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingProfileKindEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingProfileKindEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingProfileKindEnumType, v)
	}
	*j = ChargingProfileKindEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *InstallCertificateStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_InstallCertificateStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_InstallCertificateStatusEnumType_1, v)
	}
	*j = InstallCertificateStatusEnumType_1(v)
	return nil
}

type ChargingProfileKindEnumType string

const ChargingLimitSourceEnumType_6_CSO ChargingLimitSourceEnumType_6 = "CSO"
const ChargingLimitSourceEnumType_6_SO ChargingLimitSourceEnumType_6 = "SO"
const ChargingLimitSourceEnumType_6_Other ChargingLimitSourceEnumType_6 = "Other"

// UnmarshalJSON implements json.Unmarshaler.
func (j *InstallCertificateResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain InstallCertificateResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = InstallCertificateResponseJson(plain)
	return nil
}

const ChargingLimitSourceEnumType_6_EMS ChargingLimitSourceEnumType_6 = "EMS"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_60) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_60
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_60(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType_6, v)
	}
	*j = ChargingLimitSourceEnumType_6(v)
	return nil
}

type ChargingLimitSourceEnumType_6 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *UploadLogStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UploadLogStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UploadLogStatusEnumType, v)
	}
	*j = UploadLogStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_83) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_83
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_83(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PublishFirmwareStatusNotificationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain PublishFirmwareStatusNotificationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PublishFirmwareStatusNotificationRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PublishFirmwareStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PublishFirmwareStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PublishFirmwareStatusEnumType_1, v)
	}
	*j = PublishFirmwareStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PublishFirmwareStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PublishFirmwareStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PublishFirmwareStatusEnumType, v)
	}
	*j = PublishFirmwareStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_82) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_82
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_82(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PublishFirmwareResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain PublishFirmwareResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PublishFirmwareResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_5, v)
	}
	*j = GenericStatusEnumType_5(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_25) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_25
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_25(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_4, v)
	}
	*j = GenericStatusEnumType_4(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_81) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_81
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_81(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UploadLogStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_UploadLogStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_UploadLogStatusEnumType_1, v)
	}
	*j = UploadLogStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PublishFirmwareRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["checksum"]; !ok || v == nil {
		return fmt.Errorf("field checksum: required")
	}
	if v, ok := raw["location"]; !ok || v == nil {
		return fmt.Errorf("field location: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain PublishFirmwareRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = PublishFirmwareRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_80) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_80
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_80(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_79) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_79
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_79(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyReportRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["generatedAt"]; !ok || v == nil {
		return fmt.Errorf("field generatedAt: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	if v, ok := raw["seqNo"]; !ok || v == nil {
		return fmt.Errorf("field seqNo: required")
	}
	type Plain NotifyReportRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = NotifyReportRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReportDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	if v, ok := raw["variableAttribute"]; !ok || v == nil {
		return fmt.Errorf("field variableAttribute: required")
	}
	type Plain ReportDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ReportDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableCharacteristicsType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["dataType"]; !ok || v == nil {
		return fmt.Errorf("field dataType: required")
	}
	if v, ok := raw["supportsMonitoring"]; !ok || v == nil {
		return fmt.Errorf("field supportsMonitoring: required")
	}
	type Plain VariableCharacteristicsType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableCharacteristicsType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DataEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DataEnumType_1, v)
	}
	*j = DataEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableAttributeType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain VariableAttributeType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["constant"]; !ok || v == nil {
		plain.Constant = false
	}
	if v, ok := raw["persistent"]; !ok || v == nil {
		plain.Persistent = false
	}
	*j = VariableAttributeType(plain)
	return nil
}

const AttributeEnumType_5_MaxSet AttributeEnumType_5 = "MaxSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *LogStatusNotificationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain LogStatusNotificationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = LogStatusNotificationRequestJson(plain)
	return nil
}

const AttributeEnumType_5_MinSet AttributeEnumType_5 = "MinSet"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_61) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_61
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_61(plain)
	return nil
}

const AttributeEnumType_5_Target AttributeEnumType_5 = "Target"
const AttributeEnumType_5_Actual AttributeEnumType_5 = "Actual"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_62) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_62
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_62(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_5, v)
	}
	*j = AttributeEnumType_5(v)
	return nil
}

type AttributeEnumType_5 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *LocationEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LocationEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LocationEnumType, v)
	}
	*j = LocationEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MutabilityEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MutabilityEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MutabilityEnumType_1, v)
	}
	*j = MutabilityEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_6) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_6
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_6(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MutabilityEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MutabilityEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MutabilityEnumType, v)
	}
	*j = MutabilityEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DataEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DataEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DataEnumType, v)
	}
	*j = DataEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_7) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_7
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_7(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_7 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_78 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_8 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_8) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_8
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_8(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeasurandEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MeasurandEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MeasurandEnumType, v)
	}
	*j = MeasurandEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_78) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_78
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_78(plain)
	return nil
}

const AttributeEnumType_4_MaxSet AttributeEnumType_4 = "MaxSet"
const AttributeEnumType_4_MinSet AttributeEnumType_4 = "MinSet"
const AttributeEnumType_4_Target AttributeEnumType_4 = "Target"
const AttributeEnumType_4_Actual AttributeEnumType_4 = "Actual"

// UnmarshalJSON implements json.Unmarshaler.
func (j *AttributeEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AttributeEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AttributeEnumType_4, v)
	}
	*j = AttributeEnumType_4(v)
	return nil
}

type AttributeEnumType_4 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_77) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_77
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_77(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyMonitoringReportRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["generatedAt"]; !ok || v == nil {
		return fmt.Errorf("field generatedAt: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	if v, ok := raw["seqNo"]; !ok || v == nil {
		return fmt.Errorf("field seqNo: required")
	}
	type Plain NotifyMonitoringReportRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = NotifyMonitoringReportRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitoringDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	if v, ok := raw["variableMonitoring"]; !ok || v == nil {
		return fmt.Errorf("field variableMonitoring: required")
	}
	type Plain MonitoringDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MonitoringDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableMonitoringType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["severity"]; !ok || v == nil {
		return fmt.Errorf("field severity: required")
	}
	if v, ok := raw["transaction"]; !ok || v == nil {
		return fmt.Errorf("field transaction: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value: required")
	}
	type Plain VariableMonitoringType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableMonitoringType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType_1, v)
	}
	*j = MonitorEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_5(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MonitorEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MonitorEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MonitorEnumType, v)
	}
	*j = MonitorEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_6) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_6
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_6(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_6 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_76 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_7 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_7) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_7
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_7(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_76) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_76
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_76(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_75) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_75
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_75(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyEventRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["eventData"]; !ok || v == nil {
		return fmt.Errorf("field eventData: required")
	}
	if v, ok := raw["generatedAt"]; !ok || v == nil {
		return fmt.Errorf("field generatedAt: required")
	}
	if v, ok := raw["seqNo"]; !ok || v == nil {
		return fmt.Errorf("field seqNo: required")
	}
	type Plain NotifyEventRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = NotifyEventRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventTriggerEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventTriggerEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventTriggerEnumType_1, v)
	}
	*j = EventTriggerEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventNotificationEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventNotificationEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventNotificationEnumType_1, v)
	}
	*j = EventNotificationEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventDataType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["actualValue"]; !ok || v == nil {
		return fmt.Errorf("field actualValue: required")
	}
	if v, ok := raw["component"]; !ok || v == nil {
		return fmt.Errorf("field component: required")
	}
	if v, ok := raw["eventId"]; !ok || v == nil {
		return fmt.Errorf("field eventId: required")
	}
	if v, ok := raw["eventNotificationType"]; !ok || v == nil {
		return fmt.Errorf("field eventNotificationType: required")
	}
	if v, ok := raw["timestamp"]; !ok || v == nil {
		return fmt.Errorf("field timestamp: required")
	}
	if v, ok := raw["trigger"]; !ok || v == nil {
		return fmt.Errorf("field trigger: required")
	}
	if v, ok := raw["variable"]; !ok || v == nil {
		return fmt.Errorf("field variable: required")
	}
	type Plain EventDataType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EventDataType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *VariableType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain VariableType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = VariableType_4(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventTriggerEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventTriggerEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventTriggerEnumType, v)
	}
	*j = EventTriggerEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EventNotificationEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EventNotificationEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EventNotificationEnumType, v)
	}
	*j = EventNotificationEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_5(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReadingContextEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReadingContextEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReadingContextEnumType, v)
	}
	*j = ReadingContextEnumType(v)
	return nil
}

// A physical or logical component
//
type ComponentType_5 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_74 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_6 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_6) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_6
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_6(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_74) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_74
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_74(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyEVChargingScheduleResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain NotifyEVChargingScheduleResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NotifyEVChargingScheduleResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_3, v)
	}
	*j = GenericStatusEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_24) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_24
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_24(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *GenericStatusEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_GenericStatusEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_GenericStatusEnumType_2, v)
	}
	*j = GenericStatusEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_73) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_73
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_73(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyEVChargingScheduleRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingSchedule"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedule: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	if v, ok := raw["timeBase"]; !ok || v == nil {
		return fmt.Errorf("field timeBase: required")
	}
	type Plain NotifyEVChargingScheduleRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NotifyEVChargingScheduleRequestJson(plain)
	return nil
}

const CostKindEnumType_3_RenewableGenerationPercentage CostKindEnumType_3 = "RenewableGenerationPercentage"

// UnmarshalJSON implements json.Unmarshaler.
func (j *LocationEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_LocationEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_LocationEnumType_1, v)
	}
	*j = LocationEnumType_1(v)
	return nil
}

const CostKindEnumType_3_RelativePricePercentage CostKindEnumType_3 = "RelativePricePercentage"
const CostKindEnumType_3_CarbonDioxideEmission CostKindEnumType_3 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_3, v)
	}
	*j = CostKindEnumType_3(v)
	return nil
}

type CostKindEnumType_3 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingScheduleType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingRateUnit"]; !ok || v == nil {
		return fmt.Errorf("field chargingRateUnit: required")
	}
	if v, ok := raw["chargingSchedulePeriod"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedulePeriod: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain ChargingScheduleType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingScheduleType_1(plain)
	return nil
}

// Charging_ Schedule
// urn:x-oca:ocpp:uid:2:233256
// Charging schedule structure defines a list of charging periods, as used in:
// GetCompositeSchedule.conf and ChargingProfile.
//
type ChargingScheduleType_1 struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType_7 `json:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType_2 `json:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`

	// Charging_ Schedule. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569236
	// Duration of the charging schedule in seconds. If the duration is left empty,
	// the last period will continue indefinitely or until end of the transaction if
	// chargingProfilePurpose = TxProfile.
	//
	Duration *int `json:"duration,omitempty"`

	// Identifies the ChargingSchedule.
	//
	Id int `json:"id"`

	// Charging_ Schedule. Min_ Charging_ Rate. Numeric
	// urn:x-oca:ocpp:uid:1:569239
	// Minimum charging rate supported by the EV. The unit of measure is defined by
	// the chargingRateUnit. This parameter is intended to be used by a local smart
	// charging algorithm to optimize the power allocation for in the case a charging
	// process is inefficient at lower charging rates. Accepts at most one digit
	// fraction (e.g. 8.1)
	//
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`

	// SalesTariff corresponds to the JSON schema field "salesTariff".
	SalesTariff *SalesTariffType_1 `json:"salesTariff,omitempty"`

	// Charging_ Schedule. Start_ Schedule. Date_ Time
	// urn:x-oca:ocpp:uid:1:569237
	// Starting point of an absolute schedule. If absent the schedule will be relative
	// to start of charging.
	//
	StartSchedule *string `json:"startSchedule,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["salesTariffEntry"]; !ok || v == nil {
		return fmt.Errorf("field salesTariffEntry: required")
	}
	type Plain SalesTariffType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeasurandEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MeasurandEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MeasurandEnumType_1, v)
	}
	*j = MeasurandEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffEntryType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["relativeTimeInterval"]; !ok || v == nil {
		return fmt.Errorf("field relativeTimeInterval: required")
	}
	type Plain SalesTariffEntryType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffEntryType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RelativeTimeIntervalType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["start"]; !ok || v == nil {
		return fmt.Errorf("field start: required")
	}
	type Plain RelativeTimeIntervalType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RelativeTimeIntervalType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsumptionCostType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cost"]; !ok || v == nil {
		return fmt.Errorf("field cost: required")
	}
	if v, ok := raw["startValue"]; !ok || v == nil {
		return fmt.Errorf("field startValue: required")
	}
	type Plain ConsumptionCostType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ConsumptionCostType_1(plain)
	return nil
}

// Consumption_ Cost
// urn:x-oca:ocpp:uid:2:233259
//
type ConsumptionCostType_1 struct {
	// Cost corresponds to the JSON schema field "cost".
	Cost []CostType_1 `json:"cost"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`

	// Consumption_ Cost. Start_ Value. Numeric
	// urn:x-oca:ocpp:uid:1:569246
	// The lowest level of consumption that defines the starting point of this
	// consumption block. The block interval extends to the start of the next
	// interval.
	//
	StartValue float64 `json:"startValue"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["amount"]; !ok || v == nil {
		return fmt.Errorf("field amount: required")
	}
	if v, ok := raw["costKind"]; !ok || v == nil {
		return fmt.Errorf("field costKind: required")
	}
	type Plain CostType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostType_1(plain)
	return nil
}

// Cost
// urn:x-oca:ocpp:uid:2:233258
//
type CostType_1 struct {
	// Cost. Amount. Amount
	// urn:x-oca:ocpp:uid:1:569244
	// The estimated or actual cost per kWh
	//
	Amount int `json:"amount"`

	// Cost. Amount_ Multiplier. Integer
	// urn:x-oca:ocpp:uid:1:569245
	// Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The
	// final value is determined by: amount * 10 ^ amountMultiplier
	//
	AmountMultiplier *int `json:"amountMultiplier,omitempty"`

	// CostKind corresponds to the JSON schema field "costKind".
	CostKind CostKindEnumType_2 `json:"costKind"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`
}

const CostKindEnumType_2_RenewableGenerationPercentage CostKindEnumType_2 = "RenewableGenerationPercentage"
const CostKindEnumType_2_RelativePricePercentage CostKindEnumType_2 = "RelativePricePercentage"
const CostKindEnumType_2_CarbonDioxideEmission CostKindEnumType_2 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_2, v)
	}
	*j = CostKindEnumType_2(v)
	return nil
}

type CostKindEnumType_2 string

const ChargingRateUnitEnumType_7_A ChargingRateUnitEnumType_7 = "A"
const ChargingRateUnitEnumType_7_W ChargingRateUnitEnumType_7 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_7) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_7 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_7, v)
	}
	*j = ChargingRateUnitEnumType_7(v)
	return nil
}

type ChargingRateUnitEnumType_7 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType_2) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit: required")
	}
	if v, ok := raw["startPeriod"]; !ok || v == nil {
		return fmt.Errorf("field startPeriod: required")
	}
	type Plain ChargingSchedulePeriodType_2
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType_2(plain)
	return nil
}

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
//
type ChargingSchedulePeriodType_2 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_72) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_72
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_72(plain)
	return nil
}

const ChargingRateUnitEnumType_6_A ChargingRateUnitEnumType_6 = "A"
const ChargingRateUnitEnumType_6_W ChargingRateUnitEnumType_6 = "W"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_6) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_6 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_6, v)
	}
	*j = ChargingRateUnitEnumType_6(v)
	return nil
}

type ChargingRateUnitEnumType_6 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyEVChargingNeedsResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain NotifyEVChargingNeedsResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NotifyEVChargingNeedsResponseJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyEVChargingNeedsStatusEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NotifyEVChargingNeedsStatusEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NotifyEVChargingNeedsStatusEnumType_1, v)
	}
	*j = NotifyEVChargingNeedsStatusEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StatusInfoType_23) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["reasonCode"]; !ok || v == nil {
		return fmt.Errorf("field reasonCode: required")
	}
	type Plain StatusInfoType_23
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = StatusInfoType_23(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyEVChargingNeedsStatusEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NotifyEVChargingNeedsStatusEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NotifyEVChargingNeedsStatusEnumType, v)
	}
	*j = NotifyEVChargingNeedsStatusEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_71) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_71
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_71(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PhaseEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PhaseEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PhaseEnumType, v)
	}
	*j = PhaseEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyEVChargingNeedsRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingNeeds"]; !ok || v == nil {
		return fmt.Errorf("field chargingNeeds: required")
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	type Plain NotifyEVChargingNeedsRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NotifyEVChargingNeedsRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnergyTransferModeEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnergyTransferModeEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnergyTransferModeEnumType_1, v)
	}
	*j = EnergyTransferModeEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingNeedsType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["requestedEnergyTransfer"]; !ok || v == nil {
		return fmt.Errorf("field requestedEnergyTransfer: required")
	}
	type Plain ChargingNeedsType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingNeedsType(plain)
	return nil
}

// Charging_ Needs
// urn:x-oca:ocpp:uid:2:233249
//
type ChargingNeedsType struct {
	// AcChargingParameters corresponds to the JSON schema field
	// "acChargingParameters".
	AcChargingParameters *ACChargingParametersType `json:"acChargingParameters,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_70 `json:"customData,omitempty"`

	// DcChargingParameters corresponds to the JSON schema field
	// "dcChargingParameters".
	DcChargingParameters *DCChargingParametersType `json:"dcChargingParameters,omitempty"`

	// Charging_ Needs. Departure_ Time. Date_ Time
	// urn:x-oca:ocpp:uid:1:569223
	// Estimated departure time of the EV.
	//
	DepartureTime *string `json:"departureTime,omitempty"`

	// RequestedEnergyTransfer corresponds to the JSON schema field
	// "requestedEnergyTransfer".
	RequestedEnergyTransfer EnergyTransferModeEnumType `json:"requestedEnergyTransfer"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EnergyTransferModeEnumType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_EnergyTransferModeEnumType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_EnergyTransferModeEnumType, v)
	}
	*j = EnergyTransferModeEnumType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DCChargingParametersType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["evMaxCurrent"]; !ok || v == nil {
		return fmt.Errorf("field evMaxCurrent: required")
	}
	if v, ok := raw["evMaxVoltage"]; !ok || v == nil {
		return fmt.Errorf("field evMaxVoltage: required")
	}
	type Plain DCChargingParametersType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DCChargingParametersType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ACChargingParametersType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["energyAmount"]; !ok || v == nil {
		return fmt.Errorf("field energyAmount: required")
	}
	if v, ok := raw["evMaxCurrent"]; !ok || v == nil {
		return fmt.Errorf("field evMaxCurrent: required")
	}
	if v, ok := raw["evMaxVoltage"]; !ok || v == nil {
		return fmt.Errorf("field evMaxVoltage: required")
	}
	if v, ok := raw["evMinCurrent"]; !ok || v == nil {
		return fmt.Errorf("field evMinCurrent: required")
	}
	type Plain ACChargingParametersType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ACChargingParametersType(plain)
	return nil
}

// AC_ Charging_ Parameters
// urn:x-oca:ocpp:uid:2:233250
// EV AC charging parameters.
//
//
type ACChargingParametersType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_70 `json:"customData,omitempty"`

	// AC_ Charging_ Parameters. Energy_ Amount. Energy_ Amount
	// urn:x-oca:ocpp:uid:1:569211
	// Amount of energy requested (in Wh). This includes energy required for
	// preconditioning.
	//
	EnergyAmount int `json:"energyAmount"`

	// AC_ Charging_ Parameters. EV_ Max. Current
	// urn:x-oca:ocpp:uid:1:569213
	// Maximum current (amps) supported by the electric vehicle (per phase). Includes
	// cable capacity.
	//
	EvMaxCurrent int `json:"evMaxCurrent"`

	// AC_ Charging_ Parameters. EV_ Max. Voltage
	// urn:x-oca:ocpp:uid:1:569214
	// Maximum voltage supported by the electric vehicle
	//
	EvMaxVoltage int `json:"evMaxVoltage"`

	// AC_ Charging_ Parameters. EV_ Min. Current
	// urn:x-oca:ocpp:uid:1:569212
	// Minimum current (amps) supported by the electric vehicle (per phase).
	//
	EvMinCurrent int `json:"evMinCurrent"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_70) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_70
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_70(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_69) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_69
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_69(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyDisplayMessagesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	type Plain NotifyDisplayMessagesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = NotifyDisplayMessagesRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SignedMeterValueType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["encodingMethod"]; !ok || v == nil {
		return fmt.Errorf("field encodingMethod: required")
	}
	if v, ok := raw["publicKey"]; !ok || v == nil {
		return fmt.Errorf("field publicKey: required")
	}
	if v, ok := raw["signedMeterData"]; !ok || v == nil {
		return fmt.Errorf("field signedMeterData: required")
	}
	if v, ok := raw["signingMethod"]; !ok || v == nil {
		return fmt.Errorf("field signingMethod: required")
	}
	type Plain SignedMeterValueType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SignedMeterValueType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType_3, v)
	}
	*j = MessageStateEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UnitOfMeasureType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain UnitOfMeasureType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["multiplier"]; !ok || v == nil {
		plain.Multiplier = 0
	}
	if v, ok := raw["unit"]; !ok || v == nil {
		plain.Unit = "Wh"
	}
	*j = UnitOfMeasureType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType_3, v)
	}
	*j = MessagePriorityEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SampledValueType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["value"]; !ok || v == nil {
		return fmt.Errorf("field value: required")
	}
	type Plain SampledValueType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SampledValueType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageInfoType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["message"]; !ok || v == nil {
		return fmt.Errorf("field message: required")
	}
	if v, ok := raw["priority"]; !ok || v == nil {
		return fmt.Errorf("field priority: required")
	}
	type Plain MessageInfoType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageInfoType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeterValueType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["sampledValue"]; !ok || v == nil {
		return fmt.Errorf("field sampledValue: required")
	}
	if v, ok := raw["timestamp"]; !ok || v == nil {
		return fmt.Errorf("field timestamp: required")
	}
	type Plain MeterValueType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MeterValueType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageStateEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageStateEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageStateEnumType_2, v)
	}
	*j = MessageStateEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessagePriorityEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessagePriorityEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessagePriorityEnumType_2, v)
	}
	*j = MessagePriorityEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *PhaseEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_PhaseEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_PhaseEnumType_1, v)
	}
	*j = PhaseEnumType_1(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_3) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_3 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_3, v)
	}
	*j = MessageFormatEnumType_3(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageContentType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["content"]; !ok || v == nil {
		return fmt.Errorf("field content: required")
	}
	if v, ok := raw["format"]; !ok || v == nil {
		return fmt.Errorf("field format: required")
	}
	type Plain MessageContentType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MessageContentType_1(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MessageFormatEnumType_2) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_MessageFormatEnumType_2 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_MessageFormatEnumType_2, v)
	}
	*j = MessageFormatEnumType_2(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ComponentType_4) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain ComponentType_4
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ComponentType_4(plain)
	return nil
}

// A physical or logical component
//
type ComponentType_4 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_68 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_5 `json:"evse,omitempty"`

	// Name of instance in case the component exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the component. Name should be taken from the list of standardized
	// component names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *EVSEType_5) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain EVSEType_5
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = EVSEType_5(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_68) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_68
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_68(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_67) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_67
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_67(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyCustomerInformationRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["data"]; !ok || v == nil {
		return fmt.Errorf("field data: required")
	}
	if v, ok := raw["generatedAt"]; !ok || v == nil {
		return fmt.Errorf("field generatedAt: required")
	}
	if v, ok := raw["requestId"]; !ok || v == nil {
		return fmt.Errorf("field requestId: required")
	}
	if v, ok := raw["seqNo"]; !ok || v == nil {
		return fmt.Errorf("field seqNo: required")
	}
	type Plain NotifyCustomerInformationRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["tbc"]; !ok || v == nil {
		plain.Tbc = false
	}
	*j = NotifyCustomerInformationRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_66) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_66
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_66(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_65) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_65
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_65(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NotifyChargingLimitRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingLimit"]; !ok || v == nil {
		return fmt.Errorf("field chargingLimit: required")
	}
	type Plain NotifyChargingLimitRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NotifyChargingLimitRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ReadingContextEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ReadingContextEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ReadingContextEnumType_1, v)
	}
	*j = ReadingContextEnumType_1(v)
	return nil
}

const CostKindEnumType_1_RenewableGenerationPercentage CostKindEnumType_1 = "RenewableGenerationPercentage"
const CostKindEnumType_1_RelativePricePercentage CostKindEnumType_1 = "RelativePricePercentage"
const CostKindEnumType_1_CarbonDioxideEmission CostKindEnumType_1 = "CarbonDioxideEmission"

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostKindEnumType_1) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_CostKindEnumType_1 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_CostKindEnumType_1, v)
	}
	*j = CostKindEnumType_1(v)
	return nil
}

type CostKindEnumType_1 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingScheduleType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingRateUnit"]; !ok || v == nil {
		return fmt.Errorf("field chargingRateUnit: required")
	}
	if v, ok := raw["chargingSchedulePeriod"]; !ok || v == nil {
		return fmt.Errorf("field chargingSchedulePeriod: required")
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	type Plain ChargingScheduleType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingScheduleType(plain)
	return nil
}

// Charging_ Schedule
// urn:x-oca:ocpp:uid:2:233256
// Charging schedule structure defines a list of charging periods, as used in:
// GetCompositeSchedule.conf and ChargingProfile.
//
type ChargingScheduleType struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit ChargingRateUnitEnumType_5 `json:"chargingRateUnit"`

	// ChargingSchedulePeriod corresponds to the JSON schema field
	// "chargingSchedulePeriod".
	ChargingSchedulePeriod []ChargingSchedulePeriodType_1 `json:"chargingSchedulePeriod"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// Charging_ Schedule. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569236
	// Duration of the charging schedule in seconds. If the duration is left empty,
	// the last period will continue indefinitely or until end of the transaction if
	// chargingProfilePurpose = TxProfile.
	//
	Duration *int `json:"duration,omitempty"`

	// Identifies the ChargingSchedule.
	//
	Id int `json:"id"`

	// Charging_ Schedule. Min_ Charging_ Rate. Numeric
	// urn:x-oca:ocpp:uid:1:569239
	// Minimum charging rate supported by the EV. The unit of measure is defined by
	// the chargingRateUnit. This parameter is intended to be used by a local smart
	// charging algorithm to optimize the power allocation for in the case a charging
	// process is inefficient at lower charging rates. Accepts at most one digit
	// fraction (e.g. 8.1)
	//
	MinChargingRate *float64 `json:"minChargingRate,omitempty"`

	// SalesTariff corresponds to the JSON schema field "salesTariff".
	SalesTariff *SalesTariffType `json:"salesTariff,omitempty"`

	// Charging_ Schedule. Start_ Schedule. Date_ Time
	// urn:x-oca:ocpp:uid:1:569237
	// Starting point of an absolute schedule. If absent the schedule will be relative
	// to start of charging.
	//
	StartSchedule *string `json:"startSchedule,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["id"]; !ok || v == nil {
		return fmt.Errorf("field id: required")
	}
	if v, ok := raw["salesTariffEntry"]; !ok || v == nil {
		return fmt.Errorf("field salesTariffEntry: required")
	}
	type Plain SalesTariffType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *SalesTariffEntryType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["relativeTimeInterval"]; !ok || v == nil {
		return fmt.Errorf("field relativeTimeInterval: required")
	}
	type Plain SalesTariffEntryType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = SalesTariffEntryType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *MeterValuesRequestJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["evseId"]; !ok || v == nil {
		return fmt.Errorf("field evseId: required")
	}
	if v, ok := raw["meterValue"]; !ok || v == nil {
		return fmt.Errorf("field meterValue: required")
	}
	type Plain MeterValuesRequestJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = MeterValuesRequestJson(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *RelativeTimeIntervalType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["start"]; !ok || v == nil {
		return fmt.Errorf("field start: required")
	}
	type Plain RelativeTimeIntervalType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = RelativeTimeIntervalType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_63) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_63
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_63(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ConsumptionCostType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["cost"]; !ok || v == nil {
		return fmt.Errorf("field cost: required")
	}
	if v, ok := raw["startValue"]; !ok || v == nil {
		return fmt.Errorf("field startValue: required")
	}
	type Plain ConsumptionCostType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ConsumptionCostType(plain)
	return nil
}

type ChargingLimitSourceEnumType_4 string

// Consumption_ Cost
// urn:x-oca:ocpp:uid:2:233259
//
type ConsumptionCostType struct {
	// Cost corresponds to the JSON schema field "cost".
	Cost []CostType `json:"cost"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// Consumption_ Cost. Start_ Value. Numeric
	// urn:x-oca:ocpp:uid:1:569246
	// The lowest level of consumption that defines the starting point of this
	// consumption block. The block interval extends to the start of the next
	// interval.
	//
	StartValue float64 `json:"startValue"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType_4, v)
	}
	*j = ChargingLimitSourceEnumType_4(v)
	return nil
}

const ChargingLimitSourceEnumType_4_EMS ChargingLimitSourceEnumType_4 = "EMS"
const ChargingLimitSourceEnumType_4_Other ChargingLimitSourceEnumType_4 = "Other"
const ChargingLimitSourceEnumType_4_SO ChargingLimitSourceEnumType_4 = "SO"
const ChargingLimitSourceEnumType_4_CSO ChargingLimitSourceEnumType_4 = "CSO"

type ChargingLimitSourceEnumType_5 string

// UnmarshalJSON implements json.Unmarshaler.
func (j *CostType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["amount"]; !ok || v == nil {
		return fmt.Errorf("field amount: required")
	}
	if v, ok := raw["costKind"]; !ok || v == nil {
		return fmt.Errorf("field costKind: required")
	}
	type Plain CostType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CostType(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitSourceEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingLimitSourceEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingLimitSourceEnumType_5, v)
	}
	*j = ChargingLimitSourceEnumType_5(v)
	return nil
}

const ChargingLimitSourceEnumType_5_EMS ChargingLimitSourceEnumType_5 = "EMS"
const ChargingLimitSourceEnumType_5_Other ChargingLimitSourceEnumType_5 = "Other"
const ChargingLimitSourceEnumType_5_SO ChargingLimitSourceEnumType_5 = "SO"
const ChargingLimitSourceEnumType_5_CSO ChargingLimitSourceEnumType_5 = "CSO"

// Cost
// urn:x-oca:ocpp:uid:2:233258
//
type CostType struct {
	// Cost. Amount. Amount
	// urn:x-oca:ocpp:uid:1:569244
	// The estimated or actual cost per kWh
	//
	Amount int `json:"amount"`

	// Cost. Amount_ Multiplier. Integer
	// urn:x-oca:ocpp:uid:1:569245
	// Values: -3..3, The amountMultiplier defines the exponent to base 10 (dec). The
	// final value is determined by: amount * 10 ^ amountMultiplier
	//
	AmountMultiplier *int `json:"amountMultiplier,omitempty"`

	// CostKind corresponds to the JSON schema field "costKind".
	CostKind CostKindEnumType `json:"costKind"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *CustomDataType_64) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["vendorId"]; !ok || v == nil {
		return fmt.Errorf("field vendorId: required")
	}
	type Plain CustomDataType_64
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = CustomDataType_64(plain)
	return nil
}

// Charging_ Limit
// urn:x-enexis:ecdm:uid:2:234489
//
type ChargingLimitType struct {
	// ChargingLimitSource corresponds to the JSON schema field "chargingLimitSource".
	ChargingLimitSource ChargingLimitSourceEnumType_5 `json:"chargingLimitSource"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// Charging_ Limit. Is_ Grid_ Critical. Indicator
	// urn:x-enexis:ecdm:uid:1:570847
	// Indicates whether the charging limit is critical for the grid.
	//
	IsGridCritical *bool `json:"isGridCritical,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingLimitType) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["chargingLimitSource"]; !ok || v == nil {
		return fmt.Errorf("field chargingLimitSource: required")
	}
	type Plain ChargingLimitType
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingLimitType(plain)
	return nil
}

type ChargingRateUnitEnumType_4 string

const CostKindEnumTypeRenewableGenerationPercentage CostKindEnumType = "RenewableGenerationPercentage"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_4) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_4 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_4, v)
	}
	*j = ChargingRateUnitEnumType_4(v)
	return nil
}

const ChargingRateUnitEnumType_4_W ChargingRateUnitEnumType_4 = "W"
const ChargingRateUnitEnumType_4_A ChargingRateUnitEnumType_4 = "A"

// Charging_ Schedule_ Period
// urn:x-oca:ocpp:uid:2:233257
// Charging schedule period structure defines a time period in a charging schedule.
//
type ChargingSchedulePeriodType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// Charging_ Schedule_ Period. Limit. Measure
	// urn:x-oca:ocpp:uid:1:569241
	// Charging rate limit during the schedule period, in the applicable
	// chargingRateUnit, for example in Amperes (A) or Watts (W). Accepts at most one
	// digit fraction (e.g. 8.1).
	//
	Limit float64 `json:"limit"`

	// Charging_ Schedule_ Period. Number_ Phases. Counter
	// urn:x-oca:ocpp:uid:1:569242
	// The number of phases that can be used for charging. If a number of phases is
	// needed, numberPhases=3 will be assumed unless another number is given.
	//
	NumberPhases *int `json:"numberPhases,omitempty"`

	// Values: 1..3, Used if numberPhases=1 and if the EVSE is capable of switching
	// the phase connected to the EV, i.e. ACPhaseSwitchingSupported is defined and
	// true. It’s not allowed unless both conditions above are true. If both
	// conditions are true, and phaseToUse is omitted, the Charging Station / EVSE
	// will make the selection on its own.
	//
	//
	PhaseToUse *int `json:"phaseToUse,omitempty"`

	// Charging_ Schedule_ Period. Start_ Period. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569240
	// Start of the period, in seconds from the start of schedule. The value of
	// StartPeriod also defines the stop time of the previous period.
	//
	StartPeriod int `json:"startPeriod"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingSchedulePeriodType_1) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["limit"]; !ok || v == nil {
		return fmt.Errorf("field limit: required")
	}
	if v, ok := raw["startPeriod"]; !ok || v == nil {
		return fmt.Errorf("field startPeriod: required")
	}
	type Plain ChargingSchedulePeriodType_1
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = ChargingSchedulePeriodType_1(plain)
	return nil
}

type ChargingRateUnitEnumType_5 string

const CostKindEnumTypeRelativePricePercentage CostKindEnumType = "RelativePricePercentage"

// UnmarshalJSON implements json.Unmarshaler.
func (j *ChargingRateUnitEnumType_5) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_ChargingRateUnitEnumType_5 {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_ChargingRateUnitEnumType_5, v)
	}
	*j = ChargingRateUnitEnumType_5(v)
	return nil
}

const ChargingRateUnitEnumType_5_W ChargingRateUnitEnumType_5 = "W"
const ChargingRateUnitEnumType_5_A ChargingRateUnitEnumType_5 = "A"

type CostKindEnumType string

const ChargingProfileKindEnumType_2_Recurring ChargingProfileKindEnumType_2 = "Recurring"

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_1 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_10 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_100 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_101 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_102 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_103 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_104 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_105 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_106 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_107 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_108 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_109 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_11 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_110 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_111 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_112 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_113 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_114 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_115 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_116 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_117 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_118 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_119 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_12 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_120 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_121 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_122 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_123 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_124 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_125 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_126 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_127 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_13 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_14 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_15 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_16 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_17 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_18 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_19 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_2 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`

	// ConnectorType type of connector.
	ConnectorType []string `json:"connectorType"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_20 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_21 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_22 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_23 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_24 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_25 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_26 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_27 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_28 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_29 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_3 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_30 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_31 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_32 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_33 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_34 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_35 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_36 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_37 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_38 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_39 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_4 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_40 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_41 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_42 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_43 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_44 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_45 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_46 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_47 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_48 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_49 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_5 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_50 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_51 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_52 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_53 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_54 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_55 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_56 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_57 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_58 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_59 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_6 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_60 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_61 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_62 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_63 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_64 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_65 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_66 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_67 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_68 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_69 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_7 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_70 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_71 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_72 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_73 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_74 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_75 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_76 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_77 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_78 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_79 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_8 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_80 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_81 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_82 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_83 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_84 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_85 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_86 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_87 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_88 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_89 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_9 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_90 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_91 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_92 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_93 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_94 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_95 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_96 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_97 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_98 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

// This class does not get 'AdditionalProperties = false' in the schema generation,
// so it can be extended with arbitrary JSON properties to allow adding custom
// data.
type CustomDataType_99 struct {
	// VendorId corresponds to the JSON schema field "vendorId".
	VendorId string `json:"vendorId"`
}

type CustomerInformationRequestJson struct {
	// Flag indicating whether the Charging Station should clear all information about
	// the customer referred to.
	//
	Clear bool `json:"clear"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_22 `json:"customData,omitempty"`

	// CustomerCertificate corresponds to the JSON schema field "customerCertificate".
	CustomerCertificate *CertificateHashDataType `json:"customerCertificate,omitempty"`

	// A (e.g. vendor specific) identifier of the customer this request refers to.
	// This field contains a custom identifier other than IdToken and Certificate.
	// One of the possible identifiers (customerIdentifier, customerIdToken or
	// customerCertificate) should be in the request message.
	//
	CustomerIdentifier *string `json:"customerIdentifier,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken *IdTokenType_2 `json:"idToken,omitempty"`

	// Flag indicating whether the Charging Station should return
	// NotifyCustomerInformationRequest messages containing information about the
	// customer referred to.
	//
	Report bool `json:"report"`

	// The Id of the request.
	//
	//
	RequestId int `json:"requestId"`
}

type CustomerInformationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_23 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status CustomerInformationStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_8 `json:"statusInfo,omitempty"`
}

type CustomerInformationStatusEnumType string

const CustomerInformationStatusEnumTypeAccepted CustomerInformationStatusEnumType = "Accepted"
const CustomerInformationStatusEnumTypeInvalid CustomerInformationStatusEnumType = "Invalid"
const CustomerInformationStatusEnumTypeRejected CustomerInformationStatusEnumType = "Rejected"

type CustomerInformationStatusEnumType_1 string

const CustomerInformationStatusEnumType_1_Accepted CustomerInformationStatusEnumType_1 = "Accepted"
const CustomerInformationStatusEnumType_1_Invalid CustomerInformationStatusEnumType_1 = "Invalid"
const CustomerInformationStatusEnumType_1_Rejected CustomerInformationStatusEnumType_1 = "Rejected"

// DC_ Charging_ Parameters
// urn:x-oca:ocpp:uid:2:233251
// EV DC charging parameters
//
//
//
type DCChargingParametersType struct {
	// DC_ Charging_ Parameters. Bulk_ SOC. Percentage
	// urn:x-oca:ocpp:uid:1:569222
	// Percentage of SoC at which the EV considers a fast charging process to end.
	// (possible values: 0 - 100)
	//
	BulkSoC *int `json:"bulkSoC,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_70 `json:"customData,omitempty"`

	// DC_ Charging_ Parameters. Energy_ Amount. Energy_ Amount
	// urn:x-oca:ocpp:uid:1:569217
	// Amount of energy requested (in Wh). This inludes energy required for
	// preconditioning.
	//
	EnergyAmount *int `json:"energyAmount,omitempty"`

	// DC_ Charging_ Parameters. EV_ Energy_ Capacity. Numeric
	// urn:x-oca:ocpp:uid:1:569220
	// Capacity of the electric vehicle battery (in Wh)
	//
	EvEnergyCapacity *int `json:"evEnergyCapacity,omitempty"`

	// DC_ Charging_ Parameters. EV_ Max. Current
	// urn:x-oca:ocpp:uid:1:569215
	// Maximum current (amps) supported by the electric vehicle. Includes cable
	// capacity.
	//
	EvMaxCurrent int `json:"evMaxCurrent"`

	// DC_ Charging_ Parameters. EV_ Max. Power
	// urn:x-oca:ocpp:uid:1:569218
	// Maximum power (in W) supported by the electric vehicle. Required for DC
	// charging.
	//
	EvMaxPower *int `json:"evMaxPower,omitempty"`

	// DC_ Charging_ Parameters. EV_ Max. Voltage
	// urn:x-oca:ocpp:uid:1:569216
	// Maximum voltage supported by the electric vehicle
	//
	EvMaxVoltage int `json:"evMaxVoltage"`

	// DC_ Charging_ Parameters. Full_ SOC. Percentage
	// urn:x-oca:ocpp:uid:1:569221
	// Percentage of SoC at which the EV considers the battery fully charged.
	// (possible values: 0 - 100)
	//
	FullSoC *int `json:"fullSoC,omitempty"`

	// DC_ Charging_ Parameters. State_ Of_ Charge. Numeric
	// urn:x-oca:ocpp:uid:1:569219
	// Energy available in the battery (in percent of the battery capacity)
	//
	StateOfCharge *int `json:"stateOfCharge,omitempty"`
}

type DataEnumType string

const DataEnumTypeBoolean DataEnumType = "boolean"
const DataEnumTypeDateTime DataEnumType = "dateTime"
const DataEnumTypeDecimal DataEnumType = "decimal"
const DataEnumTypeInteger DataEnumType = "integer"
const DataEnumTypeMemberList DataEnumType = "MemberList"
const DataEnumTypeOptionList DataEnumType = "OptionList"
const DataEnumTypeSequenceList DataEnumType = "SequenceList"
const DataEnumTypeString DataEnumType = "string"

type DataEnumType_1 string

const DataEnumType_1_Boolean DataEnumType_1 = "boolean"
const DataEnumType_1_DateTime DataEnumType_1 = "dateTime"
const DataEnumType_1_Decimal DataEnumType_1 = "decimal"
const DataEnumType_1_Integer DataEnumType_1 = "integer"
const DataEnumType_1_MemberList DataEnumType_1 = "MemberList"
const DataEnumType_1_OptionList DataEnumType_1 = "OptionList"
const DataEnumType_1_SequenceList DataEnumType_1 = "SequenceList"
const DataEnumType_1_String DataEnumType_1 = "string"

type DataTransferRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_24 `json:"customData,omitempty"`

	// Data without specified length or format. This needs to be decided by both
	// parties (Open to implementation).
	//
	Data interface{} `json:"data,omitempty"`

	// May be used to indicate a specific message or implementation.
	//
	MessageId *string `json:"messageId,omitempty"`

	// This identifies the Vendor specific implementation
	//
	//
	VendorId string `json:"vendorId"`
}

type DataTransferResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_25 `json:"customData,omitempty"`

	// Data without specified length or format, in response to request.
	//
	Data interface{} `json:"data,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status DataTransferStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_9 `json:"statusInfo,omitempty"`
}

type DataTransferStatusEnumType string

const DataTransferStatusEnumTypeAccepted DataTransferStatusEnumType = "Accepted"
const DataTransferStatusEnumTypeRejected DataTransferStatusEnumType = "Rejected"
const DataTransferStatusEnumTypeUnknownMessageId DataTransferStatusEnumType = "UnknownMessageId"
const DataTransferStatusEnumTypeUnknownVendorId DataTransferStatusEnumType = "UnknownVendorId"

type DataTransferStatusEnumType_1 string

const DataTransferStatusEnumType_1_Accepted DataTransferStatusEnumType_1 = "Accepted"
const DataTransferStatusEnumType_1_Rejected DataTransferStatusEnumType_1 = "Rejected"
const DataTransferStatusEnumType_1_UnknownMessageId DataTransferStatusEnumType_1 = "UnknownMessageId"
const DataTransferStatusEnumType_1_UnknownVendorId DataTransferStatusEnumType_1 = "UnknownVendorId"

type DeleteCertificateRequestJson struct {
	// CertificateHashData corresponds to the JSON schema field "certificateHashData".
	CertificateHashData CertificateHashDataType_1 `json:"certificateHashData"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_26 `json:"customData,omitempty"`
}

type DeleteCertificateResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_27 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status DeleteCertificateStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_10 `json:"statusInfo,omitempty"`
}

type DeleteCertificateStatusEnumType string

const DeleteCertificateStatusEnumTypeAccepted DeleteCertificateStatusEnumType = "Accepted"
const DeleteCertificateStatusEnumTypeFailed DeleteCertificateStatusEnumType = "Failed"
const DeleteCertificateStatusEnumTypeNotFound DeleteCertificateStatusEnumType = "NotFound"

type DeleteCertificateStatusEnumType_1 string

const DeleteCertificateStatusEnumType_1_Accepted DeleteCertificateStatusEnumType_1 = "Accepted"
const DeleteCertificateStatusEnumType_1_Failed DeleteCertificateStatusEnumType_1 = "Failed"
const DeleteCertificateStatusEnumType_1_NotFound DeleteCertificateStatusEnumType_1 = "NotFound"

type DisplayMessageStatusEnumType string

const DisplayMessageStatusEnumTypeAccepted DisplayMessageStatusEnumType = "Accepted"
const DisplayMessageStatusEnumTypeNotSupportedMessageFormat DisplayMessageStatusEnumType = "NotSupportedMessageFormat"
const DisplayMessageStatusEnumTypeNotSupportedPriority DisplayMessageStatusEnumType = "NotSupportedPriority"
const DisplayMessageStatusEnumTypeNotSupportedState DisplayMessageStatusEnumType = "NotSupportedState"
const DisplayMessageStatusEnumTypeRejected DisplayMessageStatusEnumType = "Rejected"
const DisplayMessageStatusEnumTypeUnknownTransaction DisplayMessageStatusEnumType = "UnknownTransaction"

type DisplayMessageStatusEnumType_1 string

const DisplayMessageStatusEnumType_1_Accepted DisplayMessageStatusEnumType_1 = "Accepted"
const DisplayMessageStatusEnumType_1_NotSupportedMessageFormat DisplayMessageStatusEnumType_1 = "NotSupportedMessageFormat"
const DisplayMessageStatusEnumType_1_NotSupportedPriority DisplayMessageStatusEnumType_1 = "NotSupportedPriority"
const DisplayMessageStatusEnumType_1_NotSupportedState DisplayMessageStatusEnumType_1 = "NotSupportedState"
const DisplayMessageStatusEnumType_1_Rejected DisplayMessageStatusEnumType_1 = "Rejected"
const DisplayMessageStatusEnumType_1_UnknownTransaction DisplayMessageStatusEnumType_1 = "UnknownTransaction"

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_8 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_1 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_48 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_10 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_110 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_11 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_111 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_12 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_112 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_13 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_113 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_14 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_15 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_120 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_2 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_50 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_3 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_54 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_4 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_55 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_5 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_68 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_6 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_74 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_7 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_76 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_8 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_78 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

// EVSE
// urn:x-oca:ocpp:uid:2:233123
// Electric Vehicle Supply Equipment
//
type EVSEType_9 struct {
	// An id to designate a specific connector (on an EVSE) by connector index number.
	//
	ConnectorId *int `json:"connectorId,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_102 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// EVSE Identifier. This contains a number (&gt; 0) designating an EVSE of the
	// Charging Station.
	//
	Id int `json:"id"`
}

type EnergyTransferModeEnumType string

const EnergyTransferModeEnumTypeACSinglePhase EnergyTransferModeEnumType = "AC_single_phase"
const EnergyTransferModeEnumTypeACThreePhase EnergyTransferModeEnumType = "AC_three_phase"
const EnergyTransferModeEnumTypeACTwoPhase EnergyTransferModeEnumType = "AC_two_phase"
const EnergyTransferModeEnumTypeDC EnergyTransferModeEnumType = "DC"

type EnergyTransferModeEnumType_1 string

const EnergyTransferModeEnumType_1_ACSinglePhase EnergyTransferModeEnumType_1 = "AC_single_phase"
const EnergyTransferModeEnumType_1_ACThreePhase EnergyTransferModeEnumType_1 = "AC_three_phase"
const EnergyTransferModeEnumType_1_ACTwoPhase EnergyTransferModeEnumType_1 = "AC_two_phase"
const EnergyTransferModeEnumType_1_DC EnergyTransferModeEnumType_1 = "DC"

// Class to report an event notification for a component-variable.
//
type EventDataType struct {
	// Actual value (_attributeType_ Actual) of the variable.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-reporting-value-size,ReportingValueSize&gt;&gt; can be used
	// to limit GetVariableResult.attributeValue, VariableAttribute.value and
	// EventData.actualValue. The max size of these values will always remain equal.
	//
	//
	ActualValue string `json:"actualValue"`

	// Refers to the Id of an event that is considered to be the cause for this event.
	//
	//
	Cause *int `json:"cause,omitempty"`

	// _Cleared_ is set to true to report the clearing of a monitored situation, i.e.
	// a 'return to normal'.
	//
	//
	Cleared *bool `json:"cleared,omitempty"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType_5 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_74 `json:"customData,omitempty"`

	// Identifies the event. This field can be referred to as a cause by other events.
	//
	//
	EventId int `json:"eventId"`

	// EventNotificationType corresponds to the JSON schema field
	// "eventNotificationType".
	EventNotificationType EventNotificationEnumType `json:"eventNotificationType"`

	// Technical (error) code as reported by component.
	//
	TechCode *string `json:"techCode,omitempty"`

	// Technical detail information as reported by component.
	//
	TechInfo *string `json:"techInfo,omitempty"`

	// Timestamp of the moment the report was generated.
	//
	Timestamp string `json:"timestamp"`

	// If an event notification is linked to a specific transaction, this field can be
	// used to specify its transactionId.
	//
	TransactionId *string `json:"transactionId,omitempty"`

	// Trigger corresponds to the JSON schema field "trigger".
	Trigger EventTriggerEnumType `json:"trigger"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_4 `json:"variable"`

	// Identifies the VariableMonitoring which triggered the event.
	//
	VariableMonitoringId *int `json:"variableMonitoringId,omitempty"`
}

type EventNotificationEnumType string

const EventNotificationEnumTypeCustomMonitor EventNotificationEnumType = "CustomMonitor"
const EventNotificationEnumTypeHardWiredMonitor EventNotificationEnumType = "HardWiredMonitor"
const EventNotificationEnumTypeHardWiredNotification EventNotificationEnumType = "HardWiredNotification"
const EventNotificationEnumTypePreconfiguredMonitor EventNotificationEnumType = "PreconfiguredMonitor"

type EventNotificationEnumType_1 string

const EventNotificationEnumType_1_CustomMonitor EventNotificationEnumType_1 = "CustomMonitor"
const EventNotificationEnumType_1_HardWiredMonitor EventNotificationEnumType_1 = "HardWiredMonitor"
const EventNotificationEnumType_1_HardWiredNotification EventNotificationEnumType_1 = "HardWiredNotification"
const EventNotificationEnumType_1_PreconfiguredMonitor EventNotificationEnumType_1 = "PreconfiguredMonitor"

type EventTriggerEnumType string

const EventTriggerEnumTypeAlerting EventTriggerEnumType = "Alerting"
const EventTriggerEnumTypeDelta EventTriggerEnumType = "Delta"
const EventTriggerEnumTypePeriodic EventTriggerEnumType = "Periodic"

type EventTriggerEnumType_1 string

const EventTriggerEnumType_1_Alerting EventTriggerEnumType_1 = "Alerting"
const EventTriggerEnumType_1_Delta EventTriggerEnumType_1 = "Delta"
const EventTriggerEnumType_1_Periodic EventTriggerEnumType_1 = "Periodic"

type FirmwareStatusEnumType string

const FirmwareStatusEnumTypeDownloadFailed FirmwareStatusEnumType = "DownloadFailed"
const FirmwareStatusEnumTypeDownloadPaused FirmwareStatusEnumType = "DownloadPaused"
const FirmwareStatusEnumTypeDownloadScheduled FirmwareStatusEnumType = "DownloadScheduled"
const FirmwareStatusEnumTypeDownloaded FirmwareStatusEnumType = "Downloaded"
const FirmwareStatusEnumTypeDownloading FirmwareStatusEnumType = "Downloading"
const FirmwareStatusEnumTypeIdle FirmwareStatusEnumType = "Idle"
const FirmwareStatusEnumTypeInstallRebooting FirmwareStatusEnumType = "InstallRebooting"
const FirmwareStatusEnumTypeInstallScheduled FirmwareStatusEnumType = "InstallScheduled"
const FirmwareStatusEnumTypeInstallVerificationFailed FirmwareStatusEnumType = "InstallVerificationFailed"
const FirmwareStatusEnumTypeInstallationFailed FirmwareStatusEnumType = "InstallationFailed"
const FirmwareStatusEnumTypeInstalled FirmwareStatusEnumType = "Installed"
const FirmwareStatusEnumTypeInstalling FirmwareStatusEnumType = "Installing"
const FirmwareStatusEnumTypeInvalidSignature FirmwareStatusEnumType = "InvalidSignature"
const FirmwareStatusEnumTypeSignatureVerified FirmwareStatusEnumType = "SignatureVerified"

type FirmwareStatusEnumType_1 string

const FirmwareStatusEnumType_1_DownloadFailed FirmwareStatusEnumType_1 = "DownloadFailed"
const FirmwareStatusEnumType_1_DownloadPaused FirmwareStatusEnumType_1 = "DownloadPaused"
const FirmwareStatusEnumType_1_DownloadScheduled FirmwareStatusEnumType_1 = "DownloadScheduled"
const FirmwareStatusEnumType_1_Downloaded FirmwareStatusEnumType_1 = "Downloaded"
const FirmwareStatusEnumType_1_Downloading FirmwareStatusEnumType_1 = "Downloading"
const FirmwareStatusEnumType_1_Idle FirmwareStatusEnumType_1 = "Idle"
const FirmwareStatusEnumType_1_InstallRebooting FirmwareStatusEnumType_1 = "InstallRebooting"
const FirmwareStatusEnumType_1_InstallScheduled FirmwareStatusEnumType_1 = "InstallScheduled"
const FirmwareStatusEnumType_1_InstallVerificationFailed FirmwareStatusEnumType_1 = "InstallVerificationFailed"
const FirmwareStatusEnumType_1_InstallationFailed FirmwareStatusEnumType_1 = "InstallationFailed"
const FirmwareStatusEnumType_1_Installed FirmwareStatusEnumType_1 = "Installed"
const FirmwareStatusEnumType_1_Installing FirmwareStatusEnumType_1 = "Installing"
const FirmwareStatusEnumType_1_InvalidSignature FirmwareStatusEnumType_1 = "InvalidSignature"
const FirmwareStatusEnumType_1_SignatureVerified FirmwareStatusEnumType_1 = "SignatureVerified"

type FirmwareStatusNotificationRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_28 `json:"customData,omitempty"`

	// The request id that was provided in the
	// UpdateFirmwareRequest that started this firmware update.
	// This field is mandatory, unless the message was triggered by a
	// TriggerMessageRequest AND there is no firmware update ongoing.
	//
	RequestId *int `json:"requestId,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status FirmwareStatusEnumType_1 `json:"status"`
}

type FirmwareStatusNotificationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_29 `json:"customData,omitempty"`
}

// Firmware
// urn:x-enexis:ecdm:uid:2:233291
// Represents a copy of the firmware that can be loaded/updated on the Charging
// Station.
//
type FirmwareType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_126 `json:"customData,omitempty"`

	// Firmware. Install. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569462
	// Date and time at which the firmware shall be installed.
	//
	InstallDateTime *string `json:"installDateTime,omitempty"`

	// Firmware. Location. URI
	// urn:x-enexis:ecdm:uid:1:569460
	// URI defining the origin of the firmware.
	//
	Location string `json:"location"`

	// Firmware. Retrieve. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569461
	// Date and time at which the firmware shall be retrieved.
	//
	RetrieveDateTime string `json:"retrieveDateTime"`

	// Firmware. Signature. Signature
	// urn:x-enexis:ecdm:uid:1:569464
	// Base64 encoded firmware signature.
	//
	Signature *string `json:"signature,omitempty"`

	// Certificate with which the firmware was signed.
	// PEM encoded X.509 certificate.
	//
	SigningCertificate *string `json:"signingCertificate,omitempty"`
}

type GenericDeviceModelStatusEnumType string

const GenericDeviceModelStatusEnumTypeAccepted GenericDeviceModelStatusEnumType = "Accepted"
const GenericDeviceModelStatusEnumTypeEmptyResultSet GenericDeviceModelStatusEnumType = "EmptyResultSet"
const GenericDeviceModelStatusEnumTypeNotSupported GenericDeviceModelStatusEnumType = "NotSupported"
const GenericDeviceModelStatusEnumTypeRejected GenericDeviceModelStatusEnumType = "Rejected"

type GenericDeviceModelStatusEnumType_1 string

const GenericDeviceModelStatusEnumType_1_Accepted GenericDeviceModelStatusEnumType_1 = "Accepted"
const GenericDeviceModelStatusEnumType_1_EmptyResultSet GenericDeviceModelStatusEnumType_1 = "EmptyResultSet"
const GenericDeviceModelStatusEnumType_1_NotSupported GenericDeviceModelStatusEnumType_1 = "NotSupported"
const GenericDeviceModelStatusEnumType_1_Rejected GenericDeviceModelStatusEnumType_1 = "Rejected"

type GenericDeviceModelStatusEnumType_2 string

const GenericDeviceModelStatusEnumType_2_Accepted GenericDeviceModelStatusEnumType_2 = "Accepted"
const GenericDeviceModelStatusEnumType_2_EmptyResultSet GenericDeviceModelStatusEnumType_2 = "EmptyResultSet"
const GenericDeviceModelStatusEnumType_2_NotSupported GenericDeviceModelStatusEnumType_2 = "NotSupported"
const GenericDeviceModelStatusEnumType_2_Rejected GenericDeviceModelStatusEnumType_2 = "Rejected"

type GenericDeviceModelStatusEnumType_3 string

const GenericDeviceModelStatusEnumType_3_Accepted GenericDeviceModelStatusEnumType_3 = "Accepted"
const GenericDeviceModelStatusEnumType_3_EmptyResultSet GenericDeviceModelStatusEnumType_3 = "EmptyResultSet"
const GenericDeviceModelStatusEnumType_3_NotSupported GenericDeviceModelStatusEnumType_3 = "NotSupported"
const GenericDeviceModelStatusEnumType_3_Rejected GenericDeviceModelStatusEnumType_3 = "Rejected"

type GenericDeviceModelStatusEnumType_4 string

const GenericDeviceModelStatusEnumType_4_Accepted GenericDeviceModelStatusEnumType_4 = "Accepted"
const GenericDeviceModelStatusEnumType_4_EmptyResultSet GenericDeviceModelStatusEnumType_4 = "EmptyResultSet"
const GenericDeviceModelStatusEnumType_4_NotSupported GenericDeviceModelStatusEnumType_4 = "NotSupported"
const GenericDeviceModelStatusEnumType_4_Rejected GenericDeviceModelStatusEnumType_4 = "Rejected"

type GenericDeviceModelStatusEnumType_5 string

const GenericDeviceModelStatusEnumType_5_Accepted GenericDeviceModelStatusEnumType_5 = "Accepted"
const GenericDeviceModelStatusEnumType_5_EmptyResultSet GenericDeviceModelStatusEnumType_5 = "EmptyResultSet"
const GenericDeviceModelStatusEnumType_5_NotSupported GenericDeviceModelStatusEnumType_5 = "NotSupported"
const GenericDeviceModelStatusEnumType_5_Rejected GenericDeviceModelStatusEnumType_5 = "Rejected"

type GenericDeviceModelStatusEnumType_6 string

const GenericDeviceModelStatusEnumType_6_Accepted GenericDeviceModelStatusEnumType_6 = "Accepted"
const GenericDeviceModelStatusEnumType_6_EmptyResultSet GenericDeviceModelStatusEnumType_6 = "EmptyResultSet"
const GenericDeviceModelStatusEnumType_6_NotSupported GenericDeviceModelStatusEnumType_6 = "NotSupported"
const GenericDeviceModelStatusEnumType_6_Rejected GenericDeviceModelStatusEnumType_6 = "Rejected"

type GenericDeviceModelStatusEnumType_7 string

const GenericDeviceModelStatusEnumType_7_Accepted GenericDeviceModelStatusEnumType_7 = "Accepted"
const GenericDeviceModelStatusEnumType_7_EmptyResultSet GenericDeviceModelStatusEnumType_7 = "EmptyResultSet"
const GenericDeviceModelStatusEnumType_7_NotSupported GenericDeviceModelStatusEnumType_7 = "NotSupported"
const GenericDeviceModelStatusEnumType_7_Rejected GenericDeviceModelStatusEnumType_7 = "Rejected"

type GenericStatusEnumType string

const GenericStatusEnumTypeAccepted GenericStatusEnumType = "Accepted"
const GenericStatusEnumTypeRejected GenericStatusEnumType = "Rejected"

type GenericStatusEnumType_1 string

const GenericStatusEnumType_1_Accepted GenericStatusEnumType_1 = "Accepted"
const GenericStatusEnumType_1_Rejected GenericStatusEnumType_1 = "Rejected"

type GenericStatusEnumType_2 string

const GenericStatusEnumType_2_Accepted GenericStatusEnumType_2 = "Accepted"
const GenericStatusEnumType_2_Rejected GenericStatusEnumType_2 = "Rejected"

type GenericStatusEnumType_3 string

const GenericStatusEnumType_3_Accepted GenericStatusEnumType_3 = "Accepted"
const GenericStatusEnumType_3_Rejected GenericStatusEnumType_3 = "Rejected"

type GenericStatusEnumType_4 string

const GenericStatusEnumType_4_Accepted GenericStatusEnumType_4 = "Accepted"
const GenericStatusEnumType_4_Rejected GenericStatusEnumType_4 = "Rejected"

type GenericStatusEnumType_5 string

const GenericStatusEnumType_5_Accepted GenericStatusEnumType_5 = "Accepted"
const GenericStatusEnumType_5_Rejected GenericStatusEnumType_5 = "Rejected"

type GenericStatusEnumType_6 string

const GenericStatusEnumType_6_Accepted GenericStatusEnumType_6 = "Accepted"
const GenericStatusEnumType_6_Rejected GenericStatusEnumType_6 = "Rejected"

type GenericStatusEnumType_7 string

const GenericStatusEnumType_7_Accepted GenericStatusEnumType_7 = "Accepted"
const GenericStatusEnumType_7_Rejected GenericStatusEnumType_7 = "Rejected"

type GenericStatusEnumType_8 string

const GenericStatusEnumType_8_Accepted GenericStatusEnumType_8 = "Accepted"
const GenericStatusEnumType_8_Rejected GenericStatusEnumType_8 = "Rejected"

type GenericStatusEnumType_9 string

const GenericStatusEnumType_9_Accepted GenericStatusEnumType_9 = "Accepted"
const GenericStatusEnumType_9_Rejected GenericStatusEnumType_9 = "Rejected"

type Get15118EVCertificateRequestJson struct {
	// Action corresponds to the JSON schema field "action".
	Action CertificateActionEnumType_1 `json:"action"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_30 `json:"customData,omitempty"`

	// Raw CertificateInstallationReq request from EV, Base64 encoded.
	//
	ExiRequest string `json:"exiRequest"`

	// Schema version currently used for the 15118 session between EV and Charging
	// Station. Needed for parsing of the EXI stream by the CSMS.
	//
	//
	Iso15118SchemaVersion string `json:"iso15118SchemaVersion"`
}

type Get15118EVCertificateResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_31 `json:"customData,omitempty"`

	// Raw CertificateInstallationRes response for the EV, Base64 encoded.
	//
	ExiResponse string `json:"exiResponse"`

	// Status corresponds to the JSON schema field "status".
	Status Iso15118EVCertificateStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_11 `json:"statusInfo,omitempty"`
}

type GetBaseReportRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_32 `json:"customData,omitempty"`

	// ReportBase corresponds to the JSON schema field "reportBase".
	ReportBase ReportBaseEnumType_1 `json:"reportBase"`

	// The Id of the request.
	//
	RequestId int `json:"requestId"`
}

type GetBaseReportResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_33 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericDeviceModelStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_12 `json:"statusInfo,omitempty"`
}

type GetCertificateIdUseEnumType string

const GetCertificateIdUseEnumTypeCSMSRootCertificate GetCertificateIdUseEnumType = "CSMSRootCertificate"
const GetCertificateIdUseEnumTypeMORootCertificate GetCertificateIdUseEnumType = "MORootCertificate"
const GetCertificateIdUseEnumTypeManufacturerRootCertificate GetCertificateIdUseEnumType = "ManufacturerRootCertificate"
const GetCertificateIdUseEnumTypeV2GCertificateChain GetCertificateIdUseEnumType = "V2GCertificateChain"
const GetCertificateIdUseEnumTypeV2GRootCertificate GetCertificateIdUseEnumType = "V2GRootCertificate"

type GetCertificateIdUseEnumType_1 string

const GetCertificateIdUseEnumType_1_CSMSRootCertificate GetCertificateIdUseEnumType_1 = "CSMSRootCertificate"
const GetCertificateIdUseEnumType_1_MORootCertificate GetCertificateIdUseEnumType_1 = "MORootCertificate"
const GetCertificateIdUseEnumType_1_ManufacturerRootCertificate GetCertificateIdUseEnumType_1 = "ManufacturerRootCertificate"
const GetCertificateIdUseEnumType_1_V2GCertificateChain GetCertificateIdUseEnumType_1 = "V2GCertificateChain"
const GetCertificateIdUseEnumType_1_V2GRootCertificate GetCertificateIdUseEnumType_1 = "V2GRootCertificate"

type GetCertificateIdUseEnumType_2 string

const GetCertificateIdUseEnumType_2_CSMSRootCertificate GetCertificateIdUseEnumType_2 = "CSMSRootCertificate"
const GetCertificateIdUseEnumType_2_MORootCertificate GetCertificateIdUseEnumType_2 = "MORootCertificate"
const GetCertificateIdUseEnumType_2_ManufacturerRootCertificate GetCertificateIdUseEnumType_2 = "ManufacturerRootCertificate"
const GetCertificateIdUseEnumType_2_V2GCertificateChain GetCertificateIdUseEnumType_2 = "V2GCertificateChain"
const GetCertificateIdUseEnumType_2_V2GRootCertificate GetCertificateIdUseEnumType_2 = "V2GRootCertificate"

type GetCertificateIdUseEnumType_3 string

const GetCertificateIdUseEnumType_3_CSMSRootCertificate GetCertificateIdUseEnumType_3 = "CSMSRootCertificate"
const GetCertificateIdUseEnumType_3_MORootCertificate GetCertificateIdUseEnumType_3 = "MORootCertificate"
const GetCertificateIdUseEnumType_3_ManufacturerRootCertificate GetCertificateIdUseEnumType_3 = "ManufacturerRootCertificate"
const GetCertificateIdUseEnumType_3_V2GCertificateChain GetCertificateIdUseEnumType_3 = "V2GCertificateChain"
const GetCertificateIdUseEnumType_3_V2GRootCertificate GetCertificateIdUseEnumType_3 = "V2GRootCertificate"

type GetCertificateStatusEnumType string

const GetCertificateStatusEnumTypeAccepted GetCertificateStatusEnumType = "Accepted"
const GetCertificateStatusEnumTypeFailed GetCertificateStatusEnumType = "Failed"

type GetCertificateStatusEnumType_1 string

const GetCertificateStatusEnumType_1_Accepted GetCertificateStatusEnumType_1 = "Accepted"
const GetCertificateStatusEnumType_1_Failed GetCertificateStatusEnumType_1 = "Failed"

type GetCertificateStatusRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_34 `json:"customData,omitempty"`

	// OcspRequestData corresponds to the JSON schema field "ocspRequestData".
	OcspRequestData OCSPRequestDataType_1 `json:"ocspRequestData"`
}

type GetCertificateStatusResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_35 `json:"customData,omitempty"`

	// OCSPResponse class as defined in &lt;&lt;ref-ocpp_security_24, IETF RFC
	// 6960&gt;&gt;. DER encoded (as defined in &lt;&lt;ref-ocpp_security_24, IETF RFC
	// 6960&gt;&gt;), and then base64 encoded. MAY only be omitted when status is not
	// Accepted.
	//
	OcspResult *string `json:"ocspResult,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GetCertificateStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_13 `json:"statusInfo,omitempty"`
}

type GetChargingProfileStatusEnumType string

const GetChargingProfileStatusEnumTypeAccepted GetChargingProfileStatusEnumType = "Accepted"
const GetChargingProfileStatusEnumTypeNoProfiles GetChargingProfileStatusEnumType = "NoProfiles"

type GetChargingProfileStatusEnumType_1 string

const GetChargingProfileStatusEnumType_1_Accepted GetChargingProfileStatusEnumType_1 = "Accepted"
const GetChargingProfileStatusEnumType_1_NoProfiles GetChargingProfileStatusEnumType_1 = "NoProfiles"

type GetChargingProfilesRequestJson struct {
	// ChargingProfile corresponds to the JSON schema field "chargingProfile".
	ChargingProfile ChargingProfileCriterionType `json:"chargingProfile"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_36 `json:"customData,omitempty"`

	// For which EVSE installed charging profiles SHALL be reported. If 0, only
	// charging profiles installed on the Charging Station itself (the grid
	// connection) SHALL be reported. If omitted, all installed charging profiles
	// SHALL be reported.
	//
	EvseId *int `json:"evseId,omitempty"`

	// Reference identification that is to be used by the Charging Station in the
	// &lt;&lt;reportchargingprofilesrequest, ReportChargingProfilesRequest&gt;&gt;
	// when provided.
	//
	RequestId int `json:"requestId"`
}

type GetChargingProfilesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_37 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GetChargingProfileStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_14 `json:"statusInfo,omitempty"`
}

type GetCompositeScheduleRequestJson struct {
	// ChargingRateUnit corresponds to the JSON schema field "chargingRateUnit".
	ChargingRateUnit *ChargingRateUnitEnumType_1 `json:"chargingRateUnit,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_38 `json:"customData,omitempty"`

	// Length of the requested schedule in seconds.
	//
	//
	Duration int `json:"duration"`

	// The ID of the EVSE for which the schedule is requested. When evseid=0, the
	// Charging Station will calculate the expected consumption for the grid
	// connection.
	//
	EvseId int `json:"evseId"`
}

type GetCompositeScheduleResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_39 `json:"customData,omitempty"`

	// Schedule corresponds to the JSON schema field "schedule".
	Schedule *CompositeScheduleType `json:"schedule,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_15 `json:"statusInfo,omitempty"`
}

type GetDisplayMessagesRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_40 `json:"customData,omitempty"`

	// If provided the Charging Station shall return Display Messages of the given
	// ids. This field SHALL NOT contain more ids than set in
	// &lt;&lt;configkey-number-of-display-messages,NumberOfDisplayMessages.maxLimit&gt;&gt;
	//
	//
	Id []int `json:"id,omitempty"`

	// Priority corresponds to the JSON schema field "priority".
	Priority *MessagePriorityEnumType_1 `json:"priority,omitempty"`

	// The Id of this request.
	//
	RequestId int `json:"requestId"`

	// State corresponds to the JSON schema field "state".
	State *MessageStateEnumType_1 `json:"state,omitempty"`
}

type GetDisplayMessagesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_41 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GetDisplayMessagesStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_16 `json:"statusInfo,omitempty"`
}

type GetDisplayMessagesStatusEnumType string

const GetDisplayMessagesStatusEnumTypeAccepted GetDisplayMessagesStatusEnumType = "Accepted"
const GetDisplayMessagesStatusEnumTypeUnknown GetDisplayMessagesStatusEnumType = "Unknown"

type GetDisplayMessagesStatusEnumType_1 string

const GetDisplayMessagesStatusEnumType_1_Accepted GetDisplayMessagesStatusEnumType_1 = "Accepted"
const GetDisplayMessagesStatusEnumType_1_Unknown GetDisplayMessagesStatusEnumType_1 = "Unknown"

type GetInstalledCertificateIdsRequestJson struct {
	// Indicates the type of certificates requested. When omitted, all certificate
	// types are requested.
	//
	CertificateType []GetCertificateIdUseEnumType_1 `json:"certificateType,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_42 `json:"customData,omitempty"`
}

type GetInstalledCertificateIdsResponseJson struct {
	// CertificateHashDataChain corresponds to the JSON schema field
	// "certificateHashDataChain".
	CertificateHashDataChain []CertificateHashDataChainType `json:"certificateHashDataChain,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_43 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GetInstalledCertificateStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_17 `json:"statusInfo,omitempty"`
}

type GetInstalledCertificateStatusEnumType string

const GetInstalledCertificateStatusEnumTypeAccepted GetInstalledCertificateStatusEnumType = "Accepted"
const GetInstalledCertificateStatusEnumTypeNotFound GetInstalledCertificateStatusEnumType = "NotFound"

type GetInstalledCertificateStatusEnumType_1 string

const GetInstalledCertificateStatusEnumType_1_Accepted GetInstalledCertificateStatusEnumType_1 = "Accepted"
const GetInstalledCertificateStatusEnumType_1_NotFound GetInstalledCertificateStatusEnumType_1 = "NotFound"

type GetLocalListVersionRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_44 `json:"customData,omitempty"`
}

type GetLocalListVersionResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_45 `json:"customData,omitempty"`

	// This contains the current version number of the local authorization list in the
	// Charging Station.
	//
	VersionNumber int `json:"versionNumber"`
}

type GetLogRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_46 `json:"customData,omitempty"`

	// Log corresponds to the JSON schema field "log".
	Log LogParametersType `json:"log"`

	// LogType corresponds to the JSON schema field "logType".
	LogType LogEnumType_1 `json:"logType"`

	// The Id of this request
	//
	RequestId int `json:"requestId"`

	// This specifies how many times the Charging Station must try to upload the log
	// before giving up. If this field is not present, it is left to Charging Station
	// to decide how many times it wants to retry.
	//
	Retries *int `json:"retries,omitempty"`

	// The interval in seconds after which a retry may be attempted. If this field is
	// not present, it is left to Charging Station to decide how long to wait between
	// attempts.
	//
	RetryInterval *int `json:"retryInterval,omitempty"`
}

type GetLogResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_47 `json:"customData,omitempty"`

	// This contains the name of the log file that will be uploaded. This field is not
	// present when no logging information is available.
	//
	Filename *string `json:"filename,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status LogStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_18 `json:"statusInfo,omitempty"`
}

type GetMonitoringReportRequestJson struct {
	// ComponentVariable corresponds to the JSON schema field "componentVariable".
	ComponentVariable []ComponentVariableType `json:"componentVariable,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_48 `json:"customData,omitempty"`

	// This field contains criteria for components for which a monitoring report is
	// requested
	//
	MonitoringCriteria []MonitoringCriterionEnumType_1 `json:"monitoringCriteria,omitempty"`

	// The Id of the request.
	//
	RequestId int `json:"requestId"`
}

type GetMonitoringReportResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_49 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericDeviceModelStatusEnumType_3 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_19 `json:"statusInfo,omitempty"`
}

type GetReportRequestJson struct {
	// This field contains criteria for components for which a report is requested
	//
	ComponentCriteria []ComponentCriterionEnumType_1 `json:"componentCriteria,omitempty"`

	// ComponentVariable corresponds to the JSON schema field "componentVariable".
	ComponentVariable []ComponentVariableType_1 `json:"componentVariable,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_50 `json:"customData,omitempty"`

	// The Id of the request.
	//
	RequestId int `json:"requestId"`
}

type GetReportResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_51 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericDeviceModelStatusEnumType_5 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_20 `json:"statusInfo,omitempty"`
}

type GetTransactionStatusRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_52 `json:"customData,omitempty"`

	// The Id of the transaction for which the status is requested.
	//
	TransactionId *string `json:"transactionId,omitempty"`
}

type GetTransactionStatusResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_53 `json:"customData,omitempty"`

	// Whether there are still message to be delivered.
	//
	MessagesInQueue bool `json:"messagesInQueue"`

	// Whether the transaction is still ongoing.
	//
	OngoingIndicator *bool `json:"ongoingIndicator,omitempty"`
}

// Class to hold parameters for GetVariables request.
//
type GetVariableDataType struct {
	// AttributeType corresponds to the JSON schema field "attributeType".
	AttributeType *AttributeEnumType_1 `json:"attributeType,omitempty"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType_2 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_54 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_2 `json:"variable"`
}

// Class to hold results of GetVariables request.
//
type GetVariableResultType struct {
	// AttributeStatus corresponds to the JSON schema field "attributeStatus".
	AttributeStatus GetVariableStatusEnumType `json:"attributeStatus"`

	// AttributeStatusInfo corresponds to the JSON schema field "attributeStatusInfo".
	AttributeStatusInfo *StatusInfoType_21 `json:"attributeStatusInfo,omitempty"`

	// AttributeType corresponds to the JSON schema field "attributeType".
	AttributeType *AttributeEnumType_3 `json:"attributeType,omitempty"`

	// Value of requested attribute type of component-variable. This field can only be
	// empty when the given status is NOT accepted.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-reporting-value-size,ReportingValueSize&gt;&gt; can be used
	// to limit GetVariableResult.attributeValue, VariableAttribute.value and
	// EventData.actualValue. The max size of these values will always remain equal.
	//
	//
	AttributeValue *string `json:"attributeValue,omitempty"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType_3 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_55 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_3 `json:"variable"`
}

type GetVariableStatusEnumType string

const GetVariableStatusEnumTypeAccepted GetVariableStatusEnumType = "Accepted"
const GetVariableStatusEnumTypeNotSupportedAttributeType GetVariableStatusEnumType = "NotSupportedAttributeType"
const GetVariableStatusEnumTypeRejected GetVariableStatusEnumType = "Rejected"
const GetVariableStatusEnumTypeUnknownComponent GetVariableStatusEnumType = "UnknownComponent"
const GetVariableStatusEnumTypeUnknownVariable GetVariableStatusEnumType = "UnknownVariable"

type GetVariableStatusEnumType_1 string

const GetVariableStatusEnumType_1_Accepted GetVariableStatusEnumType_1 = "Accepted"
const GetVariableStatusEnumType_1_NotSupportedAttributeType GetVariableStatusEnumType_1 = "NotSupportedAttributeType"
const GetVariableStatusEnumType_1_Rejected GetVariableStatusEnumType_1 = "Rejected"
const GetVariableStatusEnumType_1_UnknownComponent GetVariableStatusEnumType_1 = "UnknownComponent"
const GetVariableStatusEnumType_1_UnknownVariable GetVariableStatusEnumType_1 = "UnknownVariable"

type GetVariablesRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_54 `json:"customData,omitempty"`

	// GetVariableData corresponds to the JSON schema field "getVariableData".
	GetVariableData []GetVariableDataType `json:"getVariableData"`
}

type GetVariablesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_55 `json:"customData,omitempty"`

	// GetVariableResult corresponds to the JSON schema field "getVariableResult".
	GetVariableResult []GetVariableResultType `json:"getVariableResult"`
}

type HashAlgorithmEnumType string

const HashAlgorithmEnumTypeSHA256 HashAlgorithmEnumType = "SHA256"
const HashAlgorithmEnumTypeSHA384 HashAlgorithmEnumType = "SHA384"
const HashAlgorithmEnumTypeSHA512 HashAlgorithmEnumType = "SHA512"

type HashAlgorithmEnumType_1 string

const HashAlgorithmEnumType_1_SHA256 HashAlgorithmEnumType_1 = "SHA256"
const HashAlgorithmEnumType_1_SHA384 HashAlgorithmEnumType_1 = "SHA384"
const HashAlgorithmEnumType_1_SHA512 HashAlgorithmEnumType_1 = "SHA512"

type HashAlgorithmEnumType_2 string

const HashAlgorithmEnumType_2_SHA256 HashAlgorithmEnumType_2 = "SHA256"
const HashAlgorithmEnumType_2_SHA384 HashAlgorithmEnumType_2 = "SHA384"
const HashAlgorithmEnumType_2_SHA512 HashAlgorithmEnumType_2 = "SHA512"

type HashAlgorithmEnumType_3 string

const HashAlgorithmEnumType_3_SHA256 HashAlgorithmEnumType_3 = "SHA256"
const HashAlgorithmEnumType_3_SHA384 HashAlgorithmEnumType_3 = "SHA384"
const HashAlgorithmEnumType_3_SHA512 HashAlgorithmEnumType_3 = "SHA512"

type HashAlgorithmEnumType_4 string

const HashAlgorithmEnumType_4_SHA256 HashAlgorithmEnumType_4 = "SHA256"
const HashAlgorithmEnumType_4_SHA384 HashAlgorithmEnumType_4 = "SHA384"
const HashAlgorithmEnumType_4_SHA512 HashAlgorithmEnumType_4 = "SHA512"

type HashAlgorithmEnumType_5 string

const HashAlgorithmEnumType_5_SHA256 HashAlgorithmEnumType_5 = "SHA256"
const HashAlgorithmEnumType_5_SHA384 HashAlgorithmEnumType_5 = "SHA384"
const HashAlgorithmEnumType_5_SHA512 HashAlgorithmEnumType_5 = "SHA512"

type HashAlgorithmEnumType_6 string

const HashAlgorithmEnumType_6_SHA256 HashAlgorithmEnumType_6 = "SHA256"
const HashAlgorithmEnumType_6_SHA384 HashAlgorithmEnumType_6 = "SHA384"
const HashAlgorithmEnumType_6_SHA512 HashAlgorithmEnumType_6 = "SHA512"

type HashAlgorithmEnumType_7 string

const HashAlgorithmEnumType_7_SHA256 HashAlgorithmEnumType_7 = "SHA256"
const HashAlgorithmEnumType_7_SHA384 HashAlgorithmEnumType_7 = "SHA384"
const HashAlgorithmEnumType_7_SHA512 HashAlgorithmEnumType_7 = "SHA512"

type HashAlgorithmEnumType_8 string

const HashAlgorithmEnumType_8_SHA256 HashAlgorithmEnumType_8 = "SHA256"
const HashAlgorithmEnumType_8_SHA384 HashAlgorithmEnumType_8 = "SHA384"
const HashAlgorithmEnumType_8_SHA512 HashAlgorithmEnumType_8 = "SHA512"

type HashAlgorithmEnumType_9 string

const HashAlgorithmEnumType_9_SHA256 HashAlgorithmEnumType_9 = "SHA256"
const HashAlgorithmEnumType_9_SHA384 HashAlgorithmEnumType_9 = "SHA384"
const HashAlgorithmEnumType_9_SHA512 HashAlgorithmEnumType_9 = "SHA512"

type HeartbeatRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_56 `json:"customData,omitempty"`
}

type HeartbeatResponseJson struct {
	// Contains the current time of the CSMS.
	//
	CurrentTime string `json:"currentTime"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_57 `json:"customData,omitempty"`
}

type IdTokenEnumType string

const IdTokenEnumTypeCentral IdTokenEnumType = "Central"
const IdTokenEnumTypeEMAID IdTokenEnumType = "eMAID"
const IdTokenEnumTypeISO14443 IdTokenEnumType = "ISO14443"
const IdTokenEnumTypeISO15693 IdTokenEnumType = "ISO15693"
const IdTokenEnumTypeKeyCode IdTokenEnumType = "KeyCode"
const IdTokenEnumTypeLocal IdTokenEnumType = "Local"
const IdTokenEnumTypeMacAddress IdTokenEnumType = "MacAddress"
const IdTokenEnumTypeNoAuthorization IdTokenEnumType = "NoAuthorization"

type IdTokenEnumType_1 string

type IdTokenEnumType_10 string

const IdTokenEnumType_10_Central IdTokenEnumType_10 = "Central"
const IdTokenEnumType_10_EMAID IdTokenEnumType_10 = "eMAID"
const IdTokenEnumType_10_ISO14443 IdTokenEnumType_10 = "ISO14443"
const IdTokenEnumType_10_ISO15693 IdTokenEnumType_10 = "ISO15693"
const IdTokenEnumType_10_KeyCode IdTokenEnumType_10 = "KeyCode"
const IdTokenEnumType_10_Local IdTokenEnumType_10 = "Local"
const IdTokenEnumType_10_MacAddress IdTokenEnumType_10 = "MacAddress"
const IdTokenEnumType_10_NoAuthorization IdTokenEnumType_10 = "NoAuthorization"

type IdTokenEnumType_11 string

const IdTokenEnumType_11_Central IdTokenEnumType_11 = "Central"
const IdTokenEnumType_11_EMAID IdTokenEnumType_11 = "eMAID"
const IdTokenEnumType_11_ISO14443 IdTokenEnumType_11 = "ISO14443"
const IdTokenEnumType_11_ISO15693 IdTokenEnumType_11 = "ISO15693"
const IdTokenEnumType_11_KeyCode IdTokenEnumType_11 = "KeyCode"
const IdTokenEnumType_11_Local IdTokenEnumType_11 = "Local"
const IdTokenEnumType_11_MacAddress IdTokenEnumType_11 = "MacAddress"
const IdTokenEnumType_11_NoAuthorization IdTokenEnumType_11 = "NoAuthorization"

type IdTokenEnumType_12 string

const IdTokenEnumType_12_Central IdTokenEnumType_12 = "Central"
const IdTokenEnumType_12_EMAID IdTokenEnumType_12 = "eMAID"
const IdTokenEnumType_12_ISO14443 IdTokenEnumType_12 = "ISO14443"
const IdTokenEnumType_12_ISO15693 IdTokenEnumType_12 = "ISO15693"
const IdTokenEnumType_12_KeyCode IdTokenEnumType_12 = "KeyCode"
const IdTokenEnumType_12_Local IdTokenEnumType_12 = "Local"
const IdTokenEnumType_12_MacAddress IdTokenEnumType_12 = "MacAddress"
const IdTokenEnumType_12_NoAuthorization IdTokenEnumType_12 = "NoAuthorization"

type IdTokenEnumType_13 string

const IdTokenEnumType_13_Central IdTokenEnumType_13 = "Central"
const IdTokenEnumType_13_EMAID IdTokenEnumType_13 = "eMAID"
const IdTokenEnumType_13_ISO14443 IdTokenEnumType_13 = "ISO14443"
const IdTokenEnumType_13_ISO15693 IdTokenEnumType_13 = "ISO15693"
const IdTokenEnumType_13_KeyCode IdTokenEnumType_13 = "KeyCode"
const IdTokenEnumType_13_Local IdTokenEnumType_13 = "Local"
const IdTokenEnumType_13_MacAddress IdTokenEnumType_13 = "MacAddress"
const IdTokenEnumType_13_NoAuthorization IdTokenEnumType_13 = "NoAuthorization"

type IdTokenEnumType_14 string

const IdTokenEnumType_14_Central IdTokenEnumType_14 = "Central"
const IdTokenEnumType_14_EMAID IdTokenEnumType_14 = "eMAID"
const IdTokenEnumType_14_ISO14443 IdTokenEnumType_14 = "ISO14443"
const IdTokenEnumType_14_ISO15693 IdTokenEnumType_14 = "ISO15693"
const IdTokenEnumType_14_KeyCode IdTokenEnumType_14 = "KeyCode"
const IdTokenEnumType_14_Local IdTokenEnumType_14 = "Local"
const IdTokenEnumType_14_MacAddress IdTokenEnumType_14 = "MacAddress"
const IdTokenEnumType_14_NoAuthorization IdTokenEnumType_14 = "NoAuthorization"

type IdTokenEnumType_15 string

const IdTokenEnumType_15_Central IdTokenEnumType_15 = "Central"
const IdTokenEnumType_15_EMAID IdTokenEnumType_15 = "eMAID"
const IdTokenEnumType_15_ISO14443 IdTokenEnumType_15 = "ISO14443"
const IdTokenEnumType_15_ISO15693 IdTokenEnumType_15 = "ISO15693"
const IdTokenEnumType_15_KeyCode IdTokenEnumType_15 = "KeyCode"
const IdTokenEnumType_15_Local IdTokenEnumType_15 = "Local"
const IdTokenEnumType_15_MacAddress IdTokenEnumType_15 = "MacAddress"
const IdTokenEnumType_15_NoAuthorization IdTokenEnumType_15 = "NoAuthorization"
const IdTokenEnumType_1_Central IdTokenEnumType_1 = "Central"
const IdTokenEnumType_1_EMAID IdTokenEnumType_1 = "eMAID"
const IdTokenEnumType_1_ISO14443 IdTokenEnumType_1 = "ISO14443"
const IdTokenEnumType_1_ISO15693 IdTokenEnumType_1 = "ISO15693"
const IdTokenEnumType_1_KeyCode IdTokenEnumType_1 = "KeyCode"
const IdTokenEnumType_1_Local IdTokenEnumType_1 = "Local"
const IdTokenEnumType_1_MacAddress IdTokenEnumType_1 = "MacAddress"
const IdTokenEnumType_1_NoAuthorization IdTokenEnumType_1 = "NoAuthorization"

type IdTokenEnumType_2 string

const IdTokenEnumType_2_Central IdTokenEnumType_2 = "Central"
const IdTokenEnumType_2_EMAID IdTokenEnumType_2 = "eMAID"
const IdTokenEnumType_2_ISO14443 IdTokenEnumType_2 = "ISO14443"
const IdTokenEnumType_2_ISO15693 IdTokenEnumType_2 = "ISO15693"
const IdTokenEnumType_2_KeyCode IdTokenEnumType_2 = "KeyCode"
const IdTokenEnumType_2_Local IdTokenEnumType_2 = "Local"
const IdTokenEnumType_2_MacAddress IdTokenEnumType_2 = "MacAddress"
const IdTokenEnumType_2_NoAuthorization IdTokenEnumType_2 = "NoAuthorization"

type IdTokenEnumType_3 string

const IdTokenEnumType_3_Central IdTokenEnumType_3 = "Central"
const IdTokenEnumType_3_EMAID IdTokenEnumType_3 = "eMAID"
const IdTokenEnumType_3_ISO14443 IdTokenEnumType_3 = "ISO14443"
const IdTokenEnumType_3_ISO15693 IdTokenEnumType_3 = "ISO15693"
const IdTokenEnumType_3_KeyCode IdTokenEnumType_3 = "KeyCode"
const IdTokenEnumType_3_Local IdTokenEnumType_3 = "Local"
const IdTokenEnumType_3_MacAddress IdTokenEnumType_3 = "MacAddress"
const IdTokenEnumType_3_NoAuthorization IdTokenEnumType_3 = "NoAuthorization"

type IdTokenEnumType_4 string

const IdTokenEnumType_4_Central IdTokenEnumType_4 = "Central"
const IdTokenEnumType_4_EMAID IdTokenEnumType_4 = "eMAID"
const IdTokenEnumType_4_ISO14443 IdTokenEnumType_4 = "ISO14443"
const IdTokenEnumType_4_ISO15693 IdTokenEnumType_4 = "ISO15693"
const IdTokenEnumType_4_KeyCode IdTokenEnumType_4 = "KeyCode"
const IdTokenEnumType_4_Local IdTokenEnumType_4 = "Local"
const IdTokenEnumType_4_MacAddress IdTokenEnumType_4 = "MacAddress"
const IdTokenEnumType_4_NoAuthorization IdTokenEnumType_4 = "NoAuthorization"

type IdTokenEnumType_5 string

const IdTokenEnumType_5_Central IdTokenEnumType_5 = "Central"
const IdTokenEnumType_5_EMAID IdTokenEnumType_5 = "eMAID"
const IdTokenEnumType_5_ISO14443 IdTokenEnumType_5 = "ISO14443"
const IdTokenEnumType_5_ISO15693 IdTokenEnumType_5 = "ISO15693"
const IdTokenEnumType_5_KeyCode IdTokenEnumType_5 = "KeyCode"
const IdTokenEnumType_5_Local IdTokenEnumType_5 = "Local"
const IdTokenEnumType_5_MacAddress IdTokenEnumType_5 = "MacAddress"
const IdTokenEnumType_5_NoAuthorization IdTokenEnumType_5 = "NoAuthorization"

type IdTokenEnumType_6 string

const IdTokenEnumType_6_Central IdTokenEnumType_6 = "Central"
const IdTokenEnumType_6_EMAID IdTokenEnumType_6 = "eMAID"
const IdTokenEnumType_6_ISO14443 IdTokenEnumType_6 = "ISO14443"
const IdTokenEnumType_6_ISO15693 IdTokenEnumType_6 = "ISO15693"
const IdTokenEnumType_6_KeyCode IdTokenEnumType_6 = "KeyCode"
const IdTokenEnumType_6_Local IdTokenEnumType_6 = "Local"
const IdTokenEnumType_6_MacAddress IdTokenEnumType_6 = "MacAddress"
const IdTokenEnumType_6_NoAuthorization IdTokenEnumType_6 = "NoAuthorization"

type IdTokenEnumType_7 string

const IdTokenEnumType_7_Central IdTokenEnumType_7 = "Central"
const IdTokenEnumType_7_EMAID IdTokenEnumType_7 = "eMAID"
const IdTokenEnumType_7_ISO14443 IdTokenEnumType_7 = "ISO14443"
const IdTokenEnumType_7_ISO15693 IdTokenEnumType_7 = "ISO15693"
const IdTokenEnumType_7_KeyCode IdTokenEnumType_7 = "KeyCode"
const IdTokenEnumType_7_Local IdTokenEnumType_7 = "Local"
const IdTokenEnumType_7_MacAddress IdTokenEnumType_7 = "MacAddress"
const IdTokenEnumType_7_NoAuthorization IdTokenEnumType_7 = "NoAuthorization"

type IdTokenEnumType_8 string

const IdTokenEnumType_8_Central IdTokenEnumType_8 = "Central"
const IdTokenEnumType_8_EMAID IdTokenEnumType_8 = "eMAID"
const IdTokenEnumType_8_ISO14443 IdTokenEnumType_8 = "ISO14443"
const IdTokenEnumType_8_ISO15693 IdTokenEnumType_8 = "ISO15693"
const IdTokenEnumType_8_KeyCode IdTokenEnumType_8 = "KeyCode"
const IdTokenEnumType_8_Local IdTokenEnumType_8 = "Local"
const IdTokenEnumType_8_MacAddress IdTokenEnumType_8 = "MacAddress"
const IdTokenEnumType_8_NoAuthorization IdTokenEnumType_8 = "NoAuthorization"

type IdTokenEnumType_9 string

const IdTokenEnumType_9_Central IdTokenEnumType_9 = "Central"
const IdTokenEnumType_9_EMAID IdTokenEnumType_9 = "eMAID"
const IdTokenEnumType_9_ISO14443 IdTokenEnumType_9 = "ISO14443"
const IdTokenEnumType_9_ISO15693 IdTokenEnumType_9 = "ISO15693"
const IdTokenEnumType_9_KeyCode IdTokenEnumType_9 = "KeyCode"
const IdTokenEnumType_9_Local IdTokenEnumType_9 = "Local"
const IdTokenEnumType_9_MacAddress IdTokenEnumType_9 = "MacAddress"
const IdTokenEnumType_9_NoAuthorization IdTokenEnumType_9 = "NoAuthorization"

// ID_ Token
// urn:x-oca:ocpp:uid:2:233247
// Contains status information about an identifier.
// It is advised to not stop charging for a token that expires during charging, as
// ExpiryDate is only used for caching purposes. If ExpiryDate is not given, the
// status has no end date.
//
type IdTokenInfoType struct {
	// ID_ Token. Expiry. Date_ Time
	// urn:x-oca:ocpp:uid:1:569373
	// Date and Time after which the token must be considered invalid.
	//
	CacheExpiryDateTime *string `json:"cacheExpiryDateTime,omitempty"`

	// Priority from a business point of view. Default priority is 0, The range is
	// from -9 to 9. Higher values indicate a higher priority. The chargingPriority in
	// &lt;&lt;transactioneventresponse,TransactionEventResponse&gt;&gt; overrules
	// this one.
	//
	ChargingPriority *int `json:"chargingPriority,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_1 `json:"customData,omitempty"`

	// Only used when the IdToken is only valid for one or more specific EVSEs, not
	// for the entire Charging Station.
	//
	//
	EvseId []int `json:"evseId,omitempty"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType_1 `json:"groupIdToken,omitempty"`

	// ID_ Token. Language1. Language_ Code
	// urn:x-oca:ocpp:uid:1:569374
	// Preferred user interface language of identifier user. Contains a language code
	// as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	//
	Language1 *string `json:"language1,omitempty"`

	// ID_ Token. Language2. Language_ Code
	// urn:x-oca:ocpp:uid:1:569375
	// Second preferred user interface language of identifier user. Don’t use when
	// language1 is omitted, has to be different from language1. Contains a language
	// code as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language2 *string `json:"language2,omitempty"`

	// PersonalMessage corresponds to the JSON schema field "personalMessage".
	PersonalMessage *MessageContentType `json:"personalMessage,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status AuthorizationStatusEnumType_1 `json:"status"`
}

// ID_ Token
// urn:x-oca:ocpp:uid:2:233247
// Contains status information about an identifier.
// It is advised to not stop charging for a token that expires during charging, as
// ExpiryDate is only used for caching purposes. If ExpiryDate is not given, the
// status has no end date.
//
type IdTokenInfoType_1 struct {
	// ID_ Token. Expiry. Date_ Time
	// urn:x-oca:ocpp:uid:1:569373
	// Date and Time after which the token must be considered invalid.
	//
	CacheExpiryDateTime *string `json:"cacheExpiryDateTime,omitempty"`

	// Priority from a business point of view. Default priority is 0, The range is
	// from -9 to 9. Higher values indicate a higher priority. The chargingPriority in
	// &lt;&lt;transactioneventresponse,TransactionEventResponse&gt;&gt; overrules
	// this one.
	//
	ChargingPriority *int `json:"chargingPriority,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_98 `json:"customData,omitempty"`

	// Only used when the IdToken is only valid for one or more specific EVSEs, not
	// for the entire Charging Station.
	//
	//
	EvseId []int `json:"evseId,omitempty"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType_5 `json:"groupIdToken,omitempty"`

	// ID_ Token. Language1. Language_ Code
	// urn:x-oca:ocpp:uid:1:569374
	// Preferred user interface language of identifier user. Contains a language code
	// as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	//
	Language1 *string `json:"language1,omitempty"`

	// ID_ Token. Language2. Language_ Code
	// urn:x-oca:ocpp:uid:1:569375
	// Second preferred user interface language of identifier user. Don’t use when
	// language1 is omitted, has to be different from language1. Contains a language
	// code as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language2 *string `json:"language2,omitempty"`

	// PersonalMessage corresponds to the JSON schema field "personalMessage".
	PersonalMessage *MessageContentType_2 `json:"personalMessage,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status AuthorizationStatusEnumType_2 `json:"status"`
}

// ID_ Token
// urn:x-oca:ocpp:uid:2:233247
// Contains status information about an identifier.
// It is advised to not stop charging for a token that expires during charging, as
// ExpiryDate is only used for caching purposes. If ExpiryDate is not given, the
// status has no end date.
//
type IdTokenInfoType_2 struct {
	// ID_ Token. Expiry. Date_ Time
	// urn:x-oca:ocpp:uid:1:569373
	// Date and Time after which the token must be considered invalid.
	//
	CacheExpiryDateTime *string `json:"cacheExpiryDateTime,omitempty"`

	// Priority from a business point of view. Default priority is 0, The range is
	// from -9 to 9. Higher values indicate a higher priority. The chargingPriority in
	// &lt;&lt;transactioneventresponse,TransactionEventResponse&gt;&gt; overrules
	// this one.
	//
	ChargingPriority *int `json:"chargingPriority,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_119 `json:"customData,omitempty"`

	// Only used when the IdToken is only valid for one or more specific EVSEs, not
	// for the entire Charging Station.
	//
	//
	EvseId []int `json:"evseId,omitempty"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType_7 `json:"groupIdToken,omitempty"`

	// ID_ Token. Language1. Language_ Code
	// urn:x-oca:ocpp:uid:1:569374
	// Preferred user interface language of identifier user. Contains a language code
	// as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	//
	Language1 *string `json:"language1,omitempty"`

	// ID_ Token. Language2. Language_ Code
	// urn:x-oca:ocpp:uid:1:569375
	// Second preferred user interface language of identifier user. Don’t use when
	// language1 is omitted, has to be different from language1. Contains a language
	// code as defined in &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language2 *string `json:"language2,omitempty"`

	// PersonalMessage corresponds to the JSON schema field "personalMessage".
	PersonalMessage *MessageContentType_4 `json:"personalMessage,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status AuthorizationStatusEnumType_5 `json:"status"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_1 `json:"type"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType_1 struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType_1 `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_1 `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_3 `json:"type"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType_2 struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType_2 `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_22 `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_5 `json:"type"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType_3 struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType_3 `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_7 `json:"type"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType_4 struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType_4 `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_92 `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_9 `json:"type"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType_5 struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType_5 `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_98 `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_10 `json:"type"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType_6 struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType_6 `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_13 `json:"type"`
}

// Contains a case insensitive identifier to use for the authorization and the type
// of authorization to support multiple forms of identifiers.
//
type IdTokenType_7 struct {
	// AdditionalInfo corresponds to the JSON schema field "additionalInfo".
	AdditionalInfo []AdditionalInfoType_7 `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_119 `json:"customData,omitempty"`

	// IdToken is case insensitive. Might hold the hidden id of an RFID tag, but can
	// for example also contain a UUID.
	//
	IdToken string `json:"idToken"`

	// Type corresponds to the JSON schema field "type".
	Type IdTokenEnumType_15 `json:"type"`
}

type InstallCertificateRequestJson struct {
	// A PEM encoded X.509 certificate.
	//
	Certificate string `json:"certificate"`

	// CertificateType corresponds to the JSON schema field "certificateType".
	CertificateType InstallCertificateUseEnumType_1 `json:"certificateType"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_58 `json:"customData,omitempty"`
}

type InstallCertificateResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_59 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status InstallCertificateStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_22 `json:"statusInfo,omitempty"`
}

type InstallCertificateStatusEnumType string

const InstallCertificateStatusEnumTypeAccepted InstallCertificateStatusEnumType = "Accepted"
const InstallCertificateStatusEnumTypeFailed InstallCertificateStatusEnumType = "Failed"
const InstallCertificateStatusEnumTypeRejected InstallCertificateStatusEnumType = "Rejected"

type InstallCertificateStatusEnumType_1 string

const InstallCertificateStatusEnumType_1_Accepted InstallCertificateStatusEnumType_1 = "Accepted"
const InstallCertificateStatusEnumType_1_Failed InstallCertificateStatusEnumType_1 = "Failed"
const InstallCertificateStatusEnumType_1_Rejected InstallCertificateStatusEnumType_1 = "Rejected"

type InstallCertificateUseEnumType string

const InstallCertificateUseEnumTypeCSMSRootCertificate InstallCertificateUseEnumType = "CSMSRootCertificate"
const InstallCertificateUseEnumTypeMORootCertificate InstallCertificateUseEnumType = "MORootCertificate"
const InstallCertificateUseEnumTypeManufacturerRootCertificate InstallCertificateUseEnumType = "ManufacturerRootCertificate"
const InstallCertificateUseEnumTypeV2GRootCertificate InstallCertificateUseEnumType = "V2GRootCertificate"

type InstallCertificateUseEnumType_1 string

const InstallCertificateUseEnumType_1_CSMSRootCertificate InstallCertificateUseEnumType_1 = "CSMSRootCertificate"
const InstallCertificateUseEnumType_1_MORootCertificate InstallCertificateUseEnumType_1 = "MORootCertificate"
const InstallCertificateUseEnumType_1_ManufacturerRootCertificate InstallCertificateUseEnumType_1 = "ManufacturerRootCertificate"
const InstallCertificateUseEnumType_1_V2GRootCertificate InstallCertificateUseEnumType_1 = "V2GRootCertificate"

type Iso15118EVCertificateStatusEnumType string

const Iso15118EVCertificateStatusEnumTypeAccepted Iso15118EVCertificateStatusEnumType = "Accepted"
const Iso15118EVCertificateStatusEnumTypeFailed Iso15118EVCertificateStatusEnumType = "Failed"

type Iso15118EVCertificateStatusEnumType_1 string

const Iso15118EVCertificateStatusEnumType_1_Accepted Iso15118EVCertificateStatusEnumType_1 = "Accepted"
const Iso15118EVCertificateStatusEnumType_1_Failed Iso15118EVCertificateStatusEnumType_1 = "Failed"

type LocationEnumType string

const LocationEnumTypeBody LocationEnumType = "Body"
const LocationEnumTypeCable LocationEnumType = "Cable"
const LocationEnumTypeEV LocationEnumType = "EV"
const LocationEnumTypeInlet LocationEnumType = "Inlet"
const LocationEnumTypeOutlet LocationEnumType = "Outlet"

type LocationEnumType_1 string

const LocationEnumType_1_Body LocationEnumType_1 = "Body"
const LocationEnumType_1_Cable LocationEnumType_1 = "Cable"
const LocationEnumType_1_EV LocationEnumType_1 = "EV"
const LocationEnumType_1_Inlet LocationEnumType_1 = "Inlet"
const LocationEnumType_1_Outlet LocationEnumType_1 = "Outlet"

type LocationEnumType_2 string

const LocationEnumType_2_Body LocationEnumType_2 = "Body"
const LocationEnumType_2_Cable LocationEnumType_2 = "Cable"
const LocationEnumType_2_EV LocationEnumType_2 = "EV"
const LocationEnumType_2_Inlet LocationEnumType_2 = "Inlet"
const LocationEnumType_2_Outlet LocationEnumType_2 = "Outlet"

type LocationEnumType_3 string

const LocationEnumType_3_Body LocationEnumType_3 = "Body"
const LocationEnumType_3_Cable LocationEnumType_3 = "Cable"
const LocationEnumType_3_EV LocationEnumType_3 = "EV"
const LocationEnumType_3_Inlet LocationEnumType_3 = "Inlet"
const LocationEnumType_3_Outlet LocationEnumType_3 = "Outlet"

type LogEnumType string

const LogEnumTypeDiagnosticsLog LogEnumType = "DiagnosticsLog"
const LogEnumTypeSecurityLog LogEnumType = "SecurityLog"

type LogEnumType_1 string

const LogEnumType_1_DiagnosticsLog LogEnumType_1 = "DiagnosticsLog"
const LogEnumType_1_SecurityLog LogEnumType_1 = "SecurityLog"

// Log
// urn:x-enexis:ecdm:uid:2:233373
// Generic class for the configuration of logging entries.
//
type LogParametersType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_46 `json:"customData,omitempty"`

	// Log. Latest_ Timestamp. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569482
	// This contains the date and time of the latest logging information to include in
	// the diagnostics.
	//
	LatestTimestamp *string `json:"latestTimestamp,omitempty"`

	// Log. Oldest_ Timestamp. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569477
	// This contains the date and time of the oldest logging information to include in
	// the diagnostics.
	//
	OldestTimestamp *string `json:"oldestTimestamp,omitempty"`

	// Log. Remote_ Location. URI
	// urn:x-enexis:ecdm:uid:1:569484
	// The URL of the location at the remote system where the log should be stored.
	//
	RemoteLocation string `json:"remoteLocation"`
}

type LogStatusEnumType string

const LogStatusEnumTypeAccepted LogStatusEnumType = "Accepted"
const LogStatusEnumTypeAcceptedCanceled LogStatusEnumType = "AcceptedCanceled"
const LogStatusEnumTypeRejected LogStatusEnumType = "Rejected"

type LogStatusEnumType_1 string

const LogStatusEnumType_1_Accepted LogStatusEnumType_1 = "Accepted"
const LogStatusEnumType_1_AcceptedCanceled LogStatusEnumType_1 = "AcceptedCanceled"
const LogStatusEnumType_1_Rejected LogStatusEnumType_1 = "Rejected"

type LogStatusNotificationRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_60 `json:"customData,omitempty"`

	// The request id that was provided in GetLogRequest that started this log upload.
	// This field is mandatory,
	// unless the message was triggered by a TriggerMessageRequest AND there is no log
	// upload ongoing.
	//
	RequestId *int `json:"requestId,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status UploadLogStatusEnumType_1 `json:"status"`
}

type LogStatusNotificationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_61 `json:"customData,omitempty"`
}

type MeasurandEnumType string

const MeasurandEnumTypeCurrentExport MeasurandEnumType = "Current.Export"
const MeasurandEnumTypeCurrentImport MeasurandEnumType = "Current.Import"
const MeasurandEnumTypeCurrentOffered MeasurandEnumType = "Current.Offered"
const MeasurandEnumTypeEnergyActiveExportInterval MeasurandEnumType = "Energy.Active.Export.Interval"
const MeasurandEnumTypeEnergyActiveExportRegister MeasurandEnumType = "Energy.Active.Export.Register"
const MeasurandEnumTypeEnergyActiveImportInterval MeasurandEnumType = "Energy.Active.Import.Interval"
const MeasurandEnumTypeEnergyActiveImportRegister MeasurandEnumType = "Energy.Active.Import.Register"
const MeasurandEnumTypeEnergyActiveNet MeasurandEnumType = "Energy.Active.Net"
const MeasurandEnumTypeEnergyApparentExport MeasurandEnumType = "Energy.Apparent.Export"
const MeasurandEnumTypeEnergyApparentImport MeasurandEnumType = "Energy.Apparent.Import"
const MeasurandEnumTypeEnergyApparentNet MeasurandEnumType = "Energy.Apparent.Net"
const MeasurandEnumTypeEnergyReactiveExportInterval MeasurandEnumType = "Energy.Reactive.Export.Interval"
const MeasurandEnumTypeEnergyReactiveExportRegister MeasurandEnumType = "Energy.Reactive.Export.Register"
const MeasurandEnumTypeEnergyReactiveImportInterval MeasurandEnumType = "Energy.Reactive.Import.Interval"
const MeasurandEnumTypeEnergyReactiveImportRegister MeasurandEnumType = "Energy.Reactive.Import.Register"
const MeasurandEnumTypeEnergyReactiveNet MeasurandEnumType = "Energy.Reactive.Net"
const MeasurandEnumTypeFrequency MeasurandEnumType = "Frequency"
const MeasurandEnumTypePowerActiveExport MeasurandEnumType = "Power.Active.Export"
const MeasurandEnumTypePowerActiveImport MeasurandEnumType = "Power.Active.Import"
const MeasurandEnumTypePowerFactor MeasurandEnumType = "Power.Factor"
const MeasurandEnumTypePowerOffered MeasurandEnumType = "Power.Offered"
const MeasurandEnumTypePowerReactiveExport MeasurandEnumType = "Power.Reactive.Export"
const MeasurandEnumTypePowerReactiveImport MeasurandEnumType = "Power.Reactive.Import"
const MeasurandEnumTypeSoC MeasurandEnumType = "SoC"
const MeasurandEnumTypeVoltage MeasurandEnumType = "Voltage"

type MeasurandEnumType_1 string

const MeasurandEnumType_1_CurrentExport MeasurandEnumType_1 = "Current.Export"
const MeasurandEnumType_1_CurrentImport MeasurandEnumType_1 = "Current.Import"
const MeasurandEnumType_1_CurrentOffered MeasurandEnumType_1 = "Current.Offered"
const MeasurandEnumType_1_EnergyActiveExportInterval MeasurandEnumType_1 = "Energy.Active.Export.Interval"
const MeasurandEnumType_1_EnergyActiveExportRegister MeasurandEnumType_1 = "Energy.Active.Export.Register"
const MeasurandEnumType_1_EnergyActiveImportInterval MeasurandEnumType_1 = "Energy.Active.Import.Interval"
const MeasurandEnumType_1_EnergyActiveImportRegister MeasurandEnumType_1 = "Energy.Active.Import.Register"
const MeasurandEnumType_1_EnergyActiveNet MeasurandEnumType_1 = "Energy.Active.Net"
const MeasurandEnumType_1_EnergyApparentExport MeasurandEnumType_1 = "Energy.Apparent.Export"
const MeasurandEnumType_1_EnergyApparentImport MeasurandEnumType_1 = "Energy.Apparent.Import"
const MeasurandEnumType_1_EnergyApparentNet MeasurandEnumType_1 = "Energy.Apparent.Net"
const MeasurandEnumType_1_EnergyReactiveExportInterval MeasurandEnumType_1 = "Energy.Reactive.Export.Interval"
const MeasurandEnumType_1_EnergyReactiveExportRegister MeasurandEnumType_1 = "Energy.Reactive.Export.Register"
const MeasurandEnumType_1_EnergyReactiveImportInterval MeasurandEnumType_1 = "Energy.Reactive.Import.Interval"
const MeasurandEnumType_1_EnergyReactiveImportRegister MeasurandEnumType_1 = "Energy.Reactive.Import.Register"
const MeasurandEnumType_1_EnergyReactiveNet MeasurandEnumType_1 = "Energy.Reactive.Net"
const MeasurandEnumType_1_Frequency MeasurandEnumType_1 = "Frequency"
const MeasurandEnumType_1_PowerActiveExport MeasurandEnumType_1 = "Power.Active.Export"
const MeasurandEnumType_1_PowerActiveImport MeasurandEnumType_1 = "Power.Active.Import"
const MeasurandEnumType_1_PowerFactor MeasurandEnumType_1 = "Power.Factor"
const MeasurandEnumType_1_PowerOffered MeasurandEnumType_1 = "Power.Offered"
const MeasurandEnumType_1_PowerReactiveExport MeasurandEnumType_1 = "Power.Reactive.Export"
const MeasurandEnumType_1_PowerReactiveImport MeasurandEnumType_1 = "Power.Reactive.Import"
const MeasurandEnumType_1_SoC MeasurandEnumType_1 = "SoC"
const MeasurandEnumType_1_Voltage MeasurandEnumType_1 = "Voltage"

type MeasurandEnumType_2 string

const MeasurandEnumType_2_CurrentExport MeasurandEnumType_2 = "Current.Export"
const MeasurandEnumType_2_CurrentImport MeasurandEnumType_2 = "Current.Import"
const MeasurandEnumType_2_CurrentOffered MeasurandEnumType_2 = "Current.Offered"
const MeasurandEnumType_2_EnergyActiveExportInterval MeasurandEnumType_2 = "Energy.Active.Export.Interval"
const MeasurandEnumType_2_EnergyActiveExportRegister MeasurandEnumType_2 = "Energy.Active.Export.Register"
const MeasurandEnumType_2_EnergyActiveImportInterval MeasurandEnumType_2 = "Energy.Active.Import.Interval"
const MeasurandEnumType_2_EnergyActiveImportRegister MeasurandEnumType_2 = "Energy.Active.Import.Register"
const MeasurandEnumType_2_EnergyActiveNet MeasurandEnumType_2 = "Energy.Active.Net"
const MeasurandEnumType_2_EnergyApparentExport MeasurandEnumType_2 = "Energy.Apparent.Export"
const MeasurandEnumType_2_EnergyApparentImport MeasurandEnumType_2 = "Energy.Apparent.Import"
const MeasurandEnumType_2_EnergyApparentNet MeasurandEnumType_2 = "Energy.Apparent.Net"
const MeasurandEnumType_2_EnergyReactiveExportInterval MeasurandEnumType_2 = "Energy.Reactive.Export.Interval"
const MeasurandEnumType_2_EnergyReactiveExportRegister MeasurandEnumType_2 = "Energy.Reactive.Export.Register"
const MeasurandEnumType_2_EnergyReactiveImportInterval MeasurandEnumType_2 = "Energy.Reactive.Import.Interval"
const MeasurandEnumType_2_EnergyReactiveImportRegister MeasurandEnumType_2 = "Energy.Reactive.Import.Register"
const MeasurandEnumType_2_EnergyReactiveNet MeasurandEnumType_2 = "Energy.Reactive.Net"
const MeasurandEnumType_2_Frequency MeasurandEnumType_2 = "Frequency"
const MeasurandEnumType_2_PowerActiveExport MeasurandEnumType_2 = "Power.Active.Export"
const MeasurandEnumType_2_PowerActiveImport MeasurandEnumType_2 = "Power.Active.Import"
const MeasurandEnumType_2_PowerFactor MeasurandEnumType_2 = "Power.Factor"
const MeasurandEnumType_2_PowerOffered MeasurandEnumType_2 = "Power.Offered"
const MeasurandEnumType_2_PowerReactiveExport MeasurandEnumType_2 = "Power.Reactive.Export"
const MeasurandEnumType_2_PowerReactiveImport MeasurandEnumType_2 = "Power.Reactive.Import"
const MeasurandEnumType_2_SoC MeasurandEnumType_2 = "SoC"
const MeasurandEnumType_2_Voltage MeasurandEnumType_2 = "Voltage"

type MeasurandEnumType_3 string

const MeasurandEnumType_3_CurrentExport MeasurandEnumType_3 = "Current.Export"
const MeasurandEnumType_3_CurrentImport MeasurandEnumType_3 = "Current.Import"
const MeasurandEnumType_3_CurrentOffered MeasurandEnumType_3 = "Current.Offered"
const MeasurandEnumType_3_EnergyActiveExportInterval MeasurandEnumType_3 = "Energy.Active.Export.Interval"
const MeasurandEnumType_3_EnergyActiveExportRegister MeasurandEnumType_3 = "Energy.Active.Export.Register"
const MeasurandEnumType_3_EnergyActiveImportInterval MeasurandEnumType_3 = "Energy.Active.Import.Interval"
const MeasurandEnumType_3_EnergyActiveImportRegister MeasurandEnumType_3 = "Energy.Active.Import.Register"
const MeasurandEnumType_3_EnergyActiveNet MeasurandEnumType_3 = "Energy.Active.Net"
const MeasurandEnumType_3_EnergyApparentExport MeasurandEnumType_3 = "Energy.Apparent.Export"
const MeasurandEnumType_3_EnergyApparentImport MeasurandEnumType_3 = "Energy.Apparent.Import"
const MeasurandEnumType_3_EnergyApparentNet MeasurandEnumType_3 = "Energy.Apparent.Net"
const MeasurandEnumType_3_EnergyReactiveExportInterval MeasurandEnumType_3 = "Energy.Reactive.Export.Interval"
const MeasurandEnumType_3_EnergyReactiveExportRegister MeasurandEnumType_3 = "Energy.Reactive.Export.Register"
const MeasurandEnumType_3_EnergyReactiveImportInterval MeasurandEnumType_3 = "Energy.Reactive.Import.Interval"
const MeasurandEnumType_3_EnergyReactiveImportRegister MeasurandEnumType_3 = "Energy.Reactive.Import.Register"
const MeasurandEnumType_3_EnergyReactiveNet MeasurandEnumType_3 = "Energy.Reactive.Net"
const MeasurandEnumType_3_Frequency MeasurandEnumType_3 = "Frequency"
const MeasurandEnumType_3_PowerActiveExport MeasurandEnumType_3 = "Power.Active.Export"
const MeasurandEnumType_3_PowerActiveImport MeasurandEnumType_3 = "Power.Active.Import"
const MeasurandEnumType_3_PowerFactor MeasurandEnumType_3 = "Power.Factor"
const MeasurandEnumType_3_PowerOffered MeasurandEnumType_3 = "Power.Offered"
const MeasurandEnumType_3_PowerReactiveExport MeasurandEnumType_3 = "Power.Reactive.Export"
const MeasurandEnumType_3_PowerReactiveImport MeasurandEnumType_3 = "Power.Reactive.Import"
const MeasurandEnumType_3_SoC MeasurandEnumType_3 = "SoC"
const MeasurandEnumType_3_Voltage MeasurandEnumType_3 = "Voltage"

// Message_ Content
// urn:x-enexis:ecdm:uid:2:234490
// Contains message details, for a message to be displayed on a Charging Station.
//
//
type MessageContentType struct {
	// Message_ Content. Content. Message
	// urn:x-enexis:ecdm:uid:1:570852
	// Message contents.
	//
	//
	Content string `json:"content"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_1 `json:"customData,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format MessageFormatEnumType `json:"format"`

	// Message_ Content. Language. Language_ Code
	// urn:x-enexis:ecdm:uid:1:570849
	// Message language identifier. Contains a language code as defined in
	// &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language *string `json:"language,omitempty"`
}

// Message_ Content
// urn:x-enexis:ecdm:uid:2:234490
// Contains message details, for a message to be displayed on a Charging Station.
//
//
type MessageContentType_1 struct {
	// Message_ Content. Content. Message
	// urn:x-enexis:ecdm:uid:1:570852
	// Message contents.
	//
	//
	Content string `json:"content"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_68 `json:"customData,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format MessageFormatEnumType_2 `json:"format"`

	// Message_ Content. Language. Language_ Code
	// urn:x-enexis:ecdm:uid:1:570849
	// Message language identifier. Contains a language code as defined in
	// &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language *string `json:"language,omitempty"`
}

// Message_ Content
// urn:x-enexis:ecdm:uid:2:234490
// Contains message details, for a message to be displayed on a Charging Station.
//
//
type MessageContentType_2 struct {
	// Message_ Content. Content. Message
	// urn:x-enexis:ecdm:uid:1:570852
	// Message contents.
	//
	//
	Content string `json:"content"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_98 `json:"customData,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format MessageFormatEnumType_4 `json:"format"`

	// Message_ Content. Language. Language_ Code
	// urn:x-enexis:ecdm:uid:1:570849
	// Message language identifier. Contains a language code as defined in
	// &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language *string `json:"language,omitempty"`
}

// Message_ Content
// urn:x-enexis:ecdm:uid:2:234490
// Contains message details, for a message to be displayed on a Charging Station.
//
//
type MessageContentType_3 struct {
	// Message_ Content. Content. Message
	// urn:x-enexis:ecdm:uid:1:570852
	// Message contents.
	//
	//
	Content string `json:"content"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_102 `json:"customData,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format MessageFormatEnumType_6 `json:"format"`

	// Message_ Content. Language. Language_ Code
	// urn:x-enexis:ecdm:uid:1:570849
	// Message language identifier. Contains a language code as defined in
	// &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language *string `json:"language,omitempty"`
}

// Message_ Content
// urn:x-enexis:ecdm:uid:2:234490
// Contains message details, for a message to be displayed on a Charging Station.
//
//
type MessageContentType_4 struct {
	// Message_ Content. Content. Message
	// urn:x-enexis:ecdm:uid:1:570852
	// Message contents.
	//
	//
	Content string `json:"content"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_119 `json:"customData,omitempty"`

	// Format corresponds to the JSON schema field "format".
	Format MessageFormatEnumType_8 `json:"format"`

	// Message_ Content. Language. Language_ Code
	// urn:x-enexis:ecdm:uid:1:570849
	// Message language identifier. Contains a language code as defined in
	// &lt;&lt;ref-RFC5646,[RFC5646]&gt;&gt;.
	//
	Language *string `json:"language,omitempty"`
}

type MessageFormatEnumType string

const MessageFormatEnumTypeASCII MessageFormatEnumType = "ASCII"
const MessageFormatEnumTypeHTML MessageFormatEnumType = "HTML"
const MessageFormatEnumTypeURI MessageFormatEnumType = "URI"
const MessageFormatEnumTypeUTF8 MessageFormatEnumType = "UTF8"

type MessageFormatEnumType_1 string

const MessageFormatEnumType_1_ASCII MessageFormatEnumType_1 = "ASCII"
const MessageFormatEnumType_1_HTML MessageFormatEnumType_1 = "HTML"
const MessageFormatEnumType_1_URI MessageFormatEnumType_1 = "URI"
const MessageFormatEnumType_1_UTF8 MessageFormatEnumType_1 = "UTF8"

type MessageFormatEnumType_2 string

const MessageFormatEnumType_2_ASCII MessageFormatEnumType_2 = "ASCII"
const MessageFormatEnumType_2_HTML MessageFormatEnumType_2 = "HTML"
const MessageFormatEnumType_2_URI MessageFormatEnumType_2 = "URI"
const MessageFormatEnumType_2_UTF8 MessageFormatEnumType_2 = "UTF8"

type MessageFormatEnumType_3 string

const MessageFormatEnumType_3_ASCII MessageFormatEnumType_3 = "ASCII"
const MessageFormatEnumType_3_HTML MessageFormatEnumType_3 = "HTML"
const MessageFormatEnumType_3_URI MessageFormatEnumType_3 = "URI"
const MessageFormatEnumType_3_UTF8 MessageFormatEnumType_3 = "UTF8"

type MessageFormatEnumType_4 string

const MessageFormatEnumType_4_ASCII MessageFormatEnumType_4 = "ASCII"
const MessageFormatEnumType_4_HTML MessageFormatEnumType_4 = "HTML"
const MessageFormatEnumType_4_URI MessageFormatEnumType_4 = "URI"
const MessageFormatEnumType_4_UTF8 MessageFormatEnumType_4 = "UTF8"

type MessageFormatEnumType_5 string

const MessageFormatEnumType_5_ASCII MessageFormatEnumType_5 = "ASCII"
const MessageFormatEnumType_5_HTML MessageFormatEnumType_5 = "HTML"
const MessageFormatEnumType_5_URI MessageFormatEnumType_5 = "URI"
const MessageFormatEnumType_5_UTF8 MessageFormatEnumType_5 = "UTF8"

type MessageFormatEnumType_6 string

const MessageFormatEnumType_6_ASCII MessageFormatEnumType_6 = "ASCII"
const MessageFormatEnumType_6_HTML MessageFormatEnumType_6 = "HTML"
const MessageFormatEnumType_6_URI MessageFormatEnumType_6 = "URI"
const MessageFormatEnumType_6_UTF8 MessageFormatEnumType_6 = "UTF8"

type MessageFormatEnumType_7 string

const MessageFormatEnumType_7_ASCII MessageFormatEnumType_7 = "ASCII"
const MessageFormatEnumType_7_HTML MessageFormatEnumType_7 = "HTML"
const MessageFormatEnumType_7_URI MessageFormatEnumType_7 = "URI"
const MessageFormatEnumType_7_UTF8 MessageFormatEnumType_7 = "UTF8"

type MessageFormatEnumType_8 string

const MessageFormatEnumType_8_ASCII MessageFormatEnumType_8 = "ASCII"
const MessageFormatEnumType_8_HTML MessageFormatEnumType_8 = "HTML"
const MessageFormatEnumType_8_URI MessageFormatEnumType_8 = "URI"
const MessageFormatEnumType_8_UTF8 MessageFormatEnumType_8 = "UTF8"

type MessageFormatEnumType_9 string

const MessageFormatEnumType_9_ASCII MessageFormatEnumType_9 = "ASCII"
const MessageFormatEnumType_9_HTML MessageFormatEnumType_9 = "HTML"
const MessageFormatEnumType_9_URI MessageFormatEnumType_9 = "URI"
const MessageFormatEnumType_9_UTF8 MessageFormatEnumType_9 = "UTF8"

// Message_ Info
// urn:x-enexis:ecdm:uid:2:233264
// Contains message details, for a message to be displayed on a Charging Station.
//
type MessageInfoType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_68 `json:"customData,omitempty"`

	// Display corresponds to the JSON schema field "display".
	Display *ComponentType_4 `json:"display,omitempty"`

	// Message_ Info. End. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569257
	// Until what date-time should this message be shown, after this date/time this
	// message SHALL be removed.
	//
	EndDateTime *string `json:"endDateTime,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Master resource identifier, unique within an exchange context. It is defined
	// within the OCPP context as a positive Integer value (greater or equal to zero).
	//
	Id int `json:"id"`

	// Message corresponds to the JSON schema field "message".
	Message MessageContentType_1 `json:"message"`

	// Priority corresponds to the JSON schema field "priority".
	Priority MessagePriorityEnumType_2 `json:"priority"`

	// Message_ Info. Start. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569256
	// From what date-time should this message be shown. If omitted: directly.
	//
	StartDateTime *string `json:"startDateTime,omitempty"`

	// State corresponds to the JSON schema field "state".
	State *MessageStateEnumType_2 `json:"state,omitempty"`

	// During which transaction shall this message be shown.
	// Message SHALL be removed by the Charging Station after transaction has
	// ended.
	//
	TransactionId *string `json:"transactionId,omitempty"`
}

// Message_ Info
// urn:x-enexis:ecdm:uid:2:233264
// Contains message details, for a message to be displayed on a Charging Station.
//
type MessageInfoType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_102 `json:"customData,omitempty"`

	// Display corresponds to the JSON schema field "display".
	Display *ComponentType_8 `json:"display,omitempty"`

	// Message_ Info. End. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569257
	// Until what date-time should this message be shown, after this date/time this
	// message SHALL be removed.
	//
	EndDateTime *string `json:"endDateTime,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// Master resource identifier, unique within an exchange context. It is defined
	// within the OCPP context as a positive Integer value (greater or equal to zero).
	//
	Id int `json:"id"`

	// Message corresponds to the JSON schema field "message".
	Message MessageContentType_3 `json:"message"`

	// Priority corresponds to the JSON schema field "priority".
	Priority MessagePriorityEnumType_4 `json:"priority"`

	// Message_ Info. Start. Date_ Time
	// urn:x-enexis:ecdm:uid:1:569256
	// From what date-time should this message be shown. If omitted: directly.
	//
	StartDateTime *string `json:"startDateTime,omitempty"`

	// State corresponds to the JSON schema field "state".
	State *MessageStateEnumType_4 `json:"state,omitempty"`

	// During which transaction shall this message be shown.
	// Message SHALL be removed by the Charging Station after transaction has
	// ended.
	//
	TransactionId *string `json:"transactionId,omitempty"`
}

type MessagePriorityEnumType string

const MessagePriorityEnumTypeAlwaysFront MessagePriorityEnumType = "AlwaysFront"
const MessagePriorityEnumTypeInFront MessagePriorityEnumType = "InFront"
const MessagePriorityEnumTypeNormalCycle MessagePriorityEnumType = "NormalCycle"

type MessagePriorityEnumType_1 string

const MessagePriorityEnumType_1_AlwaysFront MessagePriorityEnumType_1 = "AlwaysFront"
const MessagePriorityEnumType_1_InFront MessagePriorityEnumType_1 = "InFront"
const MessagePriorityEnumType_1_NormalCycle MessagePriorityEnumType_1 = "NormalCycle"

type MessagePriorityEnumType_2 string

const MessagePriorityEnumType_2_AlwaysFront MessagePriorityEnumType_2 = "AlwaysFront"
const MessagePriorityEnumType_2_InFront MessagePriorityEnumType_2 = "InFront"
const MessagePriorityEnumType_2_NormalCycle MessagePriorityEnumType_2 = "NormalCycle"

type MessagePriorityEnumType_3 string

const MessagePriorityEnumType_3_AlwaysFront MessagePriorityEnumType_3 = "AlwaysFront"
const MessagePriorityEnumType_3_InFront MessagePriorityEnumType_3 = "InFront"
const MessagePriorityEnumType_3_NormalCycle MessagePriorityEnumType_3 = "NormalCycle"

type MessagePriorityEnumType_4 string

const MessagePriorityEnumType_4_AlwaysFront MessagePriorityEnumType_4 = "AlwaysFront"
const MessagePriorityEnumType_4_InFront MessagePriorityEnumType_4 = "InFront"
const MessagePriorityEnumType_4_NormalCycle MessagePriorityEnumType_4 = "NormalCycle"

type MessagePriorityEnumType_5 string

const MessagePriorityEnumType_5_AlwaysFront MessagePriorityEnumType_5 = "AlwaysFront"
const MessagePriorityEnumType_5_InFront MessagePriorityEnumType_5 = "InFront"
const MessagePriorityEnumType_5_NormalCycle MessagePriorityEnumType_5 = "NormalCycle"

type MessageStateEnumType string

const MessageStateEnumTypeCharging MessageStateEnumType = "Charging"
const MessageStateEnumTypeFaulted MessageStateEnumType = "Faulted"
const MessageStateEnumTypeIdle MessageStateEnumType = "Idle"
const MessageStateEnumTypeUnavailable MessageStateEnumType = "Unavailable"

type MessageStateEnumType_1 string

const MessageStateEnumType_1_Charging MessageStateEnumType_1 = "Charging"
const MessageStateEnumType_1_Faulted MessageStateEnumType_1 = "Faulted"
const MessageStateEnumType_1_Idle MessageStateEnumType_1 = "Idle"
const MessageStateEnumType_1_Unavailable MessageStateEnumType_1 = "Unavailable"

type MessageStateEnumType_2 string

const MessageStateEnumType_2_Charging MessageStateEnumType_2 = "Charging"
const MessageStateEnumType_2_Faulted MessageStateEnumType_2 = "Faulted"
const MessageStateEnumType_2_Idle MessageStateEnumType_2 = "Idle"
const MessageStateEnumType_2_Unavailable MessageStateEnumType_2 = "Unavailable"

type MessageStateEnumType_3 string

const MessageStateEnumType_3_Charging MessageStateEnumType_3 = "Charging"
const MessageStateEnumType_3_Faulted MessageStateEnumType_3 = "Faulted"
const MessageStateEnumType_3_Idle MessageStateEnumType_3 = "Idle"
const MessageStateEnumType_3_Unavailable MessageStateEnumType_3 = "Unavailable"

type MessageStateEnumType_4 string

const MessageStateEnumType_4_Charging MessageStateEnumType_4 = "Charging"
const MessageStateEnumType_4_Faulted MessageStateEnumType_4 = "Faulted"
const MessageStateEnumType_4_Idle MessageStateEnumType_4 = "Idle"
const MessageStateEnumType_4_Unavailable MessageStateEnumType_4 = "Unavailable"

type MessageStateEnumType_5 string

const MessageStateEnumType_5_Charging MessageStateEnumType_5 = "Charging"
const MessageStateEnumType_5_Faulted MessageStateEnumType_5 = "Faulted"
const MessageStateEnumType_5_Idle MessageStateEnumType_5 = "Idle"
const MessageStateEnumType_5_Unavailable MessageStateEnumType_5 = "Unavailable"

type MessageTriggerEnumType string

const MessageTriggerEnumTypeBootNotification MessageTriggerEnumType = "BootNotification"
const MessageTriggerEnumTypeFirmwareStatusNotification MessageTriggerEnumType = "FirmwareStatusNotification"
const MessageTriggerEnumTypeHeartbeat MessageTriggerEnumType = "Heartbeat"
const MessageTriggerEnumTypeLogStatusNotification MessageTriggerEnumType = "LogStatusNotification"
const MessageTriggerEnumTypeMeterValues MessageTriggerEnumType = "MeterValues"
const MessageTriggerEnumTypePublishFirmwareStatusNotification MessageTriggerEnumType = "PublishFirmwareStatusNotification"
const MessageTriggerEnumTypeSignChargingStationCertificate MessageTriggerEnumType = "SignChargingStationCertificate"
const MessageTriggerEnumTypeSignCombinedCertificate MessageTriggerEnumType = "SignCombinedCertificate"
const MessageTriggerEnumTypeSignV2GCertificate MessageTriggerEnumType = "SignV2GCertificate"
const MessageTriggerEnumTypeStatusNotification MessageTriggerEnumType = "StatusNotification"
const MessageTriggerEnumTypeTransactionEvent MessageTriggerEnumType = "TransactionEvent"

type MessageTriggerEnumType_1 string

const MessageTriggerEnumType_1_BootNotification MessageTriggerEnumType_1 = "BootNotification"
const MessageTriggerEnumType_1_FirmwareStatusNotification MessageTriggerEnumType_1 = "FirmwareStatusNotification"
const MessageTriggerEnumType_1_Heartbeat MessageTriggerEnumType_1 = "Heartbeat"
const MessageTriggerEnumType_1_LogStatusNotification MessageTriggerEnumType_1 = "LogStatusNotification"
const MessageTriggerEnumType_1_MeterValues MessageTriggerEnumType_1 = "MeterValues"
const MessageTriggerEnumType_1_PublishFirmwareStatusNotification MessageTriggerEnumType_1 = "PublishFirmwareStatusNotification"
const MessageTriggerEnumType_1_SignChargingStationCertificate MessageTriggerEnumType_1 = "SignChargingStationCertificate"
const MessageTriggerEnumType_1_SignCombinedCertificate MessageTriggerEnumType_1 = "SignCombinedCertificate"
const MessageTriggerEnumType_1_SignV2GCertificate MessageTriggerEnumType_1 = "SignV2GCertificate"
const MessageTriggerEnumType_1_StatusNotification MessageTriggerEnumType_1 = "StatusNotification"
const MessageTriggerEnumType_1_TransactionEvent MessageTriggerEnumType_1 = "TransactionEvent"

// Meter_ Value
// urn:x-oca:ocpp:uid:2:233265
// Collection of one or more sampled values in MeterValuesRequest and
// TransactionEvent. All sampled values in a MeterValue are sampled at the same
// point in time.
//
type MeterValueType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_62 `json:"customData,omitempty"`

	// SampledValue corresponds to the JSON schema field "sampledValue".
	SampledValue []SampledValueType `json:"sampledValue"`

	// Meter_ Value. Timestamp. Date_ Time
	// urn:x-oca:ocpp:uid:1:569259
	// Timestamp for measured value(s).
	//
	Timestamp string `json:"timestamp"`
}

// Meter_ Value
// urn:x-oca:ocpp:uid:2:233265
// Collection of one or more sampled values in MeterValuesRequest and
// TransactionEvent. All sampled values in a MeterValue are sampled at the same
// point in time.
//
type MeterValueType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// SampledValue corresponds to the JSON schema field "sampledValue".
	SampledValue []SampledValueType_1 `json:"sampledValue"`

	// Meter_ Value. Timestamp. Date_ Time
	// urn:x-oca:ocpp:uid:1:569259
	// Timestamp for measured value(s).
	//
	Timestamp string `json:"timestamp"`
}

// Request_ Body
// urn:x-enexis:ecdm:uid:2:234744
//
type MeterValuesRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_62 `json:"customData,omitempty"`

	// Request_ Body. EVSEID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:571101
	// This contains a number (&gt;0) designating an EVSE of the Charging Station.
	// ‘0’ (zero) is used to designate the main power meter.
	//
	EvseId int `json:"evseId"`

	// MeterValue corresponds to the JSON schema field "meterValue".
	MeterValue []MeterValueType `json:"meterValue"`
}

type MeterValuesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_63 `json:"customData,omitempty"`
}

// Wireless_ Communication_ Module
// urn:x-oca:ocpp:uid:2:233306
// Defines parameters required for initiating and maintaining wireless
// communication with other devices.
//
type ModemType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_2 `json:"customData,omitempty"`

	// Wireless_ Communication_ Module. ICCID. CI20_ Text
	// urn:x-oca:ocpp:uid:1:569327
	// This contains the ICCID of the modem’s SIM card.
	//
	Iccid *string `json:"iccid,omitempty"`

	// Wireless_ Communication_ Module. IMSI. CI20_ Text
	// urn:x-oca:ocpp:uid:1:569328
	// This contains the IMSI of the modem’s SIM card.
	//
	Imsi *string `json:"imsi,omitempty"`
}

type MonitorEnumType string

const MonitorEnumTypeDelta MonitorEnumType = "Delta"
const MonitorEnumTypeLowerThreshold MonitorEnumType = "LowerThreshold"
const MonitorEnumTypePeriodic MonitorEnumType = "Periodic"
const MonitorEnumTypePeriodicClockAligned MonitorEnumType = "PeriodicClockAligned"
const MonitorEnumTypeUpperThreshold MonitorEnumType = "UpperThreshold"

type MonitorEnumType_1 string

const MonitorEnumType_1_Delta MonitorEnumType_1 = "Delta"
const MonitorEnumType_1_LowerThreshold MonitorEnumType_1 = "LowerThreshold"
const MonitorEnumType_1_Periodic MonitorEnumType_1 = "Periodic"
const MonitorEnumType_1_PeriodicClockAligned MonitorEnumType_1 = "PeriodicClockAligned"
const MonitorEnumType_1_UpperThreshold MonitorEnumType_1 = "UpperThreshold"

type MonitorEnumType_2 string

const MonitorEnumType_2_Delta MonitorEnumType_2 = "Delta"
const MonitorEnumType_2_LowerThreshold MonitorEnumType_2 = "LowerThreshold"
const MonitorEnumType_2_Periodic MonitorEnumType_2 = "Periodic"
const MonitorEnumType_2_PeriodicClockAligned MonitorEnumType_2 = "PeriodicClockAligned"
const MonitorEnumType_2_UpperThreshold MonitorEnumType_2 = "UpperThreshold"

type MonitorEnumType_3 string

const MonitorEnumType_3_Delta MonitorEnumType_3 = "Delta"
const MonitorEnumType_3_LowerThreshold MonitorEnumType_3 = "LowerThreshold"
const MonitorEnumType_3_Periodic MonitorEnumType_3 = "Periodic"
const MonitorEnumType_3_PeriodicClockAligned MonitorEnumType_3 = "PeriodicClockAligned"
const MonitorEnumType_3_UpperThreshold MonitorEnumType_3 = "UpperThreshold"

type MonitorEnumType_4 string

const MonitorEnumType_4_Delta MonitorEnumType_4 = "Delta"
const MonitorEnumType_4_LowerThreshold MonitorEnumType_4 = "LowerThreshold"
const MonitorEnumType_4_Periodic MonitorEnumType_4 = "Periodic"
const MonitorEnumType_4_PeriodicClockAligned MonitorEnumType_4 = "PeriodicClockAligned"
const MonitorEnumType_4_UpperThreshold MonitorEnumType_4 = "UpperThreshold"

type MonitorEnumType_5 string

const MonitorEnumType_5_Delta MonitorEnumType_5 = "Delta"
const MonitorEnumType_5_LowerThreshold MonitorEnumType_5 = "LowerThreshold"
const MonitorEnumType_5_Periodic MonitorEnumType_5 = "Periodic"
const MonitorEnumType_5_PeriodicClockAligned MonitorEnumType_5 = "PeriodicClockAligned"
const MonitorEnumType_5_UpperThreshold MonitorEnumType_5 = "UpperThreshold"

type MonitoringBaseEnumType string

const MonitoringBaseEnumTypeAll MonitoringBaseEnumType = "All"
const MonitoringBaseEnumTypeFactoryDefault MonitoringBaseEnumType = "FactoryDefault"
const MonitoringBaseEnumTypeHardWiredOnly MonitoringBaseEnumType = "HardWiredOnly"

type MonitoringBaseEnumType_1 string

const MonitoringBaseEnumType_1_All MonitoringBaseEnumType_1 = "All"
const MonitoringBaseEnumType_1_FactoryDefault MonitoringBaseEnumType_1 = "FactoryDefault"
const MonitoringBaseEnumType_1_HardWiredOnly MonitoringBaseEnumType_1 = "HardWiredOnly"

type MonitoringCriterionEnumType string

const MonitoringCriterionEnumTypeDeltaMonitoring MonitoringCriterionEnumType = "DeltaMonitoring"
const MonitoringCriterionEnumTypePeriodicMonitoring MonitoringCriterionEnumType = "PeriodicMonitoring"
const MonitoringCriterionEnumTypeThresholdMonitoring MonitoringCriterionEnumType = "ThresholdMonitoring"

type MonitoringCriterionEnumType_1 string

const MonitoringCriterionEnumType_1_DeltaMonitoring MonitoringCriterionEnumType_1 = "DeltaMonitoring"
const MonitoringCriterionEnumType_1_PeriodicMonitoring MonitoringCriterionEnumType_1 = "PeriodicMonitoring"
const MonitoringCriterionEnumType_1_ThresholdMonitoring MonitoringCriterionEnumType_1 = "ThresholdMonitoring"

// Class to hold parameters of SetVariableMonitoring request.
//
type MonitoringDataType struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType_6 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_76 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_5 `json:"variable"`

	// VariableMonitoring corresponds to the JSON schema field "variableMonitoring".
	VariableMonitoring []VariableMonitoringType `json:"variableMonitoring"`
}

type MutabilityEnumType string

const MutabilityEnumTypeReadOnly MutabilityEnumType = "ReadOnly"
const MutabilityEnumTypeReadWrite MutabilityEnumType = "ReadWrite"
const MutabilityEnumTypeWriteOnly MutabilityEnumType = "WriteOnly"

type MutabilityEnumType_1 string

const MutabilityEnumType_1_ReadOnly MutabilityEnumType_1 = "ReadOnly"
const MutabilityEnumType_1_ReadWrite MutabilityEnumType_1 = "ReadWrite"
const MutabilityEnumType_1_WriteOnly MutabilityEnumType_1 = "WriteOnly"

// Communication_ Function
// urn:x-oca:ocpp:uid:2:233304
// The NetworkConnectionProfile defines the functional and technical parameters of
// a communication link.
//
type NetworkConnectionProfileType struct {
	// Apn corresponds to the JSON schema field "apn".
	Apn *APNType `json:"apn,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_108 `json:"customData,omitempty"`

	// Duration in seconds before a message send by the Charging Station via this
	// network connection times-out.
	// The best setting depends on the underlying network and response times of the
	// CSMS.
	// If you are looking for a some guideline: use 30 seconds as a starting point.
	//
	MessageTimeout int `json:"messageTimeout"`

	// Communication_ Function. OCPP_ Central_ System_ URL. URI
	// urn:x-oca:ocpp:uid:1:569357
	// URL of the CSMS(s) that this Charging Station  communicates with.
	//
	OcppCsmsUrl string `json:"ocppCsmsUrl"`

	// OcppInterface corresponds to the JSON schema field "ocppInterface".
	OcppInterface OCPPInterfaceEnumType `json:"ocppInterface"`

	// OcppTransport corresponds to the JSON schema field "ocppTransport".
	OcppTransport OCPPTransportEnumType `json:"ocppTransport"`

	// OcppVersion corresponds to the JSON schema field "ocppVersion".
	OcppVersion OCPPVersionEnumType `json:"ocppVersion"`

	// This field specifies the security profile used when connecting to the CSMS with
	// this NetworkConnectionProfile.
	//
	SecurityProfile int `json:"securityProfile"`

	// Vpn corresponds to the JSON schema field "vpn".
	Vpn *VPNType `json:"vpn,omitempty"`
}

type NotifyChargingLimitRequestJson struct {
	// ChargingLimit corresponds to the JSON schema field "chargingLimit".
	ChargingLimit ChargingLimitType `json:"chargingLimit"`

	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule []ChargingScheduleType `json:"chargingSchedule,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// The charging schedule contained in this notification applies to an EVSE. evseId
	// must be &gt; 0.
	//
	EvseId *int `json:"evseId,omitempty"`
}

type NotifyChargingLimitResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_65 `json:"customData,omitempty"`
}

type NotifyCustomerInformationRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_66 `json:"customData,omitempty"`

	// (Part of) the requested data. No format specified in which the data is
	// returned. Should be human readable.
	//
	Data string `json:"data"`

	//  Timestamp of the moment this message was generated at the Charging Station.
	//
	GeneratedAt string `json:"generatedAt"`

	// The Id of the request.
	//
	//
	RequestId int `json:"requestId"`

	// Sequence number of this message. First message starts at 0.
	//
	SeqNo int `json:"seqNo"`

	// “to be continued” indicator. Indicates whether another part of the
	// monitoringData follows in an upcoming notifyMonitoringReportRequest message.
	// Default value when omitted is false.
	//
	Tbc bool `json:"tbc,omitempty"`
}

type NotifyCustomerInformationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_67 `json:"customData,omitempty"`
}

type NotifyDisplayMessagesRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_68 `json:"customData,omitempty"`

	// MessageInfo corresponds to the JSON schema field "messageInfo".
	MessageInfo []MessageInfoType `json:"messageInfo,omitempty"`

	// The id of the
	// &lt;&lt;getdisplaymessagesrequest,GetDisplayMessagesRequest&gt;&gt; that
	// requested this message.
	//
	RequestId int `json:"requestId"`

	// "to be continued" indicator. Indicates whether another part of the report
	// follows in an upcoming NotifyDisplayMessagesRequest message. Default value when
	// omitted is false.
	//
	Tbc bool `json:"tbc,omitempty"`
}

type NotifyDisplayMessagesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_69 `json:"customData,omitempty"`
}

type NotifyEVChargingNeedsRequestJson struct {
	// ChargingNeeds corresponds to the JSON schema field "chargingNeeds".
	ChargingNeeds ChargingNeedsType `json:"chargingNeeds"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_70 `json:"customData,omitempty"`

	// Defines the EVSE and connector to which the EV is connected. EvseId may not be
	// 0.
	//
	EvseId int `json:"evseId"`

	// Contains the maximum schedule tuples the car supports per schedule.
	//
	MaxScheduleTuples *int `json:"maxScheduleTuples,omitempty"`
}

type NotifyEVChargingNeedsResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_71 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status NotifyEVChargingNeedsStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_23 `json:"statusInfo,omitempty"`
}

type NotifyEVChargingNeedsStatusEnumType string

const NotifyEVChargingNeedsStatusEnumTypeAccepted NotifyEVChargingNeedsStatusEnumType = "Accepted"
const NotifyEVChargingNeedsStatusEnumTypeProcessing NotifyEVChargingNeedsStatusEnumType = "Processing"
const NotifyEVChargingNeedsStatusEnumTypeRejected NotifyEVChargingNeedsStatusEnumType = "Rejected"

type NotifyEVChargingNeedsStatusEnumType_1 string

const NotifyEVChargingNeedsStatusEnumType_1_Accepted NotifyEVChargingNeedsStatusEnumType_1 = "Accepted"
const NotifyEVChargingNeedsStatusEnumType_1_Processing NotifyEVChargingNeedsStatusEnumType_1 = "Processing"
const NotifyEVChargingNeedsStatusEnumType_1_Rejected NotifyEVChargingNeedsStatusEnumType_1 = "Rejected"

type NotifyEVChargingScheduleRequestJson struct {
	// ChargingSchedule corresponds to the JSON schema field "chargingSchedule".
	ChargingSchedule ChargingScheduleType_1 `json:"chargingSchedule"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`

	// The charging schedule contained in this notification applies to an EVSE. EvseId
	// must be &gt; 0.
	//
	EvseId int `json:"evseId"`

	// Periods contained in the charging profile are relative to this point in time.
	//
	TimeBase string `json:"timeBase"`
}

type NotifyEVChargingScheduleResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_73 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericStatusEnumType_3 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_24 `json:"statusInfo,omitempty"`
}

type NotifyEventRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_74 `json:"customData,omitempty"`

	// EventData corresponds to the JSON schema field "eventData".
	EventData []EventDataType `json:"eventData"`

	// Timestamp of the moment this message was generated at the Charging Station.
	//
	GeneratedAt string `json:"generatedAt"`

	// Sequence number of this message. First message starts at 0.
	//
	SeqNo int `json:"seqNo"`

	// “to be continued” indicator. Indicates whether another part of the report
	// follows in an upcoming notifyEventRequest message. Default value when omitted
	// is false.
	//
	Tbc bool `json:"tbc,omitempty"`
}

type NotifyEventResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_75 `json:"customData,omitempty"`
}

type NotifyMonitoringReportRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_76 `json:"customData,omitempty"`

	// Timestamp of the moment this message was generated at the Charging Station.
	//
	GeneratedAt string `json:"generatedAt"`

	// Monitor corresponds to the JSON schema field "monitor".
	Monitor []MonitoringDataType `json:"monitor,omitempty"`

	// The id of the GetMonitoringRequest that requested this report.
	//
	//
	RequestId int `json:"requestId"`

	// Sequence number of this message. First message starts at 0.
	//
	SeqNo int `json:"seqNo"`

	// “to be continued” indicator. Indicates whether another part of the
	// monitoringData follows in an upcoming notifyMonitoringReportRequest message.
	// Default value when omitted is false.
	//
	Tbc bool `json:"tbc,omitempty"`
}

type NotifyMonitoringReportResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_77 `json:"customData,omitempty"`
}

type NotifyReportRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_78 `json:"customData,omitempty"`

	// Timestamp of the moment this message was generated at the Charging Station.
	//
	GeneratedAt string `json:"generatedAt"`

	// ReportData corresponds to the JSON schema field "reportData".
	ReportData []ReportDataType `json:"reportData,omitempty"`

	// The id of the GetReportRequest  or GetBaseReportRequest that requested this
	// report
	//
	RequestId int `json:"requestId"`

	// Sequence number of this message. First message starts at 0.
	//
	SeqNo int `json:"seqNo"`

	// “to be continued” indicator. Indicates whether another part of the report
	// follows in an upcoming notifyReportRequest message. Default value when omitted
	// is false.
	//
	//
	Tbc bool `json:"tbc,omitempty"`
}

type NotifyReportResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_79 `json:"customData,omitempty"`
}

type OCPPInterfaceEnumType string

const OCPPInterfaceEnumTypeWired0 OCPPInterfaceEnumType = "Wired0"
const OCPPInterfaceEnumTypeWired1 OCPPInterfaceEnumType = "Wired1"
const OCPPInterfaceEnumTypeWired2 OCPPInterfaceEnumType = "Wired2"
const OCPPInterfaceEnumTypeWired3 OCPPInterfaceEnumType = "Wired3"
const OCPPInterfaceEnumTypeWireless0 OCPPInterfaceEnumType = "Wireless0"
const OCPPInterfaceEnumTypeWireless1 OCPPInterfaceEnumType = "Wireless1"
const OCPPInterfaceEnumTypeWireless2 OCPPInterfaceEnumType = "Wireless2"
const OCPPInterfaceEnumTypeWireless3 OCPPInterfaceEnumType = "Wireless3"

type OCPPInterfaceEnumType_1 string

const OCPPInterfaceEnumType_1_Wired0 OCPPInterfaceEnumType_1 = "Wired0"
const OCPPInterfaceEnumType_1_Wired1 OCPPInterfaceEnumType_1 = "Wired1"
const OCPPInterfaceEnumType_1_Wired2 OCPPInterfaceEnumType_1 = "Wired2"
const OCPPInterfaceEnumType_1_Wired3 OCPPInterfaceEnumType_1 = "Wired3"
const OCPPInterfaceEnumType_1_Wireless0 OCPPInterfaceEnumType_1 = "Wireless0"
const OCPPInterfaceEnumType_1_Wireless1 OCPPInterfaceEnumType_1 = "Wireless1"
const OCPPInterfaceEnumType_1_Wireless2 OCPPInterfaceEnumType_1 = "Wireless2"
const OCPPInterfaceEnumType_1_Wireless3 OCPPInterfaceEnumType_1 = "Wireless3"

type OCPPTransportEnumType string

const OCPPTransportEnumTypeJSON OCPPTransportEnumType = "JSON"
const OCPPTransportEnumTypeSOAP OCPPTransportEnumType = "SOAP"

type OCPPTransportEnumType_1 string

const OCPPTransportEnumType_1_JSON OCPPTransportEnumType_1 = "JSON"
const OCPPTransportEnumType_1_SOAP OCPPTransportEnumType_1 = "SOAP"

type OCPPVersionEnumType string

const OCPPVersionEnumTypeOCPP12 OCPPVersionEnumType = "OCPP12"
const OCPPVersionEnumTypeOCPP15 OCPPVersionEnumType = "OCPP15"
const OCPPVersionEnumTypeOCPP16 OCPPVersionEnumType = "OCPP16"
const OCPPVersionEnumTypeOCPP20 OCPPVersionEnumType = "OCPP20"

type OCPPVersionEnumType_1 string

const OCPPVersionEnumType_1_OCPP12 OCPPVersionEnumType_1 = "OCPP12"
const OCPPVersionEnumType_1_OCPP15 OCPPVersionEnumType_1 = "OCPP15"
const OCPPVersionEnumType_1_OCPP16 OCPPVersionEnumType_1 = "OCPP16"
const OCPPVersionEnumType_1_OCPP20 OCPPVersionEnumType_1 = "OCPP20"

type OCSPRequestDataType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType `json:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType_1 `json:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash"`

	// This contains the responder URL (Case insensitive).
	//
	//
	ResponderURL string `json:"responderURL"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber"`
}

type OCSPRequestDataType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_34 `json:"customData,omitempty"`

	// HashAlgorithm corresponds to the JSON schema field "hashAlgorithm".
	HashAlgorithm HashAlgorithmEnumType_7 `json:"hashAlgorithm"`

	// Hashed value of the issuers public key
	//
	IssuerKeyHash string `json:"issuerKeyHash"`

	// Hashed value of the Issuer DN (Distinguished Name).
	//
	//
	IssuerNameHash string `json:"issuerNameHash"`

	// This contains the responder URL (Case insensitive).
	//
	//
	ResponderURL string `json:"responderURL"`

	// The serial number of the certificate.
	//
	SerialNumber string `json:"serialNumber"`
}

type OperationalStatusEnumType string

const OperationalStatusEnumTypeInoperative OperationalStatusEnumType = "Inoperative"
const OperationalStatusEnumTypeOperative OperationalStatusEnumType = "Operative"

type OperationalStatusEnumType_1 string

const OperationalStatusEnumType_1_Inoperative OperationalStatusEnumType_1 = "Inoperative"
const OperationalStatusEnumType_1_Operative OperationalStatusEnumType_1 = "Operative"

type PhaseEnumType string

const PhaseEnumTypeL1 PhaseEnumType = "L1"
const PhaseEnumTypeL1L2 PhaseEnumType = "L1-L2"
const PhaseEnumTypeL1N PhaseEnumType = "L1-N"
const PhaseEnumTypeL2 PhaseEnumType = "L2"
const PhaseEnumTypeL2L3 PhaseEnumType = "L2-L3"
const PhaseEnumTypeL2N PhaseEnumType = "L2-N"
const PhaseEnumTypeL3 PhaseEnumType = "L3"
const PhaseEnumTypeL3L1 PhaseEnumType = "L3-L1"
const PhaseEnumTypeL3N PhaseEnumType = "L3-N"
const PhaseEnumTypeN PhaseEnumType = "N"

type PhaseEnumType_1 string

const PhaseEnumType_1_L1 PhaseEnumType_1 = "L1"
const PhaseEnumType_1_L1L2 PhaseEnumType_1 = "L1-L2"
const PhaseEnumType_1_L1N PhaseEnumType_1 = "L1-N"
const PhaseEnumType_1_L2 PhaseEnumType_1 = "L2"
const PhaseEnumType_1_L2L3 PhaseEnumType_1 = "L2-L3"
const PhaseEnumType_1_L2N PhaseEnumType_1 = "L2-N"
const PhaseEnumType_1_L3 PhaseEnumType_1 = "L3"
const PhaseEnumType_1_L3L1 PhaseEnumType_1 = "L3-L1"
const PhaseEnumType_1_L3N PhaseEnumType_1 = "L3-N"
const PhaseEnumType_1_N PhaseEnumType_1 = "N"

type PhaseEnumType_2 string

const PhaseEnumType_2_L1 PhaseEnumType_2 = "L1"
const PhaseEnumType_2_L1L2 PhaseEnumType_2 = "L1-L2"
const PhaseEnumType_2_L1N PhaseEnumType_2 = "L1-N"
const PhaseEnumType_2_L2 PhaseEnumType_2 = "L2"
const PhaseEnumType_2_L2L3 PhaseEnumType_2 = "L2-L3"
const PhaseEnumType_2_L2N PhaseEnumType_2 = "L2-N"
const PhaseEnumType_2_L3 PhaseEnumType_2 = "L3"
const PhaseEnumType_2_L3L1 PhaseEnumType_2 = "L3-L1"
const PhaseEnumType_2_L3N PhaseEnumType_2 = "L3-N"
const PhaseEnumType_2_N PhaseEnumType_2 = "N"

type PhaseEnumType_3 string

const PhaseEnumType_3_L1 PhaseEnumType_3 = "L1"
const PhaseEnumType_3_L1L2 PhaseEnumType_3 = "L1-L2"
const PhaseEnumType_3_L1N PhaseEnumType_3 = "L1-N"
const PhaseEnumType_3_L2 PhaseEnumType_3 = "L2"
const PhaseEnumType_3_L2L3 PhaseEnumType_3 = "L2-L3"
const PhaseEnumType_3_L2N PhaseEnumType_3 = "L2-N"
const PhaseEnumType_3_L3 PhaseEnumType_3 = "L3"
const PhaseEnumType_3_L3L1 PhaseEnumType_3 = "L3-L1"
const PhaseEnumType_3_L3N PhaseEnumType_3 = "L3-N"
const PhaseEnumType_3_N PhaseEnumType_3 = "N"

type PublishFirmwareRequestJson struct {
	// The MD5 checksum over the entire firmware file as a hexadecimal string of
	// length 32.
	//
	Checksum string `json:"checksum"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_80 `json:"customData,omitempty"`

	// This contains a string containing a URI pointing to a
	// location from which to retrieve the firmware.
	//
	Location string `json:"location"`

	// The Id of the request.
	//
	RequestId int `json:"requestId"`

	// This specifies how many times Charging Station must try
	// to download the firmware before giving up. If this field is not
	// present, it is left to Charging Station to decide how many times it wants to
	// retry.
	//
	Retries *int `json:"retries,omitempty"`

	// The interval in seconds
	// after which a retry may be
	// attempted. If this field is not
	// present, it is left to Charging
	// Station to decide how long to wait
	// between attempts.
	//
	RetryInterval *int `json:"retryInterval,omitempty"`
}

type PublishFirmwareResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_81 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericStatusEnumType_5 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_25 `json:"statusInfo,omitempty"`
}

type PublishFirmwareStatusEnumType string

const PublishFirmwareStatusEnumTypeChecksumVerified PublishFirmwareStatusEnumType = "ChecksumVerified"
const PublishFirmwareStatusEnumTypeDownloadFailed PublishFirmwareStatusEnumType = "DownloadFailed"
const PublishFirmwareStatusEnumTypeDownloadPaused PublishFirmwareStatusEnumType = "DownloadPaused"
const PublishFirmwareStatusEnumTypeDownloadScheduled PublishFirmwareStatusEnumType = "DownloadScheduled"
const PublishFirmwareStatusEnumTypeDownloaded PublishFirmwareStatusEnumType = "Downloaded"
const PublishFirmwareStatusEnumTypeDownloading PublishFirmwareStatusEnumType = "Downloading"
const PublishFirmwareStatusEnumTypeIdle PublishFirmwareStatusEnumType = "Idle"
const PublishFirmwareStatusEnumTypeInvalidChecksum PublishFirmwareStatusEnumType = "InvalidChecksum"
const PublishFirmwareStatusEnumTypePublishFailed PublishFirmwareStatusEnumType = "PublishFailed"
const PublishFirmwareStatusEnumTypePublished PublishFirmwareStatusEnumType = "Published"

type PublishFirmwareStatusEnumType_1 string

const PublishFirmwareStatusEnumType_1_ChecksumVerified PublishFirmwareStatusEnumType_1 = "ChecksumVerified"
const PublishFirmwareStatusEnumType_1_DownloadFailed PublishFirmwareStatusEnumType_1 = "DownloadFailed"
const PublishFirmwareStatusEnumType_1_DownloadPaused PublishFirmwareStatusEnumType_1 = "DownloadPaused"
const PublishFirmwareStatusEnumType_1_DownloadScheduled PublishFirmwareStatusEnumType_1 = "DownloadScheduled"
const PublishFirmwareStatusEnumType_1_Downloaded PublishFirmwareStatusEnumType_1 = "Downloaded"
const PublishFirmwareStatusEnumType_1_Downloading PublishFirmwareStatusEnumType_1 = "Downloading"
const PublishFirmwareStatusEnumType_1_Idle PublishFirmwareStatusEnumType_1 = "Idle"
const PublishFirmwareStatusEnumType_1_InvalidChecksum PublishFirmwareStatusEnumType_1 = "InvalidChecksum"
const PublishFirmwareStatusEnumType_1_PublishFailed PublishFirmwareStatusEnumType_1 = "PublishFailed"
const PublishFirmwareStatusEnumType_1_Published PublishFirmwareStatusEnumType_1 = "Published"

type PublishFirmwareStatusNotificationRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_82 `json:"customData,omitempty"`

	// Required if status is Published. Can be multiple URI’s, if the Local
	// Controller supports e.g. HTTP, HTTPS, and FTP.
	//
	Location []string `json:"location,omitempty"`

	// The request id that was
	// provided in the
	// PublishFirmwareRequest which
	// triggered this action.
	//
	RequestId *int `json:"requestId,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status PublishFirmwareStatusEnumType_1 `json:"status"`
}

type PublishFirmwareStatusNotificationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_83 `json:"customData,omitempty"`
}

type ReadingContextEnumType string

const ReadingContextEnumTypeInterruptionBegin ReadingContextEnumType = "Interruption.Begin"
const ReadingContextEnumTypeInterruptionEnd ReadingContextEnumType = "Interruption.End"
const ReadingContextEnumTypeOther ReadingContextEnumType = "Other"
const ReadingContextEnumTypeSampleClock ReadingContextEnumType = "Sample.Clock"
const ReadingContextEnumTypeSamplePeriodic ReadingContextEnumType = "Sample.Periodic"
const ReadingContextEnumTypeTransactionBegin ReadingContextEnumType = "Transaction.Begin"
const ReadingContextEnumTypeTransactionEnd ReadingContextEnumType = "Transaction.End"
const ReadingContextEnumTypeTrigger ReadingContextEnumType = "Trigger"

type ReadingContextEnumType_1 string

const ReadingContextEnumType_1_InterruptionBegin ReadingContextEnumType_1 = "Interruption.Begin"
const ReadingContextEnumType_1_InterruptionEnd ReadingContextEnumType_1 = "Interruption.End"
const ReadingContextEnumType_1_Other ReadingContextEnumType_1 = "Other"
const ReadingContextEnumType_1_SampleClock ReadingContextEnumType_1 = "Sample.Clock"
const ReadingContextEnumType_1_SamplePeriodic ReadingContextEnumType_1 = "Sample.Periodic"
const ReadingContextEnumType_1_TransactionBegin ReadingContextEnumType_1 = "Transaction.Begin"
const ReadingContextEnumType_1_TransactionEnd ReadingContextEnumType_1 = "Transaction.End"
const ReadingContextEnumType_1_Trigger ReadingContextEnumType_1 = "Trigger"

type ReadingContextEnumType_2 string

const ReadingContextEnumType_2_InterruptionBegin ReadingContextEnumType_2 = "Interruption.Begin"
const ReadingContextEnumType_2_InterruptionEnd ReadingContextEnumType_2 = "Interruption.End"
const ReadingContextEnumType_2_Other ReadingContextEnumType_2 = "Other"
const ReadingContextEnumType_2_SampleClock ReadingContextEnumType_2 = "Sample.Clock"
const ReadingContextEnumType_2_SamplePeriodic ReadingContextEnumType_2 = "Sample.Periodic"
const ReadingContextEnumType_2_TransactionBegin ReadingContextEnumType_2 = "Transaction.Begin"
const ReadingContextEnumType_2_TransactionEnd ReadingContextEnumType_2 = "Transaction.End"
const ReadingContextEnumType_2_Trigger ReadingContextEnumType_2 = "Trigger"

type ReadingContextEnumType_3 string

const ReadingContextEnumType_3_InterruptionBegin ReadingContextEnumType_3 = "Interruption.Begin"
const ReadingContextEnumType_3_InterruptionEnd ReadingContextEnumType_3 = "Interruption.End"
const ReadingContextEnumType_3_Other ReadingContextEnumType_3 = "Other"
const ReadingContextEnumType_3_SampleClock ReadingContextEnumType_3 = "Sample.Clock"
const ReadingContextEnumType_3_SamplePeriodic ReadingContextEnumType_3 = "Sample.Periodic"
const ReadingContextEnumType_3_TransactionBegin ReadingContextEnumType_3 = "Transaction.Begin"
const ReadingContextEnumType_3_TransactionEnd ReadingContextEnumType_3 = "Transaction.End"
const ReadingContextEnumType_3_Trigger ReadingContextEnumType_3 = "Trigger"

type ReasonEnumType string

const ReasonEnumTypeDeAuthorized ReasonEnumType = "DeAuthorized"
const ReasonEnumTypeEVDisconnected ReasonEnumType = "EVDisconnected"
const ReasonEnumTypeEmergencyStop ReasonEnumType = "EmergencyStop"
const ReasonEnumTypeEnergyLimitReached ReasonEnumType = "EnergyLimitReached"
const ReasonEnumTypeGroundFault ReasonEnumType = "GroundFault"
const ReasonEnumTypeImmediateReset ReasonEnumType = "ImmediateReset"
const ReasonEnumTypeLocal ReasonEnumType = "Local"
const ReasonEnumTypeLocalOutOfCredit ReasonEnumType = "LocalOutOfCredit"
const ReasonEnumTypeMasterPass ReasonEnumType = "MasterPass"
const ReasonEnumTypeOther ReasonEnumType = "Other"
const ReasonEnumTypeOvercurrentFault ReasonEnumType = "OvercurrentFault"
const ReasonEnumTypePowerLoss ReasonEnumType = "PowerLoss"
const ReasonEnumTypePowerQuality ReasonEnumType = "PowerQuality"
const ReasonEnumTypeReboot ReasonEnumType = "Reboot"
const ReasonEnumTypeRemote ReasonEnumType = "Remote"
const ReasonEnumTypeSOCLimitReached ReasonEnumType = "SOCLimitReached"
const ReasonEnumTypeStoppedByEV ReasonEnumType = "StoppedByEV"
const ReasonEnumTypeTimeLimitReached ReasonEnumType = "TimeLimitReached"
const ReasonEnumTypeTimeout ReasonEnumType = "Timeout"

type ReasonEnumType_1 string

const ReasonEnumType_1_DeAuthorized ReasonEnumType_1 = "DeAuthorized"
const ReasonEnumType_1_EVDisconnected ReasonEnumType_1 = "EVDisconnected"
const ReasonEnumType_1_EmergencyStop ReasonEnumType_1 = "EmergencyStop"
const ReasonEnumType_1_EnergyLimitReached ReasonEnumType_1 = "EnergyLimitReached"
const ReasonEnumType_1_GroundFault ReasonEnumType_1 = "GroundFault"
const ReasonEnumType_1_ImmediateReset ReasonEnumType_1 = "ImmediateReset"
const ReasonEnumType_1_Local ReasonEnumType_1 = "Local"
const ReasonEnumType_1_LocalOutOfCredit ReasonEnumType_1 = "LocalOutOfCredit"
const ReasonEnumType_1_MasterPass ReasonEnumType_1 = "MasterPass"
const ReasonEnumType_1_Other ReasonEnumType_1 = "Other"
const ReasonEnumType_1_OvercurrentFault ReasonEnumType_1 = "OvercurrentFault"
const ReasonEnumType_1_PowerLoss ReasonEnumType_1 = "PowerLoss"
const ReasonEnumType_1_PowerQuality ReasonEnumType_1 = "PowerQuality"
const ReasonEnumType_1_Reboot ReasonEnumType_1 = "Reboot"
const ReasonEnumType_1_Remote ReasonEnumType_1 = "Remote"
const ReasonEnumType_1_SOCLimitReached ReasonEnumType_1 = "SOCLimitReached"
const ReasonEnumType_1_StoppedByEV ReasonEnumType_1 = "StoppedByEV"
const ReasonEnumType_1_TimeLimitReached ReasonEnumType_1 = "TimeLimitReached"
const ReasonEnumType_1_Timeout ReasonEnumType_1 = "Timeout"

type RecurrencyKindEnumType string

const RecurrencyKindEnumTypeDaily RecurrencyKindEnumType = "Daily"
const RecurrencyKindEnumTypeWeekly RecurrencyKindEnumType = "Weekly"

type RecurrencyKindEnumType_1 string

const RecurrencyKindEnumType_1_Daily RecurrencyKindEnumType_1 = "Daily"
const RecurrencyKindEnumType_1_Weekly RecurrencyKindEnumType_1 = "Weekly"

type RecurrencyKindEnumType_2 string

const RecurrencyKindEnumType_2_Daily RecurrencyKindEnumType_2 = "Daily"
const RecurrencyKindEnumType_2_Weekly RecurrencyKindEnumType_2 = "Weekly"

type RecurrencyKindEnumType_3 string

const RecurrencyKindEnumType_3_Daily RecurrencyKindEnumType_3 = "Daily"
const RecurrencyKindEnumType_3_Weekly RecurrencyKindEnumType_3 = "Weekly"

type RecurrencyKindEnumType_4 string

const RecurrencyKindEnumType_4_Daily RecurrencyKindEnumType_4 = "Daily"
const RecurrencyKindEnumType_4_Weekly RecurrencyKindEnumType_4 = "Weekly"

type RecurrencyKindEnumType_5 string

const RecurrencyKindEnumType_5_Daily RecurrencyKindEnumType_5 = "Daily"
const RecurrencyKindEnumType_5_Weekly RecurrencyKindEnumType_5 = "Weekly"

type RegistrationStatusEnumType string

const RegistrationStatusEnumTypeAccepted RegistrationStatusEnumType = "Accepted"
const RegistrationStatusEnumTypePending RegistrationStatusEnumType = "Pending"
const RegistrationStatusEnumTypeRejected RegistrationStatusEnumType = "Rejected"

type RegistrationStatusEnumType_1 string

const RegistrationStatusEnumType_1_Accepted RegistrationStatusEnumType_1 = "Accepted"
const RegistrationStatusEnumType_1_Pending RegistrationStatusEnumType_1 = "Pending"
const RegistrationStatusEnumType_1_Rejected RegistrationStatusEnumType_1 = "Rejected"

// Relative_ Timer_ Interval
// urn:x-oca:ocpp:uid:2:233270
//
type RelativeTimeIntervalType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// Relative_ Timer_ Interval. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569280
	// Duration of the interval, in seconds.
	//
	Duration *int `json:"duration,omitempty"`

	// Relative_ Timer_ Interval. Start. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569279
	// Start of the interval, in seconds from NOW.
	//
	Start int `json:"start"`
}

// Relative_ Timer_ Interval
// urn:x-oca:ocpp:uid:2:233270
//
type RelativeTimeIntervalType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`

	// Relative_ Timer_ Interval. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569280
	// Duration of the interval, in seconds.
	//
	Duration *int `json:"duration,omitempty"`

	// Relative_ Timer_ Interval. Start. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569279
	// Start of the interval, in seconds from NOW.
	//
	Start int `json:"start"`
}

// Relative_ Timer_ Interval
// urn:x-oca:ocpp:uid:2:233270
//
type RelativeTimeIntervalType_2 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// Relative_ Timer_ Interval. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569280
	// Duration of the interval, in seconds.
	//
	Duration *int `json:"duration,omitempty"`

	// Relative_ Timer_ Interval. Start. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569279
	// Start of the interval, in seconds from NOW.
	//
	Start int `json:"start"`
}

// Relative_ Timer_ Interval
// urn:x-oca:ocpp:uid:2:233270
//
type RelativeTimeIntervalType_3 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Relative_ Timer_ Interval. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569280
	// Duration of the interval, in seconds.
	//
	Duration *int `json:"duration,omitempty"`

	// Relative_ Timer_ Interval. Start. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569279
	// Start of the interval, in seconds from NOW.
	//
	Start int `json:"start"`
}

// Relative_ Timer_ Interval
// urn:x-oca:ocpp:uid:2:233270
//
type RelativeTimeIntervalType_4 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// Relative_ Timer_ Interval. Duration. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569280
	// Duration of the interval, in seconds.
	//
	Duration *int `json:"duration,omitempty"`

	// Relative_ Timer_ Interval. Start. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569279
	// Start of the interval, in seconds from NOW.
	//
	Start int `json:"start"`
}

type ReportBaseEnumType string

const ReportBaseEnumTypeConfigurationInventory ReportBaseEnumType = "ConfigurationInventory"
const ReportBaseEnumTypeFullInventory ReportBaseEnumType = "FullInventory"
const ReportBaseEnumTypeSummaryInventory ReportBaseEnumType = "SummaryInventory"

type ReportBaseEnumType_1 string

const ReportBaseEnumType_1_ConfigurationInventory ReportBaseEnumType_1 = "ConfigurationInventory"
const ReportBaseEnumType_1_FullInventory ReportBaseEnumType_1 = "FullInventory"
const ReportBaseEnumType_1_SummaryInventory ReportBaseEnumType_1 = "SummaryInventory"

type ReportChargingProfilesRequestJson struct {
	// ChargingLimitSource corresponds to the JSON schema field "chargingLimitSource".
	ChargingLimitSource ChargingLimitSourceEnumType_7 `json:"chargingLimitSource"`

	// ChargingProfile corresponds to the JSON schema field "chargingProfile".
	ChargingProfile []ChargingProfileType `json:"chargingProfile"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// The evse to which the charging profile applies. If evseId = 0, the message
	// contains an overall limit for the Charging Station.
	//
	EvseId int `json:"evseId"`

	// Id used to match the &lt;&lt;getchargingprofilesrequest,
	// GetChargingProfilesRequest&gt;&gt; message with the resulting
	// ReportChargingProfilesRequest messages. When the CSMS provided a requestId in
	// the &lt;&lt;getchargingprofilesrequest, GetChargingProfilesRequest&gt;&gt;,
	// this field SHALL contain the same value.
	//
	RequestId int `json:"requestId"`

	// To Be Continued. Default value when omitted: false. false indicates that there
	// are no further messages as part of this report.
	//
	Tbc bool `json:"tbc,omitempty"`
}

type ReportChargingProfilesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_85 `json:"customData,omitempty"`
}

// Class to report components, variables and variable attributes and
// characteristics.
//
type ReportDataType struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType_7 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_78 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_6 `json:"variable"`

	// VariableAttribute corresponds to the JSON schema field "variableAttribute".
	VariableAttribute []VariableAttributeType `json:"variableAttribute"`

	// VariableCharacteristics corresponds to the JSON schema field
	// "variableCharacteristics".
	VariableCharacteristics *VariableCharacteristicsType `json:"variableCharacteristics,omitempty"`
}

type RequestStartStopStatusEnumType string

const RequestStartStopStatusEnumTypeAccepted RequestStartStopStatusEnumType = "Accepted"
const RequestStartStopStatusEnumTypeRejected RequestStartStopStatusEnumType = "Rejected"

type RequestStartStopStatusEnumType_1 string

const RequestStartStopStatusEnumType_1_Accepted RequestStartStopStatusEnumType_1 = "Accepted"
const RequestStartStopStatusEnumType_1_Rejected RequestStartStopStatusEnumType_1 = "Rejected"

type RequestStartStopStatusEnumType_2 string

const RequestStartStopStatusEnumType_2_Accepted RequestStartStopStatusEnumType_2 = "Accepted"
const RequestStartStopStatusEnumType_2_Rejected RequestStartStopStatusEnumType_2 = "Rejected"

type RequestStartStopStatusEnumType_3 string

const RequestStartStopStatusEnumType_3_Accepted RequestStartStopStatusEnumType_3 = "Accepted"
const RequestStartStopStatusEnumType_3_Rejected RequestStartStopStatusEnumType_3 = "Rejected"

type RequestStartTransactionRequestJson struct {
	// ChargingProfile corresponds to the JSON schema field "chargingProfile".
	ChargingProfile *ChargingProfileType_1 `json:"chargingProfile,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Number of the EVSE on which to start the transaction. EvseId SHALL be &gt; 0
	//
	EvseId *int `json:"evseId,omitempty"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType_3 `json:"groupIdToken,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken IdTokenType_3 `json:"idToken"`

	// Id given by the server to this start request. The Charging Station might return
	// this in the &lt;&lt;transactioneventrequest, TransactionEventRequest&gt;&gt;,
	// letting the server know which transaction was started for this request. Use to
	// start a transaction.
	//
	RemoteStartId int `json:"remoteStartId"`
}

type RequestStartTransactionResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_87 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status RequestStartStopStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_26 `json:"statusInfo,omitempty"`

	// When the transaction was already started by the Charging Station before the
	// RequestStartTransactionRequest was received, for example: cable plugged in
	// first. This contains the transactionId of the already started transaction.
	//
	TransactionId *string `json:"transactionId,omitempty"`
}

type RequestStopTransactionRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_88 `json:"customData,omitempty"`

	// The identifier of the transaction which the Charging Station is requested to
	// stop.
	//
	TransactionId string `json:"transactionId"`
}

type RequestStopTransactionResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_89 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status RequestStartStopStatusEnumType_3 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_27 `json:"statusInfo,omitempty"`
}

type ReservationStatusUpdateRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_90 `json:"customData,omitempty"`

	// The ID of the reservation.
	//
	ReservationId int `json:"reservationId"`

	// ReservationUpdateStatus corresponds to the JSON schema field
	// "reservationUpdateStatus".
	ReservationUpdateStatus ReservationUpdateStatusEnumType_1 `json:"reservationUpdateStatus"`
}

type ReservationStatusUpdateResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_91 `json:"customData,omitempty"`
}

type ReservationUpdateStatusEnumType string

const ReservationUpdateStatusEnumTypeExpired ReservationUpdateStatusEnumType = "Expired"
const ReservationUpdateStatusEnumTypeRemoved ReservationUpdateStatusEnumType = "Removed"

type ReservationUpdateStatusEnumType_1 string

const ReservationUpdateStatusEnumType_1_Expired ReservationUpdateStatusEnumType_1 = "Expired"
const ReservationUpdateStatusEnumType_1_Removed ReservationUpdateStatusEnumType_1 = "Removed"

type ReserveNowRequestJson struct {
	// ConnectorType corresponds to the JSON schema field "connectorType".
	ConnectorType *ConnectorEnumType_1 `json:"connectorType,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_92 `json:"customData,omitempty"`

	// This contains ID of the evse to be reserved.
	//
	EvseId *int `json:"evseId,omitempty"`

	// Date and time at which the reservation expires.
	//
	ExpiryDateTime string `json:"expiryDateTime"`

	// GroupIdToken corresponds to the JSON schema field "groupIdToken".
	GroupIdToken *IdTokenType_4 `json:"groupIdToken,omitempty"`

	// Id of reservation.
	//
	Id int `json:"id"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken IdTokenType_4 `json:"idToken"`
}

type ReserveNowResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_93 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ReserveNowStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_28 `json:"statusInfo,omitempty"`
}

type ReserveNowStatusEnumType string

const ReserveNowStatusEnumTypeAccepted ReserveNowStatusEnumType = "Accepted"
const ReserveNowStatusEnumTypeFaulted ReserveNowStatusEnumType = "Faulted"
const ReserveNowStatusEnumTypeOccupied ReserveNowStatusEnumType = "Occupied"
const ReserveNowStatusEnumTypeRejected ReserveNowStatusEnumType = "Rejected"
const ReserveNowStatusEnumTypeUnavailable ReserveNowStatusEnumType = "Unavailable"

type ReserveNowStatusEnumType_1 string

const ReserveNowStatusEnumType_1_Accepted ReserveNowStatusEnumType_1 = "Accepted"
const ReserveNowStatusEnumType_1_Faulted ReserveNowStatusEnumType_1 = "Faulted"
const ReserveNowStatusEnumType_1_Occupied ReserveNowStatusEnumType_1 = "Occupied"
const ReserveNowStatusEnumType_1_Rejected ReserveNowStatusEnumType_1 = "Rejected"
const ReserveNowStatusEnumType_1_Unavailable ReserveNowStatusEnumType_1 = "Unavailable"

type ResetEnumType string

const ResetEnumTypeImmediate ResetEnumType = "Immediate"
const ResetEnumTypeOnIdle ResetEnumType = "OnIdle"

type ResetEnumType_1 string

const ResetEnumType_1_Immediate ResetEnumType_1 = "Immediate"
const ResetEnumType_1_OnIdle ResetEnumType_1 = "OnIdle"

type ResetRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_94 `json:"customData,omitempty"`

	// This contains the ID of a specific EVSE that needs to be reset, instead of the
	// entire Charging Station.
	//
	EvseId *int `json:"evseId,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type ResetEnumType_1 `json:"type"`
}

type ResetResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_95 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ResetStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_29 `json:"statusInfo,omitempty"`
}

type ResetStatusEnumType string

const ResetStatusEnumTypeAccepted ResetStatusEnumType = "Accepted"
const ResetStatusEnumTypeRejected ResetStatusEnumType = "Rejected"
const ResetStatusEnumTypeScheduled ResetStatusEnumType = "Scheduled"

type ResetStatusEnumType_1 string

const ResetStatusEnumType_1_Accepted ResetStatusEnumType_1 = "Accepted"
const ResetStatusEnumType_1_Rejected ResetStatusEnumType_1 = "Rejected"
const ResetStatusEnumType_1_Scheduled ResetStatusEnumType_1 = "Scheduled"

// Sales_ Tariff_ Entry
// urn:x-oca:ocpp:uid:2:233271
//
type SalesTariffEntryType struct {
	// ConsumptionCost corresponds to the JSON schema field "consumptionCost".
	ConsumptionCost []ConsumptionCostType `json:"consumptionCost,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// Sales_ Tariff_ Entry. E_ Price_ Level. Unsigned_ Integer
	// urn:x-oca:ocpp:uid:1:569281
	// Defines the price level of this SalesTariffEntry (referring to
	// NumEPriceLevels). Small values for the EPriceLevel represent a cheaper
	// TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	//
	EPriceLevel *int `json:"ePriceLevel,omitempty"`

	// RelativeTimeInterval corresponds to the JSON schema field
	// "relativeTimeInterval".
	RelativeTimeInterval RelativeTimeIntervalType `json:"relativeTimeInterval"`
}

// Sales_ Tariff_ Entry
// urn:x-oca:ocpp:uid:2:233271
//
type SalesTariffEntryType_1 struct {
	// ConsumptionCost corresponds to the JSON schema field "consumptionCost".
	ConsumptionCost []ConsumptionCostType_1 `json:"consumptionCost,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`

	// Sales_ Tariff_ Entry. E_ Price_ Level. Unsigned_ Integer
	// urn:x-oca:ocpp:uid:1:569281
	// Defines the price level of this SalesTariffEntry (referring to
	// NumEPriceLevels). Small values for the EPriceLevel represent a cheaper
	// TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	//
	EPriceLevel *int `json:"ePriceLevel,omitempty"`

	// RelativeTimeInterval corresponds to the JSON schema field
	// "relativeTimeInterval".
	RelativeTimeInterval RelativeTimeIntervalType_1 `json:"relativeTimeInterval"`
}

// Sales_ Tariff_ Entry
// urn:x-oca:ocpp:uid:2:233271
//
type SalesTariffEntryType_2 struct {
	// ConsumptionCost corresponds to the JSON schema field "consumptionCost".
	ConsumptionCost []ConsumptionCostType_2 `json:"consumptionCost,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// Sales_ Tariff_ Entry. E_ Price_ Level. Unsigned_ Integer
	// urn:x-oca:ocpp:uid:1:569281
	// Defines the price level of this SalesTariffEntry (referring to
	// NumEPriceLevels). Small values for the EPriceLevel represent a cheaper
	// TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	//
	EPriceLevel *int `json:"ePriceLevel,omitempty"`

	// RelativeTimeInterval corresponds to the JSON schema field
	// "relativeTimeInterval".
	RelativeTimeInterval RelativeTimeIntervalType_2 `json:"relativeTimeInterval"`
}

// Sales_ Tariff_ Entry
// urn:x-oca:ocpp:uid:2:233271
//
type SalesTariffEntryType_3 struct {
	// ConsumptionCost corresponds to the JSON schema field "consumptionCost".
	ConsumptionCost []ConsumptionCostType_3 `json:"consumptionCost,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Sales_ Tariff_ Entry. E_ Price_ Level. Unsigned_ Integer
	// urn:x-oca:ocpp:uid:1:569281
	// Defines the price level of this SalesTariffEntry (referring to
	// NumEPriceLevels). Small values for the EPriceLevel represent a cheaper
	// TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	//
	EPriceLevel *int `json:"ePriceLevel,omitempty"`

	// RelativeTimeInterval corresponds to the JSON schema field
	// "relativeTimeInterval".
	RelativeTimeInterval RelativeTimeIntervalType_3 `json:"relativeTimeInterval"`
}

// Sales_ Tariff_ Entry
// urn:x-oca:ocpp:uid:2:233271
//
type SalesTariffEntryType_4 struct {
	// ConsumptionCost corresponds to the JSON schema field "consumptionCost".
	ConsumptionCost []ConsumptionCostType_4 `json:"consumptionCost,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// Sales_ Tariff_ Entry. E_ Price_ Level. Unsigned_ Integer
	// urn:x-oca:ocpp:uid:1:569281
	// Defines the price level of this SalesTariffEntry (referring to
	// NumEPriceLevels). Small values for the EPriceLevel represent a cheaper
	// TariffEntry. Large values for the EPriceLevel represent a more expensive
	// TariffEntry.
	//
	EPriceLevel *int `json:"ePriceLevel,omitempty"`

	// RelativeTimeInterval corresponds to the JSON schema field
	// "relativeTimeInterval".
	RelativeTimeInterval RelativeTimeIntervalType_4 `json:"relativeTimeInterval"`
}

// Sales_ Tariff
// urn:x-oca:ocpp:uid:2:233272
// NOTE: This dataType is based on dataTypes from &lt;&lt;ref-ISOIEC15118-2,ISO
// 15118-2&gt;&gt;.
//
type SalesTariffType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_64 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// SalesTariff identifier used to identify one sales tariff. An SAID remains a
	// unique identifier for one schedule throughout a charging session.
	//
	Id int `json:"id"`

	// Sales_ Tariff. Num_ E_ Price_ Levels. Counter
	// urn:x-oca:ocpp:uid:1:569284
	// Defines the overall number of distinct price levels used across all provided
	// SalesTariff elements.
	//
	NumEPriceLevels *int `json:"numEPriceLevels,omitempty"`

	// Sales_ Tariff. Sales. Tariff_ Description
	// urn:x-oca:ocpp:uid:1:569283
	// A human readable title/short description of the sales tariff e.g. for HMI
	// display purposes.
	//
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty"`

	// SalesTariffEntry corresponds to the JSON schema field "salesTariffEntry".
	SalesTariffEntry []SalesTariffEntryType `json:"salesTariffEntry"`
}

// Sales_ Tariff
// urn:x-oca:ocpp:uid:2:233272
// NOTE: This dataType is based on dataTypes from &lt;&lt;ref-ISOIEC15118-2,ISO
// 15118-2&gt;&gt;.
//
type SalesTariffType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_72 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// SalesTariff identifier used to identify one sales tariff. An SAID remains a
	// unique identifier for one schedule throughout a charging session.
	//
	Id int `json:"id"`

	// Sales_ Tariff. Num_ E_ Price_ Levels. Counter
	// urn:x-oca:ocpp:uid:1:569284
	// Defines the overall number of distinct price levels used across all provided
	// SalesTariff elements.
	//
	NumEPriceLevels *int `json:"numEPriceLevels,omitempty"`

	// Sales_ Tariff. Sales. Tariff_ Description
	// urn:x-oca:ocpp:uid:1:569283
	// A human readable title/short description of the sales tariff e.g. for HMI
	// display purposes.
	//
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty"`

	// SalesTariffEntry corresponds to the JSON schema field "salesTariffEntry".
	SalesTariffEntry []SalesTariffEntryType_1 `json:"salesTariffEntry"`
}

// Sales_ Tariff
// urn:x-oca:ocpp:uid:2:233272
// NOTE: This dataType is based on dataTypes from &lt;&lt;ref-ISOIEC15118-2,ISO
// 15118-2&gt;&gt;.
//
type SalesTariffType_2 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_84 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// SalesTariff identifier used to identify one sales tariff. An SAID remains a
	// unique identifier for one schedule throughout a charging session.
	//
	Id int `json:"id"`

	// Sales_ Tariff. Num_ E_ Price_ Levels. Counter
	// urn:x-oca:ocpp:uid:1:569284
	// Defines the overall number of distinct price levels used across all provided
	// SalesTariff elements.
	//
	NumEPriceLevels *int `json:"numEPriceLevels,omitempty"`

	// Sales_ Tariff. Sales. Tariff_ Description
	// urn:x-oca:ocpp:uid:1:569283
	// A human readable title/short description of the sales tariff e.g. for HMI
	// display purposes.
	//
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty"`

	// SalesTariffEntry corresponds to the JSON schema field "salesTariffEntry".
	SalesTariffEntry []SalesTariffEntryType_2 `json:"salesTariffEntry"`
}

// Sales_ Tariff
// urn:x-oca:ocpp:uid:2:233272
// NOTE: This dataType is based on dataTypes from &lt;&lt;ref-ISOIEC15118-2,ISO
// 15118-2&gt;&gt;.
//
type SalesTariffType_3 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_86 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// SalesTariff identifier used to identify one sales tariff. An SAID remains a
	// unique identifier for one schedule throughout a charging session.
	//
	Id int `json:"id"`

	// Sales_ Tariff. Num_ E_ Price_ Levels. Counter
	// urn:x-oca:ocpp:uid:1:569284
	// Defines the overall number of distinct price levels used across all provided
	// SalesTariff elements.
	//
	NumEPriceLevels *int `json:"numEPriceLevels,omitempty"`

	// Sales_ Tariff. Sales. Tariff_ Description
	// urn:x-oca:ocpp:uid:1:569283
	// A human readable title/short description of the sales tariff e.g. for HMI
	// display purposes.
	//
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty"`

	// SalesTariffEntry corresponds to the JSON schema field "salesTariffEntry".
	SalesTariffEntry []SalesTariffEntryType_3 `json:"salesTariffEntry"`
}

// Sales_ Tariff
// urn:x-oca:ocpp:uid:2:233272
// NOTE: This dataType is based on dataTypes from &lt;&lt;ref-ISOIEC15118-2,ISO
// 15118-2&gt;&gt;.
//
type SalesTariffType_4 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// Identified_ Object. MRID. Numeric_ Identifier
	// urn:x-enexis:ecdm:uid:1:569198
	// SalesTariff identifier used to identify one sales tariff. An SAID remains a
	// unique identifier for one schedule throughout a charging session.
	//
	Id int `json:"id"`

	// Sales_ Tariff. Num_ E_ Price_ Levels. Counter
	// urn:x-oca:ocpp:uid:1:569284
	// Defines the overall number of distinct price levels used across all provided
	// SalesTariff elements.
	//
	NumEPriceLevels *int `json:"numEPriceLevels,omitempty"`

	// Sales_ Tariff. Sales. Tariff_ Description
	// urn:x-oca:ocpp:uid:1:569283
	// A human readable title/short description of the sales tariff e.g. for HMI
	// display purposes.
	//
	SalesTariffDescription *string `json:"salesTariffDescription,omitempty"`

	// SalesTariffEntry corresponds to the JSON schema field "salesTariffEntry".
	SalesTariffEntry []SalesTariffEntryType_4 `json:"salesTariffEntry"`
}

// Sampled_ Value
// urn:x-oca:ocpp:uid:2:233266
// Single sampled value in MeterValues. Each value can be accompanied by optional
// fields.
//
// To save on mobile data usage, default values of all of the optional fields are
// such that. The value without any additional fields will be interpreted, as a
// register reading of active import energy in Wh (Watt-hour) units.
//
type SampledValueType struct {
	// Context corresponds to the JSON schema field "context".
	Context *ReadingContextEnumType `json:"context,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_62 `json:"customData,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *LocationEnumType_1 `json:"location,omitempty"`

	// Measurand corresponds to the JSON schema field "measurand".
	Measurand *MeasurandEnumType_1 `json:"measurand,omitempty"`

	// Phase corresponds to the JSON schema field "phase".
	Phase *PhaseEnumType `json:"phase,omitempty"`

	// SignedMeterValue corresponds to the JSON schema field "signedMeterValue".
	SignedMeterValue *SignedMeterValueType `json:"signedMeterValue,omitempty"`

	// UnitOfMeasure corresponds to the JSON schema field "unitOfMeasure".
	UnitOfMeasure *UnitOfMeasureType `json:"unitOfMeasure,omitempty"`

	// Sampled_ Value. Value. Measure
	// urn:x-oca:ocpp:uid:1:569260
	// Indicates the measured value.
	//
	//
	Value float64 `json:"value"`
}

// Sampled_ Value
// urn:x-oca:ocpp:uid:2:233266
// Single sampled value in MeterValues. Each value can be accompanied by optional
// fields.
//
// To save on mobile data usage, default values of all of the optional fields are
// such that. The value without any additional fields will be interpreted, as a
// register reading of active import energy in Wh (Watt-hour) units.
//
type SampledValueType_1 struct {
	// Context corresponds to the JSON schema field "context".
	Context *ReadingContextEnumType_2 `json:"context,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// Location corresponds to the JSON schema field "location".
	Location *LocationEnumType_3 `json:"location,omitempty"`

	// Measurand corresponds to the JSON schema field "measurand".
	Measurand *MeasurandEnumType_3 `json:"measurand,omitempty"`

	// Phase corresponds to the JSON schema field "phase".
	Phase *PhaseEnumType_2 `json:"phase,omitempty"`

	// SignedMeterValue corresponds to the JSON schema field "signedMeterValue".
	SignedMeterValue *SignedMeterValueType_1 `json:"signedMeterValue,omitempty"`

	// UnitOfMeasure corresponds to the JSON schema field "unitOfMeasure".
	UnitOfMeasure *UnitOfMeasureType_1 `json:"unitOfMeasure,omitempty"`

	// Sampled_ Value. Value. Measure
	// urn:x-oca:ocpp:uid:1:569260
	// Indicates the measured value.
	//
	//
	Value float64 `json:"value"`
}

type SecurityEventNotificationRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_96 `json:"customData,omitempty"`

	// Additional information about the occurred security event.
	//
	TechInfo *string `json:"techInfo,omitempty"`

	// Date and time at which the event occurred.
	//
	Timestamp string `json:"timestamp"`

	// Type of the security event. This value should be taken from the Security events
	// list.
	//
	Type string `json:"type"`
}

type SecurityEventNotificationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_97 `json:"customData,omitempty"`
}

type SendLocalListRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_98 `json:"customData,omitempty"`

	// LocalAuthorizationList corresponds to the JSON schema field
	// "localAuthorizationList".
	LocalAuthorizationList []AuthorizationData `json:"localAuthorizationList,omitempty"`

	// UpdateType corresponds to the JSON schema field "updateType".
	UpdateType UpdateEnumType_1 `json:"updateType"`

	// In case of a full update this is the version number of the full list. In case
	// of a differential update it is the version number of the list after the update
	// has been applied.
	//
	VersionNumber int `json:"versionNumber"`
}

type SendLocalListResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_99 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status SendLocalListStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_30 `json:"statusInfo,omitempty"`
}

type SendLocalListStatusEnumType string

const SendLocalListStatusEnumTypeAccepted SendLocalListStatusEnumType = "Accepted"
const SendLocalListStatusEnumTypeFailed SendLocalListStatusEnumType = "Failed"
const SendLocalListStatusEnumTypeVersionMismatch SendLocalListStatusEnumType = "VersionMismatch"

type SendLocalListStatusEnumType_1 string

const SendLocalListStatusEnumType_1_Accepted SendLocalListStatusEnumType_1 = "Accepted"
const SendLocalListStatusEnumType_1_Failed SendLocalListStatusEnumType_1 = "Failed"
const SendLocalListStatusEnumType_1_VersionMismatch SendLocalListStatusEnumType_1 = "VersionMismatch"

type SetChargingProfileRequestJson struct {
	// ChargingProfile corresponds to the JSON schema field "chargingProfile".
	ChargingProfile ChargingProfileType_2 `json:"chargingProfile"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_100 `json:"customData,omitempty"`

	// For TxDefaultProfile an evseId=0 applies the profile to each individual evse.
	// For ChargingStationMaxProfile and ChargingStationExternalConstraints an
	// evseId=0 contains an overal limit for the whole Charging Station.
	//
	EvseId int `json:"evseId"`
}

type SetChargingProfileResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_101 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status ChargingProfileStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_31 `json:"statusInfo,omitempty"`
}

type SetDisplayMessageRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_102 `json:"customData,omitempty"`

	// Message corresponds to the JSON schema field "message".
	Message MessageInfoType_1 `json:"message"`
}

type SetDisplayMessageResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_103 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status DisplayMessageStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_32 `json:"statusInfo,omitempty"`
}

type SetMonitoringBaseRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_104 `json:"customData,omitempty"`

	// MonitoringBase corresponds to the JSON schema field "monitoringBase".
	MonitoringBase MonitoringBaseEnumType_1 `json:"monitoringBase"`
}

type SetMonitoringBaseResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_105 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericDeviceModelStatusEnumType_7 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_33 `json:"statusInfo,omitempty"`
}

// Class to hold parameters of SetVariableMonitoring request.
//
type SetMonitoringDataType struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType_9 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_110 `json:"customData,omitempty"`

	// An id SHALL only be given to replace an existing monitor. The Charging Station
	// handles the generation of id's for new monitors.
	//
	//
	Id *int `json:"id,omitempty"`

	// The severity that will be assigned to an event that is triggered by this
	// monitor. The severity range is 0-9, with 0 as the highest and 9 as the lowest
	// severity level.
	//
	// The severity levels have the following meaning: +
	// *0-Danger* +
	// Indicates lives are potentially in danger. Urgent attention is needed and
	// action should be taken immediately. +
	// *1-Hardware Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to Hardware issues. Action is required. +
	// *2-System Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to software or minor hardware issues. Action is required. +
	// *3-Critical* +
	// Indicates a critical error. Action is required. +
	// *4-Error* +
	// Indicates a non-urgent error. Action is required. +
	// *5-Alert* +
	// Indicates an alert event. Default severity for any type of monitoring event.  +
	// *6-Warning* +
	// Indicates a warning event. Action may be required. +
	// *7-Notice* +
	// Indicates an unusual event. No immediate action is required. +
	// *8-Informational* +
	// Indicates a regular operational event. May be used for reporting, measuring
	// throughput, etc. No action is required. +
	// *9-Debug* +
	// Indicates information useful to developers for debugging, not useful during
	// operations.
	//
	//
	Severity int `json:"severity"`

	// Monitor only active when a transaction is ongoing on a component relevant to
	// this transaction. Default = false.
	//
	//
	Transaction bool `json:"transaction,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type MonitorEnumType_3 `json:"type"`

	// Value for threshold or delta monitoring.
	// For Periodic or PeriodicClockAligned this is the interval in seconds.
	//
	//
	Value float64 `json:"value"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_7 `json:"variable"`
}

type SetMonitoringLevelRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_106 `json:"customData,omitempty"`

	// The Charging Station SHALL only report events with a severity number lower than
	// or equal to this severity.
	// The severity range is 0-9, with 0 as the highest and 9 as the lowest severity
	// level.
	//
	// The severity levels have the following meaning: +
	// *0-Danger* +
	// Indicates lives are potentially in danger. Urgent attention is needed and
	// action should be taken immediately. +
	// *1-Hardware Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to Hardware issues. Action is required. +
	// *2-System Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to software or minor hardware issues. Action is required. +
	// *3-Critical* +
	// Indicates a critical error. Action is required. +
	// *4-Error* +
	// Indicates a non-urgent error. Action is required. +
	// *5-Alert* +
	// Indicates an alert event. Default severity for any type of monitoring event.  +
	// *6-Warning* +
	// Indicates a warning event. Action may be required. +
	// *7-Notice* +
	// Indicates an unusual event. No immediate action is required. +
	// *8-Informational* +
	// Indicates a regular operational event. May be used for reporting, measuring
	// throughput, etc. No action is required. +
	// *9-Debug* +
	// Indicates information useful to developers for debugging, not useful during
	// operations.
	//
	//
	//
	Severity int `json:"severity"`
}

type SetMonitoringLevelResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_107 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericStatusEnumType_7 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_34 `json:"statusInfo,omitempty"`
}

// Class to hold result of SetVariableMonitoring request.
//
type SetMonitoringResultType struct {
	// Component corresponds to the JSON schema field "component".
	Component ComponentType_10 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_111 `json:"customData,omitempty"`

	// Id given to the VariableMonitor by the Charging Station. The Id is only
	// returned when status is accepted. Installed VariableMonitors should have unique
	// id's but the id's of removed Installed monitors should have unique id's but the
	// id's of removed monitors MAY be reused.
	//
	Id *int `json:"id,omitempty"`

	// The severity that will be assigned to an event that is triggered by this
	// monitor. The severity range is 0-9, with 0 as the highest and 9 as the lowest
	// severity level.
	//
	// The severity levels have the following meaning: +
	// *0-Danger* +
	// Indicates lives are potentially in danger. Urgent attention is needed and
	// action should be taken immediately. +
	// *1-Hardware Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to Hardware issues. Action is required. +
	// *2-System Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to software or minor hardware issues. Action is required. +
	// *3-Critical* +
	// Indicates a critical error. Action is required. +
	// *4-Error* +
	// Indicates a non-urgent error. Action is required. +
	// *5-Alert* +
	// Indicates an alert event. Default severity for any type of monitoring event.  +
	// *6-Warning* +
	// Indicates a warning event. Action may be required. +
	// *7-Notice* +
	// Indicates an unusual event. No immediate action is required. +
	// *8-Informational* +
	// Indicates a regular operational event. May be used for reporting, measuring
	// throughput, etc. No action is required. +
	// *9-Debug* +
	// Indicates information useful to developers for debugging, not useful during
	// operations.
	//
	//
	Severity int `json:"severity"`

	// Status corresponds to the JSON schema field "status".
	Status SetMonitoringStatusEnumType `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_36 `json:"statusInfo,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type MonitorEnumType_5 `json:"type"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_8 `json:"variable"`
}

type SetMonitoringStatusEnumType string

const SetMonitoringStatusEnumTypeAccepted SetMonitoringStatusEnumType = "Accepted"
const SetMonitoringStatusEnumTypeDuplicate SetMonitoringStatusEnumType = "Duplicate"
const SetMonitoringStatusEnumTypeRejected SetMonitoringStatusEnumType = "Rejected"
const SetMonitoringStatusEnumTypeUnknownComponent SetMonitoringStatusEnumType = "UnknownComponent"
const SetMonitoringStatusEnumTypeUnknownVariable SetMonitoringStatusEnumType = "UnknownVariable"
const SetMonitoringStatusEnumTypeUnsupportedMonitorType SetMonitoringStatusEnumType = "UnsupportedMonitorType"

type SetMonitoringStatusEnumType_1 string

const SetMonitoringStatusEnumType_1_Accepted SetMonitoringStatusEnumType_1 = "Accepted"
const SetMonitoringStatusEnumType_1_Duplicate SetMonitoringStatusEnumType_1 = "Duplicate"
const SetMonitoringStatusEnumType_1_Rejected SetMonitoringStatusEnumType_1 = "Rejected"
const SetMonitoringStatusEnumType_1_UnknownComponent SetMonitoringStatusEnumType_1 = "UnknownComponent"
const SetMonitoringStatusEnumType_1_UnknownVariable SetMonitoringStatusEnumType_1 = "UnknownVariable"
const SetMonitoringStatusEnumType_1_UnsupportedMonitorType SetMonitoringStatusEnumType_1 = "UnsupportedMonitorType"

type SetNetworkProfileRequestJson struct {
	// Slot in which the configuration should be stored.
	//
	ConfigurationSlot int `json:"configurationSlot"`

	// ConnectionData corresponds to the JSON schema field "connectionData".
	ConnectionData NetworkConnectionProfileType `json:"connectionData"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_108 `json:"customData,omitempty"`
}

type SetNetworkProfileResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_109 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status SetNetworkProfileStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_35 `json:"statusInfo,omitempty"`
}

type SetNetworkProfileStatusEnumType string

const SetNetworkProfileStatusEnumTypeAccepted SetNetworkProfileStatusEnumType = "Accepted"
const SetNetworkProfileStatusEnumTypeFailed SetNetworkProfileStatusEnumType = "Failed"
const SetNetworkProfileStatusEnumTypeRejected SetNetworkProfileStatusEnumType = "Rejected"

type SetNetworkProfileStatusEnumType_1 string

const SetNetworkProfileStatusEnumType_1_Accepted SetNetworkProfileStatusEnumType_1 = "Accepted"
const SetNetworkProfileStatusEnumType_1_Failed SetNetworkProfileStatusEnumType_1 = "Failed"
const SetNetworkProfileStatusEnumType_1_Rejected SetNetworkProfileStatusEnumType_1 = "Rejected"

type SetVariableDataType struct {
	// AttributeType corresponds to the JSON schema field "attributeType".
	AttributeType *AttributeEnumType_7 `json:"attributeType,omitempty"`

	// Value to be assigned to attribute of variable.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-configuration-value-size,ConfigurationValueSize&gt;&gt; can
	// be used to limit SetVariableData.attributeValue and
	// VariableCharacteristics.valueList. The max size of these values will always
	// remain equal.
	//
	AttributeValue string `json:"attributeValue"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType_11 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_112 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_9 `json:"variable"`
}

type SetVariableMonitoringRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_110 `json:"customData,omitempty"`

	// SetMonitoringData corresponds to the JSON schema field "setMonitoringData".
	SetMonitoringData []SetMonitoringDataType `json:"setMonitoringData"`
}

type SetVariableMonitoringResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_111 `json:"customData,omitempty"`

	// SetMonitoringResult corresponds to the JSON schema field "setMonitoringResult".
	SetMonitoringResult []SetMonitoringResultType `json:"setMonitoringResult"`
}

type SetVariableResultType struct {
	// AttributeStatus corresponds to the JSON schema field "attributeStatus".
	AttributeStatus SetVariableStatusEnumType `json:"attributeStatus"`

	// AttributeStatusInfo corresponds to the JSON schema field "attributeStatusInfo".
	AttributeStatusInfo *StatusInfoType_37 `json:"attributeStatusInfo,omitempty"`

	// AttributeType corresponds to the JSON schema field "attributeType".
	AttributeType *AttributeEnumType_9 `json:"attributeType,omitempty"`

	// Component corresponds to the JSON schema field "component".
	Component ComponentType_12 `json:"component"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_113 `json:"customData,omitempty"`

	// Variable corresponds to the JSON schema field "variable".
	Variable VariableType_10 `json:"variable"`
}

type SetVariableStatusEnumType string

const SetVariableStatusEnumTypeAccepted SetVariableStatusEnumType = "Accepted"
const SetVariableStatusEnumTypeNotSupportedAttributeType SetVariableStatusEnumType = "NotSupportedAttributeType"
const SetVariableStatusEnumTypeRebootRequired SetVariableStatusEnumType = "RebootRequired"
const SetVariableStatusEnumTypeRejected SetVariableStatusEnumType = "Rejected"
const SetVariableStatusEnumTypeUnknownComponent SetVariableStatusEnumType = "UnknownComponent"
const SetVariableStatusEnumTypeUnknownVariable SetVariableStatusEnumType = "UnknownVariable"

type SetVariableStatusEnumType_1 string

const SetVariableStatusEnumType_1_Accepted SetVariableStatusEnumType_1 = "Accepted"
const SetVariableStatusEnumType_1_NotSupportedAttributeType SetVariableStatusEnumType_1 = "NotSupportedAttributeType"
const SetVariableStatusEnumType_1_RebootRequired SetVariableStatusEnumType_1 = "RebootRequired"
const SetVariableStatusEnumType_1_Rejected SetVariableStatusEnumType_1 = "Rejected"
const SetVariableStatusEnumType_1_UnknownComponent SetVariableStatusEnumType_1 = "UnknownComponent"
const SetVariableStatusEnumType_1_UnknownVariable SetVariableStatusEnumType_1 = "UnknownVariable"

type SetVariablesRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_112 `json:"customData,omitempty"`

	// SetVariableData corresponds to the JSON schema field "setVariableData".
	SetVariableData []SetVariableDataType `json:"setVariableData"`
}

type SetVariablesResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_113 `json:"customData,omitempty"`

	// SetVariableResult corresponds to the JSON schema field "setVariableResult".
	SetVariableResult []SetVariableResultType `json:"setVariableResult"`
}

type SignCertificateRequestJson struct {
	// CertificateType corresponds to the JSON schema field "certificateType".
	CertificateType *CertificateSigningUseEnumType_3 `json:"certificateType,omitempty"`

	// The Charging Station SHALL send the public key in form of a Certificate Signing
	// Request (CSR) as described in RFC 2986 [22] and then PEM encoded, using the
	// &lt;&lt;signcertificaterequest,SignCertificateRequest&gt;&gt; message.
	//
	Csr string `json:"csr"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_114 `json:"customData,omitempty"`
}

type SignCertificateResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_115 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status GenericStatusEnumType_9 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_38 `json:"statusInfo,omitempty"`
}

// Represent a signed version of the meter value.
//
type SignedMeterValueType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_62 `json:"customData,omitempty"`

	// Method used to encode the meter values before applying the digital signature
	// algorithm.
	//
	EncodingMethod string `json:"encodingMethod"`

	// Base64 encoded, sending depends on configuration variable
	// _PublicKeyWithSignedMeterValue_.
	//
	PublicKey string `json:"publicKey"`

	// Base64 encoded, contains the signed data which might contain more then just the
	// meter value. It can contain information like timestamps, reference to a
	// customer etc.
	//
	SignedMeterData string `json:"signedMeterData"`

	// Method used to create the digital signature.
	//
	SigningMethod string `json:"signingMethod"`
}

// Represent a signed version of the meter value.
//
type SignedMeterValueType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// Method used to encode the meter values before applying the digital signature
	// algorithm.
	//
	EncodingMethod string `json:"encodingMethod"`

	// Base64 encoded, sending depends on configuration variable
	// _PublicKeyWithSignedMeterValue_.
	//
	PublicKey string `json:"publicKey"`

	// Base64 encoded, contains the signed data which might contain more then just the
	// meter value. It can contain information like timestamps, reference to a
	// customer etc.
	//
	SignedMeterData string `json:"signedMeterData"`

	// Method used to create the digital signature.
	//
	SigningMethod string `json:"signingMethod"`
}

// Element providing more information about the status.
//
type StatusInfoType struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_3 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_1 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_5 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_10 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_27 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_11 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_31 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_12 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_33 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_13 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_35 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_14 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_37 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_15 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_39 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_16 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_41 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_17 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_43 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_18 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_47 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_19 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_49 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_2 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_7 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_20 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_51 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_21 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_55 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_22 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_59 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_23 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_71 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_24 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_73 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_25 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_81 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_26 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_87 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_27 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_89 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_28 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_93 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_29 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_95 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_3 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_9 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_30 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_99 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_31 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_101 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_32 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_103 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_33 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_105 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_34 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_107 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_35 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_109 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_36 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_111 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_37 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_113 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_38 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_115 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_39 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_121 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_4 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_11 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_40 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_123 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_41 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_127 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_5 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_13 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_6 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_15 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_7 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_17 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_8 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_23 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

// Element providing more information about the status.
//
type StatusInfoType_9 struct {
	// Additional text to provide detailed information.
	//
	AdditionalInfo *string `json:"additionalInfo,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_25 `json:"customData,omitempty"`

	// A predefined code for the reason why the status is returned in this response.
	// The string is case-insensitive.
	//
	ReasonCode string `json:"reasonCode"`
}

type StatusNotificationRequestJson struct {
	// The id of the connector within the EVSE for which the status is reported.
	//
	ConnectorId int `json:"connectorId"`

	// ConnectorStatus corresponds to the JSON schema field "connectorStatus".
	ConnectorStatus ConnectorStatusEnumType_1 `json:"connectorStatus"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_116 `json:"customData,omitempty"`

	// The id of the EVSE to which the connector belongs for which the the status is
	// reported.
	//
	EvseId int `json:"evseId"`

	// The time for which the status is reported. If absent time of receipt of the
	// message will be assumed.
	//
	Timestamp string `json:"timestamp"`
}

type StatusNotificationResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_117 `json:"customData,omitempty"`
}

type TransactionEventEnumType string

const TransactionEventEnumTypeEnded TransactionEventEnumType = "Ended"
const TransactionEventEnumTypeStarted TransactionEventEnumType = "Started"
const TransactionEventEnumTypeUpdated TransactionEventEnumType = "Updated"

type TransactionEventEnumType_1 string

const TransactionEventEnumType_1_Ended TransactionEventEnumType_1 = "Ended"
const TransactionEventEnumType_1_Started TransactionEventEnumType_1 = "Started"
const TransactionEventEnumType_1_Updated TransactionEventEnumType_1 = "Updated"

type TransactionEventRequestJson struct {
	// The maximum current of the connected cable in Ampere (A).
	//
	CableMaxCurrent *int `json:"cableMaxCurrent,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// EventType corresponds to the JSON schema field "eventType".
	EventType TransactionEventEnumType_1 `json:"eventType"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_14 `json:"evse,omitempty"`

	// IdToken corresponds to the JSON schema field "idToken".
	IdToken *IdTokenType_6 `json:"idToken,omitempty"`

	// MeterValue corresponds to the JSON schema field "meterValue".
	MeterValue []MeterValueType_1 `json:"meterValue,omitempty"`

	// If the Charging Station is able to report the number of phases used, then it
	// SHALL provide it. When omitted the CSMS may be able to determine the number of
	// phases used via device management.
	//
	NumberOfPhasesUsed *int `json:"numberOfPhasesUsed,omitempty"`

	// Indication that this transaction event happened when the Charging Station was
	// offline. Default = false, meaning: the event occurred when the Charging Station
	// was online.
	//
	Offline bool `json:"offline,omitempty"`

	// This contains the Id of the reservation that terminates as a result of this
	// transaction.
	//
	ReservationId *int `json:"reservationId,omitempty"`

	// Incremental sequence number, helps with determining if all messages of a
	// transaction have been received.
	//
	SeqNo int `json:"seqNo"`

	// The date and time at which this transaction event occurred.
	//
	Timestamp string `json:"timestamp"`

	// TransactionInfo corresponds to the JSON schema field "transactionInfo".
	TransactionInfo TransactionType `json:"transactionInfo"`

	// TriggerReason corresponds to the JSON schema field "triggerReason".
	TriggerReason TriggerReasonEnumType_1 `json:"triggerReason"`
}

type TransactionEventResponseJson struct {
	// Priority from a business point of view. Default priority is 0, The range is
	// from -9 to 9. Higher values indicate a higher priority. The chargingPriority in
	// &lt;&lt;transactioneventresponse,TransactionEventResponse&gt;&gt; is
	// temporarily, so it may not be set in the
	// &lt;&lt;cmn_idtokeninfotype,IdTokenInfoType&gt;&gt; afterwards. Also the
	// chargingPriority in
	// &lt;&lt;transactioneventresponse,TransactionEventResponse&gt;&gt; overrules the
	// one in &lt;&lt;cmn_idtokeninfotype,IdTokenInfoType&gt;&gt;.
	//
	ChargingPriority *int `json:"chargingPriority,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_119 `json:"customData,omitempty"`

	// IdTokenInfo corresponds to the JSON schema field "idTokenInfo".
	IdTokenInfo *IdTokenInfoType_2 `json:"idTokenInfo,omitempty"`

	// SHALL only be sent when charging has ended. Final total cost of this
	// transaction, including taxes. In the currency configured with the Configuration
	// Variable: &lt;&lt;configkey-currency,`Currency`&gt;&gt;. When omitted, the
	// transaction was NOT free. To indicate a free transaction, the CSMS SHALL send
	// 0.00.
	//
	//
	TotalCost *float64 `json:"totalCost,omitempty"`

	// UpdatedPersonalMessage corresponds to the JSON schema field
	// "updatedPersonalMessage".
	UpdatedPersonalMessage *MessageContentType_4 `json:"updatedPersonalMessage,omitempty"`
}

// Transaction
// urn:x-oca:ocpp:uid:2:233318
//
type TransactionType struct {
	// ChargingState corresponds to the JSON schema field "chargingState".
	ChargingState *ChargingStateEnumType_1 `json:"chargingState,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// The ID given to remote start request (&lt;&lt;requeststarttransactionrequest,
	// RequestStartTransactionRequest&gt;&gt;. This enables to CSMS to match the
	// started transaction to the given start request.
	//
	RemoteStartId *int `json:"remoteStartId,omitempty"`

	// StoppedReason corresponds to the JSON schema field "stoppedReason".
	StoppedReason *ReasonEnumType_1 `json:"stoppedReason,omitempty"`

	// Transaction. Time_ Spent_ Charging. Elapsed_ Time
	// urn:x-oca:ocpp:uid:1:569415
	// Contains the total time that energy flowed from EVSE to EV during the
	// transaction (in seconds). Note that timeSpentCharging is smaller or equal to
	// the duration of the transaction.
	//
	TimeSpentCharging *int `json:"timeSpentCharging,omitempty"`

	// This contains the Id of the transaction.
	//
	TransactionId string `json:"transactionId"`
}

type TriggerMessageRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_120 `json:"customData,omitempty"`

	// Evse corresponds to the JSON schema field "evse".
	Evse *EVSEType_15 `json:"evse,omitempty"`

	// RequestedMessage corresponds to the JSON schema field "requestedMessage".
	RequestedMessage MessageTriggerEnumType_1 `json:"requestedMessage"`
}

type TriggerMessageResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_121 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status TriggerMessageStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_39 `json:"statusInfo,omitempty"`
}

type TriggerMessageStatusEnumType string

const TriggerMessageStatusEnumTypeAccepted TriggerMessageStatusEnumType = "Accepted"
const TriggerMessageStatusEnumTypeNotImplemented TriggerMessageStatusEnumType = "NotImplemented"
const TriggerMessageStatusEnumTypeRejected TriggerMessageStatusEnumType = "Rejected"

type TriggerMessageStatusEnumType_1 string

const TriggerMessageStatusEnumType_1_Accepted TriggerMessageStatusEnumType_1 = "Accepted"
const TriggerMessageStatusEnumType_1_NotImplemented TriggerMessageStatusEnumType_1 = "NotImplemented"
const TriggerMessageStatusEnumType_1_Rejected TriggerMessageStatusEnumType_1 = "Rejected"

type TriggerReasonEnumType string

const TriggerReasonEnumTypeAbnormalCondition TriggerReasonEnumType = "AbnormalCondition"
const TriggerReasonEnumTypeAuthorized TriggerReasonEnumType = "Authorized"
const TriggerReasonEnumTypeCablePluggedIn TriggerReasonEnumType = "CablePluggedIn"
const TriggerReasonEnumTypeChargingRateChanged TriggerReasonEnumType = "ChargingRateChanged"
const TriggerReasonEnumTypeChargingStateChanged TriggerReasonEnumType = "ChargingStateChanged"
const TriggerReasonEnumTypeDeauthorized TriggerReasonEnumType = "Deauthorized"
const TriggerReasonEnumTypeEVCommunicationLost TriggerReasonEnumType = "EVCommunicationLost"
const TriggerReasonEnumTypeEVConnectTimeout TriggerReasonEnumType = "EVConnectTimeout"
const TriggerReasonEnumTypeEVDeparted TriggerReasonEnumType = "EVDeparted"
const TriggerReasonEnumTypeEVDetected TriggerReasonEnumType = "EVDetected"
const TriggerReasonEnumTypeEnergyLimitReached TriggerReasonEnumType = "EnergyLimitReached"
const TriggerReasonEnumTypeMeterValueClock TriggerReasonEnumType = "MeterValueClock"
const TriggerReasonEnumTypeMeterValuePeriodic TriggerReasonEnumType = "MeterValuePeriodic"
const TriggerReasonEnumTypeRemoteStart TriggerReasonEnumType = "RemoteStart"
const TriggerReasonEnumTypeRemoteStop TriggerReasonEnumType = "RemoteStop"
const TriggerReasonEnumTypeResetCommand TriggerReasonEnumType = "ResetCommand"
const TriggerReasonEnumTypeSignedDataReceived TriggerReasonEnumType = "SignedDataReceived"
const TriggerReasonEnumTypeStopAuthorized TriggerReasonEnumType = "StopAuthorized"
const TriggerReasonEnumTypeTimeLimitReached TriggerReasonEnumType = "TimeLimitReached"
const TriggerReasonEnumTypeTrigger TriggerReasonEnumType = "Trigger"
const TriggerReasonEnumTypeUnlockCommand TriggerReasonEnumType = "UnlockCommand"

type TriggerReasonEnumType_1 string

const TriggerReasonEnumType_1_AbnormalCondition TriggerReasonEnumType_1 = "AbnormalCondition"
const TriggerReasonEnumType_1_Authorized TriggerReasonEnumType_1 = "Authorized"
const TriggerReasonEnumType_1_CablePluggedIn TriggerReasonEnumType_1 = "CablePluggedIn"
const TriggerReasonEnumType_1_ChargingRateChanged TriggerReasonEnumType_1 = "ChargingRateChanged"
const TriggerReasonEnumType_1_ChargingStateChanged TriggerReasonEnumType_1 = "ChargingStateChanged"
const TriggerReasonEnumType_1_Deauthorized TriggerReasonEnumType_1 = "Deauthorized"
const TriggerReasonEnumType_1_EVCommunicationLost TriggerReasonEnumType_1 = "EVCommunicationLost"
const TriggerReasonEnumType_1_EVConnectTimeout TriggerReasonEnumType_1 = "EVConnectTimeout"
const TriggerReasonEnumType_1_EVDeparted TriggerReasonEnumType_1 = "EVDeparted"
const TriggerReasonEnumType_1_EVDetected TriggerReasonEnumType_1 = "EVDetected"
const TriggerReasonEnumType_1_EnergyLimitReached TriggerReasonEnumType_1 = "EnergyLimitReached"
const TriggerReasonEnumType_1_MeterValueClock TriggerReasonEnumType_1 = "MeterValueClock"
const TriggerReasonEnumType_1_MeterValuePeriodic TriggerReasonEnumType_1 = "MeterValuePeriodic"
const TriggerReasonEnumType_1_RemoteStart TriggerReasonEnumType_1 = "RemoteStart"
const TriggerReasonEnumType_1_RemoteStop TriggerReasonEnumType_1 = "RemoteStop"
const TriggerReasonEnumType_1_ResetCommand TriggerReasonEnumType_1 = "ResetCommand"
const TriggerReasonEnumType_1_SignedDataReceived TriggerReasonEnumType_1 = "SignedDataReceived"
const TriggerReasonEnumType_1_StopAuthorized TriggerReasonEnumType_1 = "StopAuthorized"
const TriggerReasonEnumType_1_TimeLimitReached TriggerReasonEnumType_1 = "TimeLimitReached"
const TriggerReasonEnumType_1_Trigger TriggerReasonEnumType_1 = "Trigger"
const TriggerReasonEnumType_1_UnlockCommand TriggerReasonEnumType_1 = "UnlockCommand"

// Represents a UnitOfMeasure with a multiplier
//
type UnitOfMeasureType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_62 `json:"customData,omitempty"`

	// Multiplier, this value represents the exponent to base 10. I.e. multiplier 3
	// means 10 raised to the 3rd power. Default is 0.
	//
	Multiplier int `json:"multiplier,omitempty"`

	// Unit of the value. Default = "Wh" if the (default) measurand is an "Energy"
	// type.
	// This field SHALL use a value from the list Standardized Units of Measurements
	// in Part 2 Appendices.
	// If an applicable unit is available in that list, otherwise a "custom" unit
	// might be used.
	//
	Unit string `json:"unit,omitempty"`
}

// Represents a UnitOfMeasure with a multiplier
//
type UnitOfMeasureType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_118 `json:"customData,omitempty"`

	// Multiplier, this value represents the exponent to base 10. I.e. multiplier 3
	// means 10 raised to the 3rd power. Default is 0.
	//
	Multiplier int `json:"multiplier,omitempty"`

	// Unit of the value. Default = "Wh" if the (default) measurand is an "Energy"
	// type.
	// This field SHALL use a value from the list Standardized Units of Measurements
	// in Part 2 Appendices.
	// If an applicable unit is available in that list, otherwise a "custom" unit
	// might be used.
	//
	Unit string `json:"unit,omitempty"`
}

type UnlockConnectorRequestJson struct {
	// This contains the identifier of the connector that needs to be unlocked.
	//
	ConnectorId int `json:"connectorId"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_122 `json:"customData,omitempty"`

	// This contains the identifier of the EVSE for which a connector needs to be
	// unlocked.
	//
	EvseId int `json:"evseId"`
}

type UnlockConnectorResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_123 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status UnlockStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_40 `json:"statusInfo,omitempty"`
}

type UnlockStatusEnumType string

const UnlockStatusEnumTypeOngoingAuthorizedTransaction UnlockStatusEnumType = "OngoingAuthorizedTransaction"
const UnlockStatusEnumTypeUnknownConnector UnlockStatusEnumType = "UnknownConnector"
const UnlockStatusEnumTypeUnlockFailed UnlockStatusEnumType = "UnlockFailed"
const UnlockStatusEnumTypeUnlocked UnlockStatusEnumType = "Unlocked"

type UnlockStatusEnumType_1 string

const UnlockStatusEnumType_1_OngoingAuthorizedTransaction UnlockStatusEnumType_1 = "OngoingAuthorizedTransaction"
const UnlockStatusEnumType_1_UnknownConnector UnlockStatusEnumType_1 = "UnknownConnector"
const UnlockStatusEnumType_1_UnlockFailed UnlockStatusEnumType_1 = "UnlockFailed"
const UnlockStatusEnumType_1_Unlocked UnlockStatusEnumType_1 = "Unlocked"

type UnpublishFirmwareRequestJson struct {
	// The MD5 checksum over the entire firmware file as a hexadecimal string of
	// length 32.
	//
	Checksum string `json:"checksum"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_124 `json:"customData,omitempty"`
}

type UnpublishFirmwareResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_125 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status UnpublishFirmwareStatusEnumType_1 `json:"status"`
}

type UnpublishFirmwareStatusEnumType string

const UnpublishFirmwareStatusEnumTypeDownloadOngoing UnpublishFirmwareStatusEnumType = "DownloadOngoing"
const UnpublishFirmwareStatusEnumTypeNoFirmware UnpublishFirmwareStatusEnumType = "NoFirmware"
const UnpublishFirmwareStatusEnumTypeUnpublished UnpublishFirmwareStatusEnumType = "Unpublished"

type UnpublishFirmwareStatusEnumType_1 string

const UnpublishFirmwareStatusEnumType_1_DownloadOngoing UnpublishFirmwareStatusEnumType_1 = "DownloadOngoing"
const UnpublishFirmwareStatusEnumType_1_NoFirmware UnpublishFirmwareStatusEnumType_1 = "NoFirmware"
const UnpublishFirmwareStatusEnumType_1_Unpublished UnpublishFirmwareStatusEnumType_1 = "Unpublished"

type UpdateEnumType string

const UpdateEnumTypeDifferential UpdateEnumType = "Differential"
const UpdateEnumTypeFull UpdateEnumType = "Full"

type UpdateEnumType_1 string

const UpdateEnumType_1_Differential UpdateEnumType_1 = "Differential"
const UpdateEnumType_1_Full UpdateEnumType_1 = "Full"

type UpdateFirmwareRequestJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_126 `json:"customData,omitempty"`

	// Firmware corresponds to the JSON schema field "firmware".
	Firmware FirmwareType `json:"firmware"`

	// The Id of this request
	//
	RequestId int `json:"requestId"`

	// This specifies how many times Charging Station must try to download the
	// firmware before giving up. If this field is not present, it is left to Charging
	// Station to decide how many times it wants to retry.
	//
	Retries *int `json:"retries,omitempty"`

	// The interval in seconds after which a retry may be attempted. If this field is
	// not present, it is left to Charging Station to decide how long to wait between
	// attempts.
	//
	RetryInterval *int `json:"retryInterval,omitempty"`
}

type UpdateFirmwareResponseJson struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_127 `json:"customData,omitempty"`

	// Status corresponds to the JSON schema field "status".
	Status UpdateFirmwareStatusEnumType_1 `json:"status"`

	// StatusInfo corresponds to the JSON schema field "statusInfo".
	StatusInfo *StatusInfoType_41 `json:"statusInfo,omitempty"`
}

type UpdateFirmwareStatusEnumType string

const UpdateFirmwareStatusEnumTypeAccepted UpdateFirmwareStatusEnumType = "Accepted"
const UpdateFirmwareStatusEnumTypeAcceptedCanceled UpdateFirmwareStatusEnumType = "AcceptedCanceled"
const UpdateFirmwareStatusEnumTypeInvalidCertificate UpdateFirmwareStatusEnumType = "InvalidCertificate"
const UpdateFirmwareStatusEnumTypeRejected UpdateFirmwareStatusEnumType = "Rejected"
const UpdateFirmwareStatusEnumTypeRevokedCertificate UpdateFirmwareStatusEnumType = "RevokedCertificate"

type UpdateFirmwareStatusEnumType_1 string

const UpdateFirmwareStatusEnumType_1_Accepted UpdateFirmwareStatusEnumType_1 = "Accepted"
const UpdateFirmwareStatusEnumType_1_AcceptedCanceled UpdateFirmwareStatusEnumType_1 = "AcceptedCanceled"
const UpdateFirmwareStatusEnumType_1_InvalidCertificate UpdateFirmwareStatusEnumType_1 = "InvalidCertificate"
const UpdateFirmwareStatusEnumType_1_Rejected UpdateFirmwareStatusEnumType_1 = "Rejected"
const UpdateFirmwareStatusEnumType_1_RevokedCertificate UpdateFirmwareStatusEnumType_1 = "RevokedCertificate"

type UploadLogStatusEnumType string

const UploadLogStatusEnumTypeAcceptedCanceled UploadLogStatusEnumType = "AcceptedCanceled"
const UploadLogStatusEnumTypeBadMessage UploadLogStatusEnumType = "BadMessage"
const UploadLogStatusEnumTypeIdle UploadLogStatusEnumType = "Idle"
const UploadLogStatusEnumTypeNotSupportedOperation UploadLogStatusEnumType = "NotSupportedOperation"
const UploadLogStatusEnumTypePermissionDenied UploadLogStatusEnumType = "PermissionDenied"
const UploadLogStatusEnumTypeUploadFailure UploadLogStatusEnumType = "UploadFailure"
const UploadLogStatusEnumTypeUploaded UploadLogStatusEnumType = "Uploaded"
const UploadLogStatusEnumTypeUploading UploadLogStatusEnumType = "Uploading"

type UploadLogStatusEnumType_1 string

const UploadLogStatusEnumType_1_AcceptedCanceled UploadLogStatusEnumType_1 = "AcceptedCanceled"
const UploadLogStatusEnumType_1_BadMessage UploadLogStatusEnumType_1 = "BadMessage"
const UploadLogStatusEnumType_1_Idle UploadLogStatusEnumType_1 = "Idle"
const UploadLogStatusEnumType_1_NotSupportedOperation UploadLogStatusEnumType_1 = "NotSupportedOperation"
const UploadLogStatusEnumType_1_PermissionDenied UploadLogStatusEnumType_1 = "PermissionDenied"
const UploadLogStatusEnumType_1_UploadFailure UploadLogStatusEnumType_1 = "UploadFailure"
const UploadLogStatusEnumType_1_Uploaded UploadLogStatusEnumType_1 = "Uploaded"
const UploadLogStatusEnumType_1_Uploading UploadLogStatusEnumType_1 = "Uploading"

type VPNEnumType string

const VPNEnumTypeIKEv2 VPNEnumType = "IKEv2"
const VPNEnumTypeIPSec VPNEnumType = "IPSec"
const VPNEnumTypeL2TP VPNEnumType = "L2TP"
const VPNEnumTypePPTP VPNEnumType = "PPTP"

type VPNEnumType_1 string

const VPNEnumType_1_IKEv2 VPNEnumType_1 = "IKEv2"
const VPNEnumType_1_IPSec VPNEnumType_1 = "IPSec"
const VPNEnumType_1_L2TP VPNEnumType_1 = "L2TP"
const VPNEnumType_1_PPTP VPNEnumType_1 = "PPTP"

// VPN
// urn:x-oca:ocpp:uid:2:233268
// VPN Configuration settings
//
type VPNType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_108 `json:"customData,omitempty"`

	// VPN. Group. Group_ Name
	// urn:x-oca:ocpp:uid:1:569274
	// VPN group.
	//
	Group *string `json:"group,omitempty"`

	// VPN. Key. VPN_ Key
	// urn:x-oca:ocpp:uid:1:569276
	// VPN shared secret.
	//
	Key string `json:"key"`

	// VPN. Password. Password
	// urn:x-oca:ocpp:uid:1:569275
	// VPN Password.
	//
	Password string `json:"password"`

	// VPN. Server. URI
	// urn:x-oca:ocpp:uid:1:569272
	// VPN Server Address
	//
	Server string `json:"server"`

	// Type corresponds to the JSON schema field "type".
	Type VPNEnumType `json:"type"`

	// VPN. User. User_ Name
	// urn:x-oca:ocpp:uid:1:569273
	// VPN User
	//
	User string `json:"user"`
}

// Attribute data of a variable.
//
type VariableAttributeType struct {
	// If true, value that will never be changed by the Charging Station at runtime.
	// Default when omitted is false.
	//
	Constant bool `json:"constant,omitempty"`

	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_78 `json:"customData,omitempty"`

	// Mutability corresponds to the JSON schema field "mutability".
	Mutability *MutabilityEnumType_1 `json:"mutability,omitempty"`

	// If true, value will be persistent across system reboots or power down. Default
	// when omitted is false.
	//
	Persistent bool `json:"persistent,omitempty"`

	// Type corresponds to the JSON schema field "type".
	Type *AttributeEnumType_5 `json:"type,omitempty"`

	// Value of the attribute. May only be omitted when mutability is set to
	// 'WriteOnly'.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-reporting-value-size,ReportingValueSize&gt;&gt; can be used
	// to limit GetVariableResult.attributeValue, VariableAttribute.value and
	// EventData.actualValue. The max size of these values will always remain equal.
	//
	Value *string `json:"value,omitempty"`
}

// Fixed read-only parameters of a variable.
//
type VariableCharacteristicsType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_78 `json:"customData,omitempty"`

	// DataType corresponds to the JSON schema field "dataType".
	DataType DataEnumType_1 `json:"dataType"`

	// Maximum possible value of this variable. When the datatype of this Variable is
	// String, OptionList, SequenceList or MemberList, this field defines the maximum
	// length of the (CSV) string.
	//
	MaxLimit *float64 `json:"maxLimit,omitempty"`

	// Minimum possible value of this variable.
	//
	MinLimit *float64 `json:"minLimit,omitempty"`

	// Flag indicating if this variable supports monitoring.
	//
	SupportsMonitoring bool `json:"supportsMonitoring"`

	// Unit of the variable. When the transmitted value has a unit, this field SHALL
	// be included.
	//
	Unit *string `json:"unit,omitempty"`

	// Allowed values when variable is Option/Member/SequenceList.
	//
	// * OptionList: The (Actual) Variable value must be a single value from the
	// reported (CSV) enumeration list.
	//
	// * MemberList: The (Actual) Variable value  may be an (unordered) (sub-)set of
	// the reported (CSV) valid values list.
	//
	// * SequenceList: The (Actual) Variable value  may be an ordered (priority, etc)
	// (sub-)set of the reported (CSV) valid values.
	//
	// This is a comma separated list.
	//
	// The Configuration Variable
	// &lt;&lt;configkey-configuration-value-size,ConfigurationValueSize&gt;&gt; can
	// be used to limit SetVariableData.attributeValue and
	// VariableCharacteristics.valueList. The max size of these values will always
	// remain equal.
	//
	//
	ValuesList *string `json:"valuesList,omitempty"`
}

// A monitoring setting for a variable.
//
type VariableMonitoringType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_76 `json:"customData,omitempty"`

	// Identifies the monitor.
	//
	Id int `json:"id"`

	// The severity that will be assigned to an event that is triggered by this
	// monitor. The severity range is 0-9, with 0 as the highest and 9 as the lowest
	// severity level.
	//
	// The severity levels have the following meaning: +
	// *0-Danger* +
	// Indicates lives are potentially in danger. Urgent attention is needed and
	// action should be taken immediately. +
	// *1-Hardware Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to Hardware issues. Action is required. +
	// *2-System Failure* +
	// Indicates that the Charging Station is unable to continue regular operations
	// due to software or minor hardware issues. Action is required. +
	// *3-Critical* +
	// Indicates a critical error. Action is required. +
	// *4-Error* +
	// Indicates a non-urgent error. Action is required. +
	// *5-Alert* +
	// Indicates an alert event. Default severity for any type of monitoring event.  +
	// *6-Warning* +
	// Indicates a warning event. Action may be required. +
	// *7-Notice* +
	// Indicates an unusual event. No immediate action is required. +
	// *8-Informational* +
	// Indicates a regular operational event. May be used for reporting, measuring
	// throughput, etc. No action is required. +
	// *9-Debug* +
	// Indicates information useful to developers for debugging, not useful during
	// operations.
	//
	Severity int `json:"severity"`

	// Monitor only active when a transaction is ongoing on a component relevant to
	// this transaction.
	//
	Transaction bool `json:"transaction"`

	// Type corresponds to the JSON schema field "type".
	Type MonitorEnumType_1 `json:"type"`

	// Value for threshold or delta monitoring.
	// For Periodic or PeriodicClockAligned this is the interval in seconds.
	//
	Value float64 `json:"value"`
}

// Reference key to a component-variable.
//
type VariableType struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_48 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_1 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_50 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_10 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_113 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_2 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_54 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_3 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_55 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_4 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_74 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_5 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_76 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_6 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_78 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_7 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_110 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_8 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_111 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

// Reference key to a component-variable.
//
type VariableType_9 struct {
	// CustomData corresponds to the JSON schema field "customData".
	CustomData *CustomDataType_112 `json:"customData,omitempty"`

	// Name of instance in case the variable exists as multiple instances. Case
	// Insensitive. strongly advised to use Camel Case.
	//
	Instance *string `json:"instance,omitempty"`

	// Name of the variable. Name should be taken from the list of standardized
	// variable names whenever possible. Case Insensitive. strongly advised to use
	// Camel Case.
	//
	Name string `json:"name"`
}

var enumValues_APNAuthenticationEnumType = []interface{}{
	"CHAP",
	"NONE",
	"PAP",
	"AUTO",
}
var enumValues_APNAuthenticationEnumType_1 = []interface{}{
	"CHAP",
	"NONE",
	"PAP",
	"AUTO",
}
var enumValues_AttributeEnumType = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_1 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_2 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_3 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_4 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_5 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_6 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_7 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_8 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AttributeEnumType_9 = []interface{}{
	"Actual",
	"Target",
	"MinSet",
	"MaxSet",
}
var enumValues_AuthorizationStatusEnumType = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_AuthorizationStatusEnumType_1 = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_AuthorizationStatusEnumType_2 = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_AuthorizationStatusEnumType_3 = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_AuthorizationStatusEnumType_4 = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_AuthorizationStatusEnumType_5 = []interface{}{
	"Accepted",
	"Blocked",
	"ConcurrentTx",
	"Expired",
	"Invalid",
	"NoCredit",
	"NotAllowedTypeEVSE",
	"NotAtThisLocation",
	"NotAtThisTime",
	"Unknown",
}
var enumValues_AuthorizeCertificateStatusEnumType = []interface{}{
	"Accepted",
	"SignatureError",
	"CertificateExpired",
	"CertificateRevoked",
	"NoCertificateAvailable",
	"CertChainError",
	"ContractCancelled",
}
var enumValues_AuthorizeCertificateStatusEnumType_1 = []interface{}{
	"Accepted",
	"SignatureError",
	"CertificateExpired",
	"CertificateRevoked",
	"NoCertificateAvailable",
	"CertChainError",
	"ContractCancelled",
}
var enumValues_BootReasonEnumType = []interface{}{
	"ApplicationReset",
	"FirmwareUpdate",
	"LocalReset",
	"PowerUp",
	"RemoteReset",
	"ScheduledReset",
	"Triggered",
	"Unknown",
	"Watchdog",
}
var enumValues_BootReasonEnumType_1 = []interface{}{
	"ApplicationReset",
	"FirmwareUpdate",
	"LocalReset",
	"PowerUp",
	"RemoteReset",
	"ScheduledReset",
	"Triggered",
	"Unknown",
	"Watchdog",
}
var enumValues_CancelReservationStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_CancelReservationStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_CertificateActionEnumType = []interface{}{
	"Install",
	"Update",
}
var enumValues_CertificateActionEnumType_1 = []interface{}{
	"Install",
	"Update",
}
var enumValues_CertificateSignedStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_CertificateSignedStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_CertificateSigningUseEnumType = []interface{}{
	"ChargingStationCertificate",
	"V2GCertificate",
}
var enumValues_CertificateSigningUseEnumType_1 = []interface{}{
	"ChargingStationCertificate",
	"V2GCertificate",
}
var enumValues_CertificateSigningUseEnumType_2 = []interface{}{
	"ChargingStationCertificate",
	"V2GCertificate",
}
var enumValues_CertificateSigningUseEnumType_3 = []interface{}{
	"ChargingStationCertificate",
	"V2GCertificate",
}
var enumValues_ChangeAvailabilityStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"Scheduled",
}
var enumValues_ChangeAvailabilityStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"Scheduled",
}
var enumValues_ChargingLimitSourceEnumType = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingLimitSourceEnumType_1 = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingLimitSourceEnumType_2 = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingLimitSourceEnumType_3 = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingLimitSourceEnumType_4 = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingLimitSourceEnumType_5 = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingLimitSourceEnumType_6 = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingLimitSourceEnumType_7 = []interface{}{
	"EMS",
	"Other",
	"SO",
	"CSO",
}
var enumValues_ChargingProfileKindEnumType = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfileKindEnumType_1 = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfileKindEnumType_2 = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfileKindEnumType_3 = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfileKindEnumType_4 = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfileKindEnumType_5 = []interface{}{
	"Absolute",
	"Recurring",
	"Relative",
}
var enumValues_ChargingProfilePurposeEnumType = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_1 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_2 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_3 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_4 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_5 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_6 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_7 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_8 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfilePurposeEnumType_9 = []interface{}{
	"ChargingStationExternalConstraints",
	"ChargingStationMaxProfile",
	"TxDefaultProfile",
	"TxProfile",
}
var enumValues_ChargingProfileStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_ChargingProfileStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_ChargingRateUnitEnumType = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_1 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_10 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_11 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_12 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_13 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_2 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_3 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_4 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_5 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_6 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_7 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_8 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingRateUnitEnumType_9 = []interface{}{
	"W",
	"A",
}
var enumValues_ChargingStateEnumType = []interface{}{
	"Charging",
	"EVConnected",
	"SuspendedEV",
	"SuspendedEVSE",
	"Idle",
}
var enumValues_ChargingStateEnumType_1 = []interface{}{
	"Charging",
	"EVConnected",
	"SuspendedEV",
	"SuspendedEVSE",
	"Idle",
}
var enumValues_ClearCacheStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_ClearCacheStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_ClearChargingProfileStatusEnumType = []interface{}{
	"Accepted",
	"Unknown",
}
var enumValues_ClearChargingProfileStatusEnumType_1 = []interface{}{
	"Accepted",
	"Unknown",
}
var enumValues_ClearMessageStatusEnumType = []interface{}{
	"Accepted",
	"Unknown",
}
var enumValues_ClearMessageStatusEnumType_1 = []interface{}{
	"Accepted",
	"Unknown",
}
var enumValues_ClearMonitoringStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"NotFound",
}
var enumValues_ClearMonitoringStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"NotFound",
}
var enumValues_ComponentCriterionEnumType = []interface{}{
	"Active",
	"Available",
	"Enabled",
	"Problem",
}
var enumValues_ComponentCriterionEnumType_1 = []interface{}{
	"Active",
	"Available",
	"Enabled",
	"Problem",
}
var enumValues_ConnectorEnumType = []interface{}{
	"cCCS1",
	"cCCS2",
	"cG105",
	"cTesla",
	"cType1",
	"cType2",
	"s309-1P-16A",
	"s309-1P-32A",
	"s309-3P-16A",
	"s309-3P-32A",
	"sBS1361",
	"sCEE-7-7",
	"sType2",
	"sType3",
	"Other1PhMax16A",
	"Other1PhOver16A",
	"Other3Ph",
	"Pan",
	"wInductive",
	"wResonant",
	"Undetermined",
	"Unknown",
}
var enumValues_ConnectorEnumType_1 = []interface{}{
	"cCCS1",
	"cCCS2",
	"cG105",
	"cTesla",
	"cType1",
	"cType2",
	"s309-1P-16A",
	"s309-1P-32A",
	"s309-3P-16A",
	"s309-3P-32A",
	"sBS1361",
	"sCEE-7-7",
	"sType2",
	"sType3",
	"Other1PhMax16A",
	"Other1PhOver16A",
	"Other3Ph",
	"Pan",
	"wInductive",
	"wResonant",
	"Undetermined",
	"Unknown",
}
var enumValues_ConnectorStatusEnumType = []interface{}{
	"Available",
	"Occupied",
	"Reserved",
	"Unavailable",
	"Faulted",
}
var enumValues_ConnectorStatusEnumType_1 = []interface{}{
	"Available",
	"Occupied",
	"Reserved",
	"Unavailable",
	"Faulted",
}
var enumValues_CostKindEnumType = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_1 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_2 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_3 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_4 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_5 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_6 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_7 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_8 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CostKindEnumType_9 = []interface{}{
	"CarbonDioxideEmission",
	"RelativePricePercentage",
	"RenewableGenerationPercentage",
}
var enumValues_CustomerInformationStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"Invalid",
}
var enumValues_CustomerInformationStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"Invalid",
}
var enumValues_DataEnumType = []interface{}{
	"string",
	"decimal",
	"integer",
	"dateTime",
	"boolean",
	"OptionList",
	"SequenceList",
	"MemberList",
}
var enumValues_DataEnumType_1 = []interface{}{
	"string",
	"decimal",
	"integer",
	"dateTime",
	"boolean",
	"OptionList",
	"SequenceList",
	"MemberList",
}
var enumValues_DataTransferStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownMessageId",
	"UnknownVendorId",
}
var enumValues_DataTransferStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownMessageId",
	"UnknownVendorId",
}
var enumValues_DeleteCertificateStatusEnumType = []interface{}{
	"Accepted",
	"Failed",
	"NotFound",
}
var enumValues_DeleteCertificateStatusEnumType_1 = []interface{}{
	"Accepted",
	"Failed",
	"NotFound",
}
var enumValues_DisplayMessageStatusEnumType = []interface{}{
	"Accepted",
	"NotSupportedMessageFormat",
	"Rejected",
	"NotSupportedPriority",
	"NotSupportedState",
	"UnknownTransaction",
}
var enumValues_DisplayMessageStatusEnumType_1 = []interface{}{
	"Accepted",
	"NotSupportedMessageFormat",
	"Rejected",
	"NotSupportedPriority",
	"NotSupportedState",
	"UnknownTransaction",
}
var enumValues_EnergyTransferModeEnumType = []interface{}{
	"DC",
	"AC_single_phase",
	"AC_two_phase",
	"AC_three_phase",
}
var enumValues_EnergyTransferModeEnumType_1 = []interface{}{
	"DC",
	"AC_single_phase",
	"AC_two_phase",
	"AC_three_phase",
}
var enumValues_EventNotificationEnumType = []interface{}{
	"HardWiredNotification",
	"HardWiredMonitor",
	"PreconfiguredMonitor",
	"CustomMonitor",
}
var enumValues_EventNotificationEnumType_1 = []interface{}{
	"HardWiredNotification",
	"HardWiredMonitor",
	"PreconfiguredMonitor",
	"CustomMonitor",
}
var enumValues_EventTriggerEnumType = []interface{}{
	"Alerting",
	"Delta",
	"Periodic",
}
var enumValues_EventTriggerEnumType_1 = []interface{}{
	"Alerting",
	"Delta",
	"Periodic",
}
var enumValues_FirmwareStatusEnumType = []interface{}{
	"Downloaded",
	"DownloadFailed",
	"Downloading",
	"DownloadScheduled",
	"DownloadPaused",
	"Idle",
	"InstallationFailed",
	"Installing",
	"Installed",
	"InstallRebooting",
	"InstallScheduled",
	"InstallVerificationFailed",
	"InvalidSignature",
	"SignatureVerified",
}
var enumValues_FirmwareStatusEnumType_1 = []interface{}{
	"Downloaded",
	"DownloadFailed",
	"Downloading",
	"DownloadScheduled",
	"DownloadPaused",
	"Idle",
	"InstallationFailed",
	"Installing",
	"Installed",
	"InstallRebooting",
	"InstallScheduled",
	"InstallVerificationFailed",
	"InvalidSignature",
	"SignatureVerified",
}
var enumValues_GenericDeviceModelStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericDeviceModelStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericDeviceModelStatusEnumType_2 = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericDeviceModelStatusEnumType_3 = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericDeviceModelStatusEnumType_4 = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericDeviceModelStatusEnumType_5 = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericDeviceModelStatusEnumType_6 = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericDeviceModelStatusEnumType_7 = []interface{}{
	"Accepted",
	"Rejected",
	"NotSupported",
	"EmptyResultSet",
}
var enumValues_GenericStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_2 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_3 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_4 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_5 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_6 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_7 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_8 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GenericStatusEnumType_9 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_GetCertificateIdUseEnumType = []interface{}{
	"V2GRootCertificate",
	"MORootCertificate",
	"CSMSRootCertificate",
	"V2GCertificateChain",
	"ManufacturerRootCertificate",
}
var enumValues_GetCertificateIdUseEnumType_1 = []interface{}{
	"V2GRootCertificate",
	"MORootCertificate",
	"CSMSRootCertificate",
	"V2GCertificateChain",
	"ManufacturerRootCertificate",
}
var enumValues_GetCertificateIdUseEnumType_2 = []interface{}{
	"V2GRootCertificate",
	"MORootCertificate",
	"CSMSRootCertificate",
	"V2GCertificateChain",
	"ManufacturerRootCertificate",
}
var enumValues_GetCertificateIdUseEnumType_3 = []interface{}{
	"V2GRootCertificate",
	"MORootCertificate",
	"CSMSRootCertificate",
	"V2GCertificateChain",
	"ManufacturerRootCertificate",
}
var enumValues_GetCertificateStatusEnumType = []interface{}{
	"Accepted",
	"Failed",
}
var enumValues_GetCertificateStatusEnumType_1 = []interface{}{
	"Accepted",
	"Failed",
}
var enumValues_GetChargingProfileStatusEnumType = []interface{}{
	"Accepted",
	"NoProfiles",
}
var enumValues_GetChargingProfileStatusEnumType_1 = []interface{}{
	"Accepted",
	"NoProfiles",
}
var enumValues_GetDisplayMessagesStatusEnumType = []interface{}{
	"Accepted",
	"Unknown",
}
var enumValues_GetDisplayMessagesStatusEnumType_1 = []interface{}{
	"Accepted",
	"Unknown",
}
var enumValues_GetInstalledCertificateStatusEnumType = []interface{}{
	"Accepted",
	"NotFound",
}
var enumValues_GetInstalledCertificateStatusEnumType_1 = []interface{}{
	"Accepted",
	"NotFound",
}
var enumValues_GetVariableStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownComponent",
	"UnknownVariable",
	"NotSupportedAttributeType",
}
var enumValues_GetVariableStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownComponent",
	"UnknownVariable",
	"NotSupportedAttributeType",
}
var enumValues_HashAlgorithmEnumType = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_1 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_2 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_3 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_4 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_5 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_6 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_7 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_8 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_HashAlgorithmEnumType_9 = []interface{}{
	"SHA256",
	"SHA384",
	"SHA512",
}
var enumValues_IdTokenEnumType = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_1 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_10 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_11 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_12 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_13 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_14 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_15 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_2 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_3 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_4 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_5 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_6 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_7 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_8 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_IdTokenEnumType_9 = []interface{}{
	"Central",
	"eMAID",
	"ISO14443",
	"ISO15693",
	"KeyCode",
	"Local",
	"MacAddress",
	"NoAuthorization",
}
var enumValues_InstallCertificateStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"Failed",
}
var enumValues_InstallCertificateStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"Failed",
}
var enumValues_InstallCertificateUseEnumType = []interface{}{
	"V2GRootCertificate",
	"MORootCertificate",
	"CSMSRootCertificate",
	"ManufacturerRootCertificate",
}
var enumValues_InstallCertificateUseEnumType_1 = []interface{}{
	"V2GRootCertificate",
	"MORootCertificate",
	"CSMSRootCertificate",
	"ManufacturerRootCertificate",
}
var enumValues_Iso15118EVCertificateStatusEnumType = []interface{}{
	"Accepted",
	"Failed",
}
var enumValues_Iso15118EVCertificateStatusEnumType_1 = []interface{}{
	"Accepted",
	"Failed",
}
var enumValues_LocationEnumType = []interface{}{
	"Body",
	"Cable",
	"EV",
	"Inlet",
	"Outlet",
}
var enumValues_LocationEnumType_1 = []interface{}{
	"Body",
	"Cable",
	"EV",
	"Inlet",
	"Outlet",
}
var enumValues_LocationEnumType_2 = []interface{}{
	"Body",
	"Cable",
	"EV",
	"Inlet",
	"Outlet",
}
var enumValues_LocationEnumType_3 = []interface{}{
	"Body",
	"Cable",
	"EV",
	"Inlet",
	"Outlet",
}
var enumValues_LogEnumType = []interface{}{
	"DiagnosticsLog",
	"SecurityLog",
}
var enumValues_LogEnumType_1 = []interface{}{
	"DiagnosticsLog",
	"SecurityLog",
}
var enumValues_LogStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"AcceptedCanceled",
}
var enumValues_LogStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"AcceptedCanceled",
}
var enumValues_MeasurandEnumType = []interface{}{
	"Current.Export",
	"Current.Import",
	"Current.Offered",
	"Energy.Active.Export.Register",
	"Energy.Active.Import.Register",
	"Energy.Reactive.Export.Register",
	"Energy.Reactive.Import.Register",
	"Energy.Active.Export.Interval",
	"Energy.Active.Import.Interval",
	"Energy.Active.Net",
	"Energy.Reactive.Export.Interval",
	"Energy.Reactive.Import.Interval",
	"Energy.Reactive.Net",
	"Energy.Apparent.Net",
	"Energy.Apparent.Import",
	"Energy.Apparent.Export",
	"Frequency",
	"Power.Active.Export",
	"Power.Active.Import",
	"Power.Factor",
	"Power.Offered",
	"Power.Reactive.Export",
	"Power.Reactive.Import",
	"SoC",
	"Voltage",
}
var enumValues_MeasurandEnumType_1 = []interface{}{
	"Current.Export",
	"Current.Import",
	"Current.Offered",
	"Energy.Active.Export.Register",
	"Energy.Active.Import.Register",
	"Energy.Reactive.Export.Register",
	"Energy.Reactive.Import.Register",
	"Energy.Active.Export.Interval",
	"Energy.Active.Import.Interval",
	"Energy.Active.Net",
	"Energy.Reactive.Export.Interval",
	"Energy.Reactive.Import.Interval",
	"Energy.Reactive.Net",
	"Energy.Apparent.Net",
	"Energy.Apparent.Import",
	"Energy.Apparent.Export",
	"Frequency",
	"Power.Active.Export",
	"Power.Active.Import",
	"Power.Factor",
	"Power.Offered",
	"Power.Reactive.Export",
	"Power.Reactive.Import",
	"SoC",
	"Voltage",
}
var enumValues_MeasurandEnumType_2 = []interface{}{
	"Current.Export",
	"Current.Import",
	"Current.Offered",
	"Energy.Active.Export.Register",
	"Energy.Active.Import.Register",
	"Energy.Reactive.Export.Register",
	"Energy.Reactive.Import.Register",
	"Energy.Active.Export.Interval",
	"Energy.Active.Import.Interval",
	"Energy.Active.Net",
	"Energy.Reactive.Export.Interval",
	"Energy.Reactive.Import.Interval",
	"Energy.Reactive.Net",
	"Energy.Apparent.Net",
	"Energy.Apparent.Import",
	"Energy.Apparent.Export",
	"Frequency",
	"Power.Active.Export",
	"Power.Active.Import",
	"Power.Factor",
	"Power.Offered",
	"Power.Reactive.Export",
	"Power.Reactive.Import",
	"SoC",
	"Voltage",
}
var enumValues_MeasurandEnumType_3 = []interface{}{
	"Current.Export",
	"Current.Import",
	"Current.Offered",
	"Energy.Active.Export.Register",
	"Energy.Active.Import.Register",
	"Energy.Reactive.Export.Register",
	"Energy.Reactive.Import.Register",
	"Energy.Active.Export.Interval",
	"Energy.Active.Import.Interval",
	"Energy.Active.Net",
	"Energy.Reactive.Export.Interval",
	"Energy.Reactive.Import.Interval",
	"Energy.Reactive.Net",
	"Energy.Apparent.Net",
	"Energy.Apparent.Import",
	"Energy.Apparent.Export",
	"Frequency",
	"Power.Active.Export",
	"Power.Active.Import",
	"Power.Factor",
	"Power.Offered",
	"Power.Reactive.Export",
	"Power.Reactive.Import",
	"SoC",
	"Voltage",
}
var enumValues_MessageFormatEnumType = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_1 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_2 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_3 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_4 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_5 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_6 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_7 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_8 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessageFormatEnumType_9 = []interface{}{
	"ASCII",
	"HTML",
	"URI",
	"UTF8",
}
var enumValues_MessagePriorityEnumType = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessagePriorityEnumType_1 = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessagePriorityEnumType_2 = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessagePriorityEnumType_3 = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessagePriorityEnumType_4 = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessagePriorityEnumType_5 = []interface{}{
	"AlwaysFront",
	"InFront",
	"NormalCycle",
}
var enumValues_MessageStateEnumType = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}
var enumValues_MessageStateEnumType_1 = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}
var enumValues_MessageStateEnumType_2 = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}
var enumValues_MessageStateEnumType_3 = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}
var enumValues_MessageStateEnumType_4 = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}
var enumValues_MessageStateEnumType_5 = []interface{}{
	"Charging",
	"Faulted",
	"Idle",
	"Unavailable",
}
var enumValues_MessageTriggerEnumType = []interface{}{
	"BootNotification",
	"LogStatusNotification",
	"FirmwareStatusNotification",
	"Heartbeat",
	"MeterValues",
	"SignChargingStationCertificate",
	"SignV2GCertificate",
	"StatusNotification",
	"TransactionEvent",
	"SignCombinedCertificate",
	"PublishFirmwareStatusNotification",
}
var enumValues_MessageTriggerEnumType_1 = []interface{}{
	"BootNotification",
	"LogStatusNotification",
	"FirmwareStatusNotification",
	"Heartbeat",
	"MeterValues",
	"SignChargingStationCertificate",
	"SignV2GCertificate",
	"StatusNotification",
	"TransactionEvent",
	"SignCombinedCertificate",
	"PublishFirmwareStatusNotification",
}
var enumValues_MonitorEnumType = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitorEnumType_1 = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitorEnumType_2 = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitorEnumType_3 = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitorEnumType_4 = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitorEnumType_5 = []interface{}{
	"UpperThreshold",
	"LowerThreshold",
	"Delta",
	"Periodic",
	"PeriodicClockAligned",
}
var enumValues_MonitoringBaseEnumType = []interface{}{
	"All",
	"FactoryDefault",
	"HardWiredOnly",
}
var enumValues_MonitoringBaseEnumType_1 = []interface{}{
	"All",
	"FactoryDefault",
	"HardWiredOnly",
}
var enumValues_MonitoringCriterionEnumType = []interface{}{
	"ThresholdMonitoring",
	"DeltaMonitoring",
	"PeriodicMonitoring",
}
var enumValues_MonitoringCriterionEnumType_1 = []interface{}{
	"ThresholdMonitoring",
	"DeltaMonitoring",
	"PeriodicMonitoring",
}
var enumValues_MutabilityEnumType = []interface{}{
	"ReadOnly",
	"WriteOnly",
	"ReadWrite",
}
var enumValues_MutabilityEnumType_1 = []interface{}{
	"ReadOnly",
	"WriteOnly",
	"ReadWrite",
}
var enumValues_NotifyEVChargingNeedsStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"Processing",
}
var enumValues_NotifyEVChargingNeedsStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"Processing",
}
var enumValues_OCPPInterfaceEnumType = []interface{}{
	"Wired0",
	"Wired1",
	"Wired2",
	"Wired3",
	"Wireless0",
	"Wireless1",
	"Wireless2",
	"Wireless3",
}
var enumValues_OCPPInterfaceEnumType_1 = []interface{}{
	"Wired0",
	"Wired1",
	"Wired2",
	"Wired3",
	"Wireless0",
	"Wireless1",
	"Wireless2",
	"Wireless3",
}
var enumValues_OCPPTransportEnumType = []interface{}{
	"JSON",
	"SOAP",
}
var enumValues_OCPPTransportEnumType_1 = []interface{}{
	"JSON",
	"SOAP",
}
var enumValues_OCPPVersionEnumType = []interface{}{
	"OCPP12",
	"OCPP15",
	"OCPP16",
	"OCPP20",
}
var enumValues_OCPPVersionEnumType_1 = []interface{}{
	"OCPP12",
	"OCPP15",
	"OCPP16",
	"OCPP20",
}
var enumValues_OperationalStatusEnumType = []interface{}{
	"Inoperative",
	"Operative",
}
var enumValues_OperationalStatusEnumType_1 = []interface{}{
	"Inoperative",
	"Operative",
}
var enumValues_PhaseEnumType = []interface{}{
	"L1",
	"L2",
	"L3",
	"N",
	"L1-N",
	"L2-N",
	"L3-N",
	"L1-L2",
	"L2-L3",
	"L3-L1",
}
var enumValues_PhaseEnumType_1 = []interface{}{
	"L1",
	"L2",
	"L3",
	"N",
	"L1-N",
	"L2-N",
	"L3-N",
	"L1-L2",
	"L2-L3",
	"L3-L1",
}
var enumValues_PhaseEnumType_2 = []interface{}{
	"L1",
	"L2",
	"L3",
	"N",
	"L1-N",
	"L2-N",
	"L3-N",
	"L1-L2",
	"L2-L3",
	"L3-L1",
}
var enumValues_PhaseEnumType_3 = []interface{}{
	"L1",
	"L2",
	"L3",
	"N",
	"L1-N",
	"L2-N",
	"L3-N",
	"L1-L2",
	"L2-L3",
	"L3-L1",
}
var enumValues_PublishFirmwareStatusEnumType = []interface{}{
	"Idle",
	"DownloadScheduled",
	"Downloading",
	"Downloaded",
	"Published",
	"DownloadFailed",
	"DownloadPaused",
	"InvalidChecksum",
	"ChecksumVerified",
	"PublishFailed",
}
var enumValues_PublishFirmwareStatusEnumType_1 = []interface{}{
	"Idle",
	"DownloadScheduled",
	"Downloading",
	"Downloaded",
	"Published",
	"DownloadFailed",
	"DownloadPaused",
	"InvalidChecksum",
	"ChecksumVerified",
	"PublishFailed",
}
var enumValues_ReadingContextEnumType = []interface{}{
	"Interruption.Begin",
	"Interruption.End",
	"Other",
	"Sample.Clock",
	"Sample.Periodic",
	"Transaction.Begin",
	"Transaction.End",
	"Trigger",
}
var enumValues_ReadingContextEnumType_1 = []interface{}{
	"Interruption.Begin",
	"Interruption.End",
	"Other",
	"Sample.Clock",
	"Sample.Periodic",
	"Transaction.Begin",
	"Transaction.End",
	"Trigger",
}
var enumValues_ReadingContextEnumType_2 = []interface{}{
	"Interruption.Begin",
	"Interruption.End",
	"Other",
	"Sample.Clock",
	"Sample.Periodic",
	"Transaction.Begin",
	"Transaction.End",
	"Trigger",
}
var enumValues_ReadingContextEnumType_3 = []interface{}{
	"Interruption.Begin",
	"Interruption.End",
	"Other",
	"Sample.Clock",
	"Sample.Periodic",
	"Transaction.Begin",
	"Transaction.End",
	"Trigger",
}
var enumValues_ReasonEnumType = []interface{}{
	"DeAuthorized",
	"EmergencyStop",
	"EnergyLimitReached",
	"EVDisconnected",
	"GroundFault",
	"ImmediateReset",
	"Local",
	"LocalOutOfCredit",
	"MasterPass",
	"Other",
	"OvercurrentFault",
	"PowerLoss",
	"PowerQuality",
	"Reboot",
	"Remote",
	"SOCLimitReached",
	"StoppedByEV",
	"TimeLimitReached",
	"Timeout",
}
var enumValues_ReasonEnumType_1 = []interface{}{
	"DeAuthorized",
	"EmergencyStop",
	"EnergyLimitReached",
	"EVDisconnected",
	"GroundFault",
	"ImmediateReset",
	"Local",
	"LocalOutOfCredit",
	"MasterPass",
	"Other",
	"OvercurrentFault",
	"PowerLoss",
	"PowerQuality",
	"Reboot",
	"Remote",
	"SOCLimitReached",
	"StoppedByEV",
	"TimeLimitReached",
	"Timeout",
}
var enumValues_RecurrencyKindEnumType = []interface{}{
	"Daily",
	"Weekly",
}
var enumValues_RecurrencyKindEnumType_1 = []interface{}{
	"Daily",
	"Weekly",
}
var enumValues_RecurrencyKindEnumType_2 = []interface{}{
	"Daily",
	"Weekly",
}
var enumValues_RecurrencyKindEnumType_3 = []interface{}{
	"Daily",
	"Weekly",
}
var enumValues_RecurrencyKindEnumType_4 = []interface{}{
	"Daily",
	"Weekly",
}
var enumValues_RecurrencyKindEnumType_5 = []interface{}{
	"Daily",
	"Weekly",
}
var enumValues_RegistrationStatusEnumType = []interface{}{
	"Accepted",
	"Pending",
	"Rejected",
}
var enumValues_RegistrationStatusEnumType_1 = []interface{}{
	"Accepted",
	"Pending",
	"Rejected",
}
var enumValues_ReportBaseEnumType = []interface{}{
	"ConfigurationInventory",
	"FullInventory",
	"SummaryInventory",
}
var enumValues_ReportBaseEnumType_1 = []interface{}{
	"ConfigurationInventory",
	"FullInventory",
	"SummaryInventory",
}
var enumValues_RequestStartStopStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_RequestStartStopStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_RequestStartStopStatusEnumType_2 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_RequestStartStopStatusEnumType_3 = []interface{}{
	"Accepted",
	"Rejected",
}
var enumValues_ReservationUpdateStatusEnumType = []interface{}{
	"Expired",
	"Removed",
}
var enumValues_ReservationUpdateStatusEnumType_1 = []interface{}{
	"Expired",
	"Removed",
}
var enumValues_ReserveNowStatusEnumType = []interface{}{
	"Accepted",
	"Faulted",
	"Occupied",
	"Rejected",
	"Unavailable",
}
var enumValues_ReserveNowStatusEnumType_1 = []interface{}{
	"Accepted",
	"Faulted",
	"Occupied",
	"Rejected",
	"Unavailable",
}
var enumValues_ResetEnumType = []interface{}{
	"Immediate",
	"OnIdle",
}
var enumValues_ResetEnumType_1 = []interface{}{
	"Immediate",
	"OnIdle",
}
var enumValues_ResetStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"Scheduled",
}
var enumValues_ResetStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"Scheduled",
}
var enumValues_SendLocalListStatusEnumType = []interface{}{
	"Accepted",
	"Failed",
	"VersionMismatch",
}
var enumValues_SendLocalListStatusEnumType_1 = []interface{}{
	"Accepted",
	"Failed",
	"VersionMismatch",
}
var enumValues_SetMonitoringStatusEnumType = []interface{}{
	"Accepted",
	"UnknownComponent",
	"UnknownVariable",
	"UnsupportedMonitorType",
	"Rejected",
	"Duplicate",
}
var enumValues_SetMonitoringStatusEnumType_1 = []interface{}{
	"Accepted",
	"UnknownComponent",
	"UnknownVariable",
	"UnsupportedMonitorType",
	"Rejected",
	"Duplicate",
}
var enumValues_SetNetworkProfileStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"Failed",
}
var enumValues_SetNetworkProfileStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"Failed",
}
var enumValues_SetVariableStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownComponent",
	"UnknownVariable",
	"NotSupportedAttributeType",
	"RebootRequired",
}
var enumValues_SetVariableStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"UnknownComponent",
	"UnknownVariable",
	"NotSupportedAttributeType",
	"RebootRequired",
}
var enumValues_TransactionEventEnumType = []interface{}{
	"Ended",
	"Started",
	"Updated",
}
var enumValues_TransactionEventEnumType_1 = []interface{}{
	"Ended",
	"Started",
	"Updated",
}
var enumValues_TriggerMessageStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"NotImplemented",
}
var enumValues_TriggerMessageStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"NotImplemented",
}
var enumValues_TriggerReasonEnumType = []interface{}{
	"Authorized",
	"CablePluggedIn",
	"ChargingRateChanged",
	"ChargingStateChanged",
	"Deauthorized",
	"EnergyLimitReached",
	"EVCommunicationLost",
	"EVConnectTimeout",
	"MeterValueClock",
	"MeterValuePeriodic",
	"TimeLimitReached",
	"Trigger",
	"UnlockCommand",
	"StopAuthorized",
	"EVDeparted",
	"EVDetected",
	"RemoteStop",
	"RemoteStart",
	"AbnormalCondition",
	"SignedDataReceived",
	"ResetCommand",
}
var enumValues_TriggerReasonEnumType_1 = []interface{}{
	"Authorized",
	"CablePluggedIn",
	"ChargingRateChanged",
	"ChargingStateChanged",
	"Deauthorized",
	"EnergyLimitReached",
	"EVCommunicationLost",
	"EVConnectTimeout",
	"MeterValueClock",
	"MeterValuePeriodic",
	"TimeLimitReached",
	"Trigger",
	"UnlockCommand",
	"StopAuthorized",
	"EVDeparted",
	"EVDetected",
	"RemoteStop",
	"RemoteStart",
	"AbnormalCondition",
	"SignedDataReceived",
	"ResetCommand",
}
var enumValues_UnlockStatusEnumType = []interface{}{
	"Unlocked",
	"UnlockFailed",
	"OngoingAuthorizedTransaction",
	"UnknownConnector",
}
var enumValues_UnlockStatusEnumType_1 = []interface{}{
	"Unlocked",
	"UnlockFailed",
	"OngoingAuthorizedTransaction",
	"UnknownConnector",
}
var enumValues_UnpublishFirmwareStatusEnumType = []interface{}{
	"DownloadOngoing",
	"NoFirmware",
	"Unpublished",
}
var enumValues_UnpublishFirmwareStatusEnumType_1 = []interface{}{
	"DownloadOngoing",
	"NoFirmware",
	"Unpublished",
}
var enumValues_UpdateEnumType = []interface{}{
	"Differential",
	"Full",
}
var enumValues_UpdateEnumType_1 = []interface{}{
	"Differential",
	"Full",
}
var enumValues_UpdateFirmwareStatusEnumType = []interface{}{
	"Accepted",
	"Rejected",
	"AcceptedCanceled",
	"InvalidCertificate",
	"RevokedCertificate",
}
var enumValues_UpdateFirmwareStatusEnumType_1 = []interface{}{
	"Accepted",
	"Rejected",
	"AcceptedCanceled",
	"InvalidCertificate",
	"RevokedCertificate",
}
var enumValues_UploadLogStatusEnumType = []interface{}{
	"BadMessage",
	"Idle",
	"NotSupportedOperation",
	"PermissionDenied",
	"Uploaded",
	"UploadFailure",
	"Uploading",
	"AcceptedCanceled",
}
var enumValues_UploadLogStatusEnumType_1 = []interface{}{
	"BadMessage",
	"Idle",
	"NotSupportedOperation",
	"PermissionDenied",
	"Uploaded",
	"UploadFailure",
	"Uploading",
	"AcceptedCanceled",
}
var enumValues_VPNEnumType = []interface{}{
	"IKEv2",
	"IPSec",
	"L2TP",
	"PPTP",
}
var enumValues_VPNEnumType_1 = []interface{}{
	"IKEv2",
	"IPSec",
	"L2TP",
	"PPTP",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *UpdateFirmwareResponseJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["status"]; !ok || v == nil {
		return fmt.Errorf("field status: required")
	}
	type Plain UpdateFirmwareResponseJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = UpdateFirmwareResponseJson(plain)
	return nil
}
