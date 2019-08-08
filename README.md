# Prometheus workshop

#### Example application: Read Evaluate Print Loop

Its a text processing application. 
Following is the schematic of application

```
loop;
    accept text input from STDIN;
    process the text data;
    print the processed output to STDOUT;
```

Each loop we consider it as `a text processing Job` 

The application is refereed from https://github.com/census-instrumentation

#### Lets see it in action
```
export GO111MODULE=on
go build -o repl main.go
```

![hello world](https://i.ibb.co/F6vgpZG/hello-repl.png)

#### Whats the problem then?

We want `metrics` about 

1. How many `jobs` it is processing? 
2. How much `time` it is taking to process one `job`?
3. In each `job` how much `amount of data` it is processing?