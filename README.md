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
You can change default _out folder and **scrape url** by modifying OUTPUTFILE variable inside Makefile.

Questions in questions.json file will have following parameters:

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