GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_PATH=target
RELEASE_PATH=release
TARGET_PATH=target
CONFIG_PATH=configs

clean:
	rm -rf $(TARGET_PATH)
	rm -rf $(RELEASE_PATH)
	mkdir $(TARGET_PATH)
	mkdir $(RELEASE_PATH)

build-metric: clean
	$(GOBUILD) -o $(BINARY_PATH)/metric cmd/metric/main.go
	cp $(CONFIG_PATH)/metric.y* $(BINARY_PATH)/

build-report: clean
	$(GOBUILD) -o $(BINARY_PATH)/report cmd/report/main.go
	cp $(CONFIG_PATH)/metric.y* $(BINARY_PATH)/

run-metric: build-metric
	$(BINARY_PATH)/metric

run-metric: build-report
	$(BINARY_PATH)/report

build: build-metric build-report