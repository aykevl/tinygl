all: fonts

fonts: \
    roboto/regular16.go \
    roboto/regular20.go \
    roboto/regular24.go \
    roboto/regular28.go \
    roboto/regular32.go \
    roboto/regular36.go \
    roboto/regular40.go \
    roboto/regular44.go \
    roboto/regular48.go

roboto/regular16.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 16 -o roboto/regular16.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular16.go

roboto/regular20.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 20 -o roboto/regular20.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular20.go

roboto/regular24.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 24 -o roboto/regular24.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular24.go

roboto/regular28.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 28 -o roboto/regular28.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular28.go

roboto/regular32.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 32 -o roboto/regular32.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular32.go

roboto/regular36.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 36 -o roboto/regular36.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular36.go

roboto/regular40.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 40 -o roboto/regular40.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular40.go

roboto/regular44.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 44 -o roboto/regular44.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular44.go

roboto/regular48.go: generate/main.go
	go run ./generate -font roboto/Roboto-Regular.ttf -size 48 -o roboto/regular48.go -package=roboto $(FONT_FLAGS)
	@go fmt roboto/regular48.go
