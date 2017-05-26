// Code generated by protoc-gen-go.
// source: checks.proto
// DO NOT EDIT!

package apiclient

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// SECTION: Defining checks.
// -------------------------
// Checks define a broad issue (e.g. system files with lax permissions).
// Methods define distinct cases where the issues is true (e.g. files with o+w,
//   files with g+w for many groups in xattrs).
// Probes define distinct ways of either detecting an issue, or detecting
//   aspects of an issue that need to be considered in combination.
// Filters define how host data should be processed to determine if an issue
//   exists.
type Match int32

const (
	// Quantifies how many results indicate a problem.
	Match_NONE Match = 0
	Match_ONE  Match = 1
	Match_ANY  Match = 2
	Match_ALL  Match = 3
	Match_SOME Match = 4
)

var Match_name = map[int32]string{
	0: "NONE",
	1: "ONE",
	2: "ANY",
	3: "ALL",
	4: "SOME",
}
var Match_value = map[string]int32{
	"NONE": 0,
	"ONE":  1,
	"ANY":  2,
	"ALL":  3,
	"SOME": 4,
}

func (x Match) Enum() *Match {
	p := new(Match)
	*p = x
	return p
}
func (x Match) String() string {
	return proto.EnumName(Match_name, int32(x))
}
func (x *Match) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Match_value, data, "Match")
	if err != nil {
		return err
	}
	*x = Match(value)
	return nil
}
func (Match) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type Probe_Mode int32

const (
	Probe_SERIAL   Probe_Mode = 0
	Probe_PARALLEL Probe_Mode = 1
)

var Probe_Mode_name = map[int32]string{
	0: "SERIAL",
	1: "PARALLEL",
}
var Probe_Mode_value = map[string]int32{
	"SERIAL":   0,
	"PARALLEL": 1,
}

func (x Probe_Mode) Enum() *Probe_Mode {
	p := new(Probe_Mode)
	*p = x
	return p
}
func (x Probe_Mode) String() string {
	return proto.EnumName(Probe_Mode_name, int32(x))
}
func (x *Probe_Mode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Probe_Mode_value, data, "Probe_Mode")
	if err != nil {
		return err
	}
	*x = Probe_Mode(value)
	return nil
}
func (Probe_Mode) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2, 0} }

type Probe_ResultContext int32

const (
	Probe_UNKNOWN_RESULT_CONTEXT Probe_ResultContext = 0
	Probe_PARSER                 Probe_ResultContext = 1
	Probe_ANOMALY                Probe_ResultContext = 2
	Probe_RAW                    Probe_ResultContext = 3
)

var Probe_ResultContext_name = map[int32]string{
	0: "UNKNOWN_RESULT_CONTEXT",
	1: "PARSER",
	2: "ANOMALY",
	3: "RAW",
}
var Probe_ResultContext_value = map[string]int32{
	"UNKNOWN_RESULT_CONTEXT": 0,
	"PARSER":                 1,
	"ANOMALY":                2,
	"RAW":                    3,
}

func (x Probe_ResultContext) Enum() *Probe_ResultContext {
	p := new(Probe_ResultContext)
	*p = x
	return p
}
func (x Probe_ResultContext) String() string {
	return proto.EnumName(Probe_ResultContext_name, int32(x))
}
func (x *Probe_ResultContext) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Probe_ResultContext_value, data, "Probe_ResultContext")
	if err != nil {
		return err
	}
	*x = Probe_ResultContext(value)
	return nil
}
func (Probe_ResultContext) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{2, 1} }

type Check struct {
	CheckId          *string   `protobuf:"bytes,1,opt,name=check_id" json:"check_id,omitempty"`
	Method           []*Method `protobuf:"bytes,2,rep,name=method" json:"method,omitempty"`
	Match            []Match   `protobuf:"varint,3,rep,name=match,enum=Match" json:"match,omitempty"`
	Target           *Target   `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	Hint             *Hint     `protobuf:"bytes,5,opt,name=hint" json:"hint,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Check) Reset()                    { *m = Check{} }
func (m *Check) String() string            { return proto.CompactTextString(m) }
func (*Check) ProtoMessage()               {}
func (*Check) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Check) GetCheckId() string {
	if m != nil && m.CheckId != nil {
		return *m.CheckId
	}
	return ""
}

func (m *Check) GetMethod() []*Method {
	if m != nil {
		return m.Method
	}
	return nil
}

