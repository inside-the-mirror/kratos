// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gogo.proto

package gogoproto // import "github.com/inside-the-mirror/kratos/tool/protobuf/pkg/extensions/gogoproto"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

var E_GoprotoEnumPrefix = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62001,
	Name:          "gogoproto.goproto_enum_prefix",
	Tag:           "varint,62001,opt,name=goproto_enum_prefix,json=goprotoEnumPrefix",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62021,
	Name:          "gogoproto.goproto_enum_stringer",
	Tag:           "varint,62021,opt,name=goproto_enum_stringer,json=goprotoEnumStringer",
	Filename:      "gogo.proto",
}

var E_EnumStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62022,
	Name:          "gogoproto.enum_stringer",
	Tag:           "varint,62022,opt,name=enum_stringer,json=enumStringer",
	Filename:      "gogo.proto",
}

var E_EnumCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         62023,
	Name:          "gogoproto.enum_customname",
	Tag:           "bytes,62023,opt,name=enum_customname,json=enumCustomname",
	Filename:      "gogo.proto",
}

var E_Enumdecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         62024,
	Name:          "gogoproto.enumdecl",
	Tag:           "varint,62024,opt,name=enumdecl",
	Filename:      "gogo.proto",
}

var E_EnumvalueCustomname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         66001,
	Name:          "gogoproto.enumvalue_customname",
	Tag:           "bytes,66001,opt,name=enumvalue_customname,json=enumvalueCustomname",
	Filename:      "gogo.proto",
}

var E_GoprotoGettersAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63001,
	Name:          "gogoproto.goproto_getters_all",
	Tag:           "varint,63001,opt,name=goproto_getters_all,json=goprotoGettersAll",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumPrefixAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63002,
	Name:          "gogoproto.goproto_enum_prefix_all",
	Tag:           "varint,63002,opt,name=goproto_enum_prefix_all,json=goprotoEnumPrefixAll",
	Filename:      "gogo.proto",
}

var E_GoprotoStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63003,
	Name:          "gogoproto.goproto_stringer_all",
	Tag:           "varint,63003,opt,name=goproto_stringer_all,json=goprotoStringerAll",
	Filename:      "gogo.proto",
}

var E_VerboseEqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63004,
	Name:          "gogoproto.verbose_equal_all",
	Tag:           "varint,63004,opt,name=verbose_equal_all,json=verboseEqualAll",
	Filename:      "gogo.proto",
}

var E_FaceAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63005,
	Name:          "gogoproto.face_all",
	Tag:           "varint,63005,opt,name=face_all,json=faceAll",
	Filename:      "gogo.proto",
}

var E_GostringAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63006,
	Name:          "gogoproto.gostring_all",
	Tag:           "varint,63006,opt,name=gostring_all,json=gostringAll",
	Filename:      "gogo.proto",
}

var E_PopulateAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63007,
	Name:          "gogoproto.populate_all",
	Tag:           "varint,63007,opt,name=populate_all,json=populateAll",
	Filename:      "gogo.proto",
}

var E_StringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63008,
	Name:          "gogoproto.stringer_all",
	Tag:           "varint,63008,opt,name=stringer_all,json=stringerAll",
	Filename:      "gogo.proto",
}

var E_OnlyoneAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63009,
	Name:          "gogoproto.onlyone_all",
	Tag:           "varint,63009,opt,name=onlyone_all,json=onlyoneAll",
	Filename:      "gogo.proto",
}

var E_EqualAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63013,
	Name:          "gogoproto.equal_all",
	Tag:           "varint,63013,opt,name=equal_all,json=equalAll",
	Filename:      "gogo.proto",
}

var E_DescriptionAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63014,
	Name:          "gogoproto.description_all",
	Tag:           "varint,63014,opt,name=description_all,json=descriptionAll",
	Filename:      "gogo.proto",
}

var E_TestgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63015,
	Name:          "gogoproto.testgen_all",
	Tag:           "varint,63015,opt,name=testgen_all,json=testgenAll",
	Filename:      "gogo.proto",
}

var E_BenchgenAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63016,
	Name:          "gogoproto.benchgen_all",
	Tag:           "varint,63016,opt,name=benchgen_all,json=benchgenAll",
	Filename:      "gogo.proto",
}

var E_MarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63017,
	Name:          "gogoproto.marshaler_all",
	Tag:           "varint,63017,opt,name=marshaler_all,json=marshalerAll",
	Filename:      "gogo.proto",
}

var E_UnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63018,
	Name:          "gogoproto.unmarshaler_all",
	Tag:           "varint,63018,opt,name=unmarshaler_all,json=unmarshalerAll",
	Filename:      "gogo.proto",
}

