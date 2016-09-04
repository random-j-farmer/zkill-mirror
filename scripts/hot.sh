#! /bin/bash
http http://localhost:${ZKM_PORT:-8080}/api/hot/${1:-1h}/
