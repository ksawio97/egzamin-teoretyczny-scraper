# Theoretical exam scraper
This repository is dedicated to scraping www.praktycznyegzamin.pl exam for questions with answers using gocolly.

## Prerequisites

- Go programming language installed on your machine.
- Access to the internet to download dependencies.

## Setup

1. Clone or download the Go web scraper repository.
2. Navigate to the root directory of the project.


## Installation

Run the following command to download the necessary dependencies:
```bash
make init
```
or
```bash
go mod download
```

## Usage

You can run web scraper with following command:
```bash
make run
```

## Makefile Variables for ```make run```
### 1. `OUTPUTFILE`

- **Description**: Defines the output directory where the scraped data will be stored.
- **Example**: `OUTPUTFILE=_out`

### 2. `URL`

- **Description**: Specifies the URL of the website that will be scraped.
- **Example**: `URL=https://www.praktycznyegzamin.pl/inf04/teoria/wszystko/`

### 3. `REMOVETITLEPREFIX`

- **Description**: Determines whether the scraper should remove title prefixes during scraping. 
  - `true`: Title prefixes will be removed.
  - `false`: Title prefixes will not be removed.
- **Example**: `REMOVETITLEPREFIX=true`

### 4. `REMOVEANSWERPREFIX`

- **Description**: Controls whether the scraper should remove answer prefixes during scraping.
  - `true`: Answer prefixes will be removed.
  - `false`: Answer prefixes will not be removed.
- **Example**: `REMOVEANSWERPREFIX=true`

## Output Structure
The scraper will generate the following structure in the specified output directory:
```
_out
│   ├── images
│   │   ├── <image1>.jpg
│   │   ├── <image2>.jpg
│   │   ├── ...
│   ├── questions.json
│   └── videos
│       ├── <video1>.mp4
│       ├── <video2>.mp4
│       ├── ...
```
### Questions data
*questions.json* file will contain list of **[Question](https://github.com/ksawio97/egzamin-teoretyczny-scraper/blob/master/internal/egzamin-teoretyczny-scraper/types.go#L3)** with titles, answers and sometimes images or videos for additional context. 