var E_StableMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63019,
	Name:          "gogoproto.stable_marshaler_all",
	Tag:           "varint,63019,opt,name=stable_marshaler_all,json=stableMarshalerAll",
	Filename:      "gogo.proto",
}

var E_SizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63020,
	Name:          "gogoproto.sizer_all",
	Tag:           "varint,63020,opt,name=sizer_all,json=sizerAll",
	Filename:      "gogo.proto",
}

var E_GoprotoEnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63021,
	Name:          "gogoproto.goproto_enum_stringer_all",
	Tag:           "varint,63021,opt,name=goproto_enum_stringer_all,json=goprotoEnumStringerAll",
	Filename:      "gogo.proto",
}

var E_EnumStringerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63022,
	Name:          "gogoproto.enum_stringer_all",
	Tag:           "varint,63022,opt,name=enum_stringer_all,json=enumStringerAll",
	Filename:      "gogo.proto",
}

var E_UnsafeMarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63023,
	Name:          "gogoproto.unsafe_marshaler_all",
	Tag:           "varint,63023,opt,name=unsafe_marshaler_all,json=unsafeMarshalerAll",
	Filename:      "gogo.proto",
}

var E_UnsafeUnmarshalerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63024,
	Name:          "gogoproto.unsafe_unmarshaler_all",
	Tag:           "varint,63024,opt,name=unsafe_unmarshaler_all,json=unsafeUnmarshalerAll",
	Filename:      "gogo.proto",
}

var E_GoprotoExtensionsMapAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63025,
	Name:          "gogoproto.goproto_extensions_map_all",
	Tag:           "varint,63025,opt,name=goproto_extensions_map_all,json=goprotoExtensionsMapAll",
	Filename:      "gogo.proto",
}

var E_GoprotoUnrecognizedAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63026,
	Name:          "gogoproto.goproto_unrecognized_all",
	Tag:           "varint,63026,opt,name=goproto_unrecognized_all,json=goprotoUnrecognizedAll",
	Filename:      "gogo.proto",
}

var E_GogoprotoImport = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63027,
	Name:          "gogoproto.gogoproto_import",
	Tag:           "varint,63027,opt,name=gogoproto_import,json=gogoprotoImport",
	Filename:      "gogo.proto",
}

var E_ProtosizerAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63028,
	Name:          "gogoproto.protosizer_all",
	Tag:           "varint,63028,opt,name=protosizer_all,json=protosizerAll",
	Filename:      "gogo.proto",
}

var E_CompareAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63029,
	Name:          "gogoproto.compare_all",
	Tag:           "varint,63029,opt,name=compare_all,json=compareAll",
	Filename:      "gogo.proto",
}

var E_TypedeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63030,
	Name:          "gogoproto.typedecl_all",
	Tag:           "varint,63030,opt,name=typedecl_all,json=typedeclAll",
	Filename:      "gogo.proto",
}

var E_EnumdeclAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63031,
	Name:          "gogoproto.enumdecl_all",
	Tag:           "varint,63031,opt,name=enumdecl_all,json=enumdeclAll",
	Filename:      "gogo.proto",
}

var E_GoprotoRegistration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63032,
	Name:          "gogoproto.goproto_registration",
	Tag:           "varint,63032,opt,name=goproto_registration,json=goprotoRegistration",
	Filename:      "gogo.proto",
}

var E_MessagenameAll = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         63033,
	Name:          "gogoproto.messagename_all",
	Tag:           "varint,63033,opt,name=messagename_all,json=messagenameAll",
	Filename:      "gogo.proto",
}

var E_GoprotoGetters = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64001,
	Name:          "gogoproto.goproto_getters",
	Tag:           "varint,64001,opt,name=goproto_getters,json=goprotoGetters",
	Filename:      "gogo.proto",
}

var E_GoprotoStringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64003,
	Name:          "gogoproto.goproto_stringer",
	Tag:           "varint,64003,opt,name=goproto_stringer,json=goprotoStringer",
	Filename:      "gogo.proto",
}

var E_VerboseEqual = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64004,
	Name:          "gogoproto.verbose_equal",
	Tag:           "varint,64004,opt,name=verbose_equal,json=verboseEqual",
	Filename:      "gogo.proto",
}

var E_Face = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64005,
	Name:          "gogoproto.face",
	Tag:           "varint,64005,opt,name=face",
	Filename:      "gogo.proto",
}

var E_Gostring = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64006,
	Name:          "gogoproto.gostring",
	Tag:           "varint,64006,opt,name=gostring",
	Filename:      "gogo.proto",
}

