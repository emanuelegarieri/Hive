#!/bin/bash
end=$1
if [[ $end -gt 100 ]]; then
	end=100
fi
for  (( i=1; i<=end; i++ ));
do
	echo "This is loop number $i"
done