func (m *Check) GetMatch() []Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Check) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Check) GetHint() *Hint {
	if m != nil {
		return m.Hint
	}
	return nil
}

type Method struct {
	Probe            []*Probe `protobuf:"bytes,1,rep,name=probe" json:"probe,omitempty"`
	Match            []Match  `protobuf:"varint,2,rep,name=match,enum=Match" json:"match,omitempty"`
	Resource         []*Dict  `protobuf:"bytes,3,rep,name=resource" json:"resource,omitempty"`
	Target           *Target  `protobuf:"bytes,4,opt,name=target" json:"target,omitempty"`
	Hint             *Hint    `protobuf:"bytes,5,opt,name=hint" json:"hint,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Method) Reset()                    { *m = Method{} }
func (m *Method) String() string            { return proto.CompactTextString(m) }
func (*Method) ProtoMessage()               {}
func (*Method) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *Method) GetProbe() []*Probe {
	if m != nil {
		return m.Probe
	}
	return nil
}

func (m *Method) GetMatch() []Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Method) GetResource() []*Dict {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Method) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Method) GetHint() *Hint {
	if m != nil {
		return m.Hint
	}
	return nil
}

type Probe struct {
	Artifact         *string              `protobuf:"bytes,1,opt,name=artifact" json:"artifact,omitempty"`
	Parser           []string             `protobuf:"bytes,2,rep,name=parser" json:"parser,omitempty"`
	Mode             *Probe_Mode          `protobuf:"varint,3,opt,name=mode,enum=Probe_Mode,def=0" json:"mode,omitempty"`
	Baseline         []*Filter            `protobuf:"bytes,4,rep,name=baseline" json:"baseline,omitempty"`
	Filters          []*Filter            `protobuf:"bytes,5,rep,name=filters" json:"filters,omitempty"`
	Match            []Match              `protobuf:"varint,6,rep,name=match,enum=Match" json:"match,omitempty"`
	Target           *Target              `protobuf:"bytes,7,opt,name=target" json:"target,omitempty"`
	Hint             *Hint                `protobuf:"bytes,8,opt,name=hint" json:"hint,omitempty"`
	ResultContext    *Probe_ResultContext `protobuf:"varint,9,opt,name=result_context,enum=Probe_ResultContext,def=1" json:"result_context,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *Probe) Reset()                    { *m = Probe{} }
func (m *Probe) String() string            { return proto.CompactTextString(m) }
func (*Probe) ProtoMessage()               {}
func (*Probe) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

const Default_Probe_Mode Probe_Mode = Probe_SERIAL
const Default_Probe_ResultContext Probe_ResultContext = Probe_PARSER

func (m *Probe) GetArtifact() string {
	if m != nil && m.Artifact != nil {
		return *m.Artifact
	}
	return ""
}

func (m *Probe) GetParser() []string {
	if m != nil {
		return m.Parser
	}
	return nil
}

func (m *Probe) GetMode() Probe_Mode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return Default_Probe_Mode
}

func (m *Probe) GetBaseline() []*Filter {
	if m != nil {
		return m.Baseline
	}
	return nil
}

func (m *Probe) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Probe) GetMatch() []Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Probe) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *Probe) GetHint() *Hint {
	if m != nil {
		return m.Hint
	}
	return nil
}

func (m *Probe) GetResultContext() Probe_ResultContext {
	if m != nil && m.ResultContext != nil {
		return *m.ResultContext
	}
	return Default_Probe_ResultContext
}

type Filter struct {
	Type             *string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Expression       *string `protobuf:"bytes,2,opt,name=expression" json:"expression,omitempty"`
	Hint             *Hint   `protobuf:"bytes,5,opt,name=hint" json:"hint,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *Filter) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Filter) GetExpression() string {
	if m != nil && m.Expression != nil {
		return *m.Expression
	}
	return ""
}

func (m *Filter) GetHint() *Hint {
	if m != nil {
		return m.Hint
	}
	return nil
}

// SECTION: Reporting issues.
// --------------------------
// The result of a single check. A check without anomalies was run, but did not
// detect a problem.
type CheckResult struct {
	CheckId          *string    `protobuf:"bytes,1,opt,name=check_id" json:"check_id,omitempty"`
	Anomaly          []*Anomaly `protobuf:"bytes,2,rep,name=anomaly" json:"anomaly,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *CheckResult) Reset()                    { *m = CheckResult{} }
func (m *CheckResult) String() string            { return proto.CompactTextString(m) }
func (*CheckResult) ProtoMessage()               {}
func (*CheckResult) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *CheckResult) GetCheckId() string {
	if m != nil && m.CheckId != nil {
		return *m.CheckId
	}
	return ""
}

