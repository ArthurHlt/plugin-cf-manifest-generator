#!/bin/bash

set -e
set -x

OUTDIR=$(dirname $0)/../out
BINARYNAME=plugin-cf-manifest-generator

GOARCH=amd64 GOOS=windows $(dirname $0)/build && cp $OUTDIR/$BINARYNAME $OUTDIR/$BINARYNAME-windows-amd64.exe
GOARCH=386 GOOS=windows $(dirname $0)/build && cp $OUTDIR/$BINARYNAME $OUTDIR/$BINARYNAME-windows-386.exe
GOARCH=amd64 GOOS=linux $(dirname $0)/build  && cp $OUTDIR/$BINARYNAME $OUTDIR/$BINARYNAME-linux-amd64
GOARCH=386 GOOS=linux $(dirname $0)/build  && cp $OUTDIR/$BINARYNAME $OUTDIR/$BINARYNAME-linux-386
GOARCH=amd64 GOOS=darwin $(dirname $0)/build  && cp $OUTDIR/$BINARYNAME $OUTDIR/$BINARYNAME-darwin-amd64
