#!/bin/sh
set -e

cp debian/tests/testgetifaddr.sh debian/tests/testupnppermissions.sh .
make -f Makefile.linux check
make -f Makefile.linux clean
rm -rf testgetifaddr.sh testupnppermissions.sh
