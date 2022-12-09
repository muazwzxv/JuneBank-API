
uri := "root:pass@/user?parseTime=true"

run:
	go run ./cmd

migrations.up:
	goose mysql ${uri} up

migrations.down:
	goose mysql ${uri} down