// Code generated by protoc-gen-go.
// source: artifact.proto
// DO NOT EDIT!

package apiclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ArtifactSource_SourceType int32

const (
	ArtifactSource_COLLECTOR_TYPE_UNKNOWN ArtifactSource_SourceType = 0
	ArtifactSource_FILE                   ArtifactSource_SourceType = 1
	ArtifactSource_REGISTRY_KEY           ArtifactSource_SourceType = 2
	ArtifactSource_REGISTRY_VALUE         ArtifactSource_SourceType = 3
	ArtifactSource_WMI                    ArtifactSource_SourceType = 4
	// ARTIFACT has been deprecated in favor of ARTIFACT_GROUP.
	ArtifactSource_ARTIFACT       ArtifactSource_SourceType = 5
	ArtifactSource_PATH           ArtifactSource_SourceType = 6
	ArtifactSource_DIRECTORY      ArtifactSource_SourceType = 7
	ArtifactSource_ARTIFACT_GROUP ArtifactSource_SourceType = 8
	// TODO(user): these types will likely be separated out from artifacts in
	// the future
	ArtifactSource_GRR_CLIENT_ACTION ArtifactSource_SourceType = 40
	ArtifactSource_LIST_FILES        ArtifactSource_SourceType = 41
	ArtifactSource_ARTIFACT_FILES    ArtifactSource_SourceType = 42
	ArtifactSource_GREP              ArtifactSource_SourceType = 43
	ArtifactSource_COMMAND           ArtifactSource_SourceType = 45
	ArtifactSource_REKALL_PLUGIN     ArtifactSource_SourceType = 46
)

var ArtifactSource_SourceType_name = map[int32]string{
	0:  "COLLECTOR_TYPE_UNKNOWN",
	1:  "FILE",
	2:  "REGISTRY_KEY",
	3:  "REGISTRY_VALUE",
	4:  "WMI",
	5:  "ARTIFACT",
	6:  "PATH",
	7:  "DIRECTORY",
	8:  "ARTIFACT_GROUP",
	40: "GRR_CLIENT_ACTION",
	41: "LIST_FILES",
	42: "ARTIFACT_FILES",
	43: "GREP",
	45: "COMMAND",
	46: "REKALL_PLUGIN",
}
var ArtifactSource_SourceType_value = map[string]int32{
	"COLLECTOR_TYPE_UNKNOWN": 0,
	"FILE":              1,
	"REGISTRY_KEY":      2,
	"REGISTRY_VALUE":    3,
	"WMI":               4,
	"ARTIFACT":          5,
	"PATH":              6,
	"DIRECTORY":         7,
	"ARTIFACT_GROUP":    8,
	"GRR_CLIENT_ACTION": 40,
	"LIST_FILES":        41,
	"ARTIFACT_FILES":    42,
	"GREP":              43,
	"COMMAND":           45,
	"REKALL_PLUGIN":     46,
}

func (x ArtifactSource_SourceType) Enum() *ArtifactSource_SourceType {
	p := new(ArtifactSource_SourceType)
	*p = x
	return p
}
func (x ArtifactSource_SourceType) String() string {
	return proto.EnumName(ArtifactSource_SourceType_name, int32(x))
}
func (x *ArtifactSource_SourceType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ArtifactSource_SourceType_value, data, "ArtifactSource_SourceType")
	if err != nil {
		return err
	}
	*x = ArtifactSource_SourceType(value)
	return nil
}
func (ArtifactSource_SourceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0, 0} }