var E_Populate = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64007,
	Name:          "gogoproto.populate",
	Tag:           "varint,64007,opt,name=populate",
	Filename:      "gogo.proto",
}

var E_Stringer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         67008,
	Name:          "gogoproto.stringer",
	Tag:           "varint,67008,opt,name=stringer",
	Filename:      "gogo.proto",
}

var E_Onlyone = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64009,
	Name:          "gogoproto.onlyone",
	Tag:           "varint,64009,opt,name=onlyone",
	Filename:      "gogo.proto",
}

var E_Equal = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64013,
	Name:          "gogoproto.equal",
	Tag:           "varint,64013,opt,name=equal",
	Filename:      "gogo.proto",
}

var E_Description = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64014,
	Name:          "gogoproto.description",
	Tag:           "varint,64014,opt,name=description",
	Filename:      "gogo.proto",
}

var E_Testgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64015,
	Name:          "gogoproto.testgen",
	Tag:           "varint,64015,opt,name=testgen",
	Filename:      "gogo.proto",
}

var E_Benchgen = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64016,
	Name:          "gogoproto.benchgen",
	Tag:           "varint,64016,opt,name=benchgen",
	Filename:      "gogo.proto",
}

var E_Marshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64017,
	Name:          "gogoproto.marshaler",
	Tag:           "varint,64017,opt,name=marshaler",
	Filename:      "gogo.proto",
}

var E_Unmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64018,
	Name:          "gogoproto.unmarshaler",
	Tag:           "varint,64018,opt,name=unmarshaler",
	Filename:      "gogo.proto",
}

var E_StableMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64019,
	Name:          "gogoproto.stable_marshaler",
	Tag:           "varint,64019,opt,name=stable_marshaler,json=stableMarshaler",
	Filename:      "gogo.proto",
}

var E_Sizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64020,
	Name:          "gogoproto.sizer",
	Tag:           "varint,64020,opt,name=sizer",
	Filename:      "gogo.proto",
}

var E_UnsafeMarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64023,
	Name:          "gogoproto.unsafe_marshaler",
	Tag:           "varint,64023,opt,name=unsafe_marshaler,json=unsafeMarshaler",
	Filename:      "gogo.proto",
}

var E_UnsafeUnmarshaler = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64024,
	Name:          "gogoproto.unsafe_unmarshaler",
	Tag:           "varint,64024,opt,name=unsafe_unmarshaler,json=unsafeUnmarshaler",
	Filename:      "gogo.proto",
}

var E_GoprotoExtensionsMap = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64025,
	Name:          "gogoproto.goproto_extensions_map",
	Tag:           "varint,64025,opt,name=goproto_extensions_map,json=goprotoExtensionsMap",
	Filename:      "gogo.proto",
}

var E_GoprotoUnrecognized = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64026,
	Name:          "gogoproto.goproto_unrecognized",
	Tag:           "varint,64026,opt,name=goproto_unrecognized,json=goprotoUnrecognized",
	Filename:      "gogo.proto",
}

var E_Protosizer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64028,
	Name:          "gogoproto.protosizer",
	Tag:           "varint,64028,opt,name=protosizer",
	Filename:      "gogo.proto",
}

var E_Compare = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64029,
	Name:          "gogoproto.compare",
	Tag:           "varint,64029,opt,name=compare",
	Filename:      "gogo.proto",
}

var E_Typedecl = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64030,
	Name:          "gogoproto.typedecl",
	Tag:           "varint,64030,opt,name=typedecl",
	Filename:      "gogo.proto",
}

var E_Messagename = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         64033,
	Name:          "gogoproto.messagename",
	Tag:           "varint,64033,opt,name=messagename",
	Filename:      "gogo.proto",
}

var E_Nullable = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65001,
	Name:          "gogoproto.nullable",
	Tag:           "varint,65001,opt,name=nullable",
	Filename:      "gogo.proto",
}

var E_Embed = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65002,
	Name:          "gogoproto.embed",
	Tag:           "varint,65002,opt,name=embed",
	Filename:      "gogo.proto",
}

var E_Customtype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65003,
	Name:          "gogoproto.customtype",
	Tag:           "bytes,65003,opt,name=customtype",
	Filename:      "gogo.proto",
}

var E_Customname = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65004,
	Name:          "gogoproto.customname",
	Tag:           "bytes,65004,opt,name=customname",
	Filename:      "gogo.proto",
}

var E_Jsontag = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65005,
	Name:          "gogoproto.jsontag",
	Tag:           "bytes,65005,opt,name=jsontag",
	Filename:      "gogo.proto",
}

