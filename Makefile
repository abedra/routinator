build:
	go build routinator.go
release:
	gox -osarch="openbsd/amd64"
	gox -osarch="openbsd/386"
	tar cvf templates.tar.gz templates
clean:
	rm -rf out
	rm -f routinator
	rm -f templates.tar.gz
	rm -f routinator_openbsd_386
	rm -f routinator_openbsd_amd64
.PHONY: clean