// Proto representation of an ArtifactSource.
type ArtifactSource struct {
	Type             *ArtifactSource_SourceType `protobuf:"varint,1,opt,name=type,enum=ArtifactSource_SourceType" json:"type,omitempty"`
	Attributes       *Dict                      `protobuf:"bytes,2,opt,name=attributes" json:"attributes,omitempty"`
	Conditions       []string                   `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
	ReturnedTypes    []string                   `protobuf:"bytes,4,rep,name=returned_types" json:"returned_types,omitempty"`
	SupportedOs      []string                   `protobuf:"bytes,5,rep,name=supported_os" json:"supported_os,omitempty"`
	XXX_unrecognized []byte                     `json:"-"`
}

func (m *ArtifactSource) Reset()                    { *m = ArtifactSource{} }
func (m *ArtifactSource) String() string            { return proto.CompactTextString(m) }
func (*ArtifactSource) ProtoMessage()               {}
func (*ArtifactSource) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ArtifactSource) GetType() ArtifactSource_SourceType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ArtifactSource_COLLECTOR_TYPE_UNKNOWN
}

func (m *ArtifactSource) GetAttributes() *Dict {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *ArtifactSource) GetConditions() []string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *ArtifactSource) GetReturnedTypes() []string {
	if m != nil {
		return m.ReturnedTypes
	}
	return nil
}

func (m *ArtifactSource) GetSupportedOs() []string {
	if m != nil {
		return m.SupportedOs
	}
	return nil
}

// Proto representation of an artifact.
type Artifact struct {
	Name             *string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Conditions       []string          `protobuf:"bytes,2,rep,name=conditions" json:"conditions,omitempty"`
	Doc              *string           `protobuf:"bytes,3,opt,name=doc" json:"doc,omitempty"`
	Labels           []string          `protobuf:"bytes,4,rep,name=labels" json:"labels,omitempty"`
	SupportedOs      []string          `protobuf:"bytes,5,rep,name=supported_os" json:"supported_os,omitempty"`
	Urls             []string          `protobuf:"bytes,6,rep,name=urls" json:"urls,omitempty"`
	Provides         []string          `protobuf:"bytes,8,rep,name=provides" json:"provides,omitempty"`
	Sources          []*ArtifactSource `protobuf:"bytes,9,rep,name=sources" json:"sources,omitempty"`
	ErrorMessage     *string           `protobuf:"bytes,10,opt,name=error_message" json:"error_message,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *Artifact) Reset()                    { *m = Artifact{} }
func (m *Artifact) String() string            { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()               {}
func (*Artifact) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Artifact) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Artifact) GetConditions() []string {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *Artifact) GetDoc() string {
	if m != nil && m.Doc != nil {
		return *m.Doc
	}
	return ""
}

func (m *Artifact) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Artifact) GetSupportedOs() []string {
	if m != nil {
		return m.SupportedOs
	}
	return nil
}

func (m *Artifact) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func (m *Artifact) GetProvides() []string {
	if m != nil {
		return m.Provides
	}
	return nil
}

func (m *Artifact) GetSources() []*ArtifactSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *Artifact) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

type ArtifactProcessorDescriptor struct {
	Name             *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description      *string  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	OutputTypes      []string `protobuf:"bytes,3,rep,name=output_types" json:"output_types,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ArtifactProcessorDescriptor) Reset()                    { *m = ArtifactProcessorDescriptor{} }
func (m *ArtifactProcessorDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ArtifactProcessorDescriptor) ProtoMessage()               {}
func (*ArtifactProcessorDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *ArtifactProcessorDescriptor) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ArtifactProcessorDescriptor) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *ArtifactProcessorDescriptor) GetOutputTypes() []string {
	if m != nil {
		return m.OutputTypes
	}
	return nil
}

type ArtifactDescriptor struct {
	Artifact         *Artifact                      `protobuf:"bytes,1,opt,name=artifact" json:"artifact,omitempty"`
	Dependencies     []string                       `protobuf:"bytes,2,rep,name=dependencies" json:"dependencies,omitempty"`
	PathDependencies []string                       `protobuf:"bytes,3,rep,name=path_dependencies" json:"path_dependencies,omitempty"`
	ArtifactSource   *string                        `protobuf:"bytes,4,opt,name=artifact_source" json:"artifact_source,omitempty"`
	Processors       []*ArtifactProcessorDescriptor `protobuf:"bytes,5,rep,name=processors" json:"processors,omitempty"`
	IsCustom         *bool                          `protobuf:"varint,6,opt,name=is_custom" json:"is_custom,omitempty"`
	ErrorMessage     *string                        `protobuf:"bytes,7,opt,name=error_message" json:"error_message,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *ArtifactDescriptor) Reset()                    { *m = ArtifactDescriptor{} }
func (m *ArtifactDescriptor) String() string            { return proto.CompactTextString(m) }
func (*ArtifactDescriptor) ProtoMessage()               {}
func (*ArtifactDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *ArtifactDescriptor) GetArtifact() *Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

func (m *ArtifactDescriptor) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *ArtifactDescriptor) GetPathDependencies() []string {
	if m != nil {
		return m.PathDependencies
	}
	return nil
}

func (m *ArtifactDescriptor) GetArtifactSource() string {
	if m != nil && m.ArtifactSource != nil {
		return *m.ArtifactSource
	}
	return ""
}

