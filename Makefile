terraform-provider-macaroons: *.go */*.go go.mod
	go build .

install: terraform-provider-macaroons
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/macaroons/0.1.0/linux_amd64
	cp $+ ~/.terraform.d/plugins/registry.terraform.io/BarnabyShearer/macaroons/0.1.0/linux_amd64
	-rm .terraform.lock.hcl
	terraform init
