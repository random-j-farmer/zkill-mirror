#! /bin/bash
http http://localhost:${ZKM_PORT:-8080}/api/activity/${1:-1h}/
