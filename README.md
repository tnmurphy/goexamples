
# Go Example Programs

Projects on which I improved my understanding of Golang.   

### Wordcounter
  Currently wordcounter is the first and only one.  It finds the top 20 words by frequency in a file. It is not internationalised and will only work for words comprised of ASCII  a-zA-Z.  Binary characters and punctuation are all treated as word separators.

## Getting Started

### Wordcounter
Clone.  
```
   cd wordcounter 
   go build

   ./wordcounter filename.txt 
```
wordcounter will read from stdin if no filenames are specified.

If multiple filenames are specified then all will be read and the results will be a summary.

### Prerequisites

golang and git

## Running the tests


### Break down into end to end tests

Explain what these tests test and why

```
   cd wordcounter 
   go build
   go test

```

## Authors

* **Tim Murphy** 

## License

This project is in the public domain.

