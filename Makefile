terraform-provider-macaroons: *.go */*.go go.mod  docs/index.md test
	go build .

test:
	terraform fmt -recursive
	go fmt ./...
	go vet .
	go test ./...

testacc: test
	TF_ACC=1 go test ./...

install: terraform-provider-macaroons
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/macaroons/0.1.0/linux_amd64
	cp $+ ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/macaroons/0.1.0/linux_amd64
	-rm .terraform.lock.hcl
	terraform init

docs/index.md: $(shell find -name "*.go" -or -name "*.tmpl" -or -name "*.tf")
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs
