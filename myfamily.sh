#!/bin/bash



curl -s https://acad.learn2earn.ng/assets/superhero/all.json | \jq --arg id "$HERO_ID" -r '.[] | select(.id == ($id | tonumber)) | .connections.relatives | @json' | \sed 's/^"\(.*\)"$/\1/'
