package registry

import (
	"github.com/anghami/protoc-gen-grpc-gateway-ts/data"
	"google.golang.org/protobuf/types/descriptorpb"
)

func (r *Registry) analyseEnumType(
	fileData *data.File,
	packageName, fileName string,
	parents []string,
	enum *descriptorpb.EnumDescriptorProto) {
	packageIdentifier := r.getNameOfPackageLevelIdentifier(parents, enum.GetName())
	fqName := r.getFullQualifiedName(packageName, parents, enum.GetName())
	protoType := descriptorpb.FieldDescriptorProto_TYPE_ENUM
	r.Types[fqName] = &TypeInformation{
		FullyQualifiedName: fqName,
		Package:            packageName,
		File:               fileName,
		PackageIdentifier:  packageIdentifier,
		LocalIdentifier:    enum.GetName(),
		ProtoType:          protoType,
	}

	enumData := data.NewEnum()
	enumData.Name = packageIdentifier

	for _, e := range enum.GetValue() {
		enumData.Values = append(enumData.Values, e.GetName())
	}

	fileData.Enums = append(fileData.Enums, enumData)
}