var E_Moretags = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65006,
	Name:          "gogoproto.moretags",
	Tag:           "bytes,65006,opt,name=moretags",
	Filename:      "gogo.proto",
}

var E_Casttype = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65007,
	Name:          "gogoproto.casttype",
	Tag:           "bytes,65007,opt,name=casttype",
	Filename:      "gogo.proto",
}

var E_Castkey = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65008,
	Name:          "gogoproto.castkey",
	Tag:           "bytes,65008,opt,name=castkey",
	Filename:      "gogo.proto",
}

var E_Castvalue = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         65009,
	Name:          "gogoproto.castvalue",
	Tag:           "bytes,65009,opt,name=castvalue",
	Filename:      "gogo.proto",
}

var E_Stdtime = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65010,
	Name:          "gogoproto.stdtime",
	Tag:           "varint,65010,opt,name=stdtime",
	Filename:      "gogo.proto",
}

var E_Stdduration = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         65011,
	Name:          "gogoproto.stdduration",
	Tag:           "varint,65011,opt,name=stdduration",
	Filename:      "gogo.proto",
}

func init() {
	proto.RegisterExtension(E_GoprotoEnumPrefix)
	proto.RegisterExtension(E_GoprotoEnumStringer)
	proto.RegisterExtension(E_EnumStringer)
	proto.RegisterExtension(E_EnumCustomname)
	proto.RegisterExtension(E_Enumdecl)
	proto.RegisterExtension(E_EnumvalueCustomname)
	proto.RegisterExtension(E_GoprotoGettersAll)
	proto.RegisterExtension(E_GoprotoEnumPrefixAll)
	proto.RegisterExtension(E_GoprotoStringerAll)
	proto.RegisterExtension(E_VerboseEqualAll)
	proto.RegisterExtension(E_FaceAll)
	proto.RegisterExtension(E_GostringAll)
	proto.RegisterExtension(E_PopulateAll)
	proto.RegisterExtension(E_StringerAll)
	proto.RegisterExtension(E_OnlyoneAll)
	proto.RegisterExtension(E_EqualAll)
	proto.RegisterExtension(E_DescriptionAll)
	proto.RegisterExtension(E_TestgenAll)
	proto.RegisterExtension(E_BenchgenAll)
	proto.RegisterExtension(E_MarshalerAll)
	proto.RegisterExtension(E_UnmarshalerAll)
	proto.RegisterExtension(E_StableMarshalerAll)
	proto.RegisterExtension(E_SizerAll)
	proto.RegisterExtension(E_GoprotoEnumStringerAll)
	proto.RegisterExtension(E_EnumStringerAll)
	proto.RegisterExtension(E_UnsafeMarshalerAll)
	proto.RegisterExtension(E_UnsafeUnmarshalerAll)
	proto.RegisterExtension(E_GoprotoExtensionsMapAll)
	proto.RegisterExtension(E_GoprotoUnrecognizedAll)
	proto.RegisterExtension(E_GogoprotoImport)
	proto.RegisterExtension(E_ProtosizerAll)
	proto.RegisterExtension(E_CompareAll)
	proto.RegisterExtension(E_TypedeclAll)
	proto.RegisterExtension(E_EnumdeclAll)
	proto.RegisterExtension(E_GoprotoRegistration)
	proto.RegisterExtension(E_MessagenameAll)
	proto.RegisterExtension(E_GoprotoGetters)
	proto.RegisterExtension(E_GoprotoStringer)
	proto.RegisterExtension(E_VerboseEqual)
	proto.RegisterExtension(E_Face)
	proto.RegisterExtension(E_Gostring)
	proto.RegisterExtension(E_Populate)
	proto.RegisterExtension(E_Stringer)
	proto.RegisterExtension(E_Onlyone)
	proto.RegisterExtension(E_Equal)
	proto.RegisterExtension(E_Description)
	proto.RegisterExtension(E_Testgen)
	proto.RegisterExtension(E_Benchgen)
	proto.RegisterExtension(E_Marshaler)
	proto.RegisterExtension(E_Unmarshaler)
	proto.RegisterExtension(E_StableMarshaler)
	proto.RegisterExtension(E_Sizer)
	proto.RegisterExtension(E_UnsafeMarshaler)
	proto.RegisterExtension(E_UnsafeUnmarshaler)
	proto.RegisterExtension(E_GoprotoExtensionsMap)
	proto.RegisterExtension(E_GoprotoUnrecognized)
	proto.RegisterExtension(E_Protosizer)
	proto.RegisterExtension(E_Compare)
	proto.RegisterExtension(E_Typedecl)
	proto.RegisterExtension(E_Messagename)
	proto.RegisterExtension(E_Nullable)
	proto.RegisterExtension(E_Embed)
	proto.RegisterExtension(E_Customtype)
	proto.RegisterExtension(E_Customname)
	proto.RegisterExtension(E_Jsontag)
	proto.RegisterExtension(E_Moretags)
	proto.RegisterExtension(E_Casttype)
	proto.RegisterExtension(E_Castkey)
	proto.RegisterExtension(E_Castvalue)
	proto.RegisterExtension(E_Stdtime)
	proto.RegisterExtension(E_Stdduration)
}

