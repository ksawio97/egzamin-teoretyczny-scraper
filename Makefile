OUTPUTFILE=_out
URL=https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/

init:
	go mod download

run: 
	go run ./cmd/egzamin-teoretyczny-scraper/main.go -out $(OUTPUTFILE) -url $(URL)