func (m *ArtifactDescriptor) GetProcessors() []*ArtifactProcessorDescriptor {
	if m != nil {
		return m.Processors
	}
	return nil
}

func (m *ArtifactDescriptor) GetIsCustom() bool {
	if m != nil && m.IsCustom != nil {
		return *m.IsCustom
	}
	return false
}

func (m *ArtifactDescriptor) GetErrorMessage() string {
	if m != nil && m.ErrorMessage != nil {
		return *m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*ArtifactSource)(nil), "ArtifactSource")
	proto.RegisterType((*Artifact)(nil), "Artifact")
	proto.RegisterType((*ArtifactProcessorDescriptor)(nil), "ArtifactProcessorDescriptor")
	proto.RegisterType((*ArtifactDescriptor)(nil), "ArtifactDescriptor")
	proto.RegisterEnum("ArtifactSource_SourceType", ArtifactSource_SourceType_name, ArtifactSource_SourceType_value)
}

func init() { proto.RegisterFile("artifact.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xdb, 0x46,
	0x13, 0xfd, 0x64, 0xc9, 0xb6, 0xb4, 0x76, 0x14, 0x66, 0x3e, 0xa4, 0x65, 0x93, 0x34, 0x5d, 0xb8,
	0x6d, 0xa2, 0x34, 0x09, 0x83, 0x18, 0xf9, 0x6d, 0x82, 0x14, 0xfa, 0x8b, 0xaa, 0x58, 0x91, 0x04,
	0x99, 0x4e, 0x6a, 0xa0, 0x00, 0x41, 0x91, 0x2b, 0x69, 0x53, 0x8a, 0xcb, 0xec, 0x2e, 0x63, 0xe8,
	0xb2, 0x8f, 0xd0, 0xfb, 0xbe, 0x43, 0x9f, 0xa1, 0x40, 0x9f, 0xa4, 0x7d, 0x82, 0xde, 0xf7, 0xa2,
	0xd8, 0x5d, 0x51, 0xb4, 0x8c, 0xc0, 0x57, 0xbd, 0xb2, 0xc5, 0x9d, 0x39, 0x67, 0xce, 0xce, 0x99,
	0x59, 0x54, 0xf5, 0xb9, 0xa4, 0x13, 0x3f, 0x90, 0x4e, 0xc2, 0x99, 0x64, 0x57, 0xd0, 0x3b, 0x36,
	0x16, 0xcb, 0xff, 0xab, 0x82, 0xcc, 0xfd, 0x58, 0xd2, 0xc0, 0xfc, 0xde, 0xfb, 0x75, 0x0b, 0x55,
	0xeb, 0xcb, 0xf0, 0x43, 0x96, 0xf2, 0x80, 0x40, 0x1b, 0x95, 0xe4, 0x22, 0x21, 0x76, 0x01, 0x17,
	0x6a, 0xd5, 0xfd, 0x2b, 0xce, 0xfa, 0xb1, 0x63, 0xfe, 0xb8, 0x8b, 0x84, 0x34, 0xae, 0xfe, 0xf9,
	0xcf, 0x5f, 0x7f, 0x14, 0x2e, 0xc3, 0xff, 0xdd, 0x19, 0xc1, 0x2a, 0x07, 0xb3, 0x09, 0x16, 0x26,
	0x0c, 0x7e, 0x44, 0xc8, 0x97, 0x92, 0xd3, 0x71, 0x2a, 0x89, 0xb0, 0x37, 0x70, 0xa1, 0xb6, 0xb3,
	0xbf, 0xe9, 0xb4, 0x68, 0x20, 0x1b, 0x75, 0x9d, 0xf7, 0x0c, 0x9e, 0xaa, 0xbc, 0x3c, 0x08, 0xcb,
	0x99, 0x2f, 0x71, 0x48, 0x44, 0xc0, 0xe9, 0x98, 0x60, 0x39, 0x23, 0x19, 0x18, 0x26, 0xce, 0xd4,
	0xc1, 0x13, 0x1a, 0x11, 0x9c, 0xf8, 0x72, 0x26, 0x1c, 0x78, 0x87, 0x50, 0xc0, 0xe2, 0x90, 0x4a,
	0xca, 0x62, 0x61, 0x17, 0x71, 0xb1, 0x56, 0x69, 0xbc, 0xd1, 0xb0, 0x43, 0xe8, 0x0f, 0xc6, 0xef,
	0x48, 0x20, 0x55, 0xbc, 0x24, 0x1c, 0xe7, 0x71, 0x19, 0x41, 0x40, 0x43, 0x82, 0xe9, 0x04, 0xcb,
	0x19, 0x15, 0x4b, 0x0a, 0xec, 0x27, 0x49, 0x44, 0x55, 0x0d, 0x0c, 0xfb, 0x78, 0x4a, 0x3f, 0x90,
	0x18, 0x8b, 0x85, 0x90, 0x64, 0xee, 0xc0, 0x2f, 0x05, 0x54, 0xe5, 0x44, 0xa6, 0x3c, 0x26, 0xa1,
	0xa7, 0x64, 0x0a, 0xbb, 0xa4, 0x09, 0x17, 0x9a, 0x50, 0xc0, 0xfb, 0x3a, 0x8e, 0xa8, 0x90, 0x4a,
	0xbd, 0x3e, 0x36, 0x3c, 0x73, 0x7f, 0x81, 0xc7, 0x04, 0x67, 0x99, 0x78, 0xbc, 0x38, 0x4d, 0xe8,
	0xe0, 0x7a, 0xbc, 0x90, 0x33, 0x1a, 0x4f, 0xf3, 0x08, 0x9d, 0x46, 0x05, 0x8e, 0x99, 0xc4, 0x34,
	0x36, 0xd1, 0x1a, 0xf8, 0x84, 0x46, 0x91, 0xc2, 0x32, 0x8a, 0x48, 0xe8, 0x80, 0x8f, 0x76, 0x45,
	0x9a, 0x24, 0x8c, 0x4b, 0x12, 0x7a, 0x4c, 0xd8, 0x9b, 0xba, 0xa0, 0x03, 0x5d, 0x50, 0x1b, 0x9a,
	0x79, 0x41, 0x2c, 0x21, 0xdc, 0x97, 0x8a, 0xc6, 0xc8, 0xd1, 0x0a, 0x4f, 0x66, 0x34, 0x98, 0xad,
	0xc9, 0x17, 0x33, 0x96, 0x46, 0xa1, 0xe2, 0x30, 0x17, 0x11, 0x3a, 0x7b, 0x3f, 0x6f, 0x20, 0x94,
	0x37, 0x1b, 0xae, 0xa0, 0x4f, 0x9a, 0x83, 0x5e, 0xaf, 0xdd, 0x74, 0x07, 0x23, 0xcf, 0x3d, 0x1e,
	0xb6, 0xbd, 0xa3, 0xfe, 0x41, 0x7f, 0xf0, 0xb6, 0x6f, 0xfd, 0x0f, 0xca, 0xa8, 0xf4, 0xb2, 0xdb,
	0x6b, 0x5b, 0x05, 0xb0, 0xd0, 0xee, 0xa8, 0xdd, 0xe9, 0x1e, 0xba, 0xa3, 0x63, 0xef, 0xa0, 0x7d,
	0x6c, 0x6d, 0x00, 0xa0, 0xea, 0xea, 0xcb, 0x9b, 0x7a, 0xef, 0xa8, 0x6d, 0x15, 0x61, 0x1b, 0x15,
	0xdf, 0xbe, 0xee, 0x5a, 0x25, 0xd8, 0x45, 0xe5, 0xfa, 0xc8, 0xed, 0xbe, 0xac, 0x37, 0x5d, 0x6b,
	0x53, 0xc1, 0x0c, 0xeb, 0xee, 0xf7, 0xd6, 0x16, 0x5c, 0x40, 0x95, 0x56, 0x77, 0xa4, 0xb9, 0x8e,
	0xad, 0x6d, 0x85, 0x91, 0x85, 0x79, 0x9d, 0xd1, 0xe0, 0x68, 0x68, 0x95, 0xe1, 0x32, 0xba, 0xd4,
	0x19, 0x8d, 0xbc, 0x66, 0xaf, 0xdb, 0xee, 0xbb, 0x5e, 0xbd, 0xe9, 0x76, 0x07, 0x7d, 0xab, 0x06,
	0x55, 0x84, 0x7a, 0xdd, 0x43, 0xd7, 0x53, 0xf5, 0x1c, 0x5a, 0xb7, 0xd6, 0x52, 0xcd, 0xb7, 0x6f,
	0x14, 0x4f, 0x67, 0xd4, 0x1e, 0x5a, 0xb7, 0x61, 0x07, 0x6d, 0x37, 0x07, 0xaf, 0x5f, 0xd7, 0xfb,
	0x2d, 0xeb, 0x2e, 0x5c, 0x42, 0x17, 0x46, 0xed, 0x83, 0x7a, 0xaf, 0xe7, 0x0d, 0x7b, 0x47, 0x9d,
	0x6e, 0xdf, 0x72, 0xf6, 0x7e, 0xdb, 0x44, 0xe5, 0xcc, 0xff, 0xd0, 0x45, 0xa5, 0xd8, 0x9f, 0x9b,
	0xc1, 0xa8, 0x34, 0x9e, 0xe9, 0xbb, 0x7e, 0x88, 0x76, 0xb3, 0xf3, 0xbe, 0x3f, 0x27, 0xf0, 0x75,
	0x27, 0x62, 0x63, 0x3f, 0x8a, 0x16, 0x38, 0x8d, 0xe9, 0xfb, 0x94, 0x60, 0x95, 0xa1, 0x8d, 0xa1,
	0xac, 0x9e, 0xcd, 0x25, 0xfc, 0xb0, 0x66, 0xdf, 0x0d, 0xdd, 0xbc, 0x96, 0x06, 0x7c, 0x01, 0xcf,
	0xf3, 0xe6, 0x9d, 0x67, 0xdd, 0x1c, 0x2e, 0x6b, 0x1e, 0x4f, 0x63, 0x07, 0xee, 0xa3, 0x62, 0xc8,
	0x02, 0xbb, 0xa8, 0x6b, 0xfc, 0x4a, 0x43, 0x5e, 0x87, 0x6b, 0x2d, 0x16, 0x60, 0x21, 0xb9, 0xb2,
	0xc1, 0x84, 0xf1, 0x33, 0xc5, 0xd4, 0xd1, 0x56, 0xe4, 0x8f, 0x49, 0x94, 0xd9, 0xfa, 0xbe, 0xce,
	0xba, 0x0d, 0xb7, 0xf2, 0x42, 0xcc, 0xf9, 0x3a, 0xeb, 0x98, 0x44, 0x2c, 0x9e, 0x2a, 0x3f, 0x39,
	0x30, 0xfc, 0xa8, 0x1d, 0xbf, 0xd5, 0x40, 0x0f, 0x60, 0xff, 0x5c, 0x3b, 0xae, 0x29, 0x31, 0x20,
	0xc2, 0x51, 0x5b, 0x28, 0xe5, 0x91, 0xb0, 0xb7, 0x34, 0xd2, 0x63, 0x8d, 0x74, 0x1f, 0xee, 0xe5,
	0x48, 0xea, 0xd4, 0xdc, 0xca, 0x8c, 0x44, 0x09, 0x0e, 0x59, 0x90, 0xce, 0x49, 0x2c, 0xcf, 0x68,
	0x1b, 0xa2, 0x72, 0xc2, 0xd9, 0x07, 0x1a, 0x12, 0x61, 0x97, 0x35, 0xd4, 0x0b, 0x0d, 0xf5, 0x04,
	0x1e, 0xe5, 0x50, 0x3f, 0xc5, 0xec, 0x24, 0x22, 0xe1, 0x94, 0x8c, 0x7d, 0x41, 0xf0, 0x07, 0x3f,
	0x4a, 0xf5, 0x0c, 0x53, 0x91, 0x17, 0x96, 0x81, 0xa8, 0xc2, 0xb6, 0xcd, 0xcc, 0x08, 0xbb, 0x82,
	0x8b, 0xb5, 0x9d, 0xfd, 0x8b, 0x67, 0x36, 0x64, 0xe3, 0x4b, 0xcd, 0xf0, 0x39, 0x5c, 0xcd, 0x19,
	0x72, 0x81, 0x26, 0xd5, 0x81, 0x37, 0xe8, 0x02, 0xe1, 0x9c, 0x71, 0x6f, 0x4e, 0x84, 0xf0, 0xa7,
	0xc4, 0x46, 0xba, 0x63, 0xdf, 0xe9, 0xdc, 0xa7, 0xf0, 0x58, 0xad, 0x46, 0x1d, 0x80, 0x97, 0x01,
	0xba, 0x71, 0x19, 0xce, 0x52, 0xf9, 0xc4, 0xa7, 0x11, 0x09, 0x55, 0xc5, 0x34, 0xf4, 0x95, 0x4f,
	0x9c, 0xbd, 0xbf, 0x0b, 0xe8, 0x6a, 0x56, 0xcf, 0x90, 0xb3, 0x80, 0x08, 0xc1, 0x78, 0x4b, 0xef,
	0xd4, 0x44, 0x32, 0x0e, 0xcf, 0xd7, 0x4c, 0xec, 0x68, 0xba, 0x1a, 0xdc, 0x58, 0x85, 0xde, 0x14,
	0xc6, 0xb2, 0xbe, 0xc0, 0x9c, 0x4c, 0xa9, 0xd0, 0xfb, 0x46, 0x6d, 0xa3, 0xce, 0x68, 0xe4, 0xc0,
	0x33, 0xb4, 0x13, 0x2e, 0xb1, 0x28, 0x8b, 0xf5, 0x56, 0xaf, 0x34, 0x6e, 0x68, 0x10, 0x0c, 0xd7,
	0x5b, 0xf9, 0x91, 0xb1, 0x3c, 0x15, 0xea, 0xe2, 0x0c, 0xb0, 0x03, 0x6f, 0xd1, 0x2e, 0x4b, 0x65,
	0x92, 0xca, 0xe5, 0x12, 0x35, 0x5b, 0x7b, 0xed, 0x31, 0xc8, 0x5e, 0xa7, 0xd3, 0x7b, 0x34, 0xf0,
	0x63, 0xb5, 0x97, 0x12, 0xce, 0xc2, 0x34, 0xc8, 0xf6, 0x28, 0x39, 0x05, 0xbc, 0xf7, 0x7b, 0x09,
	0x41, 0xa6, 0xf9, 0x94, 0xd4, 0x87, 0xa8, 0x9c, 0xdd, 0x97, 0x96, 0xbb, 0xb3, 0x5f, 0x59, 0xb5,
	0xaa, 0x61, 0x6b, 0x5a, 0x00, 0x2b, 0xfb, 0x82, 0xa9, 0x14, 0x24, 0x9a, 0x38, 0xf0, 0x0a, 0xed,
	0x86, 0x24, 0x21, 0x71, 0x48, 0xe2, 0x80, 0x92, 0x6c, 0x3a, 0x1f, 0xe8, 0x78, 0x07, 0xee, 0xa8,
	0x31, 0x17, 0xa7, 0x7b, 0x7a, 0xd6, 0x2a, 0x26, 0x59, 0x60, 0x16, 0x2b, 0xfb, 0x5d, 0x52, 0x0f,
	0x96, 0xb7, 0x06, 0x68, 0x74, 0x3f, 0xd4, 0x80, 0xf7, 0xe0, 0xee, 0x0a, 0xf0, 0xa0, 0x81, 0x99,
	0x7e, 0xb9, 0xce, 0x43, 0x7c, 0x82, 0x2e, 0x66, 0xdf, 0x3d, 0x63, 0x26, 0xbb, 0xa4, 0xbb, 0xf0,
	0x85, 0xc6, 0xfb, 0x0c, 0x3e, 0xcd, 0x04, 0xdd, 0x14, 0xf8, 0xd5, 0xe1, 0xa0, 0xbf, 0x7a, 0x90,
	0x63, 0x84, 0x56, 0x77, 0x66, 0x26, 0x74, 0x67, 0xff, 0x9a, 0x73, 0x8e, 0x57, 0x1a, 0xcf, 0x35,
	0xe4, 0x23, 0x78, 0xb0, 0x3a, 0x5c, 0xb6, 0x44, 0xbf, 0x47, 0x4b, 0xb0, 0xf5, 0x5a, 0x6f, 0x0a,
	0x6c, 0x7a, 0xec, 0x80, 0x8b, 0x2a, 0x54, 0x78, 0x41, 0x2a, 0x24, 0x9b, 0xdb, 0x5b, 0xb8, 0x50,
	0x2b, 0xe7, 0xbd, 0xee, 0x4e, 0xb0, 0xcb, 0x53, 0x72, 0xe7, 0x8c, 0xd0, 0x13, 0x5f, 0xe0, 0xb9,
	0x1f, 0xa7, 0x66, 0x81, 0x26, 0x11, 0xf3, 0xc3, 0xbc, 0xe1, 0xa9, 0x20, 0xfc, 0x23, 0x73, 0xb3,
	0xfd, 0x9f, 0xcc, 0xcd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x50, 0x1c, 0x3f, 0x35, 0x09,
	0x00, 0x00,
}
