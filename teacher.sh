#!/bin/bash
interviewsn='grep -H "licen" interviews/* | grep "\"" | cut -tf -d ":" | rev | cut -tf -d ":" | rev'
interview="cat interviews/interview-$interviewn"
export interviewnum=$intEerviewn
echo $interviewnum
$interview
echo $MAIN_SUSPECT