func (m *CheckResult) GetAnomaly() []*Anomaly {
	if m != nil {
		return m.Anomaly
	}
	return nil
}

// The results of all checks performed on a host from a flow. This provides a
// manifest of completed checks, which can be used to:
// - open new issues, if the check results include anomalies.
// - update existing issues, if the check results vary from previous state.
// - close existing issues, if the check indicates issues weren't found.
type CheckResults struct {
	Kb               *KnowledgeBase `protobuf:"bytes,1,opt,name=kb" json:"kb,omitempty"`
	Result           []*CheckResult `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *CheckResults) Reset()                    { *m = CheckResults{} }
func (m *CheckResults) String() string            { return proto.CompactTextString(m) }
func (*CheckResults) ProtoMessage()               {}
func (*CheckResults) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *CheckResults) GetKb() *KnowledgeBase {
	if m != nil {
		return m.Kb
	}
	return nil
}

func (m *CheckResults) GetResult() []*CheckResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// Information about what checks are looking for, what actions can be taken to
// address an issue, and template text to present finding data in a usable,
// condensed form when reporting problems.
type Hint struct {
	Problem          *string `protobuf:"bytes,1,opt,name=problem" json:"problem,omitempty"`
	Fix              *string `protobuf:"bytes,2,opt,name=fix" json:"fix,omitempty"`
	Format           *string `protobuf:"bytes,3,opt,name=format" json:"format,omitempty"`
	Summary          *string `protobuf:"bytes,4,opt,name=summary" json:"summary,omitempty"`
	MaxResults       *uint64 `protobuf:"varint,5,opt,name=max_results" json:"max_results,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Hint) Reset()                    { *m = Hint{} }
func (m *Hint) String() string            { return proto.CompactTextString(m) }
func (*Hint) ProtoMessage()               {}
func (*Hint) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *Hint) GetProblem() string {
	if m != nil && m.Problem != nil {
		return *m.Problem
	}
	return ""
}

func (m *Hint) GetFix() string {
	if m != nil && m.Fix != nil {
		return *m.Fix
	}
	return ""
}

func (m *Hint) GetFormat() string {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return ""
}

func (m *Hint) GetSummary() string {
	if m != nil && m.Summary != nil {
		return *m.Summary
	}
	return ""
}

func (m *Hint) GetMaxResults() uint64 {
	if m != nil && m.MaxResults != nil {
		return *m.MaxResults
	}
	return 0
}

