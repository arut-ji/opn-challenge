# GO-TAMBOON

This repository contains a solution to a proposed problem named [GO-TAMBOON](https://github.com/omise/challenges/tree/challenge-go)

## Build 
```bash
$ make build
```

## Execute

The executable requires two variable to be passed as flags - a public key and a private key for using Omise payment API -
with a path to designated CSV record containing what so called tamboon record.


The example usage is as follows:

```bash
$ ./bin/go-tamboon -publicKey=$OMISE_PUBLIC_KEY -secretKey=$OMISE_SECRET_KEY ./data/fng.1000.csv
```

Additional, an alternative way to run this application is use the following command with credentials
required by Omise payment API predefined as environment variables (see more in `Makefile`)

```bash
$ make run
```