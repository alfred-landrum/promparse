#!/usr/bin/env bash
# Input is one or .json Grafana dashboard files.

exec jq -r ".panels[].targets[0].expr | select(. == null | not) | gsub(\"[\\n\\t]\"; \"\")" $*
