all: hello count

hello:
	@./scripts/build.sh hello

count:
	@./scripts/build.sh count

package-hello: cleanpackage
	@scripts/package.sh hello Linux   linux   amd64

package-count: cleanpackage
	@scripts/package.sh count Linux   linux   amd64

package: cleanpackage package-hello package-count

cleanpackage:
	@rm -rf packages/
