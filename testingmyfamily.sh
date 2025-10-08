#!/bin/bash



curl https://acad.learn2earn.ng/assets/superhero/all.json | \jq --arg id =1 -r '.[] | .connections.relatives | @json' | \sed 's/^"\(.*\)"$/\1/'
