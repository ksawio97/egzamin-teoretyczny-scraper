OUTPUTFILE=_out
URL=https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/
# true or false values
REMOVETITLEPREFIX=true
REMOVEANSWERPREFIX=true

init:
	go mod download

run: 
	go run ./cmd/egzamin-teoretyczny-scraper/main.go \
	-out $(OUTPUTFILE) \
	-url $(URL) \
	$(if $(filter true,$(REMOVETITLEPREFIX)),-rmtitlep,) $(if $(filter true,$(REMOVEANSWERPREFIX)),-rmanswerp,)