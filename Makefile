.PHONY: entgo
gen:
	go get entgo.io/ent/cmd/ent
	rm -rf ./ent
	go run entgo.io/ent/cmd/ent generate --target ./ent ./schema
