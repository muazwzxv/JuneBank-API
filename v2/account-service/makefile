
run:
	go run ./cmd

migrations.up:
	tern migrate --migrations ./migrations --config ./migrations/tern.conf

migrations.down:
	tern migrate --migrations ./migrations --config ./migrations/tern.conf --destination ${step}

migrations.new:
	cd migrations && tern new $(name) && cd ..