func init() { proto.RegisterFile("gogo.proto", fileDescriptor_gogo_e935c22a8aa82c87) }

var fileDescriptor_gogo_e935c22a8aa82c87 = []byte{
	// 1260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x98, 0xc9, 0x6f, 0x1c, 0x45,
	0x17, 0xc0, 0xf5, 0xe9, 0x4b, 0x14, 0xcf, 0xf3, 0x86, 0xc7, 0xc6, 0x84, 0x08, 0x44, 0xb8, 0x71,
	0xc1, 0x73, 0x40, 0x11, 0x4a, 0x59, 0x91, 0xe5, 0x58, 0x8e, 0x15, 0x84, 0xc1, 0x98, 0xd8, 0x6c,
	0x87, 0x51, 0xcf, 0x4c, 0xb9, 0x33, 0xa4, 0xbb, 0xab, 0xe9, 0xae, 0x8e, 0xe2, 0xdc, 0x50, 0x58,
	0x84, 0x10, 0x3b, 0x12, 0x24, 0x24, 0x81, 0x1c, 0xd8, 0xd7, 0xb0, 0x73, 0xe3, 0xc2, 0x72, 0xe5,
	0x7f, 0xe0, 0x02, 0x98, 0xdd, 0x37, 0x5f, 0xd0, 0xeb, 0x7e, 0xaf, 0xa7, 0x66, 0x3c, 0x52, 0xd5,
	0xdc, 0xda, 0x76, 0xfd, 0x7e, 0xae, 0x7e, 0xaf, 0xea, 0xbd, 0x37, 0x03, 0xe0, 0x2b, 0x5f, 0xcd,
	0xc4, 0x89, 0xd2, 0xaa, 0x5a, 0xc1, 0xe7, 0xfc, 0xf1, 0xc0, 0x41, 0x5f, 0x29, 0x3f, 0x90, 0xb5,
	0xfc, 0xa7, 0x46, 0xb6, 0x51, 0x6b, 0xc9, 0xb4, 0x99, 0xb4, 0x63, 0xad, 0x92, 0x62, 0xb1, 0xb8,
	0x0b, 0x26, 0x69, 0x71, 0x5d, 0x46, 0x59, 0x58, 0x8f, 0x13, 0xb9, 0xd1, 0x3e, 0x53, 0xbd, 0x61,
	0xa6, 0x20, 0x67, 0x98, 0x9c, 0x59, 0x8c, 0xb2, 0xf0, 0xee, 0x58, 0xb7, 0x55, 0x94, 0xee, 0xbf,
	0xfa, 0xf3, 0xff, 0x0f, 0xfe, 0xef, 0x96, 0xa1, 0xd5, 0x09, 0x42, 0xf1, 0x6f, 0x2b, 0x39, 0x28,
	0x56, 0xe1, 0xda, 0x2e, 0x5f, 0xaa, 0x93, 0x76, 0xe4, 0xcb, 0xc4, 0x62, 0xfc, 0x8e, 0x8c, 0x93,
	0x86, 0xf1, 0x5e, 0x42, 0xc5, 0x02, 0x8c, 0x0e, 0xe2, 0xfa, 0x9e, 0x5c, 0x23, 0xd2, 0x94, 0x2c,
	0xc1, 0x78, 0x2e, 0x69, 0x66, 0xa9, 0x56, 0x61, 0xe4, 0x85, 0xd2, 0xa2, 0xf9, 0x21, 0xd7, 0x54,
	0x56, 0xc7, 0x10, 0x5b, 0x28, 0x29, 0x21, 0x60, 0x08, 0x7f, 0xd3, 0x92, 0xcd, 0xc0, 0x62, 0xf8,
	0x91, 0x36, 0x52, 0xae, 0x17, 0xeb, 0x30, 0x85, 0xcf, 0xa7, 0xbd, 0x20, 0x93, 0xe6, 0x4e, 0x6e,
	0xee, 0xeb, 0x59, 0xc7, 0x65, 0x2c, 0xfb, 0xe9, 0xdc, 0x9e, 0x7c, 0x3b, 0x93, 0xa5, 0xc0, 0xd8,
	0x93, 0x91, 0x45, 0x5f, 0x6a, 0x2d, 0x93, 0xb4, 0xee, 0x05, 0xfd, 0xb6, 0x77, 0xac, 0x1d, 0x94,
	0xc6, 0xf3, 0x5b, 0xdd, 0x59, 0x5c, 0x2a, 0xc8, 0xf9, 0x20, 0x10, 0x6b, 0x70, 0x5d, 0x9f, 0x53,
	0xe1, 0xe0, 0xbc, 0x40, 0xce, 0xa9, 0x5d, 0x27, 0x03, 0xb5, 0x2b, 0xc0, 0xbf, 0x2f, 0x73, 0xe9,
	0xe0, 0x7c, 0x8d, 0x9c, 0x55, 0x62, 0x39, 0xa5, 0x68, 0xbc, 0x03, 0x26, 0x4e, 0xcb, 0xa4, 0xa1,
	0x52, 0x59, 0x97, 0x8f, 0x64, 0x5e, 0xe0, 0xa0, 0xbb, 0x48, 0xba, 0x71, 0x02, 0x17, 0x91, 0x43,
	0xd7, 0x61, 0x18, 0xda, 0xf0, 0x9a, 0xd2, 0x41, 0x71, 0x89, 0x14, 0xfb, 0x70, 0x3d, 0xa2, 0xf3,
	0x30, 0xe2, 0xab, 0xe2, 0x95, 0x1c, 0xf0, 0xcb, 0x84, 0x0f, 0x33, 0x43, 0x8a, 0x58, 0xc5, 0x59,
	0xe0, 0x69, 0x97, 0x1d, 0xbc, 0xce, 0x0a, 0x66, 0x48, 0x31, 0x40, 0x58, 0xdf, 0x60, 0x45, 0x6a,
	0xc4, 0x73, 0x0e, 0x86, 0x55, 0x14, 0x6c, 0xaa, 0xc8, 0x65, 0x13, 0x57, 0xc8, 0x00, 0x84, 0xa0,
	0x60, 0x16, 0x2a, 0xae, 0x89, 0x78, 0x73, 0x8b, 0xaf, 0x07, 0x67, 0x60, 0x09, 0xc6, 0xb9, 0x40,
	0xb5, 0x55, 0xe4, 0xa0, 0x78, 0x8b, 0x14, 0x63, 0x06, 0x46, 0xaf, 0xa1, 0x65, 0xaa, 0x7d, 0xe9,
	0x22, 0x79, 0x9b, 0x5f, 0x83, 0x10, 0x0a, 0x65, 0x43, 0x46, 0xcd, 0x93, 0x6e, 0x86, 0x77, 0x38,
	0x94, 0xcc, 0xa0, 0x62, 0x01, 0x46, 0x43, 0x2f, 0x49, 0x4f, 0x7a, 0x81, 0x53, 0x3a, 0xde, 0x25,
	0xc7, 0x48, 0x09, 0x51, 0x44, 0xb2, 0x68, 0x10, 0xcd, 0x7b, 0x1c, 0x11, 0x03, 0xa3, 0xab, 0x97,
	0x6a, 0xaf, 0x11, 0xc8, 0xfa, 0x20, 0xb6, 0xf7, 0xf9, 0xea, 0x15, 0xec, 0xb2, 0x69, 0x9c, 0x85,
	0x4a, 0xda, 0x3e, 0xeb, 0xa4, 0xf9, 0x80, 0x33, 0x9d, 0x03, 0x08, 0x3f, 0x00, 0xd7, 0xf7, 0x6d,
	0x13, 0x0e, 0xb2, 0x0f, 0x49, 0x36, 0xdd, 0xa7, 0x55, 0x50, 0x49, 0x18, 0x54, 0xf9, 0x11, 0x97,
	0x04, 0xd9, 0xe3, 0x5a, 0x81, 0xa9, 0x2c, 0x4a, 0xbd, 0x8d, 0xc1, 0xa2, 0xf6, 0x31, 0x47, 0xad,
	0x60, 0xbb, 0xa2, 0x76, 0x02, 0xa6, 0xc9, 0x38, 0x58, 0x5e, 0x3f, 0xe1, 0xc2, 0x5a, 0xd0, 0x6b,
	0xdd, 0xd9, 0x7d, 0x08, 0x0e, 0x94, 0xe1, 0x3c, 0xa3, 0x65, 0x94, 0x22, 0x53, 0x0f, 0xbd, 0xd8,
	0xc1, 0x7c, 0x95, 0xcc, 0x5c, 0xf1, 0x17, 0x4b, 0xc1, 0xb2, 0x17, 0xa3, 0xfc, 0x7e, 0xd8, 0xcf,
	0xf2, 0x2c, 0x4a, 0x64, 0x53, 0xf9, 0x51, 0xfb, 0xac, 0x6c, 0x39, 0xa8, 0x3f, 0xed, 0x49, 0xd5,
	0x9a, 0x81, 0xa3, 0xf9, 0x38, 0x5c, 0x53, 0xce, 0x2a, 0xf5, 0x76, 0x18, 0xab, 0x44, 0x5b, 0x8c,
	0x9f, 0x71, 0xa6, 0x4a, 0xee, 0x78, 0x8e, 0x89, 0x45, 0x18, 0xcb, 0x7f, 0x74, 0x3d, 0x92, 0x9f,
	0x93, 0x68, 0xb4, 0x43, 0x51, 0xe1, 0x68, 0xaa, 0x30, 0xf6, 0x12, 0x97, 0xfa, 0xf7, 0x05, 0x17,
	0x0e, 0x42, 0xa8, 0x70, 0xe8, 0xcd, 0x58, 0x62, 0xb7, 0x77, 0x30, 0x7c, 0xc9, 0x85, 0x83, 0x19,
	0x52, 0xf0, 0xc0, 0xe0, 0xa0, 0xf8, 0x8a, 0x15, 0xcc, 0xa0, 0xe2, 0x9e, 0x4e, 0xa3, 0x4d, 0xa4,
	0xdf, 0x4e, 0x75, 0xe2, 0xe1, 0x6a, 0x8b, 0xea, 0xeb, 0xad, 0xee, 0x21, 0x6c, 0xd5, 0x40, 0xb1,
	0x12, 0x85, 0x32, 0x4d, 0x3d, 0x5f, 0xe2, 0xc4, 0xe1, 0xb0, 0xb1, 0x6f, 0xb8, 0x12, 0x19, 0x58,
	0x71, 0x3f, 0xc7, 0x7b, 0x66, 0x95, 0xea, 0x4d, 0xbb, 0x44, 0xcb, 0x05, 0xc3, 0xae, 0x47, 0xb7,
	0xc9, 0xd5, 0x3d, 0xaa, 0x88, 0x3b, 0xf1, 0x00, 0x75, 0x0f, 0x14, 0x76, 0xd9, 0xb9, 0xed, 0xf2,
	0x0c, 0x75, 0xcd, 0x13, 0xe2, 0x18, 0x8c, 0x76, 0x0d, 0x13, 0x76, 0xd5, 0x63, 0xa4, 0x1a, 0x31,
	0x67, 0x09, 0x71, 0x08, 0xf6, 0xe0, 0x60, 0x60, 0xc7, 0x1f, 0x27, 0x3c, 0x5f, 0x2e, 0x8e, 0xc0,
	0x10, 0x0f, 0x04, 0x76, 0xf4, 0x09, 0x42, 0x4b, 0x04, 0x71, 0x1e, 0x06, 0xec, 0xf8, 0x93, 0x8c,
	0x33, 0x82, 0xb8, 0x7b, 0x08, 0xbf, 0x7d, 0x7a, 0x0f, 0x15, 0x74, 0x8e, 0xdd, 0x2c, 0xec, 0xa3,
	0x29, 0xc0, 0x4e, 0x3f, 0x45, 0xff, 0x9c, 0x09, 0x71, 0x3b, 0xec, 0x75, 0x0c, 0xf8, 0x33, 0x84,
	0x16, 0xeb, 0xc5, 0x02, 0x0c, 0x1b, 0x9d, 0xdf, 0x8e, 0x3f, 0x4b, 0xb8, 0x49, 0xe1, 0xd6, 0xa9,
	0xf3, 0xdb, 0x05, 0xcf, 0xf1, 0xd6, 0x89, 0xc0, 0xb0, 0x71, 0xd3, 0xb7, 0xd3, 0xcf, 0x73, 0xd4,
	0x19, 0x11, 0x73, 0x50, 0x29, 0x0b, 0xb9, 0x9d, 0x7f, 0x81, 0xf8, 0x0e, 0x83, 0x11, 0x30, 0x1a,
	0x89, 0x5d, 0xf1, 0x22, 0x47, 0xc0, 0xa0, 0xf0, 0x1a, 0xf5, 0x0e, 0x07, 0x76, 0xd3, 0x4b, 0x7c,
	0x8d, 0x7a, 0x66, 0x03, 0xcc, 0x66, 0x5e, 0x4f, 0xed, 0x8a, 0x97, 0x39, 0x9b, 0xf9, 0x7a, 0xdc,
	0x46, 0x6f, 0xb7, 0xb5, 0x3b, 0x5e, 0xe1, 0x6d, 0xf4, 0x34, 0x5b, 0xb1, 0x02, 0xd5, 0xdd, 0x9d,
	0xd6, 0xee, 0x7b, 0x95, 0x7c, 0x13, 0xbb, 0x1a, 0xad, 0xb8, 0x0f, 0xa6, 0xfb, 0x77, 0x59, 0xbb,
	0xf5, 0xfc, 0x76, 0xcf, 0xe7, 0x22, 0xb3, 0xc9, 0x8a, 0x13, 0x9d, 0x72, 0x6d, 0x76, 0x58, 0xbb,
	0xf6, 0xc2, 0x76, 0x77, 0xc5, 0x36, 0x1b, 0xac, 0x98, 0x07, 0xe8, 0x34, 0x37, 0xbb, 0xeb, 0x22,
	0xb9, 0x0c, 0x08, 0xaf, 0x06, 0xf5, 0x36, 0x3b, 0x7f, 0x89, 0xaf, 0x06, 0x11, 0x78, 0x35, 0xb8,
	0xad, 0xd9, 0xe9, 0xcb, 0x7c, 0x35, 0x18, 0xc1, 0x93, 0x6d, 0x74, 0x0e, 0xbb, 0xe1, 0x0a, 0x9f,
	0x6c, 0x83, 0x12, 0xb3, 0x30, 0x14, 0x65, 0x41, 0x80, 0x07, 0xb4, 0x7a, 0x63, 0x9f, 0x76, 0x25,
	0x83, 0x16, 0xf3, 0xbf, 0xec, 0xd0, 0x0e, 0x18, 0x10, 0x87, 0x60, 0xaf, 0x0c, 0x1b, 0xb2, 0x65,
	0x23, 0x7f, 0xdd, 0xe1, 0xa2, 0x84, 0xab, 0xc5, 0x1c, 0x40, 0xf1, 0xd1, 0x1e, 0x5f, 0xc5, 0xc6,
	0xfe, 0xb6, 0x53, 0x7c, 0xcb, 0x60, 0x20, 0x1d, 0x41, 0xfe, 0xe2, 0x16, 0xc1, 0x56, 0xb7, 0x20,
	0x7f, 0xeb, 0xc3, 0xb0, 0xef, 0xe1, 0x54, 0x45, 0xda, 0xf3, 0x6d, 0xf4, 0xef, 0x44, 0xf3, 0x7a,
	0x0c, 0x58, 0xa8, 0x12, 0xa9, 0x3d, 0x3f, 0xb5, 0xb1, 0x7f, 0x10, 0x5b, 0x02, 0x08, 0x37, 0xbd,
	0x54, 0xbb, 0xbc, 0xf7, 0x9f, 0x0c, 0x33, 0x80, 0x9b, 0xc6, 0xe7, 0x53, 0x72, 0xd3, 0xc6, 0xfe,
	0xc5, 0x9b, 0xa6, 0xf5, 0xe2, 0x08, 0x54, 0xf0, 0x31, 0xff, 0x56, 0xc4, 0x06, 0xff, 0x4d, 0x70,
	0x87, 0xc0, 0xff, 0x9c, 0xea, 0x96, 0x6e, 0xdb, 0x83, 0xfd, 0x0f, 0x65, 0x9a, 0xd7, 0x8b, 0x79,
	0x18, 0x4e, 0x75, 0xab, 0x95, 0xd1, 0x7c, 0x65, 0xc1, 0xff, 0xdd, 0x29, 0x3f, 0x72, 0x97, 0xcc,
	0xd1, 0x75, 0x98, 0x6c, 0xaa, 0xb0, 0x17, 0x3c, 0x0a, 0x4b, 0x6a, 0x49, 0xad, 0xe4, 0x57, 0xf1,
	0xc1, 0xdb, 0x7c, 0x75, 0x6b, 0x53, 0x85, 0xa1, 0x8a, 0x6a, 0x5e, 0x1c, 0xd7, 0xb4, 0x52, 0x41,
	0xad, 0x11, 0xe6, 0x4b, 0x6b, 0xf1, 0x29, 0xbf, 0xd6, 0xa9, 0x46, 0xb5, 0x72, 0x2e, 0xfe, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x97, 0xb1, 0x98, 0x88, 0x13, 0x14, 0x00, 0x00,
}
