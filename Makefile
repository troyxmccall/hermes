SHELL := /bin/bash
.DEFAULT_GOAL := ci
TARGETS := $(shell ls scripts) 
.PHONY: run clean $(TARGETS)

$(TARGETS): 
	./scripts/$@

run:
	go run main.go task.go config.go screen.go download.go log.go archive.go \
	run example/00-demo.yml

examples: clean build
	./dist/hermes run example/00-demo.yml
	./dist/hermes run example/01-simple.yml
	./dist/hermes run example/02-simple-and-pretty.yml
	./dist/hermes run example/03-repetitive.yml
	./dist/hermes run example/04-repetitive-parallel.yml
	./dist/hermes run example/05-minimal.yml
	./dist/hermes run example/06-with-errors.yml || true
	# ./dist/hermes run example/07-from-url.yml || true
	./dist/hermes run example/08-complicated.yml || true
	# ./dist/hermes run example/09-stress-and-flow-control.yml
	./dist/hermes run example/10-bad-values.yml || true
	./dist/hermes run example/11-tags.yml --tags some-app1
	./dist/hermes run example/11-tags.yml --only-tags migrate
	./dist/hermes run example/12-share-variables.yml
	./dist/hermes run example/13-single-line.yml || true
	# ./dist/hermes run example/14-sudo.yml
	./dist/hermes run example/15-yaml-includes.yml
	./dist/hermes bundle example/16-bundle-manifest.yml && ./16-bundle-manifest.hermes; rm -f 16-bundle-manifest.hermes

clean:
	rm -f dist/hermes build.log
