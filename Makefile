install:
	pulumi plugin install --reinstall --server https://downloads.spacelift.io/pulumi-plugins resource spacelift 0.0.0
	go mod download

