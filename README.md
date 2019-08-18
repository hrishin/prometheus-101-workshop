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
1. Total number of `jobs` is been processed? 
2. How much `time` it is taking to process an one `job`?
3. In each `job` how much `amount of data` it is processing?

#### Step-1 : Job Application
Run the application
```
go build -o repl main.go
```

#### Step-2 : Job counter

Run the application, expose metrics and collect metrics

```
go build -o repl main.go

./repl
```

start prometheus in new terminal
```
wget https://github.com/prometheus/prometheus/releases/download/v2.11.1/prometheus-2.11.1.linux-386.tar.gz

tar -xzvf prometheus-2.11.1.linux-386.tar.gz 

sudo cp prometheus-2.11.1.linux-386/prometheus /usr/bin/prometheus

prometheus --config.file=promconfig.yaml
```
