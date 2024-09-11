#!/bin/sh

# RNG for cow chance
cow=$((1 + $RANDOM % $CHANCE))
if [ $cow != 1 ]; then
	# Use default cow
	mootd=$(fortune | cowsay)
else
	# Use random cow from /usr/share/cows/
	random_cow=$(ls /usr/share/cows/ | shuf -n1)
	mootd=$(fortune | cowsay -f "/usr/share/cows/$random_cow")
fi

# Write motd
echo "$mootd" > /srv/mootd
