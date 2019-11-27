#!/bin/sh -e

# usage: <action>_config <file> <key> <value>

replace_config () {
	if [ -s "${1}" ]; then
		sed -ri '
x
s|^$||
t find
x
b
: find
x
s|^\s*('"${2}"'=).*|\1'"${3}"'|
t end
s|^#[ \t#]*('"${2}"'=['\'\"']?'"${3}"'['\'\"']?)\s*$|\1|
t end
s|^#[ \t#]*('"${2}"'=)\s*$|\1'"${3}"'|
t end
$ a'"${2}"'='"${3}"'
b
: end
h
		' "${1}"
	else
		echo "${2}=${3}" > "${1}"
	fi
}

append_config () {
	file="${1}"
	key="${2}"
	while [ "$#" -ge 3 ]; do
		if [ -s "${file}" ]; then
			sed -ri '
x
s|^$||
t find
x
b
: find
x
s|^\s*'"${key}"'=['\'\"']?'"${3}"'['\'\"']?\s*$|&|
t end
s|^#[ \t#]*('"${key}"'=['\'\"']?'"${3}"'['\'\"']?)\s*$|\1|
t end
s|^#[ \t#]*('"${key}"'=)\s*$|\1'"${3}"'|
t end
$ a'"${key}"'='"${3}"'
b
: end
h
			' "${file}"
		else
			echo "${key}=${3}" > "${file}"
		fi
		shift
	done
}

remove_config () {
	file="${1}"
	key="${2}"
	while [ "$#" -ge 3 ]; do
		sed -ri 's|^\s*'"${key}"'=['\'\"']?'"${3}"'['\'\"']?\s*$|#&|' "${file}"
		shift
	done
}

remove_config_all () {
	sed -ri 's|^\s*'"${2}"'=.*$|#&|' "${1}"
}

read_config () {
	sed -rn '
s|^\s*'"${2}"'='\''([^'\'']+)'\''\s*$|\1|g
t hold
s|^\s*'"${2}"'='\"'([^'\"']+)'\"'\s*$|\1|g
t hold
s|^\s*'"${2}"'=(\S+)\s*$|\1|
t hold
b
: hold
p
	' "${1}"
}

f=a

echo -n > $f

echo 'append a=b c d/d'
append_config $f a b c d/d
cat $f
echo '---'

echo 'append a=c'
append_config $f a c
cat $f
echo '---'

echo 'replace a=e'
replace_config $f a e
cat $f
echo '---'

echo 'write a=f/f'
echo '   a=f/f' >> $f
cat $f
echo '---'

echo 'remove a=f/f'
remove_config $f a f/f
cat $f
echo '---'

echo 'append a=f/f'
append_config $f a f/f
cat $f
echo '---'

echo 'remove a=c f/f'
remove_config $f a c f/f
cat $f
echo '---'

echo 'write a='\'f g\'
echo '   a='\'f g\' >> $f
cat $f
echo '---'

echo 'read a'
read_config $f a
echo '---'

echo 'remove a='\'f g\'
remove_config $f a \'f' 'g\'
cat $f
echo '---'

echo 'remove_all a'
remove_config_all $f a
cat $f
echo '---'

echo 'append a=g'
append_config $f a g
cat $f
echo '---'