// SECTION: Selecting targets.
// ---------------------------
type Target struct {
	Cpe              []string `protobuf:"bytes,1,rep,name=cpe" json:"cpe,omitempty"`
	Os               []string `protobuf:"bytes,2,rep,name=os" json:"os,omitempty"`
	Label            []string `protobuf:"bytes,3,rep,name=label" json:"label,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Target) Reset()                    { *m = Target{} }
func (m *Target) String() string            { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()               {}
func (*Target) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *Target) GetCpe() []string {
	if m != nil {
		return m.Cpe
	}
	return nil
}

func (m *Target) GetOs() []string {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *Target) GetLabel() []string {
	if m != nil {
		return m.Label
	}
	return nil
}

func init() {
	proto.RegisterType((*Check)(nil), "Check")
	proto.RegisterType((*Method)(nil), "Method")
	proto.RegisterType((*Probe)(nil), "Probe")
	proto.RegisterType((*Filter)(nil), "Filter")
	proto.RegisterType((*CheckResult)(nil), "CheckResult")
	proto.RegisterType((*CheckResults)(nil), "CheckResults")
	proto.RegisterType((*Hint)(nil), "Hint")
	proto.RegisterType((*Target)(nil), "Target")
	proto.RegisterEnum("Match", Match_name, Match_value)
	proto.RegisterEnum("Probe_Mode", Probe_Mode_name, Probe_Mode_value)
	proto.RegisterEnum("Probe_ResultContext", Probe_ResultContext_name, Probe_ResultContext_value)
}

func init() { proto.RegisterFile("checks.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 1548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xcd, 0x8e, 0x24, 0x47,
	0x11, 0xde, 0x9a, 0xfe, 0x9d, 0x9c, 0xf5, 0xa8, 0x95, 0xb2, 0xa0, 0x59, 0x0c, 0x84, 0xda, 0x18,
	0xcf, 0xe2, 0x71, 0xad, 0xbc, 0xf8, 0x6f, 0xd7, 0x82, 0xa5, 0x6a, 0xb6, 0x17, 0x5b, 0xee, 0x3f,
	0xf5, 0xb4, 0xb5, 0xf8, 0x34, 0xca, 0xae, 0x8a, 0xee, 0x4a, 0xb6, 0xaa, 0xb2, 0x9d, 0x99, 0xb5,
	0xd3, 0x8d, 0xc4, 0x89, 0x07, 0x80, 0x17, 0xe0, 0x09, 0x40, 0x42, 0x08, 0x89, 0x1b, 0x37, 0x6e,
	0xdc, 0x78, 0x04, 0xb8, 0x71, 0xe5, 0xca, 0x01, 0x65, 0x54, 0xd5, 0x4c, 0xef, 0x9a, 0x1f, 0x59,
	0x02, 0xb4, 0xb7, 0xce, 0xae, 0xcc, 0x2f, 0x23, 0xbe, 0xf8, 0xe2, 0xcb, 0x60, 0x37, 0xa3, 0x04,
	0xa3, 0x27, 0xc6, 0xdf, 0x68, 0x65, 0xd5, 0xad, 0x97, 0x44, 0xae, 0x32, 0x91, 0xee, 0xaa, 0x25,
	0xfb, 0x91, 0x5a, 0xd6, 0x9f, 0x5e, 0x7e, 0x92, 0xab, 0xcb, 0x14, 0xe3, 0x35, 0x5e, 0x2c, 0x85,
	0xc1, 0xea, 0xdf, 0x63, 0x83, 0x99, 0xc8, 0xad, 0x8c, 0xca, 0xf5, 0xe0, 0xf7, 0x0d, 0xd6, 0x3a,
	0x73, 0x88, 0xfc, 0x11, 0xeb, 0x12, 0xf4, 0x85, 0x8c, 0xfb, 0x1e, 0x78, 0x27, 0x87, 0xe1, 0xdb,
	0x7f, 0xfe, 0xfb, 0x5f, 0xfe, 0xe0, 0xf9, 0xfc, 0x34, 0x80, 0x58, 0x1a, 0x2b, 0xf3, 0x75, 0x21,
	0x4d, 0x22, 0xf3, 0x35, 0xc8, 0x18, 0x73, 0x2b, 0x57, 0x12, 0x35, 0xac, 0x94, 0x06, 0x91, 0x83,
	0x88, 0x9f, 0x4a, 0xa3, 0xf4, 0xce, 0xe7, 0x1f, 0xb0, 0x76, 0x86, 0x36, 0x51, 0x71, 0xff, 0x00,
	0x1a, 0x27, 0x47, 0x77, 0x3b, 0xfe, 0x98, 0x96, 0xe1, 0x80, 0xe0, 0x5e, 0xe1, 0xb7, 0x02, 0xb8,
	0x14, 0x3b, 0xb0, 0x0a, 0x44, 0x2e, 0xd2, 0x9d, 0x41, 0xb0, 0x09, 0x42, 0xa2, 0x8c, 0xf5, 0xf9,
	0xcf, 0x3d, 0xd6, 0xca, 0x84, 0x8d, 0x92, 0x7e, 0x03, 0x1a, 0x27, 0xc7, 0x77, 0xdb, 0xfe, 0xd8,
	0xad, 0xc2, 0x1d, 0x9d, 0x35, 0xfc, 0xb3, 0x45, 0x82, 0x10, 0xa9, 0x3c, 0x96, 0x56, 0xaa, 0x1c,
	0xa4, 0x81, 0x18, 0x2d, 0x46, 0x16, 0x63, 0x90, 0x2b, 0xc2, 0x31, 0x1b, 0x8c, 0x5c, 0x64, 0x31,
	0xe4, 0x45, 0xb6, 0x44, 0x0d, 0x6a, 0x45, 0xd8, 0xa0, 0xd1, 0x14, 0xa9, 0x35, 0x80, 0x5b, 0x69,
	0x2c, 0x88, 0x95, 0x45, 0x0d, 0x1b, 0xad, 0x22, 0x34, 0x46, 0xe6, 0x6b, 0x1f, 0xce, 0xd1, 0xba,
	0x98, 0x82, 0xc9, 0xa7, 0xa7, 0xb0, 0xdc, 0x41, 0x8c, 0x2b, 0x51, 0xa4, 0xd6, 0xe7, 0x8f, 0x58,
	0xdb, 0x0a, 0xbd, 0x46, 0xdb, 0x6f, 0x82, 0x47, 0xf9, 0x2c, 0x68, 0x19, 0xbe, 0x49, 0x31, 0xbd,
	0xce, 0x5f, 0x1b, 0xc9, 0x4c, 0x5a, 0xb0, 0x89, 0x34, 0x40, 0x0c, 0x3a, 0x98, 0x2a, 0x8c, 0x08,
	0xca, 0xb3, 0xc6, 0xe7, 0x0b, 0xd6, 0x4c, 0x64, 0x6e, 0xfb, 0x2d, 0x42, 0x69, 0xf9, 0x1f, 0xca,
	0xdc, 0x86, 0x0f, 0x08, 0xe3, 0x1e, 0x7f, 0x6f, 0x81, 0x5b, 0x5b, 0x88, 0x14, 0x62, 0x34, 0x91,
	0x96, 0x1b, 0x97, 0x9d, 0x71, 0xa1, 0x0b, 0x17, 0xe3, 0x32, 0xc5, 0xec, 0x14, 0x56, 0x72, 0x0b,
	0x22, 0x8f, 0xef, 0x28, 0x0d, 0x2b, 0x99, 0xc7, 0x32, 0x5f, 0x1b, 0x7f, 0xf0, 0xc7, 0x06, 0x6b,
	0x97, 0xfc, 0xf2, 0x07, 0xac, 0xe5, 0xf6, 0x62, 0xdf, 0x23, 0xde, 0xdb, 0xfe, 0xcc, 0xad, 0xc2,
	0x13, 0xba, 0x62, 0xc0, 0xe1, 0x8a, 0xf6, 0x2a, 0x6b, 0x30, 0x2a, 0x2b, 0x79, 0x87, 0x58, 0x58,
	0xb1, 0x4f, 0xfe, 0xc1, 0x8b, 0x42, 0xfe, 0x84, 0x75, 0x35, 0x1a, 0x55, 0xe8, 0x08, 0x49, 0x11,
	0x8e, 0xb8, 0x87, 0x32, 0xb2, 0xe1, 0x3b, 0x14, 0xd3, 0x1d, 0xfe, 0xe6, 0x70, 0x6b, 0xb5, 0xa0,
	0x04, 0xe0, 0x04, 0xfd, 0xb5, 0x0f, 0x8e, 0x64, 0xb0, 0xb8, 0xb5, 0xa7, 0x10, 0xa9, 0x6c, 0x23,
	0xb4, 0x34, 0x2a, 0xa7, 0x0d, 0xb7, 0x7d, 0xfe, 0x83, 0x7f, 0x55, 0x4c, 0x9f, 0xf0, 0x4e, 0xf8,
	0xb7, 0xf6, 0x8a, 0x59, 0xca, 0xf8, 0xff, 0x59, 0xcd, 0x5f, 0x1e, 0xb2, 0x16, 0x55, 0x8d, 0x9f,
	0xb1, 0xae, 0xd0, 0x56, 0xae, 0x44, 0x64, 0xab, 0x6e, 0x7c, 0x8b, 0xc0, 0xdf, 0xe0, 0xb7, 0x5d,
	0x15, 0xea, 0x6f, 0x60, 0x13, 0x61, 0x1d, 0xe8, 0x53, 0x19, 0xa3, 0xb9, 0x6a, 0xa4, 0xaa, 0xa0,
	0x0f, 0x58, 0x7b, 0x23, 0xb4, 0x41, 0x4d, 0x05, 0x3d, 0x0c, 0xef, 0x10, 0xc4, 0x6d, 0xfe, 0xfa,
	0xe3, 0x44, 0x46, 0x09, 0x94, 0xdf, 0x0c, 0x98, 0x44, 0x15, 0x69, 0x7c, 0x25, 0x8b, 0x3d, 0x80,
	0x5f, 0x78, 0xac, 0x99, 0xa9, 0xd8, 0x71, 0xef, 0x9d, 0x1c, 0xdf, 0x3d, 0x2a, 0x25, 0xe5, 0x8f,
	0x55, 0x8c, 0xf7, 0xdb, 0xe7, 0xc3, 0xf9, 0x47, 0xc1, 0x28, 0xfc, 0x31, 0x81, 0x5a, 0xae, 0x3f,
	0x54, 0x97, 0xd4, 0xd4, 0x9b, 0x4d, 0xba, 0x83, 0x95, 0x4c, 0x2d, 0x6a, 0xe3, 0x2a, 0xab, 0xa5,
	0x48, 0x41, 0x17, 0xb9, 0x29, 0xab, 0x64, 0x13, 0xad, 0x8a, 0x75, 0x02, 0x28, 0xa2, 0xa4, 0xde,
	0x86, 0x9f, 0x15, 0xce, 0x50, 0x44, 0x9a, 0xee, 0x7c, 0x98, 0x09, 0x2d, 0xd2, 0x14, 0x53, 0x42,
	0x92, 0x68, 0xf6, 0x77, 0x82, 0xa3, 0xea, 0xa9, 0x8c, 0x0b, 0xda, 0xca, 0xff, 0xe4, 0xb1, 0xae,
	0x33, 0xb7, 0x54, 0xe6, 0xd8, 0x6f, 0x56, 0x76, 0xf3, 0x88, 0x76, 0x86, 0xbf, 0xf3, 0x28, 0xb0,
	0x5f, 0x7b, 0xfc, 0x57, 0xde, 0x34, 0x47, 0x50, 0x1a, 0x32, 0xa5, 0xb1, 0xbe, 0x12, 0x0a, 0x83,
	0x54, 0x5b, 0x74, 0x02, 0x8a, 0x2c, 0xd4, 0x38, 0xfb, 0x5a, 0x12, 0x69, 0x4a, 0x9c, 0x8a, 0x8d,
	0x93, 0xf0, 0x46, 0x19, 0x50, 0x39, 0x08, 0x30, 0x3b, 0x63, 0x31, 0xbb, 0xed, 0xc3, 0x47, 0x2b,
	0x27, 0x55, 0x99, 0x63, 0x7c, 0xfa, 0xdc, 0x79, 0xb9, 0x77, 0xc1, 0x53, 0x91, 0x16, 0xc2, 0xd6,
	0x37, 0xd7, 0xbd, 0x70, 0x0a, 0x5a, 0xd8, 0x04, 0xb5, 0x2b, 0x61, 0xbe, 0x4f, 0xfa, 0x94, 0x75,
	0xaa, 0x18, 0xfb, 0xad, 0x67, 0x53, 0xba, 0x12, 0xfd, 0x3f, 0xcb, 0xa7, 0x64, 0x8b, 0x6e, 0x7c,
	0x4e, 0x06, 0xd7, 0x7d, 0xdd, 0x7e, 0xf1, 0x4c, 0xb5, 0xf3, 0x9f, 0x4d, 0x95, 0x44, 0xf7, 0xef,
	0xdb, 0xb0, 0xfb, 0xdf, 0x6c, 0x43, 0xfe, 0x5b, 0x8f, 0x1d, 0x97, 0x99, 0x5d, 0x44, 0x2a, 0x77,
	0x5e, 0xd2, 0x3f, 0xa4, 0x06, 0x78, 0xb9, 0x6a, 0x80, 0x39, 0x7d, 0x3c, 0x2b, 0xbf, 0xdd, 0x6f,
	0xcf, 0x82, 0xf9, 0xf9, 0x70, 0x1e, 0xfe, 0x84, 0xee, 0xbd, 0xe4, 0x45, 0xd9, 0x5e, 0xc6, 0x8a,
	0x35, 0xba, 0xeb, 0x6a, 0x92, 0x64, 0x5e, 0x8a, 0xa9, 0xee, 0xdd, 0x48, 0xa5, 0x29, 0x46, 0xc4,
	0xb9, 0x55, 0x4e, 0x2e, 0xf4, 0x98, 0xba, 0x2d, 0xd5, 0x1b, 0x0f, 0x0f, 0x4b, 0xbe, 0x4c, 0x5d,
	0x54, 0x8d, 0x6b, 0x69, 0x2c, 0x6a, 0x8c, 0xaf, 0x51, 0xaa, 0x2e, 0xf6, 0x07, 0xc0, 0x9a, 0xae,
	0x31, 0x39, 0x63, 0x55, 0x6b, 0xf6, 0x6e, 0xf0, 0x9b, 0xac, 0x3b, 0x0b, 0xe6, 0xc1, 0x68, 0x34,
	0x1c, 0xf5, 0xbc, 0xc1, 0x98, 0xbd, 0xf4, 0x4c, 0xe4, 0xfc, 0x16, 0xfb, 0xd2, 0x27, 0x93, 0x8f,
	0x27, 0xd3, 0xc7, 0x93, 0x8b, 0xf9, 0xf0, 0xfc, 0x93, 0xd1, 0xe2, 0xe2, 0x6c, 0x3a, 0x59, 0x0c,
	0x7f, 0xb8, 0xe8, 0xdd, 0x70, 0x30, 0x65, 0x5e, 0x3d, 0x8f, 0x1f, 0xb1, 0x4e, 0x30, 0x99, 0x8e,
	0x83, 0xd1, 0xa7, 0xbd, 0x03, 0xde, 0x61, 0x8d, 0x79, 0xf0, 0xb8, 0xd7, 0x18, 0xfc, 0xf4, 0x80,
	0xb5, 0x4b, 0x65, 0xf2, 0x7b, 0xac, 0x69, 0x77, 0x1b, 0xac, 0xac, 0xea, 0x0d, 0x22, 0xe2, 0x35,
	0xfe, 0xaa, 0x13, 0x96, 0xfb, 0xdf, 0xb1, 0x50, 0xa9, 0xdf, 0x2a, 0x48, 0x44, 0x1e, 0x57, 0x5e,
	0xa0, 0x7c, 0x9e, 0x30, 0x86, 0xdb, 0x8d, 0x76, 0x8a, 0x51, 0x79, 0xff, 0x80, 0x00, 0x16, 0x04,
	0x30, 0xe1, 0xa3, 0x00, 0x74, 0x91, 0x62, 0xd9, 0x64, 0xa5, 0x3a, 0x9d, 0xec, 0x9d, 0x7a, 0xc1,
	0xa2, 0xce, 0xae, 0x7a, 0x6c, 0x23, 0x8c, 0x71, 0x1e, 0x63, 0x65, 0x86, 0x8e, 0x11, 0x91, 0x21,
	0x35, 0x85, 0xf3, 0xa3, 0xea, 0xe6, 0xff, 0x95, 0x67, 0xff, 0xc6, 0x63, 0x47, 0x34, 0x41, 0x95,
	0xd4, 0xf2, 0xd9, 0xe7, 0xe6, 0xa8, 0xef, 0xd1, 0x15, 0xef, 0xf3, 0x77, 0x17, 0x75, 0x69, 0x41,
	0xc6, 0xa7, 0xd7, 0x43, 0x94, 0x81, 0x4b, 0xe7, 0xe2, 0xd5, 0x17, 0x03, 0x4b, 0x74, 0x33, 0x96,
	0xf3, 0x1b, 0x6d, 0x31, 0xf6, 0xf9, 0x8c, 0x75, 0xaa, 0x31, 0xaf, 0x1a, 0xa9, 0xba, 0x7e, 0x50,
	0xae, 0xaf, 0x1d, 0xe1, 0xbc, 0x6e, 0x8d, 0x3a, 0x30, 0x72, 0xca, 0x48, 0xb8, 0xa9, 0xcd, 0xcd,
	0x67, 0xd2, 0x98, 0x02, 0xcb, 0x36, 0x35, 0xfe, 0xe0, 0x67, 0x1e, 0xbb, 0xb9, 0x17, 0xb3, 0xe1,
	0x1f, 0xb0, 0x83, 0x27, 0x4b, 0x0a, 0xf7, 0xe8, 0xee, 0xb1, 0xff, 0x71, 0x3d, 0x39, 0x86, 0xc2,
	0x60, 0xf8, 0x0d, 0xba, 0xe3, 0x2b, 0xfc, 0xcb, 0x0f, 0xd1, 0x0a, 0x99, 0x1a, 0x10, 0x4b, 0x55,
	0xd8, 0xbd, 0xa1, 0x2d, 0x64, 0xed, 0x52, 0xe2, 0x55, 0x78, 0x37, 0xfd, 0x3d, 0xec, 0xf0, 0x9b,
	0x74, 0xfc, 0xeb, 0xfc, 0x95, 0xeb, 0xec, 0xeb, 0x86, 0xa0, 0xd9, 0x91, 0x40, 0x06, 0x7f, 0x3d,
	0x60, 0x4d, 0x57, 0x0f, 0xfe, 0x36, 0xeb, 0x54, 0x7c, 0x57, 0xec, 0xbd, 0x4a, 0xe7, 0xbf, 0xc6,
	0xbf, 0x1a, 0xec, 0x97, 0xc6, 0x55, 0xc6, 0x45, 0x40, 0x79, 0xb9, 0xa1, 0xb3, 0xb1, 0x92, 0xdb,
	0x4a, 0x3d, 0x57, 0x6f, 0xf9, 0xe7, 0x4e, 0x24, 0xe5, 0x13, 0xe5, 0xea, 0xb8, 0x77, 0x78, 0xcc,
	0xda, 0x2b, 0xa5, 0x33, 0x61, 0xe9, 0x99, 0x3b, 0x0c, 0xbf, 0x4b, 0xe7, 0xdf, 0xe3, 0xef, 0x04,
	0x60, 0x31, 0xdb, 0xa4, 0xce, 0xbb, 0xaf, 0x25, 0x4a, 0x00, 0xb4, 0xbd, 0xe6, 0xba, 0x9e, 0x81,
	0x2b, 0xb8, 0x19, 0xeb, 0x98, 0x22, 0xcb, 0x84, 0xde, 0xd1, 0x90, 0x71, 0x78, 0x2d, 0xb1, 0xc0,
	0x3d, 0xb5, 0xda, 0x42, 0x2e, 0x32, 0xf2, 0xef, 0x67, 0x24, 0x5c, 0x86, 0xba, 0x2c, 0xa7, 0xe1,
	0xba, 0x69, 0x4a, 0x03, 0x9f, 0xb0, 0xa3, 0x4c, 0x6c, 0x2f, 0x2a, 0xda, 0x48, 0xbf, 0xcd, 0xf0,
	0x1e, 0xa1, 0x7e, 0x87, 0xbf, 0x35, 0x16, 0x5b, 0x99, 0x15, 0xd9, 0x9e, 0x35, 0x5f, 0x69, 0xc0,
	0x2a, 0x90, 0x79, 0x94, 0x16, 0x31, 0x3a, 0xf3, 0xa9, 0xce, 0xfb, 0x83, 0xbf, 0x79, 0xac, 0x5d,
	0xfa, 0x2d, 0x1f, 0xb1, 0x46, 0xb4, 0x29, 0x47, 0xc6, 0xbd, 0x40, 0xe7, 0x68, 0xac, 0x96, 0xd1,
	0xf3, 0x43, 0xad, 0xab, 0x94, 0x81, 0x4b, 0x69, 0x13, 0x10, 0xf9, 0xae, 0x2a, 0x81, 0x41, 0x38,
	0x9b, 0x0d, 0x41, 0xc6, 0xc6, 0xe7, 0x23, 0x76, 0xa0, 0x4c, 0x35, 0x6c, 0x7c, 0x9f, 0xc0, 0xee,
	0xf3, 0xf7, 0xbf, 0x20, 0xd8, 0xf4, 0x9c, 0xb2, 0x37, 0xae, 0x2e, 0xad, 0x54, 0x2c, 0x31, 0xa5,
	0xc9, 0x6f, 0xaf, 0x8d, 0xbe, 0x20, 0x20, 0x61, 0x18, 0xff, 0xdb, 0xef, 0xb2, 0x16, 0xbd, 0x7b,
	0xbc, 0xcb, 0x9a, 0x93, 0xe9, 0x64, 0xd8, 0xbb, 0xe1, 0xac, 0xcc, 0xfd, 0xf0, 0xdc, 0x8f, 0x60,
	0x52, 0x99, 0x5b, 0x30, 0x1a, 0xf5, 0x1a, 0x6e, 0xd3, 0xf9, 0x74, 0x3c, 0xec, 0x35, 0xff, 0x11,
	0x00, 0x00, 0xff, 0xff, 0xe3, 0x6f, 0x47, 0x08, 0x72, 0x0d, 0x00, 0x00,
}
