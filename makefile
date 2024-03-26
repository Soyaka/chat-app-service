run:
	@go run main.go
push:
	git add .
	git commit -m "$(m)"
	git push origin master
