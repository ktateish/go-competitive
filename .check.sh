#!/bin/sh

D=$(mktemp -d)
rc=0
for i in testcases/*input
do
	casename=$(basename $i .input)
	expected=$D/expected
	actual=$D/actual
	cp -f testcases/${casename}.output $expected
	./a.out < $i > $actual
	if diff -q $expected $actual > /dev/null; then
		echo "$casename: OK"
	else
		echo "$casename: NG"
		(cd $D && git --no-pager diff --no-index actual expected)
		rc=$((rc + 1))
	fi
	rm -f $actual $expected
done
rmdir $D
if [ $rc -eq 0 ]; then
	echo
	echo "=================="
	echo ">>> OK: $(basename $(pwd))"
	echo "=================="
	pbcopy < gottani.go
else
	pbcopy < /dev/null
	exit 1
fi
