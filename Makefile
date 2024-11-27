run: Run

Run:	
	go run cmd/app/main.go   

test: Test

Test:
	go test ./controllers -v    
	go test ./models -v
	go test ./views -v