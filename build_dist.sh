#!/usr/bin/env bash

make clean

export GOOS=windows
make
mkdir rel_windows
upx bin/*
mv bin/* rel_windows

export GOOS=freebsd
make
mkdir rel_freebsd
upx bin/*
mv bin/* rel_freebsd

export GOOS=openbsd
make
mkdir rel_openbsd
upx bin/*
mv bin/* rel_openbsd

export GOOS=linux
make
mkdir rel_linux
upx bin/*
mv bin/* rel_linux

mkdir rel
mv rel_* rel
