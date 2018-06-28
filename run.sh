#!/bin/zsh

JAR="maelstrom-0.1.0-standalone.jar"

make && \
	java -jar $JAR test --bin bin/server -n "n1" --log-stderr
