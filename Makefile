all: hello count

.PHONY: hello
hello:
	@./scripts/build.sh hello

.PHONY: count
count:
	@./scripts/build.sh count

.PHONY: package-hello
package-hello: cleanpackage
	@scripts/package.sh hello Linux   linux   amd64

.PHONY: package-count
package-count: cleanpackage
	@scripts/package.sh count Linux   linux   amd64

package: cleanpackage package-hello package-count

cleanpackage:
	@rm -rf packages/
