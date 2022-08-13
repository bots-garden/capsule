#!/bin/bash
capsule \
   -mode=registry \
   -files="${PWD}/registry/functions" \
   -httpPort=4999
