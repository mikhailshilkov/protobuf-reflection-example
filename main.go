package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func registerProtoFile(srcDir string, filename string) error {
	// First, convert the .proto file to a file descriptor set.
	tmpFile := path.Join(srcDir, filename + ".pb")
	cmd := exec.Command("protoc",
		"--include_source_info",
		"--descriptor_set_out=" + tmpFile,
		"--proto_path="+srcDir,
		path.Join(srcDir, filename))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return errors.Wrapf(err, "protoc")
	}

	defer os.Remove(tmpFile)

	// Now load that temporary file as a file descriptor set protobuf.
	protoFile, err := ioutil.ReadFile(tmpFile)
	if err != nil {
		return errors.Wrapf(err, "read tmp file")
	}

	pbSet := new(descriptorpb.FileDescriptorSet)
	if err := proto.Unmarshal(protoFile, pbSet); err != nil {
		return errors.Wrapf(err, "unmarshal")
	}

	// We know protoc was invoked with a single .proto file.
	pb := pbSet.GetFile()[0]

	// Initialize the file descriptor object.
	fd, err := protodesc.NewFile(pb, protoregistry.GlobalFiles)
	if err != nil {
		return errors.Wrapf(err, "NewFile")
	}

	// and finally register it.
	return protoregistry.GlobalFiles.RegisterFile(fd)
}

func main() {
	files := []string{
		"google/protobuf/any.proto",
		"google/protobuf/duration.proto",
		"google/protobuf/empty.proto",
		"google/protobuf/field_mask.proto",
		"google/protobuf/timestamp.proto",
		"google/rpc/status.proto",
		"google/type/expr.proto",
		"google/api/http.proto",
		"google/api/annotations.proto",
		"google/api/client.proto",
		"google/api/field_behavior.proto",
		"google/api/resource.proto",
		"google/iam/v1/options.proto",
		"google/iam/v1/policy.proto",
		"google/iam/v1/iam_policy.proto",
		"google/longrunning/operations.proto",
		"google/cloud/functions/v1/functions.proto",
	}
	for _, f := range files {
		err := registerProtoFile("./proto", f)
		if err != nil {
			panic(errors.Wrapf(err, f))
		}
	}
	var services []protoreflect.ServiceDescriptors
	var sourceLocations protoreflect.SourceLocations
	protoregistry.GlobalFiles.RangeFilesByPackage("google.cloud.functions.v1", func(descriptor protoreflect.FileDescriptor) bool {
		services = append(services, descriptor.Services())
		sourceLocations = descriptor.SourceLocations()
		return true
	})
	service := services[0].Get(0)
	serviceDesc := sourceLocations.ByDescriptor(service)
	fmt.Printf("\n=====\n%+v\n=====\n\n", strings.TrimSpace(serviceDesc.LeadingComments))
	for i := 0; i < service.Methods().Len(); i++ {
		method := service.Methods().Get(i)
		methodDesc := sourceLocations.ByDescriptor(method)
		fmt.Printf("%s:\n%s\n", method.Name(), methodDesc.LeadingComments)
	